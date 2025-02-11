package services

import (
	"context"
	"database/sql"
	"io"
	"os"
	"registration/api/generated/supporters"
	validator "registration/internal/validators"
	models "registration/models/generated"
	"strconv"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"

	"cloud.google.com/go/storage"
)

type SupporterService interface {
	ValidateSignUp(ctx context.Context, request *supporters.PostAuthValidateSignUpMultipartRequestBody) error
	SignUp(ctx context.Context, requestParams supporters.PostAuthSignUpMultipartRequestBody) error
}

type supporterService struct {
	db *sql.DB
}

func NewSupporterService(db *sql.DB) SupporterService {
	return &supporterService{db}
}

func (ss *supporterService) ValidateSignUp(ctx context.Context, request *supporters.PostAuthValidateSignUpMultipartRequestBody) error {
	return validator.ValidateSignUpSupporter(request)
}

func (ss *supporterService) SignUp(ctx context.Context, requestParams supporters.PostAuthSignUpMultipartRequestBody) error {
	supporter := models.Supporter{}
	supporter.FirstName = requestParams.FirstName
	supporter.LastName = requestParams.LastName
	supporter.Email = requestParams.Email
	if requestParams.Birthday != nil {
		supporter.Birthday = null.Time{Time: requestParams.Birthday.Time, Valid: true}
	}
	supporter.FrontIdentification = ""
	supporter.BackIdentification = ""

	// NOTE: パスワードをハッシュ化の上、Create処理
	hashedPassword, err := ss.encryptPassword(requestParams.Password)
	if err != nil {
		return err
	}
	supporter.Password = hashedPassword
	createErr := supporter.Insert(ctx, ss.db, boil.Infer())
	if createErr != nil {
		return createErr
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	if requestParams.FrontIdentification == nil && requestParams.BackIdentification == nil {
		return nil
	}

	supporter.Reload(ctx, ss.db)
	supporterID := strconv.Itoa(supporter.ID)

	bucket := client.Bucket(os.Getenv("STORAGE_BUCKET_NAME"))

	if requestParams.FrontIdentification != nil {
		frontIdentificationPath := "supporters/"+supporterID+"/"+requestParams.FrontIdentification.Filename()
		reader, _ :=requestParams.FrontIdentification.Reader()
		uploadFrontIdentificationErr := ss.uploadIdentification(ctx, bucket, frontIdentificationPath, reader)
		if uploadFrontIdentificationErr != nil {
			return uploadFrontIdentificationErr
		}
		supporter.FrontIdentification = frontIdentificationPath
	}
	if requestParams.BackIdentification != nil {
		backIdentificationPath := "supporters/"+supporterID+"/"+requestParams.BackIdentification.Filename()
		backIdentificationReader, _ :=requestParams.BackIdentification.Reader()
		uploadBackIdentificationErr := ss.uploadIdentification(ctx, bucket, backIdentificationPath, backIdentificationReader)
		if uploadBackIdentificationErr != nil {
			return uploadBackIdentificationErr
		}
		supporter.BackIdentification = backIdentificationPath
	}
	_, updateIdenfiticationErr := supporter.Update(ctx, ss.db, boil.Infer())
	return updateIdenfiticationErr
}

// NOTE: パスワードの文字列をハッシュ化する
func (ss *supporterService) encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (ss *supporterService) uploadIdentification(ctx context.Context, bucket *storage.BucketHandle, path string, reader io.Reader) error {
	obj := bucket.Object(path)
	writer := obj.NewWriter(ctx)
	// NOTE: ファイルをCloud Storageにコピー
	if _, err := io.Copy(writer, reader); err != nil {
		return err
	}
	// NOTE: Writerを閉じて完了
	if err := writer.Close(); err != nil {
		return err
	}
	return nil
}

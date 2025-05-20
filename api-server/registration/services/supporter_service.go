package registrationservices

import (
	registrationapi "apps/api/registration"
	models "apps/models"
	registrationvalidator "apps/registration/validators"
	"context"
	"io"
	"os"
	"strconv"

	"github.com/volatiletech/null/v8"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"cloud.google.com/go/storage"
)

type SupporterService interface {
	ValidateSignUp(request *registrationapi.PostSupporterValidateSignUpMultipartRequestBody) error
	SignUp(ctx context.Context, requestParams registrationapi.PostSupporterSignUpMultipartRequestBody) error
}

type supporterService struct {
	db *gorm.DB
}

func NewSupporterService(db *gorm.DB) SupporterService {
	return &supporterService{db}
}

func (ss *supporterService) ValidateSignUp(request *registrationapi.PostSupporterValidateSignUpMultipartRequestBody) error {
	return registrationvalidator.ValidateSignUpSupporter(request)
}

func (ss *supporterService) SignUp(ctx context.Context, requestParams registrationapi.PostSupporterSignUpMultipartRequestBody) error {
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
	ss.db.Create(&supporter)

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	if requestParams.FrontIdentification == nil && requestParams.BackIdentification == nil {
		return nil
	}

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
	return ss.db.Save(&supporter).Error
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

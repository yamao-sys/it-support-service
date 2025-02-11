package services

import (
	"context"
	"database/sql"
	"io"
	"os"
	"registration/api/generated/companies"
	validator "registration/internal/validators"
	models "registration/models/generated"
	"strconv"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"

	"cloud.google.com/go/storage"
)

type CompanyService interface {
	ValidateSignUp(ctx context.Context, request *companies.PostAuthValidateSignUpMultipartRequestBody) error
	SignUp(ctx context.Context, requestParams companies.PostAuthSignUpMultipartRequestBody) error
}

type companyService struct {
	db *sql.DB
}

func NewCompanyService(db *sql.DB) CompanyService {
	return &companyService{db}
}

func (cs *companyService) ValidateSignUp(ctx context.Context, request *companies.PostAuthValidateSignUpMultipartRequestBody) error {
	return validator.ValidateSignUpCompany(request)
}

func (cs *companyService) SignUp(ctx context.Context, requestParams companies.PostAuthSignUpMultipartRequestBody) error {
	company := models.Company{}
	company.Name = requestParams.Name
	company.Email = requestParams.Email
	company.FinalTaxReturn = ""

	// NOTE: パスワードをハッシュ化の上、Create処理
	hashedPassword, err := cs.encryptPassword(requestParams.Password)
	if err != nil {
		return err
	}
	company.Password = hashedPassword
	createErr := company.Insert(ctx, cs.db, boil.Infer())
	if createErr != nil {
		return createErr
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	if requestParams.FinalTaxReturn == nil {
		return nil
	}

	company.Reload(ctx, cs.db)
	companyID := strconv.Itoa(company.ID)

	bucket := client.Bucket(os.Getenv("STORAGE_BUCKET_NAME"))

	finalTaxReturnPath := "companies/"+companyID+"/"+requestParams.FinalTaxReturn.Filename()
	reader, _ :=requestParams.FinalTaxReturn.Reader()
	uploadFinalTaxReturnErr := cs.upload(ctx, bucket, finalTaxReturnPath, reader)
	if uploadFinalTaxReturnErr != nil {
		return uploadFinalTaxReturnErr
	}
	company.FinalTaxReturn = finalTaxReturnPath
	_, updateIdenfiticationErr := company.Update(ctx, cs.db, boil.Infer())
	return updateIdenfiticationErr
}

// NOTE: パスワードの文字列をハッシュ化する
func (cs *companyService) encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (cs *companyService) upload(ctx context.Context, bucket *storage.BucketHandle, path string, reader io.Reader) error {
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

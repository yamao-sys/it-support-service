package registrationservices

import (
	registrationapi "apps/api/registration"
	models "apps/models"
	registrationvalidator "apps/registration/validators"
	"context"
	"io"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"cloud.google.com/go/storage"
)

type CompanyService interface {
	ValidateSignUp(request *registrationapi.PostCompanyValidateSignUpMultipartRequestBody) error
	SignUp(ctx context.Context, requestParams registrationapi.PostCompanySignUpMultipartRequestBody) error
}

type companyService struct {
	db *gorm.DB
}

func NewCompanyService(db *gorm.DB) CompanyService {
	return &companyService{db}
}

func (cs *companyService) ValidateSignUp(request *registrationapi.PostCompanyValidateSignUpMultipartRequestBody) error {
	return registrationvalidator.ValidateSignUpCompany(request)
}

func (cs *companyService) SignUp(ctx context.Context, requestParams registrationapi.PostCompanySignUpMultipartRequestBody) error {
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
	cs.db.Create(&company)

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	if requestParams.FinalTaxReturn == nil {
		return nil
	}

	companyID := strconv.Itoa(company.ID)

	bucket := client.Bucket(os.Getenv("STORAGE_BUCKET_NAME"))

	finalTaxReturnPath := "companies/"+companyID+"/"+requestParams.FinalTaxReturn.Filename()
	reader, _ :=requestParams.FinalTaxReturn.Reader()
	uploadFinalTaxReturnErr := cs.upload(ctx, bucket, finalTaxReturnPath, reader)
	if uploadFinalTaxReturnErr != nil {
		return uploadFinalTaxReturnErr
	}
	company.FinalTaxReturn = finalTaxReturnPath
	return cs.db.Save(&company).Error
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

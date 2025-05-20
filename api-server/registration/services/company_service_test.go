package registrationservices

import (
	registrationapi "apps/api/registration"
	models "apps/models"
	"bytes"
	"strconv"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type TestCompanyServiceSuite struct {
	WithDBSuite
}

var testCompanyService CompanyService

func (s *TestCompanyServiceSuite) SetupTest() {
	s.SetDBCon()

	testCompanyService = NewCompanyService(DBCon)
}

func (s *TestCompanyServiceSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestCompanyServiceSuite) TestValidateSignUp_SuccessRequiredFields() {
	requestParams := registrationapi.PostCompanyValidateSignUpMultipartRequestBody{
		Name: "name",
		Email: "test@example.com",
		Password: "Password",
	}

	result := testCompanyService.ValidateSignUp(&requestParams)

	assert.Nil(s.T(), result)
}

func (s *TestCompanyServiceSuite) TestValidateSignUp_ValidationErrorRequiredFields() {
	requestParams := registrationapi.PostCompanyValidateSignUpMultipartRequestBody{
		Name: "",
		Email: "",
		Password: "",
	}

	result := testCompanyService.ValidateSignUp(&requestParams)

	assert.NotNil(s.T(), result)
	if errors, ok := result.(validation.Errors); ok {
		for field, err := range errors {
			message := err.Error()
			switch field {
			case "name":
				assert.Equal(s.T(), "企業名は必須入力です。", message)
			case "email":
				assert.Equal(s.T(), "Emailは必須入力です。", message)
			case "password":
				assert.Equal(s.T(), "パスワードは必須入力です。", message)
			}
		}
	}
}

func (s *TestCompanyServiceSuite) TestValidateSignUp_SuccessWithOptionalFields() {
	pngSignature := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	jpgSignature := []byte{0xFF, 0xD8, 0xFF, 0xE0}
	// NOTE:データを格納
	var pngBuf, jpgBuf bytes.Buffer
	pngBuf.Write(pngSignature)
	jpgBuf.Write(jpgSignature)
	
	var finalTaxReturn openapi_types.File
	finalTaxReturn.InitFromBytes(pngBuf.Bytes(), "finalTaxReturn.png")

	requestParams := registrationapi.PostCompanyValidateSignUpMultipartRequestBody{
		Name: "name",
		Email: "test@example.com",
		Password: "Password",
		FinalTaxReturn: &finalTaxReturn,
	}

	result := testCompanyService.ValidateSignUp(&requestParams)

	assert.Nil(s.T(), result)
}

func (s *TestCompanyServiceSuite) TestValidateSignUp_ValidationErrorWithOptionalFields() {
	gifSignature := []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61}
	// NOTE:データを格納
	var gifBuf bytes.Buffer
	gifBuf.Write(gifSignature)
	
	var finalTaxReturn openapi_types.File
	finalTaxReturn.InitFromBytes(gifBuf.Bytes(), "finalTaxReturn.gif")

	requestParams := registrationapi.PostCompanyValidateSignUpMultipartRequestBody{
		Name: "name",
		Email: "test@example.com",
		Password: "Password",
		FinalTaxReturn: &finalTaxReturn,
	}

	result := testCompanyService.ValidateSignUp(&requestParams)

	assert.NotNil(s.T(), result)
	if errors, ok := result.(validation.Errors); ok {
		for field, err := range errors {
			message := err.Error()
			switch field {
			case "finalTaxReturn":
				assert.Equal(s.T(), "確定申告書の拡張子はwebp, png, jpegのいずれかでお願いします。", message)
			}
		}
	}
}

func (s *TestCompanyServiceSuite) TestSignUp_SuccessRequiredFields() {
	requestParams := registrationapi.PostCompanySignUpMultipartRequestBody{
		Name: "name",
		Email: "test@example.com",
		Password: "Password",
	}

	result := testCompanyService.SignUp(ctx, requestParams)

	assert.Nil(s.T(), result)

	// NOTE: 企業が作成されていることを確認
	var company models.Company
	DBCon.Where("email = ?", "test@example.com").Take(&company)
	assert.Equal(s.T(), "name", company.Name)
	assert.Equal(s.T(), "", company.FinalTaxReturn)
}

func (s *TestCompanyServiceSuite) TestSignUp_SuccessWithOptionalFields() {
	pngSignature := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	jpgSignature := []byte{0xFF, 0xD8, 0xFF, 0xE0}
	// NOTE:データを格納
	var pngBuf, jpgBuf bytes.Buffer
	pngBuf.Write(pngSignature)
	jpgBuf.Write(jpgSignature)
	
	var finalTaxReturn openapi_types.File
	finalTaxReturn.InitFromBytes(pngBuf.Bytes(), "finalTaxReturn.png")

	requestParams := registrationapi.PostCompanySignUpMultipartRequestBody{
		Name: "name",
		Email: "test@example.com",
		Password: "Password",
		FinalTaxReturn: &finalTaxReturn,
	}

	result := testCompanyService.SignUp(ctx, requestParams)

	assert.Nil(s.T(), result)

	// NOTE: が作成されていることを確認
	var company models.Company
	DBCon.Where("email = ?", "test@example.com").Take(&company)
	assert.Equal(s.T(), "name", company.Name)
	id := strconv.Itoa(company.ID)
	assert.Equal(s.T(), "companies/"+id+"/finalTaxReturn.png", company.FinalTaxReturn)
}

func TestCompanyService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestCompanyServiceSuite))
}

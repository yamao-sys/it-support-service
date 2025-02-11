package services

import (
	"bytes"
	"registration/api/generated/companies"
	models "registration/models/generated"
	"strconv"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

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
	requestParams := companies.PostAuthValidateSignUpMultipartRequestBody{
		Name: "name",
		Email: "test@example.com",
		Password: "Password",
	}

	result := testCompanyService.ValidateSignUp(ctx, &requestParams)

	assert.Nil(s.T(), result)
}

func (s *TestCompanyServiceSuite) TestValidateSignUp_ValidationErrorRequiredFields() {
	requestParams := companies.PostAuthValidateSignUpMultipartRequestBody{
		Name: "",
		Email: "",
		Password: "",
	}

	result := testCompanyService.ValidateSignUp(ctx, &requestParams)

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

	requestParams := companies.PostAuthValidateSignUpMultipartRequestBody{
		Name: "name",
		Email: "test@example.com",
		Password: "Password",
		FinalTaxReturn: &finalTaxReturn,
	}

	result := testCompanyService.ValidateSignUp(ctx, &requestParams)

	assert.Nil(s.T(), result)
}

func (s *TestCompanyServiceSuite) TestValidateSignUp_ValidationErrorWithOptionalFields() {
	gifSignature := []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61}
	// NOTE:データを格納
	var gifBuf bytes.Buffer
	gifBuf.Write(gifSignature)
	
	var finalTaxReturn openapi_types.File
	finalTaxReturn.InitFromBytes(gifBuf.Bytes(), "finalTaxReturn.gif")

	requestParams := companies.PostAuthValidateSignUpMultipartRequestBody{
		Name: "name",
		Email: "test@example.com",
		Password: "Password",
		FinalTaxReturn: &finalTaxReturn,
	}

	result := testCompanyService.ValidateSignUp(ctx, &requestParams)

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
	requestParams := companies.PostAuthSignUpMultipartRequestBody{
		Name: "name",
		Email: "test@example.com",
		Password: "Password",
	}

	result := testCompanyService.SignUp(ctx, requestParams)

	assert.Nil(s.T(), result)

	// NOTE: 企業が作成されていることを確認
	company, err := models.Companies(
		qm.Where("email = ?", "test@example.com"),
	).One(ctx, DBCon)
	if err != nil {
		s.T().Fatalf("failed to create company %v", err)
	}
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

	requestParams := companies.PostAuthSignUpMultipartRequestBody{
		Name: "name",
		Email: "test@example.com",
		Password: "Password",
		FinalTaxReturn: &finalTaxReturn,
	}

	result := testCompanyService.SignUp(ctx, requestParams)

	assert.Nil(s.T(), result)

	// NOTE: が作成されていることを確認
	company, err := models.Companies(
		qm.Where("email = ?", "test@example.com"),
	).One(ctx, DBCon)
	if err != nil {
		s.T().Fatalf("failed to create company %v", err)
	}
	id := strconv.Itoa(company.ID)
	assert.Equal(s.T(), "companies/"+id+"/finalTaxReturn.png", company.FinalTaxReturn)
}

func TestCompanyService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestCompanyServiceSuite))
}

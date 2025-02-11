package services

import (
	"bytes"
	supporters "registration/api/generated"
	models "registration/models/generated"
	"strconv"
	"testing"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type TestSupporterServiceSuite struct {
	WithDBSuite
}

var testSupporterService SupporterService

func (s *TestSupporterServiceSuite) SetupTest() {
	s.SetDBCon()

	testSupporterService = NewSupporterService(DBCon)
}

func (s *TestSupporterServiceSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestSupporterServiceSuite) TestValidateSignUp_SuccessRequiredFields() {
	requestParams := supporters.PostAuthValidateSignUpMultipartRequestBody{
		FirstName: "first_name",
		LastName: "last_name",
		Email: "test@example.com",
		Password: "Password",
	}

	result := testSupporterService.ValidateSignUp(ctx, &requestParams)

	assert.Nil(s.T(), result)
}

func (s *TestSupporterServiceSuite) TestValidateSignUp_ValidationErrorRequiredFields() {
	requestParams := supporters.PostAuthValidateSignUpMultipartRequestBody{
		FirstName: "",
		LastName: "",
		Email: "",
		Password: "",
	}

	result := testSupporterService.ValidateSignUp(ctx, &requestParams)

	assert.NotNil(s.T(), result)
	if errors, ok := result.(validation.Errors); ok {
		for field, err := range errors {
			message := err.Error()
			switch field {
			case "firstName":
				assert.Equal(s.T(), "名は必須入力です。", message)
			case "lastName":
				assert.Equal(s.T(), "姓は必須入力です。", message)
			case "email":
				assert.Equal(s.T(), "Emailは必須入力です。", message)
			case "password":
				assert.Equal(s.T(), "パスワードは必須入力です。", message)
			}
		}
	}
}

func (s *TestSupporterServiceSuite) TestValidateSignUp_SuccessWithOptionalFields() {
	pngSignature := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	jpgSignature := []byte{0xFF, 0xD8, 0xFF, 0xE0}
	// NOTE:データを格納
	var pngBuf, jpgBuf bytes.Buffer
	pngBuf.Write(pngSignature)
	jpgBuf.Write(jpgSignature)
	
	var frontIdentificationFile, backIdentificationFile openapi_types.File
	frontIdentificationFile.InitFromBytes(pngBuf.Bytes(), "frontIdentificationFile.png")
	backIdentificationFile.InitFromBytes(jpgBuf.Bytes(), "backIdentificationFile.jpg")

	requestParams := supporters.PostAuthValidateSignUpMultipartRequestBody{
		FirstName: "first_name",
		LastName: "last_name",
		Email: "test@example.com",
		Password: "Password",
		FrontIdentification: &frontIdentificationFile,
		BackIdentification: &backIdentificationFile,
	}

	result := testSupporterService.ValidateSignUp(ctx, &requestParams)

	assert.Nil(s.T(), result)
}

func (s *TestSupporterServiceSuite) TestValidateSignUp_ValidationErrorWithOptionalFields() {
	gifSignature := []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61}
	// NOTE:データを格納
	var gifBuf bytes.Buffer
	gifBuf.Write(gifSignature)
	
	var identificationFile openapi_types.File
	identificationFile.InitFromBytes(gifBuf.Bytes(), "frontIdentificationFile.gif")

	parsedTime, _ := time.Parse("2006-01-02", "1992-07-07")
	requestParams := supporters.PostAuthValidateSignUpMultipartRequestBody{
		FirstName: "first_name",
		LastName: "last_name",
		Email: "test@example.com",
		Password: "Password",
		Birthday: &openapi_types.Date{Time: parsedTime},
		FrontIdentification: &identificationFile,
		BackIdentification: &identificationFile,
	}

	result := testSupporterService.ValidateSignUp(ctx, &requestParams)

	assert.NotNil(s.T(), result)
	if errors, ok := result.(validation.Errors); ok {
		for field, err := range errors {
			message := err.Error()
			switch field {
			case "frontIdentification":
				assert.Equal(s.T(), "身分証明書(表)の拡張子はwebp, png, jpegのいずれかでお願いします。", message)
			case "backIdentification":
				assert.Equal(s.T(), "身分証明書(裏)の拡張子はwebp, png, jpegのいずれかでお願いします。", message)
			}
		}
	}
}

func (s *TestSupporterServiceSuite) TestSignUp_SuccessRequiredFields() {
	requestParams := supporters.PostAuthSignUpMultipartRequestBody{
		FirstName: "first_name",
		LastName: "last_name",
		Email: "test@example.com",
		Password: "Password",
	}

	result := testSupporterService.SignUp(ctx, requestParams)

	assert.Nil(s.T(), result)

	// NOTE: ユーザが作成されていることを確認
	user, err := models.Supporters(
		qm.Where("email = ?", "test@example.com"),
	).One(ctx, DBCon)
	if err != nil {
		s.T().Fatalf("failed to create user %v", err)
	}
	// NOTE: Birthdayはnullとなっている
	assert.Equal(s.T(), null.Time{}, user.Birthday)
	assert.Equal(s.T(), "", user.FrontIdentification)
	assert.Equal(s.T(), "", user.BackIdentification)
}

func (s *TestSupporterServiceSuite) TestSignUp_SuccessWithOptionalFields() {
	pngSignature := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	jpgSignature := []byte{0xFF, 0xD8, 0xFF, 0xE0}
	// NOTE:データを格納
	var pngBuf, jpgBuf bytes.Buffer
	pngBuf.Write(pngSignature)
	jpgBuf.Write(jpgSignature)
	
	var frontIdentificationFile, backIdentificationFile openapi_types.File
	frontIdentificationFile.InitFromBytes(pngBuf.Bytes(), "frontIdentificationFile.png")
	backIdentificationFile.InitFromBytes(jpgBuf.Bytes(), "backIdentificationFile.jpg")
	parsedTime, _ := time.Parse("2006-01-02", "1992-07-07")

	requestParams := supporters.PostAuthSignUpMultipartRequestBody{
		FirstName: "first_name",
		LastName: "last_name",
		Email: "test@example.com",
		Password: "Password",
		Birthday: &openapi_types.Date{Time: parsedTime},
		FrontIdentification: &frontIdentificationFile,
		BackIdentification: &backIdentificationFile,
	}

	result := testSupporterService.SignUp(ctx, requestParams)

	assert.Nil(s.T(), result)

	// NOTE: が作成されていることを確認
	user, err := models.Supporters(
		qm.Where("email = ?", "test@example.com"),
	).One(ctx, DBCon)
	if err != nil {
		s.T().Fatalf("failed to create user %v", err)
	}
	id := strconv.Itoa(user.ID)
	assert.Equal(s.T(), "1992-07-07", user.Birthday.Time.Format("2006-01-02"))
	assert.Equal(s.T(), "supporters/"+id+"/frontIdentificationFile.png", user.FrontIdentification)
	assert.Equal(s.T(), "supporters/"+id+"/backIdentificationFile.jpg", user.BackIdentification)
}

func TestSupporterService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestSupporterServiceSuite))
}

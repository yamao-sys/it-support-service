package registrationservices

import (
	registrationapi "apps/api/registration"
	models "apps/models"
	"bytes"
	"strconv"
	"testing"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"

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
	requestParams := registrationapi.PostSupporterValidateSignUpMultipartRequestBody{
		FirstName: "first_name",
		LastName: "last_name",
		Email: "test@example.com",
		Password: "Password",
	}

	result := testSupporterService.ValidateSignUp(&requestParams)

	assert.Nil(s.T(), result)
}

func (s *TestSupporterServiceSuite) TestValidateSignUp_ValidationErrorRequiredFields() {
	requestParams := registrationapi.PostSupporterValidateSignUpMultipartRequestBody{
		FirstName: "",
		LastName: "",
		Email: "",
		Password: "",
	}

	result := testSupporterService.ValidateSignUp(&requestParams)

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

	requestParams := registrationapi.PostSupporterValidateSignUpMultipartRequestBody{
		FirstName: "first_name",
		LastName: "last_name",
		Email: "test@example.com",
		Password: "Password",
		FrontIdentification: &frontIdentificationFile,
		BackIdentification: &backIdentificationFile,
	}

	result := testSupporterService.ValidateSignUp(&requestParams)

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
	requestParams := registrationapi.PostSupporterValidateSignUpMultipartRequestBody{
		FirstName: "first_name",
		LastName: "last_name",
		Email: "test@example.com",
		Password: "Password",
		Birthday: &openapi_types.Date{Time: parsedTime},
		FrontIdentification: &identificationFile,
		BackIdentification: &identificationFile,
	}

	result := testSupporterService.ValidateSignUp(&requestParams)

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
	requestParams := registrationapi.PostSupporterSignUpMultipartRequestBody{
		FirstName: "first_name",
		LastName: "last_name",
		Email: "test@example.com",
		Password: "Password",
	}

	result := testSupporterService.SignUp(ctx, requestParams)

	assert.Nil(s.T(), result)

	// NOTE: サポータが作成されていることを確認
	var supporter models.Supporter
	DBCon.Where("email = ?", "test@example.com").Take(&supporter)
	// NOTE: Birthdayはnullとなっている
	assert.Equal(s.T(), null.Time{}, supporter.Birthday)
	assert.Equal(s.T(), "", supporter.FrontIdentification)
	assert.Equal(s.T(), "", supporter.BackIdentification)
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

	requestParams := registrationapi.PostSupporterSignUpMultipartRequestBody{
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

	// NOTE: サポータが作成されていることを確認
	var supporter models.Supporter
	DBCon.Where("email = ?", "test@example.com").Take(&supporter)
	id := strconv.Itoa(supporter.ID)
	assert.Equal(s.T(), "1992-07-07", supporter.Birthday.Time.Format("2006-01-02"))
	assert.Equal(s.T(), "supporters/"+id+"/frontIdentificationFile.png", supporter.FrontIdentification)
	assert.Equal(s.T(), "supporters/"+id+"/backIdentificationFile.jpg", supporter.BackIdentification)
}

func TestSupporterService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestSupporterServiceSuite))
}

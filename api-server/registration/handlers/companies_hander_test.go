package registrationhandlers

import (
	registrationapi "apps/api/registration"
	models "apps/models"
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/oapi-codegen/testutil"
)

type TestCompaniesHandlerSuite struct {
	WithDBSuite
}

func (s *TestCompaniesHandlerSuite) SetupTest() {
	s.SetDBCon()

	s.initializeHandlers()

	// NOTE: CSRFトークンのセット
	s.SetCsrfHeaderValues()
}

func (s *TestCompaniesHandlerSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestCompaniesHandlerSuite) TestPostAuthValidateSignUp_SuccessRequiredFields() {
	body := new(bytes.Buffer)
	// NOTE: フォームデータを作成する
	mw := multipart.NewWriter(body)

	w, _ := mw.CreateFormField("name")
	w.Write([]byte("name"))
	w3, _ := mw.CreateFormField("email")
	w3.Write([]byte("test@example.com"))
	w4, _ := mw.CreateFormField("password")
	w4.Write([]byte("password"))

	// NOTE: 終了メッセージを書く
	mw.Close()

	result := testutil.NewRequest().Post("/companies/validate-sign-up").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithBody(body.Bytes()).WithContentType(mw.FormDataContentType()).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res registrationapi.CompanySignUpResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.Equal(s.T(), int(http.StatusOK), res.Code)
	jsonErrors, _ := json.Marshal(res.Errors)
	assert.Equal(s.T(), "{}", string(jsonErrors))
}

func (s *TestCompaniesHandlerSuite) TestPostAuthValidateSignUp_ValidationErrorRequiredFields() {
	body := new(bytes.Buffer)
	// NOTE: フォームデータを作成する
	mw := multipart.NewWriter(body)

	w, _ := mw.CreateFormField("name")
	w.Write([]byte(""))
	w3, _ := mw.CreateFormField("email")
	w3.Write([]byte(""))
	w4, _ := mw.CreateFormField("password")
	w4.Write([]byte(""))

	// NOTE: 終了メッセージを書く
	mw.Close()

	result := testutil.NewRequest().Post("/companies/validate-sign-up").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithBody(body.Bytes()).WithContentType(mw.FormDataContentType()).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res registrationapi.CompanySignUpResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.Equal(s.T(), int(http.StatusOK), res.Code)
	assert.Equal(s.T(), &[]string{"企業名は必須入力です。"}, res.Errors.Name)
	assert.Equal(s.T(), &[]string{"Emailは必須入力です。"}, res.Errors.Email)
	assert.Equal(s.T(), &[]string{"パスワードは必須入力です。"}, res.Errors.Password)
}

func (s *TestCompaniesHandlerSuite) TestPostAuthValidateSignUp_SuccessWithOptionalFields() {
	body := new(bytes.Buffer)
	// NOTE: フォームデータを作成する
	mw := multipart.NewWriter(body)

	w, _ := mw.CreateFormField("name")
	w.Write([]byte("name"))
	w3, _ := mw.CreateFormField("email")
	w3.Write([]byte("test@example.com"))
	w4, _ := mw.CreateFormField("password")
	w4.Write([]byte("password"))

	pngSignature := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	// NOTE:データを格納
	var pngBuf bytes.Buffer
	pngBuf.Write(pngSignature)
	w6, _ := mw.CreateFormFile("finalTaxReturn", "finalTaxReturn.png")
	w6.Write(pngBuf.Bytes())

	// NOTE: 終了メッセージを書く
	mw.Close()

	result := testutil.NewRequest().Post("/companies/validate-sign-up").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithBody(body.Bytes()).WithContentType(mw.FormDataContentType()).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res registrationapi.CompanySignUpResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.Equal(s.T(), int(http.StatusOK), res.Code)
	jsonErrors, _ := json.Marshal(res.Errors)
	assert.Equal(s.T(), "{}", string(jsonErrors))
}

func (s *TestCompaniesHandlerSuite) TestPostAuthValidateSignUp_ValidationErrorWithOptionalFields() {
	body := new(bytes.Buffer)
	// NOTE: フォームデータを作成する
	mw := multipart.NewWriter(body)

	w, _ := mw.CreateFormField("name")
	w.Write([]byte("name"))
	w3, _ := mw.CreateFormField("email")
	w3.Write([]byte("test@example.com"))
	w4, _ := mw.CreateFormField("password")
	w4.Write([]byte("password"))

	gifSignature := []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61}
	// NOTE:データを格納
	var gifBuf bytes.Buffer
	gifBuf.Write(gifSignature)
	w6, _ := mw.CreateFormFile("finalTaxReturn", "finalTaxReturn.gif")
	w6.Write(gifBuf.Bytes())

	// NOTE: 終了メッセージを書く
	mw.Close()

	result := testutil.NewRequest().Post("/companies/validate-sign-up").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithBody(body.Bytes()).WithContentType(mw.FormDataContentType()).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res registrationapi.CompanySignUpResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.Equal(s.T(), int(http.StatusOK), res.Code)
	assert.Equal(s.T(), &[]string{"確定申告書の拡張子はwebp, png, jpegのいずれかでお願いします。"}, res.Errors.FinalTaxReturn)
}

func (s *TestCompaniesHandlerSuite) TestPostAuthSignUp_SuccessRequiredFields() {
	body := new(bytes.Buffer)
	// NOTE: フォームデータを作成する
	mw := multipart.NewWriter(body)

	w, _ := mw.CreateFormField("name")
	w.Write([]byte("name"))
	w3, _ := mw.CreateFormField("email")
	w3.Write([]byte("test@example.com"))
	w4, _ := mw.CreateFormField("password")
	w4.Write([]byte("password"))

	// NOTE: 終了メッセージを書く
	mw.Close()

	result := testutil.NewRequest().Post("/companies/sign-up").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithBody(body.Bytes()).WithContentType(mw.FormDataContentType()).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res registrationapi.CompanySignUpResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.Equal(s.T(), int(http.StatusOK), res.Code)
	jsonErrors, _ := json.Marshal(res.Errors)
	assert.Equal(s.T(), "{}", string(jsonErrors))

	// NOTE: Companyが作成されていることを確認
	var company models.Company
	DBCon.Where("email = ?", "test@example.com").Take(&company)
	assert.Equal(s.T(), "name", company.Name)
	assert.Equal(s.T(), "", company.FinalTaxReturn)
}

func (s *TestCompaniesHandlerSuite) TestPostAuthSignUp_SuccessWithOptionalFields() {
	body := new(bytes.Buffer)
	// NOTE: フォームデータを作成する
	mw := multipart.NewWriter(body)

	w, _ := mw.CreateFormField("name")
	w.Write([]byte("name"))
	w3, _ := mw.CreateFormField("email")
	w3.Write([]byte("test@example.com"))
	w4, _ := mw.CreateFormField("password")
	w4.Write([]byte("password"))

	pngSignature := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	// NOTE:データを格納
	var pngBuf bytes.Buffer
	pngBuf.Write(pngSignature)
	w6, _ := mw.CreateFormFile("finalTaxReturn", "finalTaxReturn.png")
	w6.Write(pngBuf.Bytes())

	// NOTE: 終了メッセージを書く
	mw.Close()

	result := testutil.NewRequest().Post("/companies/sign-up").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithBody(body.Bytes()).WithContentType(mw.FormDataContentType()).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res registrationapi.CompanySignUpResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.Equal(s.T(), int(http.StatusOK), res.Code)
	jsonErrors, _ := json.Marshal(res.Errors)
	assert.Equal(s.T(), "{}", string(jsonErrors))

	// NOTE: companyが作成されていることを確認
	var company models.Company
	DBCon.Where("email = ?", "test@example.com").Take(&company)
	id := strconv.Itoa(company.ID)
	assert.Equal(s.T(), "companies/"+id+"/finalTaxReturn.png", company.FinalTaxReturn)
}

func TestCompaniesHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(TestCompaniesHandlerSuite))
}

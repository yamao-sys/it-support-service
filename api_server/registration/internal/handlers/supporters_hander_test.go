package handlers

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"registration/api/generated/supporters"
	"registration/internal/services"
	models "registration/models/generated"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/oapi-codegen/testutil"
)

var (
	testSupportersHandler SupportersHandler
)

type TestSupportersHandlerSuite struct {
	WithDBSuite
}

func (s *TestSupportersHandlerSuite) SetupTest() {
	s.SetDBCon()

	authService := services.NewSupporterService(DBCon)

	// NOTE: テスト対象のコントローラを設定
	testSupportersHandler = NewSupportersHandler(authService)
}

func (s *TestSupportersHandlerSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestSupportersHandlerSuite) TestPostAuthValidateSignUp_SuccessRequiredFields() {
	body := new(bytes.Buffer)
	// NOTE: フォームデータを作成する
	mw := multipart.NewWriter(body)

	w, _ := mw.CreateFormField("firstName")
	w.Write([]byte("first_name"))
	w2, _ := mw.CreateFormField("lastName")
	w2.Write([]byte("last_name"))
	w3, _ := mw.CreateFormField("email")
	w3.Write([]byte("test@example.com"))
	w4, _ := mw.CreateFormField("password")
	w4.Write([]byte("password"))

	// NOTE: 終了メッセージを書く
	mw.Close()

	e := echo.New()

	strictHandler := supporters.NewStrictHandler(testSupportersHandler, nil)
	supporters.RegisterHandlers(e, strictHandler)

	result := testutil.NewRequest().Post("/supporters/validateSignUp").WithBody(body.Bytes()).WithContentType(mw.FormDataContentType()).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res supporters.SignUpResponseJSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.Equal(s.T(), int64(http.StatusOK), res.Code)
	jsonErrors, _ := json.Marshal(res.Errors)
	assert.Equal(s.T(), "{}", string(jsonErrors))
}

func (s *TestSupportersHandlerSuite) TestPostAuthValidateSignUp_ValidationErrorRequiredFields() {
	body := new(bytes.Buffer)
	// NOTE: フォームデータを作成する
	mw := multipart.NewWriter(body)

	w, _ := mw.CreateFormField("firstName")
	w.Write([]byte(""))
	w2, _ := mw.CreateFormField("lastName")
	w2.Write([]byte(""))
	w3, _ := mw.CreateFormField("email")
	w3.Write([]byte(""))
	w4, _ := mw.CreateFormField("password")
	w4.Write([]byte(""))

	// NOTE: 終了メッセージを書く
	mw.Close()

	e := echo.New()

	strictHandler := supporters.NewStrictHandler(testSupportersHandler, nil)
	supporters.RegisterHandlers(e, strictHandler)

	result := testutil.NewRequest().Post("/supporters/validateSignUp").WithBody(body.Bytes()).WithContentType(mw.FormDataContentType()).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res supporters.SignUpResponseJSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.Equal(s.T(), int64(http.StatusOK), res.Code)
	assert.Equal(s.T(), &[]string{"名は必須入力です。"}, res.Errors.FirstName)
	assert.Equal(s.T(), &[]string{"姓は必須入力です。"}, res.Errors.LastName)
	assert.Equal(s.T(), &[]string{"Emailは必須入力です。"}, res.Errors.Email)
	assert.Equal(s.T(), &[]string{"パスワードは必須入力です。"}, res.Errors.Password)
}

func (s *TestSupportersHandlerSuite) TestPostAuthValidateSignUp_SuccessWithOptionalFields() {
	body := new(bytes.Buffer)
	// NOTE: フォームデータを作成する
	mw := multipart.NewWriter(body)

	w, _ := mw.CreateFormField("firstName")
	w.Write([]byte("first_name"))
	w2, _ := mw.CreateFormField("lastName")
	w2.Write([]byte("last_name"))
	w3, _ := mw.CreateFormField("email")
	w3.Write([]byte("test@example.com"))
	w4, _ := mw.CreateFormField("password")
	w4.Write([]byte("password"))
	w5, _ := mw.CreateFormField("birthday")
	w5.Write([]byte("1992-07-07"))

	pngSignature := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	jpgSignature := []byte{0xFF, 0xD8, 0xFF, 0xE0}
	// NOTE:データを格納
	var pngBuf, jpgBuf bytes.Buffer
	pngBuf.Write(pngSignature)
	jpgBuf.Write(jpgSignature)
	w6, _ := mw.CreateFormFile("frontIdentification", "frontIdentificationFile.png")
	w6.Write(pngBuf.Bytes())
	w7, _ := mw.CreateFormFile("backIdentification", "backIdentificationFile.jpg")
	w7.Write(jpgBuf.Bytes())

	// NOTE: 終了メッセージを書く
	mw.Close()

	e := echo.New()

	strictHandler := supporters.NewStrictHandler(testSupportersHandler, nil)
	supporters.RegisterHandlers(e, strictHandler)

	result := testutil.NewRequest().Post("/supporters/validateSignUp").WithBody(body.Bytes()).WithContentType(mw.FormDataContentType()).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res supporters.SignUpResponseJSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.Equal(s.T(), int64(http.StatusOK), res.Code)
	jsonErrors, _ := json.Marshal(res.Errors)
	assert.Equal(s.T(), "{}", string(jsonErrors))
}

func (s *TestSupportersHandlerSuite) TestPostAuthValidateSignUp_ValidationErrorWithOptionalFields() {
	body := new(bytes.Buffer)
	// NOTE: フォームデータを作成する
	mw := multipart.NewWriter(body)

	w, _ := mw.CreateFormField("firstName")
	w.Write([]byte("first_name"))
	w2, _ := mw.CreateFormField("lastName")
	w2.Write([]byte("last_name"))
	w3, _ := mw.CreateFormField("email")
	w3.Write([]byte("test@example.com"))
	w4, _ := mw.CreateFormField("password")
	w4.Write([]byte("password"))
	w5, _ := mw.CreateFormField("birthday")
	w5.Write([]byte("1992-07-07"))

	gifSignature := []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61}
	// NOTE:データを格納
	var gifBuf bytes.Buffer
	gifBuf.Write(gifSignature)
	w6, _ := mw.CreateFormFile("frontIdentification", "frontIdentificationFile.gif")
	w6.Write(gifBuf.Bytes())
	w7, _ := mw.CreateFormFile("backIdentification", "backIdentificationFile.gif")
	w7.Write(gifBuf.Bytes())

	// NOTE: 終了メッセージを書く
	mw.Close()

	e := echo.New()

	strictHandler := supporters.NewStrictHandler(testSupportersHandler, nil)
	supporters.RegisterHandlers(e, strictHandler)

	result := testutil.NewRequest().Post("/supporters/validateSignUp").WithBody(body.Bytes()).WithContentType(mw.FormDataContentType()).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res supporters.SignUpResponseJSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.Equal(s.T(), int64(http.StatusOK), res.Code)
	assert.Equal(s.T(), &[]string{"身分証明書(表)の拡張子はwebp, png, jpegのいずれかでお願いします。"}, res.Errors.FrontIdentification)
	assert.Equal(s.T(), &[]string{"身分証明書(裏)の拡張子はwebp, png, jpegのいずれかでお願いします。"}, res.Errors.BackIdentification)
}

func (s *TestSupportersHandlerSuite) TestPostAuthSignUp_SuccessRequiredFields() {
	body := new(bytes.Buffer)
	// NOTE: フォームデータを作成する
	mw := multipart.NewWriter(body)

	w, _ := mw.CreateFormField("firstName")
	w.Write([]byte("first_name"))
	w2, _ := mw.CreateFormField("lastName")
	w2.Write([]byte("last_name"))
	w3, _ := mw.CreateFormField("email")
	w3.Write([]byte("test@example.com"))
	w4, _ := mw.CreateFormField("password")
	w4.Write([]byte("password"))

	// NOTE: 終了メッセージを書く
	mw.Close()

	e := echo.New()

	strictHandler := supporters.NewStrictHandler(testSupportersHandler, nil)
	supporters.RegisterHandlers(e, strictHandler)

	result := testutil.NewRequest().Post("/supporters/signUp").WithBody(body.Bytes()).WithContentType(mw.FormDataContentType()).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res supporters.SignUpResponseJSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.Equal(s.T(), int64(http.StatusOK), res.Code)
	jsonErrors, _ := json.Marshal(res.Errors)
	assert.Equal(s.T(), "{}", string(jsonErrors))

	// NOTE: Supporterが作成されていることを確認
	supporter, err := models.Supporters(
		qm.Where("email = ?", "test@example.com"),
	).One(ctx, DBCon)
	if err != nil {
		s.T().Fatalf("failed to create supporter %v", err)
	}
	// NOTE: Birthdayはnullとなっている
	assert.Equal(s.T(), null.Time{}, supporter.Birthday)
	assert.Equal(s.T(), "", supporter.FrontIdentification)
	assert.Equal(s.T(), "", supporter.BackIdentification)
}

func (s *TestSupportersHandlerSuite) TestPostAuthSignUp_SuccessWithOptionalFields() {
	body := new(bytes.Buffer)
	// NOTE: フォームデータを作成する
	mw := multipart.NewWriter(body)

	w, _ := mw.CreateFormField("firstName")
	w.Write([]byte("first_name"))
	w2, _ := mw.CreateFormField("lastName")
	w2.Write([]byte("last_name"))
	w3, _ := mw.CreateFormField("email")
	w3.Write([]byte("test@example.com"))
	w4, _ := mw.CreateFormField("password")
	w4.Write([]byte("password"))
	w5, _ := mw.CreateFormField("birthday")
	w5.Write([]byte("1992-07-07"))

	pngSignature := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	jpgSignature := []byte{0xFF, 0xD8, 0xFF, 0xE0}
	// NOTE:データを格納
	var pngBuf, jpgBuf bytes.Buffer
	pngBuf.Write(pngSignature)
	jpgBuf.Write(jpgSignature)
	w6, _ := mw.CreateFormFile("frontIdentification", "frontIdentificationFile.png")
	w6.Write(pngBuf.Bytes())
	w7, _ := mw.CreateFormFile("backIdentification", "backIdentificationFile.jpg")
	w7.Write(jpgBuf.Bytes())

	// NOTE: 終了メッセージを書く
	mw.Close()

	e := echo.New()

	strictHandler := supporters.NewStrictHandler(testSupportersHandler, nil)
	supporters.RegisterHandlers(e, strictHandler)

	result := testutil.NewRequest().Post("/supporters/signUp").WithBody(body.Bytes()).WithContentType(mw.FormDataContentType()).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res supporters.SignUpResponseJSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.Equal(s.T(), int64(http.StatusOK), res.Code)
	jsonErrors, _ := json.Marshal(res.Errors)
	assert.Equal(s.T(), "{}", string(jsonErrors))

	// NOTE: Supporterが作成されていることを確認
	supporter, err := models.Supporters(
		qm.Where("email = ?", "test@example.com"),
	).One(ctx, DBCon)
	if err != nil {
		s.T().Fatalf("failed to create supporter %v", err)
	}
	id := strconv.Itoa(supporter.ID)
	assert.Equal(s.T(), "1992-07-07", supporter.Birthday.Time.Format("2006-01-02"))
	assert.Equal(s.T(), "supporters/"+id+"/frontIdentificationFile.png", supporter.FrontIdentification)
	assert.Equal(s.T(), "supporters/"+id+"/backIdentificationFile.jpg", supporter.BackIdentification)
}

func (s *TestSupportersHandlerSuite) TestPostAuthSignUp_SuccessWithEmptyBirthday() {
	body := new(bytes.Buffer)
	// NOTE: フォームデータを作成する
	mw := multipart.NewWriter(body)

	w, _ := mw.CreateFormField("firstName")
	w.Write([]byte("first_name"))
	w2, _ := mw.CreateFormField("lastName")
	w2.Write([]byte("last_name"))
	w3, _ := mw.CreateFormField("email")
	w3.Write([]byte("test@example.com"))
	w4, _ := mw.CreateFormField("password")
	w4.Write([]byte("password"))
	w5, _ := mw.CreateFormField("birthday")
	w5.Write([]byte(""))

	pngSignature := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	jpgSignature := []byte{0xFF, 0xD8, 0xFF, 0xE0}
	// NOTE:データを格納
	var pngBuf, jpgBuf bytes.Buffer
	pngBuf.Write(pngSignature)
	jpgBuf.Write(jpgSignature)
	w6, _ := mw.CreateFormFile("frontIdentification", "frontIdentificationFile.png")
	w6.Write(pngBuf.Bytes())
	w7, _ := mw.CreateFormFile("backIdentification", "backIdentificationFile.jpg")
	w7.Write(jpgBuf.Bytes())

	// NOTE: 終了メッセージを書く
	mw.Close()

	e := echo.New()

	strictHandler := supporters.NewStrictHandler(testSupportersHandler, nil)
	supporters.RegisterHandlers(e, strictHandler)

	result := testutil.NewRequest().Post("/supporters/signUp").WithBody(body.Bytes()).WithContentType(mw.FormDataContentType()).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res supporters.SignUpResponseJSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.Equal(s.T(), int64(http.StatusOK), res.Code)
	jsonErrors, _ := json.Marshal(res.Errors)
	assert.Equal(s.T(), "{}", string(jsonErrors))

	// NOTE: Supporterが作成されていることを確認
	supporter, err := models.Supporters(
		qm.Where("email = ?", "test@example.com"),
	).One(ctx, DBCon)
	if err != nil {
		s.T().Fatalf("failed to create supporter %v", err)
	}
	id := strconv.Itoa(supporter.ID)
	// NOTE: Birthdayはnullとなっている
	assert.Equal(s.T(), null.Time{}, supporter.Birthday)
	assert.Equal(s.T(), "supporters/"+id+"/frontIdentificationFile.png", supporter.FrontIdentification)
	assert.Equal(s.T(), "supporters/"+id+"/backIdentificationFile.jpg", supporter.BackIdentification)
}

func TestSupportersHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(TestSupportersHandlerSuite))
}

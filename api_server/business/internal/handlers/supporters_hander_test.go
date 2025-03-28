package handlers

import (
	"business/api/generated/supporters"
	"business/internal/services"
	models "business/models/generated"
	"business/test/factories"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	testSupportersHandler SupportersHandler
)

type TestSupportersHandlerSuite struct {
	WithDBSuite
}

func (s *TestSupportersHandlerSuite) SetupTest() {
	s.SetDBCon()

	supporterService := services.NewSupporterService(DBCon)

	// NOTE: テスト対象のコントローラを設定
	testSupportersHandler = NewSupportersHandler(supporterService)

	// NOTE: CSRFトークンのセット
	s.SetCsrfHeaderValues()
}

func (s *TestSupportersHandlerSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestSupportersHandlerSuite) TestPostSupportersSignIn_StatusOk() {
	// NOTE: テスト用サポータの作成
	supporter := factories.SupporterFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Supporter)
	if err := supporter.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test supporter %v", err)
	}

	strictHandler := supporters.NewStrictHandler(testSupportersHandler, nil)
	supporters.RegisterHandlers(e, strictHandler)

	reqBody := supporters.SignInInput{
		Email: "test@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/supporters/signIn").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	cookieString := result.Recorder.Result().Header.Values("Set-Cookie")[0]
	assert.NotEmpty(s.T(), cookieString)
}

func (s *TestSupportersHandlerSuite) TestPostSupportersSignIn_BadRequest() {
	// NOTE: テスト用サポータの作成
	supporter := factories.SupporterFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Supporter)
	if err := supporter.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test supporter %v", err)
	}

	strictHandler := supporters.NewStrictHandler(testSupportersHandler, nil)
	supporters.RegisterHandlers(e, strictHandler)

	reqBody := supporters.SignInInput{
		Email: "test_@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/supporters/signIn").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), int(http.StatusBadRequest), result.Code())

	var res supporters.SignInBadRequestResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), []string{"メールアドレスまたはパスワードに該当するサポータが存在しません。"}, res.Errors)
}

func TestSupportersHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(TestSupportersHandlerSuite))
}

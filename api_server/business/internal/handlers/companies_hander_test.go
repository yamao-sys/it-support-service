package handlers

import (
	"business/api/generated/companies"
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
	testCompaniesHandler CompaniesHandler
)

type TestCompaniesHandlerSuite struct {
	WithDBSuite
}

func (s *TestCompaniesHandlerSuite) SetupTest() {
	s.SetDBCon()

	companyService := services.NewCompanyService(DBCon)

	// NOTE: テスト対象のコントローラを設定
	testCompaniesHandler = NewCompaniesHandler(companyService)

	// NOTE: CSRFトークンのセット
	s.SetCsrfHeaderValues()
}

func (s *TestCompaniesHandlerSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestCompaniesHandlerSuite) TestPostCompaniesSignIn_StatusOk() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	if err := company.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test company %v", err)
	}

	strictHandler := companies.NewStrictHandler(testCompaniesHandler, nil)
	companies.RegisterHandlers(e, strictHandler)

	reqBody := companies.SignInInput{
		Email: "test@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/companies/signIn").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	cookieString := result.Recorder.Result().Header.Values("Set-Cookie")[0]
	assert.NotEmpty(s.T(), cookieString)
}

func (s *TestCompaniesHandlerSuite) TestPostCompaniesSignIn_BadRequest() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	if err := company.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test company %v", err)
	}

	strictHandler := companies.NewStrictHandler(testCompaniesHandler, nil)
	companies.RegisterHandlers(e, strictHandler)

	reqBody := companies.SignInInput{
		Email: "test_@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/companies/signIn").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), int(http.StatusBadRequest), result.Code())

	var res companies.SignInBadRequestResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), []string{"メールアドレスまたはパスワードに該当する企業が存在しません。"}, res.Errors)
}

func TestCompaniesHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(TestCompaniesHandlerSuite))
}

package businesshandlers

import (
	businessapi "apps/api/business"
	businessservices "apps/business/services"
	models "apps/models"
	"apps/test/factories"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSupportersHandlerSuite struct {
	WithDBSuite
}

func (s *TestSupportersHandlerSuite) SetupTest() {
	s.SetDBCon()

	s.initializeHandlers(businessservices.NewProjectService(DBCon), businessservices.NewPlanService(DBCon))

	// NOTE: CSRFトークンのセット
	s.SetCsrfHeaderValues()
}

func (s *TestSupportersHandlerSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestSupportersHandlerSuite) TestPostSupportersSignIn_StatusOk() {
	// NOTE: テスト用サポータの作成
	supporter := factories.SupporterFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Supporter)
	DBCon.Create(supporter)

	reqBody := businessapi.SupporterSignInInput{
		Email: "test@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/supporters/sign-in").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	cookieString := result.Recorder.Result().Header.Values("Set-Cookie")[0]
	assert.NotEmpty(s.T(), cookieString)
}

func (s *TestSupportersHandlerSuite) TestPostSupportersSignIn_BadRequest() {
	// NOTE: テスト用サポータの作成
	supporter := factories.SupporterFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Supporter)
	DBCon.Create(supporter)

	reqBody := businessapi.SupporterSignInInput{
		Email: "test_@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/supporters/sign-in").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), int(http.StatusBadRequest), result.Code())

	var res businessapi.SupporterSignInBadRequestResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), []string{"メールアドレスまたはパスワードに該当するサポータが存在しません。"}, res.Errors)
}

func TestSupportersHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(TestSupportersHandlerSuite))
}

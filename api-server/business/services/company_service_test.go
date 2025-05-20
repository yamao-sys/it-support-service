package businessservices

import (
	businessapi "apps/api/business"
	models "apps/models"
	"apps/test/factories"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

func (s *TestCompanyServiceSuite) TestSignIn_StatusOK() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	DBCon.Create(company)

	requestParams := businessapi.PostCompanySignInJSONRequestBody{Email: "test@example.com", Password: "password"}

	statusCode, tokenString, err := testCompanyService.SignIn(requestParams)

	assert.Equal(s.T(), int(http.StatusOK), statusCode)
	assert.NotNil(s.T(), tokenString)
	assert.Nil(s.T(), err)
}

func (s *TestCompanyServiceSuite) TestSignIn_BadRequest() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	DBCon.Create(company)

	requestParams := businessapi.PostCompanySignInJSONRequestBody{Email: "test_@example.com", Password: "password"}

	statusCode, tokenString, err := testCompanyService.SignIn(requestParams)

	assert.Equal(s.T(), int(http.StatusBadRequest), statusCode)
	assert.Equal(s.T(), "", tokenString)
	assert.Equal(s.T(), "メールアドレスまたはパスワードに該当する企業が存在しません。", err.Error())
}

func TestCompanyService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestCompanyServiceSuite))
}

package services

import (
	"business/api/generated/companies"
	models "business/models/generated"
	"business/test/factories"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
	supporter := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	if err := supporter.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test supporter %v", err)
	}

	requestParams := companies.PostCompaniesSignInJSONRequestBody{Email: "test@example.com", Password: "password"}

	statusCode, tokenString, err := testCompanyService.SignIn(ctx, requestParams)

	assert.Equal(s.T(), int64(http.StatusOK), statusCode)
	assert.NotNil(s.T(), tokenString)
	assert.Nil(s.T(), err)
}

func (s *TestCompanyServiceSuite) TestSignIn_BadRequest() {
	// NOTE: テスト用企業の作成
	supporter := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	if err := supporter.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test user %v", err)
	}

	requestParams := companies.PostCompaniesSignInJSONRequestBody{Email: "test_@example.com", Password: "password"}

	statusCode, tokenString, err := testCompanyService.SignIn(ctx, requestParams)

	assert.Equal(s.T(), int64(http.StatusBadRequest), statusCode)
	assert.Equal(s.T(), "", tokenString)
	assert.Equal(s.T(), "メールアドレスまたはパスワードに該当する企業が存在しません。", err.Error())
}

func TestCompanyService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestCompanyServiceSuite))
}

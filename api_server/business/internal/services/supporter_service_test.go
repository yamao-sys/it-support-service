package services

import (
	"business/api/generated/supporters"
	models "business/models/generated"
	"business/test/factories"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func (s *TestSupporterServiceSuite) TestSignIn_StatusOK() {
	// NOTE: テスト用サポータの作成
	supporter := factories.SupporterFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Supporter)
	if err := supporter.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test supporter %v", err)
	}

	requestParams := supporters.PostSupportersSignInJSONRequestBody{Email: "test@example.com", Password: "password"}

	statusCode, tokenString, err := testSupporterService.SignIn(ctx, requestParams)

	assert.Equal(s.T(), int64(http.StatusOK), statusCode)
	assert.NotNil(s.T(), tokenString)
	assert.Nil(s.T(), err)
}

func (s *TestSupporterServiceSuite) TestSignIn_BadRequest() {
	// NOTE: テスト用サポータの作成
	supporter := factories.SupporterFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Supporter)
	if err := supporter.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test user %v", err)
	}

	requestParams := supporters.PostSupportersSignInJSONRequestBody{Email: "test_@example.com", Password: "password"}

	statusCode, tokenString, err := testSupporterService.SignIn(ctx, requestParams)

	assert.Equal(s.T(), int64(http.StatusBadRequest), statusCode)
	assert.Equal(s.T(), "", tokenString)
	assert.Equal(s.T(), "メールアドレスまたはパスワードに該当するサポータが存在しません。", err.Error())
}

func TestSupporterService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestSupporterServiceSuite))
}

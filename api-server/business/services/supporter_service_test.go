package businessservices

import (
	businessapi "apps/api/business"
	models "apps/models/generated"
	"apps/test/factories"
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
	supporter.Insert(ctx, DBCon, boil.Infer())

	requestParams := businessapi.PostSupportersSignInJSONRequestBody{Email: "test@example.com", Password: "password"}

	statusCode, tokenString, err := testSupporterService.SignIn(ctx, requestParams)

	assert.Equal(s.T(), int(http.StatusOK), statusCode)
	assert.NotNil(s.T(), tokenString)
	assert.Nil(s.T(), err)
}

func (s *TestSupporterServiceSuite) TestSignIn_BadRequest() {
	// NOTE: テスト用サポータの作成
	supporter := factories.SupporterFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Supporter)
	supporter.Insert(ctx, DBCon, boil.Infer())

	requestParams := businessapi.PostSupportersSignInJSONRequestBody{Email: "test_@example.com", Password: "password"}

	statusCode, tokenString, err := testSupporterService.SignIn(ctx, requestParams)

	assert.Equal(s.T(), int(http.StatusBadRequest), statusCode)
	assert.Equal(s.T(), "", tokenString)
	assert.Equal(s.T(), "メールアドレスまたはパスワードに該当するサポータが存在しません。", err.Error())
}

func TestSupporterService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestSupporterServiceSuite))
}

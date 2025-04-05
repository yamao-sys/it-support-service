package businesshandlers

import (
	businessapi "apps/api/business"
	businessmiddlewares "apps/business/middlewares"
	businessservices "apps/business/services"
	"apps/database"
	models "apps/models/generated"
	"apps/test/factories"
	"context"
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type WithDBSuite struct {
	suite.Suite
}

var (
	DBCon *sql.DB
	ctx   context.Context
	// token string
	e *echo.Echo
	csrfToken string
	csrfTokenCookie string
)

// func (s *WithDBSuite) SetupSuite()                           {} // テストスイート実施前の処理
// func (s *WithDBSuite) TearDownSuite()                        {} // テストスイート終了後の処理
// func (s *WithDBSuite) SetupTest()                            {} // テストケース実施前の処理
// func (s *WithDBSuite) TearDownTest()                         {} // テストケース終了後の処理
// func (s *WithDBSuite) BeforeTest(suiteName, testName string) {} // テストケース実施前の処理
// func (s *WithDBSuite) AfterTest(suiteName, testName string)  {} // テストケース終了後の処理

func init() {
	txdb.Register("txdb-controller", "mysql", database.GetDsn())
	ctx = context.Background()

	e = businessmiddlewares.ApplyMiddlewares(echo.New())
}

func (s *WithDBSuite) SetDBCon() {
	db, err := sql.Open("txdb-controller", "connect")
	if err != nil {
		s.T().Fatalf("failed to initialize DB: %v", err)
	}
	DBCon = db
}

func (s *WithDBSuite) CloseDB() {
	DBCon.Close()
}

func (s *WithDBSuite) SetCsrfHeaderValues() {
	result := testutil.NewRequest().Get("/csrf").GoWithHTTPHandler(s.T(), e)

	var res businessapi.GetCsrf200JSONResponse
	result.UnmarshalJsonToObject(&res)

	csrfToken = res.CsrfToken
	csrfTokenCookie = result.Recorder.Result().Header.Values("Set-Cookie")[0]
}


func (s *WithDBSuite) initializeHandlers() {
	csrfServer := NewCsrfHandler()

	supporterService := businessservices.NewSupporterService(DBCon)
	testSupportersHandler := NewSupportersHandler(supporterService)

	companyService := businessservices.NewCompanyService(DBCon)
	testCompaniesHandler := NewCompaniesHandler(companyService)

	projectService := businessservices.NewProjectService(DBCon)
	testProjectsHandler := NewProjectsHandler(projectService)

	mainHandler := NewMainHandler(csrfServer, testSupportersHandler, testCompaniesHandler, testProjectsHandler)

	strictHandler := businessapi.NewStrictHandler(mainHandler, []businessapi.StrictMiddlewareFunc{businessmiddlewares.AuthMiddleware})
	businessapi.RegisterHandlers(e, strictHandler)
}

func (s *WithDBSuite) companySignIn() (company *models.Company, cookieString string) {
	// NOTE: テスト用企業の作成
	company = factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())

	reqBody := businessapi.CompanySignInInput{
		Email: "test@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/companies/signIn").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	cookieString = result.Recorder.Result().Header.Values("Set-Cookie")[0]

	return company, cookieString
}

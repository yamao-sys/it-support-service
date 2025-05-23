package businesshandlers

import (
	businessapi "apps/api/business"
	businessmiddlewares "apps/business/middlewares"
	businessservices "apps/business/services"
	"apps/database"
	models "apps/models"
	"apps/test/factories"
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type WithDBSuite struct {
	suite.Suite
}

var (
	DBCon *gorm.DB
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
	txdb.Register("txdb-handler", "mysql", database.GetDsn())

	e = businessmiddlewares.ApplyMiddlewares(echo.New())
}

func (s *WithDBSuite) SetDBCon() {
	db, err := sql.Open("txdb-handler", "connect")
	if err != nil {
		s.T().Fatalf("failed to initialize DB: %v", err)
	}
	gormDB, err := gorm.Open(
		mysql.New(mysql.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)
	if err != nil {
		s.T().Fatalf("failed to open gorm DB: %v", err)
	}
	DBCon = gormDB
}

func (s *WithDBSuite) CloseDB() {
	database.Close(DBCon)
}

func (s *WithDBSuite) SetCsrfHeaderValues() {
	result := testutil.NewRequest().Get("/csrf").GoWithHTTPHandler(s.T(), e)

	var res businessapi.GetCsrf200JSONResponse
	result.UnmarshalJsonToObject(&res)

	csrfToken = res.CsrfToken
	csrfTokenCookie = result.Recorder.Result().Header.Values("Set-Cookie")[0]
}


func (s *WithDBSuite) initializeHandlers(projectService businessservices.ProjectService, planService businessservices.PlanService) {
	csrfServer := NewCsrfHandler()

	supporterService := businessservices.NewSupporterService(DBCon)
	testSupportersHandler := NewSupportersHandler(supporterService)

	companyService := businessservices.NewCompanyService(DBCon)
	testCompaniesHandler := NewCompaniesHandler(companyService)

	testProjectsHandler := NewProjectsHandler(projectService)
	
	testPlansHandler := NewPlansHandler(planService)

	toProjectService := businessservices.NewToProjectService(DBCon)
	testToProjectsHandler := NewToProjectsHandler(toProjectService)

	mainHandler := NewMainHandler(csrfServer, testSupportersHandler, testCompaniesHandler, testProjectsHandler, testPlansHandler, testToProjectsHandler)

	strictHandler := businessapi.NewStrictHandler(mainHandler, []businessapi.StrictMiddlewareFunc{businessmiddlewares.AuthMiddleware})
	businessapi.RegisterHandlers(e, strictHandler)
}

func (s *WithDBSuite) companySignIn() (company *models.Company, cookieString string) {
	// NOTE: テスト用企業の作成
	company = factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	DBCon.Create(company)

	reqBody := businessapi.CompanySignInInput{
		Email: "test@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/companies/sign-in").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	cookieString = result.Recorder.Result().Header.Values("Set-Cookie")[0]

	return company, cookieString
}

func (s *WithDBSuite) supporterSignIn() (supporter *models.Supporter, cookieString string) {
	// NOTE: テスト用サポータの作成
	supporter = factories.SupporterFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Supporter)
	DBCon.Create(supporter)

	reqBody := businessapi.SupporterSignInInput{
		Email: "test@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/supporters/sign-in").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	cookieString = result.Recorder.Result().Header.Values("Set-Cookie")[0]

	return supporter, cookieString
}

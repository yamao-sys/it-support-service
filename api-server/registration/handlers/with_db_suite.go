package registrationhandlers

import (
	registrationapi "apps/api/registration"
	"apps/database"
	registrationmiddlewares "apps/registration/middlewares"
	registrationservices "apps/registration/services"
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
	txdb.Register("txdb-registration-handler", "mysql", database.GetDsn())
	e = registrationmiddlewares.ApplyMiddlewares(echo.New())
}

func (s *WithDBSuite) SetDBCon() {
	db, err := sql.Open("txdb-registration-handler", "connect")
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

	var res registrationapi.GetCsrf200JSONResponse
	err := result.UnmarshalJsonToObject(&res)
	if err != nil {
		s.T().Error(err.Error())
	}

	csrfToken = res.CsrfToken
	csrfTokenCookie = result.Recorder.Result().Header.Values("Set-Cookie")[0]
}

func (s *WithDBSuite) initializeHandlers() {
	csrfServer := NewCsrfHandler()

	supporterService := registrationservices.NewSupporterService(DBCon)
	testSupportersHandler := NewSupportersHandler(supporterService)

	companyService := registrationservices.NewCompanyService(DBCon)
	testCompaniesHandler := NewCompaniesHandler(companyService)

	mainHandler := NewMainHandler(csrfServer, testSupportersHandler, testCompaniesHandler)

	strictHandler := registrationapi.NewStrictHandler(mainHandler, nil)
	registrationapi.RegisterHandlers(e, strictHandler)
}

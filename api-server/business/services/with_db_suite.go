package businessservices

import (
	"apps/database"
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type WithDBSuite struct {
	suite.Suite
}

var DBCon *gorm.DB

// func (s *WithDBSuite) SetupSuite()                           {} // テストスイート実施前の処理
// func (s *WithDBSuite) TearDownSuite()                        {} // テストスイート終了後の処理
// func (s *WithDBSuite) SetupTest()                            {} // テストケース実施前の処理
// func (s *WithDBSuite) TearDownTest()                         {} // テストケース終了後の処理
// func (s *WithDBSuite) BeforeTest(suiteName, testName string) {} // テストケース実施前の処理
// func (s *WithDBSuite) AfterTest(suiteName, testName string)  {} // テストケース終了後の処理

func init() {
	txdb.Register("txdb-service", "mysql", database.GetDsn())
}

func (s *WithDBSuite) SetDBCon() {
	db, err := sql.Open("txdb-service", "connect")
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

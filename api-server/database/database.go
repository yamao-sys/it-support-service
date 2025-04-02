package database

import (
	"crypto/tls"
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	// DBインスタンス生成
	db, err := sql.Open("mysql", GetDsn())
	if err != nil {
		panic(err)
	}
	return db
}

func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func GetDsn() string {
	baseDsn := os.Getenv("MYSQL_USER") +
				":" + os.Getenv("MYSQL_PASS") +
				"@tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")/" +
				os.Getenv("MYSQL_DBNAME") +
				"?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true&loc=Local"
	
	if os.Getenv("APP_ENV") != "production" {
		return baseDsn
	}
	
	err := mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: os.Getenv("MYSQL_HOST"),
	})
	if err != nil {
		log.Fatalf("failed to register TLS config: %v", err)
	}
	return baseDsn+"&tls=tidb"
}

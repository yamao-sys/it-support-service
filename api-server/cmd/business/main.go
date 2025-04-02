package main

import (
	businessapi "apps/api/business"
	businesshandlers "apps/business/handlers"
	"apps/business/middlewares"
	businessservices "apps/business/services"
	"apps/database"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// NOTE: デプロイ先の環境はSecret Managerで環境変数を管理する
	if os.Getenv("APP_ENV") != "production" {
		loadEnv()
	}

	dbCon := database.Init()

	// NOTE: service層のインスタンス
	supporterService := businessservices.NewSupporterService(dbCon)
	companyService := businessservices.NewCompanyService(dbCon)

	// NOTE: Handlerのインスタンス
	csrfHandler := businesshandlers.NewCsrfHandler()
	supportersHandler := businesshandlers.NewSupportersHandler(supporterService)
	companiesHandler := businesshandlers.NewCompaniesHandler(companyService)

	// NOTE: Handlerをルーティングに追加
	e := middlewares.ApplyMiddlewares(echo.New())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Registration!")
	})
	mainHandler := businesshandlers.NewMainHandler(csrfHandler, supportersHandler, companiesHandler)
	mainStrictHandler := businessapi.NewStrictHandler(mainHandler, nil)
	businessapi.RegisterHandlers(e, mainStrictHandler)

	if err := e.Start(":" + os.Getenv("BUSINESS_SERVER_PORT")); err != nil && err != http.ErrServerClosed {
		e.Logger.Errorf("Echo server error: %v", err)
	}
}

func loadEnv() {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}
	godotenv.Load(envFilePath)
}

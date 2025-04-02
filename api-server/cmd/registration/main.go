package main

import (
	registrationapi "apps/api/registration"
	"apps/database"
	registrationhandlers "apps/registration/handlers"
	registrationmiddlewares "apps/registration/middlewares"
	registrationservices "apps/registration/services"
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
	supporterService := registrationservices.NewSupporterService(dbCon)
	companyService := registrationservices.NewCompanyService(dbCon)

	// NOTE: Handlerのインスタンス
	csrfHandler := registrationhandlers.NewCsrfHandler()
	supportersHandler := registrationhandlers.NewSupportersHandler(supporterService)
	companiesHandler := registrationhandlers.NewCompaniesHandler(companyService)

	// NOTE: Handlerをルーティングに追加
	e := registrationmiddlewares.ApplyMiddlewares(echo.New())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Registration!")
	})
	mainHandler := registrationhandlers.NewMainHandler(csrfHandler, supportersHandler, companiesHandler)
	mainStrictHandler := registrationapi.NewStrictHandler(mainHandler, nil)
	registrationapi.RegisterHandlers(e, mainStrictHandler)

	if err := e.Start(":" + os.Getenv("SERVER_PORT")); err != nil && err != http.ErrServerClosed {
		e.Logger.Errorf("Echo server error: %v", err)
	}
}

func loadEnv() {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = "/app/.env"
	}
	godotenv.Load(envFilePath)
}

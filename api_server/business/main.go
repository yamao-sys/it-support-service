package main

import (
	"business/api/generated/companies"
	"business/api/generated/csrf"
	"business/api/generated/supporters"
	"business/internal/database"
	"business/internal/handlers"
	"business/internal/middlewares"
	"business/internal/services"
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
	supporterService := services.NewSupporterService(dbCon)
	companyService := services.NewCompanyService(dbCon)

	// NOTE: controllerをHandlerに追加
	supporterServer := handlers.NewSupportersHandler(supporterService)
	supporterStrictHandler := supporters.NewStrictHandler(supporterServer, nil)
	companyServer := handlers.NewCompaniesHandler(companyService)
	companyStrictHandler := companies.NewStrictHandler(companyServer, nil)

	csrfServer := handlers.NewCsrfHandler()
	csrfStrictHandler := csrf.NewStrictHandler(csrfServer, nil)

	// NOTE: Handlerをルーティングに追加
	e := middlewares.ApplyMiddlewares(echo.New())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Registration!")
	})
	supporters.RegisterHandlers(e, supporterStrictHandler)
	companies.RegisterHandlers(e, companyStrictHandler)
	csrf.RegisterHandlers(e, csrfStrictHandler)

	if err := e.Start(":" + os.Getenv("SERVER_PORT")); err != nil && err != http.ErrServerClosed {
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

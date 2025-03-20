package main

import (
	"net/http"
	"os"
	"registration/api/generated/csrf"
	"registration/api/generated/supporters"
	"registration/internal/database"
	"registration/internal/handlers"
	"registration/internal/middlewares"
	"registration/internal/services"

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

	// NOTE: controllerをHandlerに追加
	supporterServer := handlers.NewSupportersHandler(supporterService)
	supporterStrictHandler := supporters.NewStrictHandler(supporterServer, nil)

	csrfServer := handlers.NewCsrfHandler()
	csrfStrictHandler := csrf.NewStrictHandler(csrfServer, nil)

	// NOTE: Handlerをルーティングに追加
	e := middlewares.ApplyMiddlewares(echo.New())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Registration!")
	})
	supporters.RegisterHandlers(e, supporterStrictHandler)
	csrf.RegisterHandlers(e, csrfStrictHandler)

	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}

func loadEnv() {
	// TODO: godotenvで読み込むのはdevelopmentとtest環境のみにする
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}
	godotenv.Load(envFilePath)
}

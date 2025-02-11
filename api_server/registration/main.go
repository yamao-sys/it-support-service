package main

import (
	"net/http"
	"os"
	"registration/api/generated/supporters"
	"registration/internal/database"
	"registration/internal/handlers"
	"registration/internal/services"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()

	dbCon := database.Init()

	// NOTE: service層のインスタンス
	supporterService := services.NewSupporterService(dbCon)

	// NOTE: controllerをHandlerに追加
	supporterServer := handlers.NewSupportersHandler(supporterService)
	supporterStrictHandler := supporters.NewStrictHandler(supporterServer, nil)

	// NOTE: Handlerをルーティングに追加
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Registration!")
	})
	supporters.RegisterHandlers(e, supporterStrictHandler)

	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}

func loadEnv() {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}
	godotenv.Load(envFilePath)
}

package main

import (
	admin_router "apps/apps/admin/router"
	business_router "apps/apps/business/router"
	registration_router "apps/apps/registration/router"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()

	// dbCon := d.Init()

	// NOTE: Handlerをルーティングに追加
	e := echo.New()
	switch os.Getenv("APP_MODE") {
	case "registration":
		e = registration_router.SetupRouting(e)
	case "business":
		e = business_router.SetupRouting(e)
	case "admin":
		e = admin_router.SetupRouting(e)
	}

	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}

func loadEnv() {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}
	godotenv.Load(envFilePath)
}

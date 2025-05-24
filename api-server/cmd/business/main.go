package main

import (
	businessapi "apps/api/business"
	businesshandlers "apps/business/handlers"
	businessmiddlewares "apps/business/middlewares"
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
	projectService := businessservices.NewProjectService(dbCon)
	planService := businessservices.NewPlanService(dbCon)
	toProjectService := businessservices.NewToProjectService(dbCon)

	// NOTE: Handlerのインスタンス
	csrfHandler := businesshandlers.NewCsrfHandler()
	supportersHandler := businesshandlers.NewSupportersHandler(supporterService)
	companiesHandler := businesshandlers.NewCompaniesHandler(companyService)
	projectsHandler := businesshandlers.NewProjectsHandler(projectService)
	plansHandler := businesshandlers.NewPlansHandler(planService)
	toProjectsHandler := businesshandlers.NewToProjectsHandler(toProjectService)

	// NOTE: Handlerをルーティングに追加
	e := businessmiddlewares.ApplyMiddlewares(echo.New())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Registration!")
	})
	mainHandler := businesshandlers.NewMainHandler(csrfHandler, supportersHandler, companiesHandler, projectsHandler, plansHandler, toProjectsHandler)
	mainStrictHandler := businessapi.NewStrictHandler(mainHandler, []businessapi.StrictMiddlewareFunc{businessmiddlewares.AuthMiddleware})
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

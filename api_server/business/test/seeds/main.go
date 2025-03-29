package main

import (
	"business/internal/database"
	models "business/models/generated"
	"business/test/factories"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	loadEnv()

	dbCon := database.Init()
	// NOTE: DBを閉じる
	defer func(cause error) {
		fmt.Println(cause)
		if cause = dbCon.Close(); cause != nil {
			panic(cause)
		}
	}(nil)

	// NOTE: ログイン用サポータの追加
	supporter := factories.SupporterFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Supporter)
	if err := supporter.Insert(context.Background(), dbCon, boil.Infer()); err != nil {
		fmt.Println("failed to create test supporter", err)
	}

	// NOTE: ログイン用企業の追加
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	if err := company.Insert(context.Background(), dbCon, boil.Infer()); err != nil {
		fmt.Println("failed to create test company", err)
	}
}

func loadEnv() {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}
	godotenv.Load(envFilePath)
}

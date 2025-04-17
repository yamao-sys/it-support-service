package main

import (
	"apps/database"
	models "apps/models/generated"
	"apps/test/factories"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/volatiletech/null/v8"
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

	// NOTE: ログイン企業の案件の追加
	emptyBudgetProject1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	havingBudgetProject1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	emptyBudgetProject2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}, "IsActive": false}).(*models.Project)
	havingBudgetProject2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "IsActive": false}).(*models.Project)
	var projects models.ProjectSlice
	projects = append(projects, emptyBudgetProject1, havingBudgetProject1, emptyBudgetProject2, havingBudgetProject2)
	// NOTE: 無限スクロールのテスト用に11になるように登録する(案件追加のテストでさらに4件追加されるので、最大15個になるように)
	for range 6{
        projects = append(projects, factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project))
	}
	if _, err := projects.InsertAll(context.Background(), dbCon, boil.Infer()); err != nil {
		fmt.Println("failed to create projects", err)
	}
}

func loadEnv() {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}
	godotenv.Load(envFilePath)
}

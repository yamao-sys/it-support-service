package main

import (
	"apps/database"
	models "apps/models"
	"apps/test/factories"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/volatiletech/null/v8"
)

func main() {
	loadEnv()

	dbCon := database.Init()
	// NOTE: DBを閉じる
	defer func(cause error) {
		fmt.Println(cause)
		database.Close(dbCon)
	}(nil)

	// NOTE: ログイン用サポータの追加
	supporter := factories.SupporterFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Supporter)
	dbCon.Create(supporter)

	// NOTE: ログイン用企業の追加
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	dbCon.Create(company)

	// NOTE: ログイン企業の案件の追加
	emptyBudgetProject1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	dbCon.Create(emptyBudgetProject1)
	havingBudgetProject1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	dbCon.Create(havingBudgetProject1)
	emptyBudgetProject2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}, "IsActive": false}).(*models.Project)
	dbCon.Create(emptyBudgetProject2)
	havingBudgetProject2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "IsActive": false}).(*models.Project)
	dbCon.Create(havingBudgetProject2)
	var projects []models.Project
	// NOTE: 無限スクロールのテスト用に11になるように登録する(案件追加のテストでさらに4件追加されるので、最大15個になるように)
	for range 6{
		project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
        projects = append(projects, *project)
	}
	dbCon.CreateInBatches(projects, len(projects))

	// NOTE: 案件に紐づく支援計画の追加
	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": emptyBudgetProject1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	dbCon.Create(temporaryCreatingPlan)
	
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": havingBudgetProject1.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	dbCon.Create(submittedPlan)
	
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": emptyBudgetProject2.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	dbCon.Create(agreedPlan)
}

func loadEnv() {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}
	godotenv.Load(envFilePath)
}

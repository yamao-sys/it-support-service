package services

import (
	"business/api/generated/projects"
	models "business/models/generated"
	"business/test/factories"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TestProjectServiceSuite struct {
	WithDBSuite
}

var testProjectService ProjectService

func (s *TestProjectServiceSuite) SetupTest() {
	s.SetDBCon()

	testProjectService = NewProjectService(DBCon)
}

func (s *TestProjectServiceSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestProjectServiceSuite) TestProjectCreate_StatusOK() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	if err := company.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test company %v", err)
	}

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	requestParams := projects.PostProjectsJSONRequestBody{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}

	createdProject, validatorErrors, err := testProjectService.Create(ctx, &requestParams, company.ID)

	assert.Equal(s.T(), company.ID, createdProject.CompanyID)
	assert.Equal(s.T(), "test title", createdProject.Title)
	assert.Equal(s.T(), "test description", createdProject.Description)
	assert.Equal(s.T(), parsedStartDate, createdProject.StartDate)
	assert.Equal(s.T(), parsedEndDate, createdProject.EndDate)
	assert.Equal(s.T(), null.Int{Int: minBudget, Valid: true}, createdProject.MinBudget)
	assert.Equal(s.T(), null.Int{Int: maxBudget, Valid: true}, createdProject.MaxBudget)
	assert.Equal(s.T(), isActive, createdProject.IsActive)
	assert.Nil(s.T(), validatorErrors)
	assert.Nil(s.T(), err)
}

func (s *TestProjectServiceSuite) TestProjectCreate_BadRequest_Required() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	if err := company.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test company %v", err)
	}

	title := ""
	description := ""
	minBudget := 10000
	maxBudget := 20000
	requestParams := projects.PostProjectsJSONRequestBody{Title: &title, Description: &description, StartDate: nil, EndDate: nil, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: nil}

	createdProject, validatorErrors, err := testProjectService.Create(ctx, &requestParams, company.ID)
	mappedValidationErrors := testProjectService.MappingValidationErrorStruct(validatorErrors)
	
	// NOTE: projectが作られていないこと
	assert.Equal(s.T(), 0, createdProject.ID)

	titleErrorMessages := []string{"案件タイトルは必須入力です。"}
	descriptionErrorMessages := []string{"案件概要は必須入力です。"}
	startDateErrorMessages := []string{"案件開始日は必須入力です。"}
	endDateErrorMessages := []string{"案件終了日は必須入力です。"}
	isActiveErrorMessages := []string{"公開フラグは必須入力です。"}
	assert.ElementsMatch(s.T(), titleErrorMessages, *mappedValidationErrors.Title)
	assert.ElementsMatch(s.T(), descriptionErrorMessages, *mappedValidationErrors.Description)
	assert.ElementsMatch(s.T(), startDateErrorMessages, *mappedValidationErrors.StartDate)
	assert.ElementsMatch(s.T(), endDateErrorMessages, *mappedValidationErrors.EndDate)
	assert.ElementsMatch(s.T(), isActiveErrorMessages, *mappedValidationErrors.IsActive)

	assert.Nil(s.T(), err)
}

func (s *TestProjectServiceSuite) TestProjectCreate_BadRequest_Greater() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	if err := company.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test company %v", err)
	}

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10001
	maxBudget := 10000
	isActive := false
	requestParams := projects.PostProjectsJSONRequestBody{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}

	createdProject, validatorErrors, err := testProjectService.Create(ctx, &requestParams, company.ID)
	mappedValidationErrors := testProjectService.MappingValidationErrorStruct(validatorErrors)

	// NOTE: projectが作られていないこと
	assert.Equal(s.T(), 0, createdProject.ID)

	endDateErrorMessages := []string{"案件終了日と案件開始日の前後関係が不適です。"}
	maxBudgetErrorMessages := []string{"予算(上限)と予算(下限)の大小関係が不適です。"}
	assert.ElementsMatch(s.T(), endDateErrorMessages, *mappedValidationErrors.EndDate)
	assert.ElementsMatch(s.T(), maxBudgetErrorMessages, *mappedValidationErrors.MaxBudget)

	assert.Nil(s.T(), err)
}

func (s *TestProjectServiceSuite) TestProjectCreate_BadRequest_Threshold() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	if err := company.Insert(ctx, DBCon, boil.Infer()); err != nil {
		s.T().Fatalf("failed to create test company %v", err)
	}

	title := randomdata.RandStringRunes(71)
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := -1
	maxBudget := 0
	isActive := false
	requestParams := projects.PostProjectsJSONRequestBody{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}

	createdProject, validatorErrors, err := testProjectService.Create(ctx, &requestParams, company.ID)
	mappedValidationErrors := testProjectService.MappingValidationErrorStruct(validatorErrors)

	// NOTE: projectが作られていないこと
	assert.Equal(s.T(), 0, createdProject.ID)

	titleErrorMessages := []string{"案件タイトルは1 ~ 70文字での入力をお願いします。"}
	minBudgetErrorMessages := []string{"予算(下限)は正の整数での入力をお願いいたします。"}
	maxBudgetErrorMessages := []string{"予算(上限)は正の整数での入力をお願いいたします。"}
	assert.ElementsMatch(s.T(), titleErrorMessages, *mappedValidationErrors.Title)
	assert.ElementsMatch(s.T(), minBudgetErrorMessages, *mappedValidationErrors.MinBudget)
	assert.ElementsMatch(s.T(), maxBudgetErrorMessages, *mappedValidationErrors.MaxBudget)

	assert.Nil(s.T(), err)
}

func TestProjectService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestProjectServiceSuite))
}

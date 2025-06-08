package businessservices

import (
	businessapi "apps/api/business"
	models "apps/models"
	"apps/test/factories"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
)

type TestPlanServiceSuite struct {
	WithDBSuite
}

var testPlanService PlanService

func (s *TestPlanServiceSuite) SetupTest() {
	s.SetDBCon()

	testPlanService = NewPlanService(DBCon)
}

func (s *TestPlanServiceSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestPlanServiceSuite) TestPlanCreate_RequiredOnly_Success() {
	// NOTE: テスト用Company, Project, Supporterの作成
	company := factories.CompanyFactory.MustCreate().(*models.Company)
	DBCon.Create(company)
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	supporter := factories.SupporterFactory.MustCreate().(*models.Supporter)
	DBCon.Create(supporter)

	title := randomdata.RandStringRunes(70)
	description := "test description"
	unitPrice := 10000
	requestParams := businessapi.PlanStoreInput{ProjectId: project.ID, Title: title, Description: description, UnitPrice: unitPrice}

	createdPlan, validatorErrors, err := testPlanService.Create(&requestParams, supporter.ID)
	mappedValidationErrors := testPlanService.MappingValidationErrorStruct(validatorErrors)
	expectedValidationErrors := businessapi.PlanValidationError{}
	assert.Equal(s.T(), expectedValidationErrors, mappedValidationErrors)

	assert.Equal(s.T(), supporter.ID, createdPlan.SupporterID)
	assert.Equal(s.T(), project.ID, createdPlan.ProjectID)
	assert.Equal(s.T(), title, createdPlan.Title)
	assert.Equal(s.T(), description, createdPlan.Description)
	assert.Equal(s.T(), null.Time{Time: time.Time{}, Valid: false}, createdPlan.StartDate)
	assert.Equal(s.T(), null.Time{Time: time.Time{}, Valid: false}, createdPlan.EndDate)
	assert.Equal(s.T(), unitPrice, createdPlan.UnitPrice)
	assert.Equal(s.T(), null.Time{Time:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Valid:false}, createdPlan.AgreedAt)
	assert.Nil(s.T(), validatorErrors)
	assert.Nil(s.T(), err)
}

func (s *TestPlanServiceSuite) TestPlanCreate_WithOptional_Success() {
	// NOTE: テスト用Company, Project, Supporterの作成
	company := factories.CompanyFactory.MustCreate().(*models.Company)
	DBCon.Create(company)
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	supporter := factories.SupporterFactory.MustCreate().(*models.Supporter)
	DBCon.Create(supporter)

	title := randomdata.RandStringRunes(70)
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 9, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	unitPrice := 10000
	requestParams := businessapi.PlanStoreInput{ProjectId: project.ID, Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, UnitPrice: unitPrice}

	createdPlan, validatorErrors, err := testPlanService.Create(&requestParams, supporter.ID)
	mappedValidationErrors := testPlanService.MappingValidationErrorStruct(validatorErrors)
	expectedValidationErrors := businessapi.PlanValidationError{}
	assert.Equal(s.T(), expectedValidationErrors, mappedValidationErrors)

	assert.Equal(s.T(), supporter.ID, createdPlan.SupporterID)
	assert.Equal(s.T(), project.ID, createdPlan.ProjectID)
	assert.Equal(s.T(), title, createdPlan.Title)
	assert.Equal(s.T(), description, createdPlan.Description)
	assert.Equal(s.T(), null.Time{Time: time.Date(2025, 4, 9, 0, 0, 0, 0, time.Local), Valid: true}, createdPlan.StartDate)
	assert.Equal(s.T(), null.Time{Time: time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local), Valid: true}, createdPlan.EndDate)
	assert.Equal(s.T(), unitPrice, createdPlan.UnitPrice)
	assert.Equal(s.T(), null.Time{Time:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Valid:false}, createdPlan.AgreedAt)
	assert.Nil(s.T(), validatorErrors)
	assert.Nil(s.T(), err)
}

func (s *TestPlanServiceSuite) TestPlanCreate_BadRequest_Required() {
	// NOTE: テスト用Company, Project, Supporterの作成
	company := factories.CompanyFactory.MustCreate().(*models.Company)
	DBCon.Create(company)
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	supporter := factories.SupporterFactory.MustCreate().(*models.Supporter)
	DBCon.Create(supporter)

	title := ""
	description := ""
	requestParams := businessapi.PlanStoreInput{ProjectId: project.ID, Title: title, Description: description}

	createdPlan, validatorErrors, err := testPlanService.Create(&requestParams, supporter.ID)
	mappedValidationErrors := testPlanService.MappingValidationErrorStruct(validatorErrors)

	// NOTE: planが作られていないこと
	assert.Equal(s.T(), 0, createdPlan.ID)

	titleErrorMessages := []string{"支援タイトルは必須入力です。"}
	descriptionErrorMessages := []string{"支援概要は必須入力です。"}
	unitPriceErrorMessages := []string{"支援単価(税抜)は必須入力です。"}
	assert.ElementsMatch(s.T(), titleErrorMessages, *mappedValidationErrors.Title)
	assert.ElementsMatch(s.T(), descriptionErrorMessages, *mappedValidationErrors.Description)
	assert.ElementsMatch(s.T(), unitPriceErrorMessages, *mappedValidationErrors.UnitPrice)

	assert.Nil(s.T(), err)
}

func (s *TestPlanServiceSuite) TestPlanCreate_BadRequest_Greater() {
	// NOTE: テスト用Company, Project, Supporterの作成
	company := factories.CompanyFactory.MustCreate().(*models.Company)
	DBCon.Create(company)
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	supporter := factories.SupporterFactory.MustCreate().(*models.Supporter)
	DBCon.Create(supporter)

	title := randomdata.RandStringRunes(70)
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 9, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	unitPrice := 10000
	requestParams := businessapi.PlanStoreInput{ProjectId: project.ID, Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, UnitPrice: unitPrice}

	createdPlan, validatorErrors, err := testPlanService.Create(&requestParams, supporter.ID)
	mappedValidationErrors := testPlanService.MappingValidationErrorStruct(validatorErrors)

	// NOTE: planが作られていないこと
	assert.Equal(s.T(), 0, createdPlan.ID)

	endDateErrorMessages := []string{"支援終了日と支援開始日の前後関係が不適です。"}
	assert.ElementsMatch(s.T(), endDateErrorMessages, *mappedValidationErrors.EndDate)

	assert.Nil(s.T(), err)
}

func (s *TestPlanServiceSuite) TestPlanCreate_BadRequest_Threshold() {
	// NOTE: テスト用Company, Project, Supporterの作成
	company := factories.CompanyFactory.MustCreate().(*models.Company)
	DBCon.Create(company)
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	supporter := factories.SupporterFactory.MustCreate().(*models.Supporter)
	DBCon.Create(supporter)

	title := randomdata.RandStringRunes(71)
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 9, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	unitPrice := 10000
	requestParams := businessapi.PlanStoreInput{ProjectId: project.ID, Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, UnitPrice: unitPrice}

	createdPlan, validatorErrors, err := testPlanService.Create(&requestParams, supporter.ID)
	mappedValidationErrors := testPlanService.MappingValidationErrorStruct(validatorErrors)

	// NOTE: planが作られていないこと
	assert.Equal(s.T(), 0, createdPlan.ID)

	titleErrorMessages := []string{"案件タイトルは1 ~ 70文字での入力をお願いします。"}
	assert.ElementsMatch(s.T(), titleErrorMessages, *mappedValidationErrors.Title)

	assert.Nil(s.T(), err)
}

func TestPlanService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestPlanServiceSuite))
}

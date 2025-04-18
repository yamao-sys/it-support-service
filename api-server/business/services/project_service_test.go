package businessservices

import (
	businessapi "apps/api/business"
	models "apps/models/generated"
	"apps/test/factories"
	"errors"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func (s *TestProjectServiceSuite) TestProjectFetchLists_NotHavingNextPage_StatusOK() {
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())

	project1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project1.Insert(ctx, DBCon, boil.Infer())
	project2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project2.Insert(ctx, DBCon, boil.Infer())
	project3 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project3.Insert(ctx, DBCon, boil.Infer())
	project4 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project4.Insert(ctx, DBCon, boil.Infer())
	project5 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project5.Insert(ctx, DBCon, boil.Infer())
	companyProductIDs, _ := models.Projects(
		qm.Select("projects.id"),
		qm.Where("company_id = ?", company.ID),
	).All(ctx, DBCon)

	otherCompany := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.Company)
	otherCompany.Insert(ctx, DBCon, boil.Infer())
	otherCompanyProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": otherCompany.ID}).(*models.Project)
	otherCompanyProduct.Insert(ctx, DBCon, boil.Infer())

	fetchedProducts, nextPageToken, err := testProjectService.FetchLists(ctx, company.ID, 0)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs.GetIDs(), fetchedProductIDs)
	assert.Equal(s.T(), 0, nextPageToken)
	assert.Nil(s.T(), err)
}

func (s *TestProjectServiceSuite) TestProjectFetchLists_HavingNextPage_StatusOK() {
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())

	project1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project1.Insert(ctx, DBCon, boil.Infer())
	project2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project2.Insert(ctx, DBCon, boil.Infer())
	project3 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project3.Insert(ctx, DBCon, boil.Infer())
	project4 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project4.Insert(ctx, DBCon, boil.Infer())
	project5 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project5.Insert(ctx, DBCon, boil.Infer())
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project6.Insert(ctx, DBCon, boil.Infer())

	companyProductIDs, _ := models.Projects(
		qm.Select("projects.id"),
		qm.Where("company_id = ?", company.ID),
		qm.Where("id != ?", project6.ID),
	).All(ctx, DBCon)

	otherCompany := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.Company)
	otherCompany.Insert(ctx, DBCon, boil.Infer())
	otherCompanyProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": otherCompany.ID}).(*models.Project)
	otherCompanyProduct.Insert(ctx, DBCon, boil.Infer())

	fetchedProducts, nextPageToken, err := testProjectService.FetchLists(ctx, company.ID, 0)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs.GetIDs(), fetchedProductIDs)
	assert.Equal(s.T(), project6.ID, nextPageToken)
	assert.Nil(s.T(), err)
}

func (s *TestProjectServiceSuite) TestProjectFetchLists_WithPageToken_NotHavingNextPage_StatusOK() {
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())

	project1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project1.Insert(ctx, DBCon, boil.Infer())
	project2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project2.Insert(ctx, DBCon, boil.Infer())
	project3 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project3.Insert(ctx, DBCon, boil.Infer())
	project4 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project4.Insert(ctx, DBCon, boil.Infer())
	project5 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project5.Insert(ctx, DBCon, boil.Infer())
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project6.Insert(ctx, DBCon, boil.Infer())
	companyProductIDs, _ := models.Projects(
		qm.Select("projects.id"),
		qm.Where("company_id = ?", company.ID),
		qm.Where("id >= ?", project2.ID),
	).All(ctx, DBCon)

	otherCompany := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.Company)
	otherCompany.Insert(ctx, DBCon, boil.Infer())
	otherCompanyProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": otherCompany.ID}).(*models.Project)
	otherCompanyProduct.Insert(ctx, DBCon, boil.Infer())

	fetchedProducts, nextPageToken, err := testProjectService.FetchLists(ctx, company.ID, project2.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs.GetIDs(), fetchedProductIDs)
	assert.Equal(s.T(), 0, nextPageToken)
	assert.Nil(s.T(), err)
}

func (s *TestProjectServiceSuite) TestProjectFetchLists_WithPageToken_HavingNextPage_StatusOK() {
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())

	project1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project1.Insert(ctx, DBCon, boil.Infer())
	project2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project2.Insert(ctx, DBCon, boil.Infer())
	project3 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project3.Insert(ctx, DBCon, boil.Infer())
	project4 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project4.Insert(ctx, DBCon, boil.Infer())
	project5 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project5.Insert(ctx, DBCon, boil.Infer())
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project6.Insert(ctx, DBCon, boil.Infer())
	project7 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project7.Insert(ctx, DBCon, boil.Infer())
	companyProductIDs, _ := models.Projects(
		qm.Select("projects.id"),
		qm.Where("company_id = ?", company.ID),
		qm.Where("id >= ?", project2.ID),
		qm.Where("id < ?", project7.ID),
	).All(ctx, DBCon)

	otherCompany := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.Company)
	otherCompany.Insert(ctx, DBCon, boil.Infer())
	otherCompanyProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": otherCompany.ID}).(*models.Project)
	otherCompanyProduct.Insert(ctx, DBCon, boil.Infer())

	fetchedProducts, nextPageToken, err := testProjectService.FetchLists(ctx, company.ID, project2.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs.GetIDs(), fetchedProductIDs)
	assert.Equal(s.T(), project7.ID, nextPageToken)
	assert.Nil(s.T(), err)
}

func (s *TestProjectServiceSuite) TestProjectCreate_WithOnlyRequired_StatusOK() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	isActive := true
	requestParams := businessapi.ProjectStoreInput{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, IsActive: &isActive}

	createdProject, validatorErrors, err := testProjectService.Create(ctx, &requestParams, company.ID)
	mappedValidationErrors := testProjectService.MappingValidationErrorStruct(validatorErrors)
	expectedValidationErrors := businessapi.ProjectValidationError{}
	assert.Equal(s.T(), expectedValidationErrors, mappedValidationErrors)

	assert.Equal(s.T(), company.ID, createdProject.CompanyID)
	assert.Equal(s.T(), "test title", createdProject.Title)
	assert.Equal(s.T(), "test description", createdProject.Description)
	assert.Equal(s.T(), parsedStartDate, createdProject.StartDate)
	assert.Equal(s.T(), parsedEndDate, createdProject.EndDate)
	assert.Equal(s.T(), null.Int(null.Int{Int:0, Valid:false}), createdProject.MinBudget)
	assert.Equal(s.T(), null.Int(null.Int{Int:0, Valid:false}), createdProject.MaxBudget)
	assert.Equal(s.T(), isActive, createdProject.IsActive)
	assert.Nil(s.T(), validatorErrors)
	assert.Nil(s.T(), err)
}

func (s *TestProjectServiceSuite) TestProjectCreate_WithOptional_StatusOK() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	requestParams := businessapi.ProjectStoreInput{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}

	createdProject, validatorErrors, err := testProjectService.Create(ctx, &requestParams, company.ID)
	mappedValidationErrors := testProjectService.MappingValidationErrorStruct(validatorErrors)
	expectedValidationErrors := businessapi.ProjectValidationError{}
	assert.Equal(s.T(), expectedValidationErrors, mappedValidationErrors)

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
	company.Insert(ctx, DBCon, boil.Infer())

	title := ""
	description := ""
	minBudget := 10000
	maxBudget := 20000
	requestParams := businessapi.ProjectStoreInput{Title: &title, Description: &description, StartDate: nil, EndDate: nil, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: nil}

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
	company.Insert(ctx, DBCon, boil.Infer())

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10001
	maxBudget := 10000
	isActive := false
	requestParams := businessapi.ProjectStoreInput{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}

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
	company.Insert(ctx, DBCon, boil.Infer())

	title := randomdata.RandStringRunes(71)
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := -1
	maxBudget := 0
	isActive := false
	requestParams := businessapi.ProjectStoreInput{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}

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

func (s *TestProjectServiceSuite) TestProjectFetch_StatusOK() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	project.Reload(ctx, DBCon)

	fetchedProduct, err := testProjectService.Fetch(ctx, project.ID)
	
	// NOTE: レスポンスのprojectの値の確認
	assert.Equal(s.T(), project.ID, fetchedProduct.ID)
	assert.Equal(s.T(), company.ID, fetchedProduct.CompanyID)
	assert.Equal(s.T(), project.Title, fetchedProduct.Title)

	assert.Nil(s.T(), err)
}

func (s *TestProjectServiceSuite) TestProjectFetch_NotFound() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	project.Reload(ctx, DBCon)

	fetchedProduct, err := testProjectService.Fetch(ctx, project.ID+1)
	
	// NOTE: レスポンスのprojectの値の確認
	assert.Equal(s.T(), 0, fetchedProduct.ID)
	assert.Equal(s.T(), 0, fetchedProduct.CompanyID)
	assert.Equal(s.T(), "", fetchedProduct.Title)

	assert.Equal(s.T(), errors.New("not found"), err)
}

func (s *TestProjectServiceSuite) TestProjectUpdate_WithOnlyRequired_StatusOK() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	project.Reload(ctx, DBCon)

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	isActive := true
	requestParams := businessapi.ProjectStoreInput{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, IsActive: &isActive}

	updatedProject, validatorErrors, err := testProjectService.Update(ctx, &requestParams, project.ID)
	mappedValidationErrors := testProjectService.MappingValidationErrorStruct(validatorErrors)
	expectedValidationErrors := businessapi.ProjectValidationError{}
	assert.Equal(s.T(), expectedValidationErrors, mappedValidationErrors)

	// NOTE: レスポンスのprojectの値の確認
	assert.Equal(s.T(), project.ID, updatedProject.ID)
	assert.Equal(s.T(), company.ID, updatedProject.CompanyID)
	assert.Equal(s.T(), "test title", updatedProject.Title)
	assert.Equal(s.T(), "test description", updatedProject.Description)
	assert.Equal(s.T(), parsedStartDate, updatedProject.StartDate)
	assert.Equal(s.T(), parsedEndDate, updatedProject.EndDate)
	assert.Equal(s.T(), null.Int{Int: 0, Valid: false}, updatedProject.MinBudget)
	assert.Equal(s.T(), null.Int{Int: 0, Valid: false}, updatedProject.MaxBudget)
	assert.Equal(s.T(), isActive, updatedProject.IsActive)
	assert.Nil(s.T(), validatorErrors)
	assert.Nil(s.T(), err)

	// NOTE: DBの値が更新されていること
	project.Reload(ctx, DBCon)
	assert.Equal(s.T(), company.ID, project.CompanyID)
	assert.Equal(s.T(), "test title", project.Title)
	assert.Equal(s.T(), "test description", project.Description)
	assert.Equal(s.T(), parsedStartDate, project.StartDate)
	assert.Equal(s.T(), parsedEndDate, project.EndDate)
	assert.Equal(s.T(), null.Int{Int: 0, Valid: false}, project.MinBudget)
	assert.Equal(s.T(), null.Int{Int: 0, Valid: false}, project.MaxBudget)
	assert.Equal(s.T(), isActive, project.IsActive)
}

func (s *TestProjectServiceSuite) TestProjectUpdate_WithOptional_StatusOK() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	project.Reload(ctx, DBCon)

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	requestParams := businessapi.ProjectStoreInput{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}

	updatedProject, validatorErrors, err := testProjectService.Update(ctx, &requestParams, project.ID)
	mappedValidationErrors := testProjectService.MappingValidationErrorStruct(validatorErrors)
	expectedValidationErrors := businessapi.ProjectValidationError{}
	assert.Equal(s.T(), expectedValidationErrors, mappedValidationErrors)

	// NOTE: レスポンスのprojectの値の確認
	assert.Equal(s.T(), project.ID, updatedProject.ID)
	assert.Equal(s.T(), company.ID, updatedProject.CompanyID)
	assert.Equal(s.T(), "test title", updatedProject.Title)
	assert.Equal(s.T(), "test description", updatedProject.Description)
	assert.Equal(s.T(), parsedStartDate, updatedProject.StartDate)
	assert.Equal(s.T(), parsedEndDate, updatedProject.EndDate)
	assert.Equal(s.T(), null.Int{Int: minBudget, Valid: true}, updatedProject.MinBudget)
	assert.Equal(s.T(), null.Int{Int: maxBudget, Valid: true}, updatedProject.MaxBudget)
	assert.Equal(s.T(), isActive, updatedProject.IsActive)
	assert.Nil(s.T(), validatorErrors)
	assert.Nil(s.T(), err)

	// NOTE: DBの値が更新されていること
	project.Reload(ctx, DBCon)
	assert.Equal(s.T(), company.ID, project.CompanyID)
	assert.Equal(s.T(), "test title", project.Title)
	assert.Equal(s.T(), "test description", project.Description)
	assert.Equal(s.T(), parsedStartDate, project.StartDate)
	assert.Equal(s.T(), parsedEndDate, project.EndDate)
	assert.Equal(s.T(), null.Int{Int: minBudget, Valid: true}, project.MinBudget)
	assert.Equal(s.T(), null.Int{Int: maxBudget, Valid: true}, project.MaxBudget)
	assert.Equal(s.T(), isActive, project.IsActive)
}

func (s *TestProjectServiceSuite) TestProjectUpdate_BadRequest_Required() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	project.Reload(ctx, DBCon)

	title := ""
	description := ""
	minBudget := 10000
	maxBudget := 20000
	requestParams := businessapi.ProjectStoreInput{Title: &title, Description: &description, StartDate: nil, EndDate: nil, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: nil}

	updatedProject, validatorErrors, err := testProjectService.Update(ctx, &requestParams, project.ID)
	mappedValidationErrors := testProjectService.MappingValidationErrorStruct(validatorErrors)

	// NOTE: レスポンスのprojectの値の確認
	assert.Equal(s.T(), models.Project{}, updatedProject)
	
	// NOTE: DBのprojectが更新されていないこと
	project.Reload(ctx, DBCon)
	assert.NotEqual(s.T(), "", project.Title)

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

func (s *TestProjectServiceSuite) TestProjectUpdate_BadRequest_Greater() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	project.Reload(ctx, DBCon)

	title := randomdata.RandStringRunes(70)
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10001
	maxBudget := 10000
	isActive := false
	requestParams := businessapi.ProjectStoreInput{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}

	updatedProject, validatorErrors, err := testProjectService.Create(ctx, &requestParams, company.ID)
	mappedValidationErrors := testProjectService.MappingValidationErrorStruct(validatorErrors)

	// NOTE: レスポンスのprojectの値の確認
	assert.Equal(s.T(), models.Project{}, updatedProject)
	
	// NOTE: DBのprojectが更新されていないこと
	project.Reload(ctx, DBCon)
	assert.NotEqual(s.T(), title, project.Title)

	endDateErrorMessages := []string{"案件終了日と案件開始日の前後関係が不適です。"}
	maxBudgetErrorMessages := []string{"予算(上限)と予算(下限)の大小関係が不適です。"}
	assert.ElementsMatch(s.T(), endDateErrorMessages, *mappedValidationErrors.EndDate)
	assert.ElementsMatch(s.T(), maxBudgetErrorMessages, *mappedValidationErrors.MaxBudget)

	assert.Nil(s.T(), err)
}

func (s *TestProjectServiceSuite) TestProjectUpdate_BadRequest_Threshold() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	project.Reload(ctx, DBCon)

	title := randomdata.RandStringRunes(71)
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := -1
	maxBudget := 0
	isActive := false
	requestParams := businessapi.ProjectStoreInput{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}

	updatedProject, validatorErrors, err := testProjectService.Create(ctx, &requestParams, company.ID)

	mappedValidationErrors := testProjectService.MappingValidationErrorStruct(validatorErrors)
	// NOTE: レスポンスのprojectの値の確認
	assert.Equal(s.T(), models.Project{}, updatedProject)
	
	// NOTE: DBのprojectが更新されていないこと
	project.Reload(ctx, DBCon)
	assert.NotEqual(s.T(), title, project.Title)

	titleErrorMessages := []string{"案件タイトルは1 ~ 70文字での入力をお願いします。"}
	minBudgetErrorMessages := []string{"予算(下限)は正の整数での入力をお願いいたします。"}
	maxBudgetErrorMessages := []string{"予算(上限)は正の整数での入力をお願いいたします。"}
	assert.ElementsMatch(s.T(), titleErrorMessages, *mappedValidationErrors.Title)
	assert.ElementsMatch(s.T(), minBudgetErrorMessages, *mappedValidationErrors.MinBudget)
	assert.ElementsMatch(s.T(), maxBudgetErrorMessages, *mappedValidationErrors.MaxBudget)

	assert.Nil(s.T(), err)
}

func (s *TestProjectServiceSuite) TestProjectUpdate_StatusNotFound() {
	// NOTE: テスト用企業の作成
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	project.Reload(ctx, DBCon)

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	requestParams := businessapi.ProjectStoreInput{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}

	updatedProject, validatorErrors, err := testProjectService.Update(ctx, &requestParams, project.ID+1)

	// NOTE: レスポンスのprojectの値の確認
	assert.Equal(s.T(), models.Project{}, updatedProject)
	assert.Nil(s.T(), validatorErrors)
	assert.Error(s.T(), err, "not found")
}

func TestProjectService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestProjectServiceSuite))
}

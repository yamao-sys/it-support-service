package businesshandlers

import (
	businessapi "apps/api/business"
	businessservices "apps/business/services"
	models "apps/models/generated"
	"apps/test/factories"
	"context"
	"errors"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TestProjectsHandlerSuite struct {
	WithDBSuite
}

type MockProjectService struct {
    mock.Mock
}

func (m *MockProjectService) FetchLists(ctx context.Context, companyID int) (projects models.ProjectSlice, error error) {
    args := m.Called(ctx, companyID)
	return args.Get(0).(models.ProjectSlice), args.Error(1)
}

func (m *MockProjectService) Create(ctx context.Context, requestParams *businessapi.ProjectStoreInput, companyID int) (project models.Project, validatorErrors error, error error) {
    args := m.Called(ctx, requestParams, companyID)
	return args.Get(0).(models.Project), args.Error(1), args.Error(2)
}

func (m *MockProjectService) Fetch(ctx context.Context, ID int) (project models.Project, error error) {
    args := m.Called(ctx, ID)
	return args.Get(0).(models.Project), args.Error(1)
}

func (m *MockProjectService) Update(ctx context.Context, requestParams *businessapi.ProjectStoreInput, companyID int) (project models.Project, validatorErrors error, error error) {
    args := m.Called(ctx, requestParams, companyID)
	return args.Get(0).(models.Project), args.Error(1), args.Error(2)
}

func (m *MockProjectService) MappingValidationErrorStruct(err error) businessapi.ProjectValidationError {
    args := m.Called(err)
    return args.Get(0).(businessapi.ProjectValidationError)
}

func (s *TestProjectsHandlerSuite) SetupTest() {
	s.SetDBCon()

	s.initializeHandlers(businessservices.NewProjectService(DBCon))

	// NOTE: CSRFトークンのセット
	s.SetCsrfHeaderValues()
}

func (s *TestProjectsHandlerSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestProjectsHandlerSuite) TestGetProjectsFetchLists_StatusOk() {
	company, cookieString := s.companySignIn()

	var projects models.ProjectSlice
	budgetEmptyproduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	hasBudgetProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	projects = append(projects, budgetEmptyproduct, hasBudgetProduct)
	projects.InsertAll(ctx, DBCon, boil.Infer())
	companyProductIDs, _ := models.Projects(
		qm.Select("projects.id"),
		qm.Where("company_id = ?", company.ID),
	).All(ctx, DBCon)

	otherCompany := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.Company)
	otherCompany.Insert(ctx, DBCon, boil.Infer())
	otherCompanyProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": otherCompany.ID}).(*models.Project)
	otherCompanyProduct.Insert(ctx, DBCon, boil.Infer())

	result := testutil.NewRequest().Get("/projects").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		ID, _ := strconv.Atoi(*project.Id)
		projectIDs = append(projectIDs, ID)
	}
	assert.Equal(s.T(), companyProductIDs.GetIDs(), projectIDs)

	// NOTE: 予算カラムがnullの時はnull、そうでなければ値が変えること
	var resBudgetEmptyproducts []businessapi.Project
	for _, project := range res.Projects {
		if project.MinBudget == nil && project.MaxBudget == nil {
			resBudgetEmptyproducts = append(resBudgetEmptyproducts, project)
		}
	}
	assert.Len(s.T(), resBudgetEmptyproducts, 1)

	var resHaveBudgetproducts []businessapi.Project
	for _, project := range res.Projects {
		if project.MinBudget != nil && project.MaxBudget != nil {
			resHaveBudgetproducts = append(resHaveBudgetproducts, project)
		}
	}
	assert.Len(s.T(), resHaveBudgetproducts, 1)
}

func (s *TestProjectsHandlerSuite) TestGetProjectsFetchLists_StatusUnauthorized() {
	company, _ := s.companySignIn()

	var projects models.ProjectSlice
	product1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	product2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	projects = append(projects, product1, product2)
	projects.InsertAll(ctx, DBCon, boil.Infer())

	otherCompany := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.Company)
	otherCompany.Insert(ctx, DBCon, boil.Infer())
	otherCompanyProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": otherCompany.ID}).(*models.Project)
	otherCompanyProduct.Insert(ctx, DBCon, boil.Infer())

	result := testutil.NewRequest().Get("/projects").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func (s *TestProjectsHandlerSuite) TestGetProjectsFetchLists_StatusInternalServerError() {
	company, cookieString := s.companySignIn()

	var projects models.ProjectSlice
	product1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	product2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	projects = append(projects, product1, product2)
	projects.InsertAll(ctx, DBCon, boil.Infer())

	otherCompany := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.Company)
	otherCompany.Insert(ctx, DBCon, boil.Infer())
	otherCompanyProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": otherCompany.ID}).(*models.Project)
	otherCompanyProduct.Insert(ctx, DBCon, boil.Infer())

	mockProjectService := new(MockProjectService)
	mockProjectService.On("FetchLists", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("int")).Return(models.ProjectSlice{}, errors.New("internal server error"))
	mockProjectService.On("Create", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("*businessapi.ProjectStoreInput"), mock.AnythingOfType("int") ).Return(models.Project{}, nil, nil)
	mockProjectService.On("Update", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("*businessapi.ProjectStoreInput"), mock.AnythingOfType("int") ).Return(models.Project{}, nil, nil)
	mockProjectService.On("MappingValidationErrorStruct", mock.AnythingOfType("error")).Return(businessapi.ProjectValidationError{})
	s.initializeHandlers(mockProjectService)
	result := testutil.NewRequest().Get("/projects").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusInternalServerError, result.Code())
}

func (s *TestProjectsHandlerSuite) TestPostProjectsCreate_StatusOk() {
	company, cookieString := s.companySignIn()

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	reqBody := businessapi.PostProjectsJSONRequestBody{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}
	result := testutil.NewRequest().Post("/projects").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), title, *res.Project.Title)
	assert.Equal(s.T(), description, *res.Project.Description)
	assert.Equal(s.T(), startDate, *res.Project.StartDate)
	assert.Equal(s.T(), endDate, *res.Project.EndDate)
	assert.Equal(s.T(), minBudget, *res.Project.MinBudget)
	assert.Equal(s.T(), maxBudget, *res.Project.MaxBudget)
	assert.Equal(s.T(), isActive, *res.Project.IsActive)

	expectedValidationErrors := businessapi.ProjectValidationError{}
	assert.Equal(s.T(), expectedValidationErrors, res.Errors)

	// NOTE: DBの値を確認
	exists, _ := models.Projects(
		models.ProjectWhere.CompanyID.EQ(company.ID),
		models.ProjectWhere.Title.EQ(title),
	).Exists(ctx, DBCon)
	assert.True(s.T(), exists)
}

func (s *TestProjectsHandlerSuite) TestPostProjectsCreate_StatusForbidden() {
	company, cookieString := s.companySignIn()

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	reqBody := businessapi.PostProjectsJSONRequestBody{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}
	result := testutil.NewRequest().Post("/projects").WithHeader("Cookie", cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusForbidden, result.Code())

	// NOTE: DBの値を確認
	exists, _ := models.Projects(
		models.ProjectWhere.CompanyID.EQ(company.ID),
		models.ProjectWhere.Title.EQ(title),
	).Exists(ctx, DBCon)
	assert.False(s.T(), exists)
}

func (s *TestProjectsHandlerSuite) TestPostProjectsCreate_StatusUnauthorized() {
	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	reqBody := businessapi.PostProjectsJSONRequestBody{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}
	result := testutil.NewRequest().Post("/projects").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())

	// NOTE: DBの値を確認
	exists, _ := models.Projects(
		models.ProjectWhere.Title.EQ(title),
	).Exists(ctx, DBCon)
	assert.False(s.T(), exists)
}

func (s *TestProjectsHandlerSuite) TestPostProjectsCreate_StatusInternalServerError() {
	company, cookieString := s.companySignIn()

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	reqBody := businessapi.PostProjectsJSONRequestBody{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}

	mockProjectService := new(MockProjectService)
	mockProjectService.On("FetchLists", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("int")).Return(models.ProjectSlice{}, nil)
	mockProjectService.On("Create", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("*businessapi.ProjectStoreInput"), mock.AnythingOfType("int") ).Return(models.Project{}, nil, errors.New("internal server error"))
	mockProjectService.On("Update", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("*businessapi.ProjectStoreInput"), mock.AnythingOfType("int") ).Return(models.Project{}, nil, nil)
	mockProjectService.On("MappingValidationErrorStruct", mock.AnythingOfType("error")).Return(businessapi.ProjectValidationError{})
	s.initializeHandlers(mockProjectService)

	result := testutil.NewRequest().Post("/projects").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusInternalServerError, result.Code())

	// NOTE: DBの値を確認
	exists, _ := models.Projects(
		models.ProjectWhere.CompanyID.EQ(company.ID),
		models.ProjectWhere.Title.EQ(title),
	).Exists(ctx, DBCon)
	assert.False(s.T(), exists)
}

func (s *TestProjectsHandlerSuite) TestPostProjectsCreate_BadRequest_Required() {
	company, cookieString := s.companySignIn()

	title := ""
	description := ""
	minBudget := 10000
	maxBudget := 20000
	reqBody := businessapi.PostProjectsJSONRequestBody{Title: &title, Description: &description, StartDate: nil, EndDate: nil, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: nil}
	result := testutil.NewRequest().Post("/projects").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)

	titleErrorMessages := []string{"案件タイトルは必須入力です。"}
	descriptionErrorMessages := []string{"案件概要は必須入力です。"}
	startDateErrorMessages := []string{"案件開始日は必須入力です。"}
	endDateErrorMessages := []string{"案件終了日は必須入力です。"}
	isActiveErrorMessages := []string{"公開フラグは必須入力です。"}
	assert.Equal(s.T(), titleErrorMessages, *res.Errors.Title)
	assert.Equal(s.T(), descriptionErrorMessages, *res.Errors.Description)
	assert.Equal(s.T(), startDateErrorMessages, *res.Errors.StartDate)
	assert.Equal(s.T(), endDateErrorMessages, *res.Errors.EndDate)
	assert.Equal(s.T(), isActiveErrorMessages, *res.Errors.IsActive)

	// NOTE: DBの値を確認
	exists, _ := models.Projects(
		models.ProjectWhere.CompanyID.EQ(company.ID),
		models.ProjectWhere.Title.EQ(title),
	).Exists(ctx, DBCon)
	assert.False(s.T(), exists)
}

func (s *TestProjectsHandlerSuite) TestGetProjectsId_StatusOk() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	
	result := testutil.NewRequest().Get("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), project.Title, *res.Project.Title)
	assert.Equal(s.T(), project.Description, *res.Project.Description)
	assert.Equal(s.T(), project.StartDate.Format("2006-01-02"), res.Project.StartDate.Format("2006-01-02"))
	assert.Equal(s.T(), project.EndDate.Format("2006-01-02"), res.Project.EndDate.Format("2006-01-02"))
	assert.Equal(s.T(), project.MinBudget.Int, *res.Project.MinBudget)
	assert.NotNil(s.T(), *res.Project.MinBudget)
	assert.Equal(s.T(), project.MaxBudget.Int, *res.Project.MaxBudget)
	assert.NotNil(s.T(), *res.Project.MaxBudget)
	assert.Equal(s.T(), project.IsActive, *res.Project.IsActive)
}

func (s *TestProjectsHandlerSuite) TestGetProjectsId_EmptyBudget_StatusOk() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	
	result := testutil.NewRequest().Get("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), project.Title, *res.Project.Title)
	assert.Equal(s.T(), project.Description, *res.Project.Description)
	assert.Equal(s.T(), project.StartDate.Format("2006-01-02"), res.Project.StartDate.Format("2006-01-02"))
	assert.Equal(s.T(), project.EndDate.Format("2006-01-02"), res.Project.EndDate.Format("2006-01-02"))
	assert.Nil(s.T(), res.Project.MinBudget)
	assert.Nil(s.T(), res.Project.MaxBudget)
	assert.Equal(s.T(), project.IsActive, *res.Project.IsActive)
}

func (s *TestProjectsHandlerSuite) TestGetProjectsId_StatusNotFound() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	
	result := testutil.NewRequest().Get("/projects/"+strconv.Itoa(project.ID+1)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusNotFound, result.Code())

	var res businessapi.GetProjectsId404JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), http.StatusNotFound, res.Code)
}

func (s *TestProjectsHandlerSuite) TestGetProjectsId_StatusUnauthorized() {
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	
	result := testutil.NewRequest().Get("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func (s *TestProjectsHandlerSuite) TestPutProjectsId_StatusOk() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	
	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	reqBody := businessapi.PutProjectsIdJSONRequestBody{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}
	result := testutil.NewRequest().Put("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), title, *res.Project.Title)
	assert.Equal(s.T(), description, *res.Project.Description)
	assert.Equal(s.T(), startDate.Format("2006-01-02"), res.Project.StartDate.Format("2006-01-02"))
	assert.Equal(s.T(), endDate.Format("2006-01-02"), res.Project.EndDate.Format("2006-01-02"))
	assert.Equal(s.T(), minBudget, *res.Project.MinBudget)
	assert.Equal(s.T(), maxBudget, *res.Project.MaxBudget)
	assert.Equal(s.T(), isActive, *res.Project.IsActive)

	expectedValidationErrors := businessapi.ProjectValidationError{}
	assert.Equal(s.T(), expectedValidationErrors, res.Errors)

	// NOTE: DBの値を確認
	project.Reload(ctx, DBCon)
	assert.Equal(s.T(), company.ID, project.CompanyID)
	assert.Equal(s.T(), "test title", project.Title)
	assert.Equal(s.T(), "test description", project.Description)
	assert.Equal(s.T(), parsedStartDate.Format("2006-01-02"), project.StartDate.Format("2006-01-02"))
	assert.Equal(s.T(), parsedEndDate.Format("2006-01-02"), project.EndDate.Format("2006-01-02"))
	assert.Equal(s.T(), null.Int{Int: minBudget, Valid: true}, project.MinBudget)
	assert.Equal(s.T(), null.Int{Int: maxBudget, Valid: true}, project.MaxBudget)
	assert.Equal(s.T(), isActive, project.IsActive)
}

func (s *TestProjectsHandlerSuite) TestPutProjectsId_StatusForbidden() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	reqBody := businessapi.PutProjectsIdJSONRequestBody{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}
	result := testutil.NewRequest().Put("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusForbidden, result.Code())

	// NOTE: DBの値を確認
	project.Reload(ctx, DBCon)
	assert.NotEqual(s.T(), "test title", project.Title)
}

func (s *TestProjectsHandlerSuite) TestPutProjectsId_StatusUnauthorized() {
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	company.Insert(ctx, DBCon, boil.Infer())
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())
	
	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	reqBody := businessapi.PutProjectsIdJSONRequestBody{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}
	result := testutil.NewRequest().Put("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())

	// NOTE: DBの値を確認
	project.Reload(ctx, DBCon)
	assert.NotEqual(s.T(), "test title", project.Title)
}

func (s *TestProjectsHandlerSuite) TestPutProjectsId_StatusInternalServerError() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	reqBody := businessapi.PutProjectsIdJSONRequestBody{Title: &title, Description: &description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: &isActive}

	mockProjectService := new(MockProjectService)
	mockProjectService.On("FetchLists", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("int")).Return(models.ProjectSlice{}, nil)
	mockProjectService.On("Create", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("*businessapi.ProjectStoreInput"), mock.AnythingOfType("int") ).Return(models.Project{}, nil, errors.New("internal server error"))
	mockProjectService.On("Update", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("*businessapi.ProjectStoreInput"), mock.AnythingOfType("int") ).Return(models.Project{}, nil, errors.New("internal server error"))
	mockProjectService.On("MappingValidationErrorStruct", mock.AnythingOfType("error")).Return(businessapi.ProjectValidationError{})
	s.initializeHandlers(mockProjectService)

	result := testutil.NewRequest().Put("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusInternalServerError, result.Code())

	// NOTE: DBの値を確認
	project.Reload(ctx, DBCon)
	assert.NotEqual(s.T(), "test title", project.Title)
}

func (s *TestProjectsHandlerSuite) TestPutProjectsId_BadRequest_Required() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	project.Insert(ctx, DBCon, boil.Infer())

	title := ""
	description := ""
	minBudget := 10000
	maxBudget := 20000
	reqBody := businessapi.PutProjectsIdJSONRequestBody{Title: &title, Description: &description, StartDate: nil, EndDate: nil, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: nil}
	result := testutil.NewRequest().Put("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)

	titleErrorMessages := []string{"案件タイトルは必須入力です。"}
	descriptionErrorMessages := []string{"案件概要は必須入力です。"}
	startDateErrorMessages := []string{"案件開始日は必須入力です。"}
	endDateErrorMessages := []string{"案件終了日は必須入力です。"}
	isActiveErrorMessages := []string{"公開フラグは必須入力です。"}
	assert.Equal(s.T(), titleErrorMessages, *res.Errors.Title)
	assert.Equal(s.T(), descriptionErrorMessages, *res.Errors.Description)
	assert.Equal(s.T(), startDateErrorMessages, *res.Errors.StartDate)
	assert.Equal(s.T(), endDateErrorMessages, *res.Errors.EndDate)
	assert.Equal(s.T(), isActiveErrorMessages, *res.Errors.IsActive)

	// NOTE: DBの値を確認
	project.Reload(ctx, DBCon)
	assert.NotEqual(s.T(), "", project.Title)
}

func TestProjectsHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(TestProjectsHandlerSuite))
}

package businesshandlers

import (
	businessapi "apps/api/business"
	businessservices "apps/business/services"
	models "apps/models"
	"apps/test/factories"
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
)

type TestProjectsHandlerSuite struct {
	WithDBSuite
}

type MockProjectService struct {
    mock.Mock
}

func (m *MockProjectService) FetchLists(companyID int, pageToken int) (projects []models.Project, nextPageToken int, error error) {
    args := m.Called(companyID, pageToken)
	return args.Get(0).([]models.Project), args.Int(1), args.Error(2)
}

func (m *MockProjectService) Create(requestParams *businessapi.ProjectStoreInput, companyID int) (project models.Project, validatorErrors error, error error) {
    args := m.Called(requestParams, companyID)
	return args.Get(0).(models.Project), args.Error(1), args.Error(2)
}

func (m *MockProjectService) Fetch(ID int) (project models.Project, error error) {
    args := m.Called(ID)
	return args.Get(0).(models.Project), args.Error(1)
}

func (m *MockProjectService) Update(requestParams *businessapi.ProjectStoreInput, companyID int) (project models.Project, validatorErrors error, error error) {
    args := m.Called(requestParams, companyID)
	return args.Get(0).(models.Project), args.Error(1), args.Error(2)
}

func (m *MockProjectService) MappingValidationErrorStruct(err error) businessapi.ProjectValidationError {
    args := m.Called(err)
    return args.Get(0).(businessapi.ProjectValidationError)
}

func (s *TestProjectsHandlerSuite) SetupTest() {
	s.SetDBCon()

	s.initializeHandlers(businessservices.NewProjectService(DBCon), businessservices.NewPlanService(DBCon))

	// NOTE: CSRFトークンのセット
	s.SetCsrfHeaderValues()
}

func (s *TestProjectsHandlerSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestProjectsHandlerSuite) TestGetProjectsFetchLists_StatusOk() {
	company, cookieString := s.companySignIn()

	var projects []models.Project
	budgetEmptyproduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	hasBudgetProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	projects = append(projects, *budgetEmptyproduct, *hasBudgetProduct)
	DBCon.CreateInBatches(projects, len(projects))
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Pluck("id", &companyProductIDs)

	otherCompany := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.Company)
	DBCon.Create(otherCompany)
	otherCompanyProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": otherCompany.ID}).(*models.Project)
	DBCon.Create(otherCompanyProduct)

	result := testutil.NewRequest().Get("/projects").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)

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

func (s *TestProjectsHandlerSuite) TestGetProjectsFetchLists_NotHavingNextPage_StatusOK() {
	company, cookieString := s.companySignIn()

	project1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project1)
	project2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project2)
	project3 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project3)
	project4 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project4)
	project5 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project5)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Pluck("id", &companyProductIDs)

	otherCompany := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.Company)
	DBCon.Create(otherCompany)
	otherCompanyProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": otherCompany.ID}).(*models.Project)
	DBCon.Create(otherCompanyProduct)

	result := testutil.NewRequest().Get("/projects").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), "0", res.NextPageToken)
}

func (s *TestProjectsHandlerSuite) TestGetProjectsFetchLists_WithPageToken_HavingNextPage_StatusOK() {
	company, cookieString := s.companySignIn()

	project1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project1)
	project2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project2)
	project3 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project3)
	project4 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project4)
	project5 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project5)
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project6)
	project7 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project7)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ?", project2.ID).Where("id < ?", project7.ID).Pluck("id", &companyProductIDs)

	otherCompany := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.Company)
	DBCon.Create(otherCompany)
	otherCompanyProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": otherCompany.ID}).(*models.Project)
	DBCon.Create(otherCompanyProduct)

	result := testutil.NewRequest().Get("/projects?pageToken="+strconv.Itoa(project2.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), strconv.Itoa(project7.ID), res.NextPageToken)
}

func (s *TestProjectsHandlerSuite) TestGetProjectsFetchLists_StatusUnauthorized() {
	company, _ := s.companySignIn()

	var projects []models.Project
	product1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	product2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	projects = append(projects, *product1, *product2)
	DBCon.CreateInBatches(projects, len(projects))

	otherCompany := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.Company)
	DBCon.Create(otherCompany)
	otherCompanyProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": otherCompany.ID}).(*models.Project)
	DBCon.Create(otherCompanyProduct)

	result := testutil.NewRequest().Get("/projects").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func (s *TestProjectsHandlerSuite) TestGetProjectsFetchLists_StatusInternalServerError() {
	company, cookieString := s.companySignIn()

	var projects []models.Project
	product1 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	product2 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	projects = append(projects, *product1, *product2)
	DBCon.CreateInBatches(projects, len(projects))

	otherCompany := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.Company)
	DBCon.Create(otherCompany)
	otherCompanyProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": otherCompany.ID}).(*models.Project)
	DBCon.Create(otherCompanyProduct)

	mockProjectService := new(MockProjectService)
	mockProjectService.On("FetchLists", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]models.Project{}, 0, errors.New("internal server error"))
	mockProjectService.On("Create", mock.AnythingOfType("*businessapi.ProjectStoreInput"), mock.AnythingOfType("int")).Return(models.Project{}, nil, nil)
	mockProjectService.On("Update", mock.AnythingOfType("*businessapi.ProjectStoreInput"), mock.AnythingOfType("int")).Return(models.Project{}, nil, nil)
	mockProjectService.On("MappingValidationErrorStruct", mock.AnythingOfType("error")).Return(businessapi.ProjectValidationError{})
	s.initializeHandlers(mockProjectService, businessservices.NewPlanService(DBCon))
	result := testutil.NewRequest().Get("/projects").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusInternalServerError, result.Code())
}

func (s *TestProjectsHandlerSuite) TestPostProjectCreate_StatusOk() {
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
	reqBody := businessapi.PostProjectJSONRequestBody{Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: isActive}
	result := testutil.NewRequest().Post("/projects").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostProject200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), title, res.Project.Title)
	assert.Equal(s.T(), description, res.Project.Description)
	assert.Equal(s.T(), startDate, res.Project.StartDate)
	assert.Equal(s.T(), endDate, res.Project.EndDate)
	assert.Equal(s.T(), minBudget, *res.Project.MinBudget)
	assert.Equal(s.T(), maxBudget, *res.Project.MaxBudget)
	assert.Equal(s.T(), isActive, res.Project.IsActive)

	expectedValidationErrors := businessapi.ProjectValidationError{}
	assert.Equal(s.T(), expectedValidationErrors, res.Errors)

	// NOTE: DBの値を確認
	var project models.Project
	count := DBCon.Model(&project).Where("company_id = ?", company.ID).Where("title = ?", title).Take(&project).RowsAffected
	assert.Equal(s.T(), int64(1), count)
}

func (s *TestProjectsHandlerSuite) TestPostProjectCreate_StatusForbidden() {
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
	reqBody := businessapi.PostProjectJSONRequestBody{Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: isActive}
	result := testutil.NewRequest().Post("/projects").WithHeader("Cookie", cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusForbidden, result.Code())

	// NOTE: DBの値を確認
	var project models.Project
	count := DBCon.Model(&project).Where("company_id = ?", company.ID).Where("title = ?", title).Take(&project).RowsAffected
	assert.Equal(s.T(), int64(0), count)
}

func (s *TestProjectsHandlerSuite) TestPostProjectCreate_StatusUnauthorized() {
	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	reqBody := businessapi.PostProjectJSONRequestBody{Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: isActive}
	result := testutil.NewRequest().Post("/projects").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())

	// NOTE: DBの値を確認
	var project models.Project
	count := DBCon.Model(&project).Where("title = ?", title).Take(&project).RowsAffected
	assert.Equal(s.T(), int64(0), count)
}

func (s *TestProjectsHandlerSuite) TestPostProjectCreate_StatusInternalServerError() {
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
	reqBody := businessapi.PostProjectJSONRequestBody{Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: isActive}

	mockProjectService := new(MockProjectService)
	mockProjectService.On("FetchLists", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]models.Project{}, 0, nil)
	mockProjectService.On("Create", mock.AnythingOfType("*businessapi.ProjectStoreInput"), mock.AnythingOfType("int")).Return(models.Project{}, nil, errors.New("internal server error"))
	mockProjectService.On("Update", mock.AnythingOfType("*businessapi.ProjectStoreInput"), mock.AnythingOfType("int")).Return(models.Project{}, nil, nil)
	mockProjectService.On("MappingValidationErrorStruct", mock.AnythingOfType("error")).Return(businessapi.ProjectValidationError{})
	s.initializeHandlers(mockProjectService, businessservices.NewPlanService(DBCon))

	result := testutil.NewRequest().Post("/projects").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusInternalServerError, result.Code())

	// NOTE: DBの値を確認
	var project models.Project
	count := DBCon.Model(&project).Where("company_id = ?", company.ID).Where("title = ?", title).Take(&project).RowsAffected
	assert.Equal(s.T(), int64(0), count)
}

func (s *TestProjectsHandlerSuite) TestPostProjectCreate_BadRequest_Required() {
	company, cookieString := s.companySignIn()

	title := ""
	description := ""
	minBudget := 10000
	maxBudget := 20000
	reqBody := businessapi.PostProjectJSONRequestBody{Title: title, Description: description, StartDate: nil, EndDate: nil, MinBudget: &minBudget, MaxBudget: &maxBudget}
	result := testutil.NewRequest().Post("/projects").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostProject200JSONResponse
	result.UnmarshalBodyToObject(&res)

	titleErrorMessages := []string{"案件タイトルは必須入力です。"}
	descriptionErrorMessages := []string{"案件概要は必須入力です。"}
	startDateErrorMessages := []string{"案件開始日は必須入力です。"}
	endDateErrorMessages := []string{"案件終了日は必須入力です。"}
	assert.Equal(s.T(), titleErrorMessages, *res.Errors.Title)
	assert.Equal(s.T(), descriptionErrorMessages, *res.Errors.Description)
	assert.Equal(s.T(), startDateErrorMessages, *res.Errors.StartDate)
	assert.Equal(s.T(), endDateErrorMessages, *res.Errors.EndDate)

	// NOTE: DBの値を確認
	var project models.Project
	count := DBCon.Model(&project).Where("company_id = ?", company.ID).Where("title = ?", title).Take(&project).RowsAffected
	assert.Equal(s.T(),int64(0), count)
}

func (s *TestProjectsHandlerSuite) TestGetProject_StatusOk() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	
	result := testutil.NewRequest().Get("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostProject200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), project.Title, res.Project.Title)
	assert.Equal(s.T(), project.Description, res.Project.Description)
	assert.Equal(s.T(), project.StartDate.Format("2006-01-02"), res.Project.StartDate.Format("2006-01-02"))
	assert.Equal(s.T(), project.EndDate.Format("2006-01-02"), res.Project.EndDate.Format("2006-01-02"))
	assert.Equal(s.T(), project.MinBudget.Int, *res.Project.MinBudget)
	assert.NotNil(s.T(), *res.Project.MinBudget)
	assert.Equal(s.T(), project.MaxBudget.Int, *res.Project.MaxBudget)
	assert.NotNil(s.T(), *res.Project.MaxBudget)
	assert.Equal(s.T(), project.IsActive, res.Project.IsActive)
}

func (s *TestProjectsHandlerSuite) TestGetProject_EmptyBudget_StatusOk() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	DBCon.Create(project)
	
	result := testutil.NewRequest().Get("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostProject200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), project.Title, res.Project.Title)
	assert.Equal(s.T(), project.Description, res.Project.Description)
	assert.Equal(s.T(), project.StartDate.Format("2006-01-02"), res.Project.StartDate.Format("2006-01-02"))
	assert.Equal(s.T(), project.EndDate.Format("2006-01-02"), res.Project.EndDate.Format("2006-01-02"))
	assert.Nil(s.T(), res.Project.MinBudget)
	assert.Nil(s.T(), res.Project.MaxBudget)
	assert.Equal(s.T(), project.IsActive, res.Project.IsActive)
}

func (s *TestProjectsHandlerSuite) TestGetProject_StatusNotFound() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	
	result := testutil.NewRequest().Get("/projects/"+strconv.Itoa(project.ID+1)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusNotFound, result.Code())
}

func (s *TestProjectsHandlerSuite) TestGetProject_StatusUnauthorized() {
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	DBCon.Create(company)
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	
	result := testutil.NewRequest().Get("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func (s *TestProjectsHandlerSuite) TestPutProject_StatusOk() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	
	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.Local)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	reqBody := businessapi.PutProjectJSONRequestBody{Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: isActive}
	result := testutil.NewRequest().Put("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostProject200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), title, res.Project.Title)
	assert.Equal(s.T(), description, res.Project.Description)
	assert.Equal(s.T(), startDate.Format("2006-01-02"), res.Project.StartDate.Format("2006-01-02"))
	assert.Equal(s.T(), endDate.Format("2006-01-02"), res.Project.EndDate.Format("2006-01-02"))
	assert.Equal(s.T(), minBudget, *res.Project.MinBudget)
	assert.Equal(s.T(), maxBudget, *res.Project.MaxBudget)
	assert.Equal(s.T(), isActive, res.Project.IsActive)

	expectedValidationErrors := businessapi.ProjectValidationError{}
	assert.Equal(s.T(), expectedValidationErrors, res.Errors)

	// NOTE: DBの値を確認
	DBCon.Model(&project).Where("id = ?", project.ID).Take(&project)
	assert.Equal(s.T(), company.ID, project.CompanyID)
	assert.Equal(s.T(), "test title", project.Title)
	assert.Equal(s.T(), "test description", project.Description)
	assert.Equal(s.T(), parsedStartDate.Format("2006-01-02"), project.StartDate.Format("2006-01-02"))
	assert.Equal(s.T(), parsedEndDate.Format("2006-01-02"), project.EndDate.Format("2006-01-02"))
	assert.Equal(s.T(), null.Int{Int: minBudget, Valid: true}, project.MinBudget)
	assert.Equal(s.T(), null.Int{Int: maxBudget, Valid: true}, project.MaxBudget)
	assert.Equal(s.T(), isActive, project.IsActive)
}

func (s *TestProjectsHandlerSuite) TestPutProject_StatusForbidden() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	reqBody := businessapi.PutProjectJSONRequestBody{Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: isActive}
	result := testutil.NewRequest().Put("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusForbidden, result.Code())

	// NOTE: DBの値を確認
	DBCon.Model(&project).Where("id = ?", project.ID).Take(&project)
	assert.NotEqual(s.T(), "test title", project.Title)
}

func (s *TestProjectsHandlerSuite) TestPutProject_StatusUnauthorized() {
	company := factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	DBCon.Create(company)
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	
	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	reqBody := businessapi.PutProjectJSONRequestBody{Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: isActive}
	result := testutil.NewRequest().Put("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())

	// NOTE: DBの値を確認
	DBCon.Model(&project).Where("id = ?", project.ID).Take(&project)
	assert.NotEqual(s.T(), "test title", project.Title)
}

func (s *TestProjectsHandlerSuite) TestPutProject_StatusInternalServerError() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)

	title := "test title"
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	minBudget := 10000
	maxBudget := 20000
	isActive := true
	reqBody := businessapi.PutProjectJSONRequestBody{Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, MinBudget: &minBudget, MaxBudget: &maxBudget, IsActive: isActive}

	mockProjectService := new(MockProjectService)
	mockProjectService.On("FetchLists", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]models.Project{}, 0, nil)
	mockProjectService.On("Create", mock.AnythingOfType("*businessapi.ProjectStoreInput"), mock.AnythingOfType("int")).Return(models.Project{}, nil, errors.New("internal server error"))
	mockProjectService.On("Update", mock.AnythingOfType("*businessapi.ProjectStoreInput"), mock.AnythingOfType("int")).Return(models.Project{}, nil, errors.New("internal server error"))
	mockProjectService.On("MappingValidationErrorStruct", mock.AnythingOfType("error")).Return(businessapi.ProjectValidationError{})
	s.initializeHandlers(mockProjectService, businessservices.NewPlanService(DBCon))

	result := testutil.NewRequest().Put("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusInternalServerError, result.Code())

	// NOTE: DBの値を確認
	DBCon.Model(&project).Where("id = ?", project.ID).Take(&project)
	assert.NotEqual(s.T(), "test title", project.Title)
}

func (s *TestProjectsHandlerSuite) TestPutProject_BadRequest_Required() {
	company, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)

	title := ""
	description := ""
	minBudget := 10000
	maxBudget := 20000
	reqBody := businessapi.PutProjectJSONRequestBody{Title: title, Description: description, StartDate: nil, EndDate: nil, MinBudget: &minBudget, MaxBudget: &maxBudget}
	result := testutil.NewRequest().Put("/projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostProject200JSONResponse
	result.UnmarshalBodyToObject(&res)

	titleErrorMessages := []string{"案件タイトルは必須入力です。"}
	descriptionErrorMessages := []string{"案件概要は必須入力です。"}
	startDateErrorMessages := []string{"案件開始日は必須入力です。"}
	endDateErrorMessages := []string{"案件終了日は必須入力です。"}
	assert.Equal(s.T(), titleErrorMessages, *res.Errors.Title)
	assert.Equal(s.T(), descriptionErrorMessages, *res.Errors.Description)
	assert.Equal(s.T(), startDateErrorMessages, *res.Errors.StartDate)
	assert.Equal(s.T(), endDateErrorMessages, *res.Errors.EndDate)

	// NOTE: DBの値を確認
	DBCon.Model(&project).Where("id = ?", project.ID).Take(&project)
	assert.NotEqual(s.T(), "", project.Title)
}

func TestProjectsHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(TestProjectsHandlerSuite))
}

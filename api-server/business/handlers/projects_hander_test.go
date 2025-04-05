package businesshandlers

import (
	businessapi "apps/api/business"
	businessservices "apps/business/services"
	models "apps/models/generated"
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestProjectsHandlerSuite struct {
	WithDBSuite
}

type MockProjectService struct {
    mock.Mock
}

func (m *MockProjectService) Create(ctx context.Context, requestParams *businessapi.PostProjectsJSONRequestBody, companyID int) (project models.Project, validatorErrors error, error error) {
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
	mockProjectService.On("Create", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("*businessapi.PostProjectsJSONRequestBody"), mock.AnythingOfType("int") ).Return(models.Project{}, nil, errors.New("internal server error"))
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

func TestProjectsHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(TestProjectsHandlerSuite))
}

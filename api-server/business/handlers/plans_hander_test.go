package businesshandlers

import (
	businessapi "apps/api/business"
	businessservices "apps/business/services"
	models "apps/models"
	"apps/test/factories"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestPlansHandlerSuite struct {
	WithDBSuite
}

type MockPlanService struct {
    mock.Mock
}

func (m *MockPlanService) Create(requestParams *businessapi.PlanStoreInput, supporterID int) (plan models.Plan, validatorErrors error, error error) {
    args := m.Called(requestParams, supporterID)
	return args.Get(0).(models.Plan), args.Error(1), args.Error(2)
}

func (m *MockPlanService) MappingValidationErrorStruct(err error) businessapi.PlanValidationError {
    args := m.Called(err)
    return args.Get(0).(businessapi.PlanValidationError)
}

func (s *TestPlansHandlerSuite) SetupTest() {
	s.SetDBCon()

	s.initializeHandlers(businessservices.NewProjectService(DBCon), businessservices.NewPlanService(DBCon))

	// NOTE: CSRFトークンのセット
	s.SetCsrfHeaderValues()
}

func (s *TestPlansHandlerSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestPlansHandlerSuite) TestPostPlanCreate_RequiredOnly_StatusOk() {
	company := factories.CompanyFactory.MustCreate().(*models.Company)
	DBCon.Create(company)
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	supporter, cookieString := s.supporterSignIn()

	projectID := project.ID
	title := randomdata.RandStringRunes(70)
	description := "test description"
	unitPrice := 10000
	reqBody := businessapi.PostPlanJSONRequestBody{ProjectId: projectID, Title: title, Description: description, StartDate: nil, EndDate: nil, UnitPrice: unitPrice}
	result := testutil.NewRequest().Post("/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostPlan200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), projectID, res.Plan.ProjectId)
	assert.Equal(s.T(), title, res.Plan.Title)
	assert.Equal(s.T(), description, res.Plan.Description)
	assert.Nil(s.T(), res.Plan.StartDate)
	assert.Nil(s.T(), res.Plan.EndDate)
	assert.Equal(s.T(), unitPrice, res.Plan.UnitPrice)

	expectedValidationErrors := businessapi.PlanValidationError{}
	assert.Equal(s.T(), expectedValidationErrors, res.Errors)

	// NOTE: DBの値を確認
	var createdPlan models.Plan
	count := DBCon.Where("supporter_id = ? AND title = ?", supporter.ID, title).First(&createdPlan).RowsAffected
	assert.Equal(s.T(), int64(1), count)
}

func (s *TestPlansHandlerSuite) TestPostPlanCreate_WithOptional_StatusOk() {
	company := factories.CompanyFactory.MustCreate().(*models.Company)
	DBCon.Create(company)
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	supporter, cookieString := s.supporterSignIn()

	projectID := project.ID
	title := randomdata.RandStringRunes(70)
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 9, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	unitPrice := 10000
	reqBody := businessapi.PostPlanJSONRequestBody{ProjectId: projectID, Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, UnitPrice: unitPrice}
	result := testutil.NewRequest().Post("/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostPlan200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), projectID, res.Plan.ProjectId)
	assert.Equal(s.T(), title, res.Plan.Title)
	assert.Equal(s.T(), description, res.Plan.Description)
	assert.Equal(s.T(), &startDate, res.Plan.StartDate)
	assert.Equal(s.T(), &endDate, res.Plan.EndDate)
	assert.Equal(s.T(), unitPrice, res.Plan.UnitPrice)

	expectedValidationErrors := businessapi.PlanValidationError{}
	assert.Equal(s.T(), expectedValidationErrors, res.Errors)

	// NOTE: DBの値を確認
	var createdPlan models.Plan
	count := DBCon.Where("supporter_id = ? AND title = ?", supporter.ID, title).First(&createdPlan).RowsAffected
	assert.Equal(s.T(), int64(1), count)
}

func (s *TestPlansHandlerSuite) TestPostPlanCreate_StatusForbidden() {
	company := factories.CompanyFactory.MustCreate().(*models.Company)
	DBCon.Create(company)
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	supporter, cookieString := s.supporterSignIn()

	projectID := project.ID
	title := randomdata.RandStringRunes(70)
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 9, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	unitPrice := 10000
	reqBody := businessapi.PostPlanJSONRequestBody{ProjectId: projectID, Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, UnitPrice: unitPrice}
	result := testutil.NewRequest().Post("/plans").WithHeader("Cookie", cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusForbidden, result.Code())

	// NOTE: DBの値を確認
	var createdPlan models.Plan
	count := DBCon.Where("supporter_id = ? AND title = ?", supporter.ID, title).First(&createdPlan).RowsAffected
	assert.Equal(s.T(), int64(0), count)
}

func (s *TestPlansHandlerSuite) TestPostPlanCreate_StatusUnauthorized() {
	company := factories.CompanyFactory.MustCreate().(*models.Company)
	DBCon.Create(company)
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)

	projectID := project.ID
	title := randomdata.RandStringRunes(70)
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 9, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	unitPrice := 10000
	reqBody := businessapi.PostPlanJSONRequestBody{ProjectId: projectID, Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, UnitPrice: unitPrice}
	result := testutil.NewRequest().Post("/plans").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())

	// NOTE: DBの値を確認
	var createdPlan models.Plan
	count := DBCon.Where("title = ?", title).First(&createdPlan).RowsAffected
	assert.Equal(s.T(), int64(0), count)
}

func (s *TestPlansHandlerSuite) TestPostPlanCreate_StatusInternalServerError() {
	company := factories.CompanyFactory.MustCreate().(*models.Company)
	DBCon.Create(company)
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	supporter, cookieString := s.supporterSignIn()

	projectID := project.ID
	title := randomdata.RandStringRunes(70)
	description := "test description"
	parsedStartDate := time.Date(2025, 4, 9, 0, 0, 0, 0, time.UTC)
	startDate := openapi_types.Date{Time: parsedStartDate}
	parsedEndDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	endDate := openapi_types.Date{Time: parsedEndDate}
	unitPrice := 10000
	reqBody := businessapi.PostPlanJSONRequestBody{ProjectId: projectID, Title: title, Description: description, StartDate: &startDate, EndDate: &endDate, UnitPrice: unitPrice}

	mockPlanService := new(MockPlanService)
	mockPlanService.On("Create", mock.AnythingOfType("*businessapi.PlanStoreInput"), mock.AnythingOfType("int")).Return(models.Plan{}, nil, errors.New("internal server error"))
	mockPlanService.On("MappingValidationErrorStruct", mock.AnythingOfType("error")).Return(businessapi.PlanValidationError{})
	s.initializeHandlers(businessservices.NewProjectService(DBCon), mockPlanService)

	result := testutil.NewRequest().Post("/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusInternalServerError, result.Code())

	// NOTE: DBの値を確認
	var createdPlan models.Plan
	count := DBCon.Where("supporter_id = ? AND title = ?", supporter.ID, title).First(&createdPlan).RowsAffected
	assert.Equal(s.T(), int64(0), count)
}

func (s *TestPlansHandlerSuite) TestPostPlanCreate_BadRequest_Required() {
	company := factories.CompanyFactory.MustCreate().(*models.Company)
	DBCon.Create(company)
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	supporter, cookieString := s.supporterSignIn()

	projectID := project.ID
	title := ""
	description := ""
	reqBody := businessapi.PostPlanJSONRequestBody{ProjectId: projectID, Title: title, Description: description}
	result := testutil.NewRequest().Post("/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostPlan200JSONResponse
	result.UnmarshalBodyToObject(&res)

	titleErrorMessages := []string{"支援タイトルは必須入力です。"}
	descriptionErrorMessages := []string{"支援概要は必須入力です。"}
	unitPriceErrorMessages := []string{"支援単価(税抜)は必須入力です。"}
	assert.Equal(s.T(), titleErrorMessages, *res.Errors.Title)
	assert.Equal(s.T(), descriptionErrorMessages, *res.Errors.Description)
	assert.Equal(s.T(), unitPriceErrorMessages, *res.Errors.UnitPrice)

	// NOTE: DBの値を確認
	var createdPlan models.Plan
	count := DBCon.Where("supporter_id = ? AND title = ?", supporter.ID, title).First(&createdPlan).RowsAffected
	assert.Equal(s.T(), int64(0), count)
}

func TestPlansHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(TestPlansHandlerSuite))
}

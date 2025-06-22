package businesshandlers

import (
	businessapi "apps/api/business"
	businessservices "apps/business/services"
	models "apps/models"
	"apps/test/factories"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
)

type TestToProjectsHandlerSuite struct {
	WithDBSuite
}

var (
	company *models.Company
	project1 *models.Project
	project2 *models.Project
	project3 *models.Project
	project4 *models.Project
	project5 *models.Project
)

func (s *TestToProjectsHandlerSuite) SetupTest() {
	s.SetDBCon()

	s.initializeHandlers(businessservices.NewProjectService(DBCon), businessservices.NewPlanService(DBCon))

	// NOTE: CSRFトークンのセット
	s.SetCsrfHeaderValues()

	// NOTE: テスト用企業の作成
	company = factories.CompanyFactory.MustCreate().(*models.Company)
	DBCon.Create(company)

	project1 = factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 24, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 24, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project1)
	project2 = factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 25, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 25, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project2)
	project3 = factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 26, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 26, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project3)
	project4 = factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 27, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 27, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project4)
	project5 = factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 28, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 28, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project5)
}

func (s *TestToProjectsHandlerSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_StatusOk() {
	_, cookieString := s.supporterSignIn()

	budgetEmptyproduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 6, 1, 0, 0, 0, 0, time.Local),  "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	DBCon.Create(budgetEmptyproduct)
	havingBudgetProduct := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 6, 2, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(havingBudgetProduct)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ? AND id <= ?", budgetEmptyproduct.ID, havingBudgetProduct.ID).Pluck("id", &companyProductIDs)

	result := testutil.NewRequest().Get("/to-projects?startDate=2025-06-01").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)

	// NOTE: 予算カラムがnullの時はnull、そうでなければ値が返ること
	var resBudgetEmptyproducts []businessapi.ToProject
	for _, project := range res.Projects {
		if project.MinBudget == nil && project.MaxBudget == nil {
			resBudgetEmptyproducts = append(resBudgetEmptyproducts, project)
		}
	}
	assert.Len(s.T(), resBudgetEmptyproducts, 1)

	var resHaveBudgetproducts []businessapi.ToProject
	for _, project := range res.Projects {
		if project.MinBudget != nil && project.MaxBudget != nil {
			resHaveBudgetproducts = append(resHaveBudgetproducts, project)
		}
	}
	assert.Len(s.T(), resHaveBudgetproducts, 1)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_StatusOk_OnlyRequired_WithoutPlanSteps() {
	_, cookieString := s.supporterSignIn()

	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "テスト提案",
		Description: "テスト提案の概要です",
		UnitPrice:   5000,
		PlanSteps:   nil,
	}

	result := testutil.NewRequest().Post("/to-projects/"+strconv.Itoa(project1.ID)+"/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostToProjectPlan200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), "テスト提案", res.Plan.Title)
	assert.Equal(s.T(), "テスト提案の概要です", res.Plan.Description)
	assert.Equal(s.T(), 5000, res.Plan.UnitPrice)
	assert.Equal(s.T(), project1.ID, res.Plan.ProjectId)

	// NOTE: エラーがない場合は空の構造体が返される
	assert.Nil(s.T(), res.Errors.Title)
	assert.Nil(s.T(), res.Errors.Description)
	assert.Nil(s.T(), res.Errors.UnitPrice)

	// NOTE: データベースに保存されているかの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ? AND title = ?", project1.ID, "テスト提案").First(&savedPlan)
	assert.Equal(s.T(), "テスト提案", savedPlan.Title)
	assert.Equal(s.T(), "テスト提案の概要です", savedPlan.Description)
	assert.Equal(s.T(), 5000, savedPlan.UnitPrice)
	assert.Equal(s.T(), project1.ID, project1.ID)
	assert.Equal(s.T(), null.Time{Time:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Valid:false}, savedPlan.StartDate)
	assert.Equal(s.T(), null.Time{Time:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Valid:false}, savedPlan.EndDate)
	assert.Equal(s.T(), models.PlanStatusTempraryCreating, savedPlan.Status)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_StatusOk_PlanOnlyRequired_WithPlanSteps() {
	_, cookieString := s.supporterSignIn()

	planSteps := []businessapi.PlanStepInput{
		{
			Title:       "ステップ1",
			Description: "ステップ1の詳細",
			Duration:    5,
		},
		{
			Title:       "ステップ2",
			Description: "ステップ2の詳細",
			Duration:    10,
		},
	}

	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "ステップ付き提案",
		Description: "ステップ付き提案の概要です", 
		UnitPrice:   7500,
		PlanSteps:   &planSteps,
	}

	result := testutil.NewRequest().Post("/to-projects/"+strconv.Itoa(project1.ID)+"/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostToProjectPlan200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), "ステップ付き提案", res.Plan.Title)
	assert.Equal(s.T(), "ステップ付き提案の概要です", res.Plan.Description)
	assert.Equal(s.T(), 7500, res.Plan.UnitPrice)
	assert.Equal(s.T(), project1.ID, res.Plan.ProjectId)
	assert.NotNil(s.T(), res.Plan.PlanSteps)
	assert.Len(s.T(), *res.Plan.PlanSteps, 2)
	assert.Equal(s.T(), "ステップ1", (*res.Plan.PlanSteps)[0].Title)
	assert.Equal(s.T(), "ステップ1の詳細", (*res.Plan.PlanSteps)[0].Description)
	assert.Equal(s.T(), 5, (*res.Plan.PlanSteps)[0].Duration)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), "ステップ付き提案", res.Plan.Title)
	assert.Equal(s.T(), "ステップ付き提案の概要です", res.Plan.Description)
	assert.Equal(s.T(), 7500, res.Plan.UnitPrice)

	// NOTE: PlanStepsが保存されているかの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ? AND title = ?", project1.ID, "ステップ付き提案").First(&savedPlan)
	
	var savedPlanSteps []models.PlanStep
	DBCon.Where("plan_id = ?", savedPlan.ID).Find(&savedPlanSteps)
	assert.Len(s.T(), savedPlanSteps, 2)
	assert.Equal(s.T(), "ステップ1", savedPlanSteps[0].Title)
	assert.Equal(s.T(), "ステップ1の詳細", savedPlanSteps[0].Description)
	assert.Equal(s.T(), 5, savedPlanSteps[0].Duration)
	assert.Equal(s.T(), "ステップ2", savedPlanSteps[1].Title)
	assert.Equal(s.T(), "ステップ2の詳細", savedPlanSteps[1].Description)
	assert.Equal(s.T(), 10, savedPlanSteps[1].Duration)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_StatusOk_WithDates() {
	_, cookieString := s.supporterSignIn()

	stardDate := openapi_types.Date{Time: time.Date(2025, 6, 1, 0, 0, 0, 0, time.Local)}
	endDate := openapi_types.Date{Time: time.Date(2025, 6, 30, 0, 0, 0, 0, time.Local)}
	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "日付付き提案",
		Description: "日付付き提案の概要です",
		StartDate:   &stardDate,
		EndDate:     &endDate,
		UnitPrice:   5000,
		PlanSteps:   nil,
	}

	result := testutil.NewRequest().Post("/to-projects/"+strconv.Itoa(project1.ID)+"/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.PostToProjectPlan200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), "日付付き提案", res.Plan.Title)
	assert.Equal(s.T(), openapi_types.Date{Time: time.Date(stardDate.Year(), stardDate.Month(), stardDate.Day(), stardDate.Hour(), stardDate.Minute(), stardDate.Second(), stardDate.Nanosecond(), time.UTC)} , *res.Plan.StartDate)
	assert.Equal(s.T(), openapi_types.Date{Time: time.Date(endDate.Year(), endDate.Month(), endDate.Day(), endDate.Hour(), endDate.Minute(), endDate.Second(), endDate.Nanosecond(), time.UTC)}, *res.Plan.EndDate)

	// NOTE: Planが保存されているかの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ?", project1.ID).Preload("PlanSteps").First(&savedPlan)
	assert.Equal(s.T(), "日付付き提案", savedPlan.Title)
	assert.Equal(s.T(), "日付付き提案の概要です", savedPlan.Description)
	assert.Equal(s.T(), 5000, savedPlan.UnitPrice)
	assert.Equal(s.T(), null.Time{Time: stardDate.Time, Valid: true}, savedPlan.StartDate)
	assert.Equal(s.T(), null.Time{Time: endDate.Time, Valid: true}, savedPlan.EndDate)
	assert.Equal(s.T(), models.PlanStatusTempraryCreating, savedPlan.Status)
	assert.Empty(s.T(), savedPlan.PlanSteps)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_BadRequest_EmptyTitle() {
	_, cookieString := s.supporterSignIn()

	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "",
		Description: "Test plan description",
		UnitPrice:   5000,
		PlanSteps: nil,
	}

	result := testutil.NewRequest().Post("/to-projects/"+strconv.Itoa(project1.ID)+"/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusBadRequest, result.Code())

	var res businessapi.PostToProjectPlan400JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Contains(s.T(), *res.Title, "支援タイトルは必須入力です。")

	// NOTE: Planが保存されていないことの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ?", project1.ID).Preload("PlanSteps").First(&savedPlan)
	assert.Empty(s.T(), savedPlan.ID)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_BadRequest_EmptyDescription() {
	_, cookieString := s.supporterSignIn()

	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "Test Title",
		Description: "",
		UnitPrice:   5000,
		PlanSteps: nil,
	}

	result := testutil.NewRequest().Post("/to-projects/"+strconv.Itoa(project1.ID)+"/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusBadRequest, result.Code())

	var res businessapi.PostToProjectPlan400JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Contains(s.T(), *res.Description, "支援概要は必須入力です。")

	// NOTE: Planが保存されていないことの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ?", project1.ID).Preload("PlanSteps").First(&savedPlan)
	assert.Empty(s.T(), savedPlan.ID)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_BadRequest_InvalidUnitPrice_Minus() {
	_, cookieString := s.supporterSignIn()

	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "Test Title",
		Description: "Test Description", 
		UnitPrice:   -100,
		PlanSteps: nil,
	}

	result := testutil.NewRequest().Post("/to-projects/"+strconv.Itoa(project1.ID)+"/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusBadRequest, result.Code())

	var res businessapi.PostToProjectPlan400JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Contains(s.T(), *res.UnitPrice, "支援単価は1円以上で入力してください。")

	// NOTE: Planが保存されていないことの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ?", project1.ID).Preload("PlanSteps").First(&savedPlan)
	assert.Empty(s.T(), savedPlan.ID)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_BadRequest_InvalidUnitPrice_Zero() {
	_, cookieString := s.supporterSignIn()
	
	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "テスト提案",
		Description: "概要です",
		UnitPrice:   0,
		PlanSteps: nil,
	}
	
	result := testutil.NewRequest().Post("/to-projects/"+strconv.Itoa(project1.ID)+"/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusBadRequest, result.Code())

	var res businessapi.PostToProjectPlan400JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.NotNil(s.T(), res.UnitPrice)
	assert.Contains(s.T(), *res.UnitPrice, "支援単価(税抜)は必須入力です。")

	// NOTE: Planが保存されていないことの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ?", project1.ID).Preload("PlanSteps").First(&savedPlan)
	assert.Empty(s.T(), savedPlan.ID)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_BadRequest_InvalidPlanSteps() {
	_, cookieString := s.supporterSignIn()

	planSteps := []businessapi.PlanStepInput{
		{
			Title:       "",
			Description: "Description",
			Duration:    -5,
		},
	}

	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "Test Title",
		Description: "Test Description",
		UnitPrice:   5000,
		PlanSteps:   &planSteps,
	}

	result := testutil.NewRequest().Post("/to-projects/"+strconv.Itoa(project1.ID)+"/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusBadRequest, result.Code())

	var res businessapi.PostToProjectPlan400JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.NotNil(s.T(), res.PlanSteps)
	assert.Len(s.T(), *res.PlanSteps, 1)
	assert.Contains(s.T(), *(*res.PlanSteps)[0].Title, "タイトルは必須入力です。")
	assert.Contains(s.T(), *(*res.PlanSteps)[0].Duration, "支援時間は1時間以上で入力してください。")

	// NOTE: Planが保存されていないことの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ?", project1.ID).Preload("PlanSteps").First(&savedPlan)
	assert.Empty(s.T(), savedPlan.ID)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_BadRequest_DateRange() {
	_, cookieString := s.supporterSignIn()

	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "不正日付提案",
		Description: "不正日付提案の概要です",
		StartDate:   &openapi_types.Date{Time: time.Date(2025, 6, 2, 0, 0, 0, 0, time.Local)},
		EndDate:     &openapi_types.Date{Time: time.Date(2025, 6, 1, 0, 0, 0, 0, time.Local)}, // 開始日より前
		UnitPrice:   5000,
		PlanSteps:   nil,
	}

	result := testutil.NewRequest().Post("/to-projects/"+strconv.Itoa(project1.ID)+"/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusBadRequest, result.Code())

	var res businessapi.PostToProjectPlan400JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.NotNil(s.T(), res.EndDate)
	assert.Contains(s.T(), (*res.EndDate)[0], "前後関係が不適です")

	// NOTE: Planが保存されていないことの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ?", project1.ID).Preload("PlanSteps").First(&savedPlan)
	assert.Empty(s.T(), savedPlan.ID)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_BadRequest_MissingEndDate() {
	_, cookieString := s.supporterSignIn()

	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "終了日なし提案",
		Description: "終了日なし提案の概要です",
		StartDate:   &openapi_types.Date{Time: time.Date(2025, 6, 1, 0, 0, 0, 0, time.Local)},
		EndDate:     nil,
		UnitPrice:   5000,
		PlanSteps:   nil,
	}

	result := testutil.NewRequest().Post("/to-projects/"+strconv.Itoa(project1.ID)+"/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusBadRequest, result.Code())

	var res businessapi.PostToProjectPlan400JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.NotNil(s.T(), res.EndDate)
	assert.Contains(s.T(), (*res.EndDate)[0], "支援終了日は必須入力です")

	// NOTE: Planが保存されていないことの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ?", project1.ID).Preload("PlanSteps").First(&savedPlan)
	assert.Empty(s.T(), savedPlan.ID)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_BadRequest_EmptyPlanSteps() {
	_, cookieString := s.supporterSignIn()

	planSteps := []businessapi.PlanStepInput{}

	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "空ステップ提案",
		Description: "空ステップ提案の概要です",
		UnitPrice:   5000,
		PlanSteps:   &planSteps,
	}

	result := testutil.NewRequest().Post("/to-projects/"+strconv.Itoa(project1.ID)+"/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusBadRequest, result.Code())

	var res businessapi.PostToProjectPlan400JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.NotNil(s.T(), res.PlanSteps)
	assert.Contains(s.T(), (*(*res.PlanSteps)[0].Title)[0], "少なくとも1つの支援ステップを追加してください")

	// NOTE: Planが保存されていないことの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ?", project1.ID).Preload("PlanSteps").First(&savedPlan)
	assert.Empty(s.T(), savedPlan.ID)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_NotFound() {
	_, cookieString := s.supporterSignIn()

	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "テスト提案",
		Description: "概要です",
		UnitPrice:   5000,
		PlanSteps:   nil,
	}

	result := testutil.NewRequest().Post("/to-projects/0/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusNotFound, result.Code())

	// NOTE: Planが保存されていないことの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ?", 0).Preload("PlanSteps").First(&savedPlan)
	assert.Empty(s.T(), savedPlan.ID)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_Unauthorized() {
	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "テスト提案",
		Description: "概要です",
		UnitPrice:   5000,
		PlanSteps:   nil,
	}

	result := testutil.NewRequest().Post("/to-projects/"+strconv.Itoa(project1.ID)+"/plans").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())

	// NOTE: Planが保存されていないことの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ?", project1.ID).Preload("PlanSteps").First(&savedPlan)
	assert.Empty(s.T(), savedPlan.ID)
}

func (s *TestToProjectsHandlerSuite) TestPostToProjectPlan_Forbidden() {
	_, cookieString := s.companySignIn()

	reqBody := businessapi.PlanStoreWithStepsInput{
		Title:       "テスト提案",
		Description: "概要です",
		UnitPrice:   5000,
		PlanSteps:   nil,
	}

	result := testutil.NewRequest().Post("/to-projects/"+strconv.Itoa(project1.ID)+"/plans").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusForbidden, result.Code())

	// NOTE: Planが保存されていないことの確認
	var savedPlan models.Plan
	DBCon.Where("project_id = ?", project1.ID).Preload("PlanSteps").First(&savedPlan)
	assert.Empty(s.T(), savedPlan.ID)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_EmptyArgs_NotHavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), "0", res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithOnlyPageToken_NotHavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project6)

	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id != ?", project1.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?pageToken="+strconv.Itoa(project2.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), "0", res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithOnlyPageToken_HavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project6)
	project7 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project7)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ?", project2.ID).Where("id < ?", project7.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?pageToken="+strconv.Itoa(project2.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), strconv.Itoa(project7.ID), res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithOnlyStartDate_NotHavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ?", project2.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?startDate=2025-05-25").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), "0", res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithOnlyStartDate_HavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?startDate=2025-05-24").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), strconv.Itoa(project6.ID), res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithOnlyEndDate_NotHavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?endDate=2025-05-28").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), "0", res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithOnlyEndDate_HavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?endDate=2025-05-29").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), strconv.Itoa(project6.ID), res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithPageTokenAndStartDate_NotHavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ? AND id <= ?", project2.ID, project6.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?pageToken="+strconv.Itoa(project2.ID)+"&startDate=2025-05-25").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), "0", res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithPageTokenAndStartDate_HavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?pageToken="+strconv.Itoa(project1.ID)+"&startDate=2025-05-24").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), strconv.Itoa(project6.ID), res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithPageTokenAndEndDate_NotHavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ? AND id <= ?", project2.ID, project6.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?pageToken="+strconv.Itoa(project2.ID)+"&endDate=2025-05-29").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), "0", res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithPageTokenAndEndDate_HavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?pageToken="+strconv.Itoa(project1.ID)+"&endDate=2025-05-29").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), strconv.Itoa(project6.ID), res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithStartDateAndEndDate_NotHavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ? AND id <= ?", project2.ID, project6.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?startDate=2025-05-25&endDate=2025-05-29").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), "0", res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithStartDateAndEndDate_HavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?startDate=2025-05-24&endDate=2025-05-29").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), strconv.Itoa(project6.ID), res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithPageTokenAndStartDateAndEndDate_NotHavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ? AND id <= ?", project2.ID, project6.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?pageToken="+strconv.Itoa(project2.ID)+"&startDate=2025-05-25&endDate=2025-05-29").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), "0", res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_WithPageTokenAndStartDateAndEndDate_HavingNextPage_StatusOK() {
	supporter, cookieString := s.supporterSignIn()

	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	result := testutil.NewRequest().Get("/to-projects?pageToken="+strconv.Itoa(project1.ID)+"&startDate=2025-05-24&endDate=2025-05-29").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProjects200JSONResponse
	result.UnmarshalBodyToObject(&res)
	var projectIDs []int
	for _, project := range res.Projects {
		projectIDs = append(projectIDs, project.Id)
	}
	assert.Equal(s.T(), companyProductIDs, projectIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Projects[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Projects[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Projects[4].ProposalStatus)
	assert.Equal(s.T(), strconv.Itoa(project6.ID), res.NextPageToken)
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_StatusUnauthorized() {
	result := testutil.NewRequest().Get("/to-projects").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func (s *TestToProjectsHandlerSuite) TestGetToProjectsFetchLists_NotSupportersAccess_StatusForbidden() {
	_, cookieString := s.companySignIn()

	result := testutil.NewRequest().Get("/to-projects").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusForbidden, result.Code())
}

func (s *TestToProjectsHandlerSuite) TestGetToProject_StatusOk() {
	_, cookieString := s.supporterSignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	
	result := testutil.NewRequest().Get("/to-projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProject200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), project.Title, res.Project.Title)
	assert.Equal(s.T(), project.Description, res.Project.Description)
	assert.Equal(s.T(), project.StartDate.Format("2006-01-02"), res.Project.StartDate.Format("2006-01-02"))
	assert.Equal(s.T(), project.EndDate.Format("2006-01-02"), res.Project.EndDate.Format("2006-01-02"))
	assert.Equal(s.T(), project.MinBudget.Int, *res.Project.MinBudget)
	assert.NotNil(s.T(), *res.Project.MinBudget)
	assert.Equal(s.T(), project.MaxBudget.Int, *res.Project.MaxBudget)
	assert.NotNil(s.T(), *res.Project.MaxBudget)
}

func (s *TestToProjectsHandlerSuite) TestGetToProject_EmptyBudget_StatusOk() {
	_, cookieString := s.supporterSignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	DBCon.Create(project)
	
	result := testutil.NewRequest().Get("/to-projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProject200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), project.Title, res.Project.Title)
	assert.Equal(s.T(), project.Description, res.Project.Description)
	assert.Equal(s.T(), project.StartDate.Format("2006-01-02"), res.Project.StartDate.Format("2006-01-02"))
	assert.Equal(s.T(), project.EndDate.Format("2006-01-02"), res.Project.EndDate.Format("2006-01-02"))
	assert.Nil(s.T(), res.Project.MinBudget)
	assert.Nil(s.T(), res.Project.MaxBudget)
}

func (s *TestToProjectsHandlerSuite) TestGetToProject_NotProposedPlan_StatusOk() {
	_, cookieString := s.supporterSignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	DBCon.Create(project)
	
	result := testutil.NewRequest().Get("/to-projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProject200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), project.Title, res.Project.Title)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, res.Project.ProposalStatus)
}

func (s *TestToProjectsHandlerSuite) TestGetToProject_TemporaryCreatingPlan_StatusOk() {
	supporter, cookieString := s.supporterSignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	DBCon.Create(project)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	
	result := testutil.NewRequest().Get("/to-projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProject200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), project.Title, res.Project.Title)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, res.Project.ProposalStatus)
}

func (s *TestToProjectsHandlerSuite) TestGetToProject_SubmittedPlan_StatusOk() {
	supporter, cookieString := s.supporterSignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	DBCon.Create(project)

	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	
	result := testutil.NewRequest().Get("/to-projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProject200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), project.Title, res.Project.Title)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Project.ProposalStatus)
}

func (s *TestToProjectsHandlerSuite) TestGetToProject_AboveSubmittedPlan_StatusOk() {
	supporter, cookieString := s.supporterSignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "MinBudget": null.Int{Int: 0, Valid: false}, "MaxBudget": null.Int{Int: 0, Valid: false}}).(*models.Project)
	DBCon.Create(project)

	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)
	
	result := testutil.NewRequest().Get("/to-projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res businessapi.GetToProject200JSONResponse
	result.UnmarshalBodyToObject(&res)
	assert.Equal(s.T(), project.Title, res.Project.Title)
	assert.Equal(s.T(), businessapi.PROPOSED, res.Project.ProposalStatus)
}

func (s *TestToProjectsHandlerSuite) TestGetToProject_StatusNotFound() {
	_, cookieString := s.supporterSignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	
	result := testutil.NewRequest().Get("/to-projects/"+strconv.Itoa(project.ID+1)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusNotFound, result.Code())
}

func (s *TestToProjectsHandlerSuite) TestGetToProject_StatusUnauthorized() {
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	
	result := testutil.NewRequest().Get("/to-projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func (s *TestToProjectsHandlerSuite) TestGetToProject_NotSupportersAccess_StatusForbidden() {
	signedInCompany, cookieString := s.companySignIn()
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": signedInCompany.ID}).(*models.Project)
	DBCon.Create(project)
	
	result := testutil.NewRequest().Get("/to-projects/"+strconv.Itoa(project.ID)).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusForbidden, result.Code())
}

func TestToProjectsHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(TestToProjectsHandlerSuite))
}

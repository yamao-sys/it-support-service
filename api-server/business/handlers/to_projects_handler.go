package businesshandlers

import (
	businessapi "apps/api/business"
	businesshelpers "apps/business/helpers"
	businessservices "apps/business/services"
	"context"
	"strconv"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type ToProjectsHandler interface {
	GetToProjects(ctx context.Context, request businessapi.GetToProjectsRequestObject) (businessapi.GetToProjectsResponseObject, error)
	GetToProject(ctx context.Context, request businessapi.GetToProjectRequestObject) (businessapi.GetToProjectResponseObject, error)
	PostToProjectPlan(ctx context.Context, request businessapi.PostToProjectPlanRequestObject) (businessapi.PostToProjectPlanResponseObject, error)
}

type toProjectsHandler struct {
	toProjectService businessservices.ToProjectService
}

func NewToProjectsHandler(toProjectService businessservices.ToProjectService) ToProjectsHandler {
	return &toProjectsHandler{toProjectService}
}

func (tph *toProjectsHandler) GetToProjects(ctx context.Context, request businessapi.GetToProjectsRequestObject) (businessapi.GetToProjectsResponseObject, error) {
	supporterID, _ := businesshelpers.ExtractSupporterID(ctx)
	if supporterID == 0 {
		return businessapi.GetToProjects403Response{}, nil
	}

	var pageToken int
	var startDate string
	var endDate string
	if request.Params.PageToken != nil {
		pageToken, _ = strconv.Atoi(*request.Params.PageToken)
	}
	if request.Params.StartDate != nil {
		startDate = (*request.Params.StartDate).Format("2006-01-02")
	}
	if request.Params.EndDate != nil {
		endDate = (*request.Params.EndDate).Format("2006-01-02")
	}

	projects, nextPageToken := tph.toProjectService.FetchLists(pageToken, startDate, endDate, supporterID)
	var resProjects []businessapi.ToProject
	for _, project := range projects { 
		resProject := businessapi.ToProject{}

		resProject.Id = project.ID
		resProject.Title = project.Title
		resProject.Description = project.Description
		resProject.StartDate = openapi_types.Date{Time: project.StartDate}
		resProject.EndDate = openapi_types.Date{Time: project.EndDate}
		minBudget := &project.MinBudget
		if *minBudget != 0 {
			resProject.MinBudget = minBudget
		} else {
			resProject.MinBudget = nil
		}
		maxBudget := &project.MaxBudget
		if *maxBudget != 0 {
			resProject.MaxBudget = maxBudget
		} else {
			resProject.MaxBudget = nil
		}
		resProject.ProposalStatus = project.ProposalStatus

		resProjects = append(resProjects, resProject)
	}
	return businessapi.GetToProjects200JSONResponse(businessapi.ToProjectsListResponse{
		Projects: resProjects,
		NextPageToken: strconv.Itoa(nextPageToken),
	}), nil
}

func (tph *toProjectsHandler) GetToProject(ctx context.Context, request businessapi.GetToProjectRequestObject) (businessapi.GetToProjectResponseObject, error) {
	supporterID, _ := businesshelpers.ExtractSupporterID(ctx)
	if supporterID == 0 {
		return businessapi.GetToProject403Response{}, nil
	}

	projectID := request.Id
	project, err := tph.toProjectService.Fetch(projectID, supporterID)
	if err != nil {
		return businessapi.GetToProject404Response{}, nil
	}

	resProject := businessapi.ToProject{}
	resProject.Id = project.ID
	resProject.Title = project.Title
	resProject.Description = project.Description
	resProject.StartDate = openapi_types.Date{Time: project.StartDate}
	resProject.EndDate = openapi_types.Date{Time: project.EndDate}
	minBudget := &project.MinBudget
	if *minBudget != 0 {
		resProject.MinBudget = minBudget
	} else {
		resProject.MinBudget = nil
	}
	maxBudget := &project.MaxBudget
	if *maxBudget != 0 {
		resProject.MaxBudget = maxBudget
	} else {
		resProject.MaxBudget = nil
	}
	resProject.ProposalStatus = project.ProposalStatus

	return businessapi.GetToProject200JSONResponse(businessapi.ToProjectResponse{Project: resProject}), nil
}

func (tph *toProjectsHandler) PostToProjectPlan(ctx context.Context, request businessapi.PostToProjectPlanRequestObject) (businessapi.PostToProjectPlanResponseObject, error) {
	supporterID, _ := businesshelpers.ExtractSupporterID(ctx)
	if supporterID == 0 {
		return businessapi.PostToProjectPlan403Response{}, nil
	}

	projectID := request.Id
	
	createdPlan, validationErrors, err := tph.toProjectService.CreatePlan(projectID, request.Body, supporterID)
	if err != nil {
		// NOTE: プロジェクトが見つからない場合
		if err.Error() == "project not found" {
			return businessapi.PostToProjectPlan404Response{}, nil
		}
		// NOTE: その他のサーバーエラー
		return businessapi.PostToProjectPlan500Response{}, err
	}

	// NOTE: バリデーションエラーがある場合
	if validationErrors != nil {
		mappedValidationErrors := tph.toProjectService.MappingPlanWithStepsValidationErrorStruct(validationErrors)
		return businessapi.PostToProjectPlan400JSONResponse(mappedValidationErrors), nil
	}

	// NOTE: 成功レスポンスの構築
	resPlan := businessapi.Plan{
		Id:          createdPlan.ID,
		ProjectId:   createdPlan.ProjectID,
		Title:       createdPlan.Title,
		Description: createdPlan.Description,
		UnitPrice:   createdPlan.UnitPrice,
		CreatedAt:   createdPlan.CreatedAt,
	}
	if createdPlan.StartDate.Valid {
		resPlan.StartDate = &openapi_types.Date{Time: createdPlan.StartDate.Time}
	}
	if createdPlan.EndDate.Valid {
		resPlan.EndDate = &openapi_types.Date{Time: createdPlan.EndDate.Time}
	}

	// NOTE: PlanStepsをレスポンスに含める
	if len(createdPlan.PlanSteps) > 0 {
		resPlanSteps := make([]businessapi.PlanStep, 0, len(createdPlan.PlanSteps))
		for _, planStep := range createdPlan.PlanSteps {
			resPlanStep := businessapi.PlanStep{
				Id:          planStep.ID,
				PlanId:      planStep.PlanID,
				Title:       planStep.Title,
				Description: planStep.Description,
				Duration:    planStep.Duration,
			}
			resPlanSteps = append(resPlanSteps, resPlanStep)
		}
		resPlan.PlanSteps = &resPlanSteps
	}

	// NOTE: 空のバリデーションエラーを作成
	emptyValidationErrors := tph.toProjectService.MappingPlanWithStepsValidationErrorStruct(nil)

	return businessapi.PostToProjectPlan200JSONResponse(businessapi.PlanWithStepsStoreResponse{
		Plan:   resPlan,
		Errors: emptyValidationErrors,
	}), nil
}

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

	projects, nextPageToken := tph.toProjectService.FetchLists(pageToken, startDate, endDate)
	var resProjects []businessapi.ToProject
	for _, project := range projects { 
		resProject := businessapi.ToProject{}

		resProject.Id = project.ID
		resProject.Title = project.Title
		resProject.Description = project.Description
		resProject.StartDate = openapi_types.Date{Time: project.StartDate}
		resProject.EndDate = openapi_types.Date{Time: project.EndDate}
		minBudget := &project.MinBudget.Int
		if *minBudget != 0 {
			resProject.MinBudget = minBudget
		} else {
			resProject.MinBudget = nil
		}
		maxBudget := &project.MaxBudget.Int
		if *maxBudget != 0 {
			resProject.MaxBudget = maxBudget
		} else {
			resProject.MaxBudget = nil
		}

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
	project, err := tph.toProjectService.Fetch(projectID)
	if err != nil {
		return businessapi.GetToProject404Response{}, nil
	}

	resProject := businessapi.ToProject{}
	resProject.Id = project.ID
	resProject.Title = project.Title
	resProject.Description = project.Description
	resProject.StartDate = openapi_types.Date{Time: project.StartDate}
	resProject.EndDate = openapi_types.Date{Time: project.EndDate}
	minBudget := &project.MinBudget.Int
	if *minBudget != 0 {
		resProject.MinBudget = minBudget
	} else {
		resProject.MinBudget = nil
	}
	maxBudget := &project.MaxBudget.Int
	if *maxBudget != 0 {
		resProject.MaxBudget = maxBudget
	} else {
		resProject.MaxBudget = nil
	}

	return businessapi.GetToProject200JSONResponse(businessapi.ToProjectResponse{Project: resProject}), nil
}

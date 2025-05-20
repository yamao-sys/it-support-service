package businesshandlers

import (
	businessapi "apps/api/business"
	businesshelpers "apps/business/helpers"
	businessservices "apps/business/services"
	"context"
	"net/http"
	"strconv"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type ProjectsHandler interface {
	GetProjects(ctx context.Context, request businessapi.GetProjectsRequestObject) (businessapi.GetProjectsResponseObject, error)
	PostProject(ctx context.Context, request businessapi.PostProjectRequestObject) (businessapi.PostProjectResponseObject, error)
	GetProject(ctx context.Context, request businessapi.GetProjectRequestObject) (businessapi.GetProjectResponseObject, error)
	PutProject(ctx context.Context, request businessapi.PutProjectRequestObject) (businessapi.PutProjectResponseObject, error)
}

type projectsHandler struct {
	projectService businessservices.ProjectService
}

func NewProjectsHandler(projectService businessservices.ProjectService) ProjectsHandler {
	return &projectsHandler{projectService}
}

func (ph *projectsHandler) GetProjects(ctx context.Context, request businessapi.GetProjectsRequestObject) (businessapi.GetProjectsResponseObject, error) {
	companyID, _ := businesshelpers.ExtractCompanyID(ctx)
	var pageToken int
	if request.Params.PageToken != nil {
		pageToken, _ = strconv.Atoi(*request.Params.PageToken)
	}

	projects, nextPageToken, err := ph.projectService.FetchLists(companyID, pageToken)
	if err != nil {
		return businessapi.GetProjects500JSONResponse{Code: http.StatusInternalServerError}, err
	}
	var resProjects []businessapi.Project
	for _, project := range projects { 
		resProject := businessapi.Project{}

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
		resProject.IsActive = project.IsActive
		resProject.CreatedAt = project.CreatedAt

		resProjects = append(resProjects, resProject)
	}
	return businessapi.GetProjects200JSONResponse(businessapi.ProjectsListResponse{
		Projects: resProjects,
		NextPageToken: strconv.Itoa(nextPageToken),
	}), nil
}

func (ph *projectsHandler) PostProject(ctx context.Context, request businessapi.PostProjectRequestObject) (businessapi.PostProjectResponseObject, error) {
	companyID, _ := businesshelpers.ExtractCompanyID(ctx)

	inputs := businessapi.ProjectStoreInput{
		Title: request.Body.Title,
		Description: request.Body.Description,
		StartDate: request.Body.StartDate,
		EndDate: request.Body.EndDate,
		MinBudget: request.Body.MinBudget,
		MaxBudget: request.Body.MaxBudget,
		IsActive: request.Body.IsActive,
	}

	createdProject, validationErrors, err := ph.projectService.Create(&inputs, companyID)
	if err != nil {
		return businessapi.PostProject500JSONResponse{Code: http.StatusInternalServerError}, err
	}

	mappedValidationErrors := ph.projectService.MappingValidationErrorStruct(validationErrors)
	startDate := openapi_types.Date{Time: createdProject.StartDate}
	endDate := openapi_types.Date{Time: createdProject.EndDate}
	resProject := businessapi.Project{
		Id: createdProject.ID,
		Title: createdProject.Title,
		Description: createdProject.Description,
		StartDate: startDate,
		EndDate: endDate,
		MinBudget: &createdProject.MinBudget.Int,
		MaxBudget: &createdProject.MaxBudget.Int,
		IsActive: createdProject.IsActive,
		CreatedAt: createdProject.CreatedAt,
	}
	return businessapi.PostProject200JSONResponse(businessapi.ProjectStoreResponse{Errors: mappedValidationErrors, Project: resProject}), nil
}

func (ph *projectsHandler) GetProject(ctx context.Context, request businessapi.GetProjectRequestObject) (businessapi.GetProjectResponseObject, error) {
	projectID := request.Id
	project, err := ph.projectService.Fetch(projectID)
	if err != nil {
		return businessapi.GetProject404JSONResponse{Code: http.StatusNotFound}, nil
	}

	resProject := businessapi.Project{}
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
	resProject.IsActive = project.IsActive
	resProject.CreatedAt = project.CreatedAt
	return businessapi.GetProject200JSONResponse(businessapi.ProjectResponse{Project: resProject}), nil
}

func (ph *projectsHandler) PutProject(ctx context.Context, request businessapi.PutProjectRequestObject) (businessapi.PutProjectResponseObject, error) {
	projectID := request.Id
	inputs := businessapi.ProjectStoreInput{
		Title: request.Body.Title,
		Description: request.Body.Description,
		StartDate: request.Body.StartDate,
		EndDate: request.Body.EndDate,
		MinBudget: request.Body.MinBudget,
		MaxBudget: request.Body.MaxBudget,
		IsActive: request.Body.IsActive,
	}

	updatedProject, validationErrors, err := ph.projectService.Update(&inputs, projectID)
	if err != nil {
		return businessapi.PutProject500JSONResponse{Code: http.StatusInternalServerError}, err
	}

	mappedValidationErrors := ph.projectService.MappingValidationErrorStruct(validationErrors)
	startDate := openapi_types.Date{Time: updatedProject.StartDate}
	endDate := openapi_types.Date{Time: updatedProject.EndDate}
	resProject := businessapi.Project{
		Id: updatedProject.ID,
		Title: updatedProject.Title,
		Description: updatedProject.Description,
		StartDate: startDate,
		EndDate: endDate,
		MinBudget: &updatedProject.MinBudget.Int,
		MaxBudget: &updatedProject.MaxBudget.Int,
		IsActive: updatedProject.IsActive,
	}
	return businessapi.PutProject200JSONResponse(businessapi.ProjectStoreResponse{Errors: mappedValidationErrors, Project: resProject}), nil
}

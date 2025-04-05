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
	PostProjects(ctx context.Context, request businessapi.PostProjectsRequestObject) (businessapi.PostProjectsResponseObject, error)
	PutProjectsId(ctx context.Context, request businessapi.PutProjectsIdRequestObject) (businessapi.PutProjectsIdResponseObject, error)
}

type projectsHandler struct {
	projectService businessservices.ProjectService
}

func NewProjectsHandler(projectService businessservices.ProjectService) ProjectsHandler {
	return &projectsHandler{projectService}
}

func (ph *projectsHandler) GetProjects(ctx context.Context, request businessapi.GetProjectsRequestObject) (businessapi.GetProjectsResponseObject, error) {
	companyID, _ := businesshelpers.ExtractCompanyID(ctx)

	projects, err := ph.projectService.FetchLists(ctx, companyID)
	if err != nil {
		res := businessapi.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError}
		return businessapi.GetProjects500JSONResponse{InternalServerErrorResponseJSONResponse: res}, err
	}
	var resProducts []businessapi.Project
	for _, project := range projects { 
		projectID := strconv.Itoa(project.ID)
		resProducts = append(resProducts, businessapi.Project{
			Id: &projectID,
			Title: &project.Title,
			Description: &project.Description,
			StartDate: &openapi_types.Date{Time: project.StartDate},
			EndDate: &openapi_types.Date{Time: project.EndDate},
			MinBudget: &project.MinBudget.Int,
			MaxBudget: &project.MaxBudget.Int,
			IsActive: &project.IsActive,
			CreatedAt: &project.CreatedAt,
		})
	}
	return businessapi.GetProjects200JSONResponse{ProjectsListResponseJSONResponse: businessapi.ProjectsListResponseJSONResponse{
		Projects: resProducts,
	}}, nil
}

func (ph *projectsHandler) PostProjects(ctx context.Context, request businessapi.PostProjectsRequestObject) (businessapi.PostProjectsResponseObject, error) {
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

	createdProject, validationErrors, err := ph.projectService.Create(ctx, &inputs, companyID)
	if err != nil {
		res := businessapi.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError}
		return businessapi.PostProjects500JSONResponse{InternalServerErrorResponseJSONResponse: res}, err
	}

	mappedValidationErrors := ph.projectService.MappingValidationErrorStruct(validationErrors)
	projectID := strconv.Itoa(createdProject.ID)
	startDate := openapi_types.Date{Time: createdProject.StartDate}
	endDate := openapi_types.Date{Time: createdProject.EndDate}
	resProject := businessapi.Project{
		Id: &projectID,
		Title: &createdProject.Title,
		Description: &createdProject.Description,
		StartDate: &startDate,
		EndDate: &endDate,
		MinBudget: &createdProject.MinBudget.Int,
		MaxBudget: &createdProject.MaxBudget.Int,
		IsActive: &createdProject.IsActive,
		CreatedAt: &createdProject.CreatedAt,
	}

	res := businessapi.ProjectStoreResponseJSONResponse{Errors: mappedValidationErrors, Project: resProject}
	return businessapi.PostProjects200JSONResponse{ProjectStoreResponseJSONResponse: res}, nil
}

//lint:ignore ST1003 oapi-codegenの自動生成メソッド名
func (ph *projectsHandler) PutProjectsId(ctx context.Context, request businessapi.PutProjectsIdRequestObject) (businessapi.PutProjectsIdResponseObject, error) {
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

	updatedProject, validationErrors, err := ph.projectService.Update(ctx, &inputs, projectID)
	if err != nil {
		res := businessapi.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError}
		return businessapi.PutProjectsId500JSONResponse{InternalServerErrorResponseJSONResponse: res}, err
	}

	mappedValidationErrors := ph.projectService.MappingValidationErrorStruct(validationErrors)
	projectIDStr := strconv.Itoa(updatedProject.ID)
	startDate := openapi_types.Date{Time: updatedProject.StartDate}
	endDate := openapi_types.Date{Time: updatedProject.EndDate}
	resProject := businessapi.Project{
		Id: &projectIDStr,
		Title: &updatedProject.Title,
		Description: &updatedProject.Description,
		StartDate: &startDate,
		EndDate: &endDate,
		MinBudget: &updatedProject.MinBudget.Int,
		MaxBudget: &updatedProject.MaxBudget.Int,
		IsActive: &updatedProject.IsActive,
	}

	res := businessapi.ProjectStoreResponseJSONResponse{Errors: mappedValidationErrors, Project: resProject}
	return businessapi.PutProjectsId200JSONResponse{ProjectStoreResponseJSONResponse: res}, nil
}

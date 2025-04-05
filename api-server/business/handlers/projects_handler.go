package businesshandlers

import (
	businessapi "apps/api/business"
	businesshelpers "apps/business/helpers"
	businessservices "apps/business/services"
	"context"
	"errors"
	"net/http"
	"strconv"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type ProjectsHandler interface {
	PostProjects(ctx context.Context, request businessapi.PostProjectsRequestObject) (businessapi.PostProjectsResponseObject, error)
}

type projectsHandler struct {
	projectService businessservices.ProjectService
}

func NewProjectsHandler(projectService businessservices.ProjectService) ProjectsHandler {
	return &projectsHandler{projectService}
}

func (ph *projectsHandler) PostProjects(ctx context.Context, request businessapi.PostProjectsRequestObject) (businessapi.PostProjectsResponseObject, error) {
	companyID, ok := businesshelpers.ExtractCompanyID(ctx)
	if !ok {
		res := businessapi.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError}
		return businessapi.PostProjects500JSONResponse{InternalServerErrorResponseJSONResponse: res}, errors.New("fail to load context value")
	}

	inputs := businessapi.PostProjectsJSONRequestBody{
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

package businesshandlers

// import (
// 	"context"
// 	"errors"
// 	"net/http"
// 	"strconv"

// 	openapi_types "github.com/oapi-codegen/runtime/types"
// )

// type ProjectsHandler interface {
// 	PostProjects(ctx context.Context, request projects.PostProjectsRequestObject) (projects.PostProjectsResponseObject, error)
// }

// type projectsHandler struct {
// 	projectService services.ProjectService
// }

// func NewProjectsHandler(projectService services.ProjectService) ProjectsHandler {
// 	return &projectsHandler{projectService}
// }

// func (ph *projectsHandler) PostProjects(ctx context.Context, request projects.PostProjectsRequestObject) (projects.PostProjectsResponseObject, error) {
// 	ctxKey := utils.CompanyCtxKey
// 	companyID, ok := utils.ContextValue(ctx, ctxKey)
// 	if !ok {
// 		res := projects.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError}
// 		return projects.PostProjects500JSONResponse{InternalServerErrorResponseJSONResponse: res}, errors.New("fail to load context value")
// 	}

// 	inputs := projects.PostProjectsJSONRequestBody{
// 		Title: request.Body.Title,
// 		Description: request.Body.Description,
// 		StartDate: request.Body.StartDate,
// 		EndDate: request.Body.EndDate,
// 		MinBudget: request.Body.MinBudget,
// 		MaxBudget: request.Body.MaxBudget,
// 		IsActive: request.Body.IsActive,
// 	}

// 	createdProject, validationErrors, err := ph.projectService.Create(ctx, &inputs, companyID)
// 	if err != nil {
// 		res := projects.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError}
// 		return projects.PostProjects500JSONResponse{InternalServerErrorResponseJSONResponse: res}, err
// 	}

// 	mappedValidationErrors := ph.projectService.MappingValidationErrorStruct(validationErrors)
// 	projectID := strconv.Itoa(createdProject.ID)
// 	startDate := openapi_types.Date{Time: createdProject.StartDate}
// 	endDate := openapi_types.Date{Time: createdProject.EndDate}
// 	resProject := projects.Project{
// 		Id: &projectID,
// 		Title: &createdProject.Title,
// 		Description: &createdProject.Description,
// 		StartDate: &startDate,
// 		EndDate: &endDate,
// 		MinBudget: &createdProject.MinBudget.Int,
// 		MaxBudget: &createdProject.MaxBudget.Int,
// 		IsActive: &createdProject.IsActive,
// 		CreatedAt: &createdProject.CreatedAt,
// 	}

// 	res := projects.ProjectStoreResponseJSONResponse{Errors: mappedValidationErrors, Project: resProject}
// 	return projects.PostProjects200JSONResponse{ProjectStoreResponseJSONResponse: res}, nil
// }

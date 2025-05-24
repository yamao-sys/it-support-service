package businesshandlers

import (
	businessapi "apps/api/business"
	"context"
)

type MainHandler interface {
	// handlers /csrf
	GetCsrf(ctx context.Context, request businessapi.GetCsrfRequestObject) (businessapi.GetCsrfResponseObject, error)

	// handlers /supporters
	PostSupporterSignIn(ctx context.Context, request businessapi.PostSupporterSignInRequestObject) (businessapi.PostSupporterSignInResponseObject, error)

	// handlers /companies
	PostCompanySignIn(ctx context.Context, request businessapi.PostCompanySignInRequestObject) (businessapi.PostCompanySignInResponseObject, error)

	// handlers /projects
	GetProjects(ctx context.Context, request businessapi.GetProjectsRequestObject) (businessapi.GetProjectsResponseObject, error)
	PostProject(ctx context.Context, request businessapi.PostProjectRequestObject) (businessapi.PostProjectResponseObject, error)
	GetProject(ctx context.Context, request businessapi.GetProjectRequestObject) (businessapi.GetProjectResponseObject, error)
	PutProject(ctx context.Context, request businessapi.PutProjectRequestObject) (businessapi.PutProjectResponseObject, error)

	// handlers /plans
	PostPlan(ctx context.Context, request businessapi.PostPlanRequestObject) (businessapi.PostPlanResponseObject, error)

	// handlers /to_projects
	GetToProjects(ctx context.Context, request businessapi.GetToProjectsRequestObject) (businessapi.GetToProjectsResponseObject, error)
	GetToProject(ctx context.Context, request businessapi.GetToProjectRequestObject) (businessapi.GetToProjectResponseObject, error)
}

type mainHandler struct {
	csrfHandler CsrfHandler
	supportersHandler SupportersHandler
	companiesHandler CompaniesHandler
	projectsHandler ProjectsHandler
	plansHandler PlansHandler
	toProjectsHandler ToProjectsHandler
}

func NewMainHandler(
	csrfHandler CsrfHandler,
	supportersHandler SupportersHandler,
	companiesHandler CompaniesHandler,
	projectsHandler ProjectsHandler,
	plansHandler PlansHandler,
	toProjectsHandler ToProjectsHandler,
) MainHandler {
	return &mainHandler{csrfHandler, supportersHandler, companiesHandler, projectsHandler, plansHandler, toProjectsHandler}
}

func (mh *mainHandler) GetCsrf(ctx context.Context, request businessapi.GetCsrfRequestObject) (businessapi.GetCsrfResponseObject, error) {
	res, err := mh.csrfHandler.GetCsrf(ctx, request)
	return res, err
}

func (mh *mainHandler) PostSupporterSignIn(ctx context.Context, request businessapi.PostSupporterSignInRequestObject) (businessapi.PostSupporterSignInResponseObject, error) {
	res, err := mh.supportersHandler.PostSupporterSignIn(ctx, request)
	return res, err
}

func (mh *mainHandler) PostCompanySignIn(ctx context.Context, request businessapi.PostCompanySignInRequestObject) (businessapi.PostCompanySignInResponseObject, error) {
	res, err := mh.companiesHandler.PostCompanySignIn(ctx, request)
	return res, err
}

func (mh *mainHandler) GetProjects(ctx context.Context, request businessapi.GetProjectsRequestObject) (businessapi.GetProjectsResponseObject, error) {
	res, err := mh.projectsHandler.GetProjects(ctx, request)
	return res, err
}

func (mh *mainHandler) PostProject(ctx context.Context, request businessapi.PostProjectRequestObject) (businessapi.PostProjectResponseObject, error) {
	res, err := mh.projectsHandler.PostProject(ctx, request)
	return res, err
}

func (mh *mainHandler) GetProject(ctx context.Context, request businessapi.GetProjectRequestObject) (businessapi.GetProjectResponseObject, error) {
	res, err := mh.projectsHandler.GetProject(ctx, request)
	return res, err
}

func (mh *mainHandler) PutProject(ctx context.Context, request businessapi.PutProjectRequestObject) (businessapi.PutProjectResponseObject, error) {
	res, err := mh.projectsHandler.PutProject(ctx, request)
	return res, err
}

func (mh *mainHandler) PostPlan(ctx context.Context, request businessapi.PostPlanRequestObject) (businessapi.PostPlanResponseObject, error) {
	res, err := mh.plansHandler.PostPlan(ctx, request)
	return res, err
}

func (mh *mainHandler) GetToProjects(ctx context.Context, request businessapi.GetToProjectsRequestObject) (businessapi.GetToProjectsResponseObject, error) {
	res, err := mh.toProjectsHandler.GetToProjects(ctx, request)
	return res, err
}

func (mh *mainHandler) GetToProject(ctx context.Context, request businessapi.GetToProjectRequestObject) (businessapi.GetToProjectResponseObject, error) {
	res, err := mh.toProjectsHandler.GetToProject(ctx, request)
	return res, err
}

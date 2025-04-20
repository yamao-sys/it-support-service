package businesshandlers

import (
	businessapi "apps/api/business"
	"context"
)

type MainHandler interface {
	// handlers /csrf
	GetCsrf(ctx context.Context, request businessapi.GetCsrfRequestObject) (businessapi.GetCsrfResponseObject, error)

	// handlers /supporters
	PostSupportersSignIn(ctx context.Context, request businessapi.PostSupportersSignInRequestObject) (businessapi.PostSupportersSignInResponseObject, error)

	// handlers /companies
	PostCompaniesSignIn(ctx context.Context, request businessapi.PostCompaniesSignInRequestObject) (businessapi.PostCompaniesSignInResponseObject, error)

	// handlers /projects
	GetProjects(ctx context.Context, request businessapi.GetProjectsRequestObject) (businessapi.GetProjectsResponseObject, error)
	PostProjects(ctx context.Context, request businessapi.PostProjectsRequestObject) (businessapi.PostProjectsResponseObject, error)
	GetProjectsId(ctx context.Context, request businessapi.GetProjectsIdRequestObject) (businessapi.GetProjectsIdResponseObject, error)
	PutProjectsId(ctx context.Context, request businessapi.PutProjectsIdRequestObject) (businessapi.PutProjectsIdResponseObject, error)

	// handlers /plans
	PostPlans(ctx context.Context, request businessapi.PostPlansRequestObject) (businessapi.PostPlansResponseObject, error)
}

type mainHandler struct {
	csrfHandler CsrfHandler
	supportersHandler SupportersHandler
	companiesHandler CompaniesHandler
	projectsHandler ProjectsHandler
	plansHandler PlansHandler
}

func NewMainHandler(
	csrfHandler CsrfHandler,
	supportersHandler SupportersHandler,
	companiesHandler CompaniesHandler,
	projectsHandler ProjectsHandler,
	plansHandler PlansHandler,
) MainHandler {
	return &mainHandler{csrfHandler, supportersHandler, companiesHandler, projectsHandler, plansHandler}
}

func (mh *mainHandler) GetCsrf(ctx context.Context, request businessapi.GetCsrfRequestObject) (businessapi.GetCsrfResponseObject, error) {
	res, err := mh.csrfHandler.GetCsrf(ctx, request)
	return res, err
}

func (mh *mainHandler) PostSupportersSignIn(ctx context.Context, request businessapi.PostSupportersSignInRequestObject) (businessapi.PostSupportersSignInResponseObject, error) {
	res, err := mh.supportersHandler.PostSupportersSignIn(ctx, request)
	return res, err
}

func (mh *mainHandler) PostCompaniesSignIn(ctx context.Context, request businessapi.PostCompaniesSignInRequestObject) (businessapi.PostCompaniesSignInResponseObject, error) {
	res, err := mh.companiesHandler.PostCompaniesSignIn(ctx, request)
	return res, err
}

func (mh *mainHandler) GetProjects(ctx context.Context, request businessapi.GetProjectsRequestObject) (businessapi.GetProjectsResponseObject, error) {
	res, err := mh.projectsHandler.GetProjects(ctx, request)
	return res, err
}

func (mh *mainHandler) PostProjects(ctx context.Context, request businessapi.PostProjectsRequestObject) (businessapi.PostProjectsResponseObject, error) {
	res, err := mh.projectsHandler.PostProjects(ctx, request)
	return res, err
}

//lint:ignore ST1003 oapi-codegenの自動生成メソッド名
func (mh *mainHandler) GetProjectsId(ctx context.Context, request businessapi.GetProjectsIdRequestObject) (businessapi.GetProjectsIdResponseObject, error) {
	res, err := mh.projectsHandler.GetProjectsId(ctx, request)
	return res, err
}

//lint:ignore ST1003 oapi-codegenの自動生成メソッド名
func (mh *mainHandler) PutProjectsId(ctx context.Context, request businessapi.PutProjectsIdRequestObject) (businessapi.PutProjectsIdResponseObject, error) {
	res, err := mh.projectsHandler.PutProjectsId(ctx, request)
	return res, err
}

func (mh *mainHandler) PostPlans(ctx context.Context, request businessapi.PostPlansRequestObject) (businessapi.PostPlansResponseObject, error) {
	res, err := mh.plansHandler.PostPlans(ctx, request)
	return res, err
}

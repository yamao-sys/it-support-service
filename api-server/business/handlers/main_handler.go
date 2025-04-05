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
	PostProjects(ctx context.Context, request businessapi.PostProjectsRequestObject) (businessapi.PostProjectsResponseObject, error)
}

type mainHandler struct {
	csrfHandler CsrfHandler
	supportersHandler SupportersHandler
	companiesHandler CompaniesHandler
	projectsHandler ProjectsHandler
}

func NewMainHandler(
	csrfHandler CsrfHandler,
	supportersHandler SupportersHandler,
	companiesHandler CompaniesHandler,
	projectsHandler ProjectsHandler,
) MainHandler {
	return &mainHandler{csrfHandler, supportersHandler, companiesHandler, projectsHandler}
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

func (mh *mainHandler) PostProjects(ctx context.Context, request businessapi.PostProjectsRequestObject) (businessapi.PostProjectsResponseObject, error) {
	res, err := mh.projectsHandler.PostProjects(ctx, request)
	return res, err
}

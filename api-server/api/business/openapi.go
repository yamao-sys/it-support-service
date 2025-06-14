// Package businessapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package businessapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// Defines values for ToProjectProposalStatus.
const (
	NOTPROPOSED       ToProjectProposalStatus = "NOT PROPOSED"
	PROPOSED          ToProjectProposalStatus = "PROPOSED"
	TEMPORARYCREATING ToProjectProposalStatus = "TEMPORARY CREATING"
)

// CompanySignInBadRequestResponse defines model for CompanySignInBadRequestResponse.
type CompanySignInBadRequestResponse struct {
	Errors []string `json:"errors"`
}

// CompanySignInInput defines model for CompanySignInInput.
type CompanySignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CompanySignInOkResponse defines model for CompanySignInOkResponse.
type CompanySignInOkResponse struct {
	Token string `json:"token"`
}

// CsrfResponse defines model for CsrfResponse.
type CsrfResponse struct {
	CsrfToken string `json:"csrfToken"`
}

// Plan defines model for Plan.
type Plan struct {
	CreatedAt   time.Time           `json:"createdAt"`
	Description string              `json:"description"`
	EndDate     *openapi_types.Date `json:"endDate,omitempty"`
	Id          int                 `json:"id"`
	ProjectId   int                 `json:"projectId"`
	StartDate   *openapi_types.Date `json:"startDate,omitempty"`
	Title       string              `json:"title"`
	UnitPrice   int                 `json:"unitPrice"`
}

// PlanStoreInput defines model for PlanStoreInput.
type PlanStoreInput struct {
	Description string              `json:"description"`
	EndDate     *openapi_types.Date `json:"endDate,omitempty"`
	ProjectId   int                 `json:"projectId"`
	StartDate   *openapi_types.Date `json:"startDate,omitempty"`
	Title       string              `json:"title"`
	UnitPrice   int                 `json:"unitPrice"`
}

// PlanStoreResponse defines model for PlanStoreResponse.
type PlanStoreResponse struct {
	Errors PlanValidationError `json:"errors"`
	Plan   Plan                `json:"plan"`
}

// PlanValidationError defines model for PlanValidationError.
type PlanValidationError struct {
	Description *[]string `json:"description,omitempty"`
	EndDate     *[]string `json:"endDate,omitempty"`
	StartDate   *[]string `json:"startDate,omitempty"`
	Title       *[]string `json:"title,omitempty"`
	UnitPrice   *[]string `json:"unitPrice,omitempty"`
}

// Project defines model for Project.
type Project struct {
	CreatedAt   time.Time          `json:"createdAt"`
	Description string             `json:"description"`
	EndDate     openapi_types.Date `json:"endDate"`
	Id          int                `json:"id"`
	IsActive    bool               `json:"isActive"`
	MaxBudget   *int               `json:"maxBudget,omitempty"`
	MinBudget   *int               `json:"minBudget,omitempty"`
	StartDate   openapi_types.Date `json:"startDate"`
	Title       string             `json:"title"`
}

// ProjectResponse defines model for ProjectResponse.
type ProjectResponse struct {
	Project Project `json:"project"`
}

// ProjectStoreInput defines model for ProjectStoreInput.
type ProjectStoreInput struct {
	Description string              `json:"description"`
	EndDate     *openapi_types.Date `json:"endDate,omitempty"`
	IsActive    bool                `json:"isActive"`
	MaxBudget   *int                `json:"maxBudget,omitempty"`
	MinBudget   *int                `json:"minBudget,omitempty"`
	StartDate   *openapi_types.Date `json:"startDate,omitempty"`
	Title       string              `json:"title"`
}

// ProjectStoreResponse defines model for ProjectStoreResponse.
type ProjectStoreResponse struct {
	Errors  ProjectValidationError `json:"errors"`
	Project Project                `json:"project"`
}

// ProjectValidationError defines model for ProjectValidationError.
type ProjectValidationError struct {
	Description *[]string `json:"description,omitempty"`
	EndDate     *[]string `json:"endDate,omitempty"`
	IsActive    *[]string `json:"isActive,omitempty"`
	MaxBudget   *[]string `json:"maxBudget,omitempty"`
	MinBudget   *[]string `json:"minBudget,omitempty"`
	StartDate   *[]string `json:"startDate,omitempty"`
	Title       *[]string `json:"title,omitempty"`
}

// ProjectsListResponse defines model for ProjectsListResponse.
type ProjectsListResponse struct {
	NextPageToken string    `json:"nextPageToken"`
	Projects      []Project `json:"projects"`
}

// SupporterSignInBadRequestResponse defines model for SupporterSignInBadRequestResponse.
type SupporterSignInBadRequestResponse struct {
	Errors []string `json:"errors"`
}

// SupporterSignInInput defines model for SupporterSignInInput.
type SupporterSignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SupporterSignInOkResponse defines model for SupporterSignInOkResponse.
type SupporterSignInOkResponse struct {
	Token string `json:"token"`
}

// ToProject defines model for ToProject.
type ToProject struct {
	Description    string                  `json:"description"`
	EndDate        openapi_types.Date      `json:"endDate"`
	Id             int                     `json:"id"`
	MaxBudget      *int                    `json:"maxBudget,omitempty"`
	MinBudget      *int                    `json:"minBudget,omitempty"`
	ProposalStatus ToProjectProposalStatus `json:"proposalStatus"`
	StartDate      openapi_types.Date      `json:"startDate"`
	Title          string                  `json:"title"`
}

// ToProjectProposalStatus defines model for ToProject.ProposalStatus.
type ToProjectProposalStatus string

// ToProjectResponse defines model for ToProjectResponse.
type ToProjectResponse struct {
	Project ToProject `json:"project"`
}

// ToProjectsListResponse defines model for ToProjectsListResponse.
type ToProjectsListResponse struct {
	NextPageToken string      `json:"nextPageToken"`
	Projects      []ToProject `json:"projects"`
}

// GetProjectsParams defines parameters for GetProjects.
type GetProjectsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`
}

// GetToProjectsParams defines parameters for GetToProjects.
type GetToProjectsParams struct {
	PageToken *string             `form:"pageToken,omitempty" json:"pageToken,omitempty"`
	StartDate *openapi_types.Date `form:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate   *openapi_types.Date `form:"endDate,omitempty" json:"endDate,omitempty"`
}

// PostCompanySignInJSONRequestBody defines body for PostCompanySignIn for application/json ContentType.
type PostCompanySignInJSONRequestBody = CompanySignInInput

// PostPlanJSONRequestBody defines body for PostPlan for application/json ContentType.
type PostPlanJSONRequestBody = PlanStoreInput

// PostProjectJSONRequestBody defines body for PostProject for application/json ContentType.
type PostProjectJSONRequestBody = ProjectStoreInput

// PutProjectJSONRequestBody defines body for PutProject for application/json ContentType.
type PutProjectJSONRequestBody = ProjectStoreInput

// PostSupporterSignInJSONRequestBody defines body for PostSupporterSignIn for application/json ContentType.
type PostSupporterSignInJSONRequestBody = SupporterSignInInput

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Company Sign In
	// (POST /companies/sign-in)
	PostCompanySignIn(ctx echo.Context) error
	// Get Csrf
	// (GET /csrf)
	GetCsrf(ctx echo.Context) error
	// Create Plan
	// (POST /plans)
	PostPlan(ctx echo.Context) error
	// Create Project
	// (GET /projects)
	GetProjects(ctx echo.Context, params GetProjectsParams) error
	// Create Project
	// (POST /projects)
	PostProject(ctx echo.Context) error
	// Get Project
	// (GET /projects/{id})
	GetProject(ctx echo.Context, id int) error
	// Update Project
	// (PUT /projects/{id})
	PutProject(ctx echo.Context, id int) error
	// Supporter Sign In
	// (POST /supporters/sign-in)
	PostSupporterSignIn(ctx echo.Context) error
	// Get Projects for Supporters
	// (GET /to-projects)
	GetToProjects(ctx echo.Context, params GetToProjectsParams) error
	// Get Project for Supporters
	// (GET /to-projects/{id})
	GetToProject(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostCompanySignIn converts echo context to params.
func (w *ServerInterfaceWrapper) PostCompanySignIn(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostCompanySignIn(ctx)
	return err
}

// GetCsrf converts echo context to params.
func (w *ServerInterfaceWrapper) GetCsrf(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetCsrf(ctx)
	return err
}

// PostPlan converts echo context to params.
func (w *ServerInterfaceWrapper) PostPlan(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostPlan(ctx)
	return err
}

// GetProjects converts echo context to params.
func (w *ServerInterfaceWrapper) GetProjects(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProjectsParams
	// ------------- Optional query parameter "pageToken" -------------

	err = runtime.BindQueryParameter("form", false, false, "pageToken", ctx.QueryParams(), &params.PageToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageToken: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProjects(ctx, params)
	return err
}

// PostProject converts echo context to params.
func (w *ServerInterfaceWrapper) PostProject(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostProject(ctx)
	return err
}

// GetProject converts echo context to params.
func (w *ServerInterfaceWrapper) GetProject(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProject(ctx, id)
	return err
}

// PutProject converts echo context to params.
func (w *ServerInterfaceWrapper) PutProject(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutProject(ctx, id)
	return err
}

// PostSupporterSignIn converts echo context to params.
func (w *ServerInterfaceWrapper) PostSupporterSignIn(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostSupporterSignIn(ctx)
	return err
}

// GetToProjects converts echo context to params.
func (w *ServerInterfaceWrapper) GetToProjects(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetToProjectsParams
	// ------------- Optional query parameter "pageToken" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageToken", ctx.QueryParams(), &params.PageToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageToken: %s", err))
	}

	// ------------- Optional query parameter "startDate" -------------

	err = runtime.BindQueryParameter("form", true, false, "startDate", ctx.QueryParams(), &params.StartDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter startDate: %s", err))
	}

	// ------------- Optional query parameter "endDate" -------------

	err = runtime.BindQueryParameter("form", true, false, "endDate", ctx.QueryParams(), &params.EndDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter endDate: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetToProjects(ctx, params)
	return err
}

// GetToProject converts echo context to params.
func (w *ServerInterfaceWrapper) GetToProject(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetToProject(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/companies/sign-in", wrapper.PostCompanySignIn)
	router.GET(baseURL+"/csrf", wrapper.GetCsrf)
	router.POST(baseURL+"/plans", wrapper.PostPlan)
	router.GET(baseURL+"/projects", wrapper.GetProjects)
	router.POST(baseURL+"/projects", wrapper.PostProject)
	router.GET(baseURL+"/projects/:id", wrapper.GetProject)
	router.PUT(baseURL+"/projects/:id", wrapper.PutProject)
	router.POST(baseURL+"/supporters/sign-in", wrapper.PostSupporterSignIn)
	router.GET(baseURL+"/to-projects", wrapper.GetToProjects)
	router.GET(baseURL+"/to-projects/:id", wrapper.GetToProject)

}

type PostCompanySignInRequestObject struct {
	Body *PostCompanySignInJSONRequestBody
}

type PostCompanySignInResponseObject interface {
	VisitPostCompanySignInResponse(w http.ResponseWriter) error
}

type PostCompanySignIn200ResponseHeaders struct {
	SetCookie string
}

type PostCompanySignIn200JSONResponse struct {
	Body    CompanySignInOkResponse
	Headers PostCompanySignIn200ResponseHeaders
}

func (response PostCompanySignIn200JSONResponse) VisitPostCompanySignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Set-Cookie", fmt.Sprint(response.Headers.SetCookie))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostCompanySignIn400JSONResponse CompanySignInBadRequestResponse

func (response PostCompanySignIn400JSONResponse) VisitPostCompanySignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostCompanySignIn500Response struct {
}

func (response PostCompanySignIn500Response) VisitPostCompanySignInResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetCsrfRequestObject struct {
}

type GetCsrfResponseObject interface {
	VisitGetCsrfResponse(w http.ResponseWriter) error
}

type GetCsrf200JSONResponse CsrfResponse

func (response GetCsrf200JSONResponse) VisitGetCsrfResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetCsrf500Response struct {
}

func (response GetCsrf500Response) VisitGetCsrfResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type PostPlanRequestObject struct {
	Body *PostPlanJSONRequestBody
}

type PostPlanResponseObject interface {
	VisitPostPlanResponse(w http.ResponseWriter) error
}

type PostPlan200JSONResponse PlanStoreResponse

func (response PostPlan200JSONResponse) VisitPostPlanResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostPlan500Response struct {
}

func (response PostPlan500Response) VisitPostPlanResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetProjectsRequestObject struct {
	Params GetProjectsParams
}

type GetProjectsResponseObject interface {
	VisitGetProjectsResponse(w http.ResponseWriter) error
}

type GetProjects200JSONResponse ProjectsListResponse

func (response GetProjects200JSONResponse) VisitGetProjectsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetProjects500Response struct {
}

func (response GetProjects500Response) VisitGetProjectsResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type PostProjectRequestObject struct {
	Body *PostProjectJSONRequestBody
}

type PostProjectResponseObject interface {
	VisitPostProjectResponse(w http.ResponseWriter) error
}

type PostProject200JSONResponse ProjectStoreResponse

func (response PostProject200JSONResponse) VisitPostProjectResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostProject500Response struct {
}

func (response PostProject500Response) VisitPostProjectResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetProjectRequestObject struct {
	Id int `json:"id"`
}

type GetProjectResponseObject interface {
	VisitGetProjectResponse(w http.ResponseWriter) error
}

type GetProject200JSONResponse ProjectResponse

func (response GetProject200JSONResponse) VisitGetProjectResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetProject404Response struct {
}

func (response GetProject404Response) VisitGetProjectResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetProject500Response struct {
}

func (response GetProject500Response) VisitGetProjectResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type PutProjectRequestObject struct {
	Id   int `json:"id"`
	Body *PutProjectJSONRequestBody
}

type PutProjectResponseObject interface {
	VisitPutProjectResponse(w http.ResponseWriter) error
}

type PutProject200JSONResponse ProjectStoreResponse

func (response PutProject200JSONResponse) VisitPutProjectResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PutProject404Response struct {
}

func (response PutProject404Response) VisitPutProjectResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PutProject500Response struct {
}

func (response PutProject500Response) VisitPutProjectResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type PostSupporterSignInRequestObject struct {
	Body *PostSupporterSignInJSONRequestBody
}

type PostSupporterSignInResponseObject interface {
	VisitPostSupporterSignInResponse(w http.ResponseWriter) error
}

type PostSupporterSignIn200ResponseHeaders struct {
	SetCookie string
}

type PostSupporterSignIn200JSONResponse struct {
	Body    SupporterSignInOkResponse
	Headers PostSupporterSignIn200ResponseHeaders
}

func (response PostSupporterSignIn200JSONResponse) VisitPostSupporterSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Set-Cookie", fmt.Sprint(response.Headers.SetCookie))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostSupporterSignIn400JSONResponse SupporterSignInBadRequestResponse

func (response PostSupporterSignIn400JSONResponse) VisitPostSupporterSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostSupporterSignIn500Response struct {
}

func (response PostSupporterSignIn500Response) VisitPostSupporterSignInResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetToProjectsRequestObject struct {
	Params GetToProjectsParams
}

type GetToProjectsResponseObject interface {
	VisitGetToProjectsResponse(w http.ResponseWriter) error
}

type GetToProjects200JSONResponse ToProjectsListResponse

func (response GetToProjects200JSONResponse) VisitGetToProjectsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetToProjects403Response struct {
}

func (response GetToProjects403Response) VisitGetToProjectsResponse(w http.ResponseWriter) error {
	w.WriteHeader(403)
	return nil
}

type GetToProjects500Response struct {
}

func (response GetToProjects500Response) VisitGetToProjectsResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetToProjectRequestObject struct {
	Id int `json:"id"`
}

type GetToProjectResponseObject interface {
	VisitGetToProjectResponse(w http.ResponseWriter) error
}

type GetToProject200JSONResponse ToProjectResponse

func (response GetToProject200JSONResponse) VisitGetToProjectResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetToProject403Response struct {
}

func (response GetToProject403Response) VisitGetToProjectResponse(w http.ResponseWriter) error {
	w.WriteHeader(403)
	return nil
}

type GetToProject404Response struct {
}

func (response GetToProject404Response) VisitGetToProjectResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetToProject500Response struct {
}

func (response GetToProject500Response) VisitGetToProjectResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Company Sign In
	// (POST /companies/sign-in)
	PostCompanySignIn(ctx context.Context, request PostCompanySignInRequestObject) (PostCompanySignInResponseObject, error)
	// Get Csrf
	// (GET /csrf)
	GetCsrf(ctx context.Context, request GetCsrfRequestObject) (GetCsrfResponseObject, error)
	// Create Plan
	// (POST /plans)
	PostPlan(ctx context.Context, request PostPlanRequestObject) (PostPlanResponseObject, error)
	// Create Project
	// (GET /projects)
	GetProjects(ctx context.Context, request GetProjectsRequestObject) (GetProjectsResponseObject, error)
	// Create Project
	// (POST /projects)
	PostProject(ctx context.Context, request PostProjectRequestObject) (PostProjectResponseObject, error)
	// Get Project
	// (GET /projects/{id})
	GetProject(ctx context.Context, request GetProjectRequestObject) (GetProjectResponseObject, error)
	// Update Project
	// (PUT /projects/{id})
	PutProject(ctx context.Context, request PutProjectRequestObject) (PutProjectResponseObject, error)
	// Supporter Sign In
	// (POST /supporters/sign-in)
	PostSupporterSignIn(ctx context.Context, request PostSupporterSignInRequestObject) (PostSupporterSignInResponseObject, error)
	// Get Projects for Supporters
	// (GET /to-projects)
	GetToProjects(ctx context.Context, request GetToProjectsRequestObject) (GetToProjectsResponseObject, error)
	// Get Project for Supporters
	// (GET /to-projects/{id})
	GetToProject(ctx context.Context, request GetToProjectRequestObject) (GetToProjectResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// PostCompanySignIn operation middleware
func (sh *strictHandler) PostCompanySignIn(ctx echo.Context) error {
	var request PostCompanySignInRequestObject

	var body PostCompanySignInJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostCompanySignIn(ctx.Request().Context(), request.(PostCompanySignInRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostCompanySignIn")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostCompanySignInResponseObject); ok {
		return validResponse.VisitPostCompanySignInResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetCsrf operation middleware
func (sh *strictHandler) GetCsrf(ctx echo.Context) error {
	var request GetCsrfRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetCsrf(ctx.Request().Context(), request.(GetCsrfRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetCsrf")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetCsrfResponseObject); ok {
		return validResponse.VisitGetCsrfResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostPlan operation middleware
func (sh *strictHandler) PostPlan(ctx echo.Context) error {
	var request PostPlanRequestObject

	var body PostPlanJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostPlan(ctx.Request().Context(), request.(PostPlanRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostPlan")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostPlanResponseObject); ok {
		return validResponse.VisitPostPlanResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetProjects operation middleware
func (sh *strictHandler) GetProjects(ctx echo.Context, params GetProjectsParams) error {
	var request GetProjectsRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProjects(ctx.Request().Context(), request.(GetProjectsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProjects")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetProjectsResponseObject); ok {
		return validResponse.VisitGetProjectsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostProject operation middleware
func (sh *strictHandler) PostProject(ctx echo.Context) error {
	var request PostProjectRequestObject

	var body PostProjectJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProject(ctx.Request().Context(), request.(PostProjectRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProject")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProjectResponseObject); ok {
		return validResponse.VisitPostProjectResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetProject operation middleware
func (sh *strictHandler) GetProject(ctx echo.Context, id int) error {
	var request GetProjectRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProject(ctx.Request().Context(), request.(GetProjectRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProject")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetProjectResponseObject); ok {
		return validResponse.VisitGetProjectResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PutProject operation middleware
func (sh *strictHandler) PutProject(ctx echo.Context, id int) error {
	var request PutProjectRequestObject

	request.Id = id

	var body PutProjectJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PutProject(ctx.Request().Context(), request.(PutProjectRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutProject")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PutProjectResponseObject); ok {
		return validResponse.VisitPutProjectResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostSupporterSignIn operation middleware
func (sh *strictHandler) PostSupporterSignIn(ctx echo.Context) error {
	var request PostSupporterSignInRequestObject

	var body PostSupporterSignInJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostSupporterSignIn(ctx.Request().Context(), request.(PostSupporterSignInRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostSupporterSignIn")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostSupporterSignInResponseObject); ok {
		return validResponse.VisitPostSupporterSignInResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetToProjects operation middleware
func (sh *strictHandler) GetToProjects(ctx echo.Context, params GetToProjectsParams) error {
	var request GetToProjectsRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetToProjects(ctx.Request().Context(), request.(GetToProjectsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetToProjects")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetToProjectsResponseObject); ok {
		return validResponse.VisitGetToProjectsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetToProject operation middleware
func (sh *strictHandler) GetToProject(ctx echo.Context, id int) error {
	var request GetToProjectRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetToProject(ctx.Request().Context(), request.(GetToProjectRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetToProject")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetToProjectResponseObject); ok {
		return validResponse.VisitGetToProjectResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xZUW/bNhD+KwS3RyVO1w4o9OakQRFsa4w4GzAUfmCks81GIlWSymIE/u8DSVmiJMpS",
	"0sp1hr05ku743d13d7zLE454mnEGTEkcPmEZrSEl5ucFTzPCNnO6YlfsnMQ38DUHqW5AZpxJ0J9kgmcg",
	"FAUjAEJwYX5RBan5oTYZ4BBLJShb4W2we0CEIBu83QZYwNecCohx+HmnYBFgRVWivyswIAsCVSgu9aeo",
	"xFIq5ndfIFL6pBr8K5blyoM4JTTx4syIlP9wEXteNkEbHY5EN3yLog/r9X23ixW/B9aPyX7WDeT6fr/r",
	"pFh2Y4ikWN4Ow1F96mJxtXtOnyWEeU4VQBTEUxPFJRcpUTjEMVFwomjqKKpiGIOMBM0U5cwbY2DxB6Kg",
	"pdCni7pMoEzBCoThieAa91XHa6mIUIMPKTzkgZozqmaCRuA7puF0GmMX105t3R+uysBxrhMnE4eO+MwV",
	"F9CRVN/T7Ufo3ud4tuFNZNzWXQZKzw6psT8LWOIQ/zSpCvikqN4TregvktCYaEimVhpvFpnVJ9q22VKh",
	"XZ8dq/pSuomnjzdDe0iNUMOFatQZLlZyaLhIjVzP6Iw1H1feQ9Z9Pi9bZr6a2knlNFL0wU26O84T0AQM",
	"cEoez/N4BcovnFK27/X3qQy+yurP+eq8yjeOhV0ltohYdzC7K0FWRXtvPhefdZQxD5r9mWy/OVT9f3UU",
	"8bOjNMPj7b6e4Dj8m9uC1eXrDN+HTN4WUbNzALeOrFG4FBwuVaPmM8Rcyh5fM9u2w/qMxiR/p/vGRwaP",
	"akZW0DVblCytj5iD+NozeJaKgwaKNo8l0lbs5fE8zzIuFIgfOjqXKF4yPDdM+EHjc8uEzirZwDvGCN0C",
	"0zNE3/LO+9ghLlff1hs1Xi5JMldE5TbWLE+1Xz5d36LZzfXsen75AQf49vKP2fXN9OZvdHFzOb29+vQR",
	"B7h8v/AAPo6bWcNAT79acoHKmMu9Ef7mS1rFlRdc056D9PAl2DFttCLc6wFNO4hyQdVmrnFZa6cZ/Q02",
	"01ytjTUMhzji/J5qfjCSagW2ClTAjQTeaoWULXkrmfFdLikDKdF0diVxBfx893wO4sEuXR5ASCv05vRM",
	"e4dnwEhGcYjfnupHuiiqtYFqPEsYBTmRdMVOqN2QcWmopUNoWvBVrH3EpaotFLH1M0h1zuONGQg5U8CM",
	"LMmyhEZGevJF2oJkY9cXWc+CdVuPqRI5mAc2TMaSX87OxkHg1HwDox6X2zWgwgloTSSSeRQBxBCf4gCv",
	"gcRg2+8c1MmFJUH41DAlcGA1y5E+8d1YpnluEB0mShAPIFDE8yRGjCuUM22ZIixGynFBnANSHFH2oK9v",
	"SG6YIo+nmoS/WiPqqudWLdhxQWdTnqZEbBqbZWS4pshKmuXvjrJ4oUUmkRRLrbpoOnXSfgR1od+PyRZ3",
	"7/w8irzALx9BocKi0iH6T+uLLCFM7k/hYv06RuY2NrgHztr2lnOsYBQlH4ef68X+82K7qHHY7GbQbuFd",
	"hMvGqIiX0/a6+Dur+ldGBEnB9CJ9NjxmCY8Bh0uSSAhsq/mag9hUnSYr292+MrMYMy6+K8KxhKZalO2i",
	"s/P2Qt9KurOoFBwlkVrbsEPnkm879Bpi5ibV5InG2wGZ1U4sk0j6klTlkZkEevt29c+dA2TUywPz7uxd",
	"OzBunydMN/klrbd3iJEAyXMRwRjx1a2tLyFzXz7m40fy/xx/VVT6M4sHlQpZDnfD5p/GPmik8u/dkh2Y",
	"Hd2br//AFNS/ST2aOai+HqxPQhV7CzorfjLkRlmtbToK5ovukIFf2N2TVcI9C7ouZdWibbiqMTtxxwLs",
	"BVX0bZsV0ygCKRGVaMnFHY1jYCO3XenZeBVcc6nVIlvvTava1r22u1Z7GTtucI/+ajaYIka/PtDGORcJ",
	"DvFaqUyGk0nCI5KsuVTh+7P3b7A+rVDztGNBtefR1WD3UIql+7dTAp2nJQz3mRm7nQcu3O1i+28AAAD/",
	"/4cuSjayKgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

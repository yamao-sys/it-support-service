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
	BusinessAuthenticationScopes = "businessAuthentication.Scopes"
)

// Plan defines model for Plan.
type Plan struct {
	CreatedAt   *time.Time          `json:"createdAt,omitempty"`
	Description *string             `json:"description,omitempty"`
	EndDate     *openapi_types.Date `json:"endDate,omitempty"`
	Id          *string             `json:"id,omitempty"`
	ProjectId   *string             `json:"projectId,omitempty"`
	StartDate   *openapi_types.Date `json:"startDate,omitempty"`
	Title       *string             `json:"title,omitempty"`
	UnitPrice   *int                `json:"unitPrice,omitempty"`
}

// PlanValidationError defines model for PlanValidationError.
type PlanValidationError struct {
	Description *[]string `json:"description,omitempty"`
	EndDate     *[]string `json:"endDate,omitempty"`
	StartDate   *[]string `json:"startDate,omitempty"`
	Title       *[]string `json:"title,omitempty"`
	UnitPrice   *[]string `json:"unitPrice,omitempty"`
}

// Project Project
type Project struct {
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description *string             `json:"description,omitempty"`
	EndDate     *openapi_types.Date `json:"end_date,omitempty"`
	Id          *string             `json:"id,omitempty"`
	IsActive    *bool               `json:"isActive,omitempty"`
	MaxBudget   *int                `json:"max_budget,omitempty"`
	MinBudget   *int                `json:"min_budget,omitempty"`
	StartDate   *openapi_types.Date `json:"start_date,omitempty"`
	Title       *string             `json:"title,omitempty"`
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

// CompanySignInBadRequestResponse defines model for CompanySignInBadRequestResponse.
type CompanySignInBadRequestResponse struct {
	Errors []string `json:"errors"`
}

// CompanySignInOkResponse defines model for CompanySignInOkResponse.
type CompanySignInOkResponse = map[string]interface{}

// CsrfResponse defines model for CsrfResponse.
type CsrfResponse struct {
	CsrfToken string `json:"csrfToken"`
}

// InternalServerErrorResponse defines model for InternalServerErrorResponse.
type InternalServerErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NotFoundErrorResponse defines model for NotFoundErrorResponse.
type NotFoundErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// PlanStoreResponse defines model for PlanStoreResponse.
type PlanStoreResponse struct {
	Errors PlanValidationError `json:"errors"`
	Plan   Plan                `json:"plan"`
}

// ProjectResponse defines model for ProjectResponse.
type ProjectResponse struct {
	// Project Project
	Project Project `json:"project"`
}

// ProjectStoreResponse defines model for ProjectStoreResponse.
type ProjectStoreResponse struct {
	Errors ProjectValidationError `json:"errors"`

	// Project Project
	Project Project `json:"project"`
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

// SupporterSignInOkResponse defines model for SupporterSignInOkResponse.
type SupporterSignInOkResponse = map[string]interface{}

// CompanySignInInput defines model for CompanySignInInput.
type CompanySignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

// ProjectStoreInput defines model for ProjectStoreInput.
type ProjectStoreInput struct {
	Description *string             `json:"description,omitempty"`
	EndDate     *openapi_types.Date `json:"endDate,omitempty"`
	IsActive    *bool               `json:"isActive,omitempty"`
	MaxBudget   *int                `json:"maxBudget,omitempty"`
	MinBudget   *int                `json:"minBudget,omitempty"`
	StartDate   *openapi_types.Date `json:"startDate,omitempty"`
	Title       *string             `json:"title,omitempty"`
}

// SupporterSignInInput defines model for SupporterSignInInput.
type SupporterSignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// PostCompaniesSignInJSONBody defines parameters for PostCompaniesSignIn.
type PostCompaniesSignInJSONBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// PostPlansJSONBody defines parameters for PostPlans.
type PostPlansJSONBody struct {
	Description string              `json:"description"`
	EndDate     *openapi_types.Date `json:"endDate,omitempty"`
	ProjectId   int                 `json:"projectId"`
	StartDate   *openapi_types.Date `json:"startDate,omitempty"`
	Title       string              `json:"title"`
	UnitPrice   int                 `json:"unitPrice"`
}

// GetProjectsParams defines parameters for GetProjects.
type GetProjectsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`
}

// PostProjectsJSONBody defines parameters for PostProjects.
type PostProjectsJSONBody struct {
	Description *string             `json:"description,omitempty"`
	EndDate     *openapi_types.Date `json:"endDate,omitempty"`
	IsActive    *bool               `json:"isActive,omitempty"`
	MaxBudget   *int                `json:"maxBudget,omitempty"`
	MinBudget   *int                `json:"minBudget,omitempty"`
	StartDate   *openapi_types.Date `json:"startDate,omitempty"`
	Title       *string             `json:"title,omitempty"`
}

// PutProjectsIdJSONBody defines parameters for PutProjectsId.
type PutProjectsIdJSONBody struct {
	Description *string             `json:"description,omitempty"`
	EndDate     *openapi_types.Date `json:"endDate,omitempty"`
	IsActive    *bool               `json:"isActive,omitempty"`
	MaxBudget   *int                `json:"maxBudget,omitempty"`
	MinBudget   *int                `json:"minBudget,omitempty"`
	StartDate   *openapi_types.Date `json:"startDate,omitempty"`
	Title       *string             `json:"title,omitempty"`
}

// PostSupportersSignInJSONBody defines parameters for PostSupportersSignIn.
type PostSupportersSignInJSONBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// PostCompaniesSignInJSONRequestBody defines body for PostCompaniesSignIn for application/json ContentType.
type PostCompaniesSignInJSONRequestBody PostCompaniesSignInJSONBody

// PostPlansJSONRequestBody defines body for PostPlans for application/json ContentType.
type PostPlansJSONRequestBody PostPlansJSONBody

// PostProjectsJSONRequestBody defines body for PostProjects for application/json ContentType.
type PostProjectsJSONRequestBody PostProjectsJSONBody

// PutProjectsIdJSONRequestBody defines body for PutProjectsId for application/json ContentType.
type PutProjectsIdJSONRequestBody PutProjectsIdJSONBody

// PostSupportersSignInJSONRequestBody defines body for PostSupportersSignIn for application/json ContentType.
type PostSupportersSignInJSONRequestBody PostSupportersSignInJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// SignIn
	// (POST /companies/signIn)
	PostCompaniesSignIn(ctx echo.Context) error
	// Get Csrf
	// (GET /csrf)
	GetCsrf(ctx echo.Context) error
	// Plan Create
	// (POST /plans)
	PostPlans(ctx echo.Context) error
	// Project List
	// (GET /projects)
	GetProjects(ctx echo.Context, params GetProjectsParams) error
	// Project Create
	// (POST /projects)
	PostProjects(ctx echo.Context) error
	// Project Show
	// (GET /projects/{id})
	GetProjectsId(ctx echo.Context, id int) error
	// Project Update
	// (PUT /projects/{id})
	PutProjectsId(ctx echo.Context, id int) error
	// Supporter SignIn
	// (POST /supporters/signIn)
	PostSupportersSignIn(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostCompaniesSignIn converts echo context to params.
func (w *ServerInterfaceWrapper) PostCompaniesSignIn(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostCompaniesSignIn(ctx)
	return err
}

// GetCsrf converts echo context to params.
func (w *ServerInterfaceWrapper) GetCsrf(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetCsrf(ctx)
	return err
}

// PostPlans converts echo context to params.
func (w *ServerInterfaceWrapper) PostPlans(ctx echo.Context) error {
	var err error

	ctx.Set(BusinessAuthenticationScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostPlans(ctx)
	return err
}

// GetProjects converts echo context to params.
func (w *ServerInterfaceWrapper) GetProjects(ctx echo.Context) error {
	var err error

	ctx.Set(BusinessAuthenticationScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProjectsParams
	// ------------- Optional query parameter "pageToken" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageToken", ctx.QueryParams(), &params.PageToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageToken: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProjects(ctx, params)
	return err
}

// PostProjects converts echo context to params.
func (w *ServerInterfaceWrapper) PostProjects(ctx echo.Context) error {
	var err error

	ctx.Set(BusinessAuthenticationScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostProjects(ctx)
	return err
}

// GetProjectsId converts echo context to params.
func (w *ServerInterfaceWrapper) GetProjectsId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BusinessAuthenticationScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProjectsId(ctx, id)
	return err
}

// PutProjectsId converts echo context to params.
func (w *ServerInterfaceWrapper) PutProjectsId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BusinessAuthenticationScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutProjectsId(ctx, id)
	return err
}

// PostSupportersSignIn converts echo context to params.
func (w *ServerInterfaceWrapper) PostSupportersSignIn(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostSupportersSignIn(ctx)
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

	router.POST(baseURL+"/companies/signIn", wrapper.PostCompaniesSignIn)
	router.GET(baseURL+"/csrf", wrapper.GetCsrf)
	router.POST(baseURL+"/plans", wrapper.PostPlans)
	router.GET(baseURL+"/projects", wrapper.GetProjects)
	router.POST(baseURL+"/projects", wrapper.PostProjects)
	router.GET(baseURL+"/projects/:id", wrapper.GetProjectsId)
	router.PUT(baseURL+"/projects/:id", wrapper.PutProjectsId)
	router.POST(baseURL+"/supporters/signIn", wrapper.PostSupportersSignIn)

}

type CompanySignInBadRequestResponseJSONResponse struct {
	Errors []string `json:"errors"`
}

type CompanySignInOkResponseResponseHeaders struct {
	SetCookie string
}
type CompanySignInOkResponseJSONResponse struct {
	Body map[string]interface{}

	Headers CompanySignInOkResponseResponseHeaders
}

type CsrfResponseJSONResponse struct {
	CsrfToken string `json:"csrfToken"`
}

type InternalServerErrorResponseJSONResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type NotFoundErrorResponseJSONResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PlanStoreResponseJSONResponse struct {
	Errors PlanValidationError `json:"errors"`
	Plan   Plan                `json:"plan"`
}

type ProjectResponseJSONResponse struct {
	// Project Project
	Project Project `json:"project"`
}

type ProjectStoreResponseJSONResponse struct {
	Errors ProjectValidationError `json:"errors"`

	// Project Project
	Project Project `json:"project"`
}

type ProjectsListResponseJSONResponse struct {
	NextPageToken string    `json:"nextPageToken"`
	Projects      []Project `json:"projects"`
}

type SupporterSignInBadRequestResponseJSONResponse struct {
	Errors []string `json:"errors"`
}

type SupporterSignInOkResponseResponseHeaders struct {
	SetCookie string
}
type SupporterSignInOkResponseJSONResponse struct {
	Body map[string]interface{}

	Headers SupporterSignInOkResponseResponseHeaders
}

type PostCompaniesSignInRequestObject struct {
	Body *PostCompaniesSignInJSONRequestBody
}

type PostCompaniesSignInResponseObject interface {
	VisitPostCompaniesSignInResponse(w http.ResponseWriter) error
}

type PostCompaniesSignIn200JSONResponse struct {
	CompanySignInOkResponseJSONResponse
}

func (response PostCompaniesSignIn200JSONResponse) VisitPostCompaniesSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Set-Cookie", fmt.Sprint(response.Headers.SetCookie))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostCompaniesSignIn400JSONResponse struct {
	CompanySignInBadRequestResponseJSONResponse
}

func (response PostCompaniesSignIn400JSONResponse) VisitPostCompaniesSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostCompaniesSignIn500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response PostCompaniesSignIn500JSONResponse) VisitPostCompaniesSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetCsrfRequestObject struct {
}

type GetCsrfResponseObject interface {
	VisitGetCsrfResponse(w http.ResponseWriter) error
}

type GetCsrf200JSONResponse struct{ CsrfResponseJSONResponse }

func (response GetCsrf200JSONResponse) VisitGetCsrfResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetCsrf500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response GetCsrf500JSONResponse) VisitGetCsrfResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostPlansRequestObject struct {
	Body *PostPlansJSONRequestBody
}

type PostPlansResponseObject interface {
	VisitPostPlansResponse(w http.ResponseWriter) error
}

type PostPlans200JSONResponse struct{ PlanStoreResponseJSONResponse }

func (response PostPlans200JSONResponse) VisitPostPlansResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostPlans500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response PostPlans500JSONResponse) VisitPostPlansResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetProjectsRequestObject struct {
	Params GetProjectsParams
}

type GetProjectsResponseObject interface {
	VisitGetProjectsResponse(w http.ResponseWriter) error
}

type GetProjects200JSONResponse struct {
	ProjectsListResponseJSONResponse
}

func (response GetProjects200JSONResponse) VisitGetProjectsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetProjects500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response GetProjects500JSONResponse) VisitGetProjectsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostProjectsRequestObject struct {
	Body *PostProjectsJSONRequestBody
}

type PostProjectsResponseObject interface {
	VisitPostProjectsResponse(w http.ResponseWriter) error
}

type PostProjects200JSONResponse struct {
	ProjectStoreResponseJSONResponse
}

func (response PostProjects200JSONResponse) VisitPostProjectsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostProjects500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response PostProjects500JSONResponse) VisitPostProjectsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetProjectsIdRequestObject struct {
	Id int `json:"id"`
}

type GetProjectsIdResponseObject interface {
	VisitGetProjectsIdResponse(w http.ResponseWriter) error
}

type GetProjectsId200JSONResponse struct{ ProjectResponseJSONResponse }

func (response GetProjectsId200JSONResponse) VisitGetProjectsIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetProjectsId404JSONResponse struct {
	NotFoundErrorResponseJSONResponse
}

func (response GetProjectsId404JSONResponse) VisitGetProjectsIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type PutProjectsIdRequestObject struct {
	Id   int `json:"id"`
	Body *PutProjectsIdJSONRequestBody
}

type PutProjectsIdResponseObject interface {
	VisitPutProjectsIdResponse(w http.ResponseWriter) error
}

type PutProjectsId200JSONResponse struct {
	ProjectStoreResponseJSONResponse
}

func (response PutProjectsId200JSONResponse) VisitPutProjectsIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PutProjectsId500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response PutProjectsId500JSONResponse) VisitPutProjectsIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostSupportersSignInRequestObject struct {
	Body *PostSupportersSignInJSONRequestBody
}

type PostSupportersSignInResponseObject interface {
	VisitPostSupportersSignInResponse(w http.ResponseWriter) error
}

type PostSupportersSignIn200JSONResponse struct {
	SupporterSignInOkResponseJSONResponse
}

func (response PostSupportersSignIn200JSONResponse) VisitPostSupportersSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Set-Cookie", fmt.Sprint(response.Headers.SetCookie))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostSupportersSignIn400JSONResponse struct {
	SupporterSignInBadRequestResponseJSONResponse
}

func (response PostSupportersSignIn400JSONResponse) VisitPostSupportersSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostSupportersSignIn500JSONResponse struct {
	InternalServerErrorResponseJSONResponse
}

func (response PostSupportersSignIn500JSONResponse) VisitPostSupportersSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// SignIn
	// (POST /companies/signIn)
	PostCompaniesSignIn(ctx context.Context, request PostCompaniesSignInRequestObject) (PostCompaniesSignInResponseObject, error)
	// Get Csrf
	// (GET /csrf)
	GetCsrf(ctx context.Context, request GetCsrfRequestObject) (GetCsrfResponseObject, error)
	// Plan Create
	// (POST /plans)
	PostPlans(ctx context.Context, request PostPlansRequestObject) (PostPlansResponseObject, error)
	// Project List
	// (GET /projects)
	GetProjects(ctx context.Context, request GetProjectsRequestObject) (GetProjectsResponseObject, error)
	// Project Create
	// (POST /projects)
	PostProjects(ctx context.Context, request PostProjectsRequestObject) (PostProjectsResponseObject, error)
	// Project Show
	// (GET /projects/{id})
	GetProjectsId(ctx context.Context, request GetProjectsIdRequestObject) (GetProjectsIdResponseObject, error)
	// Project Update
	// (PUT /projects/{id})
	PutProjectsId(ctx context.Context, request PutProjectsIdRequestObject) (PutProjectsIdResponseObject, error)
	// Supporter SignIn
	// (POST /supporters/signIn)
	PostSupportersSignIn(ctx context.Context, request PostSupportersSignInRequestObject) (PostSupportersSignInResponseObject, error)
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

// PostCompaniesSignIn operation middleware
func (sh *strictHandler) PostCompaniesSignIn(ctx echo.Context) error {
	var request PostCompaniesSignInRequestObject

	var body PostCompaniesSignInJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostCompaniesSignIn(ctx.Request().Context(), request.(PostCompaniesSignInRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostCompaniesSignIn")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostCompaniesSignInResponseObject); ok {
		return validResponse.VisitPostCompaniesSignInResponse(ctx.Response())
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

// PostPlans operation middleware
func (sh *strictHandler) PostPlans(ctx echo.Context) error {
	var request PostPlansRequestObject

	var body PostPlansJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostPlans(ctx.Request().Context(), request.(PostPlansRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostPlans")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostPlansResponseObject); ok {
		return validResponse.VisitPostPlansResponse(ctx.Response())
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

// PostProjects operation middleware
func (sh *strictHandler) PostProjects(ctx echo.Context) error {
	var request PostProjectsRequestObject

	var body PostProjectsJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProjects(ctx.Request().Context(), request.(PostProjectsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProjects")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProjectsResponseObject); ok {
		return validResponse.VisitPostProjectsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetProjectsId operation middleware
func (sh *strictHandler) GetProjectsId(ctx echo.Context, id int) error {
	var request GetProjectsIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProjectsId(ctx.Request().Context(), request.(GetProjectsIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProjectsId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetProjectsIdResponseObject); ok {
		return validResponse.VisitGetProjectsIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PutProjectsId operation middleware
func (sh *strictHandler) PutProjectsId(ctx echo.Context, id int) error {
	var request PutProjectsIdRequestObject

	request.Id = id

	var body PutProjectsIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PutProjectsId(ctx.Request().Context(), request.(PutProjectsIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutProjectsId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PutProjectsIdResponseObject); ok {
		return validResponse.VisitPutProjectsIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostSupportersSignIn operation middleware
func (sh *strictHandler) PostSupportersSignIn(ctx echo.Context) error {
	var request PostSupportersSignInRequestObject

	var body PostSupportersSignInJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostSupportersSignIn(ctx.Request().Context(), request.(PostSupportersSignInRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostSupportersSignIn")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostSupportersSignInResponseObject); ok {
		return validResponse.VisitPostSupportersSignInResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xZ32/bNhD+VwRuj2qdbh1Q6C31tsLY1hn1tpfAKBjpbLOVSJU8pTUC/e8DSVmiJMqW",
	"fyRLgL05Eu9033dH3nfMPYlFlgsOHBWJ7omELwUofCsSBubBVGQ55dsFW/MZn/G8QP00FhyBm580z1MW",
	"U2SCTz4pwfUzFW8go/pXLkUOEitnkFGW6h+4zYFERKFkfE3KkORUqa9CJp6XZWiiYhISEt1UPhyLZbiz",
	"ELefIEZSapMEVCxZrsMi0Q5FYGEEgQVShmSeUr5AIeFcaK3veQACT36mCPrdSsiMIolIoh+EHjKk0EBm",
	"LhuMI6xB6tcKqcTRzpBhCt6QCs5wLlkMvs90WG9C2nlsU+x6G5MQzXtgiHdyYb/x5NLB1HWM7M6l6VaI",
	"FCjXbzP67W2RrAH9ycoY3/f6MrksxzBu2XVJV9rnoshzIRHkM9/iNY7uJjfOVS648pxob2nywR55H6o1",
	"52CXUkjziyFkystC9YBKSbd94NbBCSdagyOogZRhG+ufn0/CeGwo9UdCsgGagGVkAfhiKsRnBl7vdSVo",
	"91MlVxdIR6zk6i/xGfjhimuWjuJeyVUgHZpnHEFymi5A3oH8RWfxEvGLBAYOFVCKrmEELu2iWT8G3A5L",
	"YMEEBk2rqN4L/FUUPHneON8LDAwMD8JaFVz0UPhewopE5LtJI7km1kpN9Bf/oSlLjFMTkTkzU8rHmPb7",
	"tX4YHnOgOB25RYVtGxcgopIQB+FUywYUyDgsVa/zAHnktNqP+jJ7GTaOS3FLAnjIUb+zizRCDt9wTtcw",
	"dPrW8NvNchQRB1po7TjsRHEEQSrQPLQI6qikZysbeippQDh08D6UdOiFc6540JLaVox+P6+Oz06/kUAR",
	"kmvsae4XyDKv8L7oOJHs2xEz/9vHmvkqS8tcr55sZ+yeZoeGr7FV3aJxvFGLmvFmNUfjTVrkHbFXXVK7",
	"5Pk4blqD93giob+gP9LLVvTH5MySPjg4f7zdPznvfW/yPj7GfbPzLj9Vf/xDJJDuycwT2wAuzeOtWvcW",
	"R5i59xlPb5P2s9lJlurnVYcHcSEZbhe6d9gc3haKcVDqusANcKyanIlF78XYtqOQcJppZ2hERhNOzn6D",
	"rW15jK9EfzPv/AfX85kOShVZRuWWRIQ0GHaLSEjuQCpr+erllQYucuA0ZyQiP77Uj0KSU9yY2I2AopyB",
	"mijTV02JCmWSpgvVYNGthsyFwulutW3CJHSugbdDwqx1UzzxXBN3711+uLoa9lWtmwxdWJQheX20vUel",
	"lSH5aYyffRO9WzAkulm6uasZRLpWdjKtuCVLbTeJlVzpz1c7qJ2Md4BT/f4k5txLk4eG+Q4wqCKtgeo/",
	"LUY9fKr9JTc3S04otM6F/UlF1h/vH4Cv4RPkZlm2yDST99T0cIdPS2JFqDMtDRXOvBl8cippBmjE8011",
	"Xn0pQG6b4yqv56Jwj55ensSub5b8bwmu+rqOyGV4x9hSy+/hUm2IPb5ae//SKM+g9CnVbEVpv2wbUt3K",
	"ndyzpBxTvuY/TL4C1s2tqV+WEHcGRlmAp5CbyeacSm43odeH7fwXpGdzvdiIr4PlW/iqt3gUWv/fE+08",
	"/Z0ne/eE2l16jFNn9R3JGfLM+0++k2gfvhgaK9EOX6U9uEjr3Do5qWpyo5NlfGjfdsMUMiUR2SDm0WSS",
	"ipimG6EwenP15hXRpVA56Sp9rYwC4EkuGMdmrxnBVIbd1U0EHhsnvL5lLTV9H6tlaN/OKA2PjVUg5bL8",
	"NwAA//8IzlfYJSIAAA==",
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

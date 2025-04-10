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
	Projects []Project `json:"projects"`
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
	// Project List
	// (GET /projects)
	GetProjects(ctx echo.Context) error
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

// GetProjects converts echo context to params.
func (w *ServerInterfaceWrapper) GetProjects(ctx echo.Context) error {
	var err error

	ctx.Set(BusinessAuthenticationScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProjects(ctx)
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
	Projects []Project `json:"projects"`
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

type GetProjectsRequestObject struct {
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

// GetProjects operation middleware
func (sh *strictHandler) GetProjects(ctx echo.Context) error {
	var request GetProjectsRequestObject

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

	"H4sIAAAAAAAC/+xY227jNhN+FYH/f6ldZdstsNBd4rYLo+02WLe9CYyAEcc2NxapkqNsjUDvXgwl60jb",
	"8iFFA/TOFjmj+eb4jZ5ZotNMK1BoWfzMDPyZg8UbLSS4BxOdZlxtZnKppmqqshzpaaIVgnI/eZatZcJR",
	"ahV9sVrRM5usIOX0KzM6A4OVMki5XNMP3GTAYmbRSLVkRcgybu1XbYTnsAidVdKAYPFdpaMlMQ+3Evrh",
	"CyTIChIRYBMjMzKLxVsUQQkjCEogRchujSaZGWoD56LrvNKDEZT4niPQ2UKblCOLmaAH4fCutNcJyido",
	"KXrQeg1c0WnK/7rJxRKwdSwVwhKMO5Zq37FFbnC0JShxDf6wHHR75d3Aubd0uiWdszzLtEEwrzytahz9",
	"xHLKbaaV9VTRDRefyzL7XN05B7sx2rhfEiG1Xi9UD7gxfDMEXio4oYoaHEENpAi7WH99PAnjsabULwnZ",
	"CriA0iMzwDcTrR8leLXXmUDqJ9YsLhCOxJrFb/oR1OGMa66O8r01i8C03DxVCEbx9QzME5gfKIqXsF8L",
	"2NFUwFq+hBG4SEVzfwy4LZagBBM4NJ2k+qTxR50r8bpxftIYOBgehFWvvAC2rNREP/9vYMFi9r+omfJR",
	"KWaj6oUDXFvx+REN3gPE9fyLNrgRYP7gaymcXudgNwIu443wmEbZnXse59ifpb1gqLv9fxTMA1OhVnwE",
	"XBsQqg7c3qB/tZNvMOh3zL4e3peafgNzzp1/xArLDKHz26ZkvIFmYb+bGuAI4p7jgFK+QZl6eeUItnwv",
	"RtNl4dVxkEXfP+yn0XvPHY8eb+M+Il0e1X3jFy1gzQZ5WbePfps7tIqMLZDOjjJeqO3m8VKdJeYIsfZy",
	"M16ss/SMF6uDdkSP6UWzFyw7jCuZB0luJG5mVIVlDB9yKxVYe53jChRW7cLZQrWYlIUdMsVTUoaOSjbm",
	"ZPIn2JTNQ6qFHhbzVn9wfTslo2yeptxsWMxYg2F7iYXsCYwtJd+9vSLgOgPFM8li9u1bekTLE66c7W70",
	"cCXBRtZ1KJei2rqgUaI6LFNBXtIWJ9vbZTtjYes7xGbXSOt8qog83yn6S9g3V1e7dVX3ol3bSxGy90fL",
	"e+ZdEbLvxujZR+/bCcPiu3k7drUHkS9tSVMr37I5yUW0dtDrqwrqBuMjIC0a7CTPtTeol4b5ETCoLK2B",
	"0t8SY5sW7cK5JS0nYfXSuBfAvLsL3M2LjkO2o4Msajml4XFEhXcWYMcZR1be8BNacYZLu2vDv8KlE8du",
	"/E5tZ1v0LEUxJuWmwvVKw1NAR9Xuqp5O/bPp6FKwNmFFk0M4ZHE1ISHjT3Z7t8+9PyznX8jP9vVspb/u",
	"TN/cl735P+LW/2qiG6ffM7G3Jux2QxlHAOqF5gwG4P2ofJLbd29xY1nA4b33xXlAb0VshaqJDQXL6SDd",
	"ZcHkZs1itkLM4iha64SvV9pi/OHqwztGqVAp6ZNJGr4BKJFpqbCpNTeTi7B/u7HAI9MybyhZsxnfy2qm",
	"U8yLvwMAAP//gJRPadkaAAA=",
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

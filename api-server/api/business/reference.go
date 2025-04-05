// Package businessapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package businessapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
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

// ProjectStoreResponse defines model for ProjectStoreResponse.
type ProjectStoreResponse struct {
	Errors ProjectValidationError `json:"errors"`

	// Project Project
	Project Project `json:"project"`
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

// PostSupportersSignInJSONBody defines parameters for PostSupportersSignIn.
type PostSupportersSignInJSONBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// PostCompaniesSignInJSONRequestBody defines body for PostCompaniesSignIn for application/json ContentType.
type PostCompaniesSignInJSONRequestBody PostCompaniesSignInJSONBody

// PostProjectsJSONRequestBody defines body for PostProjects for application/json ContentType.
type PostProjectsJSONRequestBody PostProjectsJSONBody

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
	// Project Create
	// (POST /projects)
	PostProjects(ctx echo.Context) error
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

// PostProjects converts echo context to params.
func (w *ServerInterfaceWrapper) PostProjects(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostProjects(ctx)
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
	router.POST(baseURL+"/projects", wrapper.PostProjects)
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

type ProjectStoreResponseJSONResponse struct {
	Errors ProjectValidationError `json:"errors"`

	// Project Project
	Project Project `json:"project"`
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
	// Project Create
	// (POST /projects)
	PostProjects(ctx context.Context, request PostProjectsRequestObject) (PostProjectsResponseObject, error)
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

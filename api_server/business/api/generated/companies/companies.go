// Package companies provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package companies

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// InternalServerErrorResponse defines model for InternalServerErrorResponse.
type InternalServerErrorResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// SignInBadRequestResponse defines model for SignInBadRequestResponse.
type SignInBadRequestResponse struct {
	Errors []string `json:"errors"`
}

// SignInOkResponse defines model for SignInOkResponse.
type SignInOkResponse = map[string]interface{}

// SignInInput defines model for SignInInput.
type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// PostCompaniesSignInJSONBody defines parameters for PostCompaniesSignIn.
type PostCompaniesSignInJSONBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// PostCompaniesSignInJSONRequestBody defines body for PostCompaniesSignIn for application/json ContentType.
type PostCompaniesSignInJSONRequestBody PostCompaniesSignInJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// SignIn
	// (POST /companies/signIn)
	PostCompaniesSignIn(ctx echo.Context) error
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

}

type InternalServerErrorResponseJSONResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type SignInBadRequestResponseJSONResponse struct {
	Errors []string `json:"errors"`
}

type SignInOkResponseResponseHeaders struct {
	SetCookie string
}
type SignInOkResponseJSONResponse struct {
	Body map[string]interface{}

	Headers SignInOkResponseResponseHeaders
}

type PostCompaniesSignInRequestObject struct {
	Body *PostCompaniesSignInJSONRequestBody
}

type PostCompaniesSignInResponseObject interface {
	VisitPostCompaniesSignInResponse(w http.ResponseWriter) error
}

type PostCompaniesSignIn200JSONResponse struct{ SignInOkResponseJSONResponse }

func (response PostCompaniesSignIn200JSONResponse) VisitPostCompaniesSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Set-Cookie", fmt.Sprint(response.Headers.SetCookie))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostCompaniesSignIn400JSONResponse struct {
	SignInBadRequestResponseJSONResponse
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

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// SignIn
	// (POST /companies/signIn)
	PostCompaniesSignIn(ctx context.Context, request PostCompaniesSignInRequestObject) (PostCompaniesSignInResponseObject, error)
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

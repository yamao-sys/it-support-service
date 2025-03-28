package handlers

import (
	"business/api/generated/csrf"
	"context"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
)

type CsrfHandler interface {
	GetCsrf(ctx context.Context, request csrf.GetCsrfRequestObject) (csrf.GetCsrfResponseObject, error)
}

type csrfHandler struct {}

func NewCsrfHandler() CsrfHandler {
	return &csrfHandler{}
}

func (ch *csrfHandler) GetCsrf(ctx context.Context, request csrf.GetCsrfRequestObject) (csrf.GetCsrfResponseObject, error) {
	csrfToken, ok := ctx.Value(middleware.DefaultCSRFConfig.ContextKey).(string)
	if !ok {
		return csrf.GetCsrf500JSONResponse{InternalServerErrorResponseJSONResponse: csrf.InternalServerErrorResponseJSONResponse{
			Code: http.StatusInternalServerError,
			Message: "failed to retrieval token",
		}}, nil
	}
	
	return csrf.GetCsrf200JSONResponse{CsrfResponseJSONResponse: csrf.CsrfResponseJSONResponse{ CsrfToken: csrfToken }}, nil
}

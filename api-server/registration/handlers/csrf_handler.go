package registrationhandlers

import (
	registrationapi "apps/api/registration"
	"context"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
)

type CsrfHandler interface {
	GetCsrf(ctx context.Context, request registrationapi.GetCsrfRequestObject) (registrationapi.GetCsrfResponseObject, error)
}

type csrfHandler struct {}

func NewCsrfHandler() CsrfHandler {
	return &csrfHandler{}
}

func (ch *csrfHandler) GetCsrf(ctx context.Context, request registrationapi.GetCsrfRequestObject) (registrationapi.GetCsrfResponseObject, error) {
	csrfToken, ok := ctx.Value(middleware.DefaultCSRFConfig.ContextKey).(string)
	if !ok {
		return registrationapi.GetCsrf500JSONResponse{InternalServerErrorResponseJSONResponse: registrationapi.InternalServerErrorResponseJSONResponse{
			Code: http.StatusInternalServerError,
			Message: "failed to retrieval token",
		}}, nil
	}
	
	return registrationapi.GetCsrf200JSONResponse{CsrfResponseJSONResponse: registrationapi.CsrfResponseJSONResponse{ CsrfToken: csrfToken }}, nil
}

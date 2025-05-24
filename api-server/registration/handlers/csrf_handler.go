package registrationhandlers

import (
	registrationapi "apps/api/registration"
	"context"

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
		return registrationapi.GetCsrf500Response{}, nil
	}
	
	return registrationapi.GetCsrf200JSONResponse(registrationapi.CsrfResponse{CsrfToken: csrfToken}), nil
}

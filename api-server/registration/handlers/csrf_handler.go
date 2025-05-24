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
	csrfToken, _ := ctx.Value(middleware.DefaultCSRFConfig.ContextKey).(string)
	
	return registrationapi.GetCsrf200JSONResponse(registrationapi.CsrfResponse{CsrfToken: csrfToken}), nil
}

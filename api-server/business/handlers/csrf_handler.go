package businesshandlers

import (
	businessapi "apps/api/business"
	"context"

	"github.com/labstack/echo/v4/middleware"
)

type CsrfHandler interface {
	GetCsrf(ctx context.Context, request businessapi.GetCsrfRequestObject) (businessapi.GetCsrfResponseObject, error)
}

type csrfHandler struct {}

func NewCsrfHandler() CsrfHandler {
	return &csrfHandler{}
}

func (ch *csrfHandler) GetCsrf(ctx context.Context, request businessapi.GetCsrfRequestObject) (businessapi.GetCsrfResponseObject, error) {
	csrfToken, ok := ctx.Value(middleware.DefaultCSRFConfig.ContextKey).(string)
	if !ok {
		return businessapi.GetCsrf500Response{}, nil
	}
	
	return businessapi.GetCsrf200JSONResponse(businessapi.CsrfResponse{ CsrfToken: csrfToken }), nil
}

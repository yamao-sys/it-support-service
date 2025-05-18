package businesshandlers

import (
	businessapi "apps/api/business"
	businessservices "apps/business/services"
	"context"
	"net/http"
)

type SupportersHandler interface {
	PostSupportersSignIn(ctx context.Context, request businessapi.PostSupportersSignInRequestObject) (businessapi.PostSupportersSignInResponseObject, error)
}

type supportersHandler struct {
	supporterService businessservices.SupporterService
}

func NewSupportersHandler(supporterService businessservices.SupporterService) SupportersHandler {
	return &supportersHandler{supporterService}
}

func (sh *supportersHandler) PostSupportersSignIn(ctx context.Context, request businessapi.PostSupportersSignInRequestObject) (businessapi.PostSupportersSignInResponseObject, error) {
	inputs := businessapi.PostSupportersSignInJSONRequestBody{
		Email: request.Body.Email,
		Password: request.Body.Password,
	}

	statusCode, tokenString, err := sh.supporterService.SignIn(ctx, inputs)
	switch (statusCode) {
	case http.StatusInternalServerError:
		return businessapi.PostSupportersSignIn500JSONResponse{Code: http.StatusInternalServerError, Message: err.Error()}, nil
	case http.StatusBadRequest:
		return businessapi.PostSupportersSignIn400JSONResponse{Errors: []string{err.Error()}}, nil
	}
	
	// NOTE: Cookieにtokenをセット
	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		MaxAge:   3600 * 24,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	}
	return businessapi.PostSupportersSignIn200JSONResponse{
		Body: businessapi.SupporterSignInOkResponse{},
		Headers: businessapi.PostSupportersSignIn200ResponseHeaders{SetCookie: cookie.String()},
	}, nil
}

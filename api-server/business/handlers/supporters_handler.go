package businesshandlers

import (
	businessapi "apps/api/business"
	businessservices "apps/business/services"
	"context"
	"net/http"
)

type SupportersHandler interface {
	PostSupporterSignIn(ctx context.Context, request businessapi.PostSupporterSignInRequestObject) (businessapi.PostSupporterSignInResponseObject, error)
}

type supportersHandler struct {
	supporterService businessservices.SupporterService
}

func NewSupportersHandler(supporterService businessservices.SupporterService) SupportersHandler {
	return &supportersHandler{supporterService}
}

func (sh *supportersHandler) PostSupporterSignIn(ctx context.Context, request businessapi.PostSupporterSignInRequestObject) (businessapi.PostSupporterSignInResponseObject, error) {
	inputs := businessapi.PostSupporterSignInJSONRequestBody{
		Email: request.Body.Email,
		Password: request.Body.Password,
	}

	statusCode, tokenString, err := sh.supporterService.SignIn(inputs)
	switch (statusCode) {
	case http.StatusInternalServerError:
		return businessapi.PostSupporterSignIn500Response{}, nil
	case http.StatusBadRequest:
		return businessapi.PostSupporterSignIn400JSONResponse{Errors: []string{err.Error()}}, nil
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
	return businessapi.PostSupporterSignIn200JSONResponse{
		Body: businessapi.SupporterSignInOkResponse{
			Token: tokenString,
		},
		Headers: businessapi.PostSupporterSignIn200ResponseHeaders{SetCookie: cookie.String()},
	}, nil
}

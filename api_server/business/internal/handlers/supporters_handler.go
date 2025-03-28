package handlers

import (
	"business/api/generated/supporters"
	"business/internal/services"
	"context"
	"net/http"
)

type SupportersHandler interface {
	PostSupportersSignIn(ctx context.Context, request supporters.PostSupportersSignInRequestObject) (supporters.PostSupportersSignInResponseObject, error)
}

type supportersHandler struct {
	supporterService services.SupporterService
}

func NewSupportersHandler(supporterService services.SupporterService) SupportersHandler {
	return &supportersHandler{supporterService}
}

func (sh *supportersHandler) PostSupportersSignIn(ctx context.Context, request supporters.PostSupportersSignInRequestObject) (supporters.PostSupportersSignInResponseObject, error) {
	inputs := supporters.PostSupportersSignInJSONRequestBody{
		Email: request.Body.Email,
		Password: request.Body.Password,
	}

	statusCode, tokenString, err := sh.supporterService.SignIn(ctx, inputs)
	switch (statusCode) {
	case http.StatusInternalServerError:
		return supporters.PostSupportersSignIn500JSONResponse{InternalServerErrorResponseJSONResponse: supporters.InternalServerErrorResponseJSONResponse{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		}}, nil
	case http.StatusBadRequest:
		return supporters.PostSupportersSignIn400JSONResponse{SignInBadRequestResponseJSONResponse: supporters.SignInBadRequestResponseJSONResponse{
			Errors: []string{err.Error()},
		}}, nil
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
	return supporters.PostSupportersSignIn200JSONResponse{SignInOkResponseJSONResponse: supporters.SignInOkResponseJSONResponse{
		Headers: supporters.SignInOkResponseResponseHeaders{
			SetCookie: cookie.String(),
		},
	}}, nil
}

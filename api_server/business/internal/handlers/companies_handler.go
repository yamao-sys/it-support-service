package handlers

import (
	"business/api/generated/companies"
	"business/internal/services"
	"context"
	"net/http"
)

type CompaniesHandler interface {
	PostCompaniesSignIn(ctx context.Context, request companies.PostCompaniesSignInRequestObject) (companies.PostCompaniesSignInResponseObject, error)
}

type companiesHandler struct {
	companyService services.CompanyService
}

func NewCompaniesHandler(companyService services.CompanyService) CompaniesHandler {
	return &companiesHandler{companyService}
}

func (ch *companiesHandler) PostCompaniesSignIn(ctx context.Context, request companies.PostCompaniesSignInRequestObject) (companies.PostCompaniesSignInResponseObject, error) {
	inputs := companies.PostCompaniesSignInJSONRequestBody{
		Email: request.Body.Email,
		Password: request.Body.Password,
	}

	statusCode, tokenString, err := ch.companyService.SignIn(ctx, inputs)
	switch (statusCode) {
	case http.StatusInternalServerError:
		return companies.PostCompaniesSignIn500JSONResponse{InternalServerErrorResponseJSONResponse: companies.InternalServerErrorResponseJSONResponse{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		}}, nil
	case http.StatusBadRequest:
		return companies.PostCompaniesSignIn400JSONResponse{SignInBadRequestResponseJSONResponse: companies.SignInBadRequestResponseJSONResponse{
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
	return companies.PostCompaniesSignIn200JSONResponse{SignInOkResponseJSONResponse: companies.SignInOkResponseJSONResponse{
		Headers: companies.SignInOkResponseResponseHeaders{
			SetCookie: cookie.String(),
		},
	}}, nil
}

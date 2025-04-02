package businesshandlers

import (
	businessapi "apps/api/business"
	businessservices "apps/business/services"
	"context"
	"net/http"
)

type CompaniesHandler interface {
	PostCompaniesSignIn(ctx context.Context, request businessapi.PostCompaniesSignInRequestObject) (businessapi.PostCompaniesSignInResponseObject, error)
}

type companiesHandler struct {
	companyService businessservices.CompanyService
}

func NewCompaniesHandler(companyService businessservices.CompanyService) CompaniesHandler {
	return &companiesHandler{companyService}
}

func (ch *companiesHandler) PostCompaniesSignIn(ctx context.Context, request businessapi.PostCompaniesSignInRequestObject) (businessapi.PostCompaniesSignInResponseObject, error) {
	inputs := businessapi.PostCompaniesSignInJSONRequestBody{
		Email: request.Body.Email,
		Password: request.Body.Password,
	}

	statusCode, tokenString, err := ch.companyService.SignIn(ctx, inputs)
	switch (statusCode) {
	case http.StatusInternalServerError:
		return businessapi.PostCompaniesSignIn500JSONResponse{InternalServerErrorResponseJSONResponse: businessapi.InternalServerErrorResponseJSONResponse{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		}}, nil
	case http.StatusBadRequest:
		return businessapi.PostCompaniesSignIn400JSONResponse{CompanySignInBadRequestResponseJSONResponse: businessapi.CompanySignInBadRequestResponseJSONResponse{
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
	return businessapi.PostCompaniesSignIn200JSONResponse{CompanySignInOkResponseJSONResponse: businessapi.CompanySignInOkResponseJSONResponse{
		Headers: businessapi.CompanySignInOkResponseResponseHeaders{
			SetCookie: cookie.String(),
		},
	}}, nil
}

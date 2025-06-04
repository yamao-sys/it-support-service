package businesshandlers

import (
	businessapi "apps/api/business"
	businessservices "apps/business/services"
	"context"
	"net/http"
)

type CompaniesHandler interface {
	PostCompanySignIn(ctx context.Context, request businessapi.PostCompanySignInRequestObject) (businessapi.PostCompanySignInResponseObject, error)
}

type companiesHandler struct {
	companyService businessservices.CompanyService
}

func NewCompaniesHandler(companyService businessservices.CompanyService) CompaniesHandler {
	return &companiesHandler{companyService}
}

func (ch *companiesHandler) PostCompanySignIn(ctx context.Context, request businessapi.PostCompanySignInRequestObject) (businessapi.PostCompanySignInResponseObject, error) {
	inputs := businessapi.PostCompanySignInJSONRequestBody{
		Email: request.Body.Email,
		Password: request.Body.Password,
	}

	statusCode, tokenString, err := ch.companyService.SignIn(inputs)
	switch (statusCode) {
	case http.StatusInternalServerError:
		return businessapi.PostCompanySignIn500Response{}, nil
	case http.StatusBadRequest:
		return businessapi.PostCompanySignIn400JSONResponse{Errors: []string{err.Error()}}, nil
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
	return businessapi.PostCompanySignIn200JSONResponse{
		Body: businessapi.CompanySignInOkResponse{
			Token: tokenString,
		},
		Headers: businessapi.PostCompanySignIn200ResponseHeaders{SetCookie: cookie.String()},
	}, nil
}

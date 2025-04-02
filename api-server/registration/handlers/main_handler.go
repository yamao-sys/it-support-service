package registrationhandlers

import (
	registrationapi "apps/api/registration"
	"context"
)

type MainHandler interface {
	// handlers /csrf
	GetCsrf(ctx context.Context, request registrationapi.GetCsrfRequestObject) (registrationapi.GetCsrfResponseObject, error)

	// handlers /supporters
	PostSupporterValidateSignUp(ctx context.Context, request registrationapi.PostSupporterValidateSignUpRequestObject) (registrationapi.PostSupporterValidateSignUpResponseObject, error)
	PostSupporterSignUp(ctx context.Context, request registrationapi.PostSupporterSignUpRequestObject) (registrationapi.PostSupporterSignUpResponseObject, error)

	// handlers /companies
	PostCompanyValidateSignUp(ctx context.Context, request registrationapi.PostCompanyValidateSignUpRequestObject) (registrationapi.PostCompanyValidateSignUpResponseObject, error)
	PostCompanySignUp(ctx context.Context, request registrationapi.PostCompanySignUpRequestObject) (registrationapi.PostCompanySignUpResponseObject, error)
}

type mainHandler struct {
	csrfHandler CsrfHandler
	supportersHandler SupportersHandler
	companiesHandler CompaniesHandler
}

func NewMainHandler(
	csrfHandler CsrfHandler,
	supportersHandler SupportersHandler,
	companiesHandler CompaniesHandler,
) MainHandler {
	return &mainHandler{csrfHandler, supportersHandler, companiesHandler}
}

func (mh *mainHandler) GetCsrf(ctx context.Context, request registrationapi.GetCsrfRequestObject) (registrationapi.GetCsrfResponseObject, error) {
	res, err := mh.csrfHandler.GetCsrf(ctx, request)
	return res, err
}

func (mh *mainHandler) PostSupporterValidateSignUp(ctx context.Context, request registrationapi.PostSupporterValidateSignUpRequestObject) (registrationapi.PostSupporterValidateSignUpResponseObject, error) {
	res, err := mh.supportersHandler.PostSupporterValidateSignUp(ctx, request)
	return res, err
}

func (mh *mainHandler) PostSupporterSignUp(ctx context.Context, request registrationapi.PostSupporterSignUpRequestObject) (registrationapi.PostSupporterSignUpResponseObject, error) {
	res, err := mh.supportersHandler.PostSupporterSignUp(ctx, request)
	return res, err
}

func (mh *mainHandler) PostCompanyValidateSignUp(ctx context.Context, request registrationapi.PostCompanyValidateSignUpRequestObject) (registrationapi.PostCompanyValidateSignUpResponseObject, error) {
	res, err := mh.companiesHandler.PostCompanyValidateSignUp(ctx, request)
	return res, err
}

func (mh *mainHandler) PostCompanySignUp(ctx context.Context, request registrationapi.PostCompanySignUpRequestObject) (registrationapi.PostCompanySignUpResponseObject, error) {
	res, err := mh.companiesHandler.PostCompanySignUp(ctx, request)
	return res, err
}

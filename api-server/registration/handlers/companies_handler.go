package registrationhandlers

import (
	registrationapi "apps/api/registration"
	registrationservices "apps/registration/services"
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type CompaniesHandler interface {
	PostCompanyValidateSignUp(ctx context.Context, request registrationapi.PostCompanyValidateSignUpRequestObject) (registrationapi.PostCompanyValidateSignUpResponseObject, error)
	PostCompanySignUp(ctx context.Context, request registrationapi.PostCompanySignUpRequestObject) (registrationapi.PostCompanySignUpResponseObject, error)
}

type companiesHandler struct {
	companiesService registrationservices.CompanyService
}

func NewCompaniesHandler(companiesService registrationservices.CompanyService) CompaniesHandler {
	return &companiesHandler{companiesService}
}

func (ch *companiesHandler) PostCompanyValidateSignUp(ctx context.Context, request registrationapi.PostCompanyValidateSignUpRequestObject) (registrationapi.PostCompanyValidateSignUpResponseObject, error) {
	reader := request.Body
	// NOTE: バリデーションチェックを行う構造体
	inputStruct, mappingErr := ch.mappingInputStruct(reader)
	if mappingErr != nil {
		return registrationapi.PostCompanyValidateSignUp500JSONResponse{InternalServerErrorResponseJSONResponse: registrationapi.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError, Message: mappingErr.Error()}}, nil
	}

	err := ch.companiesService.ValidateSignUp(ctx, &inputStruct)
	validationError := ch.mappingValidationErrorStruct(err)

	res := &registrationapi.CompanySignUpResponse{
		Code: http.StatusOK,
		Errors: validationError,
	}
	return registrationapi.PostCompanyValidateSignUp200JSONResponse{CompanySignUpResponseJSONResponse: registrationapi.CompanySignUpResponseJSONResponse{Code: res.Code, Errors: res.Errors}}, nil
}

func (ch *companiesHandler) PostCompanySignUp(ctx context.Context, request registrationapi.PostCompanySignUpRequestObject) (registrationapi.PostCompanySignUpResponseObject, error) {
	reader := request.Body
	// NOTE: バリデーションチェックを行う構造体
	inputStruct, mappingErr := ch.mappingInputStruct(reader)
	if mappingErr != nil {
		return registrationapi.PostCompanySignUp500JSONResponse{InternalServerErrorResponseJSONResponse: registrationapi.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError, Message: mappingErr.Error()}}, nil
	}

	err := ch.companiesService.ValidateSignUp(ctx, &inputStruct)
	if err != nil {
		validationError := ch.mappingValidationErrorStruct(err)
	
		res := &registrationapi.CompanySignUpResponse{
			Code: http.StatusBadRequest,
			Errors: validationError,
		}
		return registrationapi.PostCompanySignUp400JSONResponse{Code: res.Code, Errors: res.Errors}, nil
	}

	signUpErr := ch.companiesService.SignUp(ctx, registrationapi.PostCompanySignUpMultipartRequestBody(inputStruct))
	if signUpErr != nil {
		return registrationapi.PostCompanySignUp500JSONResponse{InternalServerErrorResponseJSONResponse: registrationapi.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError, Message: signUpErr.Error()}}, nil
	}

	res := &registrationapi.CompanySignUpResponse{
		Code: http.StatusOK,
		Errors: registrationapi.CompanySignUpValidationError{},
	}
	return registrationapi.PostCompanySignUp200JSONResponse{CompanySignUpResponseJSONResponse: registrationapi.CompanySignUpResponseJSONResponse{Code: res.Code, Errors: res.Errors}}, nil
}

func (ch *companiesHandler) mappingInputStruct(reader *multipart.Reader) (registrationapi.PostCompanyValidateSignUpMultipartRequestBody, error) {
	var inputStruct registrationapi.PostCompanyValidateSignUpMultipartRequestBody

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			// NOTE: 全てのパートを読み終えた場合
			break
		}
		if err != nil {
			return inputStruct, fmt.Errorf("failed to read multipart part: %w", err)
		}

		// NOTE: 各パートのヘッダー情報を取得
		partName := part.FormName()
		filename := part.FileName()

		var buf bytes.Buffer
		if _, err := io.Copy(&buf, part); err != nil {
			fmt.Println(err)
			return inputStruct, fmt.Errorf("failed to copy content: %w", err)
		}

		switch partName {
		case "name":
			inputStruct.Name = buf.String()
		case "password":
			inputStruct.Password = buf.String()
		case "email":
			inputStruct.Email = buf.String()
		case "finalTaxReturn":
			var finalTaxReturn openapi_types.File
			finalTaxReturn.InitFromBytes(buf.Bytes(), filename)
			inputStruct.FinalTaxReturn = &finalTaxReturn
		}
	}

	return inputStruct, nil
}

func (ch *companiesHandler) mappingValidationErrorStruct(err error) registrationapi.CompanySignUpValidationError {
	var validationError registrationapi.CompanySignUpValidationError
	if err == nil {
		return validationError
	}

	if errors, ok := err.(validation.Errors); ok {
		// NOTE: レスポンス用の構造体にマッピング
		for field, err := range errors {
			messages := []string{err.Error()}
			switch field {
			case "name":
				validationError.Name = &messages
			case "email":
				validationError.Email = &messages
			case "password":
				validationError.Password = &messages
			case "finalTaxReturn":
				validationError.FinalTaxReturn = &messages
			}
		}
	}
	return validationError
}

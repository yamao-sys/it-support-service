package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"registration/api/generated/companies"
	"registration/internal/services"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type CompaniesHandler interface {
	PostAuthValidateSignUp(ctx context.Context, request companies.PostAuthValidateSignUpRequestObject) (companies.PostAuthValidateSignUpResponseObject, error)
	PostAuthSignUp(ctx context.Context, request companies.PostAuthSignUpRequestObject) (companies.PostAuthSignUpResponseObject, error)
}

type companiesHandler struct {
	companiesService services.CompanyService
}

func NewCompaniesHandler(companiesService services.CompanyService) CompaniesHandler {
	return &companiesHandler{companiesService}
}

func (ch *companiesHandler) PostAuthValidateSignUp(ctx context.Context, request companies.PostAuthValidateSignUpRequestObject) (companies.PostAuthValidateSignUpResponseObject, error) {
	reader := request.Body
	// NOTE: バリデーションチェックを行う構造体
	inputStruct, mappingErr := ch.mappingInputStruct(reader)
	if mappingErr != nil {
		return companies.PostAuthValidateSignUp500JSONResponse{InternalServerErrorResponseJSONResponse: companies.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError, Message: mappingErr.Error()}}, nil
	}

	err := ch.companiesService.ValidateSignUp(ctx, &inputStruct)
	validationError := ch.mappingValidationErrorStruct(err)

	res := &companies.SignUpResponse{
		Code: http.StatusOK,
		Errors: validationError,
	}
	return companies.PostAuthValidateSignUp200JSONResponse{SignUpResponseJSONResponse: companies.SignUpResponseJSONResponse{Code: res.Code, Errors: res.Errors}}, nil
}

func (ch *companiesHandler) PostAuthSignUp(ctx context.Context, request companies.PostAuthSignUpRequestObject) (companies.PostAuthSignUpResponseObject, error) {
	reader := request.Body
	// NOTE: バリデーションチェックを行う構造体
	inputStruct, mappingErr := ch.mappingInputStruct(reader)
	if mappingErr != nil {
		return companies.PostAuthSignUp500JSONResponse{InternalServerErrorResponseJSONResponse: companies.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError, Message: mappingErr.Error()}}, nil
	}

	err := ch.companiesService.ValidateSignUp(ctx, &inputStruct)
	if err != nil {
		validationError := ch.mappingValidationErrorStruct(err)
	
		res := &companies.SignUpResponse{
			Code: http.StatusBadRequest,
			Errors: validationError,
		}
		return companies.PostAuthSignUp400JSONResponse{Code: res.Code, Errors: res.Errors}, nil
	}

	signUpErr := ch.companiesService.SignUp(ctx, companies.PostAuthSignUpMultipartRequestBody(inputStruct))
	if signUpErr != nil {
		log.Fatalln(err)
	}

	res := &companies.SignUpResponse{
		Code: http.StatusOK,
		Errors: companies.SignUpValidationError{},
	}
	return companies.PostAuthSignUp200JSONResponse{SignUpResponseJSONResponse: companies.SignUpResponseJSONResponse{Code: res.Code, Errors: res.Errors}}, nil
}

func (ch *companiesHandler) mappingInputStruct(reader *multipart.Reader) (companies.PostAuthValidateSignUpMultipartRequestBody, error) {
	var inputStruct companies.PostAuthValidateSignUpMultipartRequestBody

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

func (ch *companiesHandler) mappingValidationErrorStruct(err error) companies.SignUpValidationError {
	var validationError companies.SignUpValidationError
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

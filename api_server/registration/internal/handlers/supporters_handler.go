package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"registration/api/generated/supporters"
	"registration/internal/services"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type SupportersHandler interface {
	PostAuthValidateSignUp(ctx context.Context, request supporters.PostAuthValidateSignUpRequestObject) (supporters.PostAuthValidateSignUpResponseObject, error)
	PostAuthSignUp(ctx context.Context, request supporters.PostAuthSignUpRequestObject) (supporters.PostAuthSignUpResponseObject, error)
}

type supportersHandler struct {
	supporterService services.SupporterService
}

func NewSupportersHandler(supporterService services.SupporterService) SupportersHandler {
	return &supportersHandler{supporterService}
}

func (sh *supportersHandler) PostAuthValidateSignUp(ctx context.Context, request supporters.PostAuthValidateSignUpRequestObject) (supporters.PostAuthValidateSignUpResponseObject, error) {
	reader := request.Body
	// NOTE: バリデーションチェックを行う構造体
	inputStruct, mappingErr := sh.mappingInputStruct(reader)
	if mappingErr != nil {
		return supporters.PostAuthValidateSignUp500JSONResponse{InternalServerErrorResponseJSONResponse: supporters.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError, Message: mappingErr.Error()}}, nil
	}

	err := sh.supporterService.ValidateSignUp(ctx, &inputStruct)
	validationError := sh.mappingValidationErrorStruct(err)

	res := &supporters.SignUpResponse{
		Code: http.StatusOK,
		Errors: validationError,
	}
	return supporters.PostAuthValidateSignUp200JSONResponse{SignUpResponseJSONResponse: supporters.SignUpResponseJSONResponse{Code: res.Code, Errors: res.Errors}}, nil
}

func (sh *supportersHandler) PostAuthSignUp(ctx context.Context, request supporters.PostAuthSignUpRequestObject) (supporters.PostAuthSignUpResponseObject, error) {
	reader := request.Body
	// NOTE: バリデーションチェックを行う構造体
	inputStruct, mappingErr := sh.mappingInputStruct(reader)
	if mappingErr != nil {
		return supporters.PostAuthSignUp500JSONResponse{InternalServerErrorResponseJSONResponse: supporters.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError, Message: mappingErr.Error()}}, nil
	}

	err := sh.supporterService.ValidateSignUp(ctx, &inputStruct)
	if err != nil {
		validationError := sh.mappingValidationErrorStruct(err)
	
		res := &supporters.SignUpResponse{
			Code: http.StatusBadRequest,
			Errors: validationError,
		}
		return supporters.PostAuthSignUp400JSONResponse{Code: res.Code, Errors: res.Errors}, nil
	}

	signUpErr := sh.supporterService.SignUp(ctx, supporters.PostAuthSignUpMultipartRequestBody(inputStruct))
	if signUpErr != nil {
		log.Fatalln(err)
	}

	res := &supporters.SignUpResponse{
		Code: http.StatusOK,
		Errors: supporters.SignUpValidationError{},
	}
	return supporters.PostAuthSignUp200JSONResponse{SignUpResponseJSONResponse: supporters.SignUpResponseJSONResponse{Code: res.Code, Errors: res.Errors}}, nil
}

func (sh *supportersHandler) mappingInputStruct(reader *multipart.Reader) (supporters.PostAuthValidateSignUpMultipartRequestBody, error) {
	var inputStruct supporters.PostAuthValidateSignUpMultipartRequestBody

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
		case "firstName":
			inputStruct.FirstName = buf.String()
		case "lastName":
			inputStruct.LastName = buf.String()
		case "password":
			inputStruct.Password = buf.String()
		case "email":
			inputStruct.Email = buf.String()
		case "birthday":
			birthday := buf.String()
			if birthday == "" {
				continue
			}
			
			parsedTime, parseErr := time.Parse("2006-01-02", birthday)
			if parseErr != nil {
				fmt.Println(parseErr)
				return inputStruct, fmt.Errorf("failed to parse birthday: %w", parseErr)
			}
			inputStruct.Birthday = &openapi_types.Date{Time: parsedTime}
		case "frontIdentification":
			var frontIdentification openapi_types.File
			frontIdentification.InitFromBytes(buf.Bytes(), filename)
			inputStruct.FrontIdentification = &frontIdentification
		case "backIdentification":
			var backIdentification openapi_types.File
			backIdentification.InitFromBytes(buf.Bytes(), filename)
			inputStruct.BackIdentification = &backIdentification
		}
	}

	return inputStruct, nil
}

func (sh *supportersHandler) mappingValidationErrorStruct(err error) supporters.SignUpValidationError {
	var validationError supporters.SignUpValidationError
	if err == nil {
		return validationError
	}

	if errors, ok := err.(validation.Errors); ok {
		// NOTE: レスポンス用の構造体にマッピング
		for field, err := range errors {
			messages := []string{err.Error()}
			switch field {
			case "firstName":
				validationError.FirstName = &messages
			case "lastName":
				validationError.LastName = &messages
			case "email":
				validationError.Email = &messages
			case "password":
				validationError.Password = &messages
			case "frontIdentification":
				validationError.FrontIdentification = &messages
			case "backIdentification":
				validationError.BackIdentification = &messages
			}
		}
	}
	return validationError
}

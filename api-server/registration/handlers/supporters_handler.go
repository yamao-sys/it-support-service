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
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type SupportersHandler interface {
	PostSupporterValidateSignUp(ctx context.Context, request registrationapi.PostSupporterValidateSignUpRequestObject) (registrationapi.PostSupporterValidateSignUpResponseObject, error)
	PostSupporterSignUp(ctx context.Context, request registrationapi.PostSupporterSignUpRequestObject) (registrationapi.PostSupporterSignUpResponseObject, error)
}

type supportersHandler struct {
	supporterService registrationservices.SupporterService
}

func NewSupportersHandler(supporterService registrationservices.SupporterService) SupportersHandler {
	return &supportersHandler{supporterService}
}

func (sh *supportersHandler) PostSupporterValidateSignUp(ctx context.Context, request registrationapi.PostSupporterValidateSignUpRequestObject) (registrationapi.PostSupporterValidateSignUpResponseObject, error) {
	reader := request.Body
	// NOTE: バリデーションチェックを行う構造体
	inputStruct, mappingErr := sh.mappingInputStruct(reader)
	if mappingErr != nil {
		return registrationapi.PostSupporterValidateSignUp500JSONResponse{Code: http.StatusInternalServerError, Message: mappingErr.Error()}, nil
	}

	err := sh.supporterService.ValidateSignUp(&inputStruct)
	validationError := sh.mappingValidationErrorStruct(err)
	return registrationapi.PostSupporterValidateSignUp200JSONResponse(registrationapi.SupporterSignUpResponse{Code: http.StatusOK, Errors: validationError}), nil
}

func (sh *supportersHandler) PostSupporterSignUp(ctx context.Context, request registrationapi.PostSupporterSignUpRequestObject) (registrationapi.PostSupporterSignUpResponseObject, error) {
	reader := request.Body
	// NOTE: バリデーションチェックを行う構造体
	inputStruct, mappingErr := sh.mappingInputStruct(reader)
	if mappingErr != nil {
		return registrationapi.PostSupporterSignUp500JSONResponse{Code: http.StatusInternalServerError, Message: mappingErr.Error()}, nil
	}

	err := sh.supporterService.ValidateSignUp(&inputStruct)
	if err != nil {
		validationError := sh.mappingValidationErrorStruct(err)
	
		res := &registrationapi.SupporterSignUpResponse{
			Code: http.StatusBadRequest,
			Errors: validationError,
		}
		return registrationapi.PostSupporterSignUp400JSONResponse{Code: res.Code, Errors: res.Errors}, nil
	}

	signUpErr := sh.supporterService.SignUp(ctx, registrationapi.PostSupporterSignUpMultipartRequestBody(inputStruct))
	if signUpErr != nil {
		return registrationapi.PostSupporterSignUp500JSONResponse{Code: http.StatusInternalServerError, Message: signUpErr.Error()}, nil
	}

	return registrationapi.PostSupporterSignUp200JSONResponse(registrationapi.SupporterSignUpResponse{Code: http.StatusOK, Errors: registrationapi.SupporterSignUpValidationError{}}), nil
}

func (sh *supportersHandler) mappingInputStruct(reader *multipart.Reader) (registrationapi.PostSupporterValidateSignUpMultipartRequestBody, error) {
	var inputStruct registrationapi.PostSupporterValidateSignUpMultipartRequestBody

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

func (sh *supportersHandler) mappingValidationErrorStruct(err error) registrationapi.SupporterSignUpValidationError {
	var validationError registrationapi.SupporterSignUpValidationError
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

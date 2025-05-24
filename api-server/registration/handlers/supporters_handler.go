package registrationhandlers

import (
	registrationapi "apps/api/registration"
	registrationservices "apps/registration/services"
	"bytes"
	"context"
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
	inputStruct := sh.mappingInputStruct(reader)

	err := sh.supporterService.ValidateSignUp(&inputStruct)
	validationError := sh.mappingValidationErrorStruct(err)
	return registrationapi.PostSupporterValidateSignUp200JSONResponse(registrationapi.SupporterSignUpResponse{Code: http.StatusOK, Errors: validationError}), nil
}

func (sh *supportersHandler) PostSupporterSignUp(ctx context.Context, request registrationapi.PostSupporterSignUpRequestObject) (registrationapi.PostSupporterSignUpResponseObject, error) {
	reader := request.Body
	// NOTE: バリデーションチェックを行う構造体
	inputStruct := sh.mappingInputStruct(reader)

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
		return registrationapi.PostSupporterSignUp500Response{}, nil
	}

	return registrationapi.PostSupporterSignUp200JSONResponse(registrationapi.SupporterSignUpResponse{Code: http.StatusOK, Errors: registrationapi.SupporterSignUpValidationError{}}), nil
}

func (sh *supportersHandler) mappingInputStruct(reader *multipart.Reader) (registrationapi.PostSupporterValidateSignUpMultipartRequestBody) {
	var inputStruct registrationapi.PostSupporterValidateSignUpMultipartRequestBody

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			// NOTE: 全てのパートを読み終えた場合
			break
		}

		// NOTE: 各パートのヘッダー情報を取得
		partName := part.FormName()
		filename := part.FileName()

		var buf bytes.Buffer
		io.Copy(&buf, part)

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
			
			parsedTime, _ := time.Parse("2006-01-02", birthday)
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

	return inputStruct
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

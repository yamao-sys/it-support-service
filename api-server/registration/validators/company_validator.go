package registrationvalidator

import (
	registrationapi "apps/api/registration"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/exp/slices"

	"github.com/gabriel-vasile/mimetype"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

var companyAllowedMIMEType = []string{"image/webp", "image/png", "image/jpeg"}

func ValidateSignUpCompany(input *registrationapi.PostCompanyValidateSignUpMultipartRequestBody) error {
	return validation.ValidateStruct(input,
		validation.Field(
			&input.Name,
			validation.Required.Error("企業名は必須入力です。"),
			validation.RuneLength(1, 20).Error("企業名は1 ~ 20文字での入力をお願いします。"),
		),
		validation.Field(
			&input.Email,
			validation.Required.Error("Emailは必須入力です。"),
			is.Email.Error("Emailの形式での入力をお願いします。"),
		),
		validation.Field(
			&input.Password,
			validation.Required.Error("パスワードは必須入力です。"),
			validation.Length(8, 24).Error("パスワードは8 ~ 24文字での入力をお願いします。"),
		),
		validation.Field(
			&input.FinalTaxReturn,
			validation.By(isValidCompanyFileMimeType("確定申告書")),
		),
	)
}

func isValidCompanyFileMimeType(field string) validation.RuleFunc {
	return func(value interface{}) error {
		fileInput, ok := value.(*openapi_types.File)
		if !ok {
			return fmt.Errorf("%sが正しい形式ではありません", field)
		}
		if fileInput == nil {
			return nil
		}
		
		reader, err := fileInput.Reader()
		if err != nil {
			return fmt.Errorf("%sのリーダーを開く際にエラーが発生しました: %v", field, err)
		}

		mtype, err := mimetype.DetectReader(reader)
		if err != nil {
			return fmt.Errorf("%sのMIMEタイプを判別できません: %v", field, err)
		}

		if isValid := slices.Contains(companyAllowedMIMEType, mtype.String()); !isValid {
			return fmt.Errorf("%sの拡張子はwebp, png, jpegのいずれかでお願いします。", field)
		}
		return nil
    }
}

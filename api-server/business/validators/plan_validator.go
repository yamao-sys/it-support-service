package businessvalidators

import (
	businessapi "apps/api/business"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func ValidatePlan(input *businessapi.PlanStoreInput) error {
	return validation.ValidateStruct(input,
		validation.Field(
			&input.Title,
			validation.Required.Error("支援タイトルは必須入力です。"),
			validation.RuneLength(1, 70).Error("案件タイトルは1 ~ 70文字での入力をお願いします。"),
		),
		validation.Field(
			&input.Description,
			validation.Required.Error("支援概要は必須入力です。"),
		),
		validation.Field(
			&input.EndDate,
			validation.When(input.StartDate != nil, validation.By(validatePlanDateRequired( "支援終了日"))),
			validation.When(input.StartDate != nil && input.EndDate != nil, validation.By(validatePlanDateGreater("支援終了日", "支援開始日", input.StartDate))),
		),
		validation.Field(
			&input.UnitPrice,
			validation.Required.Error("支援単価(税抜)は必須入力です。"),
		),
	)
}

func validatePlanDateRequired(field string) validation.RuleFunc {
	return func(value interface{}) error {
		if value == nil || value.(*openapi_types.Date).Time.IsZero() {
			return fmt.Errorf("%sは必須入力です。", field)
		}
		return nil
	}
}

func validatePlanDateGreater(field string, doCompareField string, doCompareDate *openapi_types.Date) validation.RuleFunc {
	return func(value interface{}) error {
		dateValue, _ := value.(*openapi_types.Date)
		if doCompareDate.After(dateValue.Time) {
			return fmt.Errorf("%sと%sの前後関係が不適です。", field, doCompareField)
		}
		return nil
	}
}

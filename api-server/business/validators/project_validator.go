package businessvalidators

import (
	businessapi "apps/api/business"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func ValidateProject(input *businessapi.ProjectStoreInput) error {
	return validation.ValidateStruct(input,
		validation.Field(
			&input.Title,
			validation.Required.Error("案件タイトルは必須入力です。"),
			validation.RuneLength(1, 70).Error("案件タイトルは1 ~ 70文字での入力をお願いします。"),
		),
		validation.Field(
			&input.Description,
			validation.Required.Error("案件概要は必須入力です。"),
		),
		validation.Field(
			&input.StartDate,
			validation.Required.Error("案件開始日は必須入力です。"),
		),
		validation.Field(
			&input.EndDate,
			validation.Required.Error("案件終了日は必須入力です。"),
			validation.By(validateDateGreater("案件終了日", "案件開始日", input.StartDate)),
		),
		validation.Field(
			&input.MinBudget,
			validation.When(input.MinBudget != nil, validation.By(validateMinValue("予算(下限)"))),
		),
		validation.Field(
			&input.MaxBudget,
			validation.When(input.MaxBudget != nil, validation.By(validateMinValue("予算(上限)"))),
			validation.When(input.MaxBudget != nil && input.MinBudget != nil, validation.By(validateGreater("予算(上限)", "予算(下限)", input.MinBudget))),
		),
	)
}

func validateDateGreater(field string, doCompareField string, doCompareDate *openapi_types.Date) validation.RuleFunc {
	return func(value interface{}) error {
		dateValue, _ := value.(*openapi_types.Date)
		if doCompareDate.After(dateValue.Time) {
			return fmt.Errorf("%sと%sの前後関係が不適です。", field, doCompareField)
		}
		return nil
	}
}

// NOTE: 0だとなぜかvalidation.Min(1)でエラーになってくれないのでメソッドを作っている
func validateMinValue(field string) validation.RuleFunc {
	return func(value interface{}) error {
		intValue, _ := value.(*int)
		if *intValue < 1 {
			return fmt.Errorf("%sは正の整数での入力をお願いいたします。", field)
		}
		return nil
	}
}

func validateGreater(field string, doCompareField string, minBudget *int) validation.RuleFunc {
	return func(value interface{}) error {
		budgetValue, _ := value.(*int)
		if *minBudget > *budgetValue {
			return fmt.Errorf("%sと%sの大小関係が不適です。", field, doCompareField)
		}
		return nil
	}
}

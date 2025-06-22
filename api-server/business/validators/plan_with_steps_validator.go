package businessvalidators

import (
	businessapi "apps/api/business"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidatePlanWithSteps(input *businessapi.PlanStoreWithStepsInput) error {
	return validation.ValidateStruct(input,
		validation.Field(
			&input.Title,
			validation.Required.Error("支援タイトルは必須入力です。"),
			validation.RuneLength(1, 70).Error("支援タイトルは1 ~ 70文字での入力をお願いします。"),
		),
		validation.Field(
			&input.Description,
			validation.Required.Error("支援概要は必須入力です。"),
		),
		validation.Field(
			&input.EndDate,
			validation.When(input.StartDate != nil, validation.By(validatePlanDateRequired("支援終了日"))),
			validation.When(input.StartDate != nil && input.EndDate != nil, validation.By(validatePlanDateGreater("支援終了日", "支援開始日", input.StartDate))),
		),
		validation.Field(
			&input.UnitPrice,
			validation.Required.Error("支援単価(税抜)は必須入力です。"),
			validation.Min(1).Error("支援単価は1円以上で入力してください。"),
		),
		validation.Field(
			&input.PlanSteps,
			validation.When(input.PlanSteps != nil, validation.By(validatePlanSteps)),
		),
	)
}

func validatePlanSteps(value interface{}) error {
	if value == nil {
		return nil
	}

	steps, ok := value.(*[]businessapi.PlanStepInput)
	if !ok {
		return fmt.Errorf("support steps must be an array")
	}

	if steps == nil {
		return nil
	}

	if len(*steps) == 0 {
		return fmt.Errorf("少なくとも1つの支援ステップを追加してください。")
	}

	for _, step := range *steps {
		if err := validation.ValidateStruct(&step,
			validation.Field(
				&step.Title,
				validation.Required.Error("タイトルは必須入力です。"),
				validation.RuneLength(1, 50).Error("タイトルは1 ~ 50文字での入力をお願いします。"),
			),
			validation.Field(
				&step.Description,
				validation.Required.Error("概要は必須入力です。"),
			),
			validation.Field(
				&step.Duration,
				validation.Required.Error("支援時間は必須入力です。"),
				validation.Min(1).Error("支援時間は1時間以上で入力してください。"),
			),
		); err != nil {
			return err
		}
	}

	return nil
}

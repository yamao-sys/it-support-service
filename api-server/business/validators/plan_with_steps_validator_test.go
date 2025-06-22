package businessvalidators

import (
	businessapi "apps/api/business"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
)

func TestValidatePlanWithSteps_Success(t *testing.T) {
	startDate := time.Date(2025, 6, 1, 0, 0, 0, 0, time.Local)
	endDate := time.Date(2025, 6, 30, 0, 0, 0, 0, time.Local)

	input := &businessapi.PlanStoreWithStepsInput{
		Title:       "テスト提案",
		Description: "テスト提案の概要",
		StartDate:   &openapi_types.Date{Time: startDate},
		EndDate:     &openapi_types.Date{Time: endDate},
		UnitPrice:   5000,
		PlanSteps:   nil,
	}

	err := ValidatePlanWithSteps(input)
	assert.Nil(t, err)
}

func TestValidatePlanWithSteps_EmptyTitle(t *testing.T) {
	input := &businessapi.PlanStoreWithStepsInput{
		Title:       "",
		Description: "概要",
		UnitPrice:   5000,
	}

	err := ValidatePlanWithSteps(input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "支援タイトルは必須入力です")
}

func TestValidatePlanWithSteps_TitleTooLong(t *testing.T) {
	input := &businessapi.PlanStoreWithStepsInput{
		Title:       randomdata.RandStringRunes(71),
		Description: "概要",
		UnitPrice:   5000,
	}

	err := ValidatePlanWithSteps(input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "支援タイトルは1 ~ 70文字での入力をお願いします")
}

func TestValidatePlanWithSteps_EmptyDescription(t *testing.T) {
	input := &businessapi.PlanStoreWithStepsInput{
		Title:       "タイトル",
		Description: "",
		UnitPrice:   5000,
	}

	err := ValidatePlanWithSteps(input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "支援概要は必須入力です")
}

func TestValidatePlanWithSteps_InvalidUnitPrice(t *testing.T) {
	input := &businessapi.PlanStoreWithStepsInput{
		Title:       "タイトル",
		Description: "概要",
		UnitPrice:   0,
	}

	err := ValidatePlanWithSteps(input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "支援単価(税抜)は必須入力です。")
}

func TestValidatePlanWithSteps_EndDateRequired(t *testing.T) {
	startDate := time.Date(2025, 6, 1, 0, 0, 0, 0, time.Local)

	input := &businessapi.PlanStoreWithStepsInput{
		Title:       "タイトル",
		Description: "概要",
		StartDate:   &openapi_types.Date{Time: startDate},
		EndDate:     nil,
		UnitPrice:   5000,
	}

	err := ValidatePlanWithSteps(input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "支援終了日は必須入力です")
}

func TestValidatePlanWithSteps_DateRangeInvalid(t *testing.T) {
	startDate := time.Date(2025, 6, 2, 0, 0, 0, 0, time.Local)
	endDate := time.Date(2025, 6, 1, 0, 0, 0, 0, time.Local) // 開始日より前

	input := &businessapi.PlanStoreWithStepsInput{
		Title:       "タイトル",
		Description: "概要",
		StartDate:   &openapi_types.Date{Time: startDate},
		EndDate:     &openapi_types.Date{Time: endDate},
		UnitPrice:   5000,
	}

	err := ValidatePlanWithSteps(input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "前後関係が不適です")
}

func TestValidatePlanWithSteps_PlanStepsNil(t *testing.T) {
	input := &businessapi.PlanStoreWithStepsInput{
		Title:       "タイトル",
		Description: "概要",
		UnitPrice:   5000,
		PlanSteps:   nil,
	}

	err := ValidatePlanWithSteps(input)
	assert.Nil(t, err)
}

func TestValidatePlanWithSteps_PlanStepsEmpty(t *testing.T) {
	planSteps := []businessapi.PlanStepInput{}

	input := &businessapi.PlanStoreWithStepsInput{
		Title:       "タイトル",
		Description: "概要",
		UnitPrice:   5000,
		PlanSteps:   &planSteps,
	}

	err := ValidatePlanWithSteps(input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "少なくとも1つの支援ステップを追加してください")
}

func TestValidatePlanWithSteps_PlanStepEmptyTitle(t *testing.T) {
	planSteps := []businessapi.PlanStepInput{
		{
			Title:       "",
			Description: "概要",
			Duration:    10,
		},
	}

	input := &businessapi.PlanStoreWithStepsInput{
		Title:       "タイトル",
		Description: "概要",
		UnitPrice:   5000,
		PlanSteps:   &planSteps,
	}

	err := ValidatePlanWithSteps(input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "タイトルは必須入力です")
}

func TestValidatePlanWithSteps_PlanStepTitleTooLong(t *testing.T) {
	planSteps := []businessapi.PlanStepInput{
		{
			Title:       randomdata.RandStringRunes(51),
			Description: "概要",
			Duration:    10,
		},
	}

	input := &businessapi.PlanStoreWithStepsInput{
		Title:       "タイトル",
		Description: "概要",
		UnitPrice:   5000,
		PlanSteps:   &planSteps,
	}

	err := ValidatePlanWithSteps(input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "タイトルは1 ~ 50文字での入力をお願いします")
}

func TestValidatePlanWithSteps_PlanStepEmptyDescription(t *testing.T) {
	planSteps := []businessapi.PlanStepInput{
		{
			Title:       "ステップ1",
			Description: "",
			Duration:    10,
		},
	}

	input := &businessapi.PlanStoreWithStepsInput{
		Title:       "タイトル",
		Description: "概要",
		UnitPrice:   5000,
		PlanSteps:   &planSteps,
	}

	err := ValidatePlanWithSteps(input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "概要は必須入力です")
}

func TestValidatePlanWithSteps_PlanStepInvalidDuration(t *testing.T) {
	planSteps := []businessapi.PlanStepInput{
		{
			Title:       "ステップ1",
			Description: "概要",
			Duration:    0,
		},
	}

	input := &businessapi.PlanStoreWithStepsInput{
		Title:       "タイトル",
		Description: "概要",
		UnitPrice:   5000,
		PlanSteps:   &planSteps,
	}

	err := ValidatePlanWithSteps(input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "支援時間は必須入力です。")
}

func TestValidatePlanSteps_InvalidType(t *testing.T) {
	// NOTE: 内部関数validatePlanStepsのテスト
	err := validatePlanSteps("invalid_type")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "support steps must be an array")
}

func TestValidatePlanSteps_NilValue(t *testing.T) {
	err := validatePlanSteps(nil)
	assert.Nil(t, err)
}

func TestValidatePlanSteps_NilPointer(t *testing.T) {
	var steps *[]businessapi.PlanStepInput = nil
	err := validatePlanSteps(steps)
	assert.Nil(t, err)
}

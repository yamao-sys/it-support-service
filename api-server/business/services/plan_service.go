package businessservices

import (
	businessapi "apps/api/business"
	businessvalidators "apps/business/validators"
	models "apps/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

type PlanService interface {
	Create(requestParams *businessapi.PlanStoreInput, supporterID int) (plan models.Plan, validatorErrors error, error error)
	MappingValidationErrorStruct(err error) businessapi.PlanValidationError
}

type planService struct {
	db *gorm.DB
}

type StatusNum int
const (
	NotStarted StatusNum = iota
	Agreed
	Rejected
	InProgress
	RequestingReward
	Rewarded
)

func NewPlanService(db *gorm.DB) PlanService {
	return &planService{db}
}

func (ps *planService) Create(requestParams *businessapi.PlanStoreInput, supporterID int) (plan models.Plan, validatorErrors error, error error) {
	// NOTE: バリデーションチェック
	validatorErrors = businessvalidators.ValidatePlan(requestParams)
	if validatorErrors != nil {
		return models.Plan{}, validatorErrors, nil
	}

	plan = models.Plan{}
	plan.SupporterID = supporterID
	plan.ProjectID = requestParams.ProjectId
	plan.Title = requestParams.Title
	plan.Description = requestParams.Description
	plan.StartDate = requestParams.StartDate.Time
	plan.EndDate = requestParams.EndDate.Time
	plan.UnitPrice = null.Int{Int: requestParams.UnitPrice, Valid: true}

	ps.db.Create(&plan)
	return plan, nil, nil
}

func (ps *planService) MappingValidationErrorStruct(err error) businessapi.PlanValidationError {
	var validationError businessapi.PlanValidationError
	if err == nil {
		return validationError
	}

	if errors, ok := err.(validation.Errors); ok {
		// NOTE: レスポンス用の構造体にマッピング
		for field, err := range errors {
			messages := []string{err.Error()}
			switch field {
			case "title":
				validationError.Title = &messages
			case "description":
				validationError.Description = &messages
			case "startDate":
				validationError.StartDate = &messages
			case "endDate":
				validationError.EndDate = &messages
			case "unitPrice":
				validationError.UnitPrice = &messages
			}
		}
	}

	return validationError
}

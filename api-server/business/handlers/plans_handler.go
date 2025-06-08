package businesshandlers

import (
	businessapi "apps/api/business"
	businesshelpers "apps/business/helpers"
	businessservices "apps/business/services"
	"context"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type PlansHandler interface {
	PostPlan(ctx context.Context, request businessapi.PostPlanRequestObject) (businessapi.PostPlanResponseObject, error)
}

type plansHandler struct {
	planService businessservices.PlanService
}

func NewPlansHandler(planService businessservices.PlanService) PlansHandler {
	return &plansHandler{planService}
}

func (ph *plansHandler) PostPlan(ctx context.Context, request businessapi.PostPlanRequestObject) (businessapi.PostPlanResponseObject, error) {
	supporterID, _ := businesshelpers.ExtractSupporterID(ctx)

	inputs := businessapi.PlanStoreInput{
		ProjectId: request.Body.ProjectId,
		Title: request.Body.Title,
		Description: request.Body.Description,
		StartDate: request.Body.StartDate,
		EndDate: request.Body.EndDate,
		UnitPrice: request.Body.UnitPrice,
	}

	createdPlan, validationErrors, err := ph.planService.Create(&inputs, supporterID)
	if err != nil {
		return businessapi.PostPlan500Response{}, err
	}

	mappedValidationErrors := ph.planService.MappingValidationErrorStruct(validationErrors)
	resPlan := businessapi.Plan{
		Id: createdPlan.ID,
		ProjectId: createdPlan.ProjectID,
		Title: createdPlan.Title,
		Description: createdPlan.Description,
		UnitPrice: createdPlan.UnitPrice,
		CreatedAt: createdPlan.CreatedAt,
	}
	if createdPlan.StartDate.Valid {
		resPlan.StartDate = &openapi_types.Date{Time: createdPlan.StartDate.Time}
	}
	if createdPlan.EndDate.Valid {
		resPlan.EndDate = &openapi_types.Date{Time: createdPlan.EndDate.Time}
	}

	return businessapi.PostPlan200JSONResponse(businessapi.PlanStoreResponse{Errors: mappedValidationErrors, Plan: resPlan}), nil
}

package businesshandlers

import (
	businessapi "apps/api/business"
	businesshelpers "apps/business/helpers"
	businessservices "apps/business/services"
	"context"
	"net/http"

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
		return businessapi.PostPlan500JSONResponse{Code: http.StatusInternalServerError}, err
	}

	mappedValidationErrors := ph.planService.MappingValidationErrorStruct(validationErrors)
	startDate := openapi_types.Date{Time: createdPlan.StartDate}
	endDate := openapi_types.Date{Time: createdPlan.EndDate}
	resPlan := businessapi.Plan{
		Id: createdPlan.ID,
		ProjectId: createdPlan.ProjectID,
		Title: createdPlan.Title,
		Description: createdPlan.Description,
		StartDate: startDate,
		EndDate: endDate,
		UnitPrice: createdPlan.UnitPrice.Int,
		CreatedAt: createdPlan.CreatedAt,
	}

	return businessapi.PostPlan200JSONResponse(businessapi.PlanStoreResponse{Errors: mappedValidationErrors, Plan: resPlan}), nil
}

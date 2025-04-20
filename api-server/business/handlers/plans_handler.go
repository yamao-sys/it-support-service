package businesshandlers

import (
	businessapi "apps/api/business"
	businesshelpers "apps/business/helpers"
	businessservices "apps/business/services"
	"context"
	"net/http"
	"strconv"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type PlansHandler interface {
	PostPlans(ctx context.Context, request businessapi.PostPlansRequestObject) (businessapi.PostPlansResponseObject, error)
}

type plansHandler struct {
	planService businessservices.PlanService
}

func NewPlansHandler(planService businessservices.PlanService) PlansHandler {
	return &plansHandler{planService}
}

func (ph *plansHandler) PostPlans(ctx context.Context, request businessapi.PostPlansRequestObject) (businessapi.PostPlansResponseObject, error) {
	supporterID, _ := businesshelpers.ExtractSupporterID(ctx)

	inputs := businessapi.PlanStoreInput{
		ProjectId: request.Body.ProjectId,
		Title: request.Body.Title,
		Description: request.Body.Description,
		StartDate: request.Body.StartDate,
		EndDate: request.Body.EndDate,
		UnitPrice: request.Body.UnitPrice,
	}

	createdPlan, validationErrors, err := ph.planService.Create(ctx, &inputs, supporterID)
	if err != nil {
		res := businessapi.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError}
		return businessapi.PostPlans500JSONResponse{InternalServerErrorResponseJSONResponse: res}, err
	}

	mappedValidationErrors := ph.planService.MappingValidationErrorStruct(validationErrors)
	planID := strconv.Itoa(createdPlan.ID)
	projectID := strconv.Itoa(createdPlan.ProjectID)
	startDate := openapi_types.Date{Time: createdPlan.StartDate}
	endDate := openapi_types.Date{Time: createdPlan.EndDate}
	resPlan := businessapi.Plan{
		Id: &planID,
		ProjectId: &projectID,
		Title: &createdPlan.Title,
		Description: &createdPlan.Description,
		StartDate: &startDate,
		EndDate: &endDate,
		UnitPrice: &createdPlan.UnitPrice.Int,
		CreatedAt: &createdPlan.CreatedAt,
	}

	res := businessapi.PlanStoreResponseJSONResponse{Errors: mappedValidationErrors, Plan: resPlan}
	return businessapi.PostPlans200JSONResponse{PlanStoreResponseJSONResponse: res}, nil
}

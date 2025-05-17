import "../../models/plan.tsp";

@summary("Plan Store Response")
model PlanStoreResponse {
  plan: Plan;
  errors: PlanValidationError;
}

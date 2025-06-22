import "./plan-step.tsp";

@summary("Plan")
model Plan {
  id: integer;
  projectId: integer;
  title: string;
  description: string;
  startDate?: plainDate;
  endDate?: plainDate;
  unitPrice: integer;

  @format("date-time")
  createdAt: string;
  planSteps?: PlanStep[];
}

@summary("Plan Validation Error")
model PlanValidationError {
  title?: string[];
  description?: string[];
  startDate?: string[];
  endDate?: string[];
  unitPrice?: string[];
}

@summary("Plan Store Input")
model PlanStoreInput {
  projectId: integer;
  title: string;
  description: string;
  startDate?: plainDate;
  endDate?: plainDate;
  unitPrice: integer;
}

@summary("Plan Validation Error")
model PlanWithStepsValidationError {
  title?: string[];
  description?: string[];
  startDate?: string[];
  endDate?: string[];
  unitPrice?: string[];
  planSteps?: PlanStepValidationError[];
}

@summary("Plan Store with Steps Input")
model PlanStoreWithStepsInput {
  title: string;
  description: string;
  startDate?: plainDate;
  endDate?: plainDate;
  unitPrice: integer;
  planSteps?: PlanStepInput[];
}

@summary("Plan Store Response")
model PlanStoreResponse {
  plan: Plan;
  errors: PlanValidationError;
}

@summary("Plan Store Response")
model PlanWithStepsStoreResponse {
  plan: Plan;
  errors: PlanWithStepsValidationError;
}

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

@summary("Plan Store Response")
model PlanStoreResponse {
  plan: Plan;
  errors: PlanValidationError;
}

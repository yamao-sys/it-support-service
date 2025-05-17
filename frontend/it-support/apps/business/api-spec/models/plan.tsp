@summary("Plan")
model Plan {
  id: integer;
  projectId: integer;
  title: string;
  description: string;
  startDate: plainDate;
  endDate: plainDate;
  unitPrice: integer;
  createdAt: plainTime;
}

@summary("Plan Validation Error")
model PlanValidationError {
  title?: string[];
  description?: string[];
  startDate?: string[];
  endDate?: string[];
  unitPrice?: string[];
}

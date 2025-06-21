@summary("Plan Step")
model PlanStep {
  id: integer;
  planId: integer;
  title: string;
  description: string;
  duration: integer;
}

@summary("PlanStep Validation Error")
model PlanStepValidationError {
  title?: string[];
  description?: string[];
  duration?: string[];
}

@summary("Plan Step Input")
model PlanStepInput {
  id?: integer;
  title: string;
  description: string;
  duration: integer;
  deleted?: boolean;
}

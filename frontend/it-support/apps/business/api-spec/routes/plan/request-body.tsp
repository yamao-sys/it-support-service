@summary("Plan Store Input")
model PlanStoreInput {
  projectId: integer;
  title: string;
  description: string;
  startDate?: plainDate;
  endDate?: plainDate;
  unitPrice: integer;
}

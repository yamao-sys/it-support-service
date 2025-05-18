@summary("Project Store Input")
model ProjectStoreInput {
  title: string;
  description: string;
  startDate?: plainDate;
  endDate?: plainDate;
  minBudget?: integer;
  maxBudget?: integer;
  isActive?: boolean;
}

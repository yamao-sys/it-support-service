@summary("Project")
model Project {
  id: string;
  name: string;
  description: string;
  startDate: plainDate;
  endDate: plainDate;
  minBudget?: integer;
  maxBudget?: integer;
  isActive: boolean;
  createdAt: string;
}

@summary("Project Validation Error")
model ProjectValidationError {
  title?: string[];
  description?: string[];
  startDate?: string[];
  endDate?: string[];
  minBudget?: string[];
  maxBudget?: string[];
  isActive?: string[];
}

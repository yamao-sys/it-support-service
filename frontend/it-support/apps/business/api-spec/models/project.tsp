@summary("Project")
model Project {
  id: integer;
  title: string;
  description: string;
  startDate: plainDate;
  endDate: plainDate;
  minBudget?: integer;
  maxBudget?: integer;
  isActive: boolean;

  @format("date-time")
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

alias ProposalStatus = "NOT PROPOSED" | "TEMPORARY CREATING" | "PROPOSED";

@summary("Project for Supporters")
model ToProject {
  id: integer;
  title: string;
  description: string;
  startDate: plainDate;
  endDate: plainDate;
  minBudget?: integer;
  maxBudget?: integer;
  proposalStatus: ProposalStatus;
}

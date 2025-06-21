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

@summary("Projects List Response for Supporters")
model ToProjectsListResponse{
  projects: ToProject[];
  nextPageToken: string;
}

@summary("Project Response for Supporters")
model ToProjectResponse{
  project: ToProject;
}

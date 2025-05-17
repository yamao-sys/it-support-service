import "../../models/project.tsp";

@summary("Project Store Response")
model ProjectStoreResponse {
  project: Project;
  errors: ProjectValidationError;
}

@summary("Projects List Response")
model ProjectsListResponse{
  projects: Project[];
  nextPageToken: string;
}

@summary("Project Response")
model ProjectResponse {
  project: Project;
}

import "../../models/project.tsp";

@summary("Project Store Response")
model ProjectStoreResponse {
  project: Project;
  errors: ProjectValidationError;
}

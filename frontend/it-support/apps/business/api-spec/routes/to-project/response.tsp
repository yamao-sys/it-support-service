import "../../models/project.tsp";

@summary("Projects List Response for Supporters")
model ToProjectsListResponse{
  projects: ToProject[];
  nextPageToken: string;
}

@summary("Project Response for Supporters")
model ToProjectResponse{
  project: ToProject;
}

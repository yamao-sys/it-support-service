import { getProjects } from "@/apis/projects.api";
import { FC } from "react";
import ProjectLists from "../ProjectLists";

const ProjectListsContainer: FC = async () => {
  const { projects, nextPageToken } = await getProjects();

  return (
    <>
      <ProjectLists initialProjects={projects} initialNextPageToken={Number(nextPageToken)} />
    </>
  );
};

export default ProjectListsContainer;

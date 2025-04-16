import { getProjects } from "@/apis/projects.api";
import { FC } from "react";
import ProjectLists from "../ProjectLists";

const ProjectListsContainer: FC = async () => {
  const projects = await getProjects();

  return (
    <>
      <ProjectLists projects={projects} />
    </>
  );
};

export default ProjectListsContainer;

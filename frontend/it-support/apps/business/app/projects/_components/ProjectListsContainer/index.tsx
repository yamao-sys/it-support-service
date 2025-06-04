import { FC } from "react";
import ProjectLists from "../ProjectLists";
import { getProjects } from "@/services/project";

const ProjectListsContainer: FC = async () => {
  // const { projects, nextPageToken } = await getProjects();
  const res = await getProjects();
  if (res === undefined) {
    throw new Error("Failed to fetch projects");
  }

  return (
    <>
      <ProjectLists
        initialProjects={res.projects}
        initialNextPageToken={Number(res.nextPageToken)}
      />
    </>
  );
};

export default ProjectListsContainer;

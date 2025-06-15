import { FC } from "react";
import ToProjectLists from "../../ToProjectLists";
import { getToProjects } from "@/services/toProject";

const ToProjectListsContainer: FC = async () => {
  const res = await getToProjects();
  if (res === undefined) {
    throw new Error("Failed to fetch projects");
  }

  return (
    <>
      <ToProjectLists
        initialProjects={res.projects}
        initialNextPageToken={Number(res.nextPageToken)}
      />
    </>
  );
};

export default ToProjectListsContainer;

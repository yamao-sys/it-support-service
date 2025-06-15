import { FC } from "react";

import ProjectEditContainer from "../../ProjectEditContainer";
import { getProject } from "@/services/project";

type Props = {
  projectId: number;
};

const ProjectEditTemplate: FC<Props> = async ({ projectId }: Props) => {
  const project = await getProject(projectId);
  if (project === undefined) {
    throw new Error("Project not found");
  }

  return (
    <>
      <ProjectEditContainer project={project} />
    </>
  );
};

export default ProjectEditTemplate;

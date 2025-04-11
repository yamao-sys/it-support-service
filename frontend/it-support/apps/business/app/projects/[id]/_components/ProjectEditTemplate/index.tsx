import { FC } from "react";
import { getProject } from "@/apis/projects.api";
import ProjectEditContainer from "../ProjectEditContainer";

type Props = {
  projectId: number;
};

const ProjectEditTemplate: FC<Props> = async ({ projectId }: Props) => {
  const project = await getProject(projectId);

  return (
    <>
      <ProjectEditContainer project={project} />
    </>
  );
};

export default ProjectEditTemplate;

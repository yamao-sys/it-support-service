import { FC } from "react";
import ProjectStoreForm from "../ProjectStoreForm";
import { getProject } from "@/apis/projects.api";

type Props = {
  projectId: number;
};

const ProjectEditTemplate: FC<Props> = async ({ projectId }: Props) => {
  const project = await getProject(projectId);

  return (
    <>
      <ProjectStoreForm project={project} />
    </>
  );
};

export default ProjectEditTemplate;

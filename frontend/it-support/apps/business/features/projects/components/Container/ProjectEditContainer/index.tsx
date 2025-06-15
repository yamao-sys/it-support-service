"use client";

import { FC } from "react";
import { useRouter } from "next/navigation";
import { Project, ProjectStoreInput } from "@/apis";
import { putUpdateProject } from "@/services/project";
import { useProjectStore } from "../../../hooks/useProjectStore";
import ProjectStoreForm from "../../ProjectStoreForm";

type Props = {
  project: Project;
};

const ProjectEditContainer: FC<Props> = ({ project }: Props) => {
  const doUpdateProjectInput: ProjectStoreInput = {
    title: project.title,
    description: project.description,
    startDate: project.startDate,
    endDate: project.endDate,
    minBudget: project.minBudget,
    maxBudget: project.maxBudget,
    isActive: project.isActive,
  };
  const { control, handleSubmit, validationErrors, setValidationErrors } =
    useProjectStore(doUpdateProjectInput);

  const router = useRouter();

  const onSubmit = handleSubmit(async (data) => {
    const resValidationErrors = await putUpdateProject(Number(project.id), data);
    if (resValidationErrors !== undefined && Object.keys(resValidationErrors).length > 0) {
      setValidationErrors(resValidationErrors);
      return;
    }

    window.alert("案件を更新しました!");
    router.push("/");
  });

  return (
    <>
      <ProjectStoreForm control={control} onSubmit={onSubmit} validationErrors={validationErrors} />
    </>
  );
};

export default ProjectEditContainer;

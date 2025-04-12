"use client";

import { useProjectStore } from "@/app/projects/_hooks/useProjectStore";
import { Project, ProjectStoreInput } from "@/types";
import { FC } from "react";
import ProjectStoreForm from "../../../_components/ProjectStoreForm";
import { useRouter } from "next/navigation";
import { putUpdateProject } from "@/apis/projects.api";

type Props = {
  project: Project;
};

const ProjectEditContainer: FC<Props> = ({ project }: Props) => {
  const doUpdateProjectInput: ProjectStoreInput = {
    title: project.title,
    description: project.description,
    startDate: project.start_date,
    endDate: project.end_date,
    minBudget: project.min_budget,
    maxBudget: project.max_budget,
    isActive: project.isActive,
  };
  const { control, handleSubmit, validationErrors, setValidationErrors } =
    useProjectStore(doUpdateProjectInput);

  const router = useRouter();

  const onSubmit = handleSubmit(async (data) => {
    const errors = await putUpdateProject(Number(project.id), data);
    if (Object.keys(errors).length > 0) {
      setValidationErrors(errors);
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

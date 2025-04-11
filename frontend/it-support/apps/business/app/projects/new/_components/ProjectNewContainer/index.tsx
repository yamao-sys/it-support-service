"use client";

import { postProjectCreate } from "@/apis/projects.api";
import ProjectStoreForm from "@/app/projects/_components/ProjectStoreForm";
import { useProjectStore } from "@/app/projects/_hooks/useProjectStore";
import { ProjectStoreInput } from "@/types";
import { useRouter } from "next/navigation";
import { FC } from "react";

const ProjectNewContainer: FC = () => {
  const doCreateProjectInput: ProjectStoreInput = {};
  const { register, control, handleSubmit, validationErrors, setValidationErrors } =
    useProjectStore(doCreateProjectInput);

  const router = useRouter();

  const onSubmit = handleSubmit(async (data) => {
    const errors = await postProjectCreate(data);
    if (Object.keys(errors).length > 0) {
      setValidationErrors(errors);
      return;
    }

    window.alert("案件を追加しました!");
    router.push("/");
  });

  return (
    <>
      <ProjectStoreForm
        register={register}
        control={control}
        onSubmit={onSubmit}
        validationErrors={validationErrors}
      />
    </>
  );
};

export default ProjectNewContainer;

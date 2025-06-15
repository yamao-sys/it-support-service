"use client";

import { ProjectStoreInput } from "@/apis";
import ProjectStoreForm from "@/app/projects/_components/ProjectStoreForm";
import { useProjectStore } from "@/app/projects/_hooks/useProjectStore";
import { postProjectCreate } from "@/services/project";
import { useRouter } from "next/navigation";
import { FC } from "react";

const ProjectNewContainer: FC = () => {
  // NOTE: requiredに設定したfieldはControlで制御する上で初期値が要るので設定
  const doCreateProjectInput: ProjectStoreInput = {
    title: "",
    description: "",
    isActive: true,
  };
  const { control, handleSubmit, validationErrors, setValidationErrors } =
    useProjectStore(doCreateProjectInput);

  const router = useRouter();

  const onSubmit = handleSubmit(async (data) => {
    const resValidationErrors = await postProjectCreate({ projectStoreInput: data });
    if (resValidationErrors !== undefined && Object.keys(resValidationErrors).length > 0) {
      setValidationErrors(resValidationErrors);
      return;
    }

    window.alert("案件を追加しました!");
    router.push("/");
  });

  return (
    <>
      <ProjectStoreForm control={control} onSubmit={onSubmit} validationErrors={validationErrors} />
    </>
  );
};

export default ProjectNewContainer;

import { ProjectStoreInput, ProjectValidationError } from "@/types";
import { useState } from "react";
import { useForm } from "react-hook-form";

const INITIAL_VALIDATION_ERRORS = {
  title: [],
  description: [],
  startDate: [],
  endDate: [],
  minBudget: [],
  maxBudget: [],
  isActive: [],
};

export const useProjectStore = (doStoreProjectInput: ProjectStoreInput) => {
  const { control, handleSubmit } = useForm<ProjectStoreInput>({
    defaultValues: doStoreProjectInput,
  });

  const [validationErrors, setValidationErrors] =
    useState<ProjectValidationError>(INITIAL_VALIDATION_ERRORS);

  return {
    control,
    handleSubmit,
    validationErrors,
    setValidationErrors,
  };
};

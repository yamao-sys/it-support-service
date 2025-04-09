"use client";

import { putUpdateProject } from "@/apis/projects.api";
import { Project, ProjectStoreInput, ProjectValidationError } from "@/types";
import { FC, useState } from "react";
import { Controller, useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
import BaseFormInput from "@repo/ui/BaseFormInput";
import BaseButton from "@repo/ui/BaseButton";
import BaseFormTextarea from "@repo/ui/BaseFormTextarea";
import ReactDatePicker from "react-datepicker";

const INITIAL_VALIDATION_ERRORS = {
  title: [],
  description: [],
  startDate: [],
  endDate: [],
  minBudget: [],
  maxBudget: [],
  isActive: [],
};

type Props = {
  project: Project;
};

const ProjectStoreForm: FC<Props> = ({ project }: Props) => {
  const doUpdateProject: ProjectStoreInput = {
    title: project.title,
    description: project.description,
    startDate: project.start_date,
    endDate: project.end_date,
    minBudget: project.min_budget,
    maxBudget: project.max_budget,
    isActive: project.isActive,
  };
  const { register, control, handleSubmit } = useForm<ProjectStoreInput>({
    defaultValues: doUpdateProject,
  });

  const [validationErrors, setValidationErrors] =
    useState<ProjectValidationError>(INITIAL_VALIDATION_ERRORS);

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
      <form onSubmit={onSubmit}>
        <div className='mt-8'>
          <Controller
            control={control}
            name='title'
            render={({ field: { ...fieldProps } }) => (
              <BaseFormInput
                id='title'
                label='案件タイトル'
                type='text'
                {...fieldProps}
                validationErrorMessages={validationErrors.title ?? []}
              />
            )}
          />
        </div>

        <div className='mt-8'>
          <Controller
            control={control}
            name='description'
            render={({ field: { ...fieldProps } }) => (
              <BaseFormTextarea
                id='description'
                label='案件概要'
                {...fieldProps}
                validationErrorMessages={validationErrors.description ?? []}
              />
            )}
          />
        </div>

        <div className='mt-8 flex items-between'>
          <div className='mr-4 w-1/3'>
            <label
              htmlFor='start-date'
              className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
            >
              <span className='font-bold'>案件開始日</span>
            </label>
            <Controller
              control={control}
              name='startDate'
              render={({ field: { value, ...fieldProps } }) => (
                <ReactDatePicker
                  {...fieldProps}
                  dateFormat='yyyy-MM-dd'
                  className='w-full px-3 py-2.5 focus-visible:outline-none bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
                  selected={value}
                />
              )}
            />
            {!!validationErrors.startDate?.length && (
              <div className='w-full pt-5 text-left'>
                {validationErrors.startDate.map((message, i) => (
                  <p key={i} className='text-red-400'>
                    {message}
                  </p>
                ))}
              </div>
            )}
          </div>
          <div className='w-1/3'>
            <label
              htmlFor='end-date'
              className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
            >
              <span className='font-bold'>案件終了日</span>
            </label>
            <Controller
              control={control}
              name='endDate'
              render={({ field: { value, ...fieldProps } }) => (
                <ReactDatePicker
                  {...fieldProps}
                  dateFormat='yyyy-MM-dd'
                  className='w-full px-3 py-2.5 focus-visible:outline-none bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
                  selected={value}
                />
              )}
            />
            {!!validationErrors.endDate?.length && (
              <div className='w-full pt-5 text-left'>
                {validationErrors.endDate.map((message, i) => (
                  <p key={i} className='text-red-400'>
                    {message}
                  </p>
                ))}
              </div>
            )}
          </div>
        </div>

        <div className='mt-8 flex items-between'>
          <div className='mr-4 w-1/3'>
            <Controller
              control={control}
              name='minBudget'
              render={({ field: { ...fieldProps } }) => (
                <BaseFormInput
                  id='minBudget'
                  label='予算(下限)'
                  type='number'
                  {...fieldProps}
                  validationErrorMessages={validationErrors.minBudget ?? []}
                />
              )}
            />
          </div>
          <div className='w-1/3'>
            <Controller
              control={control}
              name='maxBudget'
              render={({ field: { ...fieldProps } }) => (
                <BaseFormInput
                  id='maxBudget'
                  label='予算(上限)'
                  type='number'
                  {...fieldProps}
                  validationErrorMessages={validationErrors.maxBudget ?? []}
                />
              )}
            />
          </div>
        </div>

        <div className='mt-8'>
          <label className='inline-flex items-center cursor-pointer'>
            <input
              type='checkbox'
              className='sr-only peer'
              required={false}
              {...register("isActive")}
            />
            <div className="relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600 dark:peer-checked:bg-blue-600"></div>
            <span className='ms-3 text-sm font-medium text-gray-900 dark:text-gray-300'>
              公開フラグ
            </span>
          </label>
          {!!validationErrors.isActive?.length && (
            <div className='w-full pt-5 text-left'>
              {validationErrors.isActive.map((message, i) => (
                <p key={i} className='text-red-400'>
                  {message}
                </p>
              ))}
            </div>
          )}
        </div>

        <div className='w-full flex justify-center'>
          <div className='mt-16'>
            <BaseButton
              borderColor='border-green-500'
              bgColor='bg-green-500'
              label='保存する'
              type='submit'
            />
          </div>
        </div>
      </form>
    </>
  );
};

export default ProjectStoreForm;

"use client";

import { postProjectCreate } from "@/apis/projects.api";
import { ProjectStoreInput, ProjectValidationError } from "@/types";
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

const ProjectStoreForm: FC = () => {
  // TODO: register, post関数をprops化
  const { register, control, handleSubmit } = useForm<ProjectStoreInput>();

  const [validationErrors, setValidationErrors] =
    useState<ProjectValidationError>(INITIAL_VALIDATION_ERRORS);

  const router = useRouter();

  const onSubmit = handleSubmit(async (data) => {
    const errors = await postProjectCreate(data);
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
          <BaseFormInput
            id='title'
            label='案件タイトル'
            type='text'
            {...register("title", { required: "案件タイトルは必須です" })}
            validationErrorMessages={validationErrors.title ?? []}
          />
        </div>

        <div className='mt-8'>
          <BaseFormTextarea
            id='description'
            label='案件概要'
            {...register("description", { required: "案件概要は必須です" })}
            validationErrorMessages={validationErrors.description ?? []}
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
            {validationErrors.startDate?.length && (
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
            {validationErrors.endDate?.length && (
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
            <BaseFormInput
              id='minBudget'
              label='予算(下限)'
              type='number'
              {...register("minBudget")}
              validationErrorMessages={validationErrors.minBudget ?? []}
            />
          </div>
          <div className='w-1/3'>
            <BaseFormInput
              id='maxBudget'
              label='予算(上限)'
              type='number'
              {...register("maxBudget")}
              validationErrorMessages={validationErrors.maxBudget ?? []}
            />
          </div>
        </div>

        <div className='mt-8'>
          <label
            htmlFor='end-date'
            className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
          >
            <span className='font-bold'>公開フラグ</span>
          </label>

          <div className='flex space-x-2'>
            {[
              { label: "公開", value: 1 },
              { label: "非公開", value: 0 },
            ].map((option) => (
              <label key={option.label} className='cursor-pointer'>
                <input
                  type='radio'
                  value={option.value}
                  {...register("isActive", { required: "公開フラグは必須です。" })}
                  className='peer hidden'
                />
                <div className='px-4 py-2 rounded-full bg-gray-200 peer-checked:bg-blue-500 peer-checked:text-white'>
                  {option.label}
                </div>
              </label>
            ))}
          </div>
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

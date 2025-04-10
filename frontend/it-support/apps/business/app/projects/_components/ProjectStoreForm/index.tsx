"use client";

import { ProjectStoreInput, ProjectValidationError } from "@/types";
import { FC } from "react";
import { Control, Controller, UseFormRegister } from "react-hook-form";
import BaseFormInput from "@repo/ui/BaseFormInput";
import BaseButton from "@repo/ui/BaseButton";
import BaseFormTextarea from "@repo/ui/BaseFormTextarea";
import ReactDatePicker from "react-datepicker";

type Props = {
  register: UseFormRegister<ProjectStoreInput>;
  control: Control<ProjectStoreInput>;
  onSubmit: (e?: React.BaseSyntheticEvent) => Promise<void>;
  validationErrors: ProjectValidationError;
};

const ProjectStoreForm: FC<Props> = ({ register, control, onSubmit, validationErrors }: Props) => {
  return (
    <>
      <form onSubmit={onSubmit}>
        <div className='mt-8'>
          <Controller
            control={control}
            name='title'
            render={({ field }) => (
              <BaseFormInput
                id='title'
                label='案件タイトル'
                type='text'
                {...field}
                value={field.value ?? ""}
                validationErrorMessages={validationErrors.title ?? []}
              />
            )}
          />
        </div>

        <div className='mt-8'>
          <Controller
            control={control}
            name='description'
            render={({ field }) => (
              <BaseFormTextarea
                id='description'
                label='案件概要'
                {...field}
                value={field.value ?? ""}
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
              render={({ field }) => (
                <BaseFormInput
                  id='minBudget'
                  label='予算(下限)'
                  type='number'
                  {...field}
                  value={field.value ?? ""}
                  validationErrorMessages={validationErrors.minBudget ?? []}
                />
              )}
            />
          </div>
          <div className='w-1/3'>
            <Controller
              control={control}
              name='maxBudget'
              render={({ field }) => (
                <BaseFormInput
                  id='maxBudget'
                  label='予算(上限)'
                  type='number'
                  {...field}
                  value={field.value ?? ""}
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

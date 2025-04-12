"use client";

import { ProjectStoreInput, ProjectValidationError } from "@/types";
import { FC } from "react";
import { Control, Controller, UseFormRegister } from "react-hook-form";
import BaseFormInput from "@repo/ui/BaseFormInput";
import BaseButton from "@repo/ui/BaseButton";
import BaseFormTextarea from "@repo/ui/BaseFormTextarea";
import BaseFormDatePicker from "@repo/ui/BaseFormDatePicker";
import BaseFormSlider from "@repo/ui/BaseFormSlider";

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
            <BaseFormDatePicker
              label='案件開始日'
              control={control}
              name='startDate'
              validationErrors={validationErrors.startDate ?? []}
            />
          </div>
          <div className='w-1/3'>
            <BaseFormDatePicker
              label='案件終了日'
              control={control}
              name='endDate'
              validationErrors={validationErrors.endDate ?? []}
            />
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
          <BaseFormSlider
            label='公開フラグ'
            control={control}
            name='isActive'
            validationErrors={validationErrors.isActive ?? []}
          />
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

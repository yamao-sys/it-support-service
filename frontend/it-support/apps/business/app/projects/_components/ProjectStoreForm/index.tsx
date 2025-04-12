"use client";

import { ProjectStoreInput, ProjectValidationError } from "@/types";
import { FC } from "react";
import { Control } from "react-hook-form";
import BaseButton from "@repo/ui/BaseButton";
import BaseFormDatePicker from "@repo/ui/BaseFormDatePicker";
import BaseFormSlider from "@repo/ui/BaseFormSlider";
import BaseControlFormInput from "@repo/ui/BaseControlFormInput";
import BaseControlFormTextarea from "@repo/ui/BaseControlFormTextarea";

type Props = {
  control: Control<ProjectStoreInput>;
  onSubmit: (e?: React.BaseSyntheticEvent) => Promise<void>;
  validationErrors: ProjectValidationError;
};

const ProjectStoreForm: FC<Props> = ({ control, onSubmit, validationErrors }: Props) => {
  return (
    <>
      <form onSubmit={onSubmit}>
        <div className='mt-8'>
          <BaseControlFormInput
            id='title'
            label='案件タイトル'
            control={control}
            name='title'
            validationErrors={validationErrors.title ?? []}
          />
        </div>

        <div className='mt-8'>
          <BaseControlFormTextarea
            id='description'
            label='案件概要'
            control={control}
            name='description'
            validationErrors={validationErrors.description ?? []}
          />
        </div>

        <div className='mt-8 flex items-between'>
          <div className='mr-4 w-1/3'>
            <BaseFormDatePicker
              id='start-date'
              label='案件開始日'
              control={control}
              name='startDate'
              validationErrors={validationErrors.startDate ?? []}
            />
          </div>
          <div className='w-1/3'>
            <BaseFormDatePicker
              id='end-date'
              label='案件終了日'
              control={control}
              name='endDate'
              validationErrors={validationErrors.endDate ?? []}
            />
          </div>
        </div>

        <div className='mt-8 flex items-between'>
          <div className='mr-4 w-1/3'>
            <BaseControlFormInput
              id='min-budget'
              type='number'
              label='予算(下限)'
              control={control}
              name='minBudget'
              validationErrors={validationErrors.minBudget ?? []}
            />
          </div>
          <div className='w-1/3'>
            <BaseControlFormInput
              id='max-budget'
              type='number'
              label='予算(上限)'
              control={control}
              name='maxBudget'
              validationErrors={validationErrors.maxBudget ?? []}
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

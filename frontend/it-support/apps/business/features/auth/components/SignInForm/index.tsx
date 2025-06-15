"use client";

import { SignInFormType } from "@/types";
import { FC } from "react";
import FormTypeSelector from "../FormTypeSelector";
import BaseFormInput from "@repo/ui/BaseFormInput";
import BaseButton from "@repo/ui/BaseButton";
import { UseFormRegister } from "react-hook-form";
import { CompanySignInInput, SupporterSignInInput } from "@/apis";

type Props = {
  formType: SignInFormType;
  switchFormType: (newFormType: SignInFormType) => void;
  formTypeText: string;
  register: UseFormRegister<CompanySignInInput | SupporterSignInInput>;
  onSubmit: (e?: React.BaseSyntheticEvent) => Promise<void>;
  validationError?: string;
};

const SignInForm: FC<Props> = ({
  formType,
  switchFormType,
  formTypeText,
  register,
  onSubmit,
  validationError,
}: Props) => {
  return (
    <>
      <FormTypeSelector formType={formType} switchFormType={switchFormType} />

      <h3 className='mt-16 w-full text-center text-2xl font-bold'>
        {formTypeText} ログインフォーム
      </h3>

      {validationError && (
        <div className='w-full pt-5 text-center'>
          <p className='text-red-400'>{validationError}</p>
        </div>
      )}

      <form onSubmit={onSubmit}>
        <div className='mt-8'>
          <BaseFormInput
            id='email'
            label='Email'
            type='email'
            {...register("email", { required: "Emailは必須です" })}
            validationErrorMessages={[]}
          />
        </div>

        <div className='mt-8'>
          <BaseFormInput
            id='password'
            label='パスワード'
            type='password'
            {...register("password", { required: "パスワードは必須です" })}
            validationErrorMessages={[]}
          />
        </div>

        <div className='w-full flex justify-center'>
          <div className='mt-16'>
            <BaseButton
              borderColor='border-green-500'
              bgColor='bg-green-500'
              label='ログイン'
              type='submit'
            />
          </div>
        </div>
      </form>
    </>
  );
};

export default SignInForm;

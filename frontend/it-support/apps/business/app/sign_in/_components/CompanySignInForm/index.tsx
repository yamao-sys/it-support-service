"use client";

import { postCompanySignIn } from "@/apis/companies.api";
import { SignInFormType, CompanySignInInput } from "@/types";
import { FC, useState } from "react";
import { useForm } from "react-hook-form";
import FormTypeSelector from "../FormTypeSelector";
import BaseFormInput from "@repo/ui/BaseFormInput";
import { useRouter } from "next/navigation";
import BaseButton from "@repo/ui/BaseButton";

type Props = {
  formType: SignInFormType;
  switchFormType: (newFormType: SignInFormType) => void;
};

const CompanySignInForm: FC<Props> = ({ formType, switchFormType }: Props) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<CompanySignInInput>();
  const [validationError, setValidationError] = useState("");

  const router = useRouter();

  const onSubmit = handleSubmit(async (data) => {
    const response = await postCompanySignIn(data);
    if (response !== "") {
      setValidationError(response);
      return;
    }

    window.alert("企業のログインに成功しました!");
    router.push("/");
  });

  return (
    <>
      <FormTypeSelector formType={formType} switchFormType={switchFormType} />

      <h3 className='mt-16 w-full text-center text-2xl font-bold'>企業 ログインフォーム</h3>

      {validationError && (
        <div className='w-full pt-5 text-left'>
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
            validationErrorMessages={errors.email?.message ? [errors.email.message] : []}
          />
        </div>

        <div className='mt-8'>
          <BaseFormInput
            id='password'
            label='パスワード'
            type='password'
            {...register("password", { required: "パスワードは必須です" })}
            validationErrorMessages={errors.password?.message ? [errors.password.message] : []}
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

export default CompanySignInForm;

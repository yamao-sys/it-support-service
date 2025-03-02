"use client";

import { FC } from "react";
import { useForm } from "react-hook-form";
import { postValidateSignUp } from "../../_actions/supporters";
import { components } from "@/apis/generated/supporters/apiSchema";

export type SupporterSignUpInput =
  components["requestBodies"]["SignUpInput"]["content"]["multipart/form-data"];

const SupporterRegistrationForm: FC = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SupporterSignUpInput>();
  const onSubmit = handleSubmit(async (data) => {
    if (data.frontIdentification instanceof FileList && data.frontIdentification) {
      data.frontIdentification = data.frontIdentification[0];
    }
    if (data.backIdentification instanceof FileList && data.backIdentification) {
      data.backIdentification = data.backIdentification[0];
    }

    const result = await postValidateSignUp(data);
    console.log(data);
    console.log(result);
  });

  return (
    <>
      <form onSubmit={onSubmit}>
        <input type='text' placeholder='姓' {...register("lastName", { required: true })} />
        {errors.lastName && <span>姓は必須項目です。</span>}

        <input type='text' placeholder='名' {...register("firstName", { required: true })} />
        {errors.firstName && <span>名は必須項目です。</span>}

        <input type='email' placeholder='Email' {...register("email", { required: true })} />
        {errors.email && <span>Emailは必須項目です。</span>}

        <input
          type='password'
          placeholder='パスワード'
          {...register("password", { required: true })}
        />
        {errors.password && <span>パスワードは必須項目です。</span>}

        <input type='file' accept='image/*' {...register("frontIdentification")} />

        <input type='file' accept='image/*' {...register("backIdentification")} />

        <button type='submit'>確認画面へ</button>
      </form>
    </>
  );
};

export default SupporterRegistrationForm;

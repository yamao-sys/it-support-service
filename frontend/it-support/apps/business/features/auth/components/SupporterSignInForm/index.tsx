"use client";

import { SignInFormType } from "@/types";
import { FC, useState } from "react";
import { useForm } from "react-hook-form";
import FormTypeSelector from "../FormTypeSelector";
import BaseFormInput from "@repo/ui/BaseFormInput";
import { useRouter } from "next/navigation";
import BaseButton from "@repo/ui/BaseButton";
import { SupporterSignInInput } from "@/apis";
import { postSupporterSignIn } from "@/services/supporter";
import SignInForm from "../SignInForm";

type Props = {
  formType: SignInFormType;
  switchFormType: (newFormType: SignInFormType) => void;
};

const SupporterSignInForm: FC<Props> = ({ formType, switchFormType }: Props) => {
  const { register, handleSubmit } = useForm<SupporterSignInInput>();
  const [validationError, setValidationError] = useState("");

  const router = useRouter();

  const onSubmit = handleSubmit(async (data) => {
    const resValidationError = await postSupporterSignIn({ supporterSignInInput: data });
    if (resValidationError !== "" && resValidationError !== undefined) {
      setValidationError(resValidationError);
      return;
    }

    window.alert("サポータのログインに成功しました!");
    router.push("/");
  });

  return (
    <>
      <SignInForm
        formType={formType}
        switchFormType={switchFormType}
        formTypeText='サポータ'
        register={register}
        onSubmit={onSubmit}
        validationError={validationError}
      />
    </>
  );
};

export default SupporterSignInForm;

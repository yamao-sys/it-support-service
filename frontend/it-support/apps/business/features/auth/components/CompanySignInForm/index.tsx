"use client";

import { SignInFormType } from "@/types";
import { FC, useState } from "react";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
import { CompanySignInInput } from "@/apis";
import { postCompanySignIn } from "@/services/company";
import SignInForm from "../SignInForm";

type Props = {
  formType: SignInFormType;
  switchFormType: (newFormType: SignInFormType) => void;
};

const CompanySignInForm: FC<Props> = ({ formType, switchFormType }: Props) => {
  const { register, handleSubmit } = useForm<CompanySignInInput>();
  const [validationError, setValidationError] = useState("");

  const router = useRouter();

  const onSubmit = handleSubmit(async (data) => {
    const resValidationError = await postCompanySignIn({ companySignInInput: data });
    if (resValidationError !== "" && resValidationError !== undefined) {
      setValidationError(resValidationError);
      return;
    }

    window.alert("企業のログインに成功しました!");
    router.push("/");
  });

  return (
    <>
      <SignInForm
        formType={formType}
        switchFormType={switchFormType}
        formTypeText='企業'
        register={register}
        onSubmit={onSubmit}
        validationError={validationError}
      />
    </>
  );
};

export default CompanySignInForm;

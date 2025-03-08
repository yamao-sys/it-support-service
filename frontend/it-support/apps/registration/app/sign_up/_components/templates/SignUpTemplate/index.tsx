"use client";

import { FC, useEffect } from "react";
import { setCsrfToken } from "@/apis/csrf.api";
import { SupporterSignUpProvider } from "@/app/sign_up/_contexts/useSupporterSignUpContext";
import SupporterSignUpForm from "../../SupporterSignUpForm";

const SignUpTemplate: FC = () => {
  useEffect(() => {
    async function init() {
      await setCsrfToken();
    }
    init();
  }, []);

  return (
    <>
      <SupporterSignUpProvider>
        <SupporterSignUpForm />
      </SupporterSignUpProvider>
    </>
  );
};

export default SignUpTemplate;

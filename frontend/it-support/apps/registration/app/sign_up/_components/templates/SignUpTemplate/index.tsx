"use client";

import { FC, useEffect } from "react";
import { setCsrfToken } from "@/apis/csrf.api";
import { SupporterSignUpProvider } from "@/app/sign_up/_contexts/useSupporterSignUpContext";
import SignUpForm from "../../SignUpForm";
import { CompanySignUpProvider } from "@/app/sign_up/_contexts/useCompanySignUpContext";

const SignUpTemplate: FC = () => {
  useEffect(() => {
    async function init() {
      await setCsrfToken();
    }
    init();
  }, []);

  return (
    <>
      <CompanySignUpProvider>
        <SupporterSignUpProvider>
          <SignUpForm />
        </SupporterSignUpProvider>
      </CompanySignUpProvider>
    </>
  );
};

export default SignUpTemplate;

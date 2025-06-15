"use client";

import { FC, useEffect } from "react";
import SignUpForm from "../../SignUpForm";
import { setCsrfToken } from "@/services/auth";
import { CompanySignUpProvider } from "@/features/auth/contexts/useCompanySignUpContext";
import { SupporterSignUpProvider } from "@/features/auth/contexts/useSupporterSignUpContext";

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

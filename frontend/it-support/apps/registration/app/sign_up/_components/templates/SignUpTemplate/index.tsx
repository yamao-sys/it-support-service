"use client";

import { FC, useEffect } from "react";
import SupporterRegistrationForm from "../../SupporterRegistrationForm";
import { setCsrfToken } from "@/apis/csrf.api";

const SignUpTemplate: FC = () => {
  useEffect(() => {
    async function init() {
      await setCsrfToken();
    }
    init();
  }, []);

  return (
    <>
      <SupporterRegistrationForm />
    </>
  );
};

export default SignUpTemplate;

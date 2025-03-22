"use client";

import { FC } from "react";
import { FormType } from "../../_types";
import SupporterSignUpThanks from "../SupporterSignUpThanks";
import CompanySignUpThanks from "../CompanySignUpThanks";

type Props = {
  formType: FormType;
};

const SignUpThanks: FC<Props> = ({ formType }: Props) => {
  const formComponent = () => {
    switch (formType) {
      case "supporter":
        return <SupporterSignUpThanks />;
      case "company":
        return <CompanySignUpThanks />;
    }
  };

  return <>{formComponent()}</>;
};

export default SignUpThanks;

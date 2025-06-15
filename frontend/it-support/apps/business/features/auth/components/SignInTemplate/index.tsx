import { SignInFormType } from "@/types";
import { FC, useState } from "react";
import SupporterSignInForm from "../SupporterSignInForm";
import CompanySignInForm from "../CompanySignInForm";

const SignInTemplate: FC = () => {
  const [formType, setFormType] = useState<SignInFormType>("supporter");

  const formComponents = () => {
    switch (formType) {
      case "supporter":
        return (
          <>
            <SupporterSignInForm formType={formType} switchFormType={setFormType} />
          </>
        );
      case "company":
        return (
          <>
            <CompanySignInForm formType={formType} switchFormType={setFormType} />
          </>
        );
    }
  };

  return <>{formComponents()}</>;
};

export default SignInTemplate;

import { FC } from "react";
import { FormType, PhaseType } from "../../_types";
import SupporterSignUpInput from "../SupporterSignUpInput";
import CompanySignUpInput from "../CompanySignUpInput";

type Props = {
  formType: FormType;
  togglePhase: (newPhase: PhaseType) => void;
  switchFormType: (newFormType: FormType) => void;
};

const SignUpInput: FC<Props> = ({ formType, togglePhase, switchFormType }: Props) => {
  const formComponent = () => {
    switch (formType) {
      case "supporter":
        return (
          <SupporterSignUpInput
            togglePhase={togglePhase}
            formType={formType}
            switchFormType={switchFormType}
          />
        );
      case "company":
        return (
          <CompanySignUpInput
            togglePhase={togglePhase}
            formType={formType}
            switchFormType={switchFormType}
          />
        );
    }
  };

  return <>{formComponent()}</>;
};

export default SignUpInput;

import { FC } from "react";
import { FormType, PhaseType } from "../../_types";
import SupporterSignUpConfirm from "../SupporterSignUpConfirm";
import CompanySignUpConfirm from "../CompanySignUpConfirm";

type Props = {
  formType: FormType;
  togglePhase: (newPhase: PhaseType) => void;
};

const SignUpConfirm: FC<Props> = ({ formType, togglePhase }: Props) => {
  const formComponent = () => {
    switch (formType) {
      case "supporter":
        return <SupporterSignUpConfirm togglePhase={togglePhase} />;
      case "company":
        return <CompanySignUpConfirm togglePhase={togglePhase} />;
    }
  };

  return <>{formComponent()}</>;
};

export default SignUpConfirm;

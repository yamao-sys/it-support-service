import { FC, useState } from "react";
import { FormType, PhaseType } from "../../_types";
import SignUpLayout from "../SignUpLayout";
import SignUpInput from "../SignUpInput";
import SignUpConfirm from "../SignUpConfirm";
import SignUpThanks from "../SignUpThanks";

const SignUpForm: FC = () => {
  const [phase, setPhase] = useState<PhaseType>("input");
  const [formType, setFormType] = useState<FormType>("supporter");

  const togglePhase = (newPhase: PhaseType) => setPhase(newPhase);
  const switchFormType = (newFormType: FormType) => setFormType(newFormType);

  const phaseComponent = () => {
    switch (phase) {
      case "input":
        return (
          <SignUpInput
            formType={formType}
            switchFormType={switchFormType}
            togglePhase={togglePhase}
          />
        );
      case "confirm":
        return <SignUpConfirm formType={formType} togglePhase={togglePhase} />;
      case "thanks":
        return <SignUpThanks formType={formType} />;
    }
  };

  return <SignUpLayout phase={phase}>{phaseComponent()}</SignUpLayout>;
};

export default SignUpForm;

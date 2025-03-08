"use client";

import { FC, useState } from "react";
import SupporterRegistrationForm from "../SupporterRegistrationForm";
import SupporterSignUpConfirm from "../SupporterSignUpConfirm";
import SignUpLayout from "../SignUpLayout";
import SupporterSignUpThanks from "../SupporterSignUpThanks";

export type Phase = "input" | "confirm" | "thanks";

const SupporterSignUpForm: FC = () => {
  const [phase, setPhase] = useState<Phase>("input");
  const togglePhase = (newPhase: Phase) => setPhase(newPhase);

  const phaseComponent = () => {
    switch (phase) {
      case "input":
        return <SupporterRegistrationForm togglePhase={togglePhase} />;
      case "confirm":
        return <SupporterSignUpConfirm togglePhase={togglePhase} />;
      case "thanks":
        return <SupporterSignUpThanks />;
    }
  };

  return <SignUpLayout phase={phase}>{phaseComponent()}</SignUpLayout>;
};

export default SupporterSignUpForm;

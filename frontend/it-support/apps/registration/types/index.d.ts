import { components } from "@/apis/generated/apiSchema";

export type FormType = "supporter" | "company";

export type PhaseType = "input" | "confirm" | "thanks";

export type SupporterSignUpInput = components["schemas"]["SupporterSignUpInput"];

export type SupporterSignUpValidationError =
  components["schemas"]["SupporterSignUpResponse"]["errors"];

export type CompanySignUpInput = components["schemas"]["CompanySignUpInput"];

export type CompanySignUpValidationError = components["schemas"]["CompanySignUpResponse"]["errors"];

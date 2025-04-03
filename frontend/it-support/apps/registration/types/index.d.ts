import { components } from "@/apis/generated/apiSchema";

export type FormType = "supporter" | "company";

export type PhaseType = "input" | "confirm" | "thanks";

export type SupporterSignUpInput =
  components["requestBodies"]["SupporterSignUpInput"]["content"]["multipart/form-data"];

export type SupporterSignUpValidationError =
  components["responses"]["SupporterSignUpResponse"]["content"]["application/json"]["errors"];

export type CompanySignUpInput =
  components["requestBodies"]["CompanySignUpInput"]["content"]["multipart/form-data"];

export type CompanySignUpValidationError =
  components["responses"]["CompanySignUpResponse"]["content"]["application/json"]["errors"];

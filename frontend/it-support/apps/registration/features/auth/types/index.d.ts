import { components } from "@/apis/generated/supporters/apiSchema";
import { components as companyComponents } from "@/apis/generated/companies/apiSchema";

export type FormType = "supporter" | "company";

export type PhaseType = "input" | "confirm" | "thanks";

export type SupporterSignUpInput =
  components["requestBodies"]["SignUpInput"]["content"]["multipart/form-data"];

export type SupporterSignUpValidationError =
  components["responses"]["SignUpResponse"]["content"]["application/json"]["errors"];

export type CompanySignUpInput =
  companyComponents["requestBodies"]["SignUpInput"]["content"]["multipart/form-data"];

export type CompanySignUpValidationError =
  companyComponents["responses"]["SignUpResponse"]["content"]["application/json"]["errors"];

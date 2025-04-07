import { components } from "../apis/generated/apiSchema";

export type SignInFormType = "supporter" | "company";

export type SupporterSignInInput =
  components["requestBodies"]["SupporterSignInInput"]["content"]["application/json"];

export type CompanySignInInput =
  components["requestBodies"]["CompanySignInInput"]["content"]["application/json"];

export type ProjectStoreInput =
  components["requestBodies"]["ProjectStoreInput"]["content"]["application/json"];

export type ProjectValidationError = components["schemas"]["ProjectValidationError"];

import { components } from "../apis/generated/apiSchema";

export type SignInFormType = "supporter" | "company";

export type SupporterSignInInput = components["schemas"]["SupporterSignInInput"];

export type CompanySignInInput = components["schemas"]["CompanySignInInput"];

export type ProjectStoreInput = components["schemas"]["ProjectStoreInput"];

export type ProjectValidationError = components["schemas"]["ProjectValidationError"];

export type ProjectResponse = components["schemas"]["ProjectResponse"];

export type Project = components["schemas"]["Project"];

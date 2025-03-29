import { components as supporterComponents } from "@/apis/generated/supporters/apiSchema";
import { components as companyComponents } from "@/apis/generated/companies/apiSchema";

export type SignInFormType = "supporter" | "company";

export type SupporterSignInInput =
  supporterComponents["requestBodies"]["SignInInput"]["content"]["application/json"];

export type CompanySignInInput =
  companyComponents["requestBodies"]["SignInInput"]["content"]["application/json"];

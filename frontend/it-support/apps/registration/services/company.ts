"use server";

import { CompaniesApi, PostCompanySignUpRequest } from "@/apis";
import { getApiConfig } from "@/apis/client";

export async function postCompanyValidateSignUp(input: PostCompanySignUpRequest) {
  const apiConfig = await getApiConfig();
  const client = new CompaniesApi(apiConfig);

  const res = await client.postCompanyValidateSignUp(input);

  return res.errors;
}

export async function postCompanySignUp(input: PostCompanySignUpRequest) {
  const apiConfig = await getApiConfig();
  const client = new CompaniesApi(apiConfig);

  const res = await client.postCompanySignUp(input);

  return res.errors;
}

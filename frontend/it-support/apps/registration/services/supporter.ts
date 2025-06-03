"use server";

import { PostSupporterSignUpRequest, SupportersApi } from "@/apis";
import { getApiConfig } from "@/apis/client";

export async function postSupporterValidateSignUp(input: PostSupporterSignUpRequest) {
  const apiConfig = await getApiConfig();
  const client = new SupportersApi(apiConfig);

  const res = await client.postSupporterValidateSignUp(input);

  return res.errors;
}

export async function postSupporterSignUp(input: PostSupporterSignUpRequest) {
  const apiConfig = await getApiConfig();
  const client = new SupportersApi(apiConfig);

  const res = await client.postSupporterSignUp(input);

  return res.errors;
}

"use server";

import { paths } from "@/apis/generated/supporters/apiSchema";
import { cookies } from "next/headers";
import createClient from "openapi-fetch";
import { SupporterSignUpInput } from "../_components/SupporterRegistrationForm";

const client = createClient<paths>({
  baseUrl: "http://registration_api:8080/",
  credentials: "include",
});

const getRequestHeaders = async () => {
  const csrfToken = (await cookies()).get("_csrf")?.value ?? "";
  return {
    headers: {
      "X-CSRF-Token": csrfToken,
      Cookie: `_csrf=${csrfToken}`,
    },
  };
};

export async function postValidateSignUp(input: SupporterSignUpInput) {
  const body = new FormData();
  for (const [key, value] of Object.entries(input)) {
    body.append(key, value);
  }

  const { data, error } = await client.POST("/supporters/validateSignUp", {
    ...(await getRequestHeaders()),
    body: body as unknown as SupporterSignUpInput,
  });
  if (error?.code === 500 || data === undefined) {
    throw Error("Internal Server Error");
  }

  return data;
}

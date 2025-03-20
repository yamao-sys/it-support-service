"use server";

import { paths } from "@/apis/generated/supporters/apiSchema";
import { cookies } from "next/headers";
import createClient from "openapi-fetch";
import { SupporterSignUpInput } from "../_components/SupporterRegistrationForm";

const client = createClient<paths>({
  baseUrl: `${process.env.REGISTRATION_API_ENDPOINT_URI}/`,
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
  const { data, error } = await client.POST("/supporters/validateSignUp", {
    ...(await getRequestHeaders()),
    body: input,
    bodySerializer(body) {
      const formData = new FormData();

      if (body) {
        for (const [key, value] of Object.entries(input)) {
          if (value instanceof File) {
            formData.append(key, value, encodeURI(value.name));
          } else {
            formData.append(key, value);
          }
        }
      }
      return formData;
    },
  });
  if (error?.code === 500 || data === undefined) {
    throw Error("Internal Server Error");
  }

  return data;
}

export async function postSignUp(input: SupporterSignUpInput) {
  const { data, error } = await client.POST("/supporters/signUp", {
    ...(await getRequestHeaders()),
    body: input,
    bodySerializer(body) {
      const formData = new FormData();

      if (body) {
        for (const [key, value] of Object.entries(input)) {
          if (value instanceof File) {
            formData.append(key, value, encodeURI(value.name));
          } else {
            formData.append(key, value);
          }
        }
      }
      return formData;
    },
  });
  if (error?.code === 500 || data === undefined) {
    throw Error("Internal Server Error");
  }

  return data;
}

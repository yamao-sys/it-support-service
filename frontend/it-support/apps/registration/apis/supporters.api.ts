"use server";

import createClient from "openapi-fetch";
import { SupporterSignUpInput } from "../types";
import { getRequestHeaders } from "./csrf.api";
import { paths } from "./generated/apiSchema";

const client = createClient<paths>({
  baseUrl: `${process.env.REGISTRATION_API_ENDPOINT_URI}/`,
  credentials: "include",
});

export async function postSupporterValidateSignUp(input: SupporterSignUpInput) {
  const { data, error } = await client.POST("/supporters/validateSignUp", {
    ...(await getRequestHeaders()),
    body: input,
    bodySerializer(body) {
      const formData = new FormData();

      if (body) {
        for (const [key, value] of Object.entries(input)) {
          if (value instanceof File) {
            formData.append(key, value, encodeURI(value.name));
          } else if (value instanceof Date) {
            formData.append(key, value.toString());
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

export async function postSupporterSignUp(input: SupporterSignUpInput) {
  const { data, error } = await client.POST("/supporters/signUp", {
    ...(await getRequestHeaders()),
    body: input,
    bodySerializer(body) {
      const formData = new FormData();

      if (body) {
        for (const [key, value] of Object.entries(input)) {
          if (value instanceof File) {
            formData.append(key, value, encodeURI(value.name));
          } else if (value instanceof Date) {
            formData.append(key, value.toString());
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

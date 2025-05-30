"use server";

import createClient from "openapi-fetch";
import { CompanySignUpInput } from "../types";
import { getRequestHeaders } from "./csrf.api";
import { paths } from "./generated/apiSchema";

const client = createClient<paths>({
  baseUrl: `${process.env.REGISTRATION_API_ENDPOINT_URI}/`,
  credentials: "include",
});

export async function postCompanyValidateSignUp(input: CompanySignUpInput) {
  const { data, error } = await client.POST("/companies/validate-sign-up", {
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

export async function postCompanySignUp(input: CompanySignUpInput) {
  const { data, error } = await client.POST("/companies/sign-up", {
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

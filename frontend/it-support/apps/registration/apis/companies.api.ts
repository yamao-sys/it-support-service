"use server";

import { paths } from "@/apis/generated/companies/apiSchema";
import createClient from "openapi-fetch";
import { CompanySignUpInput } from "../types";
import { getRequestHeaders } from "./csrf.api";

const client = createClient<paths>({
  baseUrl: `${process.env.REGISTRATION_API_ENDPOINT_URI}/`,
  credentials: "include",
});

export async function postCompanyValidateSignUp(input: CompanySignUpInput) {
  const { data, error } = await client.POST("/companies/validateSignUp", {
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
  const { data, error } = await client.POST("/companies/signUp", {
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

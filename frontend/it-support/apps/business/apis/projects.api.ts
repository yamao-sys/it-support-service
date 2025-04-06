"use server";

import createClient from "openapi-fetch";
import { ProjectStoreInput } from "../types";
import { getRequestHeaders } from "./csrf.api";
import { paths } from "./generated/apiSchema";

const client = createClient<paths>({
  baseUrl: `${process.env.BUSINESS_API_ENDPOINT_URI}/`,
  credentials: "include",
});

export async function postProjectCreate(input: ProjectStoreInput) {
  const { data, response } = await client.POST("/projects", {
    ...(await getRequestHeaders()),
    body: input,
    bodySerializer(body) {
      const reqBody: { [key: string]: string | number | boolean } = {};
      if (body) {
        for (const [key, value] of Object.entries(input)) {
          if (value instanceof Date) {
            reqBody[key] = value
              .toLocaleDateString("ja-JP", { year: "numeric", month: "2-digit", day: "2-digit" })
              .replaceAll("/", "-");
          } else if (["minBudget", "maxBudget"].includes(key)) {
            reqBody[key] = Number(value);
          } else if (["isActive"].includes(key)) {
            reqBody[key] = Boolean(value);
          } else {
            reqBody[key] = value;
          }
        }
      }
      console.log(JSON.stringify(reqBody));
      return JSON.stringify(reqBody);
    },
  });
  console.log(response);
  if (data === undefined || response.status === 500) {
    throw Error("Internal Server Error");
  }

  return data.errors;
}

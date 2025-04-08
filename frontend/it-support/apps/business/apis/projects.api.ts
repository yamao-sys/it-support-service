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
    bodySerializer() {
      const reqBody: { [key: string]: string | number | boolean } = {};
      for (const [key, value] of Object.entries(input)) {
        if (value instanceof Date) {
          reqBody[key] = value
            .toLocaleDateString("ja-JP", { year: "numeric", month: "2-digit", day: "2-digit" })
            .replaceAll("/", "-");
        } else if (["minBudget", "maxBudget"].includes(key)) {
          if (value) {
            reqBody[key] = Number(value);
          }
        } else if (["isActive"].includes(key)) {
          reqBody[key] = value === "1" ? true : false;
        } else {
          reqBody[key] = value;
        }
      }
      return JSON.stringify(reqBody);
    },
  });
  if (data === undefined || response.status === 500) {
    throw Error("Internal Server Error");
  }

  return data.errors;
}

export async function getProject(id: number) {
  const { data, response } = await client.GET("/projects/{id}", {
    ...(await getRequestHeaders()),
    params: {
      path: {
        id,
      },
    },
  });
  if (data === undefined || response.status === 404) {
    throw Error("Not Found Error");
  }

  return data.project;
}

export async function putUpdateProject(id: number, input: ProjectStoreInput) {
  console.log(input);
  const { data, response } = await client.PUT("/projects/{id}", {
    ...(await getRequestHeaders()),
    params: {
      path: {
        id,
      },
    },
    body: input,
    bodySerializer() {
      const reqBody: { [key: string]: string | number | boolean } = {};
      for (const [key, value] of Object.entries(input)) {
        if (value instanceof Date) {
          reqBody[key] = value
            .toLocaleDateString("ja-JP", { year: "numeric", month: "2-digit", day: "2-digit" })
            .replaceAll("/", "-");
        } else if (["minBudget", "maxBudget"].includes(key)) {
          if (value) {
            reqBody[key] = Number(value);
          }
        } else {
          reqBody[key] = value;
        }
      }
      return JSON.stringify(reqBody);
    },
  });
  if (data === undefined || response.status === 404 || response.status === 500) {
    throw Error("Internal Server Error");
  }

  return data.errors;
}

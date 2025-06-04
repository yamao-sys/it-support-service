"use server";

import {
  GetProjectsRequest,
  PostProjectRequest,
  ProjectsApi,
  ProjectStoreInput,
  ResponseError,
} from "@/apis";
import { getApiConfig } from "@/apis/client";

export async function postProjectCreate(input: PostProjectRequest) {
  const apiConfig = await getApiConfig();
  const client = new ProjectsApi(apiConfig);

  try {
    const res = await client.postProject(input);
    return res.errors;
  } catch (error) {
    if (error instanceof ResponseError) {
      console.error("Error response:", error.response);
      // switch (error.response.status) {
      //   // TODO: 400でバリデーションエラーを返すようにする
      //   case 400:
      //     error.response.
      //   case 500:
      // }
      throw new Error(`Internal Server Error: ${error}`);
    } else {
      // NOTE: ネットワークエラーなどの一般的なエラー
      throw new Error(`Unexpected error: ${error}`);
    }
  }
}

export async function getProjects(nextPageToken?: string) {
  const apiConfig = await getApiConfig();
  const client = new ProjectsApi(apiConfig);

  const params: GetProjectsRequest = {};
  if (nextPageToken) {
    params.pageToken = nextPageToken;
  }
  try {
    const res = await client.getProjects(params);
    return res;
  } catch (error) {
    if (error instanceof ResponseError) {
      switch (error.response.status) {
        case 404:
          throw new Error("Not Found Error");
        case 500:
          throw new Error(`Internal Server Error: ${error}`);
      }
    } else {
      // NOTE: ネットワークエラーなどの一般的なエラー
      throw new Error(`Unexpected error: ${error}`);
    }
  }
}

export async function getProject(id: number) {
  const apiConfig = await getApiConfig();
  const client = new ProjectsApi(apiConfig);

  try {
    const res = await client.getProject({ id });
    return res.project;
  } catch (error) {
    if (error instanceof ResponseError) {
      switch (error.response.status) {
        case 404:
          throw new Error("Not Found Error");
        case 500:
          throw new Error(`Internal Server Error: ${error}`);
      }
    } else {
      // NOTE: ネットワークエラーなどの一般的なエラー
      throw new Error(`Unexpected error: ${error}`);
    }
  }
}

export async function putUpdateProject(id: number, input: ProjectStoreInput) {
  const apiConfig = await getApiConfig();
  const client = new ProjectsApi(apiConfig);

  try {
    const res = await client.putProject({ id, projectStoreInput: input });
    return res.errors;
  } catch (error) {
    if (error instanceof ResponseError) {
      switch (error.response.status) {
        case 404:
          throw new Error("Not Found Error");
        case 500:
          throw new Error(`Internal Server Error: ${error}`);
      }
    } else {
      // NOTE: ネットワークエラーなどの一般的なエラー
      throw new Error(`Unexpected error: ${error}`);
    }
  }
}

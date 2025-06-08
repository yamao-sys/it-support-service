import { GetToProjectsRequest, ResponseError, ToProjectsApi } from "@/apis";
import { getApiConfig } from "@/apis/client";

export async function getToProjects(nextPageToken?: string) {
  const apiConfig = await getApiConfig();
  const client = new ToProjectsApi(apiConfig);

  const params: GetToProjectsRequest = {};
  if (nextPageToken) {
    params.pageToken = nextPageToken;
  }
  try {
    const res = await client.getToProjects(params);
    return res;
  } catch (error) {
    if (error instanceof ResponseError) {
      switch (error.response.status) {
        case 500:
          throw new Error(`Internal Server Error: ${error}`);
      }
    } else {
      // NOTE: ネットワークエラーなどの一般的なエラー
      throw new Error(`Unexpected error: ${error}`);
    }
  }
}

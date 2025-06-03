import { cookies } from "next/headers";
import { Configuration } from "./runtime";

const getRequestHeaders = async () => {
  const csrfToken = (await cookies()).get("_csrf")?.value ?? "";
  return {
    headers: {
      "X-CSRF-Token": csrfToken,
      Cookie: `_csrf=${csrfToken}`,
    },
  };
};

export const getApiConfig = async () => {
  const headers = await getRequestHeaders();

  return new Configuration({
    basePath: process.env.REGISTRATION_API_ENDPOINT_URI,
    credentials: "include",
    headers: headers.headers,
  });
};

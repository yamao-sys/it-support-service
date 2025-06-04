import { cookies } from "next/headers";
import { Configuration } from "./runtime";

const getRequestHeaders = async () => {
  const csrfToken = (await cookies()).get("_csrf")?.value ?? "";
  const authenticateToken = (await cookies()).get("token")?.value ?? "";
  return {
    headers: {
      "X-CSRF-Token": csrfToken,
      Cookie: `_csrf=${csrfToken}; token=${authenticateToken}`,
    },
  };
};

export const getApiConfig = async () => {
  const headers = await getRequestHeaders();

  return new Configuration({
    basePath: process.env.BUSINESS_API_ENDPOINT_URI,
    credentials: "include",
    headers: headers.headers,
  });
};

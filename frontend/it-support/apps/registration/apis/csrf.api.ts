"use server";

import createClient from "openapi-fetch";
import { paths } from "./generated/csrf/apiSchema";
import { cookies } from "next/headers";

const client = createClient<paths>({
  baseUrl: "http://registration_api:8080/",
  credentials: "include",
});

export async function setCsrfToken() {
  const { data, error, response } = await client.GET("/csrf", {});
  if (error?.code === 500 || data === undefined) {
    throw Error();
  }

  // NOTE: クライアントにCookieをセット
  const setCookie = response.headers.get("set-cookie");
  if (!setCookie) {
    throw Error();
  }
  const csrfToken = setCookie?.split(";")[0]?.split("=")[1];
  if (!csrfToken) {
    throw Error();
  }

  // TODO: cookieの属性は環境変数に切り出す
  (await cookies()).set({
    name: "_csrf",
    value: csrfToken,
    secure: true,
    sameSite: "none",
    httpOnly: true,
  });
}

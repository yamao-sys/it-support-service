"use server";

import { paths } from "@/apis/generated/supporters/apiSchema";
import createClient from "openapi-fetch";
import { SupporterSignInInput } from "./types";
import { getRequestHeaders } from "./csrf.api";
import { cookies } from "next/headers";

const client = createClient<paths>({
  baseUrl: `${process.env.BUSINESS_API_ENDPOINT_URI}/`,
  credentials: "include",
});

export async function postSupporterSignIn(input: SupporterSignInInput) {
  const { response } = await client.POST("/supporters/signIn", {
    ...(await getRequestHeaders()),
    body: input,
  });
  if (response.status === 500) {
    throw Error("Internal Server Error");
  }
  if (response.status === 400) {
    return "メールアドレスまたはパスワードが正しくありません";
  }

  // NOTE: クライアントにCookieをセット
  const setCookie = response.headers.get("set-cookie");
  if (!setCookie) {
    throw Error();
  }
  const token = setCookie?.split(";")[0]?.split("=")[1];
  if (!token) {
    throw Error();
  }

  // TODO: cookieの属性は環境変数に切り出す
  (await cookies()).set({
    name: "token",
    value: token,
    secure: true,
    sameSite: "none",
    httpOnly: true,
  });

  return "";
}

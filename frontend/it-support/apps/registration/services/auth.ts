"use server";

import { CsrfApi } from "@/apis";
import { getApiConfig } from "@/apis/client";
import { cookies } from "next/headers";

export async function setCsrfToken() {
  const apiConfig = await getApiConfig();
  const client = await new CsrfApi(apiConfig).getCsrf();

  // TODO: cookieの属性は環境変数に切り出す
  (await cookies()).set({
    name: "_csrf",
    value: client.csrfToken,
    secure: true,
    sameSite: "none",
    httpOnly: true,
  });
}

"use server";

import { CompaniesApi, PostCompanySignInRequest, ResponseError } from "@/apis";
import { getApiConfig } from "@/apis/client";
import { cookies } from "next/headers";

export async function postCompanySignIn(input: PostCompanySignInRequest) {
  const apiConfig = await getApiConfig();
  const client = new CompaniesApi(apiConfig);

  try {
    const res = await client.postCompanySignIn(input);

    // TODO: cookieの属性は環境変数に切り出す
    (await cookies()).set({
      name: "token",
      value: res.token,
      secure: true,
      sameSite: "none",
      httpOnly: true,
    });

    return "";
  } catch (error) {
    if (error instanceof ResponseError) {
      switch (error.response.status) {
        case 400:
          return "メールアドレスまたはパスワードが正しくありません";
        case 500:
          throw new Error(`Internal Server Error: ${error}`);
      }
    } else {
      // NOTE: ネットワークエラーなどの一般的なエラー
      throw new Error(`Unexpected error: ${error}`);
    }
  }
}

"use server";

import { PostSupporterSignInRequest, SupportersApi } from "@/apis";
import { getApiConfig } from "@/apis/client";
import { cookies } from "next/headers";

export async function postSupporterSignIn(input: PostSupporterSignInRequest) {
  const apiConfig = await getApiConfig();
  const client = new SupportersApi(apiConfig);

  try {
    const res = await client.postSupporterSignIn(input);

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
    if (error instanceof Response) {
      switch (error.status) {
        case 400:
          return "メールアドレスまたはパスワードが正しくありません";
        case 500:
          throw new Error("Internal Server Error");
        default:
          throw new Error(`Unexpected error: ${error.status}`);
      }
    } else {
      // NOTE: ネットワークエラーなどの一般的なエラー
      console.error("Unexpected error:", error);
    }
  }
}

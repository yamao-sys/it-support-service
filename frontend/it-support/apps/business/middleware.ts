import { NextResponse, type NextRequest } from "next/server";
import { setCsrfToken } from "./services/auth";

// const needsAuthPath = ['/reading_records'];

// const needsAuth = (path: string) =>
//   needsAuthPath.some((pathPattern) => path.startsWith(pathPattern));

export async function middleware(request: NextRequest) {
  // let response = NextResponse.next({
  //   request: {
  //     headers: request.headers,
  //   },
  // });
  // // NOTE: ミドルウェアによるチェックが不要なページは、認証チェックなど不要なチェック処理が入ることを防ぐため、早期リターン
  // if (!needsAuth(request.nextUrl.pathname) && request.nextUrl.pathname !== '/') {
  //   return response;
  // }

  // const isSignedIn = await checkSignedIn();

  // // NOTE: 要認証ページで認証済みでなければログインページへ
  // if (needsAuth(request.nextUrl.pathname) && !isSignedIn) {
  //   return NextResponse.redirect(new URL('/sign_in', request.url));
  // }
  // // NOTE: TOPの向き先 認証済みの場合は読書記録一覧、そうでなければログイン画面へ
  // if (request.nextUrl.pathname === '/') {
  //   if (isSignedIn) {
  //     return NextResponse.redirect(new URL('/reading_records', request.url));
  //   }

  //   return NextResponse.redirect(new URL('/sign_in', request.url));
  // }

  // NOTE: CSRFトークンをクッキーにセット
  await setCsrfToken();

  return NextResponse.next({
    request: {
      headers: request.headers,
    },
  });
}

// NOTE: ビルドファイル, 画像最適化エンドポイント, ファビコン, 画像は対象外
export const config = {
  matcher: ["/((?!_next/static|_next/image|favicon.ico|.*\\.(?:svg|png|jpg|jpeg|gif|webp)$).*)"],
};

import { test, expect } from "@playwright/test";

test("SignUp Successfully", async ({ page }) => {
  await page.goto("/sign_up");

  // NOTE: 会員登録フォームを入力
  await page.getByRole("textbox", { name: "姓" }).fill("test_last_name");
  await page.getByRole("textbox", { name: "名" }).fill("test_first_name");
  await page.getByRole("textbox", { name: "Email" }).fill("test@example.com");
  await page.getByRole("textbox", { name: "パスワード" }).fill("password");

  await page.getByRole("button", { name: "確認画面へ" }).click();

  // NOTE: 確認画面で入力内容を表示できること
  await expect(page.getByText("test_last_name")).toBeVisible();
  await expect(page.getByText("test_first_name")).toBeVisible();
  await expect(page.getByText("test@example.com")).toBeVisible();
  await expect(page.getByText("********")).toBeVisible();

  await page.getByRole("button", { name: "登録する" }).click();

  // NOTE: 登録に成功すると、サンクス画面に遷移する
  await expect(page.getByText("会員登録が完了しました。")).toBeVisible();
});

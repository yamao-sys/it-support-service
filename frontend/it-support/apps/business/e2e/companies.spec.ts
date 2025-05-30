import { test, expect } from "@playwright/test";

test.describe("/companies/sign_in", () => {
  test("正常系", async ({ page }) => {
    await page.goto("/sign_in");

    await page.getByRole("radio", { name: "企業" }).check();

    // NOTE: ログインフォームを入力
    await page.getByRole("textbox", { name: "Email" }).fill("test@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "ログイン" }).click();

    // NOTE: ログイン成功
    page.on("dialog", async (dialog) => {
      expect(dialog.message()).toContain("企業のログインに成功しました!");
      await dialog.accept();
    });
    await page.waitForURL("/");
    await expect(page).toHaveURL("/");
  });

  test("異常系", async ({ page }) => {
    await page.goto("/sign_in");

    await page.getByRole("radio", { name: "企業" }).check();

    // NOTE: ログインフォームを入力
    await page.getByRole("textbox", { name: "Email" }).fill("test@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("passwor");
    await page.getByRole("button", { name: "ログイン" }).click();

    // NOTE: バリデーションエラーが表示されること
    await expect(page.getByText("メールアドレスまたはパスワードが正しくありません")).toBeVisible();

    // 入力し直してログインできること
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "ログイン" }).click();

    // NOTE: ログイン成功
    page.on("dialog", async (dialog) => {
      expect(dialog.message()).toContain("企業のログインに成功しました!");
      await dialog.accept();
    });
    await page.waitForURL("/");
    await expect(page).toHaveURL("/");
  });

  test("フォーム種別変更あり", async ({ page }) => {
    await page.goto("/sign_in");

    await page.getByRole("radio", { name: "企業" }).check();

    // NOTE: ログインフォームを入力
    await page.getByRole("textbox", { name: "Email" }).fill("test_@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password_");

    await page.getByRole("radio", { name: "サポータ" }).check();
    await expect(page.getByText("サポータ ログインフォーム", { exact: true })).toBeVisible();

    // NOTE: もう一度企業フォームに戻る
    await page.getByRole("radio", { name: "企業" }).check();

    // NOTE: フォーム内容を入力し直しできること
    await page.getByRole("textbox", { name: "Email" }).fill("test@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");

    await page.getByRole("button", { name: "ログイン" }).click();

    // NOTE: ログイン成功
    page.on("dialog", async (dialog) => {
      expect(dialog.message()).toContain("企業のログインに成功しました!");
      await dialog.accept();
    });
    await page.waitForURL("/");
    await expect(page).toHaveURL("/");
  });
});

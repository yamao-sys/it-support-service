import { test, expect } from "@playwright/test";

test.describe("/projects", () => {
  test.describe("POST /projects", () => {
    test("正常系", async ({ page }) => {
      await page.goto("/sign_in");

      await page.getByRole("radio", { name: "企業" }).check();

      // NOTE: ログインフォームを入力
      await page.getByRole("textbox", { name: "Email" }).fill("test@example.com");
      await page.getByRole("textbox", { name: "パスワード" }).fill("password");
      await page.getByRole("button", { name: "ログイン" }).click();

      // NOTE: ログイン成功
      await page.waitForURL("/");

      // NOTE: 案件登録ページへ遷移
      await page.goto("/projects/new");

      // NOTE: 登録内容の入力
      await page.getByRole("textbox", { name: "案件タイトル" }).fill("test title");
      await page.getByRole("textbox", { name: "案件概要" }).fill("test description\ntest");
      await page.click('input[name="startDate"]');
      await page.locator('[aria-label="Choose Monday, April 7th, 2025"]').click();
      await page.click('input[name="endDate"]');
      await page.locator('[aria-label="Choose Tuesday, April 8th, 2025"]').click();
      await page.getByRole("spinbutton", { name: "予算(下限)" }).fill("100000");
      await page.getByRole("spinbutton", { name: "予算(上限)" }).fill("100001");
      await page.locator('label[for="active"]').click();

      await page.getByRole("button", { name: "保存する" }).click();

      page.on("dialog", async (dialog) => {
        expect(dialog.message()).toContain("案件を更新しました!");
        await dialog.accept();
      });
      await page.waitForURL("/");
      await expect(page).toHaveURL("/");
    });

    // test("異常系", async ({ page }) => {
    //   await page.goto("/sign_in");

    //   await page.getByRole("radio", { name: "企業" }).check();

    //   // NOTE: ログインフォームを入力
    //   await page.getByRole("textbox", { name: "Email" }).fill("test@example.com");
    //   await page.getByRole("textbox", { name: "パスワード" }).fill("passwor");
    //   await page.getByRole("button", { name: "ログイン" }).click();

    //   // NOTE: バリデーションエラーが表示されること
    //   await expect(
    //     page.getByText("メールアドレスまたはパスワードが正しくありません"),
    //   ).toBeVisible();

    //   // 入力し直してログインできること
    //   await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    //   await page.getByRole("button", { name: "ログイン" }).click();

    //   // NOTE: ログイン成功
    //   page.on("dialog", async (dialog) => {
    //     expect(dialog.message()).toContain("企業のログインに成功しました!");
    //     await dialog.accept();
    //   });
    //   await page.waitForURL("/");
    //   await expect(page).toHaveURL("/");
    // });
  });
});

import { test, expect } from "@playwright/test";

test.describe("/supporters/sign_up", () => {
  test("正常系_必須項目のみ", async ({ page }) => {
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

  test("正常系_必須項目のみ_入力に戻って修正", async ({ page }) => {
    await page.goto("/sign_up");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "姓" }).fill("test_last_name");
    await page.getByRole("textbox", { name: "名" }).fill("test_first_name");
    await page.getByRole("textbox", { name: "Email" }).fill("test@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");

    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 入力画面で編集
    await page.getByRole("button", { name: "入力へ戻る" }).click();
    await page.getByRole("textbox", { name: "Email" }).fill("test_modified@example.com");
    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name")).toBeVisible();
    await expect(page.getByText("test_first_name")).toBeVisible();
    await expect(page.getByText("test_modified@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();

    await page.getByRole("button", { name: "登録する" }).click();

    // NOTE: 登録に成功すると、サンクス画面に遷移する
    await expect(page.getByText("会員登録が完了しました。")).toBeVisible();
  });

  test("異常系_必須項目のみ", async ({ page }) => {
    await page.goto("/sign_up");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "パスワード" }).fill("passwor");

    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: バリデーションエラーが表示されること
    await expect(page.getByText("姓は必須入力です。")).toBeVisible();
    await expect(page.getByText("名は必須入力です。")).toBeVisible();
    await expect(page.getByText("Emailは必須入力です。")).toBeVisible();
    await expect(page.getByText("パスワードは8 ~ 24文字での入力をお願いします。")).toBeVisible();

    // 入力し直して確認 → 登録できること
    await page.getByRole("textbox", { name: "姓" }).fill("test_last_name");
    await page.getByRole("textbox", { name: "名" }).fill("test_first_name");
    await page.getByRole("textbox", { name: "Email" }).fill("test_required_invalid@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name")).toBeVisible();
    await expect(page.getByText("test_first_name")).toBeVisible();
    await expect(page.getByText("test_required_invalid@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();

    await page.getByRole("button", { name: "登録する" }).click();

    // NOTE: 登録に成功すると、サンクス画面に遷移する
    await expect(page.getByText("会員登録が完了しました。")).toBeVisible();
  });

  test("異常系_必須項目のみ_入力に戻って修正", async ({ page }) => {
    await page.goto("/sign_up");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "姓" }).fill("test_last_name");
    await page.getByRole("textbox", { name: "名" }).fill("test_first_name");
    await page.getByRole("textbox", { name: "Email" }).fill("test_required_invalid_re@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");

    await page.getByRole("button", { name: "確認画面へ" }).click();

    await page.getByRole("button", { name: "入力へ戻る" }).click();

    // NOTE: 再入力内容が不適であれば、バリデーションエラーが表示されること
    await page.getByRole("textbox", { name: "パスワード" }).fill("passwor");
    await page.getByRole("button", { name: "確認画面へ" }).click();
    await expect(page.getByText("パスワードは8 ~ 24文字での入力をお願いします。")).toBeVisible();

    // 入力し直して確認 → 登録できること
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name")).toBeVisible();
    await expect(page.getByText("test_first_name")).toBeVisible();
    await expect(page.getByText("test_required_invalid_re@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();

    await page.getByRole("button", { name: "登録する" }).click();

    // NOTE: 登録に成功すると、サンクス画面に遷移する
    await expect(page.getByText("会員登録が完了しました。")).toBeVisible();
  });
});

import { test, expect } from "@playwright/test";
import * as path from "path";
import { fileURLToPath } from "url";

test.describe("/supporters/sign_up", () => {
  test("正常系_必須項目のみ", async ({ page }) => {
    await page.goto("/sign_up");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "姓" }).fill("test_last_name");
    await page.getByRole("textbox", { name: "名" }).fill("test_first_name");
    await page.getByRole("textbox", { name: "Email" }).fill("test_supporter_1@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name test_first_name", { exact: true })).toBeVisible();
    await expect(page.getByText("test_supporter_1@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
    await page.getByRole("button", { name: "登録する" }).click();

    // NOTE: 登録に成功すると、サンクス画面に遷移する
    await expect(page.getByText("サポータ登録が完了しました。")).toBeVisible();
  });

  test("正常系_必須項目のみ_入力に戻って修正", async ({ page }) => {
    await page.goto("/sign_up");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "姓" }).fill("test_last_name");
    await page.getByRole("textbox", { name: "名" }).fill("test_first_name");
    await page.getByRole("textbox", { name: "Email" }).fill("test_supporter_2@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 入力画面で編集
    await page.getByRole("button", { name: "入力へ戻る" }).click();
    await page
      .getByRole("textbox", { name: "Email" })
      .fill("test_supporter_2_modified@example.com");
    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name test_first_name", { exact: true })).toBeVisible();
    await expect(page.getByText("test_supporter_2_modified@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
    await page.getByRole("button", { name: "登録する" }).click();

    // NOTE: 登録に成功すると、サンクス画面に遷移する
    await expect(page.getByText("サポータ登録が完了しました。")).toBeVisible();
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
    await page.getByRole("textbox", { name: "Email" }).fill("test_supporter_3@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name test_first_name", { exact: true })).toBeVisible();
    await expect(page.getByText("test_supporter_3@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
    await page.getByRole("button", { name: "登録する" }).click();

    // NOTE: 登録に成功すると、サンクス画面に遷移する
    await expect(page.getByText("サポータ登録が完了しました。")).toBeVisible();
  });

  test("異常系_必須項目のみ_入力に戻って修正", async ({ page }) => {
    await page.goto("/sign_up");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "姓" }).fill("test_last_name");
    await page.getByRole("textbox", { name: "名" }).fill("test_first_name");
    await page.getByRole("textbox", { name: "Email" }).fill("test_supporter_4@example.com");
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
    await expect(page.getByText("test_last_name test_first_name", { exact: true })).toBeVisible();
    await expect(page.getByText("test_supporter_4@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
    await page.getByRole("button", { name: "登録する" }).click();

    // NOTE: 登録に成功すると、サンクス画面に遷移する
    await expect(page.getByText("サポータ登録が完了しました。")).toBeVisible();
  });

  test("正常系_任意項目含む", async ({ page }) => {
    await page.goto("/sign_up");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "姓" }).fill("test_last_name");
    await page.getByRole("textbox", { name: "名" }).fill("test_first_name");
    await page.getByRole("textbox", { name: "Email" }).fill("test_supporter_5@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");

    // NOTE: 身分証明書の表・裏のアップロード
    const fileChooserPromise = page.waitForEvent("filechooser");
    await page.getByText("+ 身分証明書(表)を選択").click();
    const frontIdentificationFileChooser = await fileChooserPromise;
    await frontIdentificationFileChooser.setFiles(
      path.join(path.dirname(fileURLToPath(import.meta.url)), "fixtures/test.jpg"),
    );

    const bfileChooserPromise = page.waitForEvent("filechooser");
    await page.getByText("+ 身分証明書(裏)を選択").click();
    const backIdentificationfileChooser = await bfileChooserPromise;
    await backIdentificationfileChooser.setFiles(
      path.join(path.dirname(fileURLToPath(import.meta.url)), "fixtures/test.jpg"),
    );

    await page.getByRole("button", { name: "確認画面へ" }).click();
    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name test_first_name", { exact: true })).toBeVisible();
    await expect(page.getByText("test_supporter_5@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(0)).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(1)).toBeVisible();
    await page.getByRole("button", { name: "登録する" }).click();

    // NOTE: 登録に成功すると、サンクス画面に遷移する
    await expect(page.getByText("サポータ登録が完了しました。")).toBeVisible();
  });

  test("正常系_任意項目含む_入力に戻って修正", async ({ page }) => {
    await page.goto("/sign_up");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "姓" }).fill("test_last_name");
    await page.getByRole("textbox", { name: "名" }).fill("test_first_name");
    await page.getByRole("textbox", { name: "Email" }).fill("test_supporter_6@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");

    // NOTE: 身分証明書の表・裏のアップロード
    const fileChooserPromise = page.waitForEvent("filechooser");
    await page.getByText("+ 身分証明書(表)を選択").click();
    const frontIdentificationFileChooser = await fileChooserPromise;
    await frontIdentificationFileChooser.setFiles(
      path.join(path.dirname(fileURLToPath(import.meta.url)), "fixtures/test.jpg"),
    );

    const bfileChooserPromise = page.waitForEvent("filechooser");
    await page.getByText("+ 身分証明書(裏)を選択").click();
    const backIdentificationfileChooser = await bfileChooserPromise;
    await backIdentificationfileChooser.setFiles(
      path.join(path.dirname(fileURLToPath(import.meta.url)), "fixtures/test.jpg"),
    );

    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name test_first_name", { exact: true })).toBeVisible();
    await expect(page.getByText("test_supporter_6@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(0)).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(1)).toBeVisible();

    // NOTE: 入力画面で編集
    await page.getByRole("button", { name: "入力へ戻る" }).click();
    await page.getByText("× キャンセル").nth(0).click();
    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name test_first_name", { exact: true })).toBeVisible();
    await expect(page.getByText("test_supporter_6@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
    await expect(page.getByText("-")).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(0)).toBeVisible();
    await page.getByRole("button", { name: "登録する" }).click();

    // NOTE: 登録に成功すると、サンクス画面に遷移する
    await expect(page.getByText("サポータ登録が完了しました。")).toBeVisible();
  });

  test("異常系_任意項目含む", async ({ page }) => {
    await page.goto("/sign_up");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "姓" }).fill("test_last_name");
    await page.getByRole("textbox", { name: "名" }).fill("test_first_name");
    await page.getByRole("textbox", { name: "Email" }).fill("test_supporter_7@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("passwor");

    // NOTE: 身分証明書の表・裏のアップロード
    const fileChooserPromise = page.waitForEvent("filechooser");
    await page.getByText("+ 身分証明書(表)を選択").click();
    const frontIdentificationFileChooser = await fileChooserPromise;
    await frontIdentificationFileChooser.setFiles(
      path.join(path.dirname(fileURLToPath(import.meta.url)), "fixtures/test.jpg"),
    );

    const bfileChooserPromise = page.waitForEvent("filechooser");
    await page.getByText("+ 身分証明書(裏)を選択").click();
    const backIdentificationfileChooser = await bfileChooserPromise;
    await backIdentificationfileChooser.setFiles(
      path.join(path.dirname(fileURLToPath(import.meta.url)), "fixtures/test.jpg"),
    );

    await page.getByRole("button", { name: "確認画面へ" }).click();

    await expect(page.getByText("パスワードは8 ~ 24文字での入力をお願いします。")).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(0)).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(1)).toBeVisible();

    // 入力し直して確認 → 登録できること
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name test_first_name", { exact: true })).toBeVisible();
    await expect(page.getByText("test_supporter_7@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(0)).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(1)).toBeVisible();
    await page.getByRole("button", { name: "登録する" }).click();

    // NOTE: 登録に成功すると、サンクス画面に遷移する
    await expect(page.getByText("サポータ登録が完了しました。")).toBeVisible();
  });

  test("異常系_任意項目含む_入力に戻って修正", async ({ page }) => {
    await page.goto("/sign_up");
    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "姓" }).fill("test_last_name");
    await page.getByRole("textbox", { name: "名" }).fill("test_first_name");
    await page.getByRole("textbox", { name: "Email" }).fill("test_supporter_8@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");

    // NOTE: 身分証明書の表・裏のアップロード
    const fileChooserPromise = page.waitForEvent("filechooser");
    await page.getByText("+ 身分証明書(表)を選択").click();
    const frontIdentificationFileChooser = await fileChooserPromise;
    await frontIdentificationFileChooser.setFiles(
      path.join(path.dirname(fileURLToPath(import.meta.url)), "fixtures/test.jpg"),
    );

    const bfileChooserPromise = page.waitForEvent("filechooser");
    await page.getByText("+ 身分証明書(裏)を選択").click();
    const backIdentificationfileChooser = await bfileChooserPromise;
    await backIdentificationfileChooser.setFiles(
      path.join(path.dirname(fileURLToPath(import.meta.url)), "fixtures/test.jpg"),
    );

    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name test_first_name", { exact: true })).toBeVisible();
    await expect(page.getByText("test_supporter_8@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(0)).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(1)).toBeVisible();

    // NOTE: 入力画面に戻って異常系データを入力する
    await page.getByRole("button", { name: "入力へ戻る" }).click();
    await page.getByRole("textbox", { name: "パスワード" }).fill("passwor");
    await page.getByRole("button", { name: "確認画面へ" }).click();
    await expect(page.getByText("パスワードは8 ~ 24文字での入力をお願いします。")).toBeVisible();

    // NOTE: 画像が残っていること
    await expect(page.getByAltText("アップロード画像").nth(0)).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(1)).toBeVisible();

    // 入力し直して確認 → 登録できること
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByText("× キャンセル").nth(1).click();
    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name test_first_name", { exact: true })).toBeVisible();
    await expect(page.getByText("test_supporter_8@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
    await expect(page.getByAltText("アップロード画像")).toBeVisible();
    await expect(page.getByText("-")).toBeVisible();
    await page.getByRole("button", { name: "登録する" }).click();

    // NOTE: 登録に成功すると、サンクス画面に遷移する
    await expect(page.getByText("サポータ登録が完了しました。")).toBeVisible();
  });

  test("フォーム種別変更あり_入力に戻って修正", async ({ page }) => {
    await page.goto("/sign_up");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "姓" }).fill("test_last_name");
    await page.getByRole("textbox", { name: "名" }).fill("test_first_name");
    await page.getByRole("textbox", { name: "Email" }).fill("test_supporter_9@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");

    // NOTE: 身分証明書の表・裏のアップロード
    const fileChooserPromise = page.waitForEvent("filechooser");
    await page.getByText("+ 身分証明書(表)を選択").click();
    const frontIdentificationFileChooser = await fileChooserPromise;
    await frontIdentificationFileChooser.setFiles(
      path.join(path.dirname(fileURLToPath(import.meta.url)), "fixtures/test.jpg"),
    );

    const bfileChooserPromise = page.waitForEvent("filechooser");
    await page.getByText("+ 身分証明書(裏)を選択").click();
    const backIdentificationfileChooser = await bfileChooserPromise;
    await backIdentificationfileChooser.setFiles(
      path.join(path.dirname(fileURLToPath(import.meta.url)), "fixtures/test.jpg"),
    );

    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name")).toBeVisible();
    await expect(page.getByText("test_first_name")).toBeVisible();
    await expect(page.getByText("test_supporter_9@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(0)).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(1)).toBeVisible();

    // NOTE: 入力画面に戻り、フォーム種別を変更
    await page.getByRole("button", { name: "入力へ戻る" }).click();
    await page.getByRole("radio", { name: "企業" }).check();
    await expect(page.getByText("企業登録フォーム", { exact: true })).toBeVisible();

    // NOTE: もう一度サポータフォームに戻る
    await page.getByRole("radio", { name: "サポータ" }).check();
    await expect(page.getByRole("textbox", { name: "姓" })).toHaveValue("test_last_name");
    await expect(page.getByRole("textbox", { name: "名" })).toHaveValue("test_first_name");
    await expect(page.getByRole("textbox", { name: "Email" })).toHaveValue(
      "test_supporter_9@example.com",
    );
    await expect(page.getByRole("textbox", { name: "パスワード" })).toHaveValue("password");
    await expect(page.getByAltText("アップロード画像").nth(0)).toBeVisible();
    await expect(page.getByAltText("アップロード画像").nth(1)).toBeVisible();

    // NOTE: フォーム内容を入力し直しできること
    await page.getByRole("textbox", { name: "パスワード" }).fill("passwor");
    await page.getByText("× キャンセル").nth(0).click();
    await page.getByRole("button", { name: "確認画面へ" }).click();
    await expect(page.getByText("パスワードは8 ~ 24文字での入力をお願いします。")).toBeVisible();
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "確認画面へ" }).click();

    // NOTE: 確認画面で入力内容を表示できること
    await expect(page.getByText("test_last_name test_first_name", { exact: true })).toBeVisible();
    await expect(page.getByText("test_supporter_9@example.com", { exact: true })).toBeVisible();
    await expect(page.getByText("********", { exact: true })).toBeVisible();
    await expect(page.getByText("-")).toBeVisible();
    await expect(page.getByAltText("アップロード画像")).toBeVisible();
    await page.getByRole("button", { name: "登録する" }).click();

    // NOTE: 登録に成功すると、サンクス画面に遷移する
    await expect(page.getByText("サポータ登録が完了しました。")).toBeVisible();
  });
});

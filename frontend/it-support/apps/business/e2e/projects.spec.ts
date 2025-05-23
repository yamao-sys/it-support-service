import { test, expect } from "@playwright/test";

test.describe("/projects", () => {
  test.describe("GET /projects", () => {
    test("案件一覧ページ", async ({ page }) => {
      await page.goto("/sign_in");

      await page.getByRole("radio", { name: "企業" }).check();

      // NOTE: ログインフォームを入力
      await page.getByRole("textbox", { name: "Email" }).fill("test@example.com");
      await page.getByRole("textbox", { name: "パスワード" }).fill("password");
      await page.getByRole("button", { name: "ログイン" }).click();

      // NOTE: ログイン成功
      await page.waitForURL("/");

      // NOTE: 案件一覧ページへ遷移
      await page.goto("/projects");

      // NOTE: 初期表示が5件であることを確認
      const initialItemsCount = await page.getByRole("link", { name: "編集" }).count();
      expect(initialItemsCount).toBe(5);
      await expect(page.getByText("これ以上データはありません。")).not.toBeVisible();

      // NOTE: スクロールしてロードをトリガー
      await page.evaluate(() => {
        window.scrollTo(0, document.body.scrollHeight);
      });
      // TODO: 特定要素の出現を待つようにする
      await page.waitForTimeout(1500);

      expect(await page.getByRole("link", { name: "編集" }).count()).toBe(10);

      // NOTE: スクロールしてロードをトリガー
      await page.evaluate(() => {
        window.scrollTo(0, document.body.scrollHeight);
      });
      await expect(page.getByText("これ以上データはありません。")).toBeVisible();
    });
  });

  test.describe("POST /projects", () => {
    test("正常系(必須項目のみ)", async ({ page }) => {
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
      await page
        .locator(".react-datepicker__day--007:not(.react-datepicker__day--outside-month)")
        .click();
      await page.click('input[name="endDate"]');
      await page
        .locator(".react-datepicker__day--008:not(.react-datepicker__day--outside-month)")
        .click();
      await page.locator('label:has-text("公開フラグ")').click();

      await page.getByRole("button", { name: "保存する" }).click();

      page.on("dialog", async (dialog) => {
        expect(dialog.message()).toContain("案件を追加しました!");
        await dialog.accept();
      });
      await page.waitForURL("/");
      await expect(page).toHaveURL("/");
    });

    test("正常系(任意項目含む)", async ({ page }) => {
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
      await page
        .locator(".react-datepicker__day--007:not(.react-datepicker__day--outside-month)")
        .click();
      await page.click('input[name="endDate"]');
      await page
        .locator(".react-datepicker__day--008:not(.react-datepicker__day--outside-month)")
        .click();
      await page.getByRole("spinbutton", { name: "予算(下限)" }).fill("100000");
      await page.getByRole("spinbutton", { name: "予算(上限)" }).fill("100001");
      await page.locator('label:has-text("公開フラグ")').click();

      await page.getByRole("button", { name: "保存する" }).click();

      page.on("dialog", async (dialog) => {
        expect(dialog.message()).toContain("案件を追加しました!");
        await dialog.accept();
      });
      await page.waitForURL("/");
      await expect(page).toHaveURL("/");
    });

    test("異常系(必須項目のみ)", async ({ page }) => {
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
      await page
        .locator(".react-datepicker__day--008:not(.react-datepicker__day--outside-month)")
        .click();
      await page.click('input[name="endDate"]');
      await page
        .locator(".react-datepicker__day--007:not(.react-datepicker__day--outside-month)")
        .click();
      await page.locator('label:has-text("公開フラグ")').click();

      await page.getByRole("button", { name: "保存する" }).click();

      await expect(
        page.getByText("案件終了日と案件開始日の前後関係が不適です。", { exact: true }),
      ).toBeVisible();

      // NOTE: 正しく入力し直して保存
      await page.click('input[name="startDate"]');
      await page
        .locator(".react-datepicker__day--007:not(.react-datepicker__day--outside-month)")
        .click();
      await page.click('input[name="endDate"]');
      await page
        .locator(".react-datepicker__day--008:not(.react-datepicker__day--outside-month)")
        .click();

      await page.getByRole("button", { name: "保存する" }).click();

      page.on("dialog", async (dialog) => {
        expect(dialog.message()).toContain("案件を追加しました!");
        await dialog.accept();
      });
      await page.waitForURL("/");
      await expect(page).toHaveURL("/");
    });

    test("異常系(任意項目含む)", async ({ page }) => {
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
      await page
        .locator(".react-datepicker__day--007:not(.react-datepicker__day--outside-month)")
        .click();
      await page.click('input[name="endDate"]');
      await page
        .locator(".react-datepicker__day--008:not(.react-datepicker__day--outside-month)")
        .click();
      await page.getByRole("spinbutton", { name: "予算(下限)" }).fill("100001");
      await page.getByRole("spinbutton", { name: "予算(上限)" }).fill("100000");
      await page.locator('label:has-text("公開フラグ")').click();

      await page.getByRole("button", { name: "保存する" }).click();

      await expect(
        page.getByText("予算(上限)と予算(下限)の大小関係が不適です。", { exact: true }),
      ).toBeVisible();

      // NOTE: 正しく入力し直して保存
      await page.getByRole("spinbutton", { name: "予算(下限)" }).fill("100000");
      await page.getByRole("spinbutton", { name: "予算(上限)" }).fill("100001");

      await page.getByRole("button", { name: "保存する" }).click();

      page.on("dialog", async (dialog) => {
        expect(dialog.message()).toContain("案件を追加しました!");
        await dialog.accept();
      });
      await page.waitForURL("/");
      await expect(page).toHaveURL("/");
    });
  });

  test.describe("PUT /projects/{id}", () => {
    test("正常系(必須項目のみ)", async ({ page }) => {
      await page.goto("/sign_in");

      await page.getByRole("radio", { name: "企業" }).check();

      // NOTE: ログインフォームを入力
      await page.getByRole("textbox", { name: "Email" }).fill("test@example.com");
      await page.getByRole("textbox", { name: "パスワード" }).fill("password");
      await page.getByRole("button", { name: "ログイン" }).click();

      // NOTE: ログイン成功
      await page.waitForURL("/");

      // NOTE: 案件登録ページへ遷移
      await page.goto("/projects/1");

      // NOTE: 初期値がフォームに入っていることの確認
      expect(await page.getByRole("textbox", { name: "案件タイトル" }).inputValue()).toBeTruthy();
      expect(await page.getByRole("textbox", { name: "案件概要" }).inputValue()).toBeTruthy();
      expect(await page.locator('input[name="startDate"]').inputValue()).toBeTruthy();
      expect(await page.locator('input[name="endDate"]').inputValue()).toBeTruthy();
      expect(await page.getByRole("spinbutton", { name: "予算(下限)" }).inputValue()).toBeFalsy();
      expect(await page.getByRole("spinbutton", { name: "予算(上限)" }).inputValue()).toBeFalsy();
      await expect(page.locator('input[name="isActive"]')).toBeChecked();

      // NOTE: 登録内容の入力
      await page.getByRole("textbox", { name: "案件タイトル" }).fill("test title");
      await page.getByRole("textbox", { name: "案件概要" }).fill("test description\ntest");
      await page.click('input[name="startDate"]');
      await page
        .locator(".react-datepicker__day--007:not(.react-datepicker__day--outside-month)")
        .click();
      await page.click('input[name="endDate"]');
      await page
        .locator(".react-datepicker__day--008:not(.react-datepicker__day--outside-month)")
        .click();
      await page.locator('label:has-text("公開フラグ")').click();

      await page.getByRole("button", { name: "保存する" }).click();

      page.on("dialog", async (dialog) => {
        expect(dialog.message()).toContain("案件を更新しました!");
        await dialog.accept();
      });
      await page.waitForURL("/");
      await expect(page).toHaveURL("/");
    });

    test("正常系(任意項目含む)", async ({ page }) => {
      await page.goto("/sign_in");

      await page.getByRole("radio", { name: "企業" }).check();

      // NOTE: ログインフォームを入力
      await page.getByRole("textbox", { name: "Email" }).fill("test@example.com");
      await page.getByRole("textbox", { name: "パスワード" }).fill("password");
      await page.getByRole("button", { name: "ログイン" }).click();

      // NOTE: ログイン成功
      await page.waitForURL("/");

      // NOTE: 案件登録ページへ遷移
      await page.goto("/projects/2");

      // NOTE: 初期値がフォームに入っていることの確認
      expect(await page.getByRole("textbox", { name: "案件タイトル" }).inputValue()).toBeTruthy();
      expect(await page.getByRole("textbox", { name: "案件概要" }).inputValue()).toBeTruthy();
      expect(await page.locator('input[name="startDate"]').inputValue()).toBeTruthy();
      expect(await page.locator('input[name="endDate"]').inputValue()).toBeTruthy();
      expect(await page.getByRole("spinbutton", { name: "予算(下限)" }).inputValue()).toBeTruthy();
      expect(await page.getByRole("spinbutton", { name: "予算(上限)" }).inputValue()).toBeTruthy();
      await expect(page.locator('input[name="isActive"]')).toBeChecked();

      // NOTE: 登録内容の入力
      await page.getByRole("textbox", { name: "案件タイトル" }).fill("test title");
      await page.getByRole("textbox", { name: "案件概要" }).fill("test description\ntest");
      await page.click('input[name="startDate"]');
      await page
        .locator(".react-datepicker__day--007:not(.react-datepicker__day--outside-month)")
        .click();
      await page.click('input[name="endDate"]');
      await page
        .locator(".react-datepicker__day--008:not(.react-datepicker__day--outside-month)")
        .click();
      await page.locator('label:has-text("公開フラグ")').click();

      await page.getByRole("button", { name: "保存する" }).click();

      page.on("dialog", async (dialog) => {
        expect(dialog.message()).toContain("案件を更新しました!");
        await dialog.accept();
      });
      await page.waitForURL("/");
      await expect(page).toHaveURL("/");
    });

    test("異常系(必須項目のみ)", async ({ page }) => {
      await page.goto("/sign_in");

      await page.getByRole("radio", { name: "企業" }).check();

      // NOTE: ログインフォームを入力
      await page.getByRole("textbox", { name: "Email" }).fill("test@example.com");
      await page.getByRole("textbox", { name: "パスワード" }).fill("password");
      await page.getByRole("button", { name: "ログイン" }).click();

      // NOTE: ログイン成功
      await page.waitForURL("/");

      // NOTE: 案件登録ページへ遷移
      await page.goto("/projects/3");

      // NOTE: 初期値がフォームに入っていることの確認
      expect(await page.getByRole("textbox", { name: "案件タイトル" }).inputValue()).toBeTruthy();
      expect(await page.getByRole("textbox", { name: "案件概要" }).inputValue()).toBeTruthy();
      expect(await page.locator('input[name="startDate"]').inputValue()).toBeTruthy();
      expect(await page.locator('input[name="endDate"]').inputValue()).toBeTruthy();
      expect(await page.getByRole("spinbutton", { name: "予算(下限)" }).inputValue()).toBeFalsy();
      expect(await page.getByRole("spinbutton", { name: "予算(上限)" }).inputValue()).toBeFalsy();
      await expect(page.locator('input[name="isActive"]')).not.toBeChecked();

      // NOTE: 登録内容の入力
      await page.getByRole("textbox", { name: "案件タイトル" }).fill("test title");
      await page.getByRole("textbox", { name: "案件概要" }).fill("test description\ntest");
      await page.click('input[name="startDate"]');
      await page
        .locator(".react-datepicker__day--008:not(.react-datepicker__day--outside-month)")
        .click();
      await page.click('input[name="endDate"]');
      await page
        .locator(".react-datepicker__day--007:not(.react-datepicker__day--outside-month)")
        .click();
      await page.locator('label:has-text("公開フラグ")').click();

      await page.getByRole("button", { name: "保存する" }).click();

      // NOTE: バリデーションエラーメッセージが表示される
      await expect(
        page.getByText("案件終了日と案件開始日の前後関係が不適です。", { exact: true }),
      ).toBeVisible();

      // NOTE: 正しく入力し直して更新
      await page.click('input[name="startDate"]');
      await page
        .locator(".react-datepicker__day--007:not(.react-datepicker__day--outside-month)")
        .click();
      await page.click('input[name="endDate"]');
      await page
        .locator(".react-datepicker__day--008:not(.react-datepicker__day--outside-month)")
        .click();

      await page.getByRole("button", { name: "保存する" }).click();

      page.on("dialog", async (dialog) => {
        expect(dialog.message()).toContain("案件を更新しました!");
        await dialog.accept();
      });
      await page.waitForURL("/");
      await expect(page).toHaveURL("/");
    });

    test("異常系(任意項目含む)", async ({ page }) => {
      await page.goto("/sign_in");

      await page.getByRole("radio", { name: "企業" }).check();

      // NOTE: ログインフォームを入力
      await page.getByRole("textbox", { name: "Email" }).fill("test@example.com");
      await page.getByRole("textbox", { name: "パスワード" }).fill("password");
      await page.getByRole("button", { name: "ログイン" }).click();

      // NOTE: ログイン成功
      await page.waitForURL("/");

      // NOTE: 案件登録ページへ遷移
      await page.goto("/projects/4");

      // NOTE: 初期値がフォームに入っていることの確認
      expect(await page.getByRole("textbox", { name: "案件タイトル" }).inputValue()).toBeTruthy();
      expect(await page.getByRole("textbox", { name: "案件概要" }).inputValue()).toBeTruthy();
      expect(await page.locator('input[name="startDate"]').inputValue()).toBeTruthy();
      expect(await page.locator('input[name="endDate"]').inputValue()).toBeTruthy();
      expect(await page.getByRole("spinbutton", { name: "予算(下限)" }).inputValue()).toBeTruthy();
      expect(await page.getByRole("spinbutton", { name: "予算(上限)" }).inputValue()).toBeTruthy();
      await expect(page.locator('input[name="isActive"]')).not.toBeChecked();

      // NOTE: 登録内容の入力
      await page.getByRole("textbox", { name: "案件タイトル" }).fill("test title");
      await page.getByRole("textbox", { name: "案件概要" }).fill("test description\ntest");
      await page.click('input[name="startDate"]');
      await page
        .locator(".react-datepicker__day--008:not(.react-datepicker__day--outside-month)")
        .click();
      await page.click('input[name="endDate"]');
      await page
        .locator(".react-datepicker__day--007:not(.react-datepicker__day--outside-month)")
        .click();
      await page.getByRole("spinbutton", { name: "予算(下限)" }).fill("20001");
      await page.getByRole("spinbutton", { name: "予算(上限)" }).fill("20000");
      await page.locator('label:has-text("公開フラグ")').click();

      await page.getByRole("button", { name: "保存する" }).click();

      // NOTE: バリデーションエラーメッセージが表示される
      await expect(
        page.getByText("案件終了日と案件開始日の前後関係が不適です。", { exact: true }),
      ).toBeVisible();
      await expect(
        page.getByText("予算(上限)と予算(下限)の大小関係が不適です。", { exact: true }),
      ).toBeVisible();

      // NOTE: 正しく入力し直して更新
      await page.click('input[name="startDate"]');
      await page
        .locator(".react-datepicker__day--007:not(.react-datepicker__day--outside-month)")
        .click();
      await page.click('input[name="endDate"]');
      await page
        .locator(".react-datepicker__day--008:not(.react-datepicker__day--outside-month)")
        .click();
      await page.getByRole("spinbutton", { name: "予算(下限)" }).fill("20000");
      await page.getByRole("spinbutton", { name: "予算(上限)" }).fill("20001");

      await page.getByRole("button", { name: "保存する" }).click();

      page.on("dialog", async (dialog) => {
        expect(dialog.message()).toContain("案件を更新しました!");
        await dialog.accept();
      });
      await page.waitForURL("/");
      await expect(page).toHaveURL("/");
    });
  });
});

import { test, expect } from "@playwright/test";

test.describe("/to_projects", () => {
  test.describe("GET /to_projects", () => {
    test("案件一覧ページ", async ({ page }) => {
      await page.goto("/sign_in");

      await page.getByRole("radio", { name: "サポータ" }).check();

      // NOTE: ログインフォームを入力
      await page.getByRole("textbox", { name: "Email" }).fill("test@example.com");
      await page.getByRole("textbox", { name: "パスワード" }).fill("password");
      await page.getByRole("button", { name: "ログイン" }).click();

      // NOTE: ログイン成功
      await page.waitForURL("/");

      // NOTE: 案件一覧ページへ遷移
      await page.goto("/to_projects");

      // NOTE: 初期表示が5件であることを確認
      const initialItemsCount = await page.getByRole("link", { name: "詳細を見る" }).count();
      expect(initialItemsCount).toBe(5);
      await expect(page.getByText("これ以上データはありません。")).not.toBeVisible();

      // NOTE: スクロールしてロードをトリガー
      await page.evaluate(() => {
        window.scrollTo(0, document.body.scrollHeight);
      });
      // TODO: 特定要素の出現を待つようにする
      await page.waitForTimeout(1500);

      expect(await page.getByRole("link", { name: "詳細を見る" }).count()).toBe(10);

      // NOTE: スクロールしてロードをトリガー
      await page.evaluate(() => {
        window.scrollTo(0, document.body.scrollHeight);
      });
      await expect(page.getByText("これ以上データはありません。")).toBeVisible();
    });
  });
});

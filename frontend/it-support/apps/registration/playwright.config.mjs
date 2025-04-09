import { defineConfig } from "@playwright/test";
import * as pkg from "../../packages/playwright-config/playwright.config.base.mjs";
const { baseConfig } = pkg;

export default defineConfig({
  ...baseConfig,
  testDir: "./e2e",
  use: {
    ...baseConfig.use,
    baseURL: "http://localhost:3100",
  },
  webServer: {
    timeout: 2 * 60 * 1000,
    command: "pnpm dev:test",
    url: "http://localhost:3100",
    reuseExistingServer: false,
  },
});

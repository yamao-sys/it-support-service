import { defineConfig } from "@playwright/test";
import * as pkg from "../../packages/playwright-config/playwright.config.base.mjs";
const { baseConfig } = pkg;

export default defineConfig({
  ...baseConfig,
  testDir: "./e2e",
  use: {
    ...baseConfig.use,
    baseURL: "http://localhost:3001",
  },
  webServer: {
    command: "pnpm dev",
    url: "http://localhost:3001",
    reuseExistingServer: true,
  },
});

import * as tailwindConfig from "@repo/ui/tailwind.config";

module.exports = {
  preset: [tailwindConfig],
  content: ["./app/**/*.tsx", "./components/**/*.tsx", "../../packages/ui/src/**/*.tsx"],
  theme: {
    extend: {},
  },
  plugins: [],
};

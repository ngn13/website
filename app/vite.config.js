import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";
import { fileURLToPath } from "url";
import { readFileSync } from "fs";

const default_env = {
  REPORT_URL: "https://github.com/ngn13/website/issues",
  SOURCE_URL: "https://github.com/ngn13/website",
  APP_URL: "http://localhost:7001",
  API_URL: "http://localhost:7002",
  DOC_URL: "http://localhost:7003",
};

const file = fileURLToPath(new URL("package.json", import.meta.url));
const json = readFileSync(file, "utf8");
const pkg = JSON.parse(json);

for (let env in default_env) {
  if (process.env["WEBSITE_" + env] === undefined) process.env["WEBSITE_" + env] = default_env[env];
}

export default defineConfig({
  plugins: [sveltekit()],
  envPrefix: "WEBSITE",
  preview: {
    port: 7001,
    strictPort: true,
  },
  server: {
    port: 7001,
    strictPort: true,
  },
  define: {
    pkg: pkg,
  },
});

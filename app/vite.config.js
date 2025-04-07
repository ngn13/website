import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";
import { fileURLToPath } from "url";
import { readFileSync } from "fs";

function env_from(prefix, object) {
  for (const [key, value] of Object.entries(object)) {
    let type = typeof value;
    let name = prefix + "_" + key.toUpperCase();

    switch (type) {
      case "object":
        env_from(name, value);
        break;

      case "string":
        if (process.env[name] === undefined) process.env[name] = value;
        break;
    }
  }
}

const default_env = {
  source_url: "https://git.ngn.tf/ngn/website",
  report_url: "https://git.ngn.tf/ngn/website/issues",
  doc_url: "http://localhost:7003",
  api: {
    url: "http://localhost:7002",
    path: "http://localhost:7002",
  },
};

const package_file = fileURLToPath(new URL("package.json", import.meta.url));
const package_json = readFileSync(package_file, "utf8");
const package_data = JSON.parse(package_json);

env_from("WEBSITE", default_env);

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
    pkg: package_data,
  },
});

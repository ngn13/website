//import adapter from '@sveltejs/adapter-auto';
import adapter from "@sveltejs/adapter-node";

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    adapter: adapter(),
  },
  onwarn: (warning, handler) => {
    if (warning.code === "a11y-click-events-have-key-events") return;
    handler(warning);
  },
};

export default config;

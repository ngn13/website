import { init, register, waitLocale } from "svelte-i18n";
import { browser_lang } from "$lib/util.js";
import { services } from "$lib/api.js";
import languages from "$lib/lang.js";

// setup the locale
for (let i = 0; i < languages.length; i++)
  register(languages[i].code, () => import(/* @vite-ignore */ languages[i].path));

init({
  fallbackLocale: languages[0].code,
  initialLocale: browser_lang(),
});

// load locales & load data from the API
export async function load({ fetch }) {
  await waitLocale();

  try {
    return {
      services: await services(fetch),
      error: null,
    };
  } catch (err) {
    return {
      error: err,
    };
  }
}

import { default_language, language, set_lang } from "$lib/util.js";
import { get_services, get_projects } from "$lib/api.js";
import languages from "$lib/lang.js";

import { init, register, waitLocale } from "svelte-i18n";
import { get } from "svelte/store";

// setup the locale
for (let i = 0; i < languages.length; i++)
  register(languages[i].code, () => import(/* @vite-ignore */ languages[i].path));

// set the language
set_lang();

init({
  fallbackLocale: default_language,
  initialLocale: get(language),
});

// load locales & load data from the API
export async function load({ fetch }) {
  await waitLocale();

  try {
    return {
      services: await get_services(fetch),
      projects: await get_projects(fetch),
      error: null,
    };
  } catch (err) {
    return {
      error: err,
    };
  }
}

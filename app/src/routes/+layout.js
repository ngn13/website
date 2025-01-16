import { default_language, language, set_lang } from "$lib/util.js";
import { api_get_services, api_get_projects } from "$lib/api.js";
import { doc_get_list } from "$lib/doc.js";
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
      services: await api_get_services(fetch),
      projects: await api_get_projects(fetch),
      docs: await doc_get_list(fetch),
      error: null,
    };
  } catch (err) {
    return {
      error: err,
    };
  }
}

import { locale, waitLocale } from "svelte-i18n";
import { init, register } from "svelte-i18n";
import { browser } from "$app/environment";
import languages from "$lib/lang.js";

const defaultLocale = languages[0].code;

for (let i = 0; i < languages.length; i++)
  register(languages[i].code, () => import(/* @vite-ignore */ languages[i].path));

init({
  fallbackLocale: defaultLocale,
  initialLocale: browser ? window.navigator.language.slice(0, 2).toLowerCase() : defaultLocale,
});

export const load = async () => {
  if (browser) locale.set(window.navigator.language);
  await waitLocale();
};

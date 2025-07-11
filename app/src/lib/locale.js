import { init, locale, register, waitLocale } from "svelte-i18n";
import { browser } from "$app/environment";
import { get, writable } from "svelte/store";

const locale_default = "en";
let locale_index = writable(0);
let locale_list = [];

function locale_setup() {
  // english
  register("en", () => import("../locales/en.json"));
  locale_list.push({ code: "en", name: "English", icon: "🇬🇧" });

  // turkish
  register("tr", () => import("../locales/tr.json"));
  locale_list.push({ code: "tr", name: "Turkish", icon: "🇹🇷" });

  init({
    fallbackLocale: locale_default,
    initialLocale: get(locale),
  });
}

function locale_from_browser() {
  if (browser) return window.navigator.language.slice(0, 2).toLowerCase();
  else return locale_default;
}

function locale_select(l = null) {
  if (l === null) {
    if (browser && null !== (l = localStorage.getItem("locale")))
      locale_select(l);
    else locale_select(locale_from_browser());
    return;
  }

  l = l.slice(0, 2);

  for (let i = 0; i < locale_list.length; i++) {
    if (l !== locale_list[i].code) continue;

    if (browser) localStorage.setItem("locale", l);

    locale.set(l);
    locale_index.set(i);

    return;
  }

  locale.set(locale_default);
  locale_index.set(0);
}

async function locale_wait() {
  await waitLocale();
}

export {
  locale,
  locale_list,
  locale_index,
  locale_default,
  locale_setup,
  locale_wait,
  locale_select,
  locale_from_browser,
};

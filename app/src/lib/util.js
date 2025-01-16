import { browser } from "$app/environment";
import { locale } from "svelte-i18n";
import languages from "$lib/lang.js";
import { writable, get } from "svelte/store";

const default_language = languages[0].code;
const colors = [
  "yellow",
  "cyan",
  "green",
  "pinkish",
  "red",
  //  "blue" (looks kinda ass)
];

let language = writable(default_language);
let colors_pos = -1;

function browser_lang() {
  if (browser) return window.navigator.language.slice(0, 2).toLowerCase();
  else return get(language);
}

function set_lang(lang) {
  language.set(default_language);
  locale.set(default_language);

  if (lang === null || lang === undefined) {
    if (browser && null !== (lang = localStorage.getItem("language"))) set_lang(lang);
    else if (browser) set_lang(browser_lang());
    return;
  }

  lang = lang.slice(0, 2);

  for (let i = 0; i < languages.length; i++) {
    if (lang !== languages[i].code) continue;

    language.set(lang);
    locale.set(lang);

    if (browser) localStorage.setItem("language", lang);
  }
}

function urljoin(url, path = null, query = {}) {
  let url_len = url.length;

  if (url[url_len - 1] != "/") url += "/";

  if (null === path || "" === path) url = new URL(url);
  else if (path[0] === "/") url = new URL(path.slice(1), url);
  else url = new URL(path, url);

  for (let k in query) url.searchParams.append(k, query[k]);

  return url.href;
}

function frontend_url(path = null, query = {}) {
  return urljoin(import.meta.env.APP_URL, path, query);
}

function color() {
  if (colors_pos < 0) colors_pos = Math.floor(Math.random() * colors.length);
  else if (colors_pos >= colors.length) colors_pos = 0;

  return colors[colors_pos];
}

function click() {
  let audio = new Audio("/click.wav");
  audio.play();
}

function time_from_ts(ts) {
  if (ts === 0 || ts === undefined) return;

  let ts_date = new Date(ts * 1000);
  let ts_zone = ts_date.toString().match(/([A-Z]+[\+-][0-9]+)/)[1];

  return (
    new Intl.DateTimeFormat(browser_lang(), {
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
    }).format(ts_date) + ` (${ts_zone})`
  );
}

function date_from_ts(ts) {
  if (ts === 0 || ts === undefined) return;

  return new Intl.DateTimeFormat(browser_lang(), {
    month: "2-digit",
    year: "2-digit",
    day: "2-digit",
  }).format(new Date(ts * 1000));
}

export {
  default_language,
  browser_lang,
  language,
  set_lang,
  urljoin,
  frontend_url,
  click,
  color,
  time_from_ts,
  date_from_ts,
};

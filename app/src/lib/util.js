import { browser } from "$app/environment";

const default_lang = "en";
const colors = [
  "yellow",
  "cyan",
  "green",
  "pinkish",
  "red",
  //  "blue" (looks kinda ass)
];

let colors_pos = -1;

function urljoin(url, path = null, query = {}) {
  let url_len = url.length;

  if (url[url_len - 1] != "/") url += "/";

  if (null === path || "" === path) url = new URL(url);
  else if (path[0] === "/") url = new URL(path.slice(1), url);
  else url = new URL(path, url);

  for (let k in query) url.searchParams.append(query[k]);

  return url.href;
}

function frontend_url(path = null, query = {}) {
  return urljoin(import.meta.env.VITE_FRONTEND_URL, path, query);
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

function browser_lang() {
  if (browser) return window.navigator.language.slice(0, 2).toLowerCase();
  return default_lang;
}

function time_from_ts(ts) {
  return new Date(ts * 1000).toLocaleTimeString();
}

export { urljoin, frontend_url, browser_lang, click, color, time_from_ts };

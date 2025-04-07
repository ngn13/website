import { locale_from_browser } from "$lib/locale.js";

const colors = [
  "yellow",
  "cyan",
  "green",
  "pinkish",
  "red",
  //  "blue" (looks kinda ass)
];

let colors_pos = -1;

function color() {
  if (colors_pos < 0) colors_pos = Math.floor(Math.random() * colors.length);
  else if (colors_pos >= colors.length) colors_pos = 0;

  return colors[colors_pos];
}

function click() {
  let audio = new Audio("/assets/click.wav");
  audio.play();
}

function urljoin(url, path = null) {
  if (undefined === url || null === url) return;

  if (url[url.length - 1] != "/") url += "/";

  if (null === path || "" === path) return url;
  if (path[0] === "/") return url + path.slice(1);
  return url + path;
}

function time_from_ts(ts) {
  if (ts === 0 || ts === undefined) return;

  let ts_date = new Date(ts * 1000);
  let ts_zone = ts_date.toString().match(/([A-Z]+[\+-][0-9]+)/)[1];

  return (
    new Intl.DateTimeFormat(locale_from_browser(), {
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
    }).format(ts_date) + ` (${ts_zone})`
  );
}

function date_from_ts(ts) {
  if (ts === 0 || ts === undefined) return;

  return new Intl.DateTimeFormat(locale_from_browser(), {
    month: "2-digit",
    year: "2-digit",
    day: "2-digit",
  }).format(new Date(ts * 1000));
}

export { color, click, urljoin, time_from_ts, date_from_ts };

import { join } from "$lib/api.js";

const colors = [
  "yellow",
  "cyan",
  "green",
  "pinkish",
  "red",
  //  "blue" (looks kinda ass)
];
let colors_pos = -1;
let api_url = join;

function color() {
  if (colors_pos < 0) colors_pos = Math.floor(Math.random() * colors.length);
  else if (colors_pos >= colors.length) colors_pos = 0;

  return colors[colors_pos];
}

function click() {
  let audio = new Audio("/click.wav");
  audio.play();
}

function frontend_url(path) {
  if (null !== path && path !== "") return new URL(path, import.meta.env.VITE_FRONTEND_URL).href;
  else return new URL(import.meta.env.VITE_FRONTEND_URL).href;
}

function time_from_ts(ts) {
  return new Date(ts * 1000).toLocaleTimeString();
}

export { api_url, frontend_url, click, color, time_from_ts };

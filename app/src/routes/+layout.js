import { locale_setup, locale_wait } from "$lib/locale.js";

export async function load() {
  locale_setup();
  await locale_wait();
}

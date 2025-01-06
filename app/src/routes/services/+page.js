import { services } from "$lib/api.js";

export async function load({ fetch }) {
  return {
    list: await services(fetch),
  };
}

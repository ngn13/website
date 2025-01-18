import { api_get_services } from "$lib/api.js";

export async function load({ fetch }) {
  try {
    let services = await api_get_services(fetch)
    return {
      services: null === services ? [] : services,
    };
  } catch (err) {
    return {
      error: err.toString(),
    };
  }
}

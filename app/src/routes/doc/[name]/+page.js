import { goto } from "$app/navigation";
import { doc_get } from "$lib/doc";

export async function load({ fetch, params }) {
  try {
    return await doc_get(fetch, params.name);
  } catch (err) {
    if (err.toString().includes("not found")) return goto("/");
    return { error: err };
  }
}

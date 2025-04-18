import { doc_get } from "$lib/doc";

export async function load({ fetch, params }) {
  try {
    return {
      doc: await doc_get(fetch, params.name),
      error: "",
    };
  } catch (err) {
    return { error: err.toString() };
  }
}

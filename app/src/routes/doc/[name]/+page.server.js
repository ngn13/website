import { doc_get_list, doc_get } from "$lib/doc";

export async function load({ fetch, params }) {
  try {
    return {
      docs: await doc_get_list(fetch),
      doc: await doc_get(fetch, params.name),
    };
  } catch (err) {
    return { error: err.toString() };
  }
}

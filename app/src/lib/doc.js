import { urljoin } from "$lib/util.js";

function doc_urljoin(path = null, query = {}) {
  return urljoin(import.meta.env.WEBSITE_DOC_URL, path, query);
}

function doc_check_err(json) {
  if ("error" in json) throw new Error(`Documentation server returned an error: ${json["error"]}`);
}

async function doc_http_get(fetch, url) {
  const res = await fetch(url);
  const json = await res.json();
  doc_check_err(json);
  return json;
}

async function doc_get_list(fetch) {
  return await doc_http_get(fetch, doc_urljoin("/list"));
}

async function doc_get(fetch, name) {
  let url = doc_urljoin("/get");
  url = urljoin(url, name);
  return await doc_http_get(fetch, url);
}

export { doc_urljoin, doc_get, doc_get_list };

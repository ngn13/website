import { urljoin } from "$lib/util.js";

const version = "v1";
const url = urljoin(import.meta.env.VITE_API_URL, version);

function api_url(path = null, query = {}) {
  return urljoin(url, path, query);
}

function check_err(json) {
  if (!("error" in json)) throw new Error('API response is missing the "error" key');

  if (json["error"] != "") throw new Error(`API returned an error: ${json["error"]}`);

  if (!("result" in json)) throw new Error('API response is missing the "result" key');
}

async function GET(fetch, url) {
  const res = await fetch(url);
  const json = await res.json();
  check_err(json);
  return json["result"];
}

async function visitor(fetch) {
  return GET(fetch, api_url("/visitor"));
}

async function services(fetch) {
  return GET(fetch, api_url("/services"));
}

export { version, api_url, visitor, services };

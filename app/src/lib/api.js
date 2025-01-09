import { urljoin } from "$lib/util.js";

const version = "v1";
const url = urljoin(import.meta.env.APP_API_URL, version);

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

async function get_metrics(fetch) {
  return GET(fetch, api_url("/metrics"));
}

async function get_services(fetch) {
  return GET(fetch, api_url("/services"));
}

async function get_projects(fetch) {
  return GET(fetch, api_url("/projects"));
}

export { version, api_url, get_metrics, get_services, get_projects };

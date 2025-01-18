import { urljoin } from "$lib/util.js";

const api_version = "v1";
const api_url = urljoin(import.meta.env.WEBSITE_API_URL, api_version);

function api_urljoin(path = null, query = {}) {
  return urljoin(api_url, path, query);
}

function api_check_err(json) {
  if (!("error" in json)) throw new Error('API response is missing the "error" key');

  if (json["error"] != "") throw new Error(`API returned an error: ${json["error"]}`);

  if (!("result" in json)) throw new Error('API response is missing the "result" key');
}

async function api_http_get(fetch, url) {
  const res = await fetch(url);
  const json = await res.json();
  api_check_err(json);
  return json["result"];
}

async function api_get_metrics(fetch) {
  return await api_http_get(fetch, api_urljoin("/metrics"));
}

async function api_get_services(fetch) {
  return await api_http_get(fetch, api_urljoin("/services"));
}

async function api_get_projects(fetch) {
  return await api_http_get(fetch, api_urljoin("/projects"));
}

export { api_version, api_urljoin, api_get_metrics, api_get_services, api_get_projects };

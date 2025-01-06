const version = "v1";
const url = new URL(version + "/", import.meta.env.VITE_API_URL).href;

function join(path) {
  if (null === path || path === "") return url;

  if (path[0] === "/") path = path.slice(1);

  return new URL(path, url).href;
}

async function services(fetch) {
  const res = await fetch(join("/services"));
  const json = await res.json();

  if (!("result" in json)) return [];

  return json.result;
}

export { version, join, services };

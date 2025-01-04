export async function load({ fetch }) {
  const api = import.meta.env.VITE_API_URL_DEV;
  const res = await fetch(api + "/blog/sum");
  const data = await res.json();

  return {
    posts: data["result"],
  };
}

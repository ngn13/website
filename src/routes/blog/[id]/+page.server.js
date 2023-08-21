export async function load({ fetch, params }) {
  const id = params.id 
  const api = import.meta.env.VITE_API_URL_DEV
  const res = await fetch(api+"/blog/get?id="+id) 
  const data = await res.json()
 
  if (data["error"] != "") {
    return {
      error: data["error"]
    }
  }

  return data["result"] 
}

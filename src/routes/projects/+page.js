export async function load({ fetch }) {
  const api = import.meta.env.VITE_API_URL_DEV
  const res = await fetch(api+"/projects/get")
  const data = await res.json()

  if (data["error"] != ""){
    return {
      error: data["error"]
    }
  }

  let all = data["result"]
  let counter = 0
  let currentlist = []
  let projects = []  

  for (let i = 0; i < all.length; i++){
    currentlist.push(all[i])
    counter += 1

    if(i == all.length-1 && counter != 3){
      projects.push(currentlist)
    }

    if (counter == 3) {
      projects.push(currentlist)
      currentlist = []
      counter = 0
    }
  }

  return {
    projects
  }
}

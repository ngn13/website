export async function load({ fetch }) {
  const api = import.meta.env.VITE_API_URL_DEV
  const res = await fetch(api+"/services/all")
  const data = await res.json()

  if (data["error"] != ""){
    return {
      error: data["error"]
    }
  }

  // Some really bad code to convert 
  // [service1, service2, service3...] 

  // to 

  // [[service1, service2], [service4, service5], [service4...]...]
  // so i can render it in the UI easily

  let all = data["result"]
  let counter = 0
  let currentlist = []  
  let services = []  

  for (let i = 0; i < all.length; i++){
    currentlist.push(all[i])
    counter += 1

    if(i == all.length-1 && counter != 2){
      services.push(currentlist)
    }

    if (counter == 2) {
      services.push(currentlist)
      currentlist = []
      counter = 0
    }
  }

  return {
    services 
  }
}

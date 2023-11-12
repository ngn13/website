<script>
  import Header from "../../../lib/header.svelte"
  import { goto } from "$app/navigation"
  import { onMount } from "svelte"
  import DOMPurify from "dompurify"
  import { marked } from "marked"

  export let data 
  let sanitized
  const api = import.meta.env.VITE_API_URL_DEV

  let upvote_status = "inactive"
  let downvote_status = "inactive"
  let voted = false
  let audio

  async function get_status() {
    const res = await fetch(api+"/blog/vote/status?id="+data.id)
    const json = await res.json()
    
    if(json["error"]!= ""){
      return
    }

    if (json["result"] == "upvote") {
      upvote_status = "active"
      downvote_status = "inactive"
    }
    else {
      downvote_status = "active"
      upvote_status = "inactive"
    }
    
    voted = true
  }

  onMount(async ()=>{
    if (data.title == undefined) 
      goto("/blog")

    sanitized = DOMPurify.sanitize(
      marked.parse(data.content, { breaks: true }),
      {
        ADD_TAGS: ["iframe"],
        ADD_ATTR: ["allow", "allowfullscreen", "frameborder", "scrolling"]
      }
    )
    
    await get_status()
  })

  async function upvote(){
    audio.play()
    const res = await fetch(api+"/blog/vote/set?id="+data.id+"&to=upvote")
    const json = await res.json()

    if(json["error"] != ""){
      return
    }

    if (voted){
      data.vote += 2
    } 
    else {
      voted = true
      data.vote += 1
    }

    await get_status()
  }

  async function downvote(){
    audio.play()
    const res = await fetch(api+"/blog/vote/set?id="+data.id+"&to=downvote")
    const json = await res.json()

    if(json["error"] != ""){
      return
    }

    if (voted){
      data.vote -= 2
    } 
    else {
      voted = true
      data.vote -= 1
    }

    await get_status()
  }
</script>

<svelte:head>
  <title>[ngn] | {data.title}</title> 
  <meta content="[ngn] | blog" property="og:title" />
  <meta content="{data.content.substring(0, 100)}..." property="og:description" />
  <meta content="https://ngn13.fun" property="og:url" />
  <meta content="#000000" data-react-helmet="true" name="theme-color" />
</svelte:head>

<Header>
  <c>{data.title}</c>
  <p>{data.author} | {data.date}</p>
</Header>

<main>
  <audio bind:this={audio} preload="auto">
    <source src="/click.wav" type="audio/mpeg" />
  </audio>
  <div class="content markdown-body">
    {@html sanitized}
  </div>
  <div class="votes">
    <h3 on:click={async ()=>{upvote()}} class="{upvote_status}">󰜷</h3>
    <p>{data.vote}</p>
    <h3 on:click={async ()=>{downvote()}} class="{downvote_status}">󰜮</h3>
  </div>
</main>

<style>
p {
  font-size: 30px;
}

main {
  padding: 50px;
  color: white;
  display: flex;
  flex-direction: row;
  gap: 20px;
  justify-content: center;
}

.content {
  max-width: 80%;
  padding: 40px;
  background: var(--dark-two);
  box-shadow: var(--box-shadow);
  border-radius: 7px;
}

.votes {
  display: flex;
  flex-direction: column;
  text-align: center;
  text-shadow: var(--text-shadow);
}

.votes h3{
  font-size: 40px;
  cursor: pointer;
}

.votes h3:hover {
  animation-name: colorAnimation;
  animation-iteration-count: infinite;
  animation-duration: 10s;
}

.active {
  animation-name: colorAnimation;
  animation-iteration-count: infinite;
  animation-duration: 10s;
}
</style>

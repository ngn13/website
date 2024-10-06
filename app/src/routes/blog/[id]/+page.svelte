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
    const res = await fetch(api+"/blog/vote/get?id="+data.id)
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
  <title>[ngn.tf] | {data.title}</title> 
  <meta content="[ngn] | blog" property="og:title" />
  <meta content="{data.content.substring(0, 100)}..." property="og:description" />
  <meta content="https://ngn.tf" property="og:url" />
  <meta content="#000000" data-react-helmet="true" name="theme-color" />
  <link href="/markdown.css" rel="stylesheet">
</svelte:head>

<Header subtitle="{data.author} | {data.date}">
  {data.title}
</Header>

<main>
  <audio bind:this={audio} preload="auto">
    <source src="/click.wav" type="audio/mpeg" />
  </audio>
  <div class="content markdown-body">
    {@html sanitized}
  </div>
  <div class="votes">
    <button on:click={async ()=>{upvote()}} class="{upvote_status}">
      <i class="nf nf-md-arrow_up_bold"></i>
    </button>
    <p>{data.vote}</p>
    <button on:click={async ()=>{downvote()}} class="{downvote_status}">
      <i class="nf nf-md-arrow_down_bold"></i>
    </button>
  </div>
</main>


<style>
main {
  padding: 50px 10% 50px 10%;
  color: white;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: start;
}

@media only screen and (max-width: 816px) {
  main {
    padding: 50px 20% 50px 20%;
  }
}

.content {
  padding: 30px;
  background: var(--dark-four);
  border-radius: var(--radius);
  border: solid 1px var(--border-color);
  box-shadow: var(--box-shadow);
  width: auto;
  width: 100%;
}

.votes {
  display: flex;
  flex-direction: column;
  text-align: center;
  text-shadow: var(--text-shadow);
  gap: 10px;
  padding: 15px 5px 15px 5px;
  margin-left: 10px;
}

.votes p {
  font-size: 25px;
  color: var(--dark-six);
}

.votes button{
  display: flex; 
  flex-direction: row;
  gap: 10px;
  background: none;
  outline: none;
  border: none;
  font-size: 30px;
  cursor: pointer;
  color: var(--dark-six);
}

.votes button:hover {
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

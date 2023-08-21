<script>
  import Header from "../../lib/header.svelte";
  import CardLink from "../../lib/card_link.svelte";
  import { onMount } from "svelte";

  export let data
  let inpt 

  let all = data.posts
  let posts = all
  let inputcls = "c"

  function show(term){
    posts = []
    for(let i = 0; i < all.length; i++) {
      if (all[i].title.toLowerCase().includes(term.toLowerCase())){
        posts.push(all[i])
      }
    }

    if (posts.length == 0){
      inputcls = "nf"
      return
    }
    inputcls = "c"
  }

  onMount(() => {
    inpt.focus()
    show(inpt.value)
  })

  function search(){
    let term = inpt.value
    show(term)
  }
</script>

<svelte:head>
  <title>[ngn] | blog</title> 
  <meta content="[ngn] | blog" property="og:title" />
  <meta content="View my blog posts" property="og:description" />
  <meta content="https://ngn13.fun" property="og:url" />
  <meta content="#000000" data-react-helmet="true" name="theme-color" />
</svelte:head>

<Header>
  <c>/dev/</c>blog
</Header>

<main>
  <input on:keyup={search} bind:this={inpt} class="{inputcls}" placeholder="Search for a post">
  {#each posts as post} 
    <CardLink url="/blog/{post.id}"  title="{post.title}">
      <p>{post.author} | {post.date}</p>
      <br>
      {post.content}...
    </CardLink>
  {/each}
</main>

<style>
main{
  display: flex;
  flex-direction: column;
  gap: 35px;
  padding: 15%;
  padding-top: 50px;
}

input {
  text-align: center;
  background: var(--dark-two);
  background: none;
  border: none;
  outline: none;
  font-size: 40px;
}

p {
  font-size: 20px;
}

.nf {
  color: rgb(120, 120, 120);
}
</style>

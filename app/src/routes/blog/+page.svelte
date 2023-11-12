<script>
  import Header from "../../lib/header.svelte";
  import CardLink from "../../lib/card_link.svelte";

  export let data
  let posts = data.posts 
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

@media only screen and (max-width: 1316px) {
  main {
    padding: 50px;
  }
}
</style>

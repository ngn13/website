<script>
  import Header from "../../lib/header.svelte";
  import CardLink from "../../lib/card_link.svelte";

  export let data
  let posts = data.posts 
</script>

<svelte:head>
  <title>[ngn.tf] | blog</title> 
  <meta content="[ngn] | blog" property="og:title" />
  <meta content="View my blog posts" property="og:description" />
  <meta content="https://ngn.tf" property="og:url" />
  <meta content="#000000" data-react-helmet="true" name="theme-color" />
  <link rel="alternate" type="application/atom+xml" href="{import.meta.env.VITE_API_URL_DEV+'/blog/feed.atom'}" title="Atom Feed">
  <link rel="alternate" type="application/rss+xml" href="{import.meta.env.VITE_API_URL_DEV+'/blog/feed.rss'}" title="RSS Feed">
</svelte:head>

<Header>
  <c>/dev/</c>blog
</Header>

<main>
  <div class="feed-list">
    <a href="{import.meta.env.VITE_API_URL_DEV+'/blog/feed.rss'}">
      <c><i class="nf nf-fa-rss_square"></i></c> <p>RSS</p>
    </a>
    <a href="{import.meta.env.VITE_API_URL_DEV+'/blog/feed.atom'}">
      <c><i class="nf nf-fae-atom"></i></c> <p>Atom</p>
    </a>
    <a href="{import.meta.env.VITE_API_URL_DEV+'/blog/feed.json'}">
      <c><i class="nf nf-seti-json"></i></c> <p>JSON</p>
    </a>
  </div>
  <div class="post-list">
    {#each posts as post} 
      <CardLink url="/blog/{post.id}"  title="{post.title}">
        <p>{post.author} | {post.date}</p>
        <br>
        {post.content}...
      </CardLink>
    {/each}
  </div>
</main>

<style>
.post-list{
  display: flex;
  flex-direction: column;
  gap: 28px;
}

main {
  padding: 15%;
  padding-top: 50px;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

p {
  font-size: 20px;
}

.feed-list{
  text-align: right;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  gap: 10px;
}

.feed-list a {
  text-decoration: none;
  padding: 10px 15px 10px 15px;
  background: var(--dark-three);
  border-radius: var(--radius);
  border: solid 1px var(--border-color);
  color: var(--white);
  font-size: 20px;
  font-weight: 900;

  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 7px;

  transition: .2s;
}

.feed-list a:hover {
  box-shadow: var(--box-shadow);
}

.feed-list a i{
  font-size: 21px;
}

@media only screen and (max-width: 1316px) {
  main {
    padding: 50px;
  }
}
</style>

<script>
  import Header from "../lib/header.svelte";
  import Card from "../lib/card.svelte";
</script>

<svelte:head>
  <title>[ngn] | homepage</title> 
  <meta content="[ngn] | Homepage" property="og:title" />
  <meta content="Homepage of my personal website" property="og:description" />
  <meta content="https://ngn13.fun" property="og:url" />
  <meta content="#141414" data-react-helmet="true" name="theme-color" />
</svelte:head>

<Header>
  <c>echo</c> 
  hello world!
</Header>

<div class="flexbox">
  <Card title="whoami">
    👋 Hello! I'm ngn! 
    <ul>
      <li>🇹🇷 I'm a high school student from Turkey</li>
      <li>🖥️ I'm interested in cyber security, programming and electronics</li>
      <li>❤️ I love and support Free/Libre and Open Source Software (FLOSS)</li>
      <li>🐧 My GNU/Linux distribution of choice is Artix</li>
      <li>😋 In my free time, I like playing Minecraft</li>
    </ul>
  </Card>

  <Card title="pwd">
    You are currently on my personal website, for all you nerds, here is some 
    techinal details you may want to know:
    <ul>
      <li><c> </c> I built the frontend app using SvelteKit 
      <li><c> </c> I deploy my app to vecel using Github actions</li>
      <li><c>󰒋 </c> Backend server is written in Go and it's hosted on my server</li>
      <li><c> </c> And yup, I also have other services hosted on it</li>
    </ul>
  </Card>
</div>

<style>
.flexbox {
  padding: 50px;
  gap: 30px;
  display: flex;
}

ul {
  padding-left: 30px;
}

li {
  padding-top: 15px;
}

@media only screen and (max-width: 1076px) {
  .flexbox {
    flex-direction: column;
  }
}
</style>

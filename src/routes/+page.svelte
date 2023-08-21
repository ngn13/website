<script>
  import Header from "../lib/header.svelte";
  import Card from "../lib/card.svelte";
</script>

<svelte:head>
  <title>[ngn] | homepage</title> 
  <meta content="[ngn] | homepage" property="og:title" />
  <meta content="Homepage of my personal website" property="og:description" />
  <meta content="https://ngn13.fun" property="og:url" />
  <meta content="#000000" data-react-helmet="true" name="theme-color" />
</svelte:head>

<Header>
  <c>echo</c> 
  hello world!
</Header>

<main>
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
        <li><c> </c> I built the frontend app using SvelteKit</li> 
        <li><c> </c> I deploy my app to vercel using Github actions</li>
        <li><c>󰒋 </c> Backend server is written in Go and it's hosted on my server</li>
        <li><c> </c> And yup, I also have other services hosted on the same server</li>
      </ul>
    </Card>
  </div>

  <div class="flexbox">
    <Card title="ps -eaf">
      I usually spend my time...
      <ul>
        <li><c> </c> building random projects</li>
        <li><c>󰓂 </c> contributing stuff that I like</li> 
        <li><c>󰈻</c> solving CTFs</li> 
        <li><c>󰌢 </c> customizing my desktop</li>
        <li><c>󱀉 </c> posting random stuff on my blog, which you should definitely check out (it's very active)</li>
      </ul>
    </Card>

    <Card title="wall">
      If you want to have a chat, you can find me on the following platforms: 
      <ul>
        <li><c> </c> <a href="https://github.com/ngn13">Github</a></li> 
        <li><c>󰫑 </c> <a href="https://mastodon.social/@ngn">Mastodon</a></li> 
        <li><c>󰙯 </c> <a href="https://discord.com/users/568131907368386565">Discord</a></li>
        <li><c> </c> <a href="mailto:ngn13proton@proton.me">Email</a></li>
      </ul>
      Or you can DM me on matrix:
      <br>
      <c>@ngn:matrix.ngn13.fun</c>
    </Card>
  </div>
</main>

<div class="version">
  <p>v3.0</p>
</div>

<style>
main{
  display: flex;
  flex-direction: column;
  gap: 35px;
  padding: 50px;
}

.flexbox {
  display: flex;
  gap: 30px;
}

ul {
  padding-left: 30px;
  margin-bottom: 20px;
}

li {
  padding-top: 15px;
}

a {
  color: white;
  text-decoration: none;
}

a:hover {
  font-weight: 900;
}

.version {
  color: var(--dark-fife);
  position: fixed;
  bottom: 10px;
  right: 10px;
  font-size: 15px;
}

@media only screen and (max-width: 1076px) {
  .flexbox {
    flex-direction: column;
  }
}
</style>

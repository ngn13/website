<script>
  import Header from "../lib/header.svelte";
  import Card from "../lib/card.svelte";
</script>

<svelte:head>
  <title>[ngn.tf] | homepage</title> 
  <meta content="[ngn] | homepage" property="og:title" />
  <meta content="Homepage of my personal website" property="og:description" />
  <meta content="https://ngn.tf" property="og:url" />
  <meta content="#000000" data-react-helmet="true" name="theme-color" />
</svelte:head>

<Header>
  <c>echo</c> 
  hello world!
</Header>

<main>
  <div class="flexbox">
    <Card title="whoami">
      <div class="whoami-box">
        <div class="whoami-pic">
          <img alt="My profile pic" src="https://files.ngn.tf/pplow.png">
          <a href="https://keyoxide.org/F9E70878C2FB389AEC2BA34CA3654DF5AD9F641D">
            <c><i class="nf nf-oct-key"></i> Keyoxide</c>
          </a>
        </div>
        <div class="whoami-text">
          👋 Hello! I'm ngn! 
          <ul>
            <li>🇹🇷 I'm a high school student from Turkey</li>
            <li>🖥️ I'm interested in cyber security and programming.</li>
            <li>❤️ I love and support Free/Libre and Open Source Software (FLOSS)</li>
            <li>🐧 My GNU/Linux distribution of choice is Artix, however I am currently running Arch</li>
          </ul>
        </div>
      </div>
    </Card>
  </div>

  <div class="flexbox">
    <Card title="ps -eaf">
      I usually spend my time...
      <ul>
        <li><c>⌨️</c> building random projects</li>
        <li><c>👥</c> contributing stuff that I like</li> 
        <li><c>🚩</c> solving CTFs</li> 
        <li><c>🖥️</c> customizing my desktop</li>
        <li><c>📑</c> posting random stuff on my blog, you should definitely check it out btw (it's very active)</li>
      </ul>
    </Card>

    <Card title="wall">
      Here are some links if you want to get in contact with me, please do not 
      use these if you want to contact about the services that I offer, I have a 
      seperate email for that: <a href="mailto:services@ngn.tf"><c>services@ngn.tf</c></a>
      <ul>
        <li>
          <c><i class="nf nf-cod-github"></i></c> 
          <a href="https://github.com/ngn13">Github</a>
        </li> 
        <li>
          <c><i class="nf nf-md-mastodon"></i></c> 
          <a href="https://mastodon.social/@ngn" rel="me">Mastodon</a>
        </li> 
        <li>
          <c><i class="nf nf-md-email"></i></c> 
          <a href="mailto:ngn@ngn.tf">Email (personal)</a>
        </li> 
        <li>
          <c><i class="nf nf-md-xmpp"></i></c> 
          <a href="xmpp:ngn@chat.ngn.tf">XMPP (my preferred way of communication)</a>
        </li> 
      </ul>
    </Card>
  </div>
</main>

<div class="version">
  <p>v4.5</p>
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

.whoami-box {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 40px;
}

.whoami-pic {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 15px;
}

.whoami-pic img {
  width: 250px;
  border-radius: 30px;

  border: solid 1px white;
  animation-name: fullBorderAnimation;
  animation-duration: 10s;
  animation-iteration-count: infinite;
}


ul {
  list-style: inside;
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

@media only screen and (max-width: 1200px) {
  .flexbox {
    flex-direction: column;
  }
}

@media only screen and (max-width: 900px) {
  .whoami-box {
    flex-direction: column;
  }
}
</style>

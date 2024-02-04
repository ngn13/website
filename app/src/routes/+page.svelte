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
      👋 Hello! I'm ngn! 
      <ul>
        <li>🇹🇷 I'm a high school student from Turkey</li>
        <li>🖥️ I'm interested in cyber security, programming and electronics</li>
        <li>❤️ I love and support Free/Libre and Open Source Software (FLOSS)</li>
        <li>🐧 My GNU/Linux distribution of choice is Artix, however I am currently running Arch</li>
        <li>🎸 I enjoy listening random stuff (mostly modern rock), so <a href="https://hp.ngn.tf/watch?v=lYBUbBu4W08"><c>here</c></a> is a random song that I like!</li>
      </ul>
    </Card>

    <Card title="pwd">
      You are currently on my personal website, for all you nerds, here is some 
      techinal details you may want to know:
      <ul>
        <li>
          <c><i class="nf nf-md-license"></i></c> 
          I built this website from scratch and its licensed under <a href="https://www.gnu.org/licenses/agpl-3.0.en.html"><c>AGPLv3</c></a>.
          All the source is avaliable over on my github.
        </li> 
        <li>
          <c><i class="nf nf-md-file_image"></i></c> 
          All assets (images etc.) are hosted on my server
        </li>
        <li>
          <c><i class="nf nf-md-eye_off"></i></c> 
          No cloudflare, no CAPTCHAs, no CDNs, this website is %100 privacy friendly
        </li>
        <li>
          <c><i class="nf nf-md-server"></i></c> 
          You can discover other services and pages I host over on the <a href="/services"><c>services page</c></a>
        </li>
      </ul>
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
      Here are some links if you want to get in contact with me: 
      <ul>
        <li>
          <c><i class="nf nf-cod-github"></i></c> 
          <a href="https://github.com/ngn13">Github</a>
        </li> 
        <li>
          <c><i class="nf nf-md-mastodon"></i></c> 
          <a href="https://mastodon.social/@ngn">Mastodon</a>
        </li> 
        <li>
          <c><i class="nf nf-md-email"></i></c> 
          <a href="mailto:ngn@ngn.tf">Email</a>
        </li> 
      </ul>
      Or you can add me on XMPP: <c>ngn@chat.ngn.tf</c>
      <br>
      And lastly my PGP key is:
      <ul>
        <li class="pgp"><a href="https://keyoxide.org/F9E70878C2FB389AEC2BA34CA3654DF5AD9F641D"><c>F9E70878C2FB389AEC2BA34CA3654DF5AD9F641D</c></a></li>
      </ul>
    </Card>
  </div>
</main>

<div class="version">
  <p>v4.4</p>
</div>

<style>
main{
  display: flex;
  flex-direction: column;
  gap: 35px;
  padding: 50px;
}

.pgp {
  word-wrap: break-word;
  white-space: pre-wrap;
  word-break: break-word;
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

@media only screen and (max-width: 1200px) {
  .flexbox {
    flex-direction: column;
  }
}
</style>

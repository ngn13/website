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
  <link
    rel="alternate"
    type="application/atom+xml"
    href={import.meta.env.VITE_API_URL_DEV + "/blog/feed.atom"}
    title="Atom Feed"
  />
  <link
    rel="alternate"
    type="application/rss+xml"
    href={import.meta.env.VITE_API_URL_DEV + "/blog/feed.rss"}
    title="RSS Feed"
  />
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
          <img alt="My profile" src="https://files.ngn.tf/pplow.png" />
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
            <li>
              🐧 My GNU/Linux distribution of choice is Artix, however I am currently running Arch
            </li>
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
        <li>
          <c>📑</c> posting random stuff on my blog, you should definitely check it out btw (it's very
          active)
        </li>
      </ul>
    </Card>

    <Card title="wall">
      Here are some links if you want to get in contact with me, I highly prefer email and I usually
      respond to emails in 1 or 2 days, just make sure to check your spam folder (turns out running
      a TOR relay gets your IP into multiple blacklists)
      <ul>
        <li>
          <c><i class="nf nf-cod-github"></i></c>
          <a href="https://github.com/ngn13">Github</a>
        </li>
        <li>
          <c><i class="nf nf-md-mastodon"></i></c>
          <a href="https://defcon.social/@ngn" rel="me">Mastodon</a>
        </li>
        <li>
          <c><i class="nf nf-md-email"></i></c>
          <a href="mailto:ngn@ngn.tf">Email</a>
        </li>
        <li>
          <c><i class="nf nf-md-xmpp"></i></c>
          <a href="xmpp:ngn@chat.ngn.tf">XMPP</a>
        </li>
      </ul>
    </Card>
  </div>
</main>

<div class="version">
  <p>v5.0</p>
</div>

<style>
  main {
    display: flex;
    flex-direction: column;
    gap: 28px;
    padding: 50px;
  }

  .flexbox {
    display: flex;
    gap: 28px;
  }

  .whoami-box {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    gap: 35px;
  }

  .whoami-pic {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 6px;

    border-right: solid 1px var(--dark-fife);
    padding: 0 35px 0 10px;
  }

  .whoami-pic img {
    width: 200px;
    border-radius: 20px;

    border: solid 1px var(--border-color);
    animation-name: fullBorderAnimation;
    animation-duration: 10s;
    animation-iteration-count: infinite;

    box-shadow:
      rgba(50, 50, 93, 1) 0px 30px 60px -12px inset,
      rgba(0, 0, 0, 1) 0px 18px 36px -18px inset;
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
      gap: 25px;
    }

    .whoami-pic {
      border-right: none;
      padding: 0;
    }
  }
</style>

<script>
  import Header from "$lib/header.svelte";
  import Head from "$lib/head.svelte";
  import Card from "$lib/card.svelte";
  import Link from "$lib/link.svelte";

  import { color, language } from "$lib/util.js";
  import { _ } from "svelte-i18n";

  const { data } = $props();
  let projects = $state(data.projects);
</script>

<Head title="home" desc="home page of my personal website" />
<Header picture="tired" title={$_("home.title")} />

<main>
  <Card title={$_("home.welcome.title")}>
    <span> 👋 {$_("home.welcome.desc")}</span>
    <ul>
      <li>🇹🇷 {$_("home.welcome.whoami")}</li>
      <li>🖥️ {$_("home.welcome.interest")}</li>
      <li>❤️ {$_("home.welcome.support")}</li>
    </ul>
  </Card>
  <Card title={$_("home.work.title")}>
    <span>{$_("home.work.desc")}</span>
    <ul>
      <li>⌨️ {$_("home.work.build")}</li>
      <li>🤦 {$_("home.work.fix")}</li>
      <li>🚩 {$_("home.work.ctf")}</li>
      <li>👥 {$_("home.work.contribute")}</li>
      <li>📑 {$_("home.work.wiki")}</li>
    </ul>
  </Card>
  <Card title={$_("home.links.title")}>
    <span>{$_("home.links.desc")}:</span>
    <ul>
      <li>
        <Link icon="nf-fa-key" link="https://keyoxide.org/F9E70878C2FB389AEC2BA34CA3654DF5AD9F641D">
          PGP
        </Link>
      </li>
      <li>
        <Link icon="nf-md-email" link="mailto:ngn@ngn.tf">Email</Link>
      </li>
      <li>
        <Link icon="nf-md-mastodon" link="https://defcon.social/@ngn">Mastodon</Link>
      </li>
    </ul>
    <span>
      {$_("home.links.prefer")}
    </span>
  </Card>
  <Card title={$_("home.services.title")}>
    <span>
      {$_("home.services.desc")}
    </span>
    <ul>
      <li>
        <i style="color: var(--{color()});" class="nf nf-md-speedometer_slow"></i>
        {$_("home.services.speed")}
      </li>
      <li>
        <i style="color: var(--{color()});" class="nf nf-fa-lock"></i>
        {$_("home.services.security")}
      </li>
      <li>
        <i style="color: var(--{color()});" class="nf nf-fa-network_wired"></i>
        {$_("home.services.privacy")}
      </li>
      <li>
        <i style="color: var(--{color()});" class="nf nf-md-eye_off"></i>
        {$_("home.services.bullshit")}
      </li>
    </ul>
    <Link linK="/services">{$_("home.services.link")}</Link>
  </Card>
  <Card title={$_("home.projects.title")}>
    <span>
      {$_("home.projects.desc")}:
    </span>
    <ul>
      {#each projects as project}
        <li>
          <Link active={true} link={project.url}>{project.name}</Link>:
          {project.desc[$language]}
        </li>
      {/each}
    </ul>
  </Card>
</main>

<style>
  main {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    align-items: stretch;

    padding: 50px;
    gap: 28px;
  }

  @media only screen and (max-width: 900px) {
    main {
      flex-direction: column;
    }
  }
</style>

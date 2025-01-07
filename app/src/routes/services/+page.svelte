<script>
  import Service from "$lib/service.svelte";
  import Header from "$lib/header.svelte";
  import Link from "$lib/link.svelte";
  import Head from "$lib/head.svelte";

  import { _, locale } from "svelte-i18n";
  import { api_url } from "$lib/api.js";

  let { data } = $props();
  let list = $state(data.services);

  function change(input) {
    let value = input.target.value.toLowerCase();
    list = [];

    if (value === "") {
      list = data.services;
      return;
    }

    data.services.forEach((s) => {
      if (s.name.toLowerCase().includes(value)) list.push(s);
      else if (s.desc[$locale.slice(0, 2)].toLowerCase().includes(value)) list.push(s);
    });
  }
</script>

<Head title="services" desc="my self-hosted services and projects" />
<Header picture="cool" title={$_("services.title")} />

<main>
  <div class="title">
    <input oninput={change} type="text" placeholder={$_("services.search")} />
    <div>
      <Link icon="nf-fa-feed" link={api_url("/news/" + $locale.slice(0, 2))}
        >{$_("services.feed")}</Link
      >
    </div>
  </div>
  <div class="services">
    {#each list as service}
      <Service {service} />
    {/each}
  </div>
</main>

<style>
  main {
    padding: 50px;
    text-align: right;
  }

  main .title {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }

  main .services {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    align-items: stretch;
    margin-top: 20px;
    gap: 28px;
  }

  @media only screen and (max-width: 1200px) {
    main .services {
      flex-direction: column;
    }
  }
</style>

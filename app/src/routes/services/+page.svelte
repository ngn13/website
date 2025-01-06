<script>
  import Service from "$lib/service.svelte";
  import Header from "$lib/header.svelte";
  import Link from "$lib/link.svelte";
  import Head from "$lib/head.svelte";

  import { api_url } from "$lib/util.js";
  import { locale } from "svelte-i18n";

  export let data;

  let list = data.list,
    services = list;
  let value = "";

  function change(input) {
    value = input.target.value.toLowerCase();
    services = [];

    if (value === "") {
      services = list;
      return;
    }

    list.forEach((s) => {
      if (s.name.toLowerCase().includes(value)) services.push(s);
    });
  }
</script>

<Head title="services" desc="my self-hosted services and projects" />
<Header title="service status" picture="cool" />

<main>
  <div class="title">
    <input on:input={change} type="text" placeholder="Search for a service" />
    <div>
      <Link icon="nf-fa-feed" link={api_url("/news/" + $locale.slice(0, 2))}>News and updates</Link>
    </div>
  </div>
  <div class="services">
    {#each services as service}
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

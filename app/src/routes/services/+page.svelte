<script>
  import Service from "$lib/service.svelte";
  import Header from "$lib/header.svelte";
  import Error from "$lib/error.svelte";
  import Link from "$lib/link.svelte";
  import Head from "$lib/head.svelte";

  import { api_urljoin } from "$lib/api.js";
  import { locale, _ } from "svelte-i18n";

  let { data } = $props();
  let services = $state(data.services);

  function change(input) {
    let value = input.target.value.toLowerCase();
    services = [];

    if (value === "") {
      services = data.services;
      return;
    }

    data.services.forEach((s) => {
      if (s.name.toLowerCase().includes(value)) services.push(s);
      else if (s.desc[$locale].toLowerCase().includes(value)) services.push(s);
    });
  }

  function get_services() {
    return services.filter((s) => {
      return s.desc[$locale] !== "" && s.desc[$locale] !== null && s.desc[$locale] !== undefined;
    });
  }
</script>

<Head title="services" desc="my self-hosted services and projects" />
<Header picture="cool" title={$_("services.title")} />

{#if data.error !== undefined}
  <Error error={data.error} />
{:else}
  <main>
    <div class="title">
      <input oninput={change} type="text" placeholder={$_("services.search")} />
      <div>
        <Link icon="nf-fa-feed" link={api_urljoin("/news/" + $locale)}>{$_("services.feed")}</Link>
      </div>
    </div>
    <div class="services">
      {#if get_services().length == 0}
        <h3 class="none">{$_("services.none")}</h3>
      {:else}
        {#each get_services() as service}
          <Service {service} />
        {/each}
      {/if}
    </div>
  </main>
{/if}

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

  main .none {
    color: var(--white-3);
  }

  main .services {
    margin-top: 20px;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    align-items: stretch;
    gap: 28px;
  }

  @media only screen and (max-width: 1200px) {
    main .services {
      flex-direction: column;
    }
  }
</style>

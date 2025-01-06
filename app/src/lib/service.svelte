<script>
  import Icon from "$lib/icon.svelte";
  import Link from "$lib/link.svelte";

  import { color, time_from_ts } from "$lib/util.js";
  import { locale } from "svelte-i18n";

  export let service = {};
  let style = "";

  if (service.check_res == 0) style = "opacity: 70%";
</script>

<main {style}>
  <div class="info">
    <div class="title">
      <h1>{service.name}</h1>
      <p>{service.desc[$locale.slice(0, 2)]}</p>
    </div>
    <div class="links">
      <Link link={service.clear}><Icon icon="nf-oct-link" /></Link>
      {#if service.onion != ""}
        <Link link={service.onion}><Icon icon="nf-linux-tor" /></Link>
      {/if}
      {#if service.i2p != ""}
        <Link link={service.i2p}><span style="color: var(--{color()})">I2P</span></Link>
      {/if}
    </div>
  </div>
  <div class="check">
    <h1>Last checked at {time_from_ts(service.check_time)}</h1>
    {#if service.check_res == 0}
      <span style="background: var(--white-2)">Down</span>
    {:else if service.check_res == 1}
      <span style="background: var(--{color()})">Up</span>
    {:else if service.check_res == 2}
      <span style="background: var(--white-2)">Slow</span>
    {/if}
  </div>
</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    background: var(--black-3);
    border: solid 1px var(--black-4);
    text-align: left;

    flex: 1;
    flex-basis: 40%;
  }

  main .info {
    padding: 25px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    color: var(--white-1);
  }

  main .info .title h1 {
    font-size: var(--size-5);
    margin-bottom: 8px;
  }

  main .info .title p {
    font-size: var(--size-4);
    color: var(--white-2);
  }

  main .info .links {
    display: flex;
    flex-direction: row;
    gap: 10px;
    font-size: var(--size-6);
  }

  main .check {
    border-top: solid 1px var(--black-4);
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    color: var(--white-1);
  }

  main .check h1 {
    padding: 15px 25px 15px 25px;
    font-size: var(--size-4);
    font-weight: 100;
  }

  main .check span {
    padding: 15px 25px 15px 25px;
    font-size: var(--size-5);
    text-transform: uppercase;
    color: var(--black-1);
    font-weight: 1000;
  }
</style>

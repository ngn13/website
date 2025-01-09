<script>
  import Link from "$lib/link.svelte";
  import { color } from "$lib/util.js";
  import { _ } from "svelte-i18n";

  export let error = "";
</script>

<main>
  <h1 style="color: var(--{color()})">{$_("error.title")}</h1>
  <code>
    {#if error === ""}
      Unknown error
    {:else}
      {error}
    {/if}
  </code>
  <Link link={import.meta.env.APP_REPORT_URL}>
    {$_("error.report")}
  </Link>
  <img src="/profile/sad.png" alt="" />
</main>

<style>
  main {
    position: fixed;

    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    z-index: 1;

    display: flex;
    flex-direction: column;
    justify-content: end;
    gap: 10px;

    padding: 50px;
    font-size: var(--size-4);

    background: var(--background);
    background-size: 50%;
  }

  main h1 {
    font-size: var(--size-6);
  }

  main code {
    font-size: var(--size-4);
    color: var(--white-2);
  }

  main img {
    width: var(--profile-size);
    position: absolute;
    right: 0;
    bottom: 0;
  }
</style>

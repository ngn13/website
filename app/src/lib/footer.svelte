<script>
  import { urljoin, color, date_from_ts } from "$lib/util.js";
  import { api_get_metrics } from "$lib/api.js";
  import Link from "$lib/link.svelte";

  import { onMount } from "svelte";
  import { _ } from "svelte-i18n";

  let data = {};

  onMount(async () => {
    data = await api_get_metrics(fetch);
  });
</script>

<footer style="border-top: solid 2px var(--{color()});">
  <div class="links">
    <span>
      <Link link={import.meta.env.WEBSITE_SOURCE_URL} bold={true}>{$_("footer.source")}</Link>
    </span>
    <span>/</span>
    <span>
      <Link link={urljoin(import.meta.env.WEBSITE_APP_URL, "doc/license")} bold={true}
        >{$_("footer.license")}</Link
      >
    </span>
    <span>/</span>
    <span>
      <Link link={urljoin(import.meta.env.WEBSITE_APP_URL, "doc/privacy")} bold={true}
        >{$_("footer.privacy")}</Link
      >
    </span>
  </div>
  <span class="counter">
    {$_("footer.number", {
      values: {
        total: data.total,
        since: date_from_ts(data.since),
      },
    })}
    {#if data.number % 1000 == 0}
      <span style="color: var(--{color()})">({$_("footer.wow")})</span>
    {/if}
  </span>
</footer>

<style>
  footer {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    background: var(--black-1);

    box-sizing: border-box;
    padding: 20px 50px 20px 50px;
  }

  div {
    display: flex;
    font-size: var(--size-2);
    flex-direction: column;
    gap: 5px;
  }

  span {
    color: var(--white-2);
    font-size: 15px;
  }

  .counter {
    text-align: right;
  }

  .links {
    text-align: left;
    display: flex;
    flex-direction: row;
    gap: 5px;
  }
</style>

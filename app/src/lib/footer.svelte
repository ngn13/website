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
  <div class="info">
    <div class="links">
      <span>
        <Link link={import.meta.env.APP_SOURCE_URL} bold={true}>{$_("footer.source")}</Link>
      </span>
      <span>/</span>
      <span>
        <Link link={urljoin(import.meta.env.APP_URL, "doc/license")} bold={true}
          >{$_("footer.license")}</Link
        >
      </span>
      <span>/</span>
      <span>
        <Link link={urljoin(import.meta.env.APP_URL, "doc/privacy")} bold={true}
          >{$_("footer.privacy")}</Link
        >
      </span>
    </div>
    <span>
      {$_("footer.powered")}
    </span>
  </div>
  <div class="useless">
    <span>
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
    <span>
      {$_("footer.version", { values: { api_version: "v1", frontend_version: pkg.version } })}
    </span>
  </div>
</footer>

<style>
  footer {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    background: var(--black-1);
  }

  div {
    display: flex;
    color: var(--white-2);
    font-size: var(--size-2);
    flex-direction: column;
    gap: 5px;
  }

  .useless {
    margin: 25px 50px 25px 0;
    text-align: right;
  }

  .info {
    margin: 25px 0 25px 50px;
    text-align: left;
  }

  .info .links {
    display: flex;
    flex-direction: row;
    gap: 5px;
  }
</style>

<script>
  import { api_urljoin } from "$lib/api.js";
  import { app_url } from "$lib/util.js";

  export let desc, title;
</script>

<svelte:head>
  <title>[ngn.tf] | {title}</title>

  <meta content="[ngn.tf] | {title}" property="og:title" />
  <meta content={desc} property="og:description" />
  <meta content={app_url()} property="og:url" />
  <meta content="#000000" data-react-helmet="true" name="theme-color" />

  <link
    rel="alternate"
    type="application/atom+xml"
    href={api_urljoin("/news/en")}
    title="Atom Feed"
  />
</svelte:head>

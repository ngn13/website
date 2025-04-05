<script>
  import { api_urljoin } from "$lib/api.js";
  import { app_url } from "$lib/util.js";

  export let desc, title;
</script>

<svelte:head>
  <title>[ngn.tf] | {title}</title>

  <meta name="description" content={desc} />
  <meta name="author" content="ngn" />
  <meta name="keywords" content="ngn,ngn13,ngn1,ngn.tf" />
  <meta name="color-scheme" content="only dark" />
  <meta name="theme-color" content="#000000" />

  <meta property="og:title" content="[ngn.tf] | {title}" />
  <meta property="og:description" content={desc} />
  <meta property="og:url" content={app_url()} />

  <link
    rel="alternate"
    type="application/atom+xml"
    href={api_urljoin("/news/en")}
    title="Service news and updates"
  />
</svelte:head>

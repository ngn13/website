<script>
  import Header from "$lib/header.svelte";
  import Error from "$lib/error.svelte";
  import Head from "$lib/head.svelte";

  import { locale, _ } from "svelte-i18n";
  import { goto } from "$app/navigation";
  import { color } from "$lib/util.js";
  import DOMPurify from "dompurify";
  import { onMount } from "svelte";
  import { marked } from "marked";

  let { data } = $props();
  marked.use({ breaks: true });

  onMount(async () => {
    for (let key in data.doc)
      data.doc[key]["content"] = DOMPurify.sanitize(data.doc[key]["content"]);

    if (undefined !== data.error && data.error.includes("not found")) goto("/");
  });
</script>

<Head title="documentation" desc="website and API documentation" />
<Header picture="reader" title={data.doc[$locale].title} />

{#if data.error.length !== 0}
  {#if !data.error.includes("not found")}
    <Error error={data.error} />
  {/if}
{:else}
  <main>
    <div class="markdown-body" style="--link-color: var(--{color()})">
      {@html marked.parse(data.doc[$locale].content)}
    </div>
  </main>
{/if}

<style>
  @import "/css/markdown.css";

  main {
    padding: 50px;
    gap: 30px;
  }

  main .markdown-body :global(a) {
    color: var(--link-color);
  }
</style>

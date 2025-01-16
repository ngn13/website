<script>
  import Header from "$lib/header.svelte";
  import Head from "$lib/head.svelte";

  import { color } from "$lib/util.js";
  import { marked } from "marked";
  import { _ } from "svelte-i18n";

  let { data } = $props();
  marked.use({ breaks: true });
</script>

<Head title="documentation" desc="website and API documentation" />
<Header picture="reader" title={$_("doc.title")} />

<main>
  <div class="markdown-body" style="--link-color: var(--{color()})">
    {@html marked.parse(data.content)}
  </div>
  <div class="docs">
    {#each data.docs as doc}
      {#if doc.title == data.title}
        <a href="/doc/{doc.name}" style="border-color: var(--{color()})">
          <h1>{doc.title}</h1>
          <h3>{doc.desc}</h3>
        </a>
      {:else}
        <a href="/doc/{doc.name}" style="border-color: var(--white-3)">
          <h1>{doc.title}</h1>
          <h3>{doc.desc}</h3>
        </a>
      {/if}
    {/each}
  </div>
</main>

<style>
  @import "/markdown.css";

  main {
    padding: 50px;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: start;
    gap: 30px;
  }

  main .docs {
    display: flex;
    flex-direction: column;
    align-items: end;
    gap: 5px;
  }

  main .docs a {
    display: flex;
    flex-direction: column;
    background: var(--black-3);
    text-decoration: none;
    box-sizing: border-box;
    border-right-style: solid;
    padding: 15px;
    width: 100%;
    gap: 4px;
  }

  main .docs a:hover {
    box-shadow: var(--box-shadow);
  }

  main .docs a h1 {
    font-size: var(--size-3);
    color: var(--white-1);
    font-weight: 900;
  }

  main .docs a h3 {
    font-size: var(--size-2);
    color: var(--white-3);
    font-weight: 100;
    text-decoration: none;
  }

  main .markdown-body :global(a) {
    color: var(--link-color);
  }

  @media only screen and (max-width: 900px) {
    main {
      flex-direction: column-reverse;
    }

    main .docs {
      width: 100%;
    }

    main .docs a {
      border-right-style: none;
      border-left-style: solid;
      width: 100%;
    }
  }
</style>

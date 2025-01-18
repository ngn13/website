<script>
  import Header from "$lib/header.svelte";
  import Error from "$lib/error.svelte";
  import Head from "$lib/head.svelte";

  import { locale, _ } from "svelte-i18n";
  import { color } from "$lib/util.js";
  import DOMPurify from "dompurify";
  import { onMount } from "svelte";
  import { marked } from "marked";

  let { data } = $props();
  marked.use({ breaks: true });

  onMount(async () => {
    for (let key in data.doc)
      data.doc[key]["content"] = DOMPurify.sanitize(data.doc[key]["content"]);
  });
</script>

<Head title="documentation" desc="website and API documentation" />
<Header picture="reader" title={$_("doc.title")} />

{#if data.error !== undefined}
  <Error error={data.error} />
{:else}
  <main>
    {#if data.doc !== undefined}
      <div class="markdown-body" style="--link-color: var(--{color()})">
        {@html marked.parse(data.doc[$locale].content)}
      </div>
      <div class="docs">
        {#each data.docs[$locale] as doc}
          {#if doc.title == data.doc[$locale].title}
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
    {/if}
  </main>
{/if}

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
    gap: 6px;
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
    box-shadow: var(--box-shadow-2);
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

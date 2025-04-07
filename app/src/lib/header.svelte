<script>
  import { browser } from "$app/environment";
  import { color } from "$lib/util.js";
  import { onMount } from "svelte";
  import { _ } from "svelte-i18n";

  export let picture = "";
  export let title = "";

  let title_cur = title;
  let show_animation = false;

  function animate(title) {
    if (!browser) return;

    let id = window.setTimeout(function () {}, 0);

    while (id--) clearTimeout(id);

    title_cur = "";

    for (let i = 0; i < title.length; i++) {
      setTimeout(() => {
        title_cur += title[i];
      }, i * 70);
    }
  }

  onMount(() => {
    show_animation = true;
  });

  $: animate(title);
</script>

<header>
  <div>
    {#if show_animation}
      <h1 class="title" style="color: var(--{color()})">{title_cur}</h1>
      <h1 class="cursor" style="color: var(--{color()})">_</h1>
    {:else}
      <h1 class="title" style="color: var(--{color()})">{title}</h1>
    {/if}
  </div>
  <img src="/profile/{picture}.png" alt="" />
</header>

<style>
  header {
    background: var(--background);
    background-size: 50%;

    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: end;
  }

  header div {
    display: flex;
    flex-direction: row;
    align-items: end;
    padding: 50px 50px 30px 50px;
    font-size: var(--size-6);
    font-family:
      Consolas,
      Monaco,
      Lucida Console,
      Liberation Mono,
      DejaVu Sans Mono,
      Bitstream Vera Sans Mono,
      Courier New,
      monospace;
    white-space: nowrap;
    justify-content: start;
    width: min-content;
  }

  header div .title {
    text-shadow: var(--text-shadow);
    overflow: hidden;
  }

  header div .cursor {
    content: "_";
    display: inline-block;
    animation: blink 1.5s steps(2) infinite;
  }

  header img {
    padding: 50px 50px 0 50px;
    width: var(--profile-size);
    bottom: 0;
    left: 0;
  }

  @media only screen and (max-width: 900px) {
    header {
      display: block;
    }

    header img {
      display: none;
    }
  }
</style>

<script>
  import { color } from "$lib/util.js";
  import { _ } from "svelte-i18n";

  export let picture = "";
  export let title = "";
</script>

<header>
  <div>
    <h1 class="title" style="color: var(--{color()})">
      {title.toLowerCase()}
    </h1>
    <h1 class="cursor" style="color: var(--{color()})">_</h1>
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
    width: 0;
    animation: typing 1s steps(20, end) forwards;
    animation-delay: 0.3s;
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

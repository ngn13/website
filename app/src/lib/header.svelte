<script>
  import { color } from "$lib/util.js";
  import { _ } from "svelte-i18n";

  export let picture = "";
  export let title = "";

  let current = "";

  for (let i = 0; i < title.length; i++) {
    setTimeout(
      () => {
        current += title[i];
      },
      100 * (i + 1)
    );
  }
</script>

<header>
  <h1 style="color: var(--{color()})">{current}</h1>
  <img src="/profile/{picture}.png" alt="" />
</header>

<style>
  header {
    background: linear-gradient(rgba(11, 11, 11, 0.808), rgba(1, 1, 1, 0.96)),
      url("https://files.ngn.tf/banner.png");
    background-size: 50%;
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: end;
  }

  img {
    padding: 50px 50px 0 50px;
    width: 220px;
    bottom: 0;
    left: 0;
  }

  h1 {
    font-size: var(--size-7);
    font-family:
      Consolas,
      Monaco,
      Lucida Console,
      Liberation Mono,
      DejaVu Sans Mono,
      Bitstream Vera Sans Mono,
      Courier New,
      monospace;
    padding: 50px 50px 30px 50px;
    white-space: nowrap;
    text-align: center;
    color: white;
    text-shadow: var(--text-shadow);
  }

  h1::after {
    content: "_";
    display: inline-block;
    animation: blink 1.5s steps(2) infinite;
  }

  @media only screen and (max-width: 900px) {
    header {
      display: block;
    }

    h1 {
      padding: 80px;
    }

    img {
      display: none;
    }
  }
</style>

<script>
  import { color, click } from "$lib/util.js";
  export let link;
</script>

<a
  style="text-decoration-color: var(--{color()})"
  data-sveltekit-preload-data
  on:click={click}
  href={link}
>
  <slot></slot>
</a>

<style>
  a {
    font-weight: 700;
    font-size: 20px;
    text-decoration: none;
    color: white;
    cursor: pointer;
  }

  a:hover {
    text-decoration: underline;
    text-shadow: 3px 4px 7px rgba(81, 67, 21, 0.8);
  }
</style>

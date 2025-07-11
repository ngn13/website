<script>
  import { color, click } from "$lib/util.js";
  import { page } from "$app/stores";

  export let link;
</script>

<a
  style="
    text-decoration-color: var(--{color()});
    {$page.url.pathname === link ? `color: var(--${color()});` : ''}
  "
  data-sveltekit-preload-data
  on:click={click}
  href={link}
>
  <slot></slot>
</a>

<style>
  a {
    font-weight: 900;
    font-size: var(--size-4);
    color: var(--white-1);
  }
</style>

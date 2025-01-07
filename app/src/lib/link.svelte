<script>
  import Icon from "$lib/icon.svelte";
  import { color } from "$lib/util.js";

  const default_color = "white-1";

  export let active = false;
  export let highlight = true;
  export let link = "";
  export let icon = "";

  let style = "";

  if (highlight) style = `text-decoration-color: var(--${color()});`;

  if (active) style += `color: var(--${color()});`;
  else style += `color: var(--${default_color});`;
</script>

{#if icon != ""}
  <Icon {icon} />
{/if}
{#if highlight}
  <a {style} href={link}>
    <slot></slot>
  </a>
{:else}
  <a {style} class="no-highlight" href={link}>
    <slot></slot>
  </a>
{/if}

<style>
  .no-highlight:hover {
    text-decoration: none;
  }
</style>

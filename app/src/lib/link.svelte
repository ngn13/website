<script>
  import Icon from "$lib/icon.svelte";
  import { color } from "$lib/util.js";

  const default_color = "white-1";

  export let active = false;
  export let link = "";
  export let icon = "";

  let style = `text-decoration-color: var(--${color()});`;

  if (active) style += `color: var(--${color()});`;
  else style += `color: var(--${default_color});`;
</script>

{#if icon != ""}
  <Icon {icon} />
{/if}
<a {style} href={link}>
  <slot></slot>
</a>

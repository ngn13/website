<script>
  import { locale } from "svelte-i18n";
  import languages from "$lib/lang.js";
  let icon = "",
    indx = 0,
    len = languages.length;

  function next() {
    if (indx >= languages.length) indx = 0;

    icon = languages[indx].icon;
    locale.set(languages[indx++].code);
  }

  for (indx = 0; indx < len; indx++) {
    if (languages[indx].code == $locale.slice(0, 2)) {
      icon = languages[indx++].icon;
      break;
    }
  }
</script>

<button on:click={next}>{icon}</button>

<style>
  button {
    background: var(--black-2);
    color: var(--white-1);
    font-size: var(--size-4);
    outline: none;
    border: none;
    transition: 0.4s;
  }

  button:hover {
    background: var(--black-1);
  }
</style>

<script>
  import { language, set_lang } from "$lib/util.js";
  import languages from "$lib/lang.js";

  let icon = null,
    indx = 0,
    len = languages.length;

  function next_indx() {
    if (indx + 1 >= len) return 0;
    return indx + 1;
  }

  function next_lang(inc) {
    let new_indx = next_indx();
    if (inc) indx = new_indx;
    return languages[new_indx];
  }

  function next() {
    set_lang(next_lang(true).code);
    icon = next_lang(false).icon;
  }

  for (indx = 0; indx < len; indx++) {
    if (languages[indx].code == $language) {
      set_lang(languages[indx].code);
      icon = next_lang(false).icon;
      break;
    }
  }
</script>

<button on:click={next}>
  {icon}
</button>

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

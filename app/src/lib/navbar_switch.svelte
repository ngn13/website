<script>
  import { locale_list, locale_select, locale_index } from "$lib/locale.js";

  let len = locale_list.length;

  function get_next(indx) {
    let new_indx = 0;

    if (indx + 1 >= len) indx = 0;
    else new_indx = indx + 1;

    return locale_list[new_indx];
  }

  function next() {
    locale_select(get_next($locale_index).code);
  }
</script>

<button on:click={next}>
  {get_next($locale_index).icon}
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

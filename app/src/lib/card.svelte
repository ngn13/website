<script>
  export let title;
</script>

<main>
  <h1 class="title">{title}</h1>
  <div>
    <slot></slot>
  </div>
</main>

<style>
  main {
    flex: 1;
    flex-basis: 30%;

    display: flex;
    flex-direction: column;
  }

  main .title {
    font-family:
      Consolas,
      Monaco,
      Lucida Console,
      Liberation Mono,
      DejaVu Sans Mono,
      Bitstream Vera Sans Mono,
      Courier New,
      monospace;
    color: var(--white-1);
  }

  main .title::before {
    content: "#";
    margin: 0 10px 0 0;
    color: var(--white-3);
  }

  main div {
    border-left: solid 1px var(--black-4);
    padding: 25px 25px 10px 25px;
    font-size: var(--size-4);
    color: var(--white-1);
    word-wrap: break-word;
    align-items: center;
    margin-left: 7px;
    flex: 1;
  }
</style>

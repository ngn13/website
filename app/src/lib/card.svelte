<script>
  export let title;

  let current = "";
  let i = 0;

  while (title.length > i) {
    let c = title[i];
    setTimeout(
      () => {
        current += c;
      },
      100 * (i + 1)
    );
    i += 1;
  }
</script>

<div class="main">
  <div class="title">
    root@ngn.tf:~# {current}
  </div>
  <div class="content">
    <slot></slot>
  </div>
</div>

<style>
  .main {
    display: flex;
    flex-direction: column;
    width: 100%;
    background: var(--dark-three);
    box-shadow: var(--box-shadow);
    border-radius: var(--radius);
    border: solid 1px var(--border-color);
  }

  .title {
    background: var(--dark-two);
    padding: 25px;
    border-radius: 7px 7px 0px 0px;
    font-size: 20px;
    font-family:
      Consolas,
      Monaco,
      Lucida Console,
      Liberation Mono,
      DejaVu Sans Mono,
      Bitstream Vera Sans Mono,
      Courier New,
      monospace;
    color: white;
  }

  .content {
    background: var(--dark-three);
    padding: 30px;
    color: white;
    border-radius: 5px;
    font-size: 25px;
  }
</style>

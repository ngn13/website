<script>
  export let title;
  export let url;
  let audio;

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

  function epicSound() {
    audio.play();
  }
</script>

<a on:click={epicSound} data-sveltekit-preload-data href={url}>
  <audio bind:this={audio} preload="auto">
    <source src="/click.wav" type="audio/mpeg" />
  </audio>
  <div class="title">
    {current}
  </div>
  <div class="content">
    <slot></slot>
  </div>
</a>

<style>
  a {
    display: flex;
    flex-direction: column;
    width: 100%;
    background: var(--dark-three);
    box-shadow: var(--box-shadow);
    border-radius: var(--radius);
    cursor: pointer;
    transition: 0.4s;
    text-decoration: none;
    border: solid 1px var(--border-color);
  }

  a:hover > .title {
    text-shadow: var(--text-shadow);
  }

  .title {
    border: solid 1px var(--dark-two);
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
    padding-top: 30px;
    color: white;
    border-radius: 5px;
    font-size: 25px;
  }
</style>

<script>
  import { page } from "$app/stores"
  export let link 
  let audio

  function epicSound() {
    audio.play()
  }
</script>

<div>
  <audio bind:this={audio} preload="auto">
    <source src="/click.wav" type="audio/mpeg" />
  </audio>
  <a data-sveltekit-preload-data on:click={epicSound} href="{link}"><slot></slot></a>
</div>

<style>
a {
  margin-left: 20px;
  font-weight: 700;
  font-size: 22px;
  text-decoration: none;
  color: white;
  cursor: pointer;
}

a:hover {
  text-decoration: underline;
  text-shadow: 3px 4px 7px rgba(81, 67, 21, 0.8);
  animation-name: underlineAnimation;
  animation-duration: 5s;
  animation-iteration-count: infinite;
}
</style>

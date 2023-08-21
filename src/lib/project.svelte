<script>
  import { goto } from "$app/navigation";

  export let desc
  export let url
  let audio

  function redirect() {
    audio.play()
    goto(url);
  }
</script>

<main on:click={redirect}>
  <audio bind:this={audio} preload="auto">
    <source src="/click.wav" type="audio/mpeg" />
  </audio>
  <h1 class="c"> </h1>
  <div>
    <h1><slot></slot></h1>
    <p>{desc}</p>
  </div>
</main>

<style>
main {
  display: flex;
  padding: 40px 60px 40px 60px;
  background: var(--dark-two);
  border-radius: 7px 7px 0px 0px;
  justify-content: center;
  align-content: center;
  flex-direction: row;
  cursor: pointer;
  transition: .4s;
  border: none;
  color: white;
  gap: 20px;
}

main:hover {
  box-shadow: var(--box-shadow);
}

h1 {
  font-size: 70px;
}

div h1 {
  font-size: 30px;
}

div p {
  margin-top: 10px;
  font-size: 20px; 
}
</style>

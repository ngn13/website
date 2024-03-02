<script>
  export let desc
  export let url

  let icon = "<i class='nf nf-md-clipboard_multiple'></i>"
  let audio

  function copy() {
    audio.play()
    navigator.clipboard.writeText(url)
    icon = "<i class='nf nf-md-clipboard_check'></i>"
    setTimeout(()=>{
      icon = "<i class='nf nf-md-clipboard_multiple'></i>"
    }, 500)
  }
</script>

<main>
  <audio bind:this={audio} preload="auto">
    <source src="/click.wav" type="audio/mpeg" />
  </audio>
  <div>
    <h1><slot></slot></h1>
    <p>{desc}</p>
  </div>
  <div>
    <button on:click={copy}>{@html icon}</button>
    <a href="{url}"><i class="nf nf-oct-link_external"></i></a>
  </div>
</main>

<style>
main {
  display: flex;
  flex-direction: row;
  padding: 30px 30px 30px 30px;
  background: var(--dark-two);
  border-radius: var(--radius);
  box-shadow: var(--box-shadow);
  justify-content: space-between;
  align-items: center;
  border: none;
  color: white;
  gap: 100px;
  transition: .4s;
  flex-grow: 1;
}

div h1 {
  animation-name: colorAnimation;
  animation-duration: 10s;
  animation-iteration-count: infinite;
  font-size: 30px;
}

div p {
  margin-top: 10px;
  font-size: 20px; 
}

a, button {
  text-align: center;
  font-size: 30px;
  text-decoration: none;
  color: white;
  border: none;
  background: none;
  outline: none;
  cursor: pointer;
}

a:hover, button:hover{
  animation-name: colorAnimation;
  animation-duration: 5s;
  animation-iteration-count: infinite;
}
</style>

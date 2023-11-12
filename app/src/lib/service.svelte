<script>
  export let desc
  export let url

  let icon = "󰅇"
  let audio

  function copy() {
    audio.play()
    navigator.clipboard.writeText(url)
    icon = "󰅎"
    setTimeout(()=>{
      icon = "󰅇"
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
    <a on:click={copy} href="#">{icon}</a>
    <a href="{url}"></a>
  </div>
</main>

<style>
main {
  display: flex;
  flex-direction: row;
  padding: 30px 30px 30px 30px;
  background: var(--dark-two);
  border-radius: 7px 7px 0px 0px;
  box-shadow: var(--box-shadow);
  justify-content: space-between;
  align-items: center;
  border: none;
  color: white;
  gap: 100px;
  transition: .4s;
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

a {
  text-align: center;
  font-size: 30px;
  text-decoration: none;
  color: white;
  border: none;
}

a:hover {
  animation-name: colorAnimation;
  animation-duration: 5s;
  animation-iteration-count: infinite;
}
</style>

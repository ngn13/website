<script>
  import Header from "../../lib/header.svelte";
  import Service from "../../lib/service.svelte";

  export let data 
</script>

<svelte:head>
  <title>[ngn] | services</title> 
  <meta content="[ngn] | services" property="og:title" />
  <meta content="Stuff that I host" property="og:description" />
  <meta content="https://ngn.tf" property="og:url" />
  <meta content="#000000" data-react-helmet="true" name="theme-color" />
</svelte:head>
<Header><c>ls -l</c> services</Header>
<main>
  <div class="info">
    <h1></h1>
    <p>
      These are free to use FOSS services that I host on my server. 
      If you want an account on any of these services, or if you are having issues with them, 
      please get in contact with me 
    </p>
  </div>
  {#each data.services as services_list}
    <div class="flexrow">
      {#each services_list as service}
        <Service url="{service.url}" desc="{service.desc}">{service.name}</Service>
      {/each} 
    </div>
  {/each}
</main>

<style>
.info h1 {
  margin-right: 40px;
  animation-name: colorAnimation;
  animation-duration: 5s;
  animation-iteration-count: infinite;
}

.info {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  padding: 40px;
  color: white;
  border-radius: var(--radius); 
  font-size: 25px;
  margin-bottom: 30px;
}

main {
  display: flex;
  flex-direction: column;
  align-content: center;
  justify-content: center;
  padding: 50px;
  gap: 25px;
}

.flexrow {
  display: flex;
  flex-direction: row;
  align-content: center;
  justify-content: center;
  gap: 25px;
}

@media only screen and (max-width: 1316px) {
  .flexrow {
    flex-direction: column;
  }
}
</style>

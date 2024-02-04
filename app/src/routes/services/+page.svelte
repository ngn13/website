<script>
  import Header from "../../lib/header.svelte";
  import Service from "../../lib/service.svelte";

  export let data 
</script>

<svelte:head>
  <title>[ngn.tf] | services</title> 
  <meta content="[ngn] | services" property="og:title" />
  <meta content="Stuff that I host" property="og:description" />
  <meta content="https://ngn.tf" property="og:url" />
  <meta content="#000000" data-react-helmet="true" name="theme-color" />
</svelte:head>
<Header><c>ls</c> services</Header>
<main>
  {#each data.services as services_list}
    <div class="flexrow">
      {#each services_list as service}
        <Service url="{service.url}" desc="{service.desc}">{service.name}</Service>
      {/each} 
    </div>
  {/each}
</main>

<style>
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

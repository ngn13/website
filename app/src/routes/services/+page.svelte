<script>
  import Header from "../../lib/header.svelte";
  import Service from "../../lib/service.svelte";
  import Card from "../../lib/card.svelte";

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
  <Card title="cat services/*/info.txt">
  <div class="flexcol">
    {#each data.services as services_list}
      <div class="flexrow">
        {#each services_list as service}
          <Service url="{service.url}" desc="{service.desc}">{service.name}</Service>
        {/each} 
      </div>
    {/each}
  </div>
  </Card>

  <Card title="cat services/details.txt">
    Here some details for all the services I offer:
    <ul>
      <li>
        <c><i class="nf nf-cod-account"></i> Registration:</c> All the services are offered for free, and all of them 
        are accessiable to public. If you want to use a service that requires 
        registration, you can email <a href="mailto:services@ngn.tf"><c>services@ngn.tf</c></a>
        for an account.
      </li>
      <li>
        <c><i class="nf nf-fa-eye_slash"></i> Privacy:</c> To protect user privacy, all the web proxy logs are cleared regularly. 
        I also do not use any kind of CDN, and provide SSL encrypted connection for all the services. You can also connect all the 
        services over TOR, as I do not block any traffic from TOR.
      </li>
      <li>
        <c><i class="nf nf-oct-graph"></i> Uptime:</c> Some services get restarted regularly, also sometimes I have 
        issues with the hosting, so I cannot provide or guarantee %100 uptime for any of the services. I also cannot guarantee 
        the that there won't be any data loss. I do take any regular backups, but your data may be lost in case of something 
        like a SSD failure. 
      </li>
      <li>
        <c><i class="nf nf-md-speedometer"></i> Speed:</c> All the services are located in Turkey, and avaliable 
        over an 233.68 Mbit/s interface. If you are close to Turkey you should have a fairly good experience. 
        If you are not, then you should probably use another provider for the service.
      </li>
    </ul>
  </Card>

</main>

<style>
main {
  display: flex;
  flex-direction: column;
  align-content: center;
  justify-content: center;
  padding: 50px;
  gap: 30px;
}

.flexcol {
  display: flex;
  flex-direction: column;
  align-content: center;
  justify-content: center;
  gap: 25px;
}

.flexrow {
  display: flex;
  flex-direction: row;
  align-content: center;
  justify-content: center;
  width: 100%;
  gap: 25px;
}

a {
  text-decoration: none;
}

ul {
  list-style: inside;
  margin-bottom: 20px;
}

li {
  padding-top: 30px;
  line-height: 35px;
}

@media only screen and (max-width: 1316px) {
  .flexrow {
    flex-direction: column;
  }
}
</style>

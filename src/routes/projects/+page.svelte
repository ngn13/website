<script>
  import Header from "../../lib/header.svelte";
  import Project from "../../lib/project.svelte";

  export let data 
  console.log(data)
</script>

<svelte:head>
  <title>[ngn] | projects</title> 
  <meta content="[ngn] | projects" property="og:title" />
  <meta content="Projects that I work on" property="og:description" />
  <meta content="https://ngn13.fun" property="og:url" />
  <meta content="#000000" data-react-helmet="true" name="theme-color" />
</svelte:head>
<Header><c>ls -l</c> projects</Header>
<main>
  {#each data.projects as projectlist}
    <div class="flexrow">
      {#each projectlist as project}
        <Project url="{project.link}" desc="{project.desc}">{project.name}</Project>
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

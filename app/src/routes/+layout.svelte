<script>
  import Navbar from "$lib/navbar.svelte";
  import Footer from "$lib/footer.svelte";
  import Error from "$lib/error.svelte";

  let { data, children } = $props();
</script>

<main>
  {#if data.error === null}
    <Navbar />
    <div class="content">
      {@render children()}
    </div>
    <Footer />
  {:else}
    <Error error={data.error} />>
  {/if}
</main>

<style>
  @import "../../static/global.css";

  main {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
  }

  .content {
    background: var(--black-1);
    flex-grow: 1;
  }
</style>

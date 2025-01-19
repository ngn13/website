<script>
  import Navbar from "$lib/navbar.svelte";
  import Footer from "$lib/footer.svelte";

  import { locale_select } from "$lib/locale.js";
  import { onMount } from "svelte";

  let { children } = $props();

  onMount(() => {
    locale_select();
  });
</script>

<main>
  <Navbar />
  <div class="content">
    {@render children()}
  </div>
  <Footer />
</main>

<style>
  @import "/global.css";

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

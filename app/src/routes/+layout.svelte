<script>
  import Navbar from "$lib/navbar.svelte";
  import Footer from "$lib/footer.svelte";
</script>

<main>
  <Navbar />
  <div class="content">
    <slot></slot>
  </div>
  <Footer />
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

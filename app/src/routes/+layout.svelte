<script>
  import Navbar from "../lib/navbar.svelte";
</script>

<main>
  <Navbar />
  <slot></slot>
</main>

<style>
  @import "../../static/global.css";
</style>

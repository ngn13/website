<script>
  import { color } from "$lib/util.js";
  import NavbarLink from "./navbar_link.svelte";
</script>

<nav style="border-bottom: solid 2px var(--{color()});">
  <h3 style="color: var(--{color()})">[ngn.tf]</h3>

  <div>
    <NavbarLink link="/">home</NavbarLink>
    <NavbarLink link="/news">news</NavbarLink>
    <NavbarLink link="/services">services</NavbarLink>
    <!-- <NavbarLink link="/donate">donate</NavbarLink> -->
  </div>
</nav>

<style>
  nav {
    background: var(--dark-one);
    padding: 20px 30px 20px 20px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    /*animation-name: borderAnimation;
  animation-duration: 10s;
  animation-iteration-count: infinite;*/
    box-shadow: var(--def-shadow);
  }

  div {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    justify-content: right;
    gap: 15px;
  }

  h3 {
    font-weight: 900;
    font-size: 20px;
  }
</style>

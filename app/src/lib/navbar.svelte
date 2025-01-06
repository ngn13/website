<script>
  import NavbarLink from "./navbar_link.svelte";
  import NavbarSwitch from "./navbar_switch.svelte";

  import { color } from "$lib/util.js";
  import { _ } from "svelte-i18n";
</script>

<nav style="border-bottom: solid 2px var(--{color()});">
  <h3 style="color: var(--{color()})">[ngn.tf]</h3>
  <div>
    <NavbarLink link="/">{$_("navbar.home")}</NavbarLink>
    <NavbarLink link="/services">{$_("navbar.services")}</NavbarLink>
    <NavbarLink link="/donate">{$_("navbar.donate")}</NavbarLink>
    <NavbarSwitch />
  </div>
</nav>

<style>
  nav {
    background: var(--black-1);
    padding: 20px 30px 20px 20px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
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
    font-size: var(--size-4);
  }
</style>

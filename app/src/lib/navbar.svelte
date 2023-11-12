<script>
  import NavbarLink from "./navbar_link.svelte";
</script>

<nav>
  <div>
    <h3>[ngn]</h3>
  </div>

  <div>
    <NavbarLink link="/">home</NavbarLink>
    <NavbarLink link="/services">services</NavbarLink>
    <NavbarLink link="/blog">blog</NavbarLink>
    <NavbarLink type="icon" link="https://github.com/ngn13/website"></NavbarLink>
  </div>
</nav>

<style>
nav {
  background: var(--dark-one);
  padding: 25px 30px 27px 25px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  border-bottom: solid 1.5px black;
  animation-name: borderAnimation;
  animation-duration: 10s;
  animation-iteration-count: infinite;
  box-shadow: var(--def-shadow);
}

div {
  display: flex;
  overflow: hidden;
  align-items: center;
  justify-content: center;
  gap: 1px;
}

h3 {
  font-weight: 900;
  font-size: 25px;
  color: red;
  animation-name: colorAnimation;
  animation-iteration-count: infinite;
  animation-duration: 10s;
}
</style>


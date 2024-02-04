<script>
  import NavbarLink from "./navbar_link.svelte";
</script>

<nav>
  <div>
    <h3>[ngn.tf]</h3>
  </div>

  <div>
    <NavbarLink link="/">home</NavbarLink>
    <NavbarLink link="/services">services</NavbarLink>
    <NavbarLink link="/blog">blog</NavbarLink>
    <NavbarLink link="/donate">donate</NavbarLink>
    <NavbarLink link="https://stats.ngn.tf">status</NavbarLink>
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
  align-items: center;
  flex-wrap: wrap;
  justify-content: right;
  gap: 15px;
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


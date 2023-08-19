<script>
  import NavbarLink from "./navbar_link.svelte";
</script>

<nav>
  <div>
    <h3>[ngn]</h3>
  </div>

  <div>
    <NavbarLink link="/">home</NavbarLink>
    <NavbarLink link="/projects">projects</NavbarLink>
    <NavbarLink link="/resources">resources</NavbarLink>
    <NavbarLink link="/blog">blog</NavbarLink>
    <NavbarLink link="https://github.com/ngn13/ngn13.fun">source</NavbarLink>
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


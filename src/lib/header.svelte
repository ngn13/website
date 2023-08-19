<script></script>

<header>
  <h1>
    <slot></slot>
  </h1>
</header>

<style>
header {
  background: 
    linear-gradient(rgba(1, 1, 1, 0.96), 
    rgba(11, 11, 11, 0.808)), 
    url("https://files.ngn13.fun/banner.png");
  background-size: 50%;
  width: 100%;
  height: 100%;
}

h1 {
  font-weight: 900;
  font-size: 5.5vw;
  padding: 120px;
  text-align: center;
  color: white;
  text-shadow: var(--text-shadow); 
  text-size-adjust: 80%;
}
</style>

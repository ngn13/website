<script>
  export let subtitle = "";
</script>

<header>
  <h1>
    <slot></slot>
  </h1>
  <h4><c>{subtitle}</c></h4>
</header>

<style>
  header {
    background: linear-gradient(rgba(11, 11, 11, 0.808), rgba(1, 1, 1, 0.96)),
      url("https://files.ngn.tf/banner.png");
    background-size: 50%;
    width: 100%;
    height: 100%;
  }

  h1 {
    font-weight: 900;
    font-size: 500%;
    padding: 120px;
    padding-bottom: 0;
    text-align: center;
    color: white;
    text-shadow: var(--text-shadow);
    text-size-adjust: 80%;
  }

  h4 {
    padding-bottom: 120px;
    font-weight: 600;
    font-size: 200%;
    text-align: center;
    color: white;
    text-size-adjust: 80%;
  }
</style>

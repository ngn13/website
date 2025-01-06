<script>
  export let class_name;
</script>

<main class={class_name}>
  <slot></slot>
</main>

<style>
  main {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    align-items: stretch;

    padding: 50px;
    gap: 28px;
  }

  @media only screen and (max-width: 900px) {
    main {
      flex-direction: column;
    }
  }
</style>

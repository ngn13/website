<template>
  <div>
    <slot></slot>
  </div>
</template>

<script>
export default {}
</script>

<style scoped>
div {
  display: flex;
  flex-direction: row;
  padding: 15px;
  align-items: center;
  justify-content: center;
  gap: 30px;
}

@media only screen and (max-width: 1416px) {
  div {
    flex-direction: column;
  }
}
</style>

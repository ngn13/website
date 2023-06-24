<template>
  <button @click="click"><slot></slot></button>
</template>

<script>
export default {
  props: ["click"]
}
</script>

<style scoped>
button {
  text-align: center;
  width: 540px;
  font-size: 25px;
  padding: 20px;
  border-radius: 20px;
  background: var(--dark-two);
  border: none;
  color: white;
  outline: none;
  cursor: pointer;
  transition: 0.4s;
}

button:hover {
  box-shadow: var(--def-shadow);
}
</style>

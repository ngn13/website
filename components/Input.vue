<template>
  <input v-on:keyup="keyup" :placeholder="placeholder" :type="type" />
</template>

<script>
export default {
  props: ["keyup", "placeholder", "type"]
}
</script>

<style scoped>
input {
  width: 500px;
  font-size: 25px;
  padding: 20px;
  border-radius: 10px;
  background: var(--dark-two);
  border: none;
  color: white;
  outline: none;
  transition: 0.4s;
}

input:focus {
  box-shadow: var(--def-shadow);
}
</style>

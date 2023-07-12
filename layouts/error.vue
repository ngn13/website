<template>
</template>

<script>
export default {
  props: ["error"],
  layout: "error",
  mounted() {
    this.$router.push({ path: "/" })
  }
}
</script>

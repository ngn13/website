<template>
  <div class="header">
    <h1>
      <slot></slot>
    </h1>
  </div>
</template>

<script>
export default {}
</script>

<style scoped>
div {
  background: linear-gradient(rgba(3, 3, 3, 0.706), rgba(22, 22, 22, 0.808)),
    url("https://www.sketchappsources.com/resources/source-image/tv-glitch-sureshmurali29.png");
  width: 100%;
  height: 100%;
}

h1 {
  font-weight: 900;
  font-size: 7.5vw;
  padding: 150px;
  text-align: center;
  color: white;
  text-shadow: 3px 4px 7px rgba(81, 67, 21, 0.8);
  text-size-adjust: 80%;
}
</style>

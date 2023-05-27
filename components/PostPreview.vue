<template>
<nuxt-link :to="url">
    <h1>{{ title }}</h1>
    <p>{{ info }}</p>
    <h2>{{ desc }}</h2> 
</nuxt-link>
</template>

<script>
export default {
    props: ["title", "desc", "info", "url"],
}
</script>

<style scoped>
a{
    padding: 30px;
    color: white;
    background: var(--dark-two);
    cursor: pointer;
    text-decoration: none;
    transition: .4s;
    width: 70%;
}

a:hover{
    box-shadow: var(--def-shadow);
}

h1{
    font-size: 40px;
    margin-bottom: 5px;
    animation-name: colorAnimation;
    animation-iteration-count: infinite;
    animation-duration: 10s;
    text-shadow: none;
}

h2{
    margin-top: 10px;
}
</style>
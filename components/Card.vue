<template>
    <div>
        <slot></slot>
    </div>
</template>

<script>
export default {

}
</script>

<style scoped>
div {
    box-shadow: var(--def-shadow);
    background: var(--dark-two);
    width: 100%;
    height: auto;
    padding: 50px;
    color: white;
}

h1 {
    text-shadow: 3px 4px 7px rgba(81, 67, 21, 0.8);
    font-size: 50px;
    margin-bottom: 30px;
    text-decoration: underline;
}

i {
    animation-name: colorAnimation;
    animation-duration: 10s;
    animation-iteration-count: infinite;
}

a:hover {
    font-weight: 900;
}

i {
    font-size: 30px;
    margin-top: 3px;
}

h2 {
    font-size: 40px;
    line-height: 60px;
}

a {
    font-size: 40px;
    color: white;
    line-height: 60px;
    text-decoration: none;
}
</style>
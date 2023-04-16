<template>
    <p>
        #<slot></slot>
    </p>
</template>

<script>
export default {

}
</script>

<style scoped>
    p{
        background: var(--dark-three);
        color: white;
        text-shadow: 1px 1px 2px white;
        padding: 5px 10px 5px 10px;
        font-size: 25px;
        border-radius: 7px;
        margin-top: 10px;
        transition: .4s;
    }
</style>
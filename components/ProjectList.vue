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
div{
    display: flex;
    flex-direction: row;
    padding: 15px;
    justify-content: center;
    gap: 30px;
}

@media only screen and (max-width: 1121px) {
    div {
        flex-direction: column;
    }
}
</style>
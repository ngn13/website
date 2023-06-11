<template>
    <div>
        <h1>Currently logged in</h1>
        <Button :click="click">Logout</Button>
    </div>
</template>

<script>
import axios from 'axios';
import Button from './Button.vue';

export default {
    methods: {
        async click(e) {
            await axios.get(`/api/auth/logout?token=${localStorage.getItem("token")}`)
            localStorage.clear()
            location.reload()
        }
    },
}
</script>

<style scoped>
h1{
    color: var(--white);
    font-size: 50px;
    margin-bottom: 20px;
    text-align: center;
}

div{
    background-color: var(--dark-three);
    padding: 50px;
    margin-top: 50px;
    display: flex;
    flex-direction: column;
    gap: 20px;
    align-items: center;
    justify-content: center;
}
</style>

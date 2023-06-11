<template>
    <div>
        <h1>Login Page</h1>
        <Input :keyup="function() { }" placeholder="Password" type="password" id="pass"/>
        <Button :click="click">Login</Button>
    </div>
</template>

<script>
import Input from '../components/Input.vue';
import Button from '../components/Button.vue';
import axios from "axios";

export default {
    methods: {
        async click(e) {
            const pass = document.getElementById("pass").value
            const res = await axios.get(`/api/auth/login?pass=${pass}`)
            if(res.data["error"]===0){
                localStorage.setItem("token", res.data["token"])
                return location.href="/"
            }
            alert("Incorrect password!")
        }
    }
}
</script>

<style scoped>
div {
    padding: 50px;
    background: var(--dark-three);
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    margin: auto;
    display: flex;
    flex-direction: column;
    gap: 20px;
    color: var(--white);
    align-items: center;
    justify-content: center;
}

h1{
    font-size: 70px;
    margin-bottom: 20px;
}
</style>

<template>
    <div>
        <h1>Add Project</h1>
        <Input :keyup="function() { }" id="name" placeholder="Project Name" type="text"/>
        <Input :keyup="function() { }" id="desc" placeholder="Project Desc" type="text"/>
        <Input :keyup="function() { }" id="url" placeholder="Project URL" type="text"/>
        <Button :click="click">Post</Button>
    </div>
</template>

<script>
import axios from 'axios';
import Input from './Input.vue';
import Button from './Button.vue';

export default {
    methods: {
        async click(e) {
            const name = document.getElementById("name").value
            const desc = document.getElementById("desc").value
            const url = document.getElementById("url").value
            const token = localStorage.getItem("token")
            const res = await axios.get(`/api/projects/add?token=${token}&name=${name}&desc=${desc}&url=${url}`)
            if(res.data["error"]!==0)
                return alert("Error!")
            alert("Project added!")
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

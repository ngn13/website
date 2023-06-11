<template>
    <div>
        <h1>Add Resource</h1>
        <Input :keyup="function(){}" id="name" placeholder="Resource Name" type="text"/>
        <Input :keyup="function(){}" id="tags" placeholder="Resource Tags (comma seperated)" type="text"/>
        <Input :keyup="function(){}" id="url" placeholder="Resource URL" type="text"/>
        <Button :click="click">Post</Button>
    </div>
</template>

<script>
import Input from './Input.vue';
import Button from './Button.vue';
import axios from 'axios';

export default {
    methods: {
        async click(e) {
            const name = document.getElementById("name").value
            const tags = document.getElementById("tags").value
            const url = document.getElementById("url").value
            const token = localStorage.getItem("token")
            const res = await axios.get(`/api/resources/add?token=${token}&name=${name}&tags=${tags}&url=${url}`)
            if(res.data["error"]!==0)
                return alert("Error!")
            alert("Resource added!")
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

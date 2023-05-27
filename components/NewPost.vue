<template>
    <main>
        <h1>Add New Post</h1>
        <div class="textareas">
            <Input :keyup="function() { }" id="title" placeholder="Post Title" type="text"/>
            <Input :keyup="function() { }" id="author" placeholder="Author" type="text"/>
        </div>
        <div class="textareas">
            <textarea nor name="contenten" id="contenten" cols="30" rows="10" placeholder="Content (EN)"></textarea>
            <textarea name="contenttr" id="contenttr" cols="30" rows="10" placeholder="Content (TR)"></textarea>
        </div>
        <Button :click="click">Post</Button>
    </main>
</template>

<script>
import axios from 'axios';
import Input from './Input.vue';
import Button from './Button.vue';

export default {
    methods: {
        async click(e) {
            const title = document.getElementById("title").value
            const author = document.getElementById("author").value
            const contenten = document.getElementById("contenten").value
            const contenttr = document.getElementById("contenttr").value
            const token = localStorage.getItem("token")
            const res = await axios.post("/api/add_post", {
                token: token,
                title: title,
                author: author,
                content: {
                    tr: contenttr,
                    en: contenten
                },
            })
            if(res.data["error"]!==0)
                return alert("Error!")
            alert("Post added!")
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

textarea{
    width: 500px;
    font-size: 15px;
    padding: 20px;
    border-radius: 20px;
    background: var(--dark-two);
    border: none;
    color: white;
    outline: none;
    resize: vertical;
    height: 200px;
    transition: .4s;
}

.textareas {
    flex-direction: row;
    display: flex;
    gap: 20px;
}

textarea:focus {
    box-shadow: var(--def-shadow);
}

main{
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
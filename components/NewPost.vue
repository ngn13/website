<template>
  <main>
    <h1>Add New Post</h1>
    <div class="textareas">
      <Input
        :keyup="function () {}"
        id="title"
        placeholder="Post Title"
        type="text"
      />
      <Input
        :keyup="function () {}"
        id="author"
        placeholder="Author"
        type="text"
      />
      <h2>
        Make the post private
        <input id="private" type="checkbox" />
      </h2>
    </div>
    <textarea
      name="content"
      id="content"
      cols="30"
      rows="10"
      placeholder="Content"
    ></textarea>
    <Button :click="click">Post</Button>
  </main>
</template>

<script>
import axios from "axios"
import Input from "./Input.vue"
import Button from "./Button.vue"

export default {
  methods: {
    async click(e) {
      const title = document.getElementById("title").value
      const author = document.getElementById("author").value
      const content = document.getElementById("content").value
      const priv = document.getElementById("private").value
      const token = localStorage.getItem("token")
      const res = await axios.post("/api/blog/add", {
        token: token,
        title: title,
        author: author,
        content: content,
        priv: priv === "on"
      })
      if (res.data["error"] !== 0) return alert("Error!")
      alert("Post added!")
      location.reload()
    }
  }
}
</script>

<style scoped>
h1 {
  color: var(--white);
  font-size: 50px;
  margin-bottom: 20px;
  text-align: center;
}

h2 {
  background: var(--dark-two);
  font-size: 25px;
  border-radius: 20px;
  border: none;
  padding: 20px;
  color: var(--white);
  display: flex;
  justify-content: space-between;
}

input[type="checkbox"] {
  -ms-transform: scale(2);
  -moz-transform: scale(2);
  -webkit-transform: scale(2);
  -o-transform: scale(2);
  transform: scale(2);
  padding: 10px;
}

textarea {
  width: 500px;
  font-size: 20px;
  padding: 20px;
  border-radius: 20px;
  background: var(--dark-two);
  border: none;
  color: white;
  outline: none;
  resize: vertical;
  height: 200px;
  transition: 0.4s;
}

.textareas {
  flex-direction: column;
  display: flex;
  gap: 20px;
}

textarea:focus {
  box-shadow: var(--def-shadow);
}

main {
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

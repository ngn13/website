<template>
    <div class="all">
        <Navbar />
        <Header>
            <glitch class="title">{{ post.title }}</glitch>
            <p>{{ post.info }}</p>
        </Header>
        <div class="postcontain">
            <main class="markdown-body" v-html="content"></main>
        </div>
    </div>
</template>

<script>
import Navbar from "../../../components/Navbar.vue";
import Header from "../../../components/Header.vue";
import axios from "axios";
import * as DOMPurify from "dompurify";
import marked from "marked";

export default {
    head() {
        return {
            title: "[ngn] | " + this.post.title,
            meta: [
                {
                    hid: "description",
                    name: "description",
                    content: "read my blog posts"
                }
            ]
        };
    },


    data() {
        return {
            post: {},
            lang: "",
            content: "",
        }
    },

    async created() {
        const res = await axios.get(`/api/blog/get?id=${this.$route.params.id}`)
        if (res.data["error"] === 3)
            return this.$router.push({ path: "/blog" })
        this.post = res.data["post"]
        this.content = DOMPurify.sanitize(
          marked.parse(this.post["content"], { breaks: true }),
          { ADD_TAGS: ["iframe"], ADD_ATTR: ['allow', 'allowfullscreen', 'frameborder', 'scrolling'] }
        )
    }

}
</script>

<style scoped>
glitch {
    font-size: 80px;
}

p {
    font-size: 30px;
}

.info {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
}

.postcontain{
    padding: 50px;
}

.markdown-body {
    font-size: 25px;
    padding: 50px;
    border-radius: 15px;
    background-color: var(--dark-three);
}
</style>

<style>
.markdown-body{
    font-family: "Ubuntu", sans-serif;
}

.markdown-body h1{
  border-bottom: 1px solid #505050;
}

.markdown-body iframe{
  display: block;
  margin: 20px 0px;
}

.markdown-body a{
  animation-name: colorAnimation;
  animation-iteration-count: infinite;
  animation-duration: 10s;
  text-shadow: none;
}

</style>

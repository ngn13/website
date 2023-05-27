<template>
    <div class="all">
        <Navbar />
        <Header>
            <glitch class="title">{{ post.title }}</glitch>
            <p>{{ post.info }}</p>
        </Header>
        <div class="postcontain">
            <div class="lang">
                <button v-if="lang==='en'" class="bl bs">EN <span class="fi fi-gb"></span></button>
                <button v-on:click="toggle_lang" v-else class="bl">EN <span class="fi fi-gb"></span></button>
                <button v-if="lang==='tr'" class="br bs">TR <span class="fi fi-tr"></span></button>
                <button v-on:click="toggle_lang" v-else class="br">TR <span class="fi fi-tr"></span></button>
            </div>
            <main class="postself" v-if="lang==='en'" v-html="en"></main>
            <main class="postself" v-else v-html="tr"></main>
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

    methods: {
        toggle_lang(){
            this.lang = this.lang==="tr" ? "en" : "tr"
            localStorage.setItem("lang", this.lang)
        }
    },

    data() {
        return {
            post: {},
            lang: "",
            en: "",
            tr: "",
        }
    },

    async created() {
        const res = await axios.get(`/api/get_post?id=${this.$route.params.id}`)
        if (res.data["error"] === 3)
            return this.$router.push({ path: "/blog" })
        this.post = res.data["post"]
        this.en = DOMPurify.sanitize(marked.parse(this.post["content"]["en"]))
        this.tr = DOMPurify.sanitize(marked.parse(this.post["content"]["tr"]))
        this.lang = localStorage.getItem("lang")!==undefined&&(localStorage.getItem("lang")==="tr"||localStorage.getItem("lang")==="en") ? localStorage.getItem("lang") : "en"
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

span{
    width: 30px;
    border-radius: 5px;
}

.info {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
}

.lang {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    margin-top: 20px;
}

button{
    padding: 10px 30px 10px 30px;
    font-size: 20px;
    background-color: var(--dark-two);
    border: none;
    cursor: pointer;
    color: white
}
.bs{
    background-color: var(--dark-three);
}
.bl{
    border-radius: 15px 0px 0px 0px;
}

.br{
    border-radius: 0px 15px 0px 0px;
}

.postcontain{
    padding: 50px;
}

.postself {
    text-align: left;
    font-size: 30px;
    color: white;
    padding: 20px 40px 70px 40px;
    line-height: 40px;
    border-radius: 15px;
    background-color: var(--dark-three);
}
</style>

<style>
.postself code {
    background: var(--dark-two);
    border-radius: 5px;
    font-size: 18px;
    padding: 5px;
    font-style: italic;
}

.postself h1{
    margin-top: 70px;
    margin-bottom: 20px;
}

.postself h2{
    margin-top: 60px;
    margin-bottom: 20px;
}

.postself h3{
    margin-top: 50px;
    margin-bottom: 20px;
}

.postself h4{
    margin-top: 40px;
    margin-bottom: 20px;
}

.postself h5{
    margin-top: 60px;
    margin-bottom: 30px;
}

.postself a{
    animation-name: colorAnimation;
    animation-iteration-count: infinite;
    animation-duration: 10s;
    text-shadow: none;
}

li{
    list-style-type: none;  
}
</style>
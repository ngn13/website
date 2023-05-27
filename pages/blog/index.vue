<template>
    <div>
        <Navbar />
        <Header>
            <glitch>/dev/</glitch>blog
        </Header>
        <div class="blogs">
            <Input :keyup="keyup" placeholder="Search post" type="text"/>
            <PostPreview v-for="post in posts" :key="post" :title="post.title" :desc="post.desc" :info="post.info" :url="post.url">
                {{ post.desc }}
            </PostPreview>
        </div>
        <NewPost v-if="logged"/>
    </div>
</template>

<script>
import Navbar from "../../components/Navbar.vue";
import Header from "../../components/Header.vue";
import NewPost from "../../components/NewPost.vue";
import PostPreview from "../../components/PostPreview.vue";
import axios from "axios";

export default {
    head() {
        return {
            title: "[ngn] | blog",
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
            logged: false,
            posts: [],
            all: []
        };
    },
    mounted: async function () {
        if (localStorage.getItem("token"))
            this.logged = true;
        const res = await axios.get("/api/get_posts");

        let array = res.data["posts"]
        let newarray = []

        for(let i=0;i<array.length;i++){
            newarray.push({
                title: array[i].title,
                desc: array[i].desc,
                info: array[i].info,
                url: `/blog/${array[i].title.toLowerCase().replaceAll(" ", "")}`
            })
        }

        this.posts = newarray;
        this.all = newarray;
    },
    methods: {
        keyup(e) {
            let val = e.target.value
            
            // search looks at name and info
            this.posts = []
            for(let i = 0; i < this.all.length; i++){
                if(this.all[i].title.toLowerCase().includes(val.toLowerCase()))
                    this.posts.push(this.all[i])
                
                else if(this.all[i].info.toLowerCase().includes(val.toLowerCase()))
                    this.posts.push(this.all[i])
            }
        }
    },
}
</script>

<style scoped>
.blogs {
    padding: 50px;
    gap: 35px;
    display: flex;
    flex-direction: column;
    gap: 30px;
    align-items: center;
}
</style>
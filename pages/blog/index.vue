<template>
    <div>
        <Navbar />
        <Header>
          <label class="glitch">/dev/</label>blog
        </Header>
        <div class="blogs">
            <Input :keyup="keyup" placeholder="Search post" type="text"/>
            <PostPreview v-for="post in posts" :key="post.title" :title="post.title" :desc="post.desc" :info="post.info" :url="post.url">
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
        const res = await axios.get("/api/blog/sum");
        let posts = []

        res.data["posts"].forEach(post=>{
            posts.push({
                title: post.title,
                desc: post.desc,
                info: post.info,
                url: `/blog/${post.title.toLowerCase().replaceAll(" ", "")}`
            })
        })

        this.posts = posts
        this.all = posts
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

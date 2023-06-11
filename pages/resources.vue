<template>
    <main>
        <Navbar />
        <Header>
            <glitch>cat</glitch> {{ header }}
        </Header>
        <div class="resources">
            <Input :keyup="keyup" placeholder="Search resource" type="text"/>
            <Resource v-for="res in show_resources" :key="res.name" :name="res.name" :tags="res.tags" :url="res.url" />
        </div>
        <NewResource v-if="logged"/>
    </main>
</template>

<script>
import axios from 'axios';
import Resource from '../components/Resource.vue';
import Input from '../components/Input.vue';
import NewResource from '../components/NewResource.vue';

export default {
    head() {
        return {
            title: "[ngn] | resources",
            meta: [
                {
                    hid: "description",
                    name: "description",
                    content: "discover new resources"
                }
            ]
        }
    },
    data() {
        return {
            header: "resources",
            logged: false,
            sum_resources: [],
            all_resources: [],
            show_resources: []
        }
    },
    methods: {
        keyup(e) {
            let search = e.target.value
            if(e.key==="Backspace" && search===""){
                this.header = "resources"
                this.show_resources = this.sum_resources
                return
            }
            if(e.key==="OS")
                return
            this.header = `resources | grep ${search}`

            // dirty asf search alg
            this.show_resources = []
            for(let i = 0; i < this.all_resources.length; i++){
                if(this.all_resources[i].name
                      .toLowerCase()
                      .includes(search.toLowerCase())
                  ){
                    this.show_resources.push(this.all_resources[i])
                    continue
                  }

                for(let e = 0; e < this.all_resources[i].tags.length; e++){
                    if(this.all_resources[i].tags[e].toLowerCase()
                          .includes(search.toLowerCase())
                    ){
                          this.show_resources.push(this.all_resources[i])
                          break
                    }
                }
            }
      }
    },
    mounted: async function(){
    		if(localStorage.getItem("token"))
            this.logged = true

        // request top 10 resources so we can
        // render the DOM as fast as possible
        let res = await axios.get("/api/resources/get?sum=1")
        this.sum_resources = res.data["resources"]
        this.show_resources = this.sum_resources

        // then we can load all the resources
        res = await axios.get("/api/resources/get")
        this.all_resources = res.data["resources"]
        console.log(res.data["resources"])
	}
}
</script>

<style scoped>
.resources {
    padding: 50px;
    padding-bottom: 60px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 40px
}
</style>

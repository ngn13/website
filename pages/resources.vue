<template>
    <main>
        <Navbar />
        <Header>
            <glitch>cat</glitch> {{ header }}
        </Header>
        <div class="resources">
            <Input :keyup="keyup" placeholder="Search resource" type="text"/>
            <Resource v-for="resource in resources" :key="resource" :name="resource.name" :tags="resource.tags" :url="resource.url" />
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
    data() {
        return {
            header: "resources",
            logged: false,
            resources: [],
            all: []
        }
    },
    methods: {
        keyup(e) {
            let val = e.target.value
            if(e.key==="Backspace" && val===""){
                this.header = "resources"
                this.resources = this.all
                return
            }
            if(e.key==="OS")
                return
            this.header = `resources | grep ${val}`

            // dirty asf search alg
            this.resources = []
            for(let i = 0; i < this.all.length; i++){
                if(this.all[i].name.toLowerCase().includes(val.toLowerCase()))
                    this.resources.push(this.all[i])
            
                for(let e = 0; e < this.all[i].tags.length; e++){
                    if(this.all[i].tags[e].toLowerCase().includes(val.toLowerCase())){
                        if(this.resources.indexOf(this.all[i])===-1)
                            this.resources.push(this.all[i])
                    }
                }
            }
        }
    },
    mounted: async function(){
		if(localStorage.getItem("token"))
            this.logged = true

        const res = await axios.get("/api/get_resources")
        this.resources = res.data["resources"]
        this.all = res.data["resources"]
	}
}
</script>

<style scoped>
.resources {
    padding: 50px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 40px
}
</style>
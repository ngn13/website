<template>
    <div>
        <Navbar />
	    <Header>
		    <glitch>ls -la</glitch> projects
	    </Header>
        <div class="projects">
            <ProjectList v-for="project in projects" :key="project">
                <Project v-if="logged" v-for="p in project" :key="p" :name="`${p.name} (${p.click})`" :desc="p.desc" :url="p.url"/>
                <Project v-if="!logged" v-for="p in project" :key="p" :name="p.name" :desc="p.desc" :url="p.url"/>
            </ProjectList>
	    </div>
        <NewProject v-if="logged"/>
    </div>
</template>

<script>
import ProjectList from "../components/ProjectList.vue";
import Project from "../components/Project.vue";
import NewProject from "../components/NewProject.vue";
import axios from "axios";

export default {
    head() {
        return {
            title: "[ngn] | projects",
            meta: [
                {
                    hid: "description",
                    name: "description",
                    content: "check out my projects"
                }
            ]
        }
    },
    data() {
		return {
		    logged: false,
        projects: []
		}
	},
    mounted: async function(){
		if(localStorage.getItem("token"))
            this.logged = true

        const res = await axios.get("/api/projects/get")

        let all = res.data["projects"]
        let projects = []
        let project = []
        for(let i = 0; i<all.length; i++){
            if(project.length!==3)
                project.push({
                    name: all[i]["name"],
                    desc: all[i]["desc"],
                    click: all[i]["click"],
                    url: `/l/${all[i]["name"].toLowerCase().replaceAll(" ", "")}`
                })
            else{
                projects.push(project)
                project = []
            }

            if(i===all.length-1){
                projects.push(project)
            }
        }
        this.projects = projects
	}
}
</script>

<style>
.projects{
    padding: 50px;
    display: flex;
    flex-direction: column;
}

@media only screen and (max-width: 1121px) {
    .projects {
        padding: 50px;
    }
}
</style>

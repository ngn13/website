<template>
    <div>
        <Navbar />
        <Header>
            <label class="glitch">echo</label> hello world!
        </Header>
        <div class="info">
            <Card>
                <h1>👋 Welcome to my website!</h1>
                <h2>
                    I am a high school student who is interested in
                    <br>
                    • cyber security
                    <br>
                    • coding
                    <br>
                    • electronics
                    <br>
                    • gaming
                    <br>
                    • simply: everything about computers!
                </h2>
            </Card>
            <Card>
                <h1>👉 Contact me</h1>
                <h2>You can contact me on the following platforms:</h2>

                <a href="https://discord.com/users/568131907368386565"><i class='bx bxl-discord-alt'></i> Discord</a>
                <br>
                <a href="https://github.com/ngn13"><i class='bx bxl-github'></i> Github</a>
                <br>
                <a href="https://mastodon.social/@ngn"><i class='bx bxl-mastodon'></i> Mastodon</a>
                <br>
                <a href="mailto:ngn13proton@proton.me"><i class='bx bxs-envelope'></i> Mail</a>
                <br>
                <h2>or private message me on matrix:</h2>
                <a><i>[matrix]</i> @ngn:matrix.ngn13.fun</a>
            </Card>
        </div>
        <Logout v-if="logged"/>
        <div class="version">
            <p>v2.3</p>
        </div>
    </div>
</template>

<script>
import Navbar from "../components/Navbar.vue";
import Header from "../components/Header.vue";
import Card from "../components/Card.vue"
import Logout from "../components/Logout.vue";

export default {
    head() {
        return {
            title: "[ngn]",
            meta: [
                {
                    hid: "description",
                    name: "description",
                    content: "homepage of my website"
                }
            ]
        }
    },
    data() {
		return {
			logged: false
		}
	},
    mounted(){
		if(localStorage.getItem("token"))
            this.logged = true
	}
}
</script>

<style scoped>
.info {
    padding: 50px;
    gap: 30px;
    display: flex;
}
.version{
    color: var(--dark-fife);
    position: fixed;
    bottom: 10px;
    right: 10px;
    font-size: 15px;
}

i{
    font-style: normal;
}

@media only screen and (max-width: 1076px) {
    .info {
        flex-direction: column;
    }
    .info div {
        width: auto;
    }
}
</style>

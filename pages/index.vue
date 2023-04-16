<template>
    <div>
        <Navbar />
        <Header>
            <glitch>echo</glitch> hello world!
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
                    • simply: everything about computers!
                </h2>
            </Card>
            <Card>
                <h1>👉 Checkout my socials</h1>
                <h2>Here you can find my socials for following platforms:</h2>
                <br>
                <a href="https://discord.com/users/568131907368386565"><i class='bx bxl-discord-alt'></i> Discord</a>
                <br>
                <a href="https://github.com/ngn13"><i class='bx bxl-github'></i> Github</a>
                <br>
                <a href="https://www.reddit.com/user/_ngn"><i class='bx bxl-reddit'></i> Reddit</a>
                <br>
                <a href="https://fosstodon.org/@ngn"><i class='bx bxl-mastodon'></i> Fosstodon</a>
            </Card>
        </div>
        <Logout v-if="logged"/>
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
                    content: "ngn's Personal Website | Home Page"
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


@media only screen and (max-width: 1076px) {
    .info {
        flex-direction: column;
    }
    .info div {
        width: auto;
    }
}
</style>
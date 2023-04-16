<template>
	<div>
		<nuxt-link v-if="!out" :class="cname" id="link" :to="url">
			<slot></slot>
		</nuxt-link>
		<a v-if="out" :class="cname" id="link" :href="url">
			<slot></slot>
		</a>
	</div>
</template>

<script>
export default {
	props: ["url", "out"],
	data() {
		return {
			cname: "notcurrent"
		}
	},
	mounted(){
		let url = window.location.pathname
		if(url===this.url){
			this.cname = "current"
		}
	}
}
</script>

<style scoped>
.notcurrent {
	margin-left: 20px;
	font-weight: 700;
	font-size: 25px;
	text-decoration: none;
	color: white;
	cursor: pointer;
}

.notcurrent:hover {
	text-decoration: underline;
	animation-name: underlineAnimation;
	animation-duration: 5s;
	animation-iteration-count: infinite;
	text-shadow: 3px 4px 7px rgba(81, 67, 21, 0.8);
}

.current{
	margin-left: 20px;
	font-weight: 700;
	font-size: 25px;
	text-decoration: none;
	color: white;
	cursor: pointer;
	text-decoration: underline;
	animation-name: underlineAnimation;
	animation-duration: 10s;
	animation-iteration-count: infinite;
	text-shadow: 3px 4px 7px rgba(81, 67, 21, 0.8);
}
</style>
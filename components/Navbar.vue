<template>
	<nav>
		<div>
			<h3>[ngn]</h3>
		</div>
		<div>
			<NavbarLink :out="false" url="/">Home</NavbarLink>
			<NavbarLink :out="false" url="/projects">Projects</NavbarLink>
			<NavbarLink :out="false" url="/resources">Resources</NavbarLink>
			<!--<NavbarLink :out="false" url="/stats">Stats</NavbarLink>-->
			<NavbarLink :out="true" url="http://github.com/ngn13/my-website">Source</NavbarLink>
		</div>
	</nav>
</template>

<script>
import NavbarLink from "./NavbarLink.vue";

export default {}
</script>

<style scoped>
nav {
	background: var(--dark-two);
	padding: 25px 35px 30px 35px;
	display: flex;
	flex-direction: row;
	align-items: center;
	justify-content: space-between;
	border-bottom: solid 3px black;
	animation-name: borderAnimation;
	animation-duration: 10s;
	animation-iteration-count: infinite;
	box-shadow: var(--def-shadow);
}

div {
	display: flex;
	overflow: hidden;
	align-items: center;
	justify-content: center;
	gap: 5px;
}

h3 {
	font-weight: 900;
	font-size: 30px;
	color: red;
	animation-name: colorAnimation;
	animation-iteration-count: infinite;
	animation-duration: 10s;
}
</style>
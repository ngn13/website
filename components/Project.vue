<template>
    <main @click="redirect()">
        <i class='bx bx-code-alt'></i>
        <div>
            <h1>{{ name }}</h1>
            <h2>{{ desc }}</h2>
        </div>
    </main>
</template>

<script>
export default {
    props: ["name", "desc", "url"],
    methods: {
        redirect(e) {
            location.href=this.url
        }
    },
}
</script>

<style scoped>
main {
    color: var(--white);
    background: var(--dark-two);
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 20px;
    padding: 40px;
    cursor: pointer;
    transition: .4s;
    height: 100px;
    width: 450px;
}

main:hover{
    box-shadow: var(--def-shadow);
}

i{
    font-size: 65px;
    animation-name: colorAnimation;
    animation-duration: 10s;
    animation-iteration-count: infinite;
}

h1{
    font-size: 35px;
    margin-bottom: 10px;
}

h2{
    font-size: 25px;
}

@media only screen and (max-width: 1416px) {
    main{
        width: 80%;
    }
}
</style>

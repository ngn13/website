<template>
    <main @click="redirect()" class="mn">
        <div class="resource">
            <h1>{{ name }}</h1>
            <div class="tags">
                <Tag v-for="tag in tags" :key="tag">{{ tag }}</Tag>
            </div>
        </div>
        <i class='bx bx-right-arrow-alt' ></i>
    </main>
</template>

<script>
import Tag from './Tag.vue';

export default {
    props: ["tags", "name", "url"],
    methods: {
        redirect(e){
            location.href = this.url
        }
    }
}
</script>

<style scoped>
main{
    background: var(--dark-two);
    padding: 30px 40px 30px 40px;
    display: flex;
    flex-direction: row;
    color: var(--white);
    justify-content: space-between;
    align-items: center;
    transition: .4s;
    width: 80%;
    cursor: pointer;
}

main:hover{
    box-shadow: var(--def-shadow);
}

.mn:hover i{
    color: white;
}

.resource{
    display: flex;
    flex-direction: column;
    gap: 10px;
}

.tags{
    display: flex;
    flex-direction: row;
    gap: 10px;
}

i{
    font-size: 70px;
    cursor: pointer;
    color: var(--dark-two);
    transition: .4s;
}
</style>
<template>
    <div class="container">
        <div class="manage-card" v-for="(doc, index) in list" :key="index">
            <h2>{{doc.title}}</h2>
            {{doc.abstract}}
        </div>
    </div>
</template>

<script>
    export default {
        name: "docs-manager",
        data() {
            return{
                list: [],
                total: 0
            }
        },
        async mounted() {
            let data = await this.$_getDocsList({start: "/doca", length: 10});
            this.list = data.data.list;
            this.total = data.data.total;
        }
    }
</script>

<style scoped>
    .container{
        width: 100%;
        margin: 0.5em;
    }
</style>
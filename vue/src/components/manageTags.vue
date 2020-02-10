<template>
    <div id="manage-tags">
        <painter-tag v-for="t in $store.state.headerTags" :key="t">{{t}}</painter-tag>
    </div>
</template>

<script>
    import PainterTag from "./tag";
    export default {
        name: "tags-manager",
        components: {PainterTag},
    }
</script>

<style scoped>
    #manage-tags{

    }
</style>
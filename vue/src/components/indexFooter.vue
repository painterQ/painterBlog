<template>
    <div class="footer">
        <span><slot></slot></span>
    </div>
</template>

<script>
    export default {
        name: "index-footer"
    }
</script>

<style scoped>
    .footer{
        padding: 1em;
        text-align: center;
        width: 100%;
        font-family: -apple-system, "Microsoft YaHei", 'Impact', 'Charcoal', sans-serif;
    }
</style>
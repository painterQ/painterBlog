<template>
    <div>
        <index-header></index-header>
        <router-view class="index-layout-main"></router-view>
    </div>
</template>

<script>
    import indexHeader from "@/components/indexHeader";

    export default {
        name: 'layout',
        components: {
            indexHeader,
        },
        props: {
            msg: String
        },
        data: function () {
            return {
                asideSwitch: false,
            }
        },

        //learn: 计算属性
        computed: {},
    }
</script>

<style scoped>

    .index-layout-main{
        max-width: 1024px;
        margin: 10px auto;
    }
    /* learn: calc()函数，css3的函数
    attr()，返回选择元素的属性值
    liner-gradient()
    radial-gradient()
    */

</style>
<!-- Add "scoped" attribute to limit CSS to this component only -->

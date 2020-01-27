<template>
    <div>
        <index-header></index-header>
        <router-view class="index-layout-main"></router-view>
        <el-backtop><i class="el-icon-caret-top"
                        style="color: orangered"></i></el-backtop>
    </div>
</template>

<script>
    import indexHeader from "@/components/indexHeader";
    import {Backtop} from "element-ui"
    import Vue from "vue"

    Vue.use(Backtop)
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

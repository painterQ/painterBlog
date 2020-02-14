<template>
    <div>
        <index-header></index-header>
        <router-view class="index-layout-main"></router-view>
        <el-backtop right="200" class="back-top-ani"><i class="el-icon-caret-top"></i></el-backtop>
        <index-footer>{{this.$store.state.ipc}}</index-footer>
    </div>
</template>

<script>
    import indexHeader from "@/components/indexHeader";
    import {Backtop} from "element-ui"
    import Vue from "vue"
    import indexFooter from "./indexFooter";
    Vue.use(Backtop)
    export default {
        name: 'layout',
        components: {
            indexHeader,
            indexFooter
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
    @keyframes backtop {
        16% {color: Blue}
        32% {color: BlueViolet}
        48% {color: Red}
        64% {color: OrangeRed}
        80% {color: Tomato}
        100% {color: OliveDrab}
    }
    .back-top-ani{
        animation: backtop 5s infinite;
        font-size: 1.5em;
    }
</style>
<!-- Add "scoped" attribute to limit CSS to this component only -->

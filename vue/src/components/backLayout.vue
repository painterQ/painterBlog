<!--
使用vw和vh做单位可以全屏，另一种方法是设置html宽度和高度为100%
el-container el-aside el-header el-footer
--->
<template>
    <el-container class="container">
        <el-header class="header">
            <painter-header></painter-header>
            <painter-login></painter-login>
        </el-header>
        <el-container class="body">
            <!--learn: 设置为auto，设置了overflow，便可以由包裹性，由内部决定--->
            <el-aside class="aside" ref="aside" width="auto">
                <painter-aside></painter-aside>
            </el-aside>
            <el-main class="main">
                <router-view></router-view>
            </el-main>
        </el-container>
    </el-container>
</template>

<script>
    import painterAside from "@/components/aside"
    import painterHeader from "@/components/header"
    import painterLogin from "@/components/login"

    export default {
        name: 'layout',
        components: {
            painterHeader,
            painterAside,
            painterLogin,
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
    /*
    learn: 关于包裹性
    如果不设置height和wight的话，block的高度是内部决定的，宽度是尽量充满
    如果指定width，百分比和长度，就会固定宽度，这时候设置margin可以水平居中
    如果指定max-width，那么...

    块格式化上下文 BFC 脱离了文档流的独立块区域
    比如浮动的元素，float不是none；由绝对定位的元素，position是absolute或者fix；
    inline-block，table-cell，flex，table-acption，inline-flex
    overflow的值不是visible
    BFC中，除了flex，都有包裹性

    overflow，规定当内部溢出的时候怎么办，默认是visible
    hidden，裁剪；scroll，滚动条；auto 必要时滚动条；
    overflow-x, 多了no-display和no-content
    overflow是大小不够了才会触发，够的时候不会,够的时候只会收缩包裹
    */
    .container {
        height: 100vh;
        overflow: hidden;
    }

    /* learn: calc()函数，css3的函数
    attr()，返回选择元素的属性值
    liner-gradient()
    radial-gradient()
    */
    .body {
        height: calc(100vh - 50px);
        overflow: hidden;
    }

    .header {
        padding: 0;
        height: 50px;
    }

    .main {
        padding: 0;
        background-color: #fafafa;
        overflow-y: auto;
        overflow-x: hidden;
        box-sizing: border-box;

    }

    .aside {
        position: relative;
        overflow-x: hidden;
        overflow-y: auto;
    }

    #login {
        width: 70%;
        margin: 0 auto;
        font-family: 微软雅黑 Sans-serif;
        font-weight: bold;
    }
</style>
<!-- Add "scoped" attribute to limit CSS to this component only -->

<template>
    <div class="index-layout-header">
        <div class="index-header-bar" v-if="showBar">
            <div class="index-header-bar-inner-wrapper">
                <el-avatar
                        shape="circle"
                        :size="40"
                        fit="contain"
                        :src="avatar"
                        @error="avatarError"></el-avatar>
                <router-link to="/list" class="bar_fount">关于我</router-link>
                <router-link to="/list" class="bar_fount">标签</router-link>
                <router-link to="/list" class="bar_fount">目录</router-link>
            </div>
            <el-form class="index-header-bar-search">
                <el-input placeholder="搜索..." prefix-icon="el-icon-search"></el-input>
            </el-form>
        </div>
        <div class="index-header-title-all">
            <div class="index-header-title-center">
                <painter-tag v-for="t in $store.state.headerTags" :key="t">{{t}}</painter-tag>
                <h1 class="index-header-title mix">{{$store.state.headerTitle}}</h1>
                <h2 class="index-header-subtitle mix">{{$store.state.headerSubTitle}}</h2>
                <span class="meta mix">{{$store.state.headerName}} on {{new Date($store.state.headerTime) | moment}}</span>
            </div>
        </div>
    </div>
</template>

<script>
    import PainterTag from "./tag";
    import vue from 'vue'
    import {Avatar, Form, Input} from 'element-ui'

    vue.use(Avatar);
    vue.use(Form);
    vue.use(Input);

    export default {
        name: "index-header",
        components: {PainterTag},
        data: function () {
            return {
                lastScroll: 0,
                showBar: true
            }
        },
        methods: {
            menu() {
                let scroll = document.documentElement.scrollTop || document.body.scrollTop;
                this.showBar = scroll <= this.lastScroll || scroll === 0;
                this.lastScroll = scroll
            },
            up2Top() {
                //回到顶部
                document.documentElement.scrollTop = 0;
                document.body.scrollTop = 0;
            },
            avatarError() {
                console.log("avatar Error");
                return true
            },
        },
        watch: {
            //如果没有immediate，避免了部分组件内路由，但是mounted还是不能省的，否则刷新不会触发
            '$route.fullPath': {
                handler: function (newFlag, /*oldFlag*/) {
                    this.$store.dispatch('setCurrentPath', newFlag);
                    this.up2Top();
                },
                //立即触发，可以省略很多mounted，因为这里刷新也会触发了
                //最佳实践就是，新建一个总会存在的components，然后在其中加route的watcher
                immediate: true
            }
        },
        computed: {
            avatar() {
                console.log("get avatar", this.$store.state.authorAvatar);
                return this.$store.state.authorAvatar
            }
        },
        mounted() {
            window.addEventListener('scroll', this.menu, true);
            this.$store.dispatch('InitAsync');
        },
    }
</script>

<style scoped>
    .index-layout-header, .index-header-bar::after {
        background: url("../../public/background.jpg") top / cover fixed;
        color: #fff;
        /*width: 100vw;*/
    }

    .index-header-title-center {
        margin: 2em auto;
        width: 50%;
    }

    .mix{
        mix-blend-mode: difference;
    }

    .meta {
        font-family: 'Lora', 'Times New Roman', serif;
        font-style: italic;
        font-weight: 300;
        font-size: 18px;
    }

    .index-header-bar {
        width: 100%;
        position: fixed;
        top: 0;
        background: rgba(255, 255, 255, .3);
        overflow: hidden;
        justify-content: space-around;
        display: flex;
        flex-wrap: wrap;
    }

    .index-header-bar::after {
        display: block;
        content: '';
        position: absolute;
        height: calc(100% + 40px);
        width: calc(100% + 40px);
        top: -20px;
        left: -20px;
        filter: blur(10px);
        z-index: -1;
    }

    .index-header-title-all {
        padding: 4em 0;
    }

    .index-header-title {
        font-size: 4em;
        margin: 0 0 10px 0;
    }

    .index-header-subtitle {
        font-size: 0.5em;
        display: block;
        margin: 0;
    }

    .index-header-bar-inner-wrapper {
        max-width: 350px;
        display: flex;
    }

    .index-header-bar-inner-wrapper > * , .index-header-bar-search {
        margin: 10px 0.3em;
    }

    .index-header-bar-search {
        padding: 0 0 2px 0;
    }

    .bar_fount {
        line-height: 1.6;
        padding: 0;
        margin: 14px 7px 0;
        text-decoration: none;
        color: #fff;
    }

</style>
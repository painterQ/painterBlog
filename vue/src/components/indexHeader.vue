<template>
    <div class="index-layout-header">
        <div class="index-header-bar" v-if="this.showBar">
            <img class="index-header-logo" src="../../public/avatar.jpeg"/>
            <router-link :to="'/doc' + this.$store.state.currentID">关于我</router-link>
            <router-link to="/list">标签</router-link>
            <router-link to="/list">目录</router-link>
        </div>
        <div class="index-header-title-all">
            <div class="index-header-title-center">
                <painter-tag v-for="t in this.$store.state.headerTags" :key="t">{{t}}</painter-tag>
                <h1 class="index-header-title">{{this.$store.state.headerTitle}}</h1>
                <h2 class="index-header-subtitle">{{this.$store.state.headerSubTitle}}</h2>
                <span class="meta">{{this.$store.state.headerName}} on {{new Date(this.$store.state.headerTime) | moment}}</span>
            </div>
        </div>
    </div>
</template>

<script>
    import PainterTag from "./tag";
    import api from "../api/rpc"

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
            setHeader() {
                //title, subTitle, time, tags, name
                this.$store.commit("setHeader", {
                    title: this.$store.state.blogTitle,
                    subTitle: this.$store.state.blogSubTitle,
                    time: this.$store.state.authorLastLogin,
                    tags: ["博客"],
                    name: this.$store.state.authorName,
                })
            },
            up2Top() {
                //回到顶部
                document.documentElement.scrollTop = 0;
                document.body.scrollTop = 0;
            },
            updateHeader(newPath){
                if (newPath.startsWith("/doc")) {
                    let currentDoc = this.$store.state.docs.docSet[this.$route.path.substr(4)]
                    this.$store.commit("setHeader",{
                        title: currentDoc.title,
                        subTitle: currentDoc.subTitle,
                        time : currentDoc.time,
                        tags: JSON.parse(JSON.stringify(currentDoc.tags)),
                        name: currentDoc.name,
                    })
                } else if (newPath.startsWith("/list")) {
                    this.$store.commit("setHeader",{
                        title: this.$store.state.blogTitle,
                        subTitle: this.$store.state.blogSubTitle,
                        time : this.$store.state.authorLastLogin,
                        tags: ["博客"],
                        name: this.$store.state.authorName,
                    })
                }
            }
        },
        watch: {
            //如果没有immediate，避免了部分组件内路由，但是mounted还是不能省的，因为不会触发
            '$route.fullPath': {
                handler:function (newFlag, /*oldFlag*/) {
                    this.up2Top();
                    this.$store.state.docs.updateList('/doca', 10).then(
                        ()=>{
                            this.updateHeader(newFlag)
                        }
                    )
                },
                //立即触发，可以省略很多mounted，因为这里刷新也会触发了
                //最佳实践就是，新建一个总会存在的components，然后在其中加route的watcher
                immediate: true
            }
        },
        mounted() {
            window.addEventListener('scroll', this.menu, true);
            //获取作者信息
            api.getAuthorInfo().then(
                (res) => {
                    let data = res.data;
                    //{avatar: "./avatar.jpeg", lastLogin: 1580263669,
                    // name: "Painter Qiao", say: "a blog for dear & love"}
                    //avatar, loginTime, name, say
                    this.$store.commit("setAuthor", data);
                    this.setHeader();
                }
            )
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
        mix-blend-mode: difference;
        margin: 2em auto;
        width: 50%;
    }

    .meta {
        font-family: 'Lora', 'Times New Roman', serif;
        font-style: italic;
        font-weight: 300;
        font-size: 18px;
    }

    .index-header-bar {
        width: 100vw;
        position: fixed;
        height: calc(1em * 1.6 + 16px);
        top: 0;
        background: rgba(255, 255, 255, .3);
        overflow: hidden;
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

    .index-header-bar > * {
        line-height: 1.6;
        padding: 0 10px 2px;
        margin: 7px;
        float: right;
        text-decoration: none;
        color: #fff;
    }

    /*a .router-link-active{*/
    /*    color: #3399ff;*/
    /*}*/
    .index-header-logo {
        width: 42px;
        height: 42px;
        margin-right: 22px;
        padding: 0 2px 10px;
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
</style>
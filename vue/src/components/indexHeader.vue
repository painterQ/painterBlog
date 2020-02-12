<template>
    <div class="header-container">
        <transition name="fade">
            <div class="header-bar" :style="bc" v-if="showBar">
                <div class="header-bar-inner-wrapper">
                    <el-avatar
                            shape="circle"
                            :size="40"
                            fit="contain"
                            :src="avatar"
                            @error="avatarError"></el-avatar>
                    <a @click="aboutMe(true)" class="bar_fount">关于我</a>
                    <router-link to="/tag" class="bar_fount">标签</router-link>
                    <router-link to="/list" class="bar_fount">目录</router-link>
                </div>
                <el-form class="index-header-bar-search">
                    <el-input placeholder="搜索..." prefix-icon="el-icon-search"></el-input>
                </el-form>
            </div>
        </transition>

        <div class="header-title-center header-title-center-before">
            <div>
            <painter-tag v-for="t in $store.state.headerTags" :key="t">
                {{t}}
            </painter-tag></div>
            <h1 class="index-header-title mix header-title-align-before">{{$store.state.headerTitle}}</h1>
            <h2 class="index-header-subtitle mix header-title-align-before">{{$store.state.headerSubTitle}}</h2>
            <div class="meta mix header-title-align-before">{{$store.state.headerName}}&nbsp;on&nbsp;
                {{new Date($store.state.headerTime) | moment}}
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
                showBar: true,
                bc: {
                    backgroundColor: 'rgba(0,0,0,0)',
                    color: '#000'
                },
                an: false,
            }
        },
        methods: {
            menu() {
                let scroll = document.documentElement.scrollTop || document.body.scrollTop;
                this.showBar = scroll <= this.lastScroll || scroll === 0;
                if (!this.showBar) {
                    this.aboutMe(false)
                }
                this.lastScroll = scroll;
                if (scroll === 0) {
                    this.bc = {
                        backgroundColor: 'rgba(0,0,0,0)',
                        color: '#000'
                    }
                } else {
                    this.bc = {
                        backgroundColor: 'rgba(0,0,0,0.5)',
                        color: '#fff'
                    }
                }
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
            aboutMe(show) {
                if(show === this.an){
                    return
                }
                this.an = show
                if (show) {
                    let htcb = document.querySelectorAll('.header-title-center-before')
                    let htab = document.querySelectorAll('.header-title-align-before')
                    this.$store.commit("setHeader", {
                        title: this.$store.state.authorName,
                        subTitle: this.$store.state.authorSay,
                        time: this.$store.state.authorLastLogin,
                        tags: [this.$store.state.github, this.$store.state.mail],
                        name: "",
                    });
                    for(let {classList} of htcb){
                        classList.remove('header-title-center-before')
                        classList.add('header-title-center-after')
                    }
                    for(let {classList}  of htab){
                        classList.remove('header-title-align-before','mix')
                        classList.add('header-title-align-after')
                    }
                } else {
                    let htcb = document.querySelectorAll('.header-title-center-after')
                    let htab = document.querySelectorAll('.header-title-align-after')
                    this.$store.dispatch("setCurrentPath",
                        this.$store.state.currentPath)
                    for(let {classList}  of htcb){
                        classList.remove('header-title-center-after')
                        classList.add('header-title-center-before')
                    }
                    for(let {classList} of htab){
                        classList.remove('header-title-align-after')
                        classList.add('header-title-align-before','mix')
                    }
                }
            }
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
        beforeDestroy() {
            window.removeEventListener('scroll', this.menu, true);
        }
    }
</script>

<style scoped>
    .header-container {
        color: #fff;
        font-family: '黑体', sans-serif;
        font-size: 10px;
    }

    .header-title-center {
        background: url("../../public/background.jpg") top / cover fixed;
        border-color: #00000000;
        background-clip: border-box;
        border-style: solid;
        width: 100%;
        box-sizing: border-box;
        white-space: nowrap;
    }

    .header-title-center-before {
        border-width: 10vh 25vw;
        font-size: 10px;
        transition: border-width 0.7s;
        transition: font-size 0.7s;
    }

    .header-title-center-after {
        border-width: 45vh 25vw;
        font-size: 20px;
        transition: border-width 0.7s;
    }

    .header-title-align-before {
        text-align: left;
        transition: text-align 0.7s;
    }

    .header-title-align-after {
        text-align: center;
        transition: text-align 0.7s;
    }

    .header-title-center h1 {
        font-size: 6em;
    }

    .header-title-center h2 {
        font-size: 2em;
    }

    .mix {
        mix-blend-mode: difference;
    }

    .meta {
        font-family: 'Caflisch Script', 'Adobe Poetica', cursive;
        font-style: italic;
        font-weight: 300;
        font-size: 1.8em;
    }

    .header-bar {
        width: 100%;
        position: fixed;
        top: 0;
        overflow: hidden;
        justify-content: space-around;
        display: flex;
        flex-wrap: wrap;
    }

    .header-bar-inner-wrapper {
        max-width: 350px;
        display: flex;
    }

    .header-bar-inner-wrapper > *, .index-header-bar-search {
        margin: 10px 0.3em;
    }

    .index-header-bar-search {
        padding: 0 0 2px 0;
    }

    .bar_fount {
        line-height: 1.6;
        padding: 0;
        margin: 14px 7px 0;
        font-size: 2em;
        text-decoration: none;
        color: inherit;
        cursor: pointer;
    }

    .index-header-title {
        font-size: 4em;
        margin: 0 0 10px 0;
    }

    .index-header-subtitle {
        display: block;
        margin: 0;
    }


    .fade-enter-active, .fade-leave-active {
        transition: opacity 1s;
    }

    .fade-enter, .fade-leave-to {
        opacity: 0;
    }

</style>
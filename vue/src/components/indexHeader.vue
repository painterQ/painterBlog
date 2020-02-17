<template>
    <div class="header-container">
        <transition name="fade">
            <div class="header-bar" :style="bc" v-show="showBar">
                <div class="header-bar-inner-wrapper">
                    <el-avatar
                            ref="avatar"
                            class="avatar-before"
                            shape="circle"
                            :size="40"
                            @error="errorHandler"
                            :src="avatar"
                            fit="cover"
                            @click.native="aboutMe"></el-avatar>
                    <router-link to="/tags" class="index_header_bar_fount">标签</router-link>
                    <router-link to="/list" class="index_header_bar_fount" ref="header_links">目录</router-link>
                </div>
                <el-form class="index-header-bar-search">
                    <el-input placeholder="搜索..." prefix-icon="el-icon-search"></el-input>
                </el-form>
            </div>
        </transition>

        <div class="header-title-center header-title-center-before" ref="header-title-center">
            <div class="header-title-align-before" ref="header-title-align-1">
            <painter-tag v-for="t in $store.state.headerTags" :key="t" :inner="t"/></div>
            <h1 class="index-header-title mix header-title-align-before" ref="header-title-align-2">{{$store.state.headerTitle}}</h1>
            <h2 class="index-header-subtitle mix header-title-align-before" ref="header-title-align-3">{{$store.state.headerSubTitle}}</h2>
            <div class="meta mix header-title-align-before" ref="header-title-align-4">{{$store.state.headerName}}&nbsp;on&nbsp;
                {{new Date($store.state.headerTime) | moment}}
            </div>
            <div class="friendLinkPad" ref="friendLinkPad" v-show="!an">
                <div>常用链接</div>
            </div>
            <div class="el-icon-arrow-down header-down" v-show="!an" @click="aboutMe"></div>
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
                an: true,
                links: null
            }
        },
        methods: {
            errorHandler(){
                return true
            },
            menu() {
                let scroll = document.documentElement.scrollTop || document.body.scrollTop;
                this.showBar = scroll <= this.lastScroll || scroll === 0;
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
            aboutMe() {
                document.body.style.overflowY = this.an?"hidden":"scroll"
                document.documentElement.scrollTop = 0
                document.body.scrollTop = 0
                let avatar = this.$refs['avatar'].$el.classList
                let htcb = this.$refs["header-title-center"].classList
                let hta1 = this.$refs['header-title-align-1'].classList
                let hta2 = this.$refs['header-title-align-2'].classList
                let hta3 = this.$refs['header-title-align-3'].classList
                let hta4 = this.$refs['header-title-align-4'].classList
                if (this.an) {
                    this.$store.commit("setHeader", {
                        title: this.$store.state.authorName,
                        subTitle: this.$store.state.authorSay,
                        time: this.$store.state.authorLastLogin,
                        tags: ["https://github.com/"+this.$store.state.github, this.$store.state.mail],
                        name: "",
                    });
                    htcb.remove('header-title-center-before')
                    htcb.add('header-title-center-after')
                    avatar.remove('avatar-before')
                    avatar.add('avatar-after')
                    hta1.remove('header-title-align-before','mix')
                    hta1.add('header-title-align-after')
                    hta2.remove('header-title-align-before','mix')
                    hta2.add('header-title-align-after')
                    hta3.remove('header-title-align-before','mix')
                    hta3.add('header-title-align-after')
                    hta4.remove('header-title-align-before','mix')
                    hta4.add('header-title-align-after')
                } else {
                    this.$store.dispatch("setCurrentPath", this.$route.path)
                    htcb.remove('header-title-center-after')
                    htcb.add('header-title-center-before')
                    avatar.remove('avatar-after')
                    avatar.add('avatar-before')
                    hta1.remove('header-title-align-after')
                    hta1.add('header-title-align-before','mix')
                    hta2.remove('header-title-align-after')
                    hta2.add('header-title-align-before','mix')
                    hta3.remove('header-title-align-after')
                    hta3.add('header-title-align-before','mix')
                    hta4.remove('header-title-align-after')
                    hta4.add('header-title-align-before','mix')
                }
                this.an = !this.an;
            }
        },
        watch: {
            //如果没有immediate，避免了部分组件内路由，但是mounted还是不能省的，否则刷新不会触发
            '$route.path': {
                async handler(newFlag, /*oldFlag*/) {
                    this.up2Top();
                    await this.$store.state.initPromise
                    this.$store.dispatch('setCurrentPath', newFlag);
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
            let links = document.querySelectorAll('.friend-links-a.friend-links-header')
            let el = this.$refs['header_links'].$el
            links.forEach((e)=>{
                e.classList.add("index_header_bar_fount")
                el.parentElement.appendChild(e)
            });

            links = document.querySelectorAll('.friend-links-a:not(.friend-links-header)')
            el = this.$refs['friendLinkPad']
            links.forEach((e)=>{
                el.appendChild(e)
            })
        },
        beforeDestroy() {
            window.removeEventListener('scroll', this.menu, true);
        }
    }
</script>

<style scoped>
    @keyframes vertical-move {
        0% {top: -5px; color: #00000000}
        25% {top:0px; color: #000000}
        100% {top: 15px; color: #00000000}
    }

    .header-down{
        animation: vertical-move 0.7s infinite;
        font-weight: bolder;
        font-size: 2em;
        position: relative;
        width: 100%;
        text-align: center;
        cursor: pointer;
    }
    .avatar-before{
        transform: rotateZ(360deg);
        transition: transform 0.7s;
        cursor: pointer;
    }
    .avatar-after{
        cursor: pointer;
        transform: rotateZ(180deg);
        transition: transform 0.7s;
    }
    .header-container {
        color: #074B72;
        font-family: -apple-system, "Microsoft YaHei", 'Impact', 'Charcoal', sans-serif;
        font-size: 10px;
    }

    .header-title-center {
        background: url("../../public/backgroundImage.jpg") top / cover fixed;
        border-color: #00000000;
        background-clip: border-box;
        border-style: solid;
        width: calc(50vw - 3px);
        box-sizing: content-box;
        white-space: nowrap;
    }

    .header-title-center-before {
        border-width: 10vh 25vw;
        font-size: 10px;
        transition: border-width 0.7s;
        transition: font-size 0.7s;
    }

    .header-title-center-after {
        border-width: 20vh 25vw 100vh;
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
        font-family: 'Merriweather', Georgia, serif;
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
        margin: 10px 7px;
    }

    .index-header-bar-search {
        padding: 0 0 2px 0;
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

    .friendLinkPad {
        background-color: #ffffff7c;
        border-radius: 10px;
        width: 50vw;
        padding: 1em;
    }
    .friendLinkPad > div{
        text-align: center;
        font-size: 1.2em;
        margin: 0 auto;
        width: 10em;
    }
    .friendLinkPad a{
        color: black;
        text-decoration: none;
    }

</style>

<style>
    .index_header_bar_fount {
        line-height: 1.6;
        padding: 0;
        margin: 10px 7px 0;
        font-size: 2em;
        text-decoration: none;
        color: #000;
        cursor: pointer;
    }

    .index_header_bar_fount:hover{
        color: #074B72;
    }
</style>
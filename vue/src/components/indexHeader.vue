<template>
    <div class="index-layout-header">
        <div class="index-header-bar" v-if="this.showBar">
            <img class="index-header-logo" src="../../public/avatar.jpeg"/>
                <router-link :to="'/doc' + this.$store.state.currentID">关于我</router-link>
                <router-link to="/list">标签</router-link>
                <router-link to="/list">目录</router-link>
        </div>
        <div  class="index-header-title-all">
            <div class="index-header-title-center">
                <div class="tags"><i>标签</i></div>
                <h1 class="index-header-title">{{this.$store.state.author.name}}</h1>
                <h2 class="index-header-subtitle">个人博客</h2>
                <span class="meta">post by 乔</span>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "index-header",
        data: function(){
          return{
              lastScroll : 0,
              showBar: true
          }
        },
        methods: {
            menu() {
                let scroll = document.documentElement.scrollTop || document.body.scrollTop;
                this.showBar = scroll <= this.lastScroll || scroll === 0;
                this.lastScroll = scroll
            }
        },
        mounted() {
            window.addEventListener('scroll', this.menu, true)
        },
    }
</script>

<style scoped>
    .index-layout-header,.index-header-bar::after{
        background: url("../../public/background.jpg") top / cover fixed;
        color: #fff;
        width: 100vw;
    }

    .index-header-title-center{
        margin: 2em auto;
        width: 50%;
    }
    .tags > i {
        display: inline-block;
        border: 1px solid rgba(255, 255, 255, .8);
        border-radius: 999em;
        padding: 0 10px;
        color: orange;
        line-height: 24px;
        font-size: 12px;
        text-decoration: none;
        margin: 0 1px 6px 1px;
    }
    .meta{
        font-family: 'Lora', 'Times New Roman', serif;
        font-style: italic;
        font-weight: 300;
        font-size: 18px;
    }

    .index-header-bar{
        width: 100vw;
        position: fixed;
        height: calc(1em * 1.6 + 16px);
        top: 0;
        background: rgba(255, 255, 255, .3);
        overflow: hidden;
    }

    .index-header-bar::after{
        display: block;
        content: '';
        position: absolute;
        height: calc(100% + 40px);
        width: calc(100% + 40px);
        top: -20px; left: -20px;
        filter: blur(10px);
        z-index: -1;
    }

    .index-header-bar > *{
        line-height: 1.6;
        padding: 0 10px 2px;
        margin: 7px;
        float: right;
        text-decoration:none;
        color: #fff;
    }

    /*a .router-link-active{*/
    /*    color: #3399ff;*/
    /*}*/
    .index-header-logo{
        width: 42px;
        height: 42px;
        margin-right: 22px;
        padding: 0 2px 10px;
    }
    .index-header-title-all{
        padding: 4em 0;
    }
    .index-header-title{
        font-size: 4em;
        margin: 0 0 10px 0;
    }
    .index-header-subtitle{
        font-size: 0.5em;
        display: block;
        margin: 0;
    }
</style>
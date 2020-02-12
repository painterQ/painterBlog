<template>
    <div class="doc-card">
        <div class="title">{{doc.title}}</div>
        <div class="subTitle">{{doc.subTitle}}</div>
        <div class="path">path:<strong>{{doc.id}}</strong></div>
        <painter-tag v-for="(t,tindex) in doc.tags" :key="tindex">{{t}}</painter-tag>
        <div class="abstract" v-popover:popover>{{this.trim(doc.abstract)}}</div>
        <el-popover
                placement="top-start"
                ref="popover"
                title="概要"
                width="200"
                trigger="hover"
                v-popover:abstract
                :content="doc.abstract">
        </el-popover>
        <div class="info">
            <i class="top el-icon-star-on" v-if="doc.attr === 1">置顶</i>
            <span>{{new Date(doc.lastTime) | moment}}</span>
        </div>
    </div>
</template>

<script>
    import painterTag from "./tag"
    import {Popover} from "element-ui"
    import vue from "vue"

    vue.use(Popover);
    export default {
        name: "doc-card",
        components: {
            painterTag
        },
        props: {
            doc: {
                type: Object,
                default() {
                    return null;
                }
            }
        },
        methods: {
            trim(s) {
                if(s.length < 14){
                    return s
                }
                return s.substr(0,10) + "..."
            }
        }
    }
</script>

<style scoped>
    .title {
        font-size: 1.2em;
        font-family: sans-serif;
        font-weight: bold;
        line-height: 2em;
        white-space: nowrap;
    }

    .top {
        color: orangered;
        margin: 0 10px 0 0;
    }

    .abstract {
        font-weight: 200;
        font-size: 20px;
        font-family: "Helvetica Neue", "Arial", "PingFang SC", "Hiragino Sans GB", "STHeiti", "Microsoft YaHei", "Microsoft JhengHei", "Source Han Sans SC", "Noto Sans CJK SC", "Source Han Sans CN", "Noto Sans SC", "Source Han Sans TC", "Noto Sans CJK TC", "WenQuanYi Micro Hei", SimSun, sans-serif;
        line-height: 1.7;
    }

    .path {
        word-wrap: break-word;
        word-break: break-word;
        color: #757d87;
        font-weight: lighter;
        font-size: 12px;
        font-style: italic;
    }

    .subTitle {
        color: #0085a1;
    }

    .info {
        font-size: 12px;
        color: #757d87;
        font-style: italic;
        font-family: Sans-serif, serif;
    }

    .doc-card {
        border-radius: 4px;
        overflow-y: hidden;
        background-color: #fff;
        box-shadow: 0 0 4px rgba(0, 0, 0, 0.1);
        box-sizing: border-box;
        padding: 0.5em;
        margin: 0.2em;
        max-width: 900px;
        min-width: 300px;
        overflow-x: auto;
        flex-grow: 1;
    }

    .doc-card:hover {
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.3);
        position: relative;
        top: 1px;
        cursor: pointer;
    }

</style>
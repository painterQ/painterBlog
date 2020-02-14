<template>
    <div class="doc-card">
        <div class="title">{{doc.title}}</div>
        <div class="subTitle">{{doc.subTitle}}</div>
        <div class="path">path:<strong>{{doc.id}}</strong></div>
        <painter-tag v-for="(t,tindex) in doc.tags" :key="tindex" :inner="t"/>
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
            <span>{{new Date(doc.time) | moment}}</span>
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
        font-size: 3em;
        font-weight: bold;
        white-space: nowrap;
    }

    .title:hover{
        color: #0085a1;
    }

    .subTitle {
        color: #0085a1;
        line-height: 2.4em;
        font-size: 2em;
    }

    .top {
        color: orangered;
        margin: 0 10px 0 0;
    }

    .abstract {
        font-weight: 200;
        font-size: 1em;
    }

    .path {
        word-wrap: break-word;
        word-break: break-word;
        color: #757d87;
        font-weight: lighter;
        font-size: 1em;
        font-family: "Lucida Console", Monaco, monospace;

    }

    .info {
        margin: 1em 0;
        font-size: 1em;
        color: #757d87;
        font-family: -apple-system, "Microsoft YaHei", 'Impact', 'Charcoal', sans-serif;
    }
    .info span{
        font-family: 'Merriweather', Georgia, serif;
    }

    .doc-card {
        border-radius: 4px;
        background-color: #fff;
        box-shadow: 0 0 4px rgba(0, 0, 0, 0.1);
        padding: 1em;
        margin: 0.2em;
        max-width: 600px;
        font-size: 8px;
        line-height: 1.5;
        overflow: hidden;
        font-family: -apple-system, "Microsoft YaHei", 'Impact', 'Charcoal', sans-serif;
        flex-grow: 1;
    }

    .doc-card:hover {
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.3);
        position: relative;
        top: 1px;
        cursor: pointer;
    }

</style>
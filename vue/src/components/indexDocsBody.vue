<template>
    <div>
        <div class="index-body-all">
            <index-aside class="index-body-aside">
                <!--                <strong class="tag-tital"></strong>-->
                <!-- 使用外部插件自动生成目录npm i katelog -S-->
                <!-- https://github.com/KELEN/katelog-->
                <h3 class="doc-cateLog-title">目录</h3>
                <div id="doc-cateLog" ref="doc-cateLog"></div>
            </index-aside>
            <div class="index-body-main">
                <!-- learn: 插入HTML-->
                <main id="doc-content" v-html="document"></main>
            </div>
        </div>
        <div class="doc-bottom">
            <h2>推荐文章</h2>
            <ul>
                <li>文章一</li>
                <li>文章二</li>
            </ul>
            <hr>
            <el-button @click="prevDoc">上一篇</el-button>
            <el-button style="float:right;" @click="nextDoc">下一篇</el-button>
            <div class="coffee">赏</div>
        </div>
    </div>
</template>

<script>

    import kateLogClass from 'katelog';
    import indexAside from "@/components/indexAside.vue";
    import Prism from 'prismjs';

    export default {
        name: "index-docs-body",
        components: {
            indexAside
        },
        data: function () {
            return {
                scroll: 0,
                kateLog: null,
                menuFloat: false,
                document: ""
            }
        },
        watch:{
            "$route.path":{
                async handler(){
                    console.log("indexDocBody activated start")
                    let mateCache = await this.$store.state.initPromise;
                    this.$store.dispatch("setCurrentPath", this.$route.path)
                    let doc = mateCache.getDocFromStore(this.$route.path)
                    if(doc instanceof Promise){
                        this.document = await doc
                        console.log("fz 1:",this.document)
                    }else if(doc === "/404"){
                        this.$router.replace("/404");
                        return
                    }else {
                        this.document = doc
                        console.log("fz 2:",this.document)
                    }

                    this.$nextTick(() => {
                        this.kateLog.rebuild();
                    });
                    console.log("indexDocBody activated end")
                },
                immediate:true
            }
        },
        methods: {
            async prevDoc() {
                //仍然在当前组件，所以只是复用，没有重新触发mounted
                let mateCache = await this.$store.state.initPromise;
                let prev = mateCache.prevDoc(this.$route.path);
                let to = "/docs" + prev;
                if (!prev || to === this.$route.path) return;
                this.$router.push("/docs" + prev);
            },
            async nextDoc() {
                let mateCache = await this.$store.state.initPromise;
                let next = mateCache.nextDoc(this.$route.path);
                let to = "/docs" + next;
                if (!next || to === this.$route.path) return;
                this.$router.push("/docs" + next);
            },
        },
        mounted() {
            Prism.highlightAll();
            this.kateLog = new kateLogClass({
                contentEl: 'doc-content',
                catelogEl: 'doc-cateLog',
                linkClass: 'k-catelog-link',
                linkActiveClass: 'k-catelog-link-active',
                // supplyTop: 20,
                selector: ['h2', 'h3'],
                active: null
            });
        },
    }
</script>

<style scoped>
    .coffee {
        font-size: 28px;
        line-height: 58px;
        position: relative;
        display: block;
        width: 60px;
        height: 60px;
        margin: 0 auto;
        padding: 0;
        text-align: center;
        vertical-align: middle;
        color: #fff;
        border: 1px solid #f1b60e;
        border-radius: 50%;
        background: #fccd60;
        background: linear-gradient(to bottom, #fccd60 0, #fbae12 100%, #2989d8 100%, #207cca 100%);
    }

    .doc-cateLog-title{
        font-family: -apple-system, "Microsoft YaHei", 'Impact', 'Charcoal', sans-serif;
        background-image: linear-gradient(top left,#f0381a, white);
        font-weight: bolder;
        font-size: 24px;
        margin: 12px 0;
    }

    #doc-cateLog{
        background: #a0a0a0;
        border-radius: 6px;
    }

</style>

<style>
    main > h2{
        font-family: -apple-system, "Microsoft YaHei", 'Impact', 'Charcoal', sans-serif;
        background-image: linear-gradient(top left,#f0381a, white);
        font-weight: bolder;
        font-size: 30px;
        margin: 12px 0;
    }

    .index-body-all {
        display: flex;
        flex-flow: column;
        flex-direction: row-reverse;
        flex-wrap: wrap;
    }

    .index-body-main {
        align-items: flex-start;
        margin: 2em;
        width: 20em;
        flex-grow: 2;
    }

    .index-body-aside {
        /*border-left: rgba(88, 88, 88, 0.1) 1px solid;*/
        margin: 2em 0 0 1em;
        flex-grow: 1;
        max-width: 20%;
        /*top: 10px;*/
        /*position: sticky;*/
        height: 50vh;
    }

    .k-catelog-level-2 {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        margin-top: 5px;
        font-size: 18px;
        cursor: pointer;
        word-break: break-all;
    }

    .k-catelog-level-3 {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        margin-top: 3px;
        font-size: 12px;
        cursor: pointer;
        word-break: break-all;
    }

    #doc-cateLog li {
        list-style-type: decimal;
    }

    #doc-cateLog ul{
        padding-left: 2em;
    }

    .k-catelog-link-active {
        color: #f0381a;
        font-weight: bolder;
    }
</style>
<template>
    <div>
        <div class="index-body-all">
            <index-aside class="index-body-aside">
                <!--                <strong class="tag-tital"></strong>-->
                <!-- 使用外部插件自动生成目录npm i katelog -S-->
                <!-- https://github.com/KELEN/katelog-->
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

    export default {
        name: "index-docs-body",
        components: {
            indexAside
        },
        data: function () {
            return {
                scroll: 0,
                kateLog: null,
                menuFloat: false
            }
        },
        computed: {
            document(){
                let doc = this.$store.getters.getDoc;
                if (doc === "/404"){
                    console.log("#####404")
                    this.$router.replace("/404");
                    return ""
                }
                this.$nextTick(() => {
                    this.kateLog.rebuild();
                });
                return doc
            },
        },
        methods: {
            prevDoc() {
                //仍然在当前组件，所以只是复用，没有重新触发mounted
                let current = this.$route.path.substr(5);
                let prev = this.$store.getters.prevDoc;
                if (!prev || prev === current) return;
                this.$router.push("/docs" + prev);
            },
            nextDoc() {
                let current = this.$route.path.substr(5);
                let next = this.$store.getters.nextDoc;
                if (!next || next === current) return;
                this.$router.push("/docs" + next);
            },
        },
        mounted() {
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


    /*#doc-cateLog{*/
    /*    position: sticky;*/
    /*    top: 200px;*/
    /*}*/

    /*p h1 h2 h3 h4 h5 h6*/
    main p {

    }

    main h1 {

    }

    main h2 {

    }

    main h3 {

    }

    main h4, main h5, main h6 {

    }
</style>

<style>
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

    .k-catelog-link {
        font-size: 0.7em;
        word-break: keep-all;
    }

    #doc-cateLog li {
        display: block;
    }

    .k-catelog-link-active {
        color: red;
    }
</style>
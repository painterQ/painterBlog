<template>
    <div>
        <div class="index-body-all">
            <index-aside class="index-body-aside">
                <!--                <strong class="tag-tital"></strong>-->
                <!-- 使用外部插件自动生成目录npm i katelog -S-->
                <!-- https://github.com/KELEN/katelog-->
                <div id="doc-cateLog" ref="doc-cateLog" ></div>
            </index-aside>
            <div  class="index-body-main">
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
            <div class="share">
                <a class="el-icon-eleme"></a>
            </div>
        </div>
    </div>
</template>

<script>

    import kateLogClass from 'katelog';
    import indexAside from "@/components/indexAside.vue";

    export default {
        name: "index-docs-body",
        components:{
            indexAside
        },
        data: function () {
            return {
                scroll : 0,
                kateLog: null,
                document: "",
                menuFloat: false
            }
        },
        watch: {
            document() {
                this.$nextTick(() => {
                    this.kateLog.rebuild();
                })
            },
        },
        methods: {
            prevDoc(){
                //仍然在当前组件，所以只是复用，没有重新触发mounted
                let current = this.$route.path.substr(4);
                let pref = this.$store.state.docs.docSet[current].pref()
                this.$router.push("/doc" + pref);
            },
            nextDoc(){
                let current = this.$route.path.substr(4)
                let next = this.$store.state.docs.docSet[current].next()
                this.$router.push("/doc" + next);
            },
            render(path) {
                if(!/^\/doc\/.*/.test(path)){
                    return
                }
                path = path.substr(4)
                console.log('change docs', "id:" + path);
                this.$store.state.docs.get(path).then(
                    (data)=>{
                        if(data === null){
                            this.$router.replace("/404");
                            return
                        }
                        this.document = data;
                    }
                )
            }
        },
        mounted() {
            console.log("doc page mounted ")
            this.kateLog = new kateLogClass({
                contentEl: 'doc-content',
                catelogEl: 'doc-cateLog',
                linkClass: 'k-catelog-link',
                linkActiveClass: 'k-catelog-link-active',
                // supplyTop: 20,
                selector: ['h2', 'h3'],
                active: null
            });
            this.render(this.$route.path)
        },
        //组件内的路由守卫
        beforeRouteUpdate(to,from,next){
            this.render(to.path)
            next();
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

    .tag-tital {

    }

    .tag-tital:before {
        content: "\f02e";
        margin-right: 0.5em;
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
    .index-body-all{
        display: flex;
        flex-flow: column;
        flex-direction: row-reverse;
        flex-wrap: wrap;
    }

    .index-body-main {
        align-items: flex-start;
        margin: 2em;
        width: 20em;
        flex-grow:2;
    }

    .index-body-aside {
        border-left: rgba(88, 88, 88, 0.1) 1px solid;
        margin: 0 0 0 1em;
        flex-grow:1;
        max-width: 10em;
        /*top: 10px;*/
        /*position: sticky;*/
        height: 400px;
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
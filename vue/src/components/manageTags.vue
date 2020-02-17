<template>
    <div id="manage-tags">
        <div id="tags-container">
            <span>tags:&nbsp;</span>
            <painter-tag ref="painter-tags" v-for="t in getTagsSlice"
                         :key="t" @clickTag="choseTag(t)" :inner="t" :selected="getTagFromPath($route.path) === t"/>
        </div>
        <div id="docList" class="scroll">
            <doc-card v-for="t in show" :key="t" :doc="t" @selectCard="clickDoc"></doc-card>
        </div>
        <el-dialog
                title="提示"
                :visible.sync="dialogVisible"
                width="30%">
            是否要开始编辑《{{getDocTitle}}》？这可能导致编辑器已有的内容被覆盖
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="editAgain">编 辑</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
    import PainterTag from "./tag";
    import DocCard from "./docCard";
    import DocListClass from "../page_index/docList";
    import vue from 'vue'
    import api from '../api/rpc'
    import {Dialog} from 'element-ui'

    vue.use(Dialog);
    vue.use(api);
    export default {
        name: "tags-manager",
        components: {DocCard, PainterTag},
        data() {
            return {
                tagsMap: null,
                show: [],
                tag: "",
                docList: null,
                forecourt: this.$store.state.docs instanceof DocListClass,

                dialogVisible: false,
                transferDoc: null,

                inRender: false
            }
        },
        methods: {
            choseTag(tag) {
                this.tag = tag;
            },
            editAgain(){
                this.$store.commit("changeCurrentDoc", this.transferDoc)
                this.$router.push('/document')
            },
            clickDoc(doc){
                if(this.forecourt){
                    this.$router.push('/docs'+doc.id)
                    return
                }
                this.dialogVisible = true
                this.transferDoc = doc
            },
            getTagFromPath(path){
                let i = path.indexOf("/tags/")
                return i>=0? path.substr(6):""
            },
            async initDocListAndTags(){
                if (this.forecourt){
                    this.tagsMap = (await this.$_getTags()).data;
                    let mateCache = await this.$store.state.initPromise;
                    this.docList = mateCache.docMateList();
                }else {
                    let promiseTags = this.$_getTags();
                    let tmp = (await this.$_getDocsList({start: "/doca", length: 10})).data;
                    this.docList = tmp.list;
                    this.tagsMap = (await promiseTags).data;
                }
            },
            render(){
                console.log("render start")
                if (this.inRender) return;
                this.inRender = true
                if (this.tag === ""){
                    this.show = this.docList;
                    this.inRender = false
                    console.log("render end at 1")
                    return
                }

                let thisTagIncludeDoc = this.tagsMap[this.tag];
                if(this.tag !== "" && this.tagsMap!==null && !thisTagIncludeDoc){
                    this.$router.push('/404')
                    this.inRender = false
                    console.log("render end at 2")
                    return
                }

                if(!this.docList){
                    this.inRender = false
                    console.log("render end at 3")
                    return
                }
                let ret = [];
                for (let doc of this.docList) {
                    for (let docIndex in thisTagIncludeDoc) {
                        if (doc.id === thisTagIncludeDoc[docIndex]){
                            doc.time = Number.parseInt(doc.time) * 1000
                            ret.push(doc)
                        }
                    }
                }
                this.show = ret
                this.inRender = false
                console.log("render end at 4")
            }
        },
        computed: {
            getTagsSlice() {
                let ret = [];
                for (let k in this.tagsMap) {
                    ret.push(k)
                }
                return ret
            },
            getDocTitle() {
                return this.transferDoc?this.transferDoc.title:""
            },
            showDocList: {
                get() {
                    return this._docNeedRefresh
                },
                set(v) {
                    this._docNeedRefresh = v
                }
            }
        },
        async mounted(){
            await this.initDocListAndTags()
        },
        watch:{
            "$route.path":{
                handler(path){
                    this.tag =  this.getTagFromPath(path)
                },
                immediate: true
            },
            tag() {
                console.log("tag change cause")
                this.render()
            },
            docList(){
                console.log("docList change cause")
                this.render()
            },
            tagsMap:{
                handler(){
                    console.log("tagsMap change cause")
                    this.render()
                },
                deep: true
            }
        },
    }
</script>

<style scoped>
    #manage-tags {
        background-color: white;
        width: 80%;
        padding: 2em;
        box-sizing: border-box;
        height: 100%;
        margin: 0 auto;
    }

    #docList {
        display: flex;
        flex-wrap: wrap;
        margin-top: 40px;
        box-sizing: border-box;
    }

    #docList:after {
        content: '';
        flex-grow: 99999;
    }

    #tags-container{
        position: relative;
        top: 0;
        margin: 0 1em;
        height: 20%;
        max-width: 100%;
        box-sizing: border-box;
    }
    #tags-container span{
        font-style: italic;
        color: #757d87;
        font-weight: bold;
        font-size: 20px;
    }
</style>
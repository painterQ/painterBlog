<template>
    <div id="manage-tags">
        <div id="tags-container">
            <span>tags:&nbsp;</span>
            <painter-tag ref="painter-tags" v-for="t in getTagsSlice"
                         :key="t" @clickTag="choseTag(t)" :inner="t" :selected="getCurrentTag === t"/>
        </div>
        <div id="docList" class="scroll">
            <doc-card v-for="t in show" :key="t" :doc="t" @selectCard="clickDoc"></doc-card>
        </div>
        <el-dialog
                title="提示"
                :visible.sync="dialogVisible"
                width="30%">
            是否要开始编辑《{{getDoc}}》？这可能导致编辑器已有的内容被覆盖
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
    vue.use(api);
    export default {
        name: "tags-manager",
        components: {DocCard, PainterTag},
        data() {
            return {
                tagsMap: null,
                show: [],
                docList: [],
                forecourt: this.$store.state.docs instanceof DocListClass,

                dialogVisible: false,
                transferDoc: null,
            }
        },
        methods: {
            choseTag(tag) {
                this.tag = tag;
                for (let i in this.$refs['painter-tags']) {
                    if (tag === this.getTagsSlice[i]) {
                        this.$refs['painter-tags'][i].select(true)
                    } else {
                        this.$refs['painter-tags'][i].select(false)
                    }
                }
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
                let p = (path.substr(0,1) === "/")?path.substr(1):path
                let i = p.indexOf("/")
                return i>0? p.substr(i+1):""
            },
            isFormSelf(path){
                if(!path) return false;
                return path.startsWith("/tags")
            },
            async initDocListAndTags(){
                if (this.forecourt){
                    this.tagsMap = (await this.$_getTags()).data;
                    let p = new Promise((resolve)=>{
                        let n = setInterval(()=>{
                            if( this.$store.getters.docMateList.length > 0){
                                clearInterval(n)
                                resolve(this.$store.getters.docMateList)
                            }
                        },100)
                    });
                    this.docList = await p;
                }else {
                    let promiseTags = this.$_getTags();
                    let tmp = (await this.$_getDocsList({start: "/doca", length: 10})).data;
                    this.docList = tmp.list;
                    this.tagsMap = (await promiseTags).data;
                }
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
            getDoc() {
                return this.transferDoc?this.transferDoc.title:""
            },
            getCurrentTag(){
                return this.getTagFromPath(this.$route.path)
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
        watch: {
            "$route.path": {
                async handler(newPath, oldPath) {
                    let tag = this.getTagFromPath(newPath)
                    let fromOther = tag === ""; //example from /document， newPath is /tags
                    let fromNothing = oldPath === ""; //example refresh
                    let fromSelf = this.isFormSelf(oldPath);//example from /tags/example
                    if (!fromSelf){
                        console.log("tags init")
                        await this.initDocListAndTags()
                    }

                    if(!fromOther && !this.tagsMap[tag]){
                        this.$router.push('/404')
                        return
                    }

                    let thisTagIncludeDoc = []
                    if (fromNothing || fromOther){
                        //immediate
                        this.show = this.docList;
                        return
                    }

                    thisTagIncludeDoc = this.tagsMap[tag];
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
                },
                immediate: true
            },
        }
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
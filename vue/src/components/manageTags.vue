<template>
    <div id="manage-tags">
        <div id="tags-container">
            <span>tags:&nbsp;</span>
            <painter-tag ref="painter-tags" v-for="t in getTagsSlice"
                         :key="t" @click.native="choseTag(t)" :inner="t"/>
        </div>
        <div id="docList" class="scroll">
            <doc-card v-for="t in show" :key="t" :doc="t" @click.native="clickDoc"></doc-card>
        </div>
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
                tag: "",
                show: [],
            }
        },
        methods: {
            choseTag(tag) {
                this.tag = tag;
                for (let i in this.$refs['painter-tags']) {
                    console.log("?", this.tagsMap[tag], i, tag)
                    if (tag === this.getTagsSlice[i]) {
                        this.$refs['painter-tags'][i].select(true)
                    } else {
                        this.$refs['painter-tags'][i].select(false)
                    }
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
            "tag": {
                async handler(newTag) {
                    let docList = [];
                    if (this.$store.state.docs instanceof DocListClass){
                        this.tagsMap = (await this.$_getTags()).data;
                        let p = new Promise((resolve)=>{
                            let n = setInterval(()=>{
                                if( this.$store.getters.docMateList.length > 0){
                                    clearInterval(n)
                                    resolve(this.$store.getters.docMateList)
                                }
                            },100)
                        });
                        docList = await p;
                        console.log('immediate',thisTagIncludeDoc)
                    }else {
                        let promiseTags = this.$_getTags();
                        let tmp = (await this.$_getDocsList({start: "/doca", length: 10})).data;
                        docList = tmp.list;
                        this.tagsMap = (await promiseTags).data;
                    }

                    let thisTagIncludeDoc = []
                    if (!newTag){
                        //immediate
                        this.show = docList
                        return
                    }
                    thisTagIncludeDoc = this.tagsMap[newTag]
                    let ret = []
                    for (let doc of docList) {
                        for (let docIndex in thisTagIncludeDoc) {
                            if (doc.id === thisTagIncludeDoc[docIndex]){
                                if(doc.lastTime &&! doc.time) {
                                    doc.time = Number.parseInt(doc.lastTime) * 1000
                                }
                                ret.push(doc)
                            }
                        }
                    }
                    this.show = ret
                },
                immediate: true
            }
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

    #docList:after {
        content: '';
        flex-grow: 99999;
    }
</style>
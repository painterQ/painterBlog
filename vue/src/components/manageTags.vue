<template>
    <div id="manage-tags">
        <div id="tags-container">
            <span>Tags:&nbsp;</span>
            <painter-tag ref="painter-tags" v-for="t in getTagsSlice"
                         :key="t" @click.native="choseTag(t)">{{t}}
            </painter-tag>
        </div>
        <div id="docList">
            <doc-card v-for="t in show" :key="t" :doc="t"></doc-card>
        </div>
    </div>
</template>

<script>
    import PainterTag from "./tag";
    import DocCard from "./docCard";

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
                    let promiseDocsList = this.$_getDocsList({start: "/doca", length: 10})
                    let {data} = await this.$_getTags();
                    this.tagsMap = data;
                    let {data: docsList} = await promiseDocsList;
                    let ret = [];
                    let list = docsList.list;

                    let thisTagIncludeDoc = []
                    if (newTag) {
                        thisTagIncludeDoc = this.tagsMap[newTag]
                    } else {
                        for (let k in data) {
                            thisTagIncludeDoc = data[k];
                            break
                        }
                    }
                    for (let doc of list) {
                        for (let docIndex in thisTagIncludeDoc) {
                            if (doc.id === thisTagIncludeDoc[docIndex]) ret.push(doc)
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
        border-radius: 10px;
        width: 80%;
        padding: 2em;
        height: 100%;
        margin: 1em auto;
    }

    #docList {
        overflow-y: scroll;
        display: flex;
        flex-wrap: wrap;
    }

    #tags-container{
        margin: 1em;
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
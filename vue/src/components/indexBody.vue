<template>
    <div>
        <div class="index-body-all">
            <index-aside class="index-body-aside">
                <motto style="margin: 10px 0"></motto>
            </index-aside>
            <div class="index-body-main">
                <!--list-->
                <div v-for="arts of docList"
                     :key="arts.id"
                     class="index-body-docs-item"
                     @click="selectDoc(arts.id)">
                    <h2 class="art">
                        <span v-if="arts.attr" class="arts-top">[置顶]</span>
                        <span>{{arts.title}}</span>
                    </h2>
                    <p>{{arts.abstract}}</p>
                    <div>Post by 乔沛杨 on {{arts.time}}</div>
                    <hr/>
                </div>
                <el-pagination
                        background
                        layout="prev, pager, next"
                        :current-page="this.currentPage"
                        @current-change="this.totalNum"
                        :total=getPaginationNum>
                </el-pagination>
            </div>
        </div>
    </div>
</template>

<script>
    import Vue from 'vue'
    import {Pagination} from "element-ui";
    import indexAside from "@/components/indexAside.vue";
    import Motto from "./motto";

    Vue.use(Pagination);

    export default {
        name: 'index-body',
        components: {
            indexAside,
            Motto
        },
        data() {
            return {
                //pagination
                currentPage: 0,
                //[currentPage * pageSize, (currentPage + 1) * pageSize)
                pageSize: 5,

            }
        },
        computed: {
            totalNum() {
                return this.$store.state.total
            },
            //[currentPage * pageSize, (currentPage + 1) * pageSize)
            docList() {
                //1. 计算范围
                //2. 判断范围内的是都都在中，如果没有，则获取
                //3. 渲染
                let currentList = [];
                let self = this;
                return function () {
                    if (!self.$store.state.docsUpdate) {
                        return currentList
                    }
                    self.$store.commit("setDocListUpdateState", false)
                    let output = [];
                    for (let e of self.$store.state.docs) {
                        output.push(e)
                    }
                    currentList = output;
                    return output
                }()
            }
        },
        methods: {
            selectDoc(artID) {
                console.log("selectdoc", artID)
                this.$router.push('/doc' + artID)
            },
            //pagination
            handleCurrentChange() {
                console.log("handleCurrentChange")
            },
        },
    }
</script>

<style scoped>
    .index-body-docs-item > div {
        font-family: 'Lora', 'Times New Roman', serif;
        color: gray;
        font-style: italic;
        margin: 0 0 1em 0;
    }

    .index-body-docs-item:hover {
        color: #1c6ca1;
        cursor: pointer;
    }

    .index-body-docs-item > p {
        color: #a3a3a3;
        font-size: 0.7em;
        font-style: italic;
    }

    .arts-top {
        color: orangered;
        margin: 0 10px 0 0;
    }

    .art {
        font-size: 1.1em;
        margin: 0 0 10px 0;
    }
</style>
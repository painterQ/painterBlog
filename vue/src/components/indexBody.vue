<template>
    <div>
        <div class="index-body-all">
            <index-aside class="index-body-aside">
                <introduction></introduction>
            </index-aside>
            <div class="index-body-main">
                <!--list-->
                <div v-for="arts of docList"
                     :key="arts.id"
                     class="index-body-docs-item"
                     @click="selectDoc(arts.id)">
                    <h2 class="art">
                        <i class="arts-top el-icon-star-on" v-if="arts.attr === 1">置顶</i>
                        <span>{{arts.title}}</span>
                    </h2>
                    <p>{{arts.abstract}}</p>
                    <painter-tag v-for="t of arts.tags" :key="t" :inner="t"/>
                    <div class="time">{{$store.state.authorName}} on {{new Date (arts.time) | moment}}</div>
                    <hr/>
                </div>
                <el-pagination
                        background
                        layout="prev, pager, next"
                        :hide-on-single-page="true"
                        :page-size="pageSize"
                        :current-page.sync="currentPage"
                        @current-change="up2Top"
                        :total="this.$store.state.total">
                </el-pagination>
            </div>
        </div>
    </div>
</template>

<script>
    import Vue from 'vue'
    import {Pagination} from "element-ui";
    import indexAside from "@/components/indexAside.vue";
    import introduction from "./introduction";
    import PainterTag from "./tag";

    Vue.use(Pagination);

    export default {
        name: 'index-body',
        components: {
            indexAside,
            introduction,
            PainterTag
        },
        data() {
            return {
                //pagination
                currentPage: 1,
                pageSize: 5,

            }
        },
        computed: {
            docList() {
                let all = this.$store.getters.docMateList;
                //分页逻辑 [(currentPage -1) * pageSize, currentPage * pageSize)
                return all.slice((this.currentPage - 1) * this.pageSize,
                    this.currentPage * this.pageSize)
            }
        },
        methods: {
            selectDoc(artID) {
                this.$router.push('/docs' + artID)
            },
            //pagination
            up2Top() {
                //回到顶部
                document.documentElement.scrollTop = 0;
                document.body.scrollTop = 0;
            },
        },
    }
</script>

<style scoped>
    .index-body-docs-item > div {
        font-family: 'Lora', 'Times New Roman', serif;
        color: gray;
        margin: 0 0 1em 0;
    }

    .index-body-docs-item:hover {
        color: #1c6ca1;
        cursor: pointer;
    }

    .index-body-docs-item > p {
        color: #a3a3a3;
        font-size: 0.7em;
    }

    .arts-top {
        color: orangered;
        margin: 0 10px 0 0;
    }

    .art {
        font-size: 1.1em;
        margin: 0 0 10px 0;
    }

    .time{
        font-family: 'Merriweather', Georgia, serif;
    }

    .art span, .art i{
        font-size: 2em;
        font-weight: bold;
    }
</style>
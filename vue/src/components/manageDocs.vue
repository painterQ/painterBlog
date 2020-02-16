<template>
    <div class="doc-container">
        <el-dialog
                title="提示"
                :visible.sync="dialogVisible"
                width="30%">
            是否要编辑《{{getDocTitle}}》，可能导致覆盖已有文章
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="editAgain">编 辑</el-button>
            </span>
        </el-dialog>
        <div id="docList" class="scroll">
        <doc-card @click.native="handleDoc(index)"
                  v-for="(doc, index) in list" :key="index" :doc="doc">
        </doc-card>
        </div>
    </div>
</template>

<script>
    import DocCard from "./docCard";

    export default {
        name: "docs-manager",
        components: {
            DocCard,
        },
        computed: {
            getDocTitle() {
                return this.transferDoc?this.transferDoc.title:""
            }
        },
        data() {
            return {
                list: [],
                total: 0,
                dialogVisible: false,
                transferDoc: null,
            }
        },
        async mounted() {
            let data = await this.$_getDocsList({start: "/doca", length: 10});
            this.list = data.data.list;
            this.total = data.data.total;
        },
        methods: {
            handleDoc(index) {
                this.dialogVisible = true;
                this.transferDoc = this.list[index]
            },
            editAgain() {
                this.$store.commit("changeCurrentDoc", this.transferDoc)
                this.$router.push('/document')
            },
        }
    }
</script>

<style scoped>
    .doc-container {
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
</style>
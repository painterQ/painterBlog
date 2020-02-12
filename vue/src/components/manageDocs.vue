<template>
    <div class="container">
        <el-dialog
                title="提示"
                :visible.sync="dialogVisible"
                width="30%">
            是否要编辑《{{getDoc}}》
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="editAgain">编 辑</el-button>
            </span>
        </el-dialog>
        <doc-card @click.native="handleDoc(index)"
                  v-for="(doc, index) in list" :key="index" :doc="doc">
        </doc-card>
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
            getDoc() {
                if (this.$store.state.currentDoc && this.$store.state.currentDoc.title) {
                    return this.$store.state.currentDoc.title
                }
                return "null"
            }
        },
        data() {
            return {
                list: [],
                total: 0,
                dialogVisible: false,
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
                this.$store.commit("changeCurrentDoc", this.list[index])
            },
            editAgain() {
                this.$router.push('/document')
            },
        }
    }
</script>

<style scoped>
    .container {
        width: 80%;
        padding: 1em;
        margin: 1em auto;
        background-color: white;
        overflow-y: scroll;
        display: flex;
        flex-wrap: wrap;
    }
</style>
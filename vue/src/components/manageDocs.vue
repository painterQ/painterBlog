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
        <div class="manage-card card-ex" @click="handleDoc(index)" v-for="(doc, index) in list" :key="index">
            <h3>{{doc.title}}</h3>
            <div>{{doc.subTitle}}</div>
            <div>id:{{doc.id}}</div>
            <painter-tag v-for="(t,tindex) in doc.tags" :key="tindex">{{t}}</painter-tag>
            <p>{{doc.abstract}}</p>
            <div>
                <span>time: {{new Date(doc.lastTime) | moment}}</span>
                <span v-if="doc.attr === 1">置顶</span>
            </div>
        </div>
    </div>
</template>

<script>
    import painterTag from './tag'

    export default {
        name: "docs-manager",
        components: {
            painterTag
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
        width: 100%;
        margin: 0.5em;
        overflow-y: scroll;
        height: 100%;
    }

    .card-ex {
        padding: 0.5em;
        margin: 0.2em auto;
        max-width: 900px;
        width: 80%;
        min-width: 300px;
    }
</style>
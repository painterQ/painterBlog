<template>
    <div>
        <el-divider class="line" content-position="left">{{this.title}}</el-divider>
        <router-view></router-view>
    </div>
</template>

<script>
    import vue from 'vue'
    import {Drawer} from "element-ui";
    vue.use(Drawer);

    export default {
        name: 'background-manage',
        components: {},
        data() {
            return {

            }
        },
        computed: {
            title(){
                let path = this.$route.path.substr(8);
                switch (path) {
                    case 'image':
                        return "管理图片";
                    case 'tags':
                        return "管理标签";
                    default:
                        return "管理文章";
                }
            }
        },
        methods: {},
    }
</script>

<style scoped>
    .line {
        margin: 2em 0 1em 0;
    }

    .line > * {
        font-size: 1.5em;
        background-color: #fafafa;
    }
</style>

<style>
    .manage-card {
        border-radius: 4px;
        width: 100%;
        overflow-y: hidden;
        /*padding: 1em;*/
        /*margin: 6px 30px;*/
        background-color: #fff;
        box-shadow: 0 0 4px rgba(0, 0, 0, 0.1);
        box-sizing: border-box;
    }

    .manage-card:hover {
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.3);
        position: relative;
        top: 1px;
    }

    .manage-card > * {
        display: block;
        cursor: pointer;
    }
</style>
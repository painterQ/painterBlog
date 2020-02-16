<template>
    <el-menu :default-active="active" class="menu"
             @open="handleOpen" @close="handleClose" @select="handleSelect"
             background-color="#545c64" text-color="#fff" :collapse="isCollapse">
        <el-menu-item index="1">
            <i class="el-icon-setting"></i>
            <span slot="title">设置</span>
        </el-menu-item>
        <el-menu-item index="2">
            <i class="el-icon-document"></i>
            <span slot="title">写文章</span>
        </el-menu-item>
        <el-submenu index="3">
            <template slot="title">
                <i class="el-icon-menu"></i>
                <span slot="title">管理</span>
            </template>
            <el-menu-item index="3-1">管理文章</el-menu-item>
            <el-menu-item index="3-2">管理图片</el-menu-item>
            <el-menu-item index="3-3">管理标签</el-menu-item>
        </el-submenu>
        <el-menu-item index="4">
            <i class="el-icon-location"></i>
            <span slot="title">消息</span>
        </el-menu-item>
        <div id="aside-switch" @click="switchAside">|</div>
        <div class="footer">{{$store.state.author.IPC}}</div>
    </el-menu>
</template>


<script>
    import Vue from 'vue'
    import {Menu, Submenu, MenuItem} from "element-ui";

    let map = [
        {path: '/home', index: '1'},
        {path: '/document', index: '2'},
        {path: '/docs', index: '3-1'},
        {path: '/image', index: '3-2'},
        {path: '/tags', index: '3-3'},
    ]
    Vue.use(Menu);
    Vue.use(MenuItem);
    Vue.use(Submenu);
    export default {
        name: 'painter-aside',
        data() {
            return {
                isCollapse: true,
                active: "1"
            }
        },
        methods: {
            handleOpen(key, keyPath) {
                console.log(key, keyPath);
            },
            handleClose(key, keyPath) {
                console.log(key, keyPath);
            },
            handleSelect(index) {
                let currentPath = this.$route.path
                let path = "";
                for (let e of map) {
                    if (e.index === index) {
                        path = e.path
                    }
                }
                if (path === '' || currentPath === path) return;
                try {
                    let replace = this.$router.history.getCurrentLocation() === "/404"
                    /*learn: 向 history 栈添加一个新的记录*/
                    /*当用户点击浏览器后退按钮时，则回到之前的 URL*/
                    replace ?
                        this.$router.replace(path) :
                        this.$router.push(path);
                } catch (e) {
                    console.log(e)
                }
            },
            switchAside() {
                this.isCollapse = !this.isCollapse
            },
        },
        watch: {
            "$route.path": {
                handler(newPath){
                    let index = "";
                    for (let e of map) {
                        if (e.path === newPath) {
                            index = e.index
                        }
                    }
                    if (index !== "") {
                        this.active = index
                    }
                },
                immediate: true
            }
        }
    }
</script>

<style scoped>
    /*router-link的路由匹配成功后自动添加下面的CSS Class*/
    .router-link-active {

    }

    /*.menu > * {*/
    /*    width: 20vw;*/
    /*}*/

    /*learn: position 相对于上一个非static的祖先元素*/
    /*learn: z-index 越大越靠近用户*/
    #aside-switch {
        position: absolute;
        width: 20px;
        height: 60px;
        top: 49%;
        right: -10px;
        z-index: 1;
        color: #f0f0f0;
    }

    /*learn: 伪类hover cursor设置鼠标指针形状*/
    #aside-switch:hover {
        cursor: pointer;
        color: #3399ff;
    }

    .footer {
        padding: 0.5em;
        position: absolute;
        bottom: 0;
        text-align: center;
        font-size: 0.5em;
        box-shadow: 0 -2px 4px rgba(0, 0, 0, 0.1);
        z-index: 10;
        width: 100%;
    }

    .menu:not(.el-menu--collapse) {
        width: 20vw;
    }

    .menu {
        height: 100%;
    }
</style>
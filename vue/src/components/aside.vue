<template>
    <el-menu
            class="menu"
            ref="menu"
            @open="handleOpen"
            @close="handleClose"
            @select="handleSelect"
            background-color="#545c64"
            text-color="#fff">
        <!--active-text-color="#ffd04b" 加在css中是不管用的--->
        <el-menu-item index="1">
            <i class="el-icon-setting"></i>
            <span slot="title">设置</span>
        </el-menu-item>
        <el-menu-item index="2">
            <i class="el-icon-document"></i>
            <span slot="title">写文章</span>
        </el-menu-item>
        <el-menu-item index="3">
            <i class="el-icon-menu"></i>
            <span slot="title">管理</span>
        </el-menu-item>
        <el-menu-item index="4">
            <i class="el-icon-location"></i>
            <span slot="title">消息</span>
        </el-menu-item>
        <div id="aside-switch" @click="switchAside">|</div>
        <div class="footer">Copyright Painter Qiao.<br/> 浙ICP备19042136号</div>
    </el-menu>
</template>


<script>
    import Vue from 'vue'
    import {Menu, Submenu, MenuItemGroup, MenuItem} from "element-ui";

    Vue.use(Menu);
    Vue.use(MenuItem);
    Vue.use(MenuItemGroup);
    Vue.use(Submenu);
    export default {
        name: 'painter-aside',
        // data(){
        //   return{
        //   }
        // },
        methods: {
            handleOpen(key, keyPath) {
                console.log(key, keyPath);
            },
            handleClose(key, keyPath) {
                console.log(key, keyPath);
            },
            handleSelect(index, indexPath) {
                try {
                    let replace = this.$router.history.getCurrentLocation() === "/404"
                    switch (index) {
                        case "1":
                            replace ?
                                this.$router.replace('/home') :
                                this.$router.push('/home');
                            this.$store.commit('changeIndex',index);
                            return;
                        case "2":
                            /*learn: 向 history 栈添加一个新的记录*/
                            /*当用户点击浏览器后退按钮时，则回到之前的 URL*/
                            replace ?
                                this.$router.replace('/document') :
                                this.$router.push('/document');
                            this.$store.commit('changeIndex',index);
                            return;
                        case "3":
                            replace ?
                                this.$router.replace('/manage/docs') :
                                this.$router.push('/manage/docs');
                            this.$store.commit('changeIndex',index);
                            return;
                        default:
                            console.log(index, indexPath);
                            break;
                    }
                } catch (e) {
                    console.log(e)
                }
            },
            switchAside(){
                this.$store.commit('changeAsideSwitch');
            },
        },
        /*learn: computed watch*/
        watch:{
            "$store.state.viewState.asideSwitch":function(){
                let m = this.$refs['menu'].$el;
                let ec = m.getElementsByTagName("span")
                if(!this.$store.state.viewState.asideSwitch){
                    for(let e of ec){
                        e.style.display = "none";
                        e.parentElement.style.width = 'auto';
                    }
                    return null;
                }
                for(let e of ec){
                    e.style.display = "inline";
                    e.parentElement.style.width = '20vw';
                }
                return null;
            }
        },
        mounted(){
            this.$refs['menu'].defaultActive = this.$store.state.viewState.index
        }
    }
</script>

<style scoped>
    /*router-link的路由匹配成功后自动添加下面的CSS Class*/
    .router-link-active {

    }

    .menu {
        height: 100%;
    }

    .menu > *{
        width: 20vw;
    }

    /*learn: position 相对于上一个非static的祖先元素*/
    /*learn: z-index 越大越靠近用户*/
    #aside-switch{
        position: absolute;
        width: 20px;
        height: 60px;
        top: 49%;
        right: -10px;
        z-index: 1;
        color: #f0f0f0;
    }

    /*learn: 伪类hover cursor设置鼠标指针形状*/
    #aside-switch:hover{
        cursor:pointer;
        color: #3399ff;
    }

    .footer{
        height: 40px;
        position: absolute;
        bottom: 0;
        text-align: center;
        font-size: 0.5em;
        box-shadow : 0 -2px 4px rgba(0,0,0,0.1);
        z-index: 10;
    }
</style>
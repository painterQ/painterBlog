<template>
    <div class="painter-header">
        <div :style="logo"/>
        <div class="header-profile">
            <el-form class="header-search">
                <el-input placeholder="搜索..." prefix-icon="el-icon-search"></el-input>
            </el-form>

            <el-dropdown @command="logout">
                <el-avatar
                        class="header-avatar"
                        shape="circle"
                        ref="avatar"
                        :size="40"
                        fit="contain"
                        :src="avatar"
                        @error="avatarError"
                ></el-avatar>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item icon="el-icon-warning-outline" command="logout">退出登出</el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
        </div>
    </div>
</template>

<script>
    import api from '../api/rpc'
    import Vue from 'vue'
    import {Dropdown, DropdownItem, DropdownMenu} from "element-ui"
    import VueCookies from 'vue-cookies'
    import constVar from '../api/const'
    Vue.use(VueCookies);
    Vue.use(Dropdown);
    Vue.use(DropdownItem);
    Vue.use(DropdownMenu);
    export default {
        name: "index-body",
        data: function () {
            return {
                logo: {
                    backgroundImage: "url(" + require("@/assets/logo.png") + ")",
                    backgroundRepeat: "no-repeat",
                    backgroundSize: "100% 100%",
                    height: "60px",
                    width: "154px",
                    margin: "0 2em 0 2em"
                },
            }
        },
        computed: {
            avatar: {
                get() {
                    console.log("header get avatar", this.$store.state.author.avatar)
                    return this.$store.state.author.avatar
                }
            },
        },
        methods: {
            avatarError() {
                //todo
                return true
            },
            logout(command) {
                switch (command) {
                    case "logout":
                        console.log(this);
                        this.$cookies.remove(constVar.cookieKey)
                        this.$store.commit("changeLogin", false)
                        console.log(this.$store.state.login, this.$cookies.get(constVar.cookieKey))
                        break
                }
            }
        },
        mounted() {
            api.getAuthorInfo().then(
                (res) => {
                    this.$store.commit("changeAvatar", res.data.avatar);
                    this.$store.commit("changeName", res.data.name);
                    this.$store.commit("changeMotto", res.data.say);
                    this.$store.commit("changeIPC", res.data.ipc);
                    this.$store.commit("changeTitle", res.data.title);
                    this.$store.commit("changeSubTitle", res.data.subTitle);
                    this.$store.commit("changeMail", res.data.email);
                    this.$store.commit("changeGithub", res.data.github);
                }
            )
        }
    }
</script>

<style scoped>
    .painter-header {
        justify-content: space-between;
        display: flex;
        flex-wrap: wrap;
        /*learn: 渐变色，可以指定多个节点，可以追加百分比*/
        background: linear-gradient(to right, #545c64 20%, #fff);
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1)
    }

    .header-profile {
        width: 30%;
        display: flex;
    }

    .header-profile > * {
        margin: 10px 1em;
    }

</style>
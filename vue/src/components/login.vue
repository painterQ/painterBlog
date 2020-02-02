<template>
    <div class="login_occlusion_layer" @click.stop=""
         :style="{display: display}">
        <el-form id="login" ref="loginForm" :model="this" :rules="rules">
            <el-form-item label="邮箱" prop="mail" style="font-size: 1.5em">
                <el-input prefix-icon="el-icon-message" v-model="mail"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="pwd" style="font-size: 1.5em">
                <el-input prefix-icon="el-icon-lock" v-model="pwd" :type="lock?'password':'text'">
                    <template slot="append">
                        <i :class="lock?'el-icon-key':'el-icon-view'" @click="lock=!lock" id="lock"></i>
                    </template>
                </el-input>
            </el-form-item>
            <div class="login_button_wrap">
                <el-button @click="loginClear" class="login_button_center">清 空</el-button>
                <el-button type="primary" @click="loginSubmit" class="login_button_center">登 录</el-button>
            </div>
        </el-form>
    </div>
</template>

<script>
    import vue from 'vue'
    import {Dialog} from 'element-ui'
    import message from "../api/message";
    import api from "../api/rpc";
    import VueCookies from 'vue-cookies'
    import constVar from '../api/const'

    vue.use(api);
    vue.use(VueCookies);
    vue.use(Dialog);
    export default {
        name: "painter-login",
        data: function () {
            return {
                mail: "",
                pwd: "",
                lock: true,
                rules: {
                    mail: [
                        {required: true, message: '请输入邮箱地址', trigger: 'change'},
                    ],
                    pwd: [
                        {required: true, message: '请输入密码', trigger: 'change'},
                    ],

                }
            }
        },
        methods: {
            loginClear() {
                this.mail = '';
                this.pwd = '';
                this.$refs["loginForm"].resetFields();
            },
            loginSubmit() {
                this.$refs["loginForm"].validate(async valid => {
                    if (valid) {
                        let res = await this.$_login({'name': this.mail, 'password': this.pwd})
                        if (res.data.status === 1) {
                            this.$store.commit("changeLogin", true);
                            console.log("console.log(commit(\"changeLogin\", true);)\n")
                        } else {
                            message(this, "登录失败:" + res.data.message);
                        }
                    }
                });
            },
        },
        computed: {
            display() {
                return this.$store.state.login ? 'none' : 'flex'
            }
        },
        mounted() {
            let cookie = this.$cookies.get(constVar.cookieKey);
            console.log("cookie", cookie);
            if (cookie && cookie.indexOf("clear") < 0) {
                this.$store.commit("changeLogin", true);
            }
        }
    }
</script>

<style scoped>
    .login_occlusion_layer {
        position: fixed;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        overflow: auto;
        margin: 0;
        z-index: 2000;
        background-color: #30303099;
        /*实现垂直居中*/
        align-items: center;
        /*实现水平居中*/
        justify-content: center;
    }

    #login {
        background-color: white;
        display: block;
        overflow: scroll;
        border-radius: 2px;
        box-shadow: 0 1px 3px rgba(0, 0, 0, .3);
        box-sizing: border-box;
        padding: 1em 2em;
        width: 15em;
        max-width: 100%;
        min-height: 15px;
    }

    #lock {
        box-sizing: border-box;
    }

    #lock:hover {
        cursor: pointer;
    }

    .login_button_wrap {
        display: flex;
        flex-direction: row;
        justify-content: space-around;
        flex-wrap: nowrap;
    }

    .login_button_center {
        max-width: 10em;
    }
</style>
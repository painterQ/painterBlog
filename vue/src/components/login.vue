<template>
    <el-dialog
            title="登 录"
            :visible.sync="!this.$store.state.login"
            width="30%"
            id="dialog"
            style="font-size: 1.5em;overflow:scroll"
            :close-on-click-modal="false"
            :show-close="false"
            center
            destroy-on-close="true">
        <el-form id="login" :rules="rules">
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
<!--            <motto></motto>-->
            <el-button @click="loginClear">清 空</el-button>
            <el-button type="primary" @click="loginSubmit">登 录</el-button>
        </el-form>
    </el-dialog>
</template>

<script>
    import vue from 'vue'
    import {Dialog} from 'element-ui'
    import message from "../api/message";
    import api from "../api/rpc";
    // import Motto from "./motto";

    vue.use(Dialog);
    export default {
        name: "painter-login",
        // components: {Motto},
        data: function () {
            return {
                mail: "",
                pwd: "",
                lock: true,
                rules: {
                    mail: [
                        {required: true, message: '请输入邮箱地址', trigger: 'change'},
                        // {type: "email", message: '不符合邮箱格式', trigger: 'change'}
                    ],
                    pwd: [
                        {required: true, message: '请输入密码', trigger: 'change'},
                        {min: 6, message: '长度不小于6', trigger: 'change'},
                    ],

                }
            }
        },
        methods: {
            loginClear() {
                this.mail = '';
                this.pwd = '';
            },
            loginSubmit() {
                api.login({'name': this.mail, 'password': this.pwd}).then(
                    (res) => {
                        if (res.data.status === 1){
                            this.$store.commit("changeLogin", true);
                        }else {
                            message(this, "登录失败:" + res.data.message);
                        }
                    }
                ).catch(e => {
                    message(this, "登录失败:" + e, "error");
                })
            },
        }
    }
</script>

<style scoped>
    #lock {
        box-sizing: border-box;
    }

    #lock:hover {
        cursor: pointer;
    }

    #login {
        height: 40vh;
    }

</style>
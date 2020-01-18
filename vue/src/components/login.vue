<template>
    <el-dialog
            title="登 录"
            :visible.sync="!this.$store.state.login"
            width="30%"
            id="dialog"
            style="font-size: 1.5em"
            :close-on-click-modal="false"
            :show-close="false"
            center>
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
            <motto></motto>
        </el-form>
        <span slot="footer" class="dialog-footer">
                    <el-button @click="loginClear">清 空</el-button>
                    <el-button type="primary" @click="loginSubmit">登 录</el-button>
                </span>
    </el-dialog>
</template>

<script>
    import vue from 'vue'
    import {Dialog} from 'element-ui'
    import message from "../api/message";
    import api from "../api/rpc";
    import Motto from "./motto";

    vue.use(Dialog);
    export default {
        name: "painter-login",
        components: {Motto},
        data: function(){
          return{
              mail: "",
              pwd: "",
              lock: true,
              rules: {
                  mail: [
                      {required: true, message: '请输入邮箱地址', trigger: 'blur'},
                      {type: "email", message: '不符合邮箱格式', trigger: 'blur'}
                  ],
                  pwd: [
                      {required: true, message: '请输入密码', trigger: 'blur'},
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
                api.login({'mail':this.mail, 'password':this.pwd}).then(
                    ()=>{this.$store.commit("changeLogin", true);}
                ).catch(e=>{
                    message(this, "登录失败:"+e, "error");
                })
            },
        }
    }
</script>

<style scoped>
    #lock{
        box-sizing: border-box;
    }

    #lock:hover{
        cursor:pointer;
    }

    #login{
        height: 40vh;
    }

    #motto > p{
        text-align: center;
        font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
        color: #1a222f;
    }

    #motto span{
        display: block;
        margin: 0.5em 2em;
        text-align: right;
        font-family: 'Avenir', Helvetica, Arial, sans-serif ;
        font-style: italic;
        font-size: 0.8em;
    }
</style>
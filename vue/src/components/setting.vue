<template>
    <div class="setting">
        <!--learn: 通过给元素绑定ref=“XXX”，然后通过this.$refs.XXX或者this.refs['XXX']来获取元素-->
        <el-form ref="baseInfoForm" :model="baseInfoForm" label-width="80px" :rules="rules" status-icon>
            <el-divider class="line" content-position="left">个人信息</el-divider>
            <div id="owner-info">
                <el-upload
                        class="avatar-uploader"
                        action="https://jsonplaceholder.typicode.com/posts/"
                        :show-file-list="false"
                        :on-success="handleAvatarSuccess"
                        :before-upload="beforeAvatarUpload">
                    <img v-if="imageUrl" :src="imageUrl" class="avatar">
                    <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                </el-upload>
                <div>
                    <el-form-item label="个人邮箱" prop="mail">
                        <span>用于发送告警邮件及其它通知, 建议填写, 如: example@163.com.</span>
                        <el-input v-model="baseInfoForm.mail" prefix-icon="el-icon-message"></el-input>
                    </el-form-item>
                    <el-form-item label="github" prop="github">
                        <span>选择填写, 如: https://github.com/painterQ</span>
                        <el-input v-model="baseInfoForm.github">
                            <template slot="prepend">Http://github.com/</template>
                        </el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="submitForm('baseInfoForm')">提交</el-button>
                        <el-button>取消</el-button>
                    </el-form-item>
                </div>
            </div>
        </el-form>
        <el-form ref="blogInfoForm" :model="blogInfoForm" label-width="80px" :rules="rules" status-icon>
            <el-divider class="line" content-position="left">博客信息</el-divider>
            <el-form-item label="博客昵称" prop="nickName">
                <el-input v-model="blogInfoForm.nickName"></el-input>
                <span>用户昵称可以与用户名不同, 用于前台显示.如果你将此项留空, 将默认使用登录用户名.</span>
            </el-form-item>
            <el-form-item label="博客标题" prop="title">
                <el-input v-model="blogInfoForm.title"></el-input>
                <span>用于所有页面的title组成, 如: Painter's Blog</span>
            </el-form-item>
            <el-form-item label="座右铭" prop="motto">
                <el-input v-model="blogInfoForm.motto" type="textarea" autosize></el-input>
                <span>简介或格言, 如: 生活百般滋味, 人生需要笑对.</span>
            </el-form-item>
            <el-form-item label="IPC" prop="ipc">
                <el-input v-model="blogInfoForm.IPC"></el-input>
                <span>用于底部显示, 不添加则不显示</span>
            </el-form-item>
            <el-form-item label="归档前言" prop="archive">
                <el-input v-model="blogInfoForm.beforeArchive" type="textarea" autosize></el-input>
                <span>此文字用于专题前述, 会在专题最前方显示.</span>
            </el-form-item>
            <el-form-item label="专题前言" prop="topic">
                <el-input v-model="blogInfoForm.beforeTopic" type="textarea" autosize></el-input>
                <span>此文字用于归档前述, 会在归档最前方显示.</span>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm('blogInfoForm')">提交</el-button>
                <el-button>取消</el-button>
            </el-form-item>
        </el-form>

        <el-form ref="pwdChangeForm" :model="pwdChangeForm" label-width="80px" :rules="rules" status-icon>
            <el-divider class="line" content-position="left">密码修改</el-divider>
            <el-form-item label="密码" prop="pwd">
                <el-input v-model="pwdChangeForm.pwd" type="password" auto-complete="off"></el-input>
            </el-form-item>
            <el-form-item label="密码确认" prop="repwd">
                <el-input v-model="pwdChangeForm.rePWD" type="password" auto-complete="off">
                </el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm('pwdChangeForm')">提交</el-button>
                <el-button>取消</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    import vue from 'vue'
    import {
        Form, FormItem, Select, Option, OptionGroup,
        Input, Button, Checkbox, CheckboxGroup, Switch,
        Avatar, Divider,Upload
    } from 'element-ui'
    import {changeBaseInfo, changeBlogInfo, changePwdChange} from "../api/rpc";

    {
        vue.use(Form);
        vue.use(FormItem);
        vue.use(Select);
        vue.use(Option);
        vue.use(OptionGroup);
        vue.use(Input);
        vue.use(Button);
        vue.use(CheckboxGroup);
        vue.use(Checkbox);
        vue.use(Avatar);
        vue.use(Switch);
        vue.use(Divider);
        vue.use(Upload);
    }

    export default {
        name: "painter-setting",
        data() {
            let validatePass = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请输入密码'));
                } else {
                    if (this.pwdChangeForm.repwd !== '') {
                        this.$refs.pwdChangeForm.validateField('repwd');
                    }
                    callback();
                }
            };
            let validatePass2 = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请再次输入密码'));
                } else if (value !== this.pwdChangeForm.pwd) {
                    callback(new Error('两次输入密码不一致'));
                } else {
                    callback();
                }
            };
            return {
                imageUrl: 'avatar.jpeg',
                baseInfoForm: {
                    mail: '',
                    github: '',
                },
                blogInfoForm: {
                    nickName: '',
                    title: '',
                    motto: '',
                    IPC: '',
                    beforeArchive: '',
                    beforeTopic: '',
                },
                pwdChangeForm: {
                    pwd: '',
                    rePWD: '',
                },
                rules: {
                    mail: [
                        {required: true, message: '请输入邮箱地址', trigger: 'blur'},
                        {type: "email", message: '不符合邮箱格式', trigger: 'blur'}
                    ],
                    nickName: [
                        {max: 10, message: '长度不超过10个字符', trigger: 'change'},
                        {required: true, message: '请输入昵称', trigger: 'blur'},
                    ],
                    title: [
                        {max: 10, message: '长度不超过10个字符', trigger: 'change'},
                        {required: true, message: '请输入博客标题', trigger: 'blur'},
                    ],
                    motto: [
                        {max: 64, message: '长度不超过64个字符', trigger: 'change'},
                    ],
                    archive: [
                        {max: 255, message: '长度不超过255个字符', trigger: 'change'},
                    ],
                    topic: [
                        {max: 255, message: '长度不超过255个字符', trigger: 'change'},
                    ],
                    pwd: [
                        {validator: validatePass, trigger: 'blur'},
                    ],
                    repwd: [
                        {validator: validatePass2, trigger: 'change'},
                    ],
                }
            }
        },

        methods: {
            submitForm(name) {
                switch (name) {
                    case  'baseInfoForm': changeBaseInfo(this.baseInfoForm);return;
                    case  'blogInfoForm': changeBlogInfo(this.blogInfoForm);return;
                    case  'pwdChangeForm': changePwdChange(this.pwdChangeForm);return;
                    default: return;
                }
            },
            avatarError() {
                this.$refs['avatar'].src = `${this.baseUrl}/avatar.jpeg`
            },
            handleAvatarSuccess(res, file) {
                this.imageUrl = URL.createObjectURL(file.raw);
            },
            beforeAvatarUpload(file) {
                const isJPG = file.type === 'image/jpeg';
                const isLt2M = file.size / 1024 / 1024 < 2;

                if (!isJPG) {
                    this.$message.error('上传头像图片只能是 JPG 格式!');
                }
                if (!isLt2M) {
                    this.$message.error('上传头像图片大小不能超过 2MB!');
                }
                return isJPG && isLt2M;
            }
        },

    }
</script>

<style scoped>
    HR {
        background-color: #5CCCCC;
        size: 3px;
        filter: alpha(opacity=100, finishopacity=0, style=3);
        margin: 2em 0;
    }

    .avatar-uploader .el-upload {
        border: 1px dashed #d9d9d9;
        border-radius: 6px;
        cursor: pointer;
        position: relative;
        overflow: hidden;
    }
    .avatar-uploader .el-upload:hover {
        border-color: #409EFF;
    }
    .avatar-uploader-icon {
        font-size: 28px;
        color: #8c939d;
        width: 178px;
        height: 178px;
        line-height: 178px;
        text-align: center;
    }
    .avatar {
        width: 178px;
        height: 178px;
        display: block;
    }

    #owner-info {
        display: flex;
        flex-wrap: wrap;
    }

    .setting {
        margin: 3em auto;
        max-width: 990px;
    }


    span {
        margin: 1em 1em 0 0;
        color: #999;
    }

    .line {
        margin: 2em 0;
    }

    .line > * {
        font-size: 1.5em;
        font-size: large;
        background-color: #fafafa;;
    }
</style>
<template>
    <div class="setting">
        <!--learn: 通过给元素绑定ref=“XXX”，然后通过this.$refs.XXX或者this.refs['XXX']来获取元素-->
        <el-form ref="baseInfoForm" label-width="80px" status-icon :model="this" :rules="rules">
            <el-divider class="line" content-position="left">个人信息</el-divider>
            <div id="owner-info">
                <el-upload
                        class="avatar-uploader"
                        action="https://jsonplaceholder.typicode.com/posts/"
                        :show-file-list="false"
                        @success="this.handleAvatarSuccess"
                        @error="this.avatarError"
                        :before-upload="this.beforeAvatarUpload">
                    <img v-if="this.avatar" :src="this.avatar" class="avatar">
                    <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                </el-upload>
                <div>
                    <el-form-item label="个人邮箱" prop="mail">
                        <span>用于发送告警邮件及其它通知, 建议填写, 如: example@163.com.</span>
                        <el-input v-model="mail" prefix-icon="el-icon-message"></el-input>
                    </el-form-item>
                    <el-form-item label="github">
                        <span>选择填写, 如: https://github.com/painterQ</span>
                        <el-input v-model="github">
                            <template slot="prepend">Http://github.com/</template>
                        </el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="submitForm('baseInfoForm')">提交</el-button>
                        <el-button @click="clear('baseInfoForm')">取消</el-button>
                    </el-form-item>
                </div>
            </div>
        </el-form>
        <el-form ref="blogInfoForm" label-width="80px" status-icon :model="this" :rules="rules">
            <el-divider class="line" content-position="left">博客信息</el-divider>
            <el-form-item label="博客昵称" prop="name">
                <el-input v-model="name"></el-input>
                <span>用户昵称可以与用户名不同, 用于前台显示.如果你将此项留空, 将默认使用登录用户名.</span>
            </el-form-item>
            <el-form-item label="博客标题" prop="title">
                <el-input v-model="title"></el-input>
                <span>用于所有页面的title组成, 如: Painter's Blog</span>
            </el-form-item>
            <el-form-item label="格言" prop="motto">
                <el-input v-model="motto" type="textarea" autosize></el-input>
                <span>格言, 如: 生活百般滋味, 人生需要笑对.</span>
            </el-form-item>
            <el-form-item label="副标题" prop="subTitle">
                <el-input v-model="subTitle" type="textarea" autosize></el-input>
                <span>介绍博客内容, 如: 为了爱与和平.</span>
            </el-form-item>
            <el-form-item label="IPC" prop="ipc">
                <el-input v-model="IPC"></el-input>
                <span>用于底部显示, 不添加则不显示</span>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm('blogInfoForm')">提交</el-button>
                <el-button @click="clear('blogInfoForm')">取消</el-button>
            </el-form-item>
        </el-form>

        <el-form ref="pwdChangeForm" label-width="80px" status-icon :model="this" :rules="rules">
            <el-divider class="line" content-position="left">密码修改</el-divider>
            <el-form-item label="密码" prop="pwd">
                <el-input v-model="pwd" type="password" auto-complete="off"></el-input>
            </el-form-item>
            <el-form-item label="密码确认" prop="rePWD">
                <el-input v-model="rePWD" type="password" auto-complete="off">
                </el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm('pwdChangeForm')">提交</el-button>
                <el-button @click="clear('pwdChangeForm')">取消</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    import vue from 'vue'
    import api from '../api/rpc'
    import message from "../api/message";
    vue.use(message);
    vue.use(api);
    import {
        Form, FormItem, Select, Option, OptionGroup,
        Input, Checkbox, CheckboxGroup, Switch,
        Avatar, Divider, Upload
    } from 'element-ui'

    {
        vue.use(Form);
        vue.use(FormItem);
        vue.use(Select);
        vue.use(Option);
        vue.use(OptionGroup);
        vue.use(Input);
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
            return {
                pwd: '',
                rePWD: '',
                rules: {
                    mail: [
                        {required: true, message: '请输入邮箱地址'},
                        {type: "email", message: '不符合邮箱格式'}
                    ],
                    name: [
                        {max: 10, message: '长度不超过10个字符', trigger: 'change'},
                        {required: true, message: '请输入昵称', trigger: 'blur'},
                    ],
                    title: [
                        {max: 10, message: '长度不超过10个字符', trigger: 'change'},
                        {required: true, message: '请输入博客标题', trigger: 'blur'},
                    ],
                    subTitle: [
                        {required: true, message: '请输入博客副标题', trigger: 'blur'},
                    ],
                    motto: [
                        {max: 64, message: '长度不超过64个字符', trigger: 'change'},
                    ],
                    pwd: [
                        {required: true, message: '请输入密码', trigger: 'blur'},
                    ],
                    rePWD: [
                        {
                            validator: (rule, value, callback) => {
                                console.log("validator", value)
                                if (value === '') {
                                    callback(new Error('请再次输入密码'));
                                } else if (value !== this.pwd) {
                                    callback(new Error('两次输入密码不一致!'));
                                } else {
                                    callback();
                                }
                            }, trigger: 'change'
                        },
                    ],
                }
            }
        },
        computed: {
            avatar: {
                get() {
                    return this.$store.state.author.avatar
                },
                set(v) {
                    this.$store.commit("changeAvatar", v);
                }
            },
            mail: {
                get() {
                    return this.$store.state.author.mail
                },
                set(v) {
                    this.$store.commit("changeMail", v);
                }
            },
            github: {
                get() {
                    return this.$store.state.author.github
                },
                set(v) {
                    this.$store.commit("changeGithub", v);
                }
            },
            name: {
                get() {
                    return this.$store.state.author.name
                },
                set(v) {
                    this.$store.commit("changeName", v);
                }
            },
            title: {
                get() {
                    return this.$store.state.author.title
                },
                set(v) {
                    this.$store.commit("changeTitle", v);
                }
            },
            subTitle: {
                get() {
                    return this.$store.state.author.subTitle
                },
                set(v) {
                    this.$store.commit("changeSubTitle", v);
                }
            },
            motto: {
                get() {
                    return this.$store.state.author.motto
                },
                set(v) {
                    this.$store.commit("changeMotto", v);
                }
            },
            IPC: {
                get() {
                    return this.$store.state.author.IPC
                },
                set(v) {
                    this.$store.commit("changeIPC", v);
                }
            },
        },
        methods: {
            submitForm(name) {
                switch (name) {
                    case  'baseInfoForm':
                        this.$refs['baseInfoForm'].validate((valid) => {
                            if (valid) {
                                this.$_changeBaseInfo({
                                    mail: this.mail,
                                    github: this.github,
                                });
                            }
                        });
                        return;
                    case  'blogInfoForm':
                        this.$refs['blogInfoForm'].validate((valid) => {
                            if (valid) {
                                this.$_changeBlogInfo({
                                    name: this.name,
                                    title: this.title,
                                    subTitle: this.motto,
                                    IPC: this.IPC,
                                });
                            }
                        });
                        return;
                    case  'pwdChangeForm':
                        this.$refs["pwdChangeForm"].validate((valid) => {
                            if (valid) {
                                this.$_changePwdChange({
                                    pwd: this.pwd,
                                });
                            }
                        });
                        return;
                    default:
                        return;
                }
            },
            async clear(name) {
                if (name === 'pwdChangeForm') {
                    this.pwd = '';
                    this.rePWD = '';
                    this.$refs["pwdChangeForm"].resetFields();
                } else if (name === 'baseInfoForm' || name === 'blogInfoForm') {
                    let res = await this.$_getAuthorInfo();
                    switch (name) {
                        case  'baseInfoForm':
                            this.$store.commit("changeMail", res.data.email);
                            this.$store.commit("changeGithub", res.data.github);
                            break;
                        case  'blogInfoForm':
                            this.$store.commit("changeAvatar", res.data.avatar);
                            this.$store.commit("changeName", res.data.name);
                            this.$store.commit("changeMotto", res.data.say);
                            this.$store.commit("changeIPC", res.data.ipc);
                            this.$store.commit("changeTitle", res.data.title);
                            this.$store.commit("changeSubTitle", res.data.subTitle);
                            break
                    }
                }
            },
        },
        avatarError() {
            this.$refs['avatar'].src = `${this.baseUrl}/static/avatar.jpeg`
        },
        handleAvatarSuccess(res, file) {
            this.avatar = URL.createObjectURL(file.raw);
        },
        beforeAvatarUpload(file) {
            const supportType = ['image/jpeg','image/png','image/gif'].indexOf(file.type) > -1
            const sizeLimit = file.size / 1024 / 1024 < 3;

            if (!supportType) {
                message.message(this,'Only support jpeg/png/gif','warning');
            }
            if (!sizeLimit) {
                message.message(this,'Image size limit: 3MByte','warning');
            }
            return supportType && sizeLimit;
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
        background-color: #fafafa;;
    }
</style>
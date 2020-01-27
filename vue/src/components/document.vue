<template>
    <div id="document">
        <el-divider class="line" content-position="left">撰写文章</el-divider>
        <el-form ref="form" label-width="80px" id="tinymce-editor-form">
            <div class="operation-flex-container">
                <el-button id="release" @click="this.release">发布</el-button>
                <el-button id="save">保存草稿</el-button>
                <el-button id="delete">删除</el-button>
                <el-button id="more" @click="this.more">更多选项</el-button>
                <el-input type="text" placeholder="请输入标题"
                          id="title" maxlength="30" show-word-limit
                          v-model="title"/>
            </div>
            <el-input type="text" id="path" v-model="path" class="editor-input"
                      placeholder="请输入路径">
                <template slot="prepend">http://www.xixi201314.cn/</template>
                <template slot="append">.html</template>
            </el-input>
            <editor v-model="myValue"
                    :init="init"
                    :disabled="disabled"
                    name="document">
            </editor>
            <!--            <el-input ref="submit" type="submit"></el-input>-->
        </el-form>
        <el-drawer
                :visible.sync="drawer"
                :with-header="false">
            <el-form class="document-drawer">
                <p style="margin: 0">标签</p>
                <el-select v-model="tag" multiple placeholder="选择标签" class="document-drawer-input">
                    <el-option
                            v-for="tag of commonTags"
                            :key="tag"
                            :value="tag">
                    </el-option>
                </el-select>

                <el-input type="text" v-model="newTag" class="document-drawer-input"
                          placeholder="新建tag" maxlength="10"
                          show-word-limit>
                    <template slot="append">
                        <i class="el-icon-edit drawer-append-edit" @click="addTag"></i>
                    </template>
                </el-input>
                <hr/> <p style="margin: 0">摘要</p>
                <el-input type="textarea"  :rows="6" v-model="abstract"
                          placeholder="请输入摘要">
                </el-input>
            </el-form>
        </el-drawer>
    </div>
</template>
<script>
    import 'tinymce/tinymce'
    import Editor from '@tinymce/tinymce-vue'
    import api from '../api/rpc'
    import 'tinymce/themes/silver'
    import Vue from 'vue'
    // 编辑器插件plugins
    // 更多插件参考：https://www.tiny.cloud/docs/plugins/
    import 'tinymce/plugins/image'// 插入上传图片插件
    import 'tinymce/plugins/media'// 插入视频插件
    import 'tinymce/plugins/table'// 插入表格插件
    import 'tinymce/plugins/lists'// 列表插件
    import 'tinymce/plugins/wordcount'// 字数统计插件
    import {Drawer, Tag} from "element-ui";

    Vue.use(Drawer);
    Vue.use(Tag);
    export default {
        components: {
            Editor,
        },
        props: {
            value: {
                type: String,
                default: ''
            },
            // 基本路径，默认为空根目录，如果你的项目发布后的地址为目录形式，
            // 即abc.com/tinymce，baseUrl需要配置成tinymce，不然发布后资源会找不到
            baseUrl: {
                type: String,
                default: ''
            },
            disabled: {
                type: Boolean,
                default: false
            },
            plugins: {
                type: [String, Array],
                default: 'lists image media table wordcount'
            },
            toolbar: {
                type: [String, Array],
                default: 'undo redo |  formatselect | bold italic forecolor backcolor | alignleft aligncenter alignright alignjustify | bullist numlist outdent indent | lists image media table | removeformat'
            }
        },
        data() {
            return {
                drawer: false,
                init: {
                    //todo 这里路径里的sttic导致移植性变差
                    language_url: `${this.baseUrl}/static/tinymce/langs/zh_CN.js`,
                    language: 'zh_CN',
                    skin_url: `${this.baseUrl}/static/tinymce/skins/ui/oxide`,
                    content_css: `${this.baseUrl}/static/tinymce/skins/content/default/content.css`,
                    // skin_url: `${this.baseUrl}/tinymce/skins/ui/oxide-dark`, // 暗色系
                    // content_css: `${this.baseUrl}/tinymce/skins/content/dark/content.css`, // 暗色系
                    height: 500,
                    plugins: this.plugins,
                    toolbar: this.toolbar,
                    branding: false,
                    menubar: false,
                    //粘贴图片
                    paste_data_images: true,
                    //TinyMCE 会将所有的 font 元素转换成 span 元素
                    convert_fonts_to_spans: true,
                    //换行符会被转换成 br 元素
                    convert_newlines_to_brs: true,
                    //在换行处 TinyMCE 会用 BR 元素而不是插入段落
                    force_br_newlines: false,
                    //当返回或进入 Mozilla/Firefox 时，这个选项可以打开/关闭段落的建立
                    force_p_newlines: false,
                    //这个选项控制是否将换行符从输出的 HTML 中去除。选项默认打开，因为许多服务端系统将换行转换成 <br />，因为文本是在无格式的 textarea 中输入的。使用这个选项可以让所有内容在同一行。
                    remove_linebreaks: false,
                    //不能把这个设置去掉，不然图片路径会出错
                    relative_urls: false,
                    //不允许拖动大小
                    resize: true,
                    // 此处为图片上传处理函数，这个直接用了base64的图片形式上传图片，
                    // 如需ajax上传可参考https://www.tiny.cloud/docs/configure/file-image-upload/#images_upload_handler
                    images_upload_handler: (blobInfo, success, /*failure*/) => {
                        // if (this.allowedFileTypes.indexOf(blobInfo.blob.type) > -1) {
                        uploadPic()
                        // } else {
                        //     console.log('图片格式错误')
                        // }
                        function uploadPic() {
                            const img = 'data:image/jpeg;base64,' + blobInfo.base64()
                            api.uploadImage(img).then((res) => {
                                // 这里返回的是你图片的地址
                                success(res.data.url)
                            }).catch(() => {
                                console.log('上传失败')
                            })
                        }
                    },

                },
                title: '',
                path: '',
                myValue: '',
                abstract: '',
                newTag: "", //新建的tag
                tag: [], //本次选择的tag
                commonTags: [], //常用的tag
                allowedFileTypes: ["image/png", "image/jpeg", "image/gif"],
            }
        },
        methods: {
            // 添加相关的事件，可用的事件参照文档=> https://github.com/tinymce/tinymce-vue => All available events
            // 需要什么事件可以自己增加
            release() {
                if(!this.path.startsWith('/')){
                    this.path = "/" + this.path
                }
                api.postDoc({
                    title: this.title,
                    path: this.path,
                    abstract: this.abstract,
                    tag: this.tag,
                    document: this.myValue,
                })
            },
            // 可以添加一些自己的自定义事件，如清空内容
            clear() {
                this.myValue = ''
            },
            more() {
                this.drawer = true
            },
            addTag(){
                if (this.newTag === ""){
                    return
                }
                //获取常用tag get /docs/tag
                api.addTag([this.newTag]).then(
                    () => {
                        this.tag.push(this.newTag);
                        this.tag  = [...new Set(this.tag)];
                        this.commonTags.push(this.newTag);
                        this.commonTags  = [...new Set(this.commonTags)];
                        this.newTag=""
                    }).catch(() => {
                    console.log('新建tag失败')
                })
            }
        },
        mounted() {
            let self = this
            //获取常用tag get /docs/tag
            api.getTags().then(
                (res) => {
                    self.commonTags = res.data
                }).catch(() => {
                console.log('获取tag失败')
            })
        }
    }
</script>

<style scoped>
    #document {
        max-width: 1024px;
        margin: 3em auto;
    }

    .operation-flex-container {
        display: -webkit-flex;
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;
    }

    .operation-flex-container > div {
        width: 60%;
    }

    #tinymce-editor-form > div {
        margin: 0 0 7px 0;
    }

    .line {
        margin: 2em 0 1em 0;
    }

    .line > * {
        font-size: 1.5em;
        background-color: #fafafa;
    }
    .document-drawer{
        margin: 1em;
    }

    .document-drawer-input{
        margin: 5px 0;
        width: 100%;
    }

    .drawer-append-edit{
        color: #0080ff;
    }
    .drawer-append-edit:hover{
        color: green;
        font-weight: bold;
        cursor: pointer;
    }
</style>
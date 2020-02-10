<template>
    <el-upload
            :action="action"
            class="upload-container"
            name="avatar"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
            :headers="header"
            :before-upload="beforeAvatarUpload">
        <div class="upload-background">
            <div v-if="this.default" :style="getStyle"></div>
            <div v-else :style="getStyle">
                <svg t="1581212999524" class="icon" viewBox="0 0 1024 1024" version="1.1"
                     xmlns="http://www.w3.org/2000/svg" p-id="2099"
                     xmlns:xlink="http://www.w3.org/1999/xlink" width="64" height="64">
                    <defs>
                        <style type="text/css"></style>
                    </defs>
                    <path d="M554.688 500.352v256H469.312v-256h-128L512 314.24l170.688 186.24h-128zM1024 640.192C1024 782.912 919.872 896 787.648 896h-512C123.904 896 0 761.6 0 597.504 0 451.968 94.656 331.52 226.432 302.976 284.16 195.456 391.808 128 512 128c152.32 0 282.112 108.416 323.392 261.12C941.888 413.44 1024 519.04 1024 640.192z m-259.2-205.312c-24.448-129.024-128.896-222.72-252.8-222.72-97.28 0-183.04 57.344-224.64 147.456l-9.28 20.224-20.928 2.944c-103.36 14.4-178.368 104.32-178.368 214.72 0 117.952 88.832 214.4 196.928 214.4h512c88.32 0 157.504-75.136 157.504-171.712 0-88.064-65.92-164.928-144.96-171.776l-29.504-2.56-5.888-30.976z"
                          fill="#0080ff" p-id="2100"></path>
                </svg>
            </div>
        </div>
    </el-upload>
</template>
<script>
    import vue from 'vue'
    import message from "../api/message";

    vue.use(message);
    export default {
        name: "painter-image",
        props: ['default', 'commit', 'callback', 'width', 'height'],
        data() {
            return {
                action: document.location.protocol + `//` +
                    window.location.host + '/docs/image/filter',
            }
        },
        computed: {
            header() {
                if (this.commit) {
                    return {
                        "avatar": true
                    }
                }
                return {}
            },
            getStyle() {
                let width = this.width;
                let height = this.height;
                if (!width || (typeof width !== 'number' && typeof width !== 'string')) {
                    width = '200';
                }
                if (!height || (typeof height !== 'number' && typeof height !== 'string')) {
                    height = '200';
                }
                if (this.default) {
                    //url("image.png") 140px 40px/200px 100px no-repeat content-box #58a;
                    //url("../image/23.png") 50% 50%/200px 200px no-repeat padding-box #ffffff00
                    let url = `url("`+this.default+`") `;
                    let positionAndSize = `50% 50%/` + width + 'px ' + height + 'px';
                    let other = ` no-repeat padding-box #ffffff00`;
                    return {
                        width: width + 'px',
                        height: height + 'px',
                        background: url + positionAndSize + other
                    }
                }
                let w = (Number.parseInt(width) - 64) / 2;
                let h = (Number.parseInt(height) - 64) / 2;
                return {
                    width: width + 'px',
                    height: height + 'px',
                    padding: h + 'px ' + w + 'px',
                    boxSizing: 'border-box'
                }
            }
        },
        methods: {
            handleAvatarSuccess(res) {
                if (this.default !== "" && res.list && res.list[0]) {
                    if (res.list[0].small) {
                        this.default = 'data:image/' + res.type + ';base64,' + res.small
                    } else {
                        this.default = res.src
                    }
                    this.width = res.list[0].width;
                    this.height = res.list[0].height;
                }

                if (this.commit) {
                    this.$store.commit('changeAvatar', this.default)
                }
                message.message(this, 'upload [' + res.name + '] success', 'success');
                if (this.callback && res.list[0]) {
                    let img = JSON.parse(JSON.stringify(res.list[0]))
                    img.src = res.webDN + '/' + res.list[0].src
                    this.callback(img)
                }
            },
            beforeAvatarUpload(file) {
                const supportType = ['image/jpeg', 'image/png', 'image/gif'].indexOf(file.type) > -1
                const sizeLimit = file.size / 1024 / 1024 < 3;
                if (!supportType) {
                    message.message(this, 'Only support jpeg/png/gif', 'warning');
                }
                if (!sizeLimit) {
                    message.message(this, 'Image size limit: 3MByte', 'warning');
                }
                return supportType && sizeLimit;
            },
        },
    }
</script>

<style scoped>

    .upload-background {
        border: 1px dashed #757d87;
        border-radius: 4px;
    }

    .upload-background:hover {
        border: 1px dashed #0080ff;
    }


</style>
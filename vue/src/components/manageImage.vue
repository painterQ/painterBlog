<template>
    <div id="manager" ref="manager">
        <el-drawer
                ref="uploadDrawer"
                :visible.sync="drawer"
                destroy-on-close
                withHeader="false"
                direction="rtl">
            <painter-image :callback="uploadCallback" width="300" height="300"
                           style="width: 300px;margin:0 auto">
            </painter-image>
        </el-drawer>
        <el-button @click="drawer=true" type="primary" style="margin: 1em;">
            上传图片
        </el-button>
        <waterfall :col='5' :data="image" ref="waterfall"
                   @scroll="scroll" :loadDistance="50"
                   class="image-manager-waterfall">
            <template>
                <div class="cart-container" v-for="(item,index) in image" :key="index">
                    <div class="manage-card">
                        <img v-if="getBaseURL(item)" :src="getBaseURL(item)" alt="加载错误" class="card-image">
                        <div style="padding: 14px;">
                            <span :title="item.name" style="word-break: break-all">
                                {{trimName(item.name)}}</span>
                            <div class="card-bottom card-clearfix">
                                <span>type: {{item.type}}   id: {{item.id}} size:{{item.width}}X{{item.height}}</span>
                            </div>
                        </div>
                    </div>
                </div>
            </template>
        </waterfall>
    </div>
</template>

<script>
    import Vue from 'vue'
    import waterfall from 'vue-waterfall2'
    import {Upload} from 'element-ui'
    import painterImage from './Image'

    Vue.use(waterfall);
    Vue.use(Upload);
    export default {
        name: "image-manager",
        components: {
            painterImage
        },
        data: function () {
            return {
                image: [],
                drawer: false,
                getting: false,
            }
        },
        async created() {
            let res = await this._get();
            for (let i in res.data.list) {
                res.data.list[i].src = res.data.webDN + '/' + res.data.list[i].src
            }
            this.image = res.data.list;
            this.$refs.waterfall.$el.addEventListener('scroll', this.scroll, true)
        },
        beforeDestroy() {
            this.$refs.waterfall.$el.removeEventListener('scroll', this.scroll, true)
        },
        methods: {
            _get() {
                //{["list":{"id":"","name":"","type":"","src":"",},],"webDN":""}
                return this.$_getImageList({
                    //index start with 1
                    'start': this.image.length + 1,
                    'limit': 20,
                });
            },
            async scroll() {
                if(this.getting) return;
                let el = this.$refs.waterfall.$el;
                if (el.offsetHeight + el.scrollTop + 30 > el.scrollHeight) {
                    this.getting = true
                    let res = await this._get()
                    for (let i in res.data.list) {
                        res.data.list[i].src = res.data.webDN + '/' + res.data.list[i].src
                        this.image.push(res.data.list[i])
                    }
                    this.$nextTick(()=>{this.getting = false})
                }
            },
            trimName(name) {
                if (name.length > 20) {
                    return name.substr(0, 20) + "..."
                }
                return name
            },
            getBaseURL(imgItem) {
                if (imgItem.small) {
                    return 'data:image/' + imgItem.type + ';base64,' + imgItem.small
                } else {
                    return imgItem.src
                }
            },
            uploadCallback() {
                this._get()
            },
        },
    }
</script>

<style scoped>

    #manager {
        width: 100%;
        max-width: 1024px;
        margin: 0 auto;
        position: relative;
        height: 100%;
    }

    .card-bottom {
        margin-top: 13px;
        line-height: 12px;
    }

    .card-bottom > span {
        font-size: 13px;
        color: #999;
    }

    .card-image {
        width: 100%;
        display: block;
    }

    .card-clearfix:before,
    .card-clearfix:after {
        display: table;
        content: "";
    }

    .card-clearfix:after {
        clear: both
    }

    .cart-container {
        box-sizing: border-box;
        background-color: #ffffff00;
        padding: 5px;
    }

    .image-manager-waterfall {
        height: 80%;
        overflow-y: scroll;
    }
</style>
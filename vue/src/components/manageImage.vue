<template>
    <!--    todo-->
    <div id="manager" ref="manager">
        <waterfall :col='5' :data="image" @loadmore="scrollHandle">
            <template>
                <div class="cart-container" v-for="(item,index) in image" :key="index">
                    <div class="manage-card">
                        <img v-if="item.src" :lazy-src="item.src" alt="加载错误" class="card-image">
                        <div style="padding: 14px;">
                            <span :title="item.name" style="word-break: break-all">
                                {{trimName(item.name)}}</span>
                            <div class="card-bottom card-clearfix">
                                <span>type: {{item.type}}   id: {{item.id}}</span>
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

    Vue.use(waterfall);
    export default {
        name: "image-manager",
        data: function () {
            return {
                image: [],
            }
        },
        async created() {
            let res = await this._get();
            this.image = res.data;
            this.$waterfall.forceUpdate()
        },

        methods: {
            _get() {
                //[{"id":"","name":"","type":"","src":"",}]
                return this.$_getImageList({
                    'start': this.image.length,
                    'limit': this.image.length + 20,
                });
            },
            async scrollHandle() {
                let res = await this._get();
                this.image = this.image.concat(res.data)
            },
            trimName(name) {
                if (name.length > 20) {
                    return name.substr(0, 20) + "..."
                }
                return name
            }
        },
    }
</script>

<style scoped>
    #const-size-container {
        height: 100%;
        width: 100%;
        overflow: scroll;
    }

    #manager {
        width: 100%;
        max-width: 1024px;
        margin: 3em auto;
        position: relative;
        overflow-y: auto;
        overflow-x: visible;
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
</style>
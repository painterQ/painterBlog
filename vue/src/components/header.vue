<template>
    <div class="painter-header">
        <div :style="logo"/>
        <div class="header-profile">
            <el-form class="header-search">
                <el-input placeholder="搜索..." prefix-icon="el-icon-search"></el-input>
            </el-form>
            <el-avatar
                    class="header-avatar"
                    shape="circle"
                    ref="avatar"
                    :size="40"
                    fit="contain"
                    src="./avatar.jpeg"
                    @error="avatarError"
            ></el-avatar>
        </div>
    </div>
</template>

<script>
    export default {
        name: "index-body",
        data: function () {
            return {
                logo:{
                    backgroundImage:"url("+require("@/assets/logo.png")+")",
                    backgroundRepeat: "no-repeat",
                    backgroundSize: "100% 100%",
                    height: "60px",
                    width: "154px",
                    margin: "0 2em 0 2em"
                }
            }
        },
        methods: {
            avatarError(){
                //todo
                return true
            }
        }
    }
</script>

<style scoped>
    .painter-header {
        justify-content: space-between;
        display: flex;
        flex-wrap: wrap;
        /*learn: 渐变色，可以指定多个节点，可以追加百分比*/
        background: linear-gradient(to right,#545c64 20%,#fff);
        box-shadow : 0 2px 4px rgba(0,0,0,0.1)
    }

    .header-profile{
        width: 30%;
        display: flex;
    }

    .header-profile > *{
        margin: 10px 1em;
    }

</style>
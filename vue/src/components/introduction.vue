<template>
    <div id="introduction-container">
        <ul class="introduction-ul block">
            <li><span>作者:</span><strong>{{this.$store.state.authorName}}</strong></li>
            <li><span>邮箱:</span><strong>{{this.$store.state.mail}}</strong></li>
            <li><span>交友:</span><strong>{{this.$store.state.github}}</strong></li>
        </ul>
        <div class="motto block">{{this.$store.state.authorSay}}</div>
    </div>
</template>

<script>
    export default {
        name: "introduction"
    }
</script>

<style scoped>
    li span{
        color: #909090;
    }
    .motto{
        text-align: center;
        font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
        color: #1a222f;
    }

    .introduction-ul{
        font-size: 18px;
        line-height: 20px;
        list-style:none;
        margin:0;
        padding: 1em;
        overflow-x: hidden;
        font-family: Verdana,serif;
        white-space: nowrap;
    }

    .block{
        display: block;
        border-radius: 4px;
        border: 1px solid #f0a591;
    }
</style>
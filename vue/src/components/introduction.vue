<template>
    <div id="introduction-container">
        <ul class="introduction-ul block">
            <li><span>by&nbsp;&nbsp;:</span><strong>{{this.$store.state.authorName}}</strong></li>
            <li><span>mail:</span><strong>{{this.$store.state.mail}}</strong></li>
            <a style="display: block">http://github.com/{{this.$store.state.github}}</a>
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
        font-family: 'Merriweather', Georgia, serif;
        color: #1a222f;
    }

    .introduction-ul{
        font-size: 12px;
        line-height: 20px;
        list-style:none;
        padding: 8px;
        overflow-x: hidden;
        font-family: "Lucida Console", Monaco, monospace;
        white-space: nowrap;
    }

    .block{
        display: block;
        margin: 1em 0;
        border-radius: 4px;
        border: 1px solid #f0381a;
    }
</style>
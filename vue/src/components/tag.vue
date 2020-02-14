<template>
    <div class="painter-tags" :style="borderStyle"
         @mouseenter="_hover()" @mouseout="_hout()" v-html="content">
    </div>
</template>

<script>
    let colorList = ["Blue", "BlueViolet", "Red", "OrangeRed", "Tomato", "OliveDrab"];
    export default {
        name: "painter-tag",
        props: {
            selected: {
                type: Boolean,
                default: false,
            },
            inner: {
                type: String,
                default: ""
            }
        },
        data() {
            return {
                curveHover: false,
            }
        },
        computed: {
            content(){
                if(this.inner === ""){
                    return "empty tag"
                }
                if(this.inner.startsWith('http')){
                    var a = document.createElement('a');
                    a.href = this.inner;
                    let url = a.host + a.pathname.replace(/^([^/])/,'/$1');
                    if(!url){
                        return this.inner
                    }
                    a.innerText = url.startsWith("www.")?url.subStr(0,4):url
                    return a.outerHTML
                }
                return this.inner
            },
            borderStyle() {
                if (this.selected || this.curveHover) {
                    return {
                        backgroundColor: 'darkgray',
                        fontStyle: 'normal',
                        color: "#ffffff",
                    };
                }
                let r = Math.floor(Math.random() * colorList.length);
                return {
                    color: colorList[r],
                };
            },
        },
        methods: {
            _hover() {
                this.curveHover = true
            },
            _hout() {
                this.curveHover = false
            },
            select(v) {
                this.selected = Boolean(v)
            }
        }
    }
</script>

<style scoped>
    .painter-tags ,.painter-tags a{
        font-family: -apple-system, "Microsoft YaHei", 'Impact', 'Charcoal', sans-serif;
        display: inline-block;
        border: 1px solid;
        border-radius: 999em;
        padding: 0 10px;
        line-height: 24px;
        font-size: 12px;
        text-decoration: none;
        margin: 2px;
    }

    .painter-tags:hover{
        cursor: pointer;
    }
</style>
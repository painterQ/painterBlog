<template>
    <a class="painter-tags" :style="borderStyle" @click.stop="_click"
         @mouseenter="_hover()" @mouseout="_hout()" :href="href">
        {{innerText}}
    </a>
</template>

<script>
    let colorList = ["Blue", "BlueViolet", "Red", "OrangeRed", "Tomato", "OliveDrab"];
    export default {
        name: "painter-tag",
        props: {
            /*api*/
            selected: {
                type: Boolean,
                default: false,
            },
            /*api*/
            inner: {
                type: String,
                default: ""
            }
        },
        data() {
            return {
                curveHover: false,
                href: "",
                innerText:""
            }
        },
        watch: {
            "inner":{
                handler(newValue){
                    if(newValue === ""){
                        this.innerText = "empty tag";
                        return
                    }
                    if(newValue.startsWith('http')){
                        var a = document.createElement('a');
                        let url = a.host + a.pathname.replace(/^([^/])/,'/$1');
                        if(!url){
                            this.innerText = newValue
                            return;
                        }
                        a.href = newValue;
                        this.href = url;
                        this.innerText = url.startsWith("www.")?url.subStr(0,4):url;
                        return
                    }
                    this.innerText = newValue
                    this.href = "#/tags/"+ newValue;
                },
                immediate: true
            },
        },
        computed:{
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
            _click(){
                /*api*/
                this.$emit("clickTag", this.innerText, this.url)
            },
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
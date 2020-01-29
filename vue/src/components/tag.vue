<template>
    <div class="painter-tags" :style="{color: borderColor}"
         @click.stop="click($event)"
         @mouseenter="hover()" @mouseout="hout()">
        <slot></slot>
   </div>
</template>

<script>
    let colorList = ["Blue","BlueViolet","HotPink" ,
            "Lime","OrangeRed","Tomato","OliveDrab"];
    export default {
        name: "painter-tag",
        data(){
            return{
                borderColor: "",
                tmp: "",
            }
        },
        mounted() {
            let r = Math.floor(Math.random() * colorList.length);
            this.borderColor = colorList[r]
        },
        methods:{
            click(e){
                var eventdom = e.target
                console.log("click tag", eventdom.innerHTML)
            },
            hover(){
                this.tmp = this.borderColor
                this.borderColor = "#ffffff"
            },
            hout(){
                this.borderColor = this.tmp
            }
        }
    }
</script>

<style scoped>
    .painter-tags {
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
        background-color: darkgray;
        font-style: normal;
    }
</style>
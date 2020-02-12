<template>
    <div class="painter-tags" :style="borderStyle"
         @mouseenter="_hover()" @mouseout="_hout()">
        <slot></slot>
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
            }
        },
        data() {
            return {
                curveHover: false,
            }
        },
        computed: {
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
                console.log("change selected", v)
                this.selected = Boolean(v)
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
</style>
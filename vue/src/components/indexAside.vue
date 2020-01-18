<template>
    <div class="header_sticky" :id="randomId">
        <slot></slot>
    </div>
</template>

<script>
    export default {
        name: 'index-aside',
        computed: {
            randomId: function(){
                return 'randomId_' + Number(Math.random().toString().substr(3,3) + Date.now()).toString(36);
            },
        },

        watch: {
            oldToNew(newVal, oldVal) {
                if(newVal.length !== oldVal.length) {
                    this.$refs.sticky_.sticky_()
                }
            }
        },
        mounted() {
            if (!CSS.supports('position', 'sticky') && !CSS.supports('position', '-webkit-sticky')) {
                let self = this.$el;
                let origOffsetY = self.offsetTop;

                document.addEventListener('scroll', ()=>{
                    window.scrollY >= origOffsetY ? self.classList.add('sticky') :
                        self.classList.remove('sticky');
                });
            }
        }
    }
</script>

<style scoped>
    .header_sticky {
        width: 100%;
        position: sticky;
        top: 0;
        z-index: 100;
        transition: height 1s;
        padding: 2em 0 0 0;
    }

    .sticky {
        width: 100%;
        position: fixed;
        top: 0;
        z-index: 100;
    }
</style>
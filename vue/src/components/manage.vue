<template>
    <!--想要外层没有滚动条，所以包了一层container，设置固定大小，设置overflow-->
    <div id="const-size-container"  ref="const-size-container" @scroll="scrollHandle">
        <div id="manager" ref="manager">
            <!--        想要更新显示，1.修改依赖的状态 2.状态被Vue感知到-->
            <div class="col" v-for="(arts, colIndex) in colItems" :key="colIndex"
                 ref="col"
                 :style="{width: colWidth + 'px'}">
                <div class="art" v-for="(artItem, index) in arts" :key="index">
                    <div>
                        {{artItem.text}}
                    </div>
                </div>
            </div>
            <div class="bottom" ref="bottom">加载中...</div>
        </div>
    </div>
</template>

<script>
    import util from '@/api/axios.config.js'
    export default {
        name: "painter-manager",
        data: function () {
            return {
                colItems: [], //每一项是一个art数组，
                colWidth: 0, //Number 列宽 px
                arts: [],
            }
        },
        created: function(){
            util.post("/arts",{
                'start':0,
                'end':40,
            }).then( res=>{
                this.arts = res.data
            })
        },
        mounted: function () {
            /*
            * learn: dom 节点的克隆和追加
            * cloneNode得到的是Node对象
            *
            * 这样实现瀑布流的缺点是，每列宽度是固定的，有可能导致很宽的横向间隔
            * */
            let w = this.$refs["manager"].clientWidth;
            let colNum = Math.floor(w / 250);
            this.colWidth = 250 + Math.floor((w - colNum * 262) / colNum);
            this.colItems = new Array(colNum);

            //现在的问题是，在mount中修改了元素，但是获取不到受此影响的其他元素的尺寸
            //生产art

            this.add();
        },
        watch:{
          'arts': function () {
              this.add()
          }
        },
        methods : {
            scrollHandle: function(){
                let tmp = this.$refs['const-size-container']
                //判定滚动条到底了
                //clientHeight  scrollTop  scrollHeight
                if (tmp.clientHeight + tmp.scrollTop === tmp.scrollHeight ){
                    util.post("/arts",{
                        'start': this.arts.length,
                        'end': this.arts.length + 20,
                    }).then( res=>{
                        this.arts = this.arts.concat(res.data)
                    }).catch(()=>{
                        this.$refs['bottom'].innerHTML = "到底了"
                    })
                }
            },
            min : function() {  //不能使用箭头函数，因为箭头函数没有this，vue就不能把this绑定到当前实例
                let min = 65535, r = 0;
                //当ref和v-for搭配的时候，$ref[~]会返回数组
                //不要吧ref绑定变量
                let cols = this.$refs['col'];
                for (let i in cols) {
                    if (cols[i].offsetHeight < min) {
                        min = cols[i].offsetHeight;
                        r = i;
                    }
                }
                //计算属性中不允许产生修改其他状态的副作用
                return r;
            },
            /*
            * 异步函数，async和await关键字
            * async是作为关键字放在function前，定义一个异步函数，异步函数其实就是返回promise的函数
            * 返回值会被当做resolve的参数，throw的error会被当作reject的参数
            *
            * await后面应该跟一个promise对象的实例，然后逻辑会停在这里，直到promise对象有了resolve或者reject状态之一
            * 下面两个的效果相同
            */
/*
function log(i) {
    setTimeout(
        () => {
            console.log(i)
        }, 100
    )
}

Promise.resolve(() => {}).
then(
    () => { for (let i = 0; i < 2; i++) log(1); }).
then(
    () => { for (let i = 0; i < 2; i++) log(2); }).
then(
    () => { for (let i = 0; i < 2; i++) log(3); });
log(4);

*/

            /* log(4);
            *
            * async function(){
            *   await Promise.resolve(()=>{ for(let i = 0;i< 10;i++)log(1);});
            *   await Promise.resolve(()=>{ for(let i = 0;i< 10;i++)log(2);});
            *   await Promise.resolve(()=>{ for(let i = 0;i< 10;i++)log(3);});
            * }()
            * log(4);
            * */
            add :async function () {
                for (let i in this.arts) {
                    let currentCol = this.min();
                    if (!(this.colItems[currentCol] instanceof Array)) {
                        this.colItems[currentCol] = [];
                    }
                    //learn: Array.from()
                    this.colItems = Array.from(this.colItems);
                    this.colItems[currentCol].push({text:this.arts[i]});
                    console.log("添加元素，当前总长",this.arts.length, i)
                    await this.$nextTick();
                }
            }
        }
    }
</script>

<style scoped>
    #const-size-container{
        height: 100%;
        width: 100%;
        overflow: scroll;
    }
    .head{
        max-width: 1024px;
        margin: 3em auto;
    }
    #manager {
        max-width: 1024px;
        margin: 3em auto;
        position: relative;
        overflow-y: auto;
        overflow-x: visible;
    }
    .bottom{
        float: left;
        width: 100%;
        text-align: center;
        height: 3em;
        padding: 1em;
        box-sizing: border-box;
    }
    .col {
        float: left;
        box-sizing: border-box;
        margin-left: 12px;
    }

    .art {
        border-radius: 4px;
        width: 100%;
        overflow-y: hidden;
        padding: 1em;
        margin-top: 1em;
        background-color: #fff;
        box-shadow: 0 0 4px rgba(0, 0, 0, 0.1);
        box-sizing: border-box;
    }

    .art:hover {
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.3);
        position: relative;
        top: 1px;
    }

    .art > * {
        display: block;
        cursor: pointer;
    }
</style>
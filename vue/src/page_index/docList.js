import Doc from "./doc";
class DocListClass {
    docSet = {}; //文章元信息
    cache = {};  //文章内容的缓存

    ErrNeedGetMateList = 0;
    ErrNeedGetDoc = 1;
    constructor(set = []) {
        for (let index in set){
            this.docSet[set[index].id] = new Doc(set[index])
        }
    }

    //设置缓存
    set(k, doc) {
        if (typeof doc === "string") {
            this.cache[k] = doc
        }
    }
    //获取文章元信息
    getDocMate(key){
        return this.docSet[key]
    }

    //获取文章内容，使用了缓存
    get(k) {
        if (this.cache[k] && this.cache[k].length > 0) {
            console.log("get doc cache")
            return this.cache[k];
        }

        if (!this.getDocMate(k)){
            console.log("not in list too, 404")
            throw this.ErrNeedGetMateList
        }
        console.log("in doclist ,get doc by api")
        throw this.ErrNeedGetDoc
        /*
        * await 表达式会暂停当前 async function 的执行，等待 Promise 处理完成。若 Promise 正常处理(fulfilled)，
        * 其回调的resolve函数参数作为 await 表达式的值，继续执行 async function。
        * 若 Promise 处理异常(rejected)，await 表达式会把 Promise 的异常原因抛出.
        * 另外，如果 await 操作符后的表达式的值不是一个 Promise，
        * 那么该值将被转换为一个已正常处理的 Promise。*/
    }

    //更新元信息
    //初始化时、查找文章内容发现没有改文章时、点击目录时


    next(key){
        if (typeof key !== "string" || key===""){
            return key
        }
        let flag = 0;
        for (let e of this){
            if (flag === 1){
                return e.id
            }
            if(e.id === key) {
                flag = 1
            }
        }
        return key
    }

    prev(key){
        if (typeof key !== "string" || key===""){
            return key
        }
        let arr = [];
        let i = 0;
        let flag = 0;
        for (let e of this){
            arr.push(e)
            if (e.id === key){
                flag =1
                break
            }
            i++
        }
        if(flag === 1){
            if (i === 0){
                return arr[0].id
            }
            return arr[i - 1].id
        }
        return key
    }

    [Symbol.iterator]() {
        let myDocs = this.docSet;
        let propUp = [];
        let prop = [];
        for (let key in myDocs) {
            myDocs[key].attr === "top" ? propUp.push(key) : prop.push(key)
        }
        let sort = (a, b) => {
            if (a === b) return 0;
            return a > b ? 1 : -1
        };
        prop.sort(sort);
        propUp.sort(sort);
        let items = propUp.concat(prop);
        let i = 0;
        return {
            next: function () {
                return {
                    done: i >= items.length,
                    value: myDocs[items[i++]]
                };
            }
        };

    }
}

export default DocListClass
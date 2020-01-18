import api from '../api/rpc'
import store from "./store";
import Doc from "./doc";
class DocListClass {
    docSet = {};
    cache = {};

    constructor(set = []) {
        for (let index in set){
            this.docSet[set[index].id] = new Doc(set[index])
        }
        this.updateList('/docc',10)
    }

    set(k, doc) {
        if (typeof doc === "string") {
            console.log("set", k, doc)
            this.cache[k] = doc
        }
    }
    findDocByKey(key){
        for(let k in this.docSet){
            if(k === key && this.docSet[k]){
                return this.docSet[k]
            }
        }
        return null
    }

    async get(k) {
        console.log("###get doc",k)
        if (this.cache[k] && this.cache[k].length > 0) {
            console.log("get doc cache")
            return this.cache[k];
        }
        if (this.findDocByKey(k) === null){
            console.log("not in list, updateList",k,10)
            await this.updateList(k,10)
        }
        if (this.findDocByKey(k) === null){
            console.log("not in list too, 404")
            return null
        }
        console.log("in doclist ,get doc by api")
        /*
        * await 表达式会暂停当前 async function 的执行，等待 Promise 处理完成。若 Promise 正常处理(fulfilled)，
        * 其回调的resolve函数参数作为 await 表达式的值，继续执行 async function。
        * 若 Promise 处理异常(rejected)，await 表达式会把 Promise 的异常原因抛出.
        * 另外，如果 await 操作符后的表达式的值不是一个 Promise，
        * 那么该值将被转换为一个已正常处理的 Promise。*/
        try{
            let ret = await api.getDoc({doc: k});
            this.set(k,ret.data)
            /*
            * async函数的返回值，如果是普通值，会被包装成promise
            * 要么用then处理，要么用await处理
            * */
            return ret.data
        }catch (e) {
            console.log(e)
        }
    }

    async updateList(start='',length=10) {
        if (start === '') return;
        api.getDocsList({start:start,length: length}).then(
            data =>{
                let set = data.data.list
                for (let index in set){
                    this.docSet[set[index].id] = new Doc(set[index])
                }
                console.log(data.data)
                store.commit("setTotalDocs",data.data.total)
                store.commit('setDocListUpdateState',set.length > 0);
            }
        ).catch(err=>{console.log("getDocsList,err:",err)});
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
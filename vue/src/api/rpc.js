import util from './axios.config'

let nil = {$router: null, $store: null, $cookies: null}
function login(user) {
    let vue = this === undefined || this === window ||!this._isVue? nil:this;
    return util.post("/login", user, vue)
}

//uploadImage 上传图片
//method: Post
//path /docs/image/filter
//data: { img: img, type: blobInfo.blob.type, name:"" }
//return: {'url':"http://localhost:8080/public/img/background.0ed615ed.jpg"}
// @router /image/filter [post]
function uploadImage(param, config) {
    let vue = this === undefined || this === window ||!this._isVue? nil:this;
    return util.post('/docs/image/filter', param, vue, config);
}

//修改基础信息
// {mail: "", github: "",}
function changeBaseInfo(info) {
    let vue = this === undefined || this === window ||!this._isVue? nil:this;
    console.log("changeBaseInfo",this,this === undefined,this === window ,!this.data)
    return util.post("/login/base/filter", info, vue)
}

// { name: "", title: "", subTitle: "", IPC: ""}
function changeBlogInfo(info) {
    let vue = this === undefined || this === window ||!this._isVue? nil:this;
    return util.post("/login/blog/filter", info, vue)
}

// {pwd: ""}
function changePwdChange(info) {
    let vue = this === undefined || this === window ||!this._isVue? nil:this;
    return util.post("/login/pwd/filter", info, vue)
}

//获取文章元信息
//{"id":"first","title":"first","subTitle":"blog",
// "tags":["blog","document"],"attr":0,"time":"",
// "abstract":"PHA+Zmlyc3Q8L3A+"},
function getDocsList(info) {
    console.log("api, 获取文章元信息");
    let vue = this === undefined || this === window ||!this._isVue? nil:this;
    return util.post("/docs", info, vue)
}

/*
* info: {doc: '/doc0'}
* data: string
* */
function getDoc(info) {
    console.log("api, 获取文章内容");
    let vue = this === undefined || this === window ||!this._isVue? nil:this;
    return util.get("/docs", info, vue)
}

//发表文章
function postDoc(info) {
    console.log("api, 创建新的文章");
    let vue = this === undefined || this === window ||!this._isVue? nil:this;
    //info: "{"title":"first","path":"first","document":"<p>第一篇文章</p>\n<p>&nbsp;</p>"}"}
    return util.post("/docs/doc/filter", info, vue)
}

//GetTags 获取全部tag
//method: GET
//path /docs/tag
//para: nil
//return: {"tag1":["",""],"tag2":["",""],"tag3":[]}
// @router /tag [get]
function getTags() {
    console.log("api, 获取全部tag");
    let vue = this === undefined || this === window ||!this._isVue? nil:this;
    return util.get("/docs/tag", undefined, vue)
}


//AddTag 新增tag
//method: Post
//path /docs/tag/filter
//data: ["tag1","tag2","tag3"]
//return: nil
// @router /tag/filter [post]
function addTag(info) {
    console.log("api, 新增tag")
    let vue = this === undefined || this === window ||!this._isVue? nil:this;
    return util.post("/docs/tag/filter", info, vue)
}

//GetAuthorInfo 获取作者信息（author）和博客信息（header）
//method: get
//path /login
//data: nil
//{title: "Painter Qiao",
//          subTitle: "for dear & love",
//          avatar: "./avatar.jpeg",
//          lastLogin: 123213213,
//          name: "Painter Qiao",
//			ipc: "",
//			github: "",
//          say: "a blog for dear & love"
//          email: ""}
// @router / [get]
function getAuthorInfo() {
    console.log("api, 获取作者信息");
    let vue = this === undefined || this === window ||!this._isVue? nil:this;
    return util.get("/login", undefined, vue)
}

//GetImageList 获取图片列表
//method: Get
//path /docs/image/filter
//para: start、limit
//return: [{"id":"","name":"","type":"","src":"",}]
// @router /image/filter [get]
function getImageList({start,limit}) {
    console.log("api, 获取图片列表");
    let vue = this === undefined || this === window ||!this._isVue? nil:this;
    return util.get('/docs/image/filter',{start:start,limit:limit},vue)
}

//DeleteDoc 删除文章
//method: DELETE
//path /docs/doc/filter
//data {"id":"/doc1"}
//return 200
// @router /doc/filter [delete]
function deleteDoc(id) {
    console.log("api, 删除文章");
    let vue = this === undefined || this === window ||!this._isVue? nil:this;
    return util.delete('/docs/doc/filter',{id:id},vue)
}



//导出 default的含义
export default {
    install(vue) {
        //传入的是vue的构造函数
        vue.prototype.$_login = login;
        vue.prototype.$_changeBaseInfo = changeBaseInfo;
        vue.prototype.$_changeBlogInfo = changeBlogInfo;
        vue.prototype.$_changePwdChange = changePwdChange;
        vue.prototype.$_getDocsList = getDocsList;
        vue.prototype.$_getDoc = getDoc;
        vue.prototype.$_uploadImage = uploadImage;
        vue.prototype.$_postDoc = postDoc;
        vue.prototype.$_getTags = getTags;
        vue.prototype.$_addTag = addTag;
        vue.prototype.$_getAuthorInfo = getAuthorInfo;
        vue.prototype.$_getImageList = getImageList
        vue.prototype.$_deleteDoc = deleteDoc
    },
    login,
    changeBaseInfo,
    changeBlogInfo,
    changePwdChange,
    getDocsList,
    getDoc,
    uploadImage,
    postDoc,
    getTags, addTag,
    getAuthorInfo,
    getImageList,
    deleteDoc
}
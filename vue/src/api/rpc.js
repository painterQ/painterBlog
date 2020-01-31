import util from './axios.config'

function login(user) {
    return util.post("/login",user)
}

function uploadImage(data) {
   return util.post('/doc/image',data)
}

//修改基础信息
// {mail: "", github: "",}
function changeBaseInfo(info, vue) {
    return util.post("/login/base/filter", info)
}

// { name: "", title: "", subTitle: "", IPC: ""}
function changeBlogInfo(info) {
    return util.post("/login/blog/filter", info)
}

// {pwd: ""}
function changePwdChange(info) {
    return util.post("/login/pwd/filter", info)
}

//获取文章元信息
function getDocsList(info) {
    console.log("api, 获取文章元信息")
    return util.post("/docs", info)
}
/*
* info: {id: '/doc0'}
* data: {content: '文章'}  or {error: '...'}
* */
function getDoc(info) {
    console.log("api, 获取文章内容")
    return util.get("/docs", info)
}

//发表文章
function postDoc(info){
    console.log("api, 创建新的文章")
    //info: "{"title":"first","path":"first","document":"<p>第一篇文章</p>\n<p>&nbsp;</p>"}"}
    return util.post("/docs/doc/filter", info)
}

//GetTags 获取全部tag
//method: GET
//path /docs/tag
//para: nil
//return: ["tag1","tag2","tag3"]
// @router /tag [get]
function getTags(){
    console.log("api, 获取全部tag")
    return util.get("/docs/tag")
}


//AddTag 新增tag
//method: Post
//path /docs/tag/filter
//data: ["tag1","tag2","tag3"]
//return: nil
// @router /tag/filter [post]
function addTag(info){
    console.log("api, 新增tag")
    return util.post("/docs/tag/filter", info)
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
function getAuthorInfo(){
    console.log("api, 获取作者信息");
    return util.get("/login",)
}


//导出 default的含义
export default {
    login,
    changeBaseInfo,
    changeBlogInfo,
    changePwdChange,
    getDocsList,
    getDoc,
    uploadImage,
    postDoc,
    getTags,addTag,
    getAuthorInfo}
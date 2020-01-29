import util from './axios.config'

function login(user) {
    return util.post("/login",user)
}

function uploadImage(data) {
   return util.post('/doc/image',data)
}

function changeBaseInfo(info) {
    return util.post("/info/base", info)
}

function changeBlogInfo(info) {
    return util.post("/info/blog", info)
}

function changePwdChange(info) {
    return util.post("/info/pwd", info)
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
    return util.post("/docs/doc", info)
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
//path /docs/tag
//data: ["tag1","tag2","tag3"]
//return: nil
// @router /tag [post]
function addTag(info){
    console.log("api, 新增tag")
    return util.post("/docs/tag", info)
}

//GetAuthorInfo 获取作者信息（author）和博客信息（header）
//method: get
//path /login
//data: nil
//return: {title: "Painter Qiao",
//          subTitle: "for dear & love",
//          avatar: "./avatar.jpeg",
//          lastLogin: 123213213,
//          name: "Painter Qiao",
//          say: "a blog for dear & love"}
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
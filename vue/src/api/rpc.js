import util from './axios.config'

function login(user) {
    console.log("login")
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

//导出 default的含义
export default {
    login,
    changeBaseInfo,
    changeBlogInfo,
    changePwdChange,
    getDocsList,
    getDoc,
    uploadImage,
    postDoc}
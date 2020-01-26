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

function getDocsList(info) {
    return util.post("/docs", info)
}
/*
* info: {id: '/doc0'}
* data: {content: '文章'}  or {error: '...'}
* */
function getDoc(info) {
    return util.get("/doc", info)
}

//发表文章
function postDoc(info){
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
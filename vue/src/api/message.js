import {Message} from 'element-ui'

//type 可以是 success message warning error
let message = function (obj, msg = '', type = '') {
    console.log("message")
    type = type || "error";
    if (!obj || !obj._isVue) {
        console.log("generate message error: [type]:" + type + "[message]:" + msg);
        return;
    }
    const h = obj.$createElement;
    let fun = type === "error" ? obj.$message.error : obj.$message;
    fun({
        type: type,
        message: h('p', null, [
            h('span', null, '内容可以是 '),
            h('i', {style: 'color: teal'}, msg)
        ])
    });
};

export default {
    install(vue) {
        vue.prototype.$message = Message;
    },
    //message(vue,"登陆弹框",'warning');
    message
}
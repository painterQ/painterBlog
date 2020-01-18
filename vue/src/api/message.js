import Vue from 'vue'
import {Message,MessageBox} from 'element-ui'

Vue.prototype.$message = Message;
Vue.prototype.$messageBox = MessageBox;

//type 可以是 success message warning error
let message = function(obj, msg='',type='') {
    if(!obj) return;
    const h = obj.$createElement;
    let fun = type==="error"?obj.$message.error:obj.$message;
    fun({
        type: type,
        message: h('p', null, [
            h('span', null, '内容可以是 '),
            h('i', { style: 'color: teal' }, msg)
        ])
    });
};

/*
* 使用方法
* message(this,"登陆弹框",'warning');
* */

export default message
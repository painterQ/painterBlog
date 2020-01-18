import axios from "axios"
import {Message} from "element-ui"

//  这个baseUrl要根据实际情况进行改变
axios.defaults.baseURL = "/";
axios.defaults.headers.common["Content-Type"] =
    "application/json; charset=UTF-8";
axios.defaults.headers.common["Access-Control-Allow-Origin"] = "*";

// 请求拦截器 添加token
axios.interceptors.request.use(
    config => {
        return config
    },
    error => {
        return Promise.reject(error)
    }
)

// 响应拦截器即异常处理
axios.interceptors.response.use(
    response => {
        return response
    },
    error => {
        Message({
            message: error.message,
            type: "error",
            duration: 5000,
        })
        return Promise.resolve(error)
    }
)

export default {
    // get请求
    get(url, param) {
        return new Promise((resolve, reject) => {
            axios({
                method: "get",
                url,
                params: param,
            })
                .then(res => {
                    resolve(res)
                })
                .catch(error => {
                    Message({
                        message: error,
                        type: "error",
                        duration: 5000,
                    });
                    reject(error)
                })
        })
    },
    // post请求
    post(url, param) {
        return new Promise((resolve, reject) => {
            axios({
                method: "post",
                url,
                data: param,
            })
                .then(res => {
                    resolve(res)
                })
                .catch(error => {
                    Message({
                        message: error,
                        type: "error",
                        duration: 5000,
                    })
                    reject(error)
                })
        })
    },
    // all get
    allGet(fnArr) {
        return axios.all(fnArr)
    },
}

/*
* promise对象
*  该对象有三个状态，等待执行完pending，成功resolved，失败rejected；只能使用pending转换到另外两个
*  状态意味着，只要在这个状态，再添加回调函数也会执行的，不是错过了就没了。（也就是说，可以在一个已经resolve的promise上继续挂then）
*  构造器：接受一个函数function(resolve, reject){}作为参数，两个参数在成功和失败的时候调用
*       关于这两个函数，resolve接受一个参数，这个参数会传给后面的then中的回调; resolve也接受第二个参数
*                     reject接受一个error(), 这个参数会传给后面的catch中的回调
*       快捷构造器 Promise.resolve() Promise.reject()
*       特殊构造器 Promise.all() Promise.race()
*  then方法，返回一个promise对象，因此可以链式调用
*  catch方法，返回一个promise对象，可以链式调用，当链中的某个then失败了，会调用离得最近的cache
*
*  then和catch一定是异步调用的，也就是最快也是下一个周期调用的
*  初始化promise时传入的函数却会同步立即执行（也就是写了一个死循环的话就永远不会执行下面的代码了，那么页面UI都会停住没有响应）
*
*  每次then都是返回全新的promise对象
*
*  一般来说，只有给new promise()传入一个异步操作才有意义，这样一开始的promise采用pending这个状态，不然直接就得到一个resolve或者reject状态的promise
*  虽然这也有写用，但是基本上主要用处还是传入一个异步操作。
*/
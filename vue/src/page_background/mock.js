import Mock from 'mockjs'

Mock.mock('/login', 'post', (options) => {
    console.log('options:', options)
    let data = JSON.parse(options.body)
    let name = data.name
    let password = data.password
    if (name === 'admin' && password === 'admin') {
        return {
            status: 1,
            message: '登录成功'
        }
    } else {
        return {
            status: 0,
            message: '账号或者密码错误'
        }
    }
});

Mock.mock('/doc/image', 'post', (options) => {
    console.log('options:', options)
    console.log("上传图片")
    return {'url':"http://localhost:8080/public/img/background.0ed615ed.jpg"}
});

let r = [];
for(let i = 0;i<160;i++){
    r.push(i + "曦曦 爱你".repeat(Math.ceil(Math.random() * 20)))
}

Mock.mock('/arts', 'post', (options) => {
    console.log('options:', options);
    let data = JSON.parse(options.body);
    let start = data.start;
    let end = data.end;
    if(end < start) return null;
    if(end > 160) return null;
    return r.slice(start,end);
});

Mock.mock('/info/base', 'post', (options) => {
    console.log('options:', options);
    return {ok:true};
});

Mock.mock('/info/blog', 'post', (options) => {
    console.log('options:', options);
    return {ok:true};
});

Mock.mock('/info/pwd', 'post', (options) => {
    console.log('options:', options);
    return {ok:true};
});

Mock.mock('/docs/doc', 'post', (options) => {
    console.log('options:', options);
    return {ok:true};
});


let tags = ["tag1","tag2","tag3","tag4","tag5","tag6"];

//GetTags 获取全部tag
//method: GET
//path /docs/tag
//para: nil
//return: ["tag1","tag2","tag3"]
// @router /tag [get]
Mock.mock("/docs/tag", 'get', () => {
    console.log("api mock, 获取全部tag")
    return tags
});

//AddTag 新增tag
//method: Post
//path /docs/tag
//data: ["tag1","tag2","tag3"]
//return: nil
// @router /tag [post]
Mock.mock("/docs/tag", 'post', (req)=>{
    console.log("api mock, 新增tag",req)
    let r = JSON.parse(req.body);
    if(r instanceof Array && r.length >0){
        tags.concat(r)
        tags = [...new Set(tags)]
    }
});
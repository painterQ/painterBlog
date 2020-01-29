import Vue from 'vue'
import Vuex from 'vuex'
import DocListClass from './docList.js'

Vue.use(Vuex);

const init = [
    // {
    //     id: '/doc',
    //     title: "第一篇来自editor的文章",
    //     subTitle: "随笔",
    //     tags: ['原创', 'editor'],
    //     attr: 'top',
    //     time: Date.now(),
    //     pref: '/doc',
    //     next: '/doca',
    //     abstract: '第一篇文章，即将开始美好生活',
    // }, {
    //     id: '/doca',
    //     title: "第二篇文章",
    //     subTitle: "随笔",
    //     tags: ['原创', 'editor'],
    //     time: Date.now(),
    //     pref: '/doc',
    //     next: '/docb',
    //     abstract: '第一篇文章，即将开始美好生活',
    // }, {
    //     id: '/docb',
    //     title: "第三篇文章",
    //     subTitle: "随笔",
    //     tags: ['原创', 'editor'],
    //     attr: 'top',
    //     time: Date.now(),
    //     pref: '/doca',
    //     next: '/docc',
    //     abstract: '第三篇文章，即将开始美好生活',
    // }, {
    //     id: '/docc',
    //     title: "第四篇文章",
    //     subTitle: "随笔",
    //     tags: ['原创', 'editor'],
    //     time: Date.now(),
    //     pref: '/docb',
    //     next: '/docd',
    //     abstract: '第四篇文章，即将开始美好生活',
    // }
];

const store = new Vuex.Store({
    state: {
        /*header，页面切换会变动*/
        headerTitle: "",
        headerSubTitle: "",
        headerTags: [],
        headerTime: Number(new Date().getDate()),
        headerName: "",

        /*作者相关,初始化后不再变动*/
        authorAvatar: "",
        blogTitle: "Blog",
        blogSubTitle: "subTitle",
        authorLastLogin: Number(new Date().getDate()),
        authorName: "Someone",
        authorSay: "for dear & love",

        docs: new DocListClass(init),
        docsUpdate: false,
        total: 0
    },

    mutations: {
        setDocListUpdateState: (state, b) => {
            console.log("setDocListUpdateState", b);
            state.docsUpdate = b
        },
        setTotalDocs: (state, num) => {
            state.total = num
        },
        setHeader: (state, header) => {
            state.headerTitle = header.title || state.headerTitle;
            state.headerSubTitle = header.subTitle || state.headerSubTitle;
            state.headerTags = JSON.parse(JSON.stringify(header.tags)) || [];
            state.headerTime = header.time || state.headerTime;
            state.headerName = header.name || state.headerName;
        },
        setAuthor: (state, author) => {
//          {title: "Painter Qiao",
//          subTitle: "for dear & love",
//          avatar: "./avatar.jpeg",
//          lastLogin: 123123132,
//          name: "Painter Qiao",
//          say: "a blog for dear & love"}
            state.authorAvatar = author.avatar || state.authorAvatar;
            //api请求来的time是s为单位的，js中需要ms为单位
            state.authorLastLogin = Number(author.lastLogin) * 1000;
            state.authorName = author.name || state.authorName;
            state.authorSay = author.say || state.authorSay ;
            state.blogTitle = author.title || state.blogTitle;
            state.blogSubTitle = author.subTitle || state.blogSubTitle;
        }
    }
});

export default store


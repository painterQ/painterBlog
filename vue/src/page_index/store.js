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
        author: {
            name: "Painter Qiao",
            avatar: "./",
            lastLogin: Date.now(),
            notice: "通知：这是一条重要通知哦",
            aboutMe: "自我介绍"
        },
        docs: new DocListClass(init),
        docsUpdate: false,
        total: 0
    },

    mutations: {
        setDocListUpdateState: (state, b) => {
            console.log("setDocListUpdateState", b)
            state.docsUpdate = b
        },
        setTotalDocs: (state, num)=>{
            state.total = num
        }
    }
});

export default store


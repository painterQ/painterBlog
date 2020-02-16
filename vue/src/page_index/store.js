import Vue from 'vue'
import Vuex from 'vuex'
import DocListClass from './docList.js'
import api from '../api/rpc'
import Doc from "./doc";

Vue.use(Vuex);
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
            ipc: "备案",
            github: "https://www.github.com/painterQ",
            mail: "example@mail.com",

            /*文章相关*/
            docs: new DocListClass(),
            currentPath: "",
            _docNeedRefresh: false,
            _initFinish: false,
            initPromise: null,
            total: 0
        },
        getters: {
            docMateList: state => {
                if (!state._initFinish) return [];
                let output = [];
                for (let e of state.docs) {
                    output.push(e)
                }
                console.log("store docMateList",output)
                return output
            },
            getDocFromStore: state => {
                if (!state._initFinish) return "";
                if (!/^\/docs\/.*/.test(state.currentPath)) {
                    return ""
                }
                //_docNeedRefresh既是依赖也是变动项，但是这不会导致再调用一次
                let docID = state.currentPath.substr(5);
                try {
                    return state.docs.get(docID)
                } catch (e) {
                    if (e === state.docs.ErrNeedGetDoc) {
                        return new Promise(async (resolve)=>{
                            let res = await api.getDoc({doc: docID})
                            state.docs.set(docID, res.data);
                            state._docNeedRefresh = true //getDoc需要被再调用一次
                            resolve(res.data)
                        });
                    }
                    return "/404" //e === state.docs.ErrNeedGetMateList
                }
            },
            prevDoc: state => {
                if (!state.currentPath) {
                    return ""
                }
                return state.docs.prev(state.currentPath.substr(5));
            },
            nextDoc: state => {
                if (!state.currentPath) {
                    return ""
                }
                return state.docs.next(state.currentPath.substr(4));
            }
        },
        mutations: {
            setHeader: (state, header) => {
                state.headerTitle = header.title || state.headerTitle;
                state.headerSubTitle = header.subTitle || state.headerSubTitle;
                state.headerTags = JSON.parse(JSON.stringify(header.tags)) || [];
                state.headerTime = header.time || state.headerTime;
                state.headerName = header.name || state.headerName;
            },
            setAuthor: (state, author) => {
                state.authorAvatar = author.avatar || state.authorAvatar;
                //api请求来的time是s为单位的，js中需要ms为单位
                state.authorLastLogin = Number(author.lastLogin) * 1000;
                state.authorName = author.name || state.authorName;
                state.authorSay = author.say || state.authorSay;
                state.blogTitle = author.title || state.blogTitle;
                state.blogSubTitle = author.subTitle || state.blogSubTitle;
                state.ipc = author.ipc || state.ipc;
                state.github = author.github || state.github;
                state.mail = author.mail || state.mail;
            },
        },
        actions: {
            setCurrentPath({state, commit}, path) {
                state.currentPath = path;
                if (state.currentPath.startsWith("/docs")) {
                    let currentDoc = state.docs.docSet[state.currentPath.substr(5)]
                    if (!currentDoc) return; //docSet没有初始化完成
                    commit("setHeader", {
                        title: currentDoc.title,
                        subTitle: currentDoc.subTitle,
                        time: currentDoc.time,
                        tags: JSON.parse(JSON.stringify(currentDoc.tags)),
                        name: currentDoc.name,
                    })
                } else if (state.currentPath.startsWith("/list")) {
                    commit("setHeader", {
                        title: state.blogTitle,
                        subTitle: state.blogSubTitle,
                        time: state.authorLastLogin,
                        tags: ["博客"],
                        name: state.authorName,
                    })
                }else if (state.currentPath.startsWith("/tags")) {
                    commit("setHeader", {
                        title: "按标签分类",
                        subTitle: state.currentPath.substr(5),
                        time: state.authorLastLogin,
                        tags: [],
                        name: state.authorName,
                    })
                }
            }
        }
    });


store.state.initPromise = new Promise(async (resolve)=>{
        console.log("init store");
        let data1 = api.getAuthorInfo();
        let data2 = api.getDocsList({start: "/doca", length: 10});
        let {data: authorInfo} = await data1;
        //title,subTitle, avatat, lastLogin, name, ipc,github,say,email
        store.state.authorAvatar = authorInfo.avatar
        //api请求来的time是s为单位的，js中需要ms为单位
        store.state.authorLastLogin = Number(authorInfo.lastLogin) * 1000;
        store.state.authorName = authorInfo.name;
        store.state.authorSay = authorInfo.say;
        store.state.blogTitle = authorInfo.title;
        store.state.blogSubTitle = authorInfo.subTitle;
        store.state.ipc = authorInfo.ipc;
        store.state.github = authorInfo.github;
        store.state.mail = authorInfo.mail;
        //title, subTitle, name, time, tags
        store.state.headerTitle = authorInfo.title;
        store.state.headerSubTitle = authorInfo.subTitle;
        store.state.headerTags = ["博客"];
        store.state.headerTime = authorInfo.lastLogin;
        store.state.headerName = authorInfo.name;
        let {data: {list: set}} = await data2;
        for (let i in set) {
            store.state.docs.docSet[set[i].id] = new Doc(set[i]);
        }
        store.state.total = Number(set.length);
        store.state._initFinish = true
        console.log("init finish");
        resolve(store.state.docs)
})

export default store


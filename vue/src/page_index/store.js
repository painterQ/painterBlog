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
            initPromise: null,
            total: 0,
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
                if (path.startsWith("/docs")) {
                    let currentDoc = state.docs.docSet[path.substr(5)]
                    if (!currentDoc) return; //docSet没有初始化完成
                    commit("setHeader", {
                        title: currentDoc.title,
                        subTitle: currentDoc.subTitle,
                        time: currentDoc.time,
                        tags: JSON.parse(JSON.stringify(currentDoc.tags)),
                        name: currentDoc.name,
                    })
                } else if (path.startsWith("/list")) {
                    commit("setHeader", {
                        title: state.blogTitle,
                        subTitle: state.blogSubTitle,
                        time: state.authorLastLogin,
                        tags: ["博客"],
                        name: state.authorName,
                    })
                }else if (path.startsWith("/tags")) {
                    commit("setHeader", {
                        title: "按标签分类",
                        subTitle: path.substr(5),
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
        store.state.mail = authorInfo.email;
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
        console.log("init finish");

        store.state.docs.docMateList= () => {
            let output = [];
            for (let e of store.state.docs) {
                output.push(e)
            }
            console.log("store docMateList",output)
            return output
        };

        store.state.docs.getDocFromStore= (currentPath) => { //返回文章内容或者一个promise
            if(! currentPath.startsWith("/docs")) return;
            let docID = currentPath.substr(5);
            try {
                return store.state.docs.get(docID)
            } catch (e) {
                if (e === store.state.docs.ErrNeedGetDoc) {
                    return new Promise(async (resolve)=>{
                        let res = await api.getDoc({doc: docID})
                        store.state.docs.set(docID, res.data);
                        resolve(res.data)
                    });
                }
                return "/404" //e === state.docs.ErrNeedGetMateList
            }
        };

        store.state.docs.prevDoc = (currentPath) => {
            return store.state.docs.prev(currentPath.substr(5));
        },
        store.state.docs.nextDoc = (currentPath) => {
            return store.state.docs.next(currentPath.substr(5));
        }
        resolve(store.state.docs)
})

export default store


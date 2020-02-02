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
            total: 0
        },
        getters: {
            docMateList: state => {
                if (!state._initFinish) return [];
                let output = [];
                for (let e of state.docs) {
                    output.push(e)
                }
                return output
            },
            getDoc: state => {
                if (!state._initFinish) return "";
                if (!/^\/doc\/.*/.test(state.currentPath)) {
                    return ""
                }
                //_docNeedRefresh既是依赖也是变动项，但是这不会导致再调用一次
                if(state._docNeedRefresh) state._docNeedRefresh = false;
                let docID = state.currentPath.substr(4);
                try {
                    return state.docs.get(docID)
                } catch (e) {
                    if (e === state.docs.ErrNeedGetDoc) {
                        (async ()=>{
                            let res = await api.getDoc({doc: docID})
                            state.docs.set(docID, res.data);
                            state._docNeedRefresh = true //getDoc需要被再调用一次
                        })();
                        return "" //getting document
                    }
                    return "/404" //e === state.docs.ErrNeedGetMateList
                }
            },
            prevDoc: state => {
                if (!state.currentPath) {
                    return ""
                }
                return state.docs.prev(state.currentPath.substr(4));
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
            async InitAsync({state, commit, dispatch}) {
                console.log("this",this)
                console.log("api getDocsList...", 0 + new Date());
                let data1 = api.getAuthorInfo();
                let data2 = api.getDocsList({start: "/doca", length: 10});
                let {data: authorInfo} = await data1;
                //title,subTitle, avatat, lastLogin, name, ipc,github,say,email
                commit("setAuthor", authorInfo);
                //title, subTitle, name, time, tags
                authorInfo.time = this.lastLogin;
                authorInfo.tags = ["博客"];
                commit("setHeader", authorInfo)

                let {data: {list: set}} = await data2;
                for (let i in set) {
                    state.docs.docSet[set[i].id] = new Doc(set[i]);
                }
                dispatch("setCurrentPath",state.currentPath);
                state.total = Number(set.length);
                state._initFinish = true
            },

            setCurrentPath({state, commit}, path) {
                state.currentPath = path;
                if (state.currentPath.startsWith("/doc")) {
                    let currentDoc = state.docs.docSet[state.currentPath.substr(4)]
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
                }
            }
        }
    })
;

export default store


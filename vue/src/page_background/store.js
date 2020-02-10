import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {
        login: false,

        author: {
            avatar: '',
            mail: '',
            github: '',
            name: '',
            title: '',
            subTitle: '',
            motto: '',
            IPC: '',
        },
        tagList: [],

        currentDoc: null
    },

    mutations: {
        changeCurrentDoc:(state, doc)=>{
            state.currentDoc = doc
        },

        changeLogin: (state, b) => {
            state.login = b
        },

        //author
        changeAvatar: (state, info) => {
            state.author.avatar = info
        },
        changeMail: (state, info) => {
            state.author.mail = info;
        },
        changeGithub: (state, info) => {
            state.author.github = info;
        },
        changeName: (state, info) => {
            state.author.name = info;
        },
        changeTitle: (state, info) => {
            state.author.title = info;
        },
        changeSubTitle: (state, info) => {
            state.author.subTitle = info;
        },
        changeMotto: (state, info) => {
            state.author.motto = info;
        },
        changeIPC: (state, info) => {
            state.author.IPC = info;
        }
    }
});

export default store


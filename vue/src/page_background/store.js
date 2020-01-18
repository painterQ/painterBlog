import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {
        viewState: {
            index: 1,
            asideSwitch: true,
        },
        login: false,
        logInState:{
            email: "",
            lastLogin: [],  //最近10次的登录ip，时间
            avatar: null, //image对象
            baseInfoForm: {
                mail: '',
                github: '',
            },
            blogInfoForm: {
                nickName: '',
                title: '',
                motto: '',
                IPC: '',
                beforeArchive: '',
                beforeTopic: '',
            },
        }
    },
    // getters: {
    //     doneTodos: state => {
    //         return state.todos.filter(todo => todo.done)
    //     }
    // },
    mutations: {
        changeIndex: (state, index) => {
            state.viewState.index = index
        },
        changeAsideSwitch: state => {
            state.viewState.asideSwitch = !state.viewState.asideSwitch;
        },
        changeLogin: state => {
            state.login = state
        }
    }
});

export default store


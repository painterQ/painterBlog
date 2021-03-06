import Vue from 'vue'
import VueRouter from 'vue-router'
import {Aside, Button, Col, Container, Footer, Header, Main, Row} from "element-ui";

import painterSetting from "../components/setting"
import document from "../components/document"
import notFound from "../components/page404"
import imageManager from "../components/manageImage"
import tagsManager from "../components/manageTags"
import docsManager from "../components/manageDocs"

const routes = [
    {
        path: '/',
        redirect: '/home',
    },
    {
        path: '/home',
        component: painterSetting,
        meta: {
            keepAlive: true
        }
    },
    {
        path: '/document',
        component: document,
        meta: {
            keepAlive: true
        }
    },
    {
        path: '/docs',
        component: docsManager,
    },
    {
        path: '/tags/*',
        component: tagsManager,
    },
    {
        path: '/tags',
        component: tagsManager,
        meta: {
            keepAlive: true
        }
    },
    {
        path: '/image',
        component: imageManager,
        meta: {
            keepAlive: true
        }
    },
    {
        path: '/404',
        component: notFound
    },
    /*含有通配符的路由应该放在最后,谁先定义的，谁的优先级就最高*/
    // {
    //     path: '*',
    //     redirect: '/404'
    // },
];


Vue.use(Button);    //针对组件选项对象中的install方法
Vue.use(Row);
Vue.use(Col);
Vue.use(Header);
Vue.use(Footer);
Vue.use(Main);
Vue.use(Aside);
Vue.use(Container);


Vue.use(VueRouter);

const router = new VueRouter({
    "routes": routes // routes(缩写) 相当于 routes: routes
})

export default router
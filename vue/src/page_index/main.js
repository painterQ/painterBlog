import Vue from 'vue'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue'
// import './mock.js'
import router from './router.js'
import store from './store.js'

//浏览器调试工具,强制打开（否则编译后是关闭的）
Vue.config.devtools = true
Vue.config.productionTip = false;


new Vue({
  router,
  render: h => h(App),
  store,
}).$mount('#app')

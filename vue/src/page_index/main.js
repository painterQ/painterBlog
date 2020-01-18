import Vue from 'vue'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue'
import './mock.js'
import router from './router.js'
import store from './store.js'

Vue.config.productionTip = false;


new Vue({
  router,
  render: h => h(App),
  store,
}).$mount('#app')

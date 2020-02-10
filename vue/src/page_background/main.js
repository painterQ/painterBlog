import Vue from 'vue'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue'
// import './mock.js'
import router from './router.js'
import store from './store.js'
import {moment} from '../api/time'

Vue.config.productionTip = false;
Vue.filter('moment', function (value, formatString) {
  formatString = formatString || 'yyyy-MM-dd'; //'yyyy-MM-dd hh:mm:ss'
  return moment(value, formatString); // 这是时间戳转时间
});

new Vue({
  router,
  render: h => h(App),
  store,
}).$mount('#background_app')

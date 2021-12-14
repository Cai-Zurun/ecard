import Vue from 'vue'
import App from './App.vue'
import router from './router';
import axios from 'axios';
import api from './utils/api.js';
import common from './utils/common.js';
import ElementUI from 'element-ui';
import dataV from '@jiaminghi/data-view'
import 'vue-g2'
import AmapVue from '@amap/amap-vue';

import 'element-ui/lib/theme-chalk/index.css'; // 默认主题
import './assets/css/icon.css';
import './assets/font/iconfont.css';

Vue.config.productionTip = false
AmapVue.config.version = '2.0';  // 默认2.0，这里可以不修改
AmapVue.config.key = 'd184a2ba0c39c25573391072cd6c5658';

Vue.use(ElementUI, {size: 'small'});
Vue.use(dataV);
Vue.use(AmapVue);

/* 挂载到原型属性上 */
Vue.prototype.$api = api;
Vue.prototype.$common = common;
Vue.prototype.$axios = axios;



// http request拦截器 添加一个请求拦截器login路由跳过
axios.interceptors.request.use(function (config) {
  let filterUrl = config.url.substring(config.url.lastIndexOf("/") + 1, config.url.length);
  if ('login' !== filterUrl) {
    let token = window.localStorage.getItem("token");
    if (token) {
      config.headers.Authorization = token; //将token放到请求头发送给服务器
      return config;
    }
  }
  return config;
}, function (error) {
  return Promise.reject(error);
})

//使用钩子函数对路由进行跳转权限判断
router.beforeEach((to, from, next) => {
  if (!localStorage.getItem('token') && to.name !== 'Login') {
    next({
      name: 'Login'
    })
    return
  }
  if (localStorage.getItem('token') && to.name === 'Login') {
    next({
      name: 'Dashboard'
    })
    return
  }
  if (to.path === '/') {
    next({
      name: 'Dashboard'
    })
    return
  }
  next()
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')

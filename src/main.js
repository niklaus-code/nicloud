// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import axios from 'axios'
import qs from 'qs'
import Vue from 'vue'
import App from './App'
import cookie from './cookie/cookie'
import router from './router'
import store from './store'
// import $ from 'jquery'
Vue.prototype.$http = axios
Vue.prototype.$qs = qs
Vue.prototype.cookie = cookie

Vue.config.productionTip = false

// axios.defaults.baseURL = 'http://192.168.1.252:1992/'

axios.interceptors.response.use(
  function(response) {
    if (response.data.err != null) {
      if (response.data.err.Message === '认证过期，请重新登陆') {
        router.push('/login')
      }
    }

    return response
  },
  function(error) {
    return Promise.reject(error)
  }
)

axios.interceptors.request.use(
  function(config) {
    let token = sessionStorage.getItem('token')
    config.headers.token = token
    return config
  },
  function(error) {
    router.push('/login')
    return Promise.reject(error)
  }
)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})

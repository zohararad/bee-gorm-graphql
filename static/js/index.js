import Vue from 'vue'
import App from './App'
import store from './store'
import router from './routes'

Vue.config.debug = process.env.NODE_ENV !== 'production'

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: { App }
})

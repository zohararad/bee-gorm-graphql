import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const mode = 'history'

const routes = [

]

const router = new VueRouter({
  mode: mode,
  routes, // short for routes: routes
  scrollBehavior (to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { x: 0, y: 0 }
    }
  }
})

export default router

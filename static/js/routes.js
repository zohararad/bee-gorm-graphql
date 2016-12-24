import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from './components/sections/Home'
import Users from './components/sections/Users'
import UsersForm from './components/sections/users/Form'
Vue.use(VueRouter)

const routes = [
  { path: '/', component: Home, name: 'home' },
  { path: '/users', component: Users, name: 'users' },
  { path: '/users/:id/edit', component: UsersForm, name: 'editUser' },
  { path: '/users/create', component: UsersForm, name: 'createUser' }
]

const router = new VueRouter({
  mode: 'history',
  routes,
  scrollBehavior (to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { x: 0, y: 0 }
    }
  }
})

export default router

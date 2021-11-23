import Vue from 'vue'
import Router from 'vue-router'
import login from '@/view/login'
import serveroom from '@/view/machine'
import createmachine from '@/view/createmachine'
import nicloud from '@/view/nicloud'
import store from '../store'


Vue.use(Router)

if (sessionStorage.getItem('token')) {
	store.commit('set_token', sessionStorage.getItem('token'))
}

const router = new Router({
  routes: [
    {
      path: '/nicloud',
      name: 'nicloud',
      component: nicloud
    },
    {
      path: '/createmachine',
      name: 'createmachine',
      component: createmachine
    },
    {
      path: '/serveroom',
      name: 'serveroom',
      component: serveroom
    },
    {
      path: '/',
      name: 'nicloud',
      component: nicloud
    },
    {
      path: '/login',
      name: 'login',
      component: login
    },
  ]
})


router.beforeEach(function(to, from, next) {
	var token = store.state.token
    if (!token) {
        if (to.path !== '/login') {
            return next('/login')
        }
    }
    next()
})

export default router

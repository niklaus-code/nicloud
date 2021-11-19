import Vue from 'vue'
import Router from 'vue-router'
import login from '@/view/login'
import serveroom from '@/view/machine'
import createmachine from '@/view/createmachine'
import nicloud from '@/view/nicloud'


Vue.use(Router)

export default new Router({
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

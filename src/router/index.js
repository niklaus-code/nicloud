import Vue from 'vue'
import Router from 'vue-router'
import blog from '@/view/index'
import details from '@/view/details'
import whisper from '@/view/whisper'
import read from '@/view/read'
import about from '@/view/about'
import login from '@/view/login'
import nicloud from '@/view/vm'
import createvm from '@/view/createvm'
import machine from '@/view/machine'


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/machine',
      name: 'machine',
      component: machine
    },
    {
      path: '/nicloud',
      name: 'nicloud',
      component: nicloud
    },
    {
      path: '/',
      name: 'whisper',
      component: whisper
    },
    {
      path: '/login',
      name: 'whisper',
      component: login
    },
    {
      path: '/blog',
      name: 'blog',
      component: blog
    },
    {
      path: '/details/',
      name: 'details',
      component: details
    },
    {
      path: '/whisper',
      name: 'whisper',
      component: whisper
    },
    {
      path: '/about',
      name: 'about',
      component: about
    },
    {
      path: '/read',
      name: 'read',
      component: read
    }
  ]
})

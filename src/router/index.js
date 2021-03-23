import Vue from 'vue'
import Router from 'vue-router'
import blog from '@/view/index'
import details from '@/view/details'
import whisper from '@/view/whisper'
import read from '@/view/read'
import about from '@/view/about'
import datetime from '@/view/datetime'
import login from '@/view/login'
import clipboard from '@/view/clipboard'


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'whisper',
      component: whisper
    },
    {
      path: '/clipboard',
      name: 'whisper',
      component: clipboard
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
    },
    {
      path: '/datetime',
      name: 'datetime',
      component: datetime
    }
  ]
})

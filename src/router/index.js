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
import gocloud from '@/view/vm'
import createvm from '@/view/createvm'


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/createvm',
      name: 'createvm',
      component: createvm
    },
    {
      path: '/gocloud',
      name: 'gocloud',
      component: gocloud
    },
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

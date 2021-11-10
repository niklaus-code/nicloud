import Vue from 'vue'
import Router from 'vue-router'
import blog from '@/view/index'
import details from '@/view/details'
import whisper from '@/view/whisper'
import read from '@/view/read'
import about from '@/view/about'
import login from '@/view/login'
import serveroom from '@/view/machine'
import createmachine from '@/view/createmachine'
import createhost from '@/view/createhost'
import migratevm from '@/view/migratevm'
import createvlan from '@/view/createvlan'
import createip from '@/view/createip'
import ips from '@/view/ips'
import createosimage from '@/view/createosimage'
import createceph from '@/view/createceph'
import createcloudrive from '@/view/createcloudrive'
import mountcloudrive from '@/view/mountcloudrive'
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
      path: '/mountcloudrive',
      name: 'mountcloudrive',
      component: mountcloudrive
    },
    {
      path: '/createcloudrive',
      name: 'createcloudrive',
      component: createcloudrive
    },
    {
      path: '/createceph',
      name: 'createceph',
      component: createceph
    },
    {
      path: '/createosimage',
      name: 'createosimage',
      component: createosimage
    },
    {
      path: '/ips',
      name: 'ips',
      component: ips
    },
    {
      path: '/createip',
      name: 'createip',
      component: createip
    },
    {
      path: '/createvlan',
      name: 'createvlan',
      component: createvlan
    },
    {
      path: '/migratevm',
      name: 'migratevm',
      component: migratevm
    },
    {
      path: '/createhost',
      name: 'createhost',
      component: createhost
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

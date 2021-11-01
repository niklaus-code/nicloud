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
import serveroom from '@/view/machine'
import createmachine from '@/view/createmachine'
import createhost from '@/view/createhost'
import hosts from '@/view/hosts'
import migratevm from '@/view/migratevm'
import network from '@/view/network'
import createvlan from '@/view/createvlan'
import createip from '@/view/createip'
import ips from '@/view/ips'
import osimage from '@/view/osimage'


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/osimage',
      name: 'osimage',
      component: osimage
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
      path: '/network',
      name: 'network',
      component: network
    },
    {
      path: '/migratevm',
      name: 'migratevm',
      component: migratevm
    },
    {
      path: '/hosts',
      name: 'hosts',
      component: hosts
    },
    {
      path: '/createhost',
      name: 'createhost',
      component: createhost
    },
    {
      path: '/createvm',
      name: 'createvm',
      component: createvm
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

<template>
  <div>
    <serverhead @toParent="getMag"></serverhead>
    <component @toParent="getMag" v-bind:is="router"></component>
  </div>
</template>

<script>
import server from '@/components/server'
import serverhead from '@/components/serverhead'
import addserver from '@/components/addserver'

export default {
  data() {
    return {
      router: sessionStorage.getItem('router')
    }
  },

  components: {
    server,
    addserver,
    serverhead
  },

  mounted: function() {
    this.getMag('server')
  },

  methods: {
    getMag(router) {
      sessionStorage.setItem('router', router)
      this.router = router
    }
  }
}
</script>

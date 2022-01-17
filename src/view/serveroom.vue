<template>
	<div>
		<serverhead @toParent="getMag"> </serverhead>
		 <component @toParent="getMag" v-bind:is="router"></component>
	<div>
</template>

<script>
import server from '@/components/server'
import serverhead from '@/components/serverhead'
import addserver from '@/components/addserver'

export default {
    data () {
        return {
            router: sessionStorage.getItem('router'),
        }
    },

    components: {
       	server, serverhead, addserver
    },
	
	mounted: function () {
        var router = sessionStorage.getItem('router')
		this.getMag(router)
        },


	methods: {
        getMag(router) {
            sessionStorage.setItem('router', router);
            this.router = router;
            },
        },

}

</script>

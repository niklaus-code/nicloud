<template>
	<div>
        <server></server>
	<div>
</template>

<script>
import server from '@/components/server'
import addserver from '@/components/addserver'

export default {
    data () {
        return {
            router: sessionStorage.getItem('router'),
        }
    },

    components: {
       	server,addserver
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

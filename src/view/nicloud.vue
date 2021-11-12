<template>
<div>
    <nicloudhead></nicloudhead>
	<vmleft @toParent="getMag"></vmleft>
	<component @toParent="getMag" v-bind:is="router"></component>
  </div>
</template>

<script>

import foot from '@/components/footer'
import nicloudhead from '@/components/nicloudhead'
import vmleft from '@/components/vmleft'
import vm from '@/components/vm'
import disk from '@/components/disk'
import osimage from '@/components/osimage'
import network from '@/components/network'
import hosts from '@/components/hosts'
import storage from '@/components/storage'
import updateosimage from '@/components/updateosimage'
import createvm from '@/components/createvm'
import createvdisk from '@/components/createvdisk'
import mountvdisk from '@/components/mountvdisk'
import createosimage from '@/components/createosimage'
import createvlan from '@/components/createvlan'
import ips from '@/components/ips'
import createip from '@/components/createip'
import createhost from '@/components/createhost'
import createceph from '@/components/createceph'

export default {
    data () {
        return {
			router: localStorage.getItem('router'),
        }
    },

    components: {
        foot, nicloudhead, vmleft, vm, disk, osimage, network, hosts, storage, createvm, updateosimage, createvdisk, mountvdisk, createosimage, createvlan, ips, createip, createhost, createceph
    },

	methods: {
    	getMag(router) {
			localStorage.setItem('router', router);
      		this.router = router;
    		},
  		},
  }
</script>

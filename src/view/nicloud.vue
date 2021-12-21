<template>
<div>
    <nicloudhead></nicloudhead>
	<vmleft @toParent="getMag"></vmleft>
	<div style="padding-right:15px" class="col-xs-11">
	   	<div class="content1 whisper-content leacots-content details-content col-xs-12" style="background-color:white; float:left">
			<component @toParent="getMag" v-bind:is="router"></component>
		</div>
	</div>
    <vmbottom></vmbottom>
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
import migratevm from '@/components/migratevm'
import changeparam from '@/components/changeparam'
import createdatacenter from '@/components/createdatacenter'
import datacenter from '@/components/datacenter'
import vmbottom from '@/components/vmbottom'

var initroute 
if (sessionStorage.getItem('router')) {
	
	} else {
		sessionStorage.setItem('router', "vm");
		}

export default {
    data () {
        return {
			router: sessionStorage.getItem('router'),
        }
    },

    components: {
        vmbottom, datacenter, createdatacenter, changeparam, foot, nicloudhead, vmleft, vm, disk, osimage, network, hosts, storage, createvm, updateosimage, createvdisk, mountvdisk, createosimage, createvlan, ips, createip, createhost, createceph, migratevm
    },


	mounted: function () {
		
		},

	methods: {
    	getMag(router) {
			sessionStorage.setItem('router', router);
      		this.router = router;
    		},
  		},
  }
</script>

<style>
.content1 {
    box-shadow: 0 0 10px rgba(0,0,0,8);
    border-radius: 10px/10px;
    padding: 30px 0px 30px 0px;
    margin-left: 0px;
    margin-top: 30px;
    margin-bottom: 50px;
}
</style>

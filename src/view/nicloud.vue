<template>
<div class="contain col-md-12">
    <div class="head">
        <nicloudhead @toParent="getMag"></nicloudhead>
	</div>
    <div class="left col-md-1">
        <vmleft @toParent="getMag"></vmleft>
    </div>

	<div  class="mid col-md-11">
        <component @toParent="getMag" v-bind:is="router"></component>
	</div>
    <div class="foot">
         <vmbottom></vmbottom>
    </div>
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
import migratevmlive from '@/components/migratevmlive'
import changeparam from '@/components/changeparam'
import createdatacenter from '@/components/createdatacenter'
import datacenter from '@/components/datacenter'
import vmbottom from '@/components/vmbottom'
import snap from '@/components/snap'
import createsnap from '@/components/createsnap'
import restorevm from '@/components/restorevm'

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
        migratevmlive, restorevm, createsnap, snap, vmbottom, datacenter, createdatacenter, changeparam, foot, nicloudhead, vmleft, vm, disk, osimage, network, hosts, storage, createvm, updateosimage, createvdisk, mountvdisk, createosimage, createvlan, ips, createip, createhost, createceph, migratevm
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

<style scoped>
.contain {
  min-height:100%;
  height: auto !important;
  height: 100%; /*IE6不识别min-height*/
  position: relative;
}

.head {
    background: #ff0;
    background-color: #e3e3e3;
    padding-top:10px;
    padding-bottom:20px;
}

.mid {
    padding-top:30px;
    padding-bottom: 100px;/*等于footer的高度*/
}

.foot {
    position: absolute;
    bottom: 0;
    width: 100%;
    height: 39px;/*脚部的高度*/
    clear:both;
}

.col-md-12 {
    padding-left:0;
    padding-right:0;
}

.left {
    background-color: #778899;
    bottom: 0;
    min-height:100%;
    height: auto !important;
    height: 100%; /*IE6不识别min-height*/
    position: absolute;
    border-left: 1px solid #e3e3e3;
    min-height:100%;
}

.col-md-11, .col-md-1, .col-md-12 {
    display:inline-block;
}
</style>

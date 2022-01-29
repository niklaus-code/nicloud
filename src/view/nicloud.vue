<template>
<div class="contain col-md-12">
    <div class="head">
        <nicloudhead @toParent="getMag"></nicloudhead>
	</div>
    <div class="col-md-1 left">
        <vmleft @toParent="getMag"></vmleft>
    </div>

	<div  class="mid col-md-11">
        <div class="mm">
            <component @toParent="getMag" v-bind:is="router"></component>
	    </div>
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
        height: 100%;
    margin: 0;
    padding: 0;
}

.head {
    border-top-right-radius: 4px;
    border-bottom-right-radius: 4px;
    border-bottom-left-radius: 4px;
    font-size: 15px;

    color: #FFF;
    background-color: #5B5B5B;
    padding-top:3px;
    padding-bottom:3px;
}

.mid {
    min-height: 100%;
    height: auto !important;
    height: 100%;
}

.mm {
    padding-bottom: 60px;
}

.foot {
    position: relative;
    margin-top: -60px; /*等于footer的高度*/
    height: 30px;
    clear:both;
    background: #e3e3e3;
}

.col-md-12 {
    padding-left:0;
    padding-right:0;
}

.left {
    padding-left:0;
    padding-right:0;
    background-color: #778899;
    bottom: 0;
    height: auto !important;
    height: 100%; /*IE6不识别min-height*/
    border-left: 1px solid #e3e3e3;
    min-height:100%;
}

.col-md-11, .col-md-1, .col-md-12 {
    display:inline-block;
}

</style>

<template>
<div class="contain col-md-12">
    <div class="head col-md-12" style="z-index: 100">
        <div class="col-md-5" style="float:left; padding-left:10px">
            <nicloudhead @toParent="getMag"></nicloudhead>
        </div>
        <div class="col-md-2" style="float:right; padding-right:0;margin-right:10px; text-align:right">
            <div class="dropdown">
                <button type="button" @mouseover="mouseover" :style="active" class="btn dropdown-toggle" id="dropdownMenu1" data-toggle="dropdown">
                     <span class="glyphicon glyphicon-user"></span>
                        {{username}}
                    <span class="caret"></span>
                </button>
                <ul class="dropdown-menu  pull-right" role="menu" aria-labelledby="dropdownMenu1">
                    <li role="presentation">
                        <a role="menuitem" tabindex="-1" href="#">修改密码</a>
                    </li>
                    <li role="presentation">
                        <a @click="usermanage" role="menuitem" tabindex="-1">用户管理</a>
                    </li>
                    <li role="presentation">
                        <a @click="logout" role="menuitem" tabindex="-1" href="#"><span class="glyphicon glyphicon-log-out"></span>&nbspLogout</a>
                    </li>
                    <li role="presentation" class="divider"></li>
                    <li role="presentation">
                        <a role="menuitem" tabindex="-1" href="#" style="color: blue">关于NICLOUD</a>
                    </li>
                </ul>
               </div>
           </div>
    </div>

    <div class="left">
        <vmleft @toParent="getMag"></vmleft>
    </div>

	<div  class="mid" >
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
import restorevm from '@/components/restorevm'
import user from '@/components/user'
import createuser from '@/components/createuser'

var initroute 
if (sessionStorage.getItem('router')) {
	
	} else {
		sessionStorage.setItem('router', "vm");
		}

export default {
    data () {
        return {
             active: "",
            username: "",
			router: sessionStorage.getItem('router'),
        }
    },

    components: {
        createuser, user, migratevmlive, restorevm, snap, vmbottom, datacenter, createdatacenter, changeparam, foot, nicloudhead, vmleft, vm, disk, osimage, network, hosts, storage, createvm, updateosimage, createvdisk, mountvdisk, createosimage, createvlan, ips, createip, createhost, createceph, migratevm
    },


	mounted: function () {
        this.getuser()
		
		},

	methods: {
        usermanage: function () {
    	    this.getMag("user");
            },

        mouseover: function () {
            this.active = "color: white";
            },

        logout: function () {
            sessionStorage.removeItem("token");
            this.$router.push({name:"login"});
            },

        getuser: function () {
            var u = this.$store.state.username
            if (u === null || typeof u === 'undefined' || u === '' || u === "undefined") {
                this.username = sessionStorage.getItem('username')
                } else {
                sessionStorage.setItem('username', this.$store.state.username)
                this.username =  this.$store.state.username
                }
                },

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
    height: 35px;
    border-bottom: 1px solid #a49595;
    border-top-right-radius: 2px;
    border-bottom-right-radius: 2px;
    border-bottom-left-radius: 2px;
    font-size: 15px;

    color: #FFF;
    padding-top:5px;
    padding-bottom:5px;
    padding-right:3px;
    padding-left:3px;
}

.mid {
    margin-left:140px;
    padding-right:0;
    padding-left:0;
    padding-top:20px;
    padding-bottom: 50px;
    min-height: 100%;
    height: auto !important;
    height: 100%;
}

.mm {
    padding-top:30px;
    padding-bottom: 60px;
}

.foot {
    margin-left:140px;
    position: relative;
    margin-top: -30px; /*等于footer的高度*/
    height: 30px;
    clear:both;
    #background: #e3e3e3;
}

.dropdown-menu  li  a {
    padding-right: 13px;
    padding-left: 13px;
    }

.col-md-12 {
    padding-left:0;
    padding-right:0;
}

.left {
    z-index: 2500;
    position: fixed;
    float: left;
    height: 100%;

    width: 140px;

    padding-left:0;
    padding-right:0;
    background-color: #393f44;
    bottom: 0;
    height: auto !important;
    height: 100%; /*IE6不识别min-height*/
    min-height:100%;
}

.col-md-11, .col-md-1, .col-md-12 {
    display:inline-block;
}
p {
    margin-bottom:0
}

#dropdownMenu1 {
    background-color: #5b5b5b;
}

.btn {
    padding:0;
}

ul {
    min-width:0;
}
</style>

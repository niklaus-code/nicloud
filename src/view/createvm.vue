<template>
	<div>
	<nicloudhead></nicloudhead>
	<vmleft></vmleft>
  	<div class="content whisper-content leacots-content details-content col-md-11 col-md-offset-2" style="background-color:white; float:left">
		<div  class="col-sm-3 col-sm-offset-4" style="margin-top:20px">
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>cpu/内存</label>
				</div>
				<div class="col-sm-9">
        			<select class="col-sm-10" v-model="flavorvalue">
  						<option  v-for="f in flavorlist" :value="f">
							{{ f.Cpu}}核/ {{f.Mem}}G
						</option>
        			</select>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>VLAN</label>
				</div>
				<div class="col-sm-9">
        			<select class="col-sm-10" v-model="vlanvalue" @change="getip">
  						<option  v-for="v in vlanlist" :value="v">
							{{ v.Vlan }}
						</option>
        			</select>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>IP</label>
				</div>
				<div class="col-sm-9">
	                <select class="col-sm-10" v-model="ipvalue">
                    	<option  v-for="ip in iplist" :value="ip">
                        	{{ ip.Ipv4 }}
						</option>
        			</select>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>宿主机</label>
				</div>
				<div class="col-sm-9">
        			<select class="col-sm-10" v-model="hostvalue">
  						<option  v-for="host in hostlist" :value="host.Ipv4">
							{{ host.Ipv4 }}
						</option>
        			</select>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>镜像</label>
				</div>
				<div class="col-sm-9">
        			<select class="col-sm-10" v-model="imagevalue">
  						<option  v-for="image in imagelist" :value="image.Osname">
							{{ image.Osname }}
						</option>
        			</select>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-2 col-sm-offset-4">
					<button class="btn btn-success btn-sm"  @click="createvm()">创建</button>
				</div>
    		</div>
		</div>
	</div>		
	</div>		
</template>
<script>
import foot from '@/components/footer'
import nicloudhead from '@/components/nicloudhead'
import vmleft from '@/components/vmleft'


export default {
    data () {
        return {
			vlanvalue: "",
			vlanlist: [],

			imagevalue: "",
            imagelist: [],

			hostvalue: "",
            hostlist: [],

			ipvalue: "",
            iplist: [],

			flavorvalue: {},
			flavorlist: [],
			status: {
				0: "关机",
				1: "运行"
				}
        }
    },

	created (){
     this.ipvalue = this.iplist[1]
     this.flavorvalue =this.flavorlist[1]
 
	},

    components: {
        foot, nicloudhead, vmleft
    },

    created: function () {
		this.getflavor()
		this.gethosts()
		this.getimage()
		this.getvlan()
		this.getip()
    },

    methods: {
		createvm: function () {
            var apiurl = `/api/vm/create`

			if (typeof this.ipvalue === 'undefined' || this.ipvalue == null || this.ipvalue === '') {
				alert("缺少信息!")
				return
			}
            this.$http.get(apiurl, { params: { cpu: this.flavorvalue.Cpu, mem:this.flavorvalue.Mem, ip: this.ipvalue.Ipv4, mac: this.ipvalue.Macaddr, host: this.hostvalue, image: this.imagevalue} }).then(response => {
				if (response.data.res) {
					alert("创建成功! 是否查看虚拟机列表")
					this.$router.push('/nicloud')
				} else {
					alert("创建失败('"+response.data.err.Message+"')")
					}
			})
			},

        getimage: function () {
            var apiurl = `/api/vm/getimage`
            this.$http.get(apiurl).then(response => {
            this.imagelist = response.data.res
			this.imagevalue = response.data.res[0].Osname
            })
        },

        gethosts: function () {
            var apiurl = `/api/vm/gethosts`
            this.$http.get(apiurl).then(response => {
            this.hostlist = response.data.res
			this.hostvalue = response.data.res[0].Ipv4
            })
        },

        getvlan: function () {
            var apiurl = `/api/networks/getvlan`
            this.$http.get(apiurl).then(response => {
            this.vlanlist = response.data.res
			this.vlanvalue = response.data.res[0]
            })
        },

        getip: function () {
			var v = ""
			if (typeof this.vlanvalue.Vlan === 'undefined' || this.vlanvalue.Vlan == null || this.vlanvalue.Vlan === '') {
				v = "vlan82"
				} else {
				v = this.vlanvalue.Vlan
				}
            var apiurl = `/api/networks/getip`
            this.$http.get(apiurl, { params: { vlan: v}}).then(response => {
            this.iplist = response.data.res
			this.ipvalue = response.data.res[0]
            })
        },

        getflavor: function () {
            var apiurl = `/api/vm/getflavor`
            this.$http.get(apiurl).then(response => {
            this.flavorlist = response.data.res
			this.flavorvalue = response.data.res[0]
            })
        }
    }
  }
</script>
<style scoped>
label {
    font-weight : 400;
    margin-top: 2px;
}

select{
    font-family: "微软雅黑";
    border: 1px #1a1a1a solid;
    border-radius: 5px;
}

.content {
    box-shadow: 0 0 10px rgba(0,0,0,8);
    border-radius: 10px/10px;
    z-index: -1;
    padding: 70px 0px 70px 0px;
    margin-left: 0px;
    margin-TOP: 50px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}
</style>

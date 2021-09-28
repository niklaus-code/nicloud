<template>
	<div>
	<nicloudhead></nicloudhead>
  	<div class="content whisper-content leacots-content details-content">
		<div  class="col-sm-3 col-sm-offset-4" style="margin-top:20px">
	 		<div class="form-group">
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
		</div>
		<div  class="col-sm-3 col-sm-offset-4" style="margin-top:20px">
	 		<div class="form-group">
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
		</div>
		<div  class="col-sm-3 col-sm-offset-4" style="margin-top:20px">
	 		<div class="form-group">
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
		</div>

		<div  class="col-sm-3 col-sm-offset-4" style="margin-top:20px">
	 		<div class="form-group">
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
		</div>
		<div class="col-sm-3 col-sm-offset-4" style="margin-top:20px">
			<button class="layui-btn"  @click="createvm()">创建</button>
		</div>
	</div>		
	</div>		
</template>
<script>
import foot from '@/components/footer'
import nicloudhead from '@/components/nicloudhead'


export default {
    data () {
        return {
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
        foot, nicloudhead
    },

    created: function () {
		this.getflavor()
		this.getip()
		this.gethost()
		this.getimage()
        var id = this.$route.params.id
        if (!id) {
            this.blogid = window.sessionStorage.getItem('blogid')
        } else {
            this.blogid = id
            window.sessionStorage.setItem('blogid', this.blogid)
        }
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
					alert("创建失败("+response.data.err+")")
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

        gethost: function () {
            var apiurl = `/api/vm/gethost`
            this.$http.get(apiurl).then(response => {
            this.hostlist = response.data.res
			this.hostvalue = response.data.res[0].Ipv4
            })
        },

        getip: function () {
            var apiurl = `/api/vm/getip`
            this.$http.get(apiurl).then(response => {
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
<style>
    .details-content .article-cont p {
    padding:30px 0 0 5px
}
</style>

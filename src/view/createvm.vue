<template>
	<div>
	<headd></headd>
  	<div class="content whisper-content leacots-content details-content">
		<div  class="col-sm-3 col-sm-offset-4" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-2">
        			<label>内存</label>
				</div>
				<div class="col-sm-9">
        			<select class="col-sm-10">
            			<option value="4">2</option>
  						<option  v-for="m in mems" :value="m.v">
							{{ m.v }}
						</option>
        			</select>
				</div>
    		</div>
		</div>
		<div  class="col-sm-3 col-sm-offset-4" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-2">
        			<label>CPU</label>
				</div>
				<div class="col-sm-9">
        			<select class="col-sm-10">
            			<option value="4">2</option>
  						<option  v-for="c in cpus" :value="c.v">
							{{ c.v }}
						</option>
        			</select>
				</div>
    		</div>
		</div>
		<div  class="col-sm-3 col-sm-offset-4" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-2">
        			<label>IP</label>
				</div>
				<div class="col-sm-9">
        			<select class="col-sm-10" v-model="ipvalue">
  						<option  v-for="ip in iplist" :value="ip.Ipv4">
							{{ ip.Ipv4 }}
						</option>
        			</select>
				</div>
    		</div>
		</div>
		<div  class="col-sm-3 col-sm-offset-4" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-2">
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

		<div class="col-sm-3 col-sm-offset-4" style="margin-top:20px">
			<button class="layui-btn"  @click="createvm()">创建</button>
		</div>
	</div>		
	</div>		
</template>
<script>

import foot from '@/components/footer'
import headd from '@/components/head'

export default {
    data () {
        return {
			hostvalue: "",
            hostlist: [],
			ipvalue: "",
            iplist: [],
			memvalue: 2,
			mems: [
     			{v: 2 },
     			{v: 4 },
     			{v: 8 },
     			{v: 16 },
     			{v: 32 }
   				],
			cpuvalue: 2,
			cpus: [
     			{v: 2 },
     			{v: 4 },
     			{v: 8 },
     			{v: 16 },
     			{v: 32 }
   				],
			status: {
				0: "关机",
				1: "运行"
				}
        }
    },

	created (){
     this.ipvalue = this.iplist[1].Ipv4
	},

    components: {
        foot, headd
    },

    created: function () {
		this.getip()
		this.gethost()
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
            this.$http.get(apiurl, { params: { cpu: this.cpuvalue, mem: this.memvalue, ip: this.ipvalue, host: this.hostvalue} }).then(response => {
				if (response.data.res) {
					alert("创建成功! 是否查看虚拟机列表")
					this.$router.push('/gocloud')
				} else {
					alert(response.data.err.Message)
					}
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
			this.ipvalue = response.data.res[0].Ipv4
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

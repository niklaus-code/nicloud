<template>
<div>
    <div class="col-sm-6 col-sm-offset-2" style="margin-top:20px">
	    <div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label style="margin-top:0 ">数据中心</label>
				</div>
				<div class="col-sm-6">
                       {{ hostinfo.Datacenter }}
    		    </div>
			</div>
    	</div>
		<div class="col-sm-12" style="margin-top: 10px">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label style="margin-top:0 ">IP</label>
				</div>
				<div class="col-sm-6">
                       {{ hostinfo.Ipv4 }}
				</div>
    		</div>
    	</div>
		<div class="col-sm-12" style="margin-top: 10px">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>VLAN子网</label>
				</div>
                <div class="col-sm-8" style="float: left">
                    <div style="width: 45%; float: left ">
                        <div class="panel panel-success">
                            <div class="panel-heading">可选VLAN <span class="glyphicon glyphicon-question-sign"></span></div>
                            <ul class="panel-body">
                                <li v-for="(item, index) in vlanlist" @click=movevlan(index)>{{item}}</li>
                            </ul>
                        </div>
                    </div>
                    <div style="width: 10%;float: left; text-align: center; padding-top: 45px">
                        <p style="margin-bottom:0"><span class="glyphicon glyphicon-chevron-left"></span></p>
                        <p><span class="glyphicon glyphicon-chevron-right"></span></p>
                    </div>
                    <div style="width: 45%;float: left ">
                        <div class="panel panel-success">
                            <div class="panel-heading">已选VLAN <span class="glyphicon glyphicon-question-sign"></span></div>
                            <ul class="panel-body">
                                <li v-for="(item, index) in choosen_vlan" @click=backvlan(index)>{{item}}</li>
                            </ul>
                        </div>
                    </div>
                </div>
    		</div>
    	</div>
		<div class="col-sm-12" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-4" style="padding-right: 8px">
        			<label>CPU（核）</label>
				</div>
				<div class="col-sm-4">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="cpu" placeholder="100">
  						</div>
					</form>
				</div>
    		</div>
    		</div>
		<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4" style="padding-right: 8px">
        			<label>内存（G）</label>
				</div>
				<div class="col-sm-4">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="mem" placeholder="300">
  						</div>
					</form>
				</div>
    		</div>
    	</div>
		<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>可创建数量</label>
				</div>
				<div class="col-sm-4">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="num" placeholder="20">
  						</div>
					</form>
				</div>
    		</div>
    	</div>
        <div class="col-sm-12">
            <div class="col-sm-9 col-md-offset-4" style="color: #C0C0C0">
                *创建宿主机之前需要配置好免密登陆
            </div>
        </div>

        <div class="col-sm-12">
		    <div class="form-group">
			    <div class="col-sm-3 col-md-offset-4" style="margin-top:20px" >
  				    <button type="submit" @click="commit" class="btn btn-success btn-sm">提交</button>
			    </div>
		    </div>
		</div>
	</div>
</div>
</template>
<script>
export default {
    data () {
        return {
            hostinfo: {},
            hostip: "",
            mem: "",
            cpu: "",
            num: "",
            vlanlist: [],
            choosen_vlan: [],
            ip: "",
        }
    },

    components: {
        },

	created: function () {
		this.getdatacenter()
        this.gethostinfo()
		},

    methods: {
        gethostbyip: function (ip) {
            var host = this.$route.query.host
            var apiurl = `/api/hosts/gethostsbyip`
            this.$http.get(apiurl, { params: { ip: ip} } ).then(response => {
                this.hostinfo = response.data.res
                this.ip = response.data.res["Ipv4"]
                this.cpu = response.data.res["Cpu"]
                this.mem = response.data.res["Mem"]
                this.num = response.data.res["Max_vms"]
                for (var k in response.data.res.vlan) {
                    this.choosen_vlan.push(response.data.res.vlan[k].Vlan)
                    }
                for (var k in response.data.res.unselectvlan) {
                    this.vlanlist.push(response.data.res.unselectvlan[k].Vlan)
                    }
            })
        },

        gethostinfo: function () {
            var ipv4 = this.$store.state.host.ip
            if (ipv4 === null || typeof ipv4 === "undefined" || ipv4 === "" || ipv4 === "undefined") {
                this.hostip = sessionStorage.getItem('hostip')
                this.gethostbyip(sessionStorage.getItem('hostip'))
                } else {
                    this.hostip = this.$store.state.host.ip
                    this.hostip = sessionStorage.setItem('hostip', ipv4)
                    this.gethostbyip(this.$store.state.host.ip)
                    }
            },

        movevlan: function (index) {
            this.choosen_vlan.push(this.vlanlist[index])
            this.vlanlist.splice(index, 1)
        },

        backvlan: function (index) {
            this.vlanlist.push(this.choosen_vlan[index])
            this.choosen_vlan.splice(index, 1)
        },

    	getdatacenter: function () {
            var apiurl = `/api/datacenter/getdatacenter`
            
            this.$http.get(apiurl).then(response => {
                if (response.data.err === null) {
                    this.datacenter = response.data.res
                    this.centervalue = response.data.res[0].Datacenter
                } else {
                    alert("获取数据失败(" + response.data.err.Message+ ")" )
                    }
            })
            },

		commit: function () {
            var apiurl = `/api/hosts/updatehost`

            this.$http.post(apiurl, this.$qs.stringify({cpu: this.cpu, mem:this.mem, ip: this.ip, num: this.num, vlan: JSON.stringify(this.choosen_vlan)})).then(response => {
				if (response.data.err === null) {
					alert("修改成功! 是否查看宿主机列表")
					this.$emit("toParent", "hosts");
				} else {
					alert(response.data.err.Message)
					}
			})
			},

        }
  }
</script>

<style scoped>
.form-control {
	height:30px;
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
}

select {
    background-color: #fff;
    border-radius: 4px;
    box-shadow: 0 1px 1px rgba(0,0,0,.05);
    width: 100%;
    height: 30px;
    border: 1px #ccc solid;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
    font-weight : 600;
	margin-top: 5px;
    float: right;
}
.panel {
    margin-bottom: 0;
    min-height: 160px;
}

</style>

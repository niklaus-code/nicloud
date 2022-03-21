<template>
<div>
    <div class="col-sm-6 col-sm-offset-2" style="margin-top:20px">
	    <div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>数据中心</label>
				</div>
				<div class="col-sm-6">
				    <select class="col-sm-12" v-model="centervalue">
                        <option  v-for="c in datacenter" :value="c.Datacenter">
                            {{ c.Datacenter }}
                        </option>
                    </select>
    		    </div>
			</div>
    	</div>
		<div class="col-sm-12" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>IP</label>
				</div>
				<div class="col-sm-6">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="ip" placeholder="192.168.1.1">
  						</div>
					</form>
				</div>
    		</div>
    	</div>
		<div class="col-sm-12" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>VLAN子网</label>
				</div>
                <div class="col-sm-8" style="float: left">
                    <div style="width: 45%; float: left ">
                        <div class="panel panel-success">
                            <div class="panel-heading">可选VLAN</div>
                            <ul class="panel-body">
                                <li v-for="(item, index) in vlanlist" @click=movevlan(index)>{{item}}</li>
                            </ul>
                        </div>
                    </div>
                    <div style="width: 10%;float: left; text-align: center; padding-top: 30px">
                        <p style="margin-bottom:0"><span class="glyphicon glyphicon-chevron-left"></span></p>
                        <p><span class="glyphicon glyphicon-chevron-right"></span></p>
                    </div>
                    <div style="width: 45%;float: left ">
                        <div class="panel panel-success">
                            <div class="panel-heading">已选VLAN</div>
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
            choosen_vlan: [],

           	centervalue: "",
            datacenter: [],

            vlanvalue: "",
            vlanlist: [],

			cpu: "",
			mem: "",
			ip: "",
			num: "",
        }
    },

    components: {
        },

	mounted: function () {
		this.getvlan();
		this.getdatacenter()
		},

    methods: {
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

        getvlan: function () {
            var apiurl = `/api/networks/getvlan`
            this.$http.get(apiurl).then(response => {
               for (var i in response.data.res) {
                    console.log(response.data.res[i].Vlan)
                    this.vlanlist.push(response.data.res[i].Vlan)
                    }
                this.vlanvalue = response.data.res[0].Vlan
            })
        },

		commit: function () {
            var apiurl = `/api/hosts/createhost`

            this.$http.post(apiurl, this.$qs.stringify({datacenter:this.centervalue, cpu: this.cpu, mem:this.mem, ipv4: this.ip, max_vms: this.num, vlan: JSON.stringify(this.choosen_vlan)})).then(response => {
				if (response.data.err === null) {
					alert("创建成功! 是否查看宿主机列表")
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
	font-weight : 500;
	margin-top: 5px;
    float: right;
}
.panel {
    margin-bottom: 0;
    min-height: 160px;
}

</style>

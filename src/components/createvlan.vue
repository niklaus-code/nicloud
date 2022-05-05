<template>
<div>
  <div class="col-sm-12 form-group" style="margin-top:30px; border-bottom: 1px green solid">
                <h4>创建网络</h4>
            </div>

		<div class="col-sm-6 col-sm-offset-2" style="margin-top:20px">
		<div class="col-sm-12" style="margin-top:20px">

				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-5">
        			<label>数据中心</label>
				</div>
				<div class="col-sm-7">
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
				<div class="col-sm-5">
        			<label>VLAN子网</label>
				</div>
				<div class="col-sm-7">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="vlan" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
				</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-5">
        			<label>网桥</label>
				</div>
				<div class="col-sm-7">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="bridge" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-5">
        			<label>地址段</label>
				</div>
				<div class="col-sm-7">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="network" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-5">
        			<label>子网掩码</label>
				</div>
				<div class="col-sm-7">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="prefix" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-5">
        			<label>网关</label>
				</div>
				<div class="col-sm-7">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="gateway" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
    		</div>
				<div class="col-sm-12">
		<div class="form-group" style="margin-top:20px" >
			<div class="col-sm-2 col-sm-offset-5">
  				<button type="submit" @click="createvlan" class="btn btn-success btn-sm">提交</button>
			</div>
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
			centervalue: "",
			datacenter: [],
			vlan: "",
			bridge: "",
			network: "",
			prefix: "",
			gateway: "",
        }
    },

	mounted: function() {
		this.getdatacenter()
		},

    methods: {
		check: function (datacenter, vlan, bridge, network, prefix, gateway) {
			if (typeof datacenter === 'undefined' || datacenter === null || datacenter === ''|| typeof vlan === 'undefined' || vlan === null || vlan === ''|| typeof bridge === 'undefined' || bridge === null || bridge === '' || typeof network === 'undefined' || network === null || network === '' ||typeof prefix === 'undefined' || prefix === null || prefix === ''|| typeof gateway === 'undefined' || gateway === null || gateway === '' ) {

				alert("缺少信息")
                return true
            } else {
				return false
				}
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

		createvlan: function () {
			if (this.check(this.datacenter, this.vlan, this.bridge, this.network, this.prefix,this.gateway)) {
				return 
				}
		

            var apiurl = `/api/networks/createvlan`

            this.$http.post(apiurl,  this.$qs.stringify({datacenter: this.centervalue,  vlan: this.vlan, bridge:this.bridge, network: this.network, prefix: this.prefix, gateway: this.gateway})).then(response => {
				if (response.data.err === null) {
					alert("创建成功! 是否查看网络列表")
					this.$emit("toParent", "network");
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
}

.col-sm-5 label{
	float: right;
}
select{
	height:30px;
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
	font-weight : 400;
	margin-top: 5px;
}
</style>

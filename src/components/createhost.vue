<template>
<div>
      <div class="col-sm-12 form-group" style="border-bottom: 1px green solid">
                <h4>创建宿主机</h4>
            </div>

		<div class="col-sm-4 col-sm-offset-4" style="margin-top:20px">
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-3">
        			<label>数据中心</label>
				</div>
				<div class="col-sm-9">
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
				<div class="col-sm-3">
        			<label>网络</label>
				</div>
				<div class="col-sm-9">
				    <select class="col-sm-10" v-model="vlanvalue">
                        <option  v-for="v in vlanlist" :value="v">
                            {{ v.Vlan }}
                        </option>
                    </select>
				</div>
    		</div>
    		</div>
				<div class="col-sm-12" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-3">
        			<label>cpu</label>
				</div>
				<div class="col-sm-9">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="cpu" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-3">
        			<label>内存</label>
				</div>
				<div class="col-sm-9">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="mem" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-3">
        			<label>IP</label>
				</div>
				<div class="col-sm-9">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="ip" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-3">
        			<label>可创建数量</label>
				</div>
				<div class="col-sm-9">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="num" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
    		</div>
		<div class="form-group">
			<div class="col-sm-3 col-sm-offset-6" style="margin-top:20px" >
  				<button type="submit" @click="commit" class="btn btn-success btn-sm">提交</button>
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

            vlanvalue: "",
            vlanlist: [],

			cpu: "",
			mem: "",
			ip: "",
			num: "",
        }
    },

	mounted: function () {
		this.getvlan();
		this.getdatacenter()
		},

    methods: {
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
            this.vlanlist = response.data.res
            this.vlanvalue = response.data.res[0]
            })
        },

		commit: function () {
            var apiurl = `/api/hosts/createhost`

            this.$http.get(apiurl, { params: {datacenter:this.centervalue, cpu: this.cpu, mem:this.mem, ip: this.ip, num: this.num, vlan: this.vlanvalue.Vlan} }).then(response => {
				if (response.data.res === null) {
					alert("创建成功! 是否查看宿主机列表")
					this.$emit("toParent", "hosts");
				} else {
					alert("插入数据失败(" + response.data.res.Message+ ")" )
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
	width: 100%;
	height:	30px;
    border: 1px #ccc solid;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
	font-weight : 400;
	margin-top: 5px;
    float: right;
}
</style>

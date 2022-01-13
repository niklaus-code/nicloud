<template>
	<div>
      <div class="col-sm-12 form-group" style="border-bottom: 1px green solid">
                <h4>创建IP</h4>
            </div>

		<div class="col-sm-8 col-sm-offset-2" style="margin-top:30px;" >
			<div class="col-sm-12">
			<span class="vlaninfo">vlan信息</span>
			</div>
			<div class="col-sm-3">
				vlan名称：{{vlan}}
			</div>
			<div class="col-sm-3">
				网桥名称：{{bridge}}
			</div>
			<div class="col-sm-3">
				地址段：{{network}}/{{prefix}}
			</div>
			<div class="col-sm-3">
				网关：{{gateway}}
			</div>
		</div>
		<div class="col-sm-8 col-sm-offset-2"  style="margin-top:30px; margin-bottom:30px" >
			<div class="col-sm-12">
				<span class="createip">生成IP & MAC</span>
			</div>
			<div class="col-sm-5 startip">
				<form class="form-inline" style="width:100%" >
  					<div class="form-group" style="width:100%" >
    					<label class="sr-only" for="exampleInputAmount">Amount (in dollars)</label>
    					<div class="input-group" style="width:100%" >
      						<div class="input-group-addon">起始IP</div>
      						<input style="width:100%" type="text" class="form-control" v-model="startip" placeholder="10.0.0.1">
    					</div>
  					</div>
				</form>
			</div>
			<div  class="col-sm-5 endip">
				<form class="form-inline" style="width:90%" >
  					<div class="form-group" style="width:100%" >
    					<label class="sr-only" for="exampleInputAmount">Amount (in dollars)</label>
    					<div class="input-group"  style="width:100%" >
      						<div class="input-group-addon">结束IP</div>
      						<input type="text" class="form-control" v-model="endip" placeholder="10.0.0.254">
    					</div>
  					</div>
				</form>
			</div>
			<div  class="col-sm-2">
  					<button type="submit" class="btn btn-sm btn-success" @click="create">生成IP & 入库</button>
  					<button type="submit" class="btn btn-sm btn-info" @click="ips">查看IP</button>
			</div>
		</div>
	</div>
</div>
</template>
<script>
export default {
    data () {
        return {
			vlan: "",
			bridge: "",
			network: "",
			prefix: "",
			gateway: "",
			startip: "",
			endip: "",
        }
    },
	
    created: function () {
        this.vlaninfo()
    },


    methods: {
       ips: function () {
            this.$emit("toParent", "ips");
            this.$store.state.network.vlan = this.vlan
            },

		create: function () {
			if (typeof this.startip === 'undefined' || this.startip === null || this.startip === ''|| typeof this.endip === 'undefined' || this.endip === null || this.endip === '') {
				alert("输入为空")
				return
				}

            var apiurl = `/api/networks/createip`
            this.$http.get(apiurl, { params: {startip: this.startip, endip: this.endip, vlan: this.vlan, prefix: this.prefix, gateway: this.gateway} }).then(response => {
				if (response.data.err === null) {
					alert("创建成功")
					} else {
					alert(response.data.res.Message)
					}
				})
			},

		vlaninfo: function () {
			var v = this.$store.state.network.vlan
			if (v === null || typeof v === 'undefined' || v === '' || v === "undefined") {
			 	this.vlan = sessionStorage.getItem('vlan', this.$store.state.network.vlan)
			 	this.bridge = sessionStorage.getItem('bridge', this.$store.state.network.bridge)
			 	this.network = sessionStorage.getItem('network', this.$store.state.network.network)
			 	this.prefix = sessionStorage.getItem('prefix', this.$store.state.network.prefix)
			 	this.gateway = sessionStorage.getItem('gateway', this.$store.state.network.gateway)
				} else {
				    this.vlan = this.$store.state.network.vlan
				    this.bridge = this.$store.state.network.bridge
				    this.network = this.$store.state.network.network
				    this.prefix = this.$store.state.network.prefix
				    this.gateway = this.$store.state.network.gateway
			 	    sessionStorage.setItem('vlan', this.$store.state.network.vlan)
			 	    sessionStorage.setItem('bridge', this.$store.state.network.bridge)
			 	    sessionStorage.setItem('network', this.$store.state.network.network)
			 	    sessionStorage.setItem('prefix', this.$store.state.network.prefix)
			 	    sessionStorage.setItem('gateway', this.$store.state.network.gateway)
				}
			},
        }
  }
</script>
<style scoped>

.createip {
	font-weight:500
}

.vlaninfo {
	font-weight:501
}
.col-sm-2 {
	padding-left:0;
	margin-top:10px;
}

.col-sm-3 {
	margin-top:10px;
}
.col-sm-8 {
	padding: 10px;
	border-style: solid;
	border-color: #ddd;
	border-width: 1px;
	border-radius: 4px 4px 0 0;
}

.startip {
	margin-top: 10px;
	padding-right: 0px;
}

.endip {
	margin-top: 10px;
	padding-right: 0px;
	padding-left: 0px;
}

.col-sm-4 label{
	float: right;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
	font-weight : 400;
	margin-top: 5px;
}
</style>

<template>
	<div>
		<div class="col-sm-8 col-sm-offset-2" >
			<div class="col-sm-12">
			<span class="vlaninfo">vlan信息</span>
			</div>
			<div class="col-sm-3">
				vlan名称：{{vlan.Vlan}}
			</div>
			<div class="col-sm-3">
				网桥名称：{{vlan.Bridge}}
			</div>
			<div class="col-sm-3">
				地址段：{{vlan.Network}}/{{vlan.Prefix}}
			</div>
			<div class="col-sm-3">
				网关：{{vlan.Gateway}}
			</div>
		</div>
		<div class="col-sm-8 col-sm-offset-2" >
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
  					<button type="submit" class="btn btn-sm btn-primary" @click="create">生成IP</button>
  					<button type="submit" class="btn btn-sm btn-primary">导入数据库</button>
			</div>
		</div>
	</div>
</template>
<script>
export default {
    data () {
        return {
			vlan: "",
			startip: "",
			endip: "",
        }
    },
	
    created: function () {
        this.vlaninfo()
    },


    methods: {
		create: function () {
			if (typeof this.startip === 'undefined' || this.startip === null || this.startip === ''|| typeof this.endip === 'undefined' || this.endip === null || this.endip === '') {
				alert("输入为空")
				return
				}

            var apiurl = `/api/networks/createip`
            this.$http.get(apiurl, { params: {startip: this.startip, endip: this.endip, vlan: this.vlan.Vlan} }).then(response => {
				if (response.data.res === null) {
					alert("创建成功")
					} else {
					alert("创建失败'(" + response.data.res.Message+"')")
					}
				})
			},

		vlaninfo: function () {
			this.vlan = this.$store.state.network.vlan
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
select{
    font-family: "微软雅黑";
    border: 1px #1a1a1a solid;
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

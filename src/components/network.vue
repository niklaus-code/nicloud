<template>
		<div class="col-sm-12" style="margin-top: 10px; padding-left:0; padding-right: 0">
			<button class="btn btn-success btn-sm" type="button" @click="createvlan"  style="margin-right: 20px; margin-bottom: 12px; float: right">创建网络<span class="glyphicon glyphicon-plus" style="margin-left: 5px"></span></button>
			<table class="table table-condensed .table-hover" style="text-align: center;">
    			<thead>
      				<tr>
        				<th>
        					<label class="checkbox-inline" style="border:red 1px">
            					<input type="checkbox" v-model="checkvalue" @click="checkbox()">
        					</label>
        				</th>
        				<th>vlan</th>
        				<th>网桥</th>
        				<th>地址段</th>
						<th>网关</th>
        				<th>数据中心</th>
						<th>状态</th>
						<th>操作</th>
      				</tr>
    			</thead>

				<tbody v-for="(item, index) in data">
      				<tr v-show="item.Status" class="table-dark text-dark" :id="item.Uuid">
        				<label class="checkbox-inline">
            				<input type="checkbox" v-model="item.Checkout">
        				</label>
        				<td>{{item.Vlan}}</td>
        				<td>{{item.Bridge}}</td>
        				<td>{{item.Network}}/{{item.Prefix}}</td>
        				<td>{{item.Gateway}}</td>
        				<td>{{item.Datacenter}}</td>
						 <td>
                            <span class="glyphicon glyphicon-ok"></span>
                        </td>
		    			<td>
							<button class="btn btn-success btn-xs" type="button" @click="createip(item.Vlan, item.Bridge, item.Network, item.Prefix, item.Gateway)">
                                <span class="glyphicon glyphicon-plus"></span>
                				增加IP
            				</button>
							<button class="btn btn-primary btn-xs" type="button" @click="ips(item.Vlan)">
                                <span class="glyphicon glyphicon-zoom-in"></span>
                				查看IP
            				</button>
							<button class="btn btn-danger btn-xs" type="button" @click="deletevlan(item.Vlan, index)">
                                <span class="glyphicon glyphicon-trash"></span>
                				删除
            				</button>
        				</td>
					</tr>
				</tbody>
			</table>
		</div>
</template>
<script>

export default {
    data () {
        return {
			data: [],
        }
    },

	mounted: function () {
		this.getvlan()
		},

    methods: {
		createvlan: function () {
			this.$emit("toParent", "createvlan");
			},

		ips: function (vlan) {
			this.$emit("toParent", "ips");
			this.$store.state.network.vlan = vlan
			},

		createip: function (vlan, bridge, network, prefix, gateway) {
			this.$emit("toParent", "createip");
			this.$store.state.network.vlan = vlan
			this.$store.state.network.bridge = bridge
			this.$store.state.network.network = network
			this.$store.state.network.prefix = prefix
			this.$store.state.network.gateway = gateway
			},

		deletevlan: function (vlan, index) {
            var apiurl = `/api/networks/delete`
            this.$http.get(apiurl, { params: { vlan: vlan} } ).then(response => {
            	if (response.data.err === null) {
					alert("删除成功")
					this.data[index].Status = 0
					} else {
					alert("删除失败('"+response.data.err.Message+"')")	
				}
            })
        },

		getvlan: function () {
            var apiurl = `/api/networks/getvlan`
            this.$http.get(apiurl).then(response => {
            	if (response.data.err === null) {
            	    this.data = response.data.res
                } else {
					alert(response.data.err.Message)
                    }
            })
        },
        }
  }
</script>
<style scoped>

select{
    font-family: "微软雅黑";
    border: 1px #1a1a1a solid;
    border-radius: 5px;
}

.checkbox-inline {
    margin-bottom: 30px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

input {
	margin-top: 2px;
}

label {
	font-weight : 400;
}

.table tbody tr td {
    vertical-align: "middle";
}

th {
    background-color: #e3e3e3;
	font-weight: bold;
    color: black;
    text-align: center;
}

</style>

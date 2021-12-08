<template>
	<div>
		<div class="col-sm-10 col-sm-offset-1 vlan" style="margin-top:20px;">
			<span>{{vlan}}&nbsp&&nbsp IP列表</span>
		</div>
		<div class="col-sm-10 col-sm-offset-1" style="margin-top:20px;">
			<table class="table table-hover" style="text-align: center;">
    			<thead>
      				<tr>
        				<th>
        					<label class="checkbox-inline" style="border:red 1px">
            					<input type="checkbox" v-model="checkvalue" @click="checkbox()">
        					</label>
        				</th>
						<th>IPv4</th>
						<th>MACADDR</th>
						<th>VLAN</th>
						<th>状态</th>
						<th>操作</th>
      				</tr>
    			</thead>

				<tbody v-for="(item, index) in ips">
      				<tr v-if="item.Exist" class="table-dark text-dark" :id="item.Ipv4">
        				<label class="checkbox-inline">
            				<input type="checkbox" v-model="item.Checkout">
        				</label>
        				<td>{{item.Ipv4}}</td>
        				<td>{{item.Macaddr}}</td>
        				<td>{{vlan}}</td>
        				<td>
							<span v-if="item.Status" class="glyphicon glyphicon-remove"></span>
							<span v-else  class="glyphicon glyphicon-ok"></span>
						</td>
		    			<td>
							<button v-if="item.Status" class="btn btn-success btn-xs" type="button" @click="upip(index, item.Ipv4)">
                				UP
            				</button>
							<button v-else class="btn btn-warning btn-xs" type="button" @click="downip(index, item.Ipv4)">
                				DOWN
            				</button>
							<button class="btn btn-danger btn-xs" type="button" @click="deleteip(index, item.Ipv4)">
                				删除
            				</button>
        				</td>
					</tr>
				</tbody>
			</table>
		</div>
	</div>
</template>
<script>
export default {
    data () {
        return {
			vlan: "",
			ips: [],
        }
    },

	created: function () {
        this.vlaninfo()
    },

	mounted: function () {
		this.getip()
		},

    methods: {
	   	vlaninfo: function () {
			var initroute 
			if (this.$store.state.network.vlan) {
				this.vlan = this.$store.state.network.vlan
				sessionStorage.setItem('vlan', this.$store.state.network.vlan)
    			} else {
				this.vlan = sessionStorage.getItem('vlan')
        		}

            },

		 deleteip: function (index, ip) {
            var apiurl = `/api/networks/deleteip`
            this.$http.get(apiurl, { params: {ip: ip, vlan: this.vlan}}).then(response => {
            	if (response.data.err != null) {
					alert("删除失败'("+response.data.err.Message+")'")
				} else {
					alert("删除成功")
					this.ips[index].Exist = false
				}
            })
        },

		 upip: function (index, ip) {
            var apiurl = `/api/networks/upip`
            this.$http.get(apiurl, { params: {ipv4: ip, vlan: this.vlan}}).then(response => {
            	if (response.data.res != null) {
					alert("操作失败'("+response.data.res+")'")
					} else {
						alert("已置为可用状态")
						this.ips[index].Status = 0
						}
            })
        },
		 downip: function (index, ip) {
            var apiurl = `/api/networks/downip`
            this.$http.get(apiurl, { params: {ipv4: ip, vlan: this.vlan}}).then(response => {
            	if (response.data.res != null) {
					alert("操作失败'("+response.data.res+")'")
					} else {
						alert("已置为不可用状态")
						this.ips[index].Status = 1
						}
            })
        },
		 getip: function () {
            var apiurl = `/api/networks/getallip`
            this.$http.get(apiurl, { params: { vlan: this.vlan}}).then(response => {
            	this.ips = response.data.res
            })
        },
  }
  }
</script>
<style>
.vlan {
	font-size: 18px;
	font-weight: 600;
	border-bottom: 3px solid green
}

.checkbox-inline {
    margin-bottom: 30px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

.table tbody tr td {
    padding: 12px;
    vertical-align: "middle";
}

th {
	font-weight: bold;
    color: black;
    text-align: center;
}

</style>

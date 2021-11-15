<template>
	<div>
  	<div class="content whisper-content leacots-content details-content col-md-11 col-md-offset-2" style="background-color:white; float:left">
		<div class="col-sm-10 col-sm-offset-1" style="margin-top:20px">
			<button class="btn btn-success btn-sm" @click="createhost" type="button">创建</button>
			<table class="table table-hover" style="text-align: center;">
    			<thead>
      				<tr>
        				<th>
        					<label class="checkbox-inline" style="border:red 1px">
            					<input type="checkbox" v-model="checkvalue" @click="checkbox()">
        					</label>
        				</th>
        				<th>数据中心</th>
        				<th>网络</th>
        				<th>ip地址</th>
        				<th>cpu</th>
        				<th>内存</th>
        				<th>可创建数量</th>
						<th>备注</th>
						<th>状态</th>
						<th>操作</th>
      				</tr>
    			</thead>

				<tbody v-for="(item, index) in data">
      				<tr v-if="item.Status" class="table-dark text-dark" :id="item.Uuid">
        				<label class="checkbox-inline" style="width:10px">
            				<input type="checkbox" v-model="item.Checkout">
        				</label>
        				<td>{{item.Datacenter}}</td>
        				<td>{{item.Vlan}}</td>
        				<td>{{item.Ipv4}}</td>
        				<td>{{item.Usedcpu}}核/{{item.Cpu}}核</td>
        				<td>{{item.Usedmem}}G/{{item.Mem}}G</td>
        				<td>{{item.count}}/{{item.Max_vms}}</td>
        				<td>test</td>
                        <td>
                            <span class="glyphicon glyphicon-ok"></span>
                        </td>

		    			<td>
							<button class="btn btn-info btn-xs" type="button" @click="deletehost(item.Ipv4, index)">
                				删除
            				</button>
        				</td>
					</tr>
				</tbody>
			</table>
		</div>
	</div>		
</div>		
</template>
<script>

export default {
    data () {
        return {
			data: [],
			cpu: "",
			mem: "",
			ip: "",
			num: "",
        }
    },

	mounted: function () {
		this.gethost()
		},

    methods: {
		createhost: function () {
			this.$emit("toParent", "createhost");
			},

		deletehost: function (ip, index) {
            var apiurl = `/api/hosts/delete`
            this.$http.get(apiurl, { params: {ip: ip} } ).then(response => {
			   if (response.data.res === null) {
                    alert("删除成功")
					this.data[index].Status=0
                    } else {
                    alert("删除失败('"+response.data.res.Message+"')")  
                }
            })
        },

		gethost: function (ip) {
            var apiurl = `/api/hosts/gethosts`
            this.$http.get(apiurl).then(response => {
            	this.data = response.data.res
            })
        },
        }
  }
</script>
<style>

select{
    font-family: "微软雅黑";
    border: 1px #1a1a1a solid;
    border-radius: 5px;
}

.content {
    box-shadow: 0 0 10px rgba(0,0,0,8);
    border-radius: 10px/10px;
    z-index: -1;
    padding: 50px 0px 50px 0px;
    margin-left: 0px;
    margin-TOP: 50px;
}
.checkbox-inline {
    margin-bottom: 30px;
}


.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
	font-weight : 400;
	margin-top: 5px;
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

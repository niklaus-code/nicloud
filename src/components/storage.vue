<template>
		<div class="col-sm-12" style="margin-top:20px;padding-right:0; padding-left:0">
			<button @click="createstorage" class="btn btn-success btn-sm" type="button"  style="margin-bottom:20px;margin-left: 5px">创建存储集群<span class="glyphicon glyphicon-plus" style="margin-left: 5px"></span></button>
			<table class="table table-hover" style="text-align: center;">
    			<thead>
      				<tr>
        				<th>
        					<label class="checkbox-inline" style="border:red 1px">
            					<input type="checkbox" v-model="checkvalue" @click="checkbox()">
        					</label>
        				</th>
        				<th>名称</th>
        				<th>UUID</th>
        				<th>数据池</th>
        				<th>secret</th>
        				<th>hosts/port</th>
        				<th>数据中心</th>
						<th>备注</th>
						<th>状态</th>
						<th>操作</th>
      				</tr>
    			</thead>

				<tbody v-for="(item, index) in data">
      				<tr v-if="item.Status"  class="table-dark text-dark" :id="item.Uuid">
        				<label class="checkbox-inline">
            				<input type="checkbox" v-model="item.Checkout">
        				</label>
        				<td>{{item.Name}}</td>
        				<td>{{item.Uuid}}</td>
        				<td>{{item.Pool}}</td>
        				<td>{{item.Ceph_secret}}</td>
        				<td>{{item.Ips}}/{{item.Port}}</td>
        				<td>{{item.Datacenter}}</td>
        				<td>{{item.Comment}}</td>
						<td>
							<span class="glyphicon glyphicon-ok"></span>
						</td>
		    			<td>
							<button class="btn btn-danger btn-xs" type="button" @click="deletestorage(item.Uuid, index)">
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
		this.getceph()
		},

    methods: {
		createstorage: function () {
            this.$emit("toParent", "createceph");
			},

		deletestorage: function (name, index) {
			this.data[index].Status = 0
            var apiurl = `/api/storage/delete`
            this.$http.get(apiurl, { params: {name: name} } ).then(response => {
			   if (response.data.err === null) {
                    alert("删除成功")
                    } else {
                    alert(response.data.err.Message)  
                }
            })
        },

		getceph: function (ip) {
            var apiurl = `/api/storage/get`
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

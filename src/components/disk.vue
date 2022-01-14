<template>
		<div class="col-sm-12" style="margin-top:20px">
			<button class="btn btn-success btn-sm" @click="create('createvdisk')" type="button">创建</button>
			<table class="table table-hover" style="text-align: center;">
    			<thead>
      				<tr>
        				<th>
        					<label class="checkbox-inline" style="border:red 1px">
            					<input type="checkbox" v-model="checkvalue" @click="checkbox()">
        					</label>
        				</th>
        				<th>云盘</th>
        				<th>容量</th>
        				<th>存储池</th>
        				<th>挂载云主机</th>
        				<th>创建者</th>
        				<th>备注</th>
        				<th>状态</th>
        				<th>创建时间</th>
						<th>操作</th>
      				</tr>
    			</thead>

				<tbody v-for="(item, index) in data" v-if="item.Exist">
      				<tr class="table-dark text-dark" :id="item.Uuid">
        				<label class="checkbox-inline" style="width:10px">
            				<input type="checkbox" v-model="item.Checkout">
        				</label>
        				<td>{{item.Vdiskid}}</td>
        				<td>{{item.Contain}}G</td>
        				<td>{{item.Pool}}</td>
        				<td>{{item.Vm_ip}}</td>
        				<td>{{item.username}}</td>
        				<td>{{item.Comment}}</td>
				      <td>
                            <span v-if="item.Status"  class="glyphicon glyphicon-ok"></span>
                            <span v-else class="glyphicon glyphicon-remove"></span>
                        </td>
        				<td>{{item.Createtime}}</td>
		    			<td>
							<!-- <button v-if="item.Status" class="btn btn-success btn-xs" type="button" @click="mount(, item.Storage, item.Pool, index)"> -->
							<button v-if="item.Status" class="btn btn-success btn-xs" type="button" @click="mount('mountvdisk', item.Vdiskid, item.storage, item.pool)">
                				挂载
            				</button>
							<button v-else class="btn btn-warning btn-xs" type="button" @click="umount(item.Vm_ip, item.Vdiskid, index)">
                				卸载
            				</button>
							<button v-if="item.Status" class="btn btn-danger btn-xs" type="button" @click="deletevdisk(item.Vdiskid, item.Comment, index)">
                				销毁
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
			cpu: "",
			mem: "",
			ip: "",
			num: "",
        }
    },

	mounted: function () {
		this.getvdisk()
		},

    methods: {
	    create: function () {
                this.$emit("toParent", "createvdisk");
                },

		deletevdisk: function (uuid, comment, index) {
            var apiurl = `/api/vdisk/deletevdisk`
            this.$http.get(apiurl , { params: { uuid: uuid, comment: comment} }).then(response => {
				if (response.data.err === null ) {
					this.data[index].Exist = 0
					alert("删除成功")
					} else {
					alert ("获取数据失败（"+response.data.err.Message+")")
					}
			})
			},

		mount: function (router, vdiskid, storage, pool) {
			this.$store.state.vdisk.vdiskid = vdiskid
			this.$store.state.vdisk.storage = storage
			this.$store.state.vdisk.pool = pool
            this.$emit("toParent", router);
			},

		umount: function (vmip, vdiskid, index) {
            var apiurl = `/api/vdisk/umountdisk`
            this.$http.get(apiurl , { params: { vmip: vmip, vdiskid: vdiskid} }).then(response => {
				if (response.data.err === null ) {
					this.data[index].Status = 1
					this.data[index].Vm_ip = ""
					alert("卸载成功")
				} else {
					alert ("获取数据失败（'"+response.data.err.Message+"')")
					}
            })
        },

		getvdisk: function (ip) {
            var apiurl = `/api/vdisk/getvdisk`
            this.$http.get(apiurl).then(response => {
				if (response.data.err === null ) {
            		this.data = response.data.res
				} else {
					alert (response.data.err.Message)
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

input {
    margin-top:2px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
	font-weight : 400;
}

.table tbody tr td {
    vertical-align: "middle";
}

th {
	font-weight: bold;
    color: black;
    text-align: center;
}

</style>

<template>
		<div style="margin-top: 10px">
            <div class="col-md-7" style="padding-left: 24px">
		        <h5><strong>云盘列表（{{countdisk}}）</strong></h5>
            </div>

            <div style="float: right; margin-bottom: 12px" class="col-md-5">
			    <button class="btn btn-success btn-sm" @click="create('createvdisk')" type="button" style="float: right; margin-right: 5px; display:inline;">创建云盘<span class="glyphicon glyphicon-plus" style="margin-left: 5px"></span></button>
            </div>
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
                        <td>
                            <span v-if='item.flag2' @click="c(index)">
                                {{item.Comment}}
                            </span>
                            <li v-if='item.flag'><span class="glyphicon glyphicon-calendar" @click="edit(index)"></span></li>
                            <div v-if='item.flag1'>
                                <div><input type="text" v-model="comments"></div>
                                <div><span  @click="input(index, item.Vdiskid)" class="glyphicon glyphicon-calendar"></span></div>
                            </div>
                        </td>
				        <td>
                            <span v-if="item.Status"  class="glyphicon glyphicon-ok"></span>
                            <span v-else class="glyphicon glyphicon-remove"></span>
                        </td>
        				<td>{{item.Createtime}}</td>
		    			<td>
							<button v-if="item.Status" class="btn btn-primary btn-xs" type="button" @click="mount('mountvdisk', item.Vdiskid, item.storage, item.pool)">
                         	<span class="glyphicon glyphicon-floppy-open"></span>
                				挂载
            				</button>
							<button v-else class="btn btn-warning btn-xs" type="button" @click="umount(item.Vm_ip, item.Vdiskid, index)">
                         	<span class="glyphicon glyphicon-floppy-save"></span>
                				卸载
            				</button>
							<button v-if="item.Status" class="btn btn-danger btn-xs" type="button" @click="deletevdisk(item.Vdiskid, item.Comment, index)">
                                 <span class="glyphicon glyphicon glyphicon-trash"></span>
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
            comments: "",
            countdisk: 0,
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
        c: function (index) {
            this.data[index].flag2 = false
            this.data[index].flag1 = true
            this.comments = this.data[index].Comment
            },

        edit: function (index) {
            this.data[index].flag = false
            this.data[index].flag1 = true
            },

        input: function (index, uuid) {
            var apiurl = `/api/vdisk/addcomment`
            this.$http.post(apiurl,  this.$qs.stringify({uuid: uuid, comment: this.comments} )).then(response => {
                if (response.data) {
                    this.data[index].Comment = this.comments
                    }
            })
            this.data[index].flag = false
            this.data[index].flag1 = false
            this.data[index].flag2 = true
            },

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
					alert (response.data.err.Message)
					}
            })
        },

        comment: function(res) {
            var d = new Array()
            for (var v in res) {
                if (res[v]["Comment"].length > 0) {
                    res[v]["flag"] = false
                    res[v]["flag2"] = true
                    } else {
                        res[v]["flag2"] = false
                        res[v]["flag"] = true
                    }
                res[v]["flag1"] = false
                d.push(res[v])
            this.data = d
                }
            },

		getvdisk: function (ip) {
            var apiurl = `/api/vdisk/getvdisk`
            this.$http.get(apiurl).then(response => {
				if (response.data.err === null ) {
                    this.comment(response.data.res)
                    this.countdisk = response.data.res.length
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
    background-color: #e3e3e3;
	font-weight: bold;
    color: black;
    text-align: center;
}

</style>

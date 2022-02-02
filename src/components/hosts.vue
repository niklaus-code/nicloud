<template>
		<div class="col-sm-12" style="margin-top:20px;padding-right:0; padding-left:0">
			<button class="btn btn-success btn-sm" @click="createhost" type="button"  style="margin-bottom:20px; margin-left:3px">创建</button>
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
						<th>操作</th>
      				</tr>
    			</thead>

				<tbody v-for="(item, index) in data">
      				<tr class="table-dark text-dark" :id="item.Uuid">
        				<label class="checkbox-inline" style="width:10px">
            				<input type="checkbox" v-model="item.Checkout">
        				</label>
        				<td>{{item.Datacenter}}</td>
        				<td>{{item.Vlan}}</td>
        				<td>{{item.Ipv4}}</td>
        				<td>{{item.Usedcpu}}核/{{item.Cpu}}核</td>
        				<td>{{item.Usedmem}}G/{{item.Mem}}G</td>
        				<td>{{item.count}}/{{item.Max_vms}}</td>
                        <td>
                            <span v-if='item.flag2' @click="c(index)">
                                {{item.Comment}}
                            </span>
                        <li v-if='item.flag'><span class="glyphicon glyphicon-calendar" @click="edit(index)"></span></li>
                            <div v-if='item.flag1'>
                                <div><input type="text" v-model="comments"></div>
                                <div><span  @click="input(index, item.Ipv4)" class="glyphicon glyphicon-calendar"></span></div>
                            </div>
                        </td>

		    			<td>
							<button class="btn btn-danger btn-xs" type="button" @click="deletehost(item.Ipv4, index)">
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
			cpu: "",
			mem: "",
			ip: "",
			num: "",
            comments: "",
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
			   if (response.data.err === null) {
                    alert("删除成功")
					this.data[index].Status=0
                    } else {
                    alert(response.data.err.Message)  
                }
            })
        },

        c: function (index) {
            this.data[index].flag2 = false
            this.data[index].flag1 = true
            this.comments = this.data[index].Comment
            },

        edit: function (index) {
            this.data[index].flag = false
            this.data[index].flag1 = true
            },

        input: function (index, ip) {
            var apiurl = `/api/hosts/addcomment`
            this.$http.get(apiurl, { params: { ip: ip, comment: this.comments} } ).then(response => {
                if (response.data) {
                    this.data[index].Comment = this.comments
                    }
            })
            this.data[index].flag = false
            this.data[index].flag1 = false
            this.data[index].flag2 = true
            },


        getvmstatus: function (uuid, host) {
            var apiurl = `/api/vm/getstatus`
            return this.$http.get(apiurl, { params: { uuid: uuid, host: host} } ).then(response => {
                return response.data.res
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
                for (let v in this.data) {
                    var r = this.getvmstatus(this.data[v].Uuid, this.data[v].Host)
                    r.then(value => {
                        this.data[v].Status = value
                        },
                    )}
            },

		gethost: function (ip) {
            var apiurl = `/api/hosts/gethosts`
            this.$http.get(apiurl).then(response => {
            if (response.data.err === null ) {
                 this.comment(response.data.res)
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

<template>
	<div>
		<div class="btn-group col-md-2 col-md-offset-10" >
			<input class="col-md-5" type="text" id="name" placeholder="" v-model="content">
			<button class="btn btn-default btn-sm" style="margin-right:5px" @click="search()">
				 <span class="glyphicon glyphicon-search"></span>筛选
			</button>
		</div>
		<div>
		<table class="table table-hover" style="text-align: center;">
    		<thead>
      			<tr>
					<th>
						<label class="checkbox-inline" style="border:red 1px">
  							<input type="checkbox" v-model="checkvalue" @click="checkbox()"> 
						</label>
					</th>
        			<th>实例名称</th>
        			<th>镜像</th>
        			<th>IP地址</th>
        			<th>CPU/内存</th>
        			<th>所属者</th>
        			<th>状态</th>
        			<th>备注</th>
        			<th>操作</th>
      			</tr>
    		</thead>
    		<tbody v-for="(item, index) in data">
      			<tr class="table-dark text-dark" :id="item.Uuid">
					<label class="checkbox-inline">
  						<input type="checkbox" v-model="item.Checkout"> 
					</label>
        		<td>{{item.Uuid}}</td>
        		<td>{{item.Os}}</td>
        		<td>{{item.Ip}}</td>
        		<td>{{item.Cpu}}核 / {{item.Mem}}G</td>
        		<td>{{item.Owner}}</td>
				<td>
        			<span>{{item.Status}}</span>
        		</td>
				<td>
                		{{item.Comment}}
				</td>
        		<td class="dropdown">
					<button class="btn btn-success btn-xs" @click="mount(item.Uuid, item.Ip, item.Host,item.Datacenter)" type="button">
						挂载
					</button>
					<button @click="shutdown(item.Uuid, index, item.Host)" class="btn btn-warning btn-xs" type="button">
						关机
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
            vdiskid: "",
            storage: "",
            pool: "",
			vmleft: "vmleft",
			active: "",
			content: "",
			stat: {
				"运行": "btn btn-success btn-xs", 
				"关机": "btn btn-warning btn-xs", 
				},
			statclass: "btn btn-danger",
            data: "",
			status: {
				0: "关机",
				1: "运行",
				2: "已删除",
			},
        }
    },

	created: function () {
        this.vdiskinfo()
    },

    mounted: function () {
		this.getvm()
    },

    methods: {
		mount: function (uuid, ip, host, datacenter ) {
         	var apiurl = `/api/vdisk/mountdisk`
            this.$http.get(apiurl, { params: { vmid: uuid, ip: ip, host: host, storage: sessionStorage.getItem('storage'), pool: sessionStorage.getItem('pool') , datacenter: datacenter, vdiskid: sessionStorage.getItem('vdiskid')}} ).then(response => {

	    	if (response.data.err === null) {
            	alert("挂载成功")
				this.$emit("toParent", "disk");
            } else {
                 alert("创建失败('"+response.data.err.Message+"')")  
               }
			})
		},

	    vdiskinfo: function () {
			var vdiskid = this.$store.state.vdisk.vdiskid
			if (vdiskid === null || typeof vdiskid === 'undefined' || vdiskid === '' ) {
                this.vdiskid = sessionStorage.getItem('vdiskid')
                this.storage = sessionStorage.getItem('storage')
                this.pool = sessionStorage.getItem('pool')
				} else {
				this.vdiskid = this.$store.state.vdisk.vdiskid
				this.storage = this.$store.state.vdisk.storage
				this.pool = this.$store.state.vdisk.pool
				sessionStorage.setItem('vdiskid', this.$store.state.vdisk.vdiskid)
				sessionStorage.setItem('storage', this.$store.state.vdisk.storage)
				sessionStorage.setItem('pool', this.$store.state.vdisk.pool)
				}
            },

        search: function (content) {
            var apiurl = `/api/vm/search`
            this.$http.get(apiurl, { params: { content: this.content} } ).then(response => {
            	this.data = response.data.res
            })
		},

		shutdown: function (uuid, index, host) {
            var apiurl = `/api/vm/operation/0`
            this.$http.get(apiurl, { params: { uuid: uuid, host: host } }).then(response => {
                if (response.data.err === null) {
                    this.data[index].Status = response.data.res.Status
                    if (response.data.res.Comment.length > 0) {
                        this.data[index].flag2 = true
                        }
                    if (response.data.res.Comment.length == 0) {
                        this.data[index].flag = true
                        }
                    } else {
                        alert("关机错误（'"+response.data.err.Message+"'）")
                    }
            })
        },

		getvmstatus: function (uuid, host) {
            var apiurl = `/api/vm/getstatus`
            return this.$http.get(apiurl, { params: { uuid: uuid, host: host} } ).then(response => {
            	return response.data.res
            	})
			},
	
        getvm: function () {
            var apiurl = `/api/vm/getvm`
            this.$http.get(apiurl).then(response => {
            var d = new Array()
            for (var v in response.data.res) {
                if (response.data.res[v]["Comment"].length > 0) {
                    response.data.res[v]["flag"] = false
                    response.data.res[v]["flag2"] = true
                    } else {
                    	response.data.res[v]["flag2"] = false
                        response.data.res[v]["flag"] = true
                    }
                response.data.res[v]["flag1"] = false
                d.push(response.data.res[v])
                }

			this.data = d
			for (let v in this.data) {
				var r = this.getvmstatus(this.data[v].Uuid, this.data[v].Host)
				r.then(value => {
					this.data[v].Status = value
					},
				)}
            })
        },
    }
  }
</script>


<style scoped>

.checkbox-inline {
	margin-bottom: 30px;
}

input{
	margin-right: 5px;
	border-color: #adadad;
	height: 30px;
}

.modal {
  display: block;
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

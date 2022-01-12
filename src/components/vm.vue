<template>
<div>
        <div class="col-md-7">
            <h5>云主机列表({{vmcount}})<h5>
        </div>
        <div class="col-md-5" style="float: right; margin-bottom:20px">
		    <button class="btn btn-default btn-sm" @click="create()" style="float: right">
		        <span class="glyphicon glyphicon-cog"></span>创建实例
		    </button>
		    <button class="btn btn-default btn-sm" style="margin-right:5px" @click="search()" style="float: right">
	            <span class="glyphicon glyphicon-search"></span>筛选
		    </button>
		    <input class="col-md-6" type="text" id="name" placeholder="" v-model="content" style="float: right">
        </div>
	<div style="margin-top:10px">
	<table class="table table-hover" style="text-align: center;">
    	<thead>
    		<tr>
				<th>
					<label class="checkbox-inline" style="border:red 1px">
  						<input type="checkbox" v-model="checkvalue" @click="checkbox()"> 
					</label>
				</th>
        		<th>实例名称</th>
        		<th>IP地址</th>
        		<th>镜像</th>
        		<th>所属宿主机</th>
        		<th>CPU/内存</th>
        		<th>云盘</th>
        		<th>所属者</th>
        		<th>备注</th>
        		<th>状态</th>
        		<th>操作</th>
      		</tr>
    	</thead>
    	<tbody v-for="(item, index) in data" v-if="item.Exist">
      		<tr class="table-dark text-dark" :id="item.Uuid">
				<label class="checkbox-inline">
  					<input type="checkbox" v-model="item.Checkout"> 
				</label>
       			<td width="300px">{{item.Name}}</td>
       			<td>{{item.Ip}}</td>
       			<td>{{item.Os}}</td>
       			<td>{{item.Host}}</td>
       			<td  style="min-width:80px">{{item.Cpu}}核 / {{item.Mem}}G</td>
       			<td>
					<ul>
						<li v-for="(k, v) in item.disk">
							{{k.Diskname}}&nbsp{{k.contain}}G
						</li>
					</ul>
				</td>
       			<td>{{item.Owner}}</td>
				<td>
	    			<span v-if='item.flag2' @click="c(index)">
               			{{item.Comment}}
           			</span>
					<li v-if='item.flag'><span class="glyphicon glyphicon-calendar" @click="edit(index)"></span></li>
						<div v-if='item.flag1'>
							<div><input type="text" v-model="comments"></div>
							<div><span  @click="input(index, item.Uuid)" class="glyphicon glyphicon-calendar"></span></div>
						</div>
				</td>
				<td>
					<button  v-if="item.Status === '运行'" type="button" class="btn btn-success btn-xs">{{item.Status}}</button>
       				<button v-else type="button" class="btn btn-warning btn-xs">{{item.Status}}</button>
       			</td>
       			<td :class="[{'dropup': (index > 10)}, {'dropdown': (index <= 10)}]" style="min-width:90px">
					<button class="btn btn-info btn-xs dropdown-toggle" type="button" id="menu1" data-toggle="dropdown">
						操作<span class="caret"></span>
					</button>
					<ul class="dropdown-menu" role="menu" aria-labelledby="menu1">
      					<li @click="start(item.Uuid, index, item.Host)" style="background-color: green;" role="presentation"><a role="menuitem" tabindex="-1">开机</a></li>
      					<li @click="reboot(item.Uuid, index, item.Host)" style="background-color: green; border-bottom: 1px white solid" role="presentation"><a role="menuitem" tabindex="-1">重启</a></li>
      					<li @click="pause(item.Uuid, index, item.Host)" style="background-color: #D2B48C;" role="presentation"><a role="menuitem" tabindex="-1">暂停</a></li>
      					<li @click="shutdown(item.Uuid, index, item.Host)" style="background-color: #D2B48C;" role="presentation"><a role="menuitem" tabindex="-1">关机</a></li>
      					<li @click="destroy(item.Uuid, index, item.Host)" style="background-color: #D2B48C; border-bottom: 1px white solid"  role="presentation"><a role="menuitem" tabindex="-1">强制断电</a></li>
      					<li style="background-color: rgb(255, 211, 0)"  role="presentation">
							<a @click="migrate(item.Uuid, item.Host, item.Cpu, item.Mem, item.Os, item.Owner, item.Ip, item.Comment)" role="menuitem" tabindex="-1">迁移</a>
						</li>
      					<li style="background-color: rgb(255, 211, 0)"  role="presentation">
							<a @click="migratelive(item.Uuid, item.Host, item.Cpu, item.Mem, item.Os, item.Owner, item.Ip, item.Comment)" style="border-bottom: 1px white solid" role="menuitem" tabindex="-1">热迁移</a>
						</li>
                        <li style="background-color: #C0C0C0;"  role="presentation">
                            <a @click="changeparam(item.Uuid, item.Ip, item.Os, item.Host, item.Cpu, item.Mem, item.Owner, item.Comment)" role="menuitem" tabindex="-1">修改配置</a>
                        </li>
      					<li @click="createsnap(item.Uuid, item.Ip,  item.Os, item.Host, item.Datacenter, item.Storage, item.Owner, item.Comment)" style="background-color: #C0C0C0; border-bottom: 1px white solid" role="presentation"><a role="menuitem" tabindex="-1">创建 & 恢复快照</a></li>
      					<li @click="restore(item.Uuid, item.Ip,  item.Os, item.Host, item.Datacenter, item.Storage, item.Owner, item.Comment)" style="background-color: #CD5C5C" role="presentation"><a role="menuitem" tabindex="-1">重置镜像</a></li>
      					<li @click="deletevm(item.Uuid, item.Datacenter, item.Storage,  index)" style="background-color: #CD5C5C; border-bottom: 1px white solid" role="presentation"><a role="menuitem" tabindex="-1">删除</a></li>
    				</ul>
					<button type="button" class="btn btn-info btn-xs" @click="vnc(item.vncid)"> <span class="glyphicon glyphicon-facetime-video"></span></button>
				</td>
      		</tr>
    	</tbody>
	</table>
    <div class="btn-group col-md-6" style="margin-top:20px">
        <ul class="pagination">
            <li><a @click="down()">&laquo;</a></li>
            <li  v-for="(item, index) in totalpagenumber"><a @click="getvm(item)">{{item}}</a></li>
            <li><a @click="up()">&raquo;</a></li>
        </ul>
	</div>
</div>
</template>
<script>

export default {
    data () {
        return {
            vmcount: 0,
            comments: "",
            totalpagenumber: "pagenumber",
            pagenumber: 1,
            dropup: "dropup",
            dropdown: "dropdown",
			active: "",
			checkvalue: false,
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
			}
        }
    },

    mounted: function () {
		this.getvm(1)
    },

    methods: {
        createsnap: function (uuid, ip, os, host,datacenter, storage , owner, comment) {
            this.$emit("toParent", "createsnap");
			this.$store.state.changeparam.uuid = uuid
			this.$store.state.changeparam.ip = ip
			this.$store.state.changeparam.os = os
			this.$store.state.changeparam.datacenter = datacenter
			this.$store.state.changeparam.storage = storage
			this.$store.state.changeparam.host = host
			this.$store.state.changeparam.owner = owner
			this.$store.state.changeparam.comment = comment
            },

        snap: function (uuid, ip, os, host,datacenter, storage , owner, comment) {
            this.$emit("toParent", "snap");
			this.$store.state.changeparam.uuid = uuid
			this.$store.state.changeparam.ip = ip
			this.$store.state.changeparam.os = os
			this.$store.state.changeparam.datacenter = datacenter
			this.$store.state.changeparam.storage = storage
			this.$store.state.changeparam.host = host
			this.$store.state.changeparam.owner = owner
			this.$store.state.changeparam.comment = comment
            },

        changeparam: function (uuid, ip, os, host, cpu, mem, owner, comment) {
            this.$emit("toParent", "changeparam");
			this.$store.state.changeparam.uuid = uuid
			this.$store.state.changeparam.ip = ip
			this.$store.state.changeparam.os = os
			this.$store.state.changeparam.host = host
			this.$store.state.changeparam.cpu = cpu
			this.$store.state.changeparam.mem = mem
			this.$store.state.changeparam.owner = owner
			this.$store.state.changeparam.comment = comment
            },

        migratelive: function (uuid, host, cpu, mem, os, owner, ip, comment) {
            this.$emit("toParent", "migratevmlive");
			this.$store.state.vm.uuid = uuid
			this.$store.state.vm.host = host
			this.$store.state.vm.cpu = cpu
			this.$store.state.vm.mem = mem
			this.$store.state.vm.os = os
			this.$store.state.vm.owner = owner
			this.$store.state.vm.ip = ip
			this.$store.state.vm.comment = comment
            },

       	migrate: function (uuid, host, cpu, mem, os, owner, ip, comment) {
            this.$emit("toParent", "migratevm");
			this.$store.state.vm.uuid = uuid
			this.$store.state.vm.host = host
			this.$store.state.vm.cpu = cpu
			this.$store.state.vm.mem = mem
			this.$store.state.vm.os = os
			this.$store.state.vm.owner = owner
			this.$store.state.vm.ip = ip
			this.$store.state.vm.comment = comment
            },

        create: function () {
                this.$emit("toParent", "createvm");
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

		input: function (index, uuid) {
            var apiurl = `/api/vm/addcomment`
            this.$http.get(apiurl, { params: { uuid: uuid, comment: this.comments} } ).then(response => {
                if (response.data) {
                    this.data[index].Comment = this.comments
                    }
            })
			this.data[index].flag = false
            this.data[index].flag1 = false
            this.data[index].flag2 = true
			},

		checkbox: function () {
			if (this.checkvalue) {
				for (var k in this.data) {
					this.checkvalue = true
					this.data[k].Checkout = false
					}
				} else {
				this.checkvalue = false
				for (var k in this.data) {
					this.data[k].Checkout = true
					}
				}
			},

        /*
        migratelive: function (content) {
            var apiurl = `/api/vm/migratelive`
            this.$http.get(apiurl, { params: { content: this.content} } ).then(response => {
                if (response.data.res === null) {
                    alert("未查询到")
                    } else {
            	    this.comment(response.data.res)
                    }
            })
		},
        */

        search: function (content) {
            var apiurl = `/api/vm/search`
            this.$http.get(apiurl, { params: { content: this.content} } ).then(response => {
                if (response.data.res === null) {
                    alert("未查询到")
                    } else {
            	    this.comment(response.data.res)
                    }
            })
		},

        vnc: function (vncid) {
            var apiurl = `/api/vm/vnc`
			window.open("http://10.0.85.90:8787/vnc.html?path=websockify/?vncid="+vncid, '_blank');
		},

		getvmstatus: function (uuid, host) {
            var apiurl = `/api/vm/getstatus`
            return this.$http.get(apiurl, { params: { uuid: uuid, host: host} } ).then(response => {
            	return response.data.res
            	})
			},

        up: function() {
            if (Number(this.pagenumber) < Number(this.totalpagenumber)) {
                this.getvm(Number(this.pagenumber)+1)
                }
            },

        down: function() {
            if (Number(this.pagenumber) > 1 ) {
                this.getvm(Number(this.pagenumber)-1)
                }
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
	
        getvm: function (start) {
            var apiurl = `/api/vm/getvm`
            this.$http.get(apiurl, { params: { start: start} }).then(response => {
			if (response.data.err === null ) {
                this.totalpagenumber = response.data.pagenumber
                this.vmcount = response.data.vmcount
                this.pagenumber = start
                this.comment(response.data.res)
				} else {
					alert(response.data.err)
					this.$router.push({name:"login"})
					} 
            })   
        },

        deletevm: function (uuid, datacenter, storage, index) {
            var apiurl = `/api/vm/delete`
            this.$http.get(apiurl, { params: { uuid: uuid, datacenter:datacenter, storage: storage} }).then(response => {
				if (response.data.err == null) {
					alert("删除成功")
					//this.data[index].Exist=0	
				} else {	
					alert(response.data.err.Message)
				}
            })
        },
    
        restore: function (uuid, ip, os, host,datacenter, storage , owner, comment) {
            this.$emit("toParent", "restorevm");
			this.$store.state.changeparam.uuid = uuid
			this.$store.state.changeparam.ip = ip
			this.$store.state.changeparam.os = os
			this.$store.state.changeparam.datacenter = datacenter
			this.$store.state.changeparam.storage = storage
			this.$store.state.changeparam.host = host
			this.$store.state.changeparam.owner = owner
			this.$store.state.changeparam.comment = comment
            },

        shutdown: function (uuid, index, host) {
            var apiurl = `/api/vm/operation/0`
            this.$http.get(apiurl, { params: { uuid: uuid, host: host } }).then(response => {
                if (response.data.err === null) {
                    this.data[index].Status = "关机"
                    } else {
                        alert("关机错误（'"+response.data.err.Message+"'）")
                    }
            })
        },


        destroy: function (uuid, index, host) {
            var apiurl = `/api/vm/operation/1`
            this.$http.get(apiurl, { params: { uuid: uuid, host: host } }).then(response => {
				if (response.data.err === null) {
					this.data[index].Status = "关机"
					} else {
						alert("关机错误（'"+response.data.err.Message+"'）")
					}
            })
        },

        pause: function (uuid, index, host) {
            var apiurl = `/api/vm/operation/3`
			
            this.$http.get(apiurl, { params: { uuid: uuid, host: host } }).then(response => {
				if (response.data.err === null) {
					this.data[index].Status = "暂停"
					} else {
						alert("暂停错误（'"+response.data.err.Message+"'）")
					}
            })
        },

        reboot: function (uuid, index, host) {
            var apiurl = `/api/vm/operation/4`
			
            this.$http.get(apiurl, { params: { uuid: uuid, host: host } }).then(response => {
				if (response.data.err === null) {
					this.data[index].Status = "正在重启"
					} else {
						alert("暂停错误（'"+response.data.err.Message+"'）")
					}
            })
        },

        start: function (uuid, index, host) {
            var apiurl = `/api/vm/operation/2`
			
            this.$http.get(apiurl, { params: { uuid: uuid, host: host } }).then(response => {
				if (response.data.err === null) {
					this.data[index].Status = "运行"
					} else {
						alert("开机错误（'"+response.data.err.Message+"'）")
					}
            })
        },
    }
  }
</script>

<style scoped>
h5 {
    font-weight: 600;
}

.checkbox-inline {
	margin-bottom: 30px;
}

input{
	margin-right: 5px;
	border-color: #adadad;
	height: 30px;
    margin-top: 1px;
}

.modal {
  display: block;
}

table {
    margin-bottom: 0px;
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

.glyphicon {
	caret-color: rgba(0, 0, 0, 0)
}

.pagination {
    margin-top: 0;
    display: block;
}

.pagination li a {
    color: #000;
}
</style>

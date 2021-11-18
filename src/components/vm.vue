<template>
<div>
  	<div class="content whisper-content leacots-content details-content col-md-11 col-md-offset-2" style="background-color:white; float:left">
		<div class="btn-group col-md-3 col-md-offset-9" >
			<input class="col-md-5" type="text" id="name" placeholder="" v-model="content">
			<button class="btn btn-default btn-sm" style="margin-right:5px" @click="search()">
				 <span class="glyphicon glyphicon-search"></span>筛选
			</button>
			<button class="btn btn-default btn-sm" @click="create()">
				 <span class="glyphicon glyphicon-cog"></span>创建实例
			</button>
		</div>
		<div style="margin-top:40px">
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
    		<tbody v-for="(item, index) in data">
      			<tr v-if="item.Exist" class="table-dark text-dark" :id="item.Uuid">
					<label class="checkbox-inline">
  						<input type="checkbox" v-model="item.Checkout"> 
					</label>
        			<td>{{item.Name}}</td>
        			<td>{{item.Ip}}</td>
        			<td>{{item.Os}}</td>
        			<td>{{item.Host}}</td>
        			<td>{{item.Cpu}}核 / {{item.Mem}}G</td>
        			<td>
						<ul>
							<li v-for="(k, v) in item.disk">
								{{k.Diskname}}&nbsp{{k.Contain}}G
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
								<div><input type="text" v-model="comment"></div>
								<div><span  @click="input(index, item.Uuid)" class="glyphicon glyphicon-calendar"></span></div>
							</div>
					</td>
					<td>
						<button  v-if="item.Status === '运行'" type="button" class="btn btn-success btn-xs">{{item.Status}}</button>
        				<button v-else type="button" class="btn btn-warning btn-xs">{{item.Status}}</button>
        			</td>
        			<td class="dropdown">
						<button class="btn btn-info btn-xs dropdown-toggle" type="button" id="menu1" data-toggle="dropdown">
							操作<span class="caret"></span>
						</button>
						<ul class="dropdown-menu" role="menu" aria-labelledby="menu1" style="">
      						<li @click="start(item.Uuid, index, item.Host)" style="background-color: green" role="presentation"><a role="menuitem" tabindex="-1">开机</a></li>
      						<li @click="shutdown(item.Uuid, index, item.Host)" style="background-color: #e56b6b"  role="presentation"><a role="menuitem" tabindex="-1">关机</a></li>
      						<li @click="pause(item.Uuid, index, item.Host)" style="background-color: rgb(255, 211, 0)" role="presentation"><a role="menuitem" tabindex="-1">暂停</a></li>
      						<li style="background-color: greenyellow"  role="presentation">
								<a @click="migrate(item.Uuid, item.Host)" role="menuitem" tabindex="-1">迁移</a>
							</li>
      						<li @click="deletevm(item.Uuid, index)" style="background-color: #808080" role="presentation"><a role="menuitem" tabindex="-1">删除</a></li>
    					</ul>
						<button type="button" class="btn btn-info btn-xs" @click="vnc(item.Uuid, item.Host)"> <span class="glyphicon glyphicon-facetime-video"></span></button>
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
		this.getvm()
    },

    methods: {
       	migrate: function (uuid, host) {
            this.$emit("toParent", "migratevm");
			this.$store.state.vm.uuid = uuid
			this.$store.state.vm.host = host
            },

        create: function () {
                this.$emit("toParent", "createvm");
                },

		c: function (index) {
			this.data[index].flag2 = false
			this.data[index].flag1 = true
			this.comment = this.data[index].Comment
			},

		edit: function (index) {
			this.data[index].flag = false
            this.data[index].flag1 = true
			},

		input: function (index, uuid) {
            var apiurl = `/api/vm/addcomment`
            this.$http.get(apiurl, { params: { uuid: uuid, comment: this.comment} } ).then(response => {
                if (response.data) {
                    this.data[index].Comment = this.comment
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
        search: function (content) {
            var apiurl = `/api/vm/search`
            this.$http.get(apiurl, { params: { content: this.content} } ).then(response => {
            	this.data = response.data.res
            })
		},

        vnc: function (uuid, host) {
            var apiurl = `/api/vm/vnc`
            this.$http.get(apiurl, { params: { uuid: uuid, host: host} } ).then(response => {
				var	url = response.data
				window.open(url, '_blank');
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

        deletevm: function (uuid, index) {
            var apiurl = `/api/vm/delete`
            this.$http.get(apiurl, { params: { uuid: uuid} }).then(response => {
				if (response.data.err == null) {
					alert("删除成功")
					this.data[index].Exist=0	
				} else {	
					alert(response.data.err.Message)
				}
            })
        },

        shutdown: function (uuid, index, host) {
            var apiurl = `/api/vm/operation/0`
            this.$http.get(apiurl, { params: { uuid: uuid, host: host } }).then(response => {
				if (response.data.err == null) {
					this.data[index].Status = response.data.res.Status
					//this.$set(this.data, index, response.data.res)
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

        pause: function (uuid, index, host) {
            var apiurl = `/api/vm/operation/3`
			
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
						alert("暂停错误（'"+response.data.err.Message+"'）")
					}
            })
        },

        start: function (uuid, index, host) {
            var apiurl = `/api/vm/operation/1`
			
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
						alert("开机错误（'"+response.data.err.Message+"'）")
					}
            })
        },
    }
  }
</script>


<style scoped>
.content {
	box-shadow: 0 0 10px rgba(0,0,0,8);
  	border-radius: 10px/10px;
  	z-index: -1;
	padding: 20px 0px 0px 0px;
	margin-left: 0px;
	margin-top: 50px;
}

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

.dropdown-menu {
	top: 40px;
}

.glyphicon {
	caret-color: rgba(0, 0, 0, 0)
}
</style>

<template>
<div>
    <nicloudhead></nicloudhead>

	<component v-bind:is="vmleft"></component>
  	<div class="content whisper-content leacots-content details-content col-md-11 col-md-offset-2" style="background-color:white; float:left">
		<div class="btn-group col-md-2 col-md-offset-10" >
			<input class="col-md-5" type="text" id="name" placeholder="" v-model="content">
			<button class="btn btn-default btn-sm" style="margin-right:5px" @click="search()">
				 <span class="glyphicon glyphicon-search"></span>筛选
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
        			<th>镜像</th>
        			<th>IP地址</th>
        			<th>CPU/内存</th>
        			<th>所属者</th>
        			<th>存储集群</th>
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
        		<td>{{item.Storage}}</td>
				<td>
					<button  v-if="item.Status === '运行'" type="button" class="btn btn-success btn-xs">{{item.Status}}</button>
        			<button v-else type="button" class="btn btn-warning btn-xs">{{item.Status}}</button>
        		</td>
				<td>
		    		<span v-if='item.flag2' @click="c(index)">
                		{{item.Comment}}
            		</span>
					<li v-if='item.flag'><span class="glyphicon glyphicon-calendar" @click="edit(index)"></span></li>
					<div v-if='item.flag1'>
						<input type="text" v-model="comment">
						<span  @click="input(index, item.Uuid)" class="glyphicon glyphicon-calendar"></span>
					</div>
				</td>
        		<td class="dropdown">
					<button class="btn btn-info btn-xs" @click="mount(item.Uuid, item.Ip, item.Host, item.Storage, item.Datacenter)" type="button">
						挂载
					</button>
					<button class="btn btn-info btn-xs" type="button">
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

import foot from '@/components/footer'
import nicloudhead from '@/components/nicloudhead'
import vmleft from '@/components/vmleft'

export default {
    data () {
        return {
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
			cloudriveid: "",
			pool: "",
			
        }
    },

    components: {
        foot, nicloudhead, vmleft
    },

	created: function () {
        this.cloudriveinfo()
    },

    mounted: function () {
		this.getvm()
    },

    methods: {
		mount: function (uuid, ip, host, storage, datacenter ) {
         	var apiurl = `/api/vm/mountdisk`
            this.$http.get(apiurl, { params: { vmid: uuid, ip: ip, host: host, storage: storage, pool: this.pool, datacenter: datacenter, cloudriveid: this.cloudriveid}} ).then(response => {

	    	if (response.data.err === null) {
            	alert("挂载成功")
            } else {
                 alert("创建失败('"+response.data.err.Message+"')")  
               }
			})
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

	    cloudriveinfo: function () {
            this.cloudriveid = this.$route.query.cloudriveid
            this.storage = this.$route.query.storage
            this.pool = this.$route.query.pool
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

        search: function (content) {
            var apiurl = `/api/vm/search`
            this.$http.get(apiurl, { params: { content: this.content} } ).then(response => {
            	this.data = response.data.res
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
.content {
	box-shadow: 0 0 10px rgba(0,0,0,8);
  	border-radius: 10px/10px;
  	z-index: -1;
	padding: 20px 0px 0px 0px;
	margin-left: 0px;
	margin-TOP: 50px;
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
</style>

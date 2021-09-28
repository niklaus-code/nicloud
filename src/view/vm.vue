<template>
<div>
    <nicloudhead></nicloudhead>

  <div class="content whisper-content leacots-content details-content">
	
	<div class="btn-group col-md-3 col-md-offset-9" >
			<input class="col-md-5" type="text" id="name" placeholder="" v-model="content">
			<button class="btn btn-default btn-sm" style="margin-right:5px" @click="search()">
				 <span class="glyphicon glyphicon-search"></span>筛选
			</button>
		<router-link :to="{name:'createvm'}">
			<button class="btn btn-default btn-sm">
				 <span class="glyphicon glyphicon-cog"></span>创建实例
			</button>
		</router-link>
	</div>
<div style="margin-top:80px">
	<table class="table table-striped" style="text-align: center;">
    <thead>
      <tr>
		<th>
		<label class="checkbox-inline" style="border:red 1px">
  			<input type="checkbox" v-model="checkvalue" @click="checkbox()"> 
		</label>
		</th>
        <th>实例名称</th>
        <th>IP地址</th>
        <th>所属宿主机</th>
        <th>CPU</th>
        <th>内存</th>
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
        <td>{{item.Ip}}</td>
        <td>{{item.Host}}</td>
        <td>{{item.Cpu}}</td>
        <td>{{item.Mem}}</td>
        <td>{{item.Owner}}</td>
        <td><button type="button" :class=stat[item.Status]>{{item.Status}}</button></td>
        <td>{{item.Comment}}</td>
        <td class="default">
			<button type="button" class="btn btn-info btn-sm" @click="start(item.Uuid, index, item.Host)">开机</button>
			<button type="button" class="btn btn-info btn-sm" @click="shutdown(item.Uuid, index, item.Host)">关机</button>
			<button type="button" class="btn btn-info btn-sm" @click="deletevm(item.Uuid, item.Ip, item.Host)">删除</button>
			<button type="button" class="btn btn-info btn-sm" @click="vnc(item.Uuid, item.Host)"> <span class="glyphicon glyphicon-facetime-video"></span></button>
		</td>
      </tr>
    </tbody>
	</table>
	<div>
  </div>
	<foot></foot>
  </div>
</template>
<script>

import foot from '@/components/footer'
import nicloudhead from '@/components/nicloudhead'

export default {
    data () {
        return {
			a : "",
			checkvalue: false,
			content: "",
			stat: {
				"运行": "btn btn-success btn-sm", 
				"关机": "btn btn-warning btn-sm", 
				},
			statclass: "btn btn-danger",
            data: [],
			status: {
				0: "关机",
				1: "运行",
				2: "已删除",
			}
        }
    },

    components: {
        foot, nicloudhead
    },

    mounted: function () {
		this.getvm()
        var id = this.$route.params.id
        if (!id) {
            this.blogid = window.sessionStorage.getItem('blogid')
        } else {
            this.blogid = id
            window.sessionStorage.setItem('blogid', this.blogid)
        }
    },

    methods: {
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
				this.data = response.data.res
				for (let v in this.data) {
					var r = this.getvmstatus(this.data[v].Uuid, this.data[v].Host)
					r.then(value => {
						this.data[v].Status = value
						},
					)
				}
            })
        },

        deletevm: function (uuid, ip, host) {
            var apiurl = `/api/vm/delete`
            this.$http.get(apiurl, { params: { uuid: uuid, ip: ip, host: host} }).then(response => {
				if (response.data.err == null) {
            		this.data = response.data.res
					} else {
						alert(response.data.err.Message)
					}
            })
        },

        shutdown: function (uuid, index, host) {
            var apiurl = `/api/vm/operation/0`
            this.$http.get(apiurl, { params: { uuid: uuid, host: host } }).then(response => {
				if (response.data.err == null) {
					this.$set(this.data, index, response.data.res)
					}
				alert(response.data.err.Message)
            })
        },

        start: function (uuid, index, host) {
            var apiurl = `/api/vm/operation/1`
			
            this.$http.get(apiurl, { params: { uuid: uuid, host: host } }).then(response => {
				if (response.data.err == null) {
					this.$set(this.data, index, response.data.res)
					}
				alert(response.data.err.Message)
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

.layui-table th {
	font-weight: bold;
	color: black;
}

.layui-table td {
	display: table-cell;
	vertical-align: "middle";
}

th {
	text-align: center;
}
</style>

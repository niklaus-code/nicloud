<template>
<div>
    <headd></headd>

  <div class="content whisper-content leacots-content details-content">
	<router-link :to="{name:'createvm'}">
	<button >创建</button>
	</router-link>
<table class="layui-table" lay-even lay-skin="line" lay-size="lg" style="text-align: center;"  style="text-align: center;">
    <thead>
      <tr>
        <th>uuid</th>
        <th>IP</th>
        <th>宿主机</th>
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
        <td>{{item.Uuid}}</td>
        <td>{{item.Ip}}</td>
        <td>{{item.Host}}</td>
        <td>{{item.Cpu}}</td>
        <td>{{item.Mem}}</td>
        <td>{{item.Owner}}</td>
        <td><button type="button" :class=stat[item.Status]>{{item.Status}}</button></td>
        <td>{{item.Comment}}</td>
        <td class="default">
			<button type="button" class="btn btn-info btn-sm" @click="start(item.Uuid, index, item.host)">开机</button>
			<button type="button" class="btn btn-info btn-sm" @click="shutdown(item.Uuid, index, item.host)">关机</button>
			<button type="button" class="btn btn-info btn-sm" @click="deletevm(item.Uuid, item.Ip, item.Host)">删除</button>
		</td>
      </tr>
    </tbody>
</table>

  </div>
	<foot></foot>
  </div>
</template>
<script>

import foot from '@/components/footer'
import headd from '@/components/head'

export default {
    data () {
        return {
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
        foot, headd
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
        getvm: function () {
            var apiurl = `/api/vm/getvm`
            this.$http.get(apiurl).then(response => {
            	this.data = response.data.res
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
.modal {
  display: block;
}

.layui-table th {
	padding:10px 30px;
}

.layui-table td {
	padding:10px 30px;
}

th {
	text-align: center;
}
</style>

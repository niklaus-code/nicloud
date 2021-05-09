<template>
<div>
    <headd></headd>

  <div class="content whisper-content leacots-content details-content">
	<router-link :to="{name:'createvm'}">
	<button >创建</button>
	</router-link>
<table class="layui-table" lay-even lay-skin="line" lay-size="lg">
    <thead>
      <tr>
        <th>uuid</th>
        <th>名称</th>
        <th>cpu</th>
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
        <td>{{item.Name}}</td>
        <td>{{item.Cpu}}</td>
        <td>{{item.Mem}}</td>
        <td>{{item.Owner}}</td>
        <td>{{item.Status}}</td>
        <td>{{item.Comment}}</td>
        <td class="default">
			<button type="button" class="btn btn-success" @click="start(item.Uuid,index)">开机</button>
			<button type="button" class="btn btn-info" @click="shutdown(item.Uuid, index)">关机</button>
			<button type="button" class="btn btn-info" @click="deletevm(item.Uuid, index)">删除</button>
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
            data: '',
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

        deletevm: function (uuid, index) {
            var apiurl = `/api/vm/delete`
            this.$http.get(apiurl, { params: { uuid: uuid} }).then(response => {
            	this.data = response.data.res
				//this.$set(this.data, index , response.data.res)
            })
        },

        shutdown: function (uuid, index) {
            var apiurl = `/api/vm/operation/0`
            this.$http.get(apiurl, { params: { uuid: uuid } }).then(response => {
				this.$set(this.data, index , response.data.res)
            })
        },

        start: function (uuid, index) {
            var apiurl = `/api/vm/operation/1`
			
            this.$http.get(apiurl, { params: { uuid: uuid } }).then(response => {
				if (response.data.err) {
					alert(response.data.err.Message)
					}
				this.$set(this.data, index , response.data.res)
            })
        },
    }
  }
</script>

<style>
	 .modal {
  display: block;
}
</style>

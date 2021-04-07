<template>
<div>
    <headd></headd>

  <div class="content whisper-content leacots-content details-content">
	<router-link :to="{name:'createvm'}">
	<button>创建</button>
	</router-link>
	
<table class="table" style="margin-top:20px">
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
        <td class="default">{{item.Uuid}}</td>
        <td class="info">{{item.Name}}</td>
        <td class="active">{{item.Cpu}}</td>
        <td class="success">{{item.Mem}}</td>
        <td class="warning">{{item.Owner}}</td>
        <td class="active">{{item.Status}}</td>
        <td class="info">{{item.Comment}}</td>
        <td class="default">
			<button type="button" class="btn btn-success" @click="start(index)">启动</button>
			<button type="button" class="btn btn-info" @click="shutdown(index)">停止</button>
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
				1: "运行"
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
            var apiurl = `/api/vm/getvmlist`
            this.$http.get(apiurl).then(response => {
            this.data = response.data.res
            })
        },
        shutdown: function (index) {
            var apiurl = `/api/vm/operation/0`
            this.$http.get(apiurl, { params: { uuid: "31a803b2-5f11-4f14-875f-d14347db13fb" } }).then(response => {
				this.$set(this.data, index , response.data.res)
            })
        },
        start: function (index) {
            var apiurl = `/api/vm/operation/1`
			
            this.$http.get(apiurl, { params: { uuid: "31a803b2-5f11-4f14-875f-d14347db13fb" } }).then(response => {
				this.$set(this.data, index , response.data.res)
            })
        },
    }
  }
</script>
<style>
    .details-content .article-cont p {
    padding:30px 0 0 5px
}
</style>

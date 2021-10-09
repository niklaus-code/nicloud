<template>
	<div>
		<nicloudhead> </nicloudhead>
		<div class="btn-group col-md-1 col-md-offset-11" >
			<strong>总数:<span>{{total}}<span></strong>
		</div>
		
		<div class="btn-group col-md-2 col-md-offset-10" style="margin-top:10px">
            <input class="col-md-5" type="text" id="name" placeholder="" v-model="content">
            <button class="btn btn-default btn-sm" style="margin-right:5px" @click="search()">
                 <span class="glyphicon glyphicon-search"></span>筛选
            </button>
        <router-link :to="{name:'createmachine'}">
            <button class="btn btn-default btn-sm">
                 <span class="glyphicon glyphicon-cog"></span>增加机器
            </button>
        </router-link>
    </div>
		<div class="machine">
			<table class="table table-bordered">
			   <thead>
      				<tr>
        				<th>资产名称</th>
        				<th>品牌</th>
        				<th>型号</th>
        				<th>原厂序列号</th>
        				<th>资产标签</th>
        				<th>单位</th>
        				<th>所属部门</th>
        				<th>责任部门</th>
        				<th>责任人</th>
        				<th>机房</th>
        				<th>机柜</th>
        				<th>机柜资产标签</th>
        				<th>机柜位置</th>
        				<th>高度</th>
        				<th>设备状态</th>
        				<th>额定功率</th>
        				<th>用电等级</th>
        				<th>管理IP</th>
        				<th>业务IP</th>
        				<th>备注</th>
        				<th>
							<span>操作</span>
						</th>
      				</tr>
    			</thead>
				<tbody v-for="(item, index) in data">
					<tr>
						<td>{{item.Zichanmingcheng}}</td>
						<td>{{item.Pinpai}}</td>
						<td>{{item.Xinghao}}</td>
						<td>{{item.Xuliehao}}</td>
						<td>{{item.Zichanbiaoqian}}</td>
						<td>{{item.Danwei}}</td>
						<td>{{item.Suoshubumen}}</td>
						<td>{{item.Zichanzerenbumen}}</td>
						<td>{{item.Zerenren}}</td>
						<td>{{item.Suoshujifang}}</td>
						<td>{{item.Jigui}}</td>
						<td>{{item.Jiguizichanbiaoqian}}</td>
						<td>{{item.Weizhi}}</td>
						<td>{{item.Gaodu}}</td>
						<td>{{item.Shebeizhuangtai}}</td>
						<td>{{item.Edinggonglv}}</td>
						<td>{{item.Yongdiandengji}}</td>
						<td>{{item.Guanliip}}</td>
						<td>{{item.Yewuip}}</td>
						<td>{{item.Beizhu}}</td>
						<td>
							<button type="button" class="btn btn-primary btn-xs">修改</button>
							<button type="button" @click="delmachine(item.Id)" class="btn btn-primary btn-xs">删除</button>
						</td>
					</tr>
				</tbody>
			</table>
		</div>
		<div class="btn-group col-md-2  col-md-offset-5" style="margin-top:10px; padding-left:0">
			<ul class="pagination">
    			<li><a>&laquo;</a></li>
    			<li v-for="(item, index) in allpage"><a @click="getmachinelist(item, 100)" >{{item}}</a></li>
    			<li><a>&raquo;</a></li>
			</ul>
		</div>
	</div>
</template>
<script>
import nicloudhead from '@/components/nicloudhead'

export default {
    data () {
        return {
			total: 0,
			allpage: 0,
			data: [],
			}
		},
    components: {
        nicloudhead
    },
	mounted: function () {
		this.getmachinelist(1, 100)
		this.getpagenumber()
		},

	methods: {
		getpagenumber: function () {
			var apiurl = `/api/machine/getpage`
			this.$http.get(apiurl).then(response => {
				this.allpage = response.data.pagenumber
				this.total = response.data.totalnumber
            })
			},

		getmachinelist: function (startpage, offset) {
			var apiurl = `/api/machine/getmachinelist`
			this.$http.get(apiurl, { params: { startpage: startpage, offset: offset }} ).then(response => {
                this.data = response.data.res
            })
			},

		delmachine: function (mid) {
			var apiurl = `/api/machine/delmachine`
			this.$http.get(apiurl, { params: { id: mid }} ).then(response => {
                this.data = response.data.res
            })
			},
		}
	}
</script>

<style>
input {
	margin-right: 5px;
	height: 29px;
	}

.pagination {
	margin: 0
}
.machine {
	background-color: white;
}
</style>

<template>
	<div>
        <div class="col-sm-12" style="margin-top: 30px; margin-bottom: 30px;">
           <strong style="margin-left: 20px">总机器数:<span>{{total}}<span></strong>
            <button  style="float: right;" class="btn btn-default btn-sm" @click="addserver">
                <span class="glyphicon glyphicon-cog"></span>增加机器
            </button>
            <button style="float: right;"  class="btn btn-default btn-sm" @click="search()">
                <span class="glyphicon glyphicon-search"></span>筛选
            </button>
            <input style="float: right;" class="col-md-5" type="text" id="name" placeholder="" v-model="content">
        </div>
		<div class="col-sm-12">
			<table class="table table-bordered">
			   <thead>
      				<tr>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp资产名称</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp品牌</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp型号</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp原厂序列号</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp资产标签</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp单位</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp所属部门</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp责任部门</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp责任人</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp机房</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp机柜</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp机柜资产标签</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp机柜位置</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp高度</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp设备状态</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp额定功率</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp用电等级</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp管理IP</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp业务IP</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp备注</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>
							<span>&nbsp操作</span>
						</th>
        				<th><span class="glyphicon glyphicon-edit"> </span>&nbsp状态</th>
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
						<td @dblclick="dblclick(index)">
							<div v-if="item.cm">
                				<input type="text" v-model="comment">
            				</div>
							<div v-else>
								{{item.Beizhu}}
							</div>
						</td>
						<td>
							<button  v-if="item.cm" type="button" class="btn btn-primary btn-xs" @click="save(item.Id, index)">保存</button>
							<button v-else type="button" @click="delmachine(item.Id, index)" class="btn btn-primary btn-xs">删除</button>
						</td>
						<td>
							<button  v-if="item.Status === '运行'" type="button" class="btn btn-success btn-xs">{{item.Status}}</button>
							<button v-else type="button" class="btn btn-warning btn-xs">{{item.Status}}</button>
						</td>
					</tr>
				</tbody>
			</table>
		</div>
		<div class="btn-group col-md-4  col-md-offset-5" style="margin-top:10px; padding-left:0">
			<ul class="pagination">
    			<li><a>&laquo;</a></li>
    			<li v-for="(item, index) in allpage"><a @click="getmachinelist(item, 50)" >{{item}}</a></li>
    			<li><a>&raquo;</a></li>
			</ul>
		</div>
	</div>
</template>
<script>

export default {
    data () {
        return {
			comment: "",
			total: 0,
			allpage: 0,
			onpage: 1,
			data: [],
			}
		},

	mounted: function () {
        var p = sessionStorage.getItem('serverpage')
        if (typeof p === 'undefined' || p === null || p === '') {
            p = 1
        }
		this.getmachinelist(p, 50)
		this.getpagenumber()
		this.refresh
		},

	methods: {
        refresh: function () {
            this.$emit("toParent", "server");
            },

        addserver: function () {
            this.$emit("toParent", "addserver");
            },

		save: function (id, index) {
			var apiurl = `/api/machine/update`
            this.$http.get(apiurl, { params: { id: id, c: this.comment } } ).then(response => {
                if (response.data.res) {
					this.data[index].cm = false
                    this.data[index].Beizhu = this.comment
                    }
                })
			},

		dblclick: function (index) {
			this.data[index].cm = true
			this.comment = this.data[index].Beizhu
			},

		search: function () {
			var apiurl = `/api/machine/search`
			this.$http.get(apiurl, { params: { content: this.content }} ).then(response => {
				this.data = response.data.res
            })
			},

		getpagenumber: function () {
			var apiurl = `/api/machine/getpage`
			this.$http.get(apiurl).then(response => {
				this.allpage = response.data.pagenumber
				this.total = response.data.totalnumber
            })
			},

		ping: function (ip) {
			var apiurl = `/api/machine/ping`
			return this.$http.get(apiurl, { params: { ip: ip }} ).then(response => {
				return response.data.res
            })
			},

		getmachinelist: function (startpage, offset) {
            sessionStorage.setItem('router', "server")
            sessionStorage.setItem('serverpage', startpage)
			this.onpage = startpage
			var apiurl = `/api/machine/getmachinelist`
			this.$http.get(apiurl, { params: { startpage: startpage, offset: offset }} ).then(response => {
				var d = new Array()
                for (var i in response.data.res) {
					response.data.res[i]["cm"] = false

					d.push(response.data.res[i])
					}

				this.data = d
				for (let b in this.data) {
					var r = this.ping(this.data[b].Yewuip)
					r.then(value => {
						this.data[b].Status = value
						},
					)}
            })
			},

		delmachine: function (mid, index) {
			this.data[index].Status = false
			var apiurl = `/api/machine/delmachine`
			this.$http.get(apiurl, { params: { id: mid, startpage: this.onpage, offset: 50}} ).then(response => {
                this.data = response.data.res
            })
			},
		}
	}
</script>

<style scoped>
.col-sm-12 {
    padding-right: 0;
    padding-left: 0;
    }

input {
	margin-right: 5px;
	height: 29px;
	}

.pagination {
	margin: 0
}

table {background-color: white}

.machine {
	background-color: white;
}
</style>

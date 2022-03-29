<template>
<div>
    <div class="col-md-7" style="padding-left: 24px">
        <h5>云主机归档列表({{vmcount}})<h5>
    </div>
    <div class="col-md-5" style="float: right; margin-bottom: 20px">
		<button class="btn btn-default btn-sm" style="margin-right: 5px" @click="search()" style="float: right">
	        <span class="glyphicon glyphicon-search"></span> 筛选
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
        	<!--	<th>实例名称</th>-->
        		<th>UUID</th>
        		<th>IP地址</th>
        		<th>镜像</th>
        		<th>宿主机</th>
        		<th>所属者</th>
        		<th>存储集群</th>
        		<th>数据中心</th>
        		<th>创建时间</th>
        		<th>归档时间</th>
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
       			<td>{{item.osname}}</td>
       			<td>{{item.Host}}</td>
       			<td>{{item.owner}}</td>
       			<td>{{item.Storage}}</td>
       			<td>{{item.Datacenter}}</td>
       			<td>{{item.Create_time}}</td>
       			<td>{{item.Archive_time}}</td>
				<td>{{item.Comment}}</td>
       			<td :class="[{'dropup': (index > 10)}, {'dropdown': (index <= 10)}]" style="min-width:90px">
					<button @click="delvmpermanent(item.Uuid, item.Storage)" class="btn btn-danger btn-xs dropdown-toggle" type="button" id="menu1" data-toggle="dropdown">
					<span class="glyphicon glyphicon-trash"></span> 永久删除  
					</button>
				</td>
      		</tr>
    	</tbody>
	</table>
    <div class="btn-group col-md-6" style="margin-top:20px; margin-bottom:30px">
        <ul class="pagination">
            <li><a @click="down()">&laquo;</a></li>
            <li v-for="(item, index) in totalpagenumber"><a @click="getvmarchive(item)">{{item}}</a></li>
            <li><a @click="up()">&raquo;</a></li>
        </ul>
	</div>
</div>
</template>
<script>

export default {
    data () {
        return {
            data: [],
            vmcount: 0,
            cpagenumber: 1,
			}
    },

    mounted: function () {
        let self = this
        document.onkeydown = function(e) {
            let ev = document.all ? window.event : e
            if (ev.keyCode === 13) {
                self.search()
            }
        }
        var p = sessionStorage.getItem('pagenumber')
        if (typeof p === 'undefined' || p === null || p === '') {
            p = 1
            this.cpagenumber = 1
        }
        this.cpagenumber = p

		this.getvmarchive(this.cpagenumber)
    },

    methods: {
        delvmpermanent: function (id, storage) {
            var apiurl = `/api/vm/delvmpermanent`
            this.$http.get(apiurl, { params: { uuid: id , storage: storage} }).then(response => {
			    if (response.data.err === null ) {
                    alert("删除成功")
                    } else {
					alert(response.data.err.Message)
                    }
                })
            },

        getvmarchive: function (start) {
            this.sortitem = item
            this.cpagenumber = start
            sessionStorage.setItem('pagenumber', start)
            var apiurl = `/api/vm/getvmarchive`
            this.$http.get(apiurl, { params: { startpage: start} }).then(response => {
			if (response.data.err === null ) {
                this.data = response.data.res
                this.totalpagenumber = response.data.pagenumber
                this.vmcount = response.data.vmcount
                this.pagenumber = start
				} else {
					alert(response.data.err.Message)
					this.$router.push({name:"login"})
					} 
                })   
            },
        }
  }
</script>

<style scoped>
.dropdown-menu {
    margin-bottom: 0;
    left: -60px;
    right:0;
    min-width: 140px;
}

.col-md-5 {
    padding-right:0;
}

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

.table  thead  tr  th {
    border-bottom: 2px solid #846d6d;
}

th {
	background-color: #77BB95;
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

.glyphicon {
    top: 2px
}

.btn-xs {
    padding-right: 2px;
    padding-left: 2px;
    }
</style>

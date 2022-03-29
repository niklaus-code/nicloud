<template>
<div>
    <div class="col-md-7" style="padding-left: 24px">
        <h5>云盘归档列表({{count}})<h5>
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
       			<td>{{item.Vdiskid}}</td>
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
</div>
</template>
<script>

export default {
    data () {
        return {
            data: [],
            count: 0,
			}
    },

    mounted: function () {
		this.getvdiskarchive()
    },

    methods: {
        delvmpermanent: function (id, storage) {
            alert("暂不开放")
            },

        getvdiskarchive: function (start) {
            var apiurl = `/api/vdisk/getvdiskarchives`
            this.$http.get(apiurl).then(response => {
			if (response.data.err === null ) {
                this.data = response.data.res
                this.count = response.data.res.length
				} else {
					alert(response.data.err.Message)
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

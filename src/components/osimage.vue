<template>
	<div>
  	<div class="content whisper-content leacots-content details-content col-md-11 col-md-offset-2" style="background-color:white; float:left">
		<div class="col-sm-10 col-sm-offset-1" style="margin-top:20px">
			<router-link :to="{name:'createosimage'}">
				<button class="btn btn-success btn-sm" type="button">创建</button>
			</router-link>
			<table class="table table-hover" style="text-align: center;">
    			<thead>
      				<tr>
        				<th>
        					<label class="checkbox-inline" style="border:red 1px">
            					<input type="checkbox" v-model="checkvalue" @click="checkbox()">
        					</label>
        				</th>
        				<th>镜像名称</th>
        				<th>块设备</th>
        				<th>快照名称</th>
        				<th>XML</th>
        				<th>存储集群</th>
        				<th>数据中心</th>
						<th>状态</th>
						<th>操作</th>
      				</tr>
    			</thead>

				<tbody v-for="(item, index) in data">
      				<tr class="table-dark text-dark" :id="item.Uuid">
        				<label class="checkbox-inline">
            				<input type="checkbox" v-model="item.Checkout">
        				</label>
        				<td>{{item.Osname}}</td>
        				<td>{{item.Cephblockdevice}}</td>
        				<td>{{item.Snapimage}}</td>
        				<td class="tdxml" width="35%">{{item.Xml}}</td>
        				<td>{{item.Storage}}</td>
        				<td>{{item.Datacenter}}</td>
						<td>
                            <span v-if="item.Status"  class="glyphicon glyphicon-ok"></span>
                            <span v-else class="glyphicon glyphicon-remove"></span>
		    			</td>
		    			<td>
							<button class="btn btn-info btn-xs" type="button" @click="editosimage(item.Id, item.Osname, item.Cephblockdevice, item.Snapimage, item.Xml)">
                				编辑
            				</button>
							<button class="btn btn-info btn-xs" type="button" @click="delosimage(item.Osname, index)">
                				删除
            				</button>
        				</td>
					</tr>
				</tbody>
			</table>
		</div>
	</div>		
</div>		
</template>
<script>
export default {
    data () {
        return {
			data: [],
        }
    },

	mounted: function () {
		this.getosimage()
		},

    methods: {
		editosimage: function (id, osname, cephblockdevice, snapimage, xml) {
			this.$emit("toParent", "updateosimage");
				/*
            this.$router.push({
            path: '/updateosimage',
                query: { 
                    'osimage': osname,
					"cephblockdevice": cephblockdevice,
					"snapimage" : snapimage,
					"xml": xml,
					"id": id,
                }
            }) 
				*/
            },

		delosimage: function (osname, index) {
            var apiurl = `/api/osimage/delimage`
            this.$http.get(apiurl, { params: {osname: osname} } ).then(response => {
            	if (response.data.res === null) {
					alert("删除成功")
					this.data[index].Status = 0
					} else {
					alert("删除失败(' "+ response.data.res.Message+"')")
					}
            })
        },

		getosimage: function () {
            var apiurl = `/api/osimage/getimage`
            this.$http.get(apiurl).then(response => {
            	this.data = response.data.res
            })
        },
        }
  }
</script>
<style scoped>

select{
    font-family: "微软雅黑";
    border: 1px #1a1a1a solid;
    border-radius: 5px;
}

.content {
    box-shadow: 0 0 10px rgba(0,0,0,8);
    border-radius: 10px/10px;
    z-index: -1;
    padding: 50px 0px 50px 0px;
    margin-left: 0px;
    margin-TOP: 50px;
}
.checkbox-inline {
    margin-bottom: 30px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
	font-weight : 400;
	margin-top: 5px;
}

.table tbody tr td {
    padding: 12px;
    vertical-align: "middle";
}

.tdxml {
	max-width: 100px;
 	overflow: hidden; 
	text-overflow:ellipsis;
	white-space: nowrap;
}

th {
	font-weight: bold;
    color: black;
    text-align: center;
}

</style>

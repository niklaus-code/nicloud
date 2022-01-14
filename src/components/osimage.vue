<template>
		<div class="col-sm-12" style="margin-top:20px">
			<button class="btn btn-success btn-sm" type="button" @click="toParent">创建</button>
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
      				<tr class="table-dark text-dark" :id="item.Uuid" v-if="item.Status">
        				<label class="checkbox-inline">
            				<input type="checkbox" v-model="item.Checkout">
        				</label>
        				<td>{{item.Osname}}</td>
        				<td>{{item.Cephblockdevice}}</td>
        				<td>{{item.Snapimage}}</td>
        				<td class="tdxml" width="30%">{{item.Xml}}</td>
        				<td>{{item.Storage}}</td>
        				<td>{{item.Datacenter}}</td>
						<td>
                            <span v-if="item.Status"  class="glyphicon glyphicon-ok"></span>
                            <span v-else class="glyphicon glyphicon-remove"></span>
		    			</td>
		    			<td style="min-width:92px">
							<button class="btn btn-info btn-xs" type="button" @click="editosimage(item.Id, item.Osname, item.Cephblockdevice, item.Snapimage, item.Xml)">
                				编辑
            				</button>
							<button class="btn btn-danger btn-xs" type="button" @click="delosimage(item.Osname, index)">
                				删除
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
        }
    },

	mounted: function () {
		this.getosimage()
		},

    methods: {
	   	toParent: function (item) {
            this.$emit("toParent", "createosimage");
            },

		editosimage: function (id, osname, cephblockdevice, snapimage, xml) {
			this.$store.state.osimage.id = id
            this.$store.state.osimage.osname = osname
            this.$store.state.osimage.cephblockdevice = cephblockdevice
            this.$store.state.osimage.snap = snapimage
            this.$store.state.osimage.xml = xml
			this.$emit("toParent", "updateosimage");
            },

		delosimage: function (osname, index) {
            var apiurl = `/api/osimage/delimage`
            this.$http.get(apiurl, { params: {osname: osname} } ).then(response => {
            	if (response.data.err === null) {
					alert("删除成功")
					this.data[index].Status = 0
					} else {
					alert(response.data.err.Message)
					}
            })
        },

		getosimage: function () {
            var apiurl = `/api/osimage/getimage`
            this.$http.get(apiurl).then(response => {
            	if (response.data.err === null) {
            	    this.data = response.data.res
                } else {
					alert(response.data.err.Message)
                    }
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

.checkbox-inline {
    margin-bottom: 30px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

input {
	margin-top: 2px;
}

label {
	font-weight : 400;
}

.table tbody tr td {
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

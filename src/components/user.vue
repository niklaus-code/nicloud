<template>
		<div class="col-sm-12" style="margin-top:20px; padding-right:0; padding-left:0">
			<button class="btn btn-success btn-sm" type="button" @click="toParent"  style="margin-bottom:20px; margin-left:5px">创建用户</button>
			<table class="table table-hover" style="text-align: center;">
    			<thead>
      				<tr>
        				<th>
        					<label class="checkbox-inline" style="border:red 1px">
            					<input type="checkbox" v-model="checkvalue" @click="checkbox()">
        					</label>
        				</th>
        				<th>用户名</th>
        				<th>密码</th>
        				<th>邮箱</th>
        				<th>角色</th>
        				<th>手机号</th>
        				<th>创建时间</th>
        				<th>操作</th>
      				</tr>
    			</thead>

				<tbody v-for="(item, index) in data">
      				<tr class="table-dark text-dark" :id="item.Uuid">
        				<label class="checkbox-inline">
            				<input type="checkbox" v-model="item.Checkout">
        				</label>
        				<td>{{item.Username}}</td>
        				<td><span>●●●●●●●●●●●●●</span></td>
        				<td>{{item.Email}}</td>
        				<td>{{item.Role}}</td>
        				<td>{{item.Mobile}}</td>
        				<td>{{item.Create_time}}</td>
		    			<td style="min-width:92px">
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
            this.$emit("toParent", "createuser");
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
            var apiurl = `/api/user/getuser`
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
    background-color: #e3e3e3;
	font-weight: bold;
    color: black;
    text-align: center;
}

</style>

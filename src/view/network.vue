<template>
	<div>
	<nicloudhead></nicloudhead>
	<vmleft></vmleft>
  	<div class="content whisper-content leacots-content details-content col-md-11 col-md-offset-2" style="background-color:white; float:left">
		<div class="col-sm-6 col-sm-offset-3" style="margin-top:20px;">
			<router-link :to="{name:'createvlan'}">
				<button class="btn btn-info btn-sm" type="button">创建</button>
			</router-link>
			<table class="table table-hover" style="text-align: center;">
    			<thead>
      				<tr>
        				<th>
        					<label class="checkbox-inline" style="border:red 1px">
            					<input type="checkbox" v-model="checkvalue" @click="checkbox()">
        					</label>
        				</th>
        				<th>vlan</th>
        				<th>网桥</th>
        				<th>地址段</th>
						<th>网关</th>
						<th>状态</th>
						<th>操作</th>
      				</tr>
    			</thead>

				<tbody v-for="(item, index) in data">
      				<tr v-if="item.Status" class="table-dark text-dark" :id="item.Uuid">
        				<label class="checkbox-inline">
            				<input type="checkbox" v-model="item.Checkout">
        				</label>
        				<td>{{item.Vlan}}</td>
        				<td>{{item.Bridge}}</td>
        				<td>{{item.Network}}/{{item.Prefix}}</td>
        				<td>{{item.Gateway}}</td>
        				<td>{{item.Status}}</td>
		    			<td>
							<button class="btn btn-success btn-xs" type="button" @click="">
                				增加IP
            				</button>
							<button class="btn btn-success btn-xs" type="button" @click="">
                				查看IP
            				</button>
							<button class="btn btn-danger btn-xs" type="button" @click="">
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
import foot from '@/components/footer'
import nicloudhead from '@/components/nicloudhead'
import vmleft from '@/components/vmleft'


export default {
    data () {
        return {
			data: [],
			cpu: "",
			mem: "",
			ip: "",
			num: "",
        }
    },

    components: {
        foot, nicloudhead, vmleft
    },

	mounted: function () {
		this.getvlan()
		},

    methods: {
		delhost: function (ip, index) {
			this.data[index].Status = 0
            var apiurl = `/api/vm/delhost`
            this.$http.get(apiurl, { params: {ip: ip} } ).then(response => {
            	if (response.data.res) {
					alert("删除成功")
					}
            })
        },

		getvlan: function (ip) {
            var apiurl = `/api/networks/getvlan`
            this.$http.get(apiurl).then(response => {
            	this.data = response.data.res
            })
        },
        }
  }
</script>
<style>

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

th {
	font-weight: bold;
    color: black;
    text-align: center;
}

</style>

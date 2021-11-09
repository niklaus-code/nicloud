<template>
	<div>
	<nicloudhead></nicloudhead>
	<vmleft></vmleft>
  	<div class="content whisper-content leacots-content details-content col-md-11 col-md-offset-2" style="background-color:white; float:left">
		<div class="col-sm-10 col-sm-offset-1" style="margin-top:20px">
			<router-link :to="{name:'createcloudrive'}">
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
        				<th>CLOUDRIVEID</th>
        				<th>容量</th>
        				<th>存储池</th>
        				<th>挂载云主机</th>
        				<th>存储集群</th>
						<th>数据中心</th>
        				<th>用户</th>
        				<th>状态`</th>
						<th>操作</th>
      				</tr>
    			</thead>

				<tbody v-for="(item, index) in data">
      				<tr class="table-dark text-dark" :id="item.Uuid">
        				<label class="checkbox-inline" style="width:10px">
            				<input type="checkbox" v-model="item.Checkout">
        				</label>
        				<td>{{item.Cloudriveid}}</td>
        				<td>{{item.Contain}}G</td>
        				<td>{{item.Pool}}</td>
        				<td>{{item.Vm_ip}}</td>
        				<td>{{item.Storage}}</td>
        				<td>{{item.Datacenter}}</td>
        				<td>{{item.User}}</td>
				      <td>
                            <span v-if="item.Status"  class="glyphicon glyphicon-ok"></span>
                            <span v-else class="glyphicon glyphicon-remove"></span>
                        </td>
		    			<td>
							<button v-if="item.Status" class="btn btn-success btn-xs" type="button" @click="mount(item.Cloudriveid, item.Storage, item.Pool, index)">
                				挂载
            				</button>
							<button v-else class="btn btn-info btn-xs" type="button" @click="umount(item.Vm_ip, item.Storage, item.Datacenter, item.Cloudriveid, index)">
                				卸载
            				</button>
							<button class="btn btn-danger btn-xs" type="button" @click="restore(item.Ipv4, item.Status, index)">
                				销毁
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
		this.getcloudrive()
		},

    methods: {
		mount: function(cloudriveid, storage, pool, index) {
			this.$router.push({
            path: '/mountcloudrive',
                query: { 
					cloudriveid: cloudriveid,
					storage: storage,
					pool: pool
                }
            }) 
			},

		umount: function (vmip, storage, datacenter, cloudriveid) {
            var apiurl = `/api/vm/umountdisk`
            this.$http.get(apiurl , { params: { vmip: vmip, storage: storage, datacenter: datacenter, cloudriveid: cloudriveid} }).then(response => {
				if (response.data.err === null ) {
					alert("卸载成功")
				} else {
					alert ("获取数据失败（"+response.data.err.Message+")")
					}
            })
        },

		getcloudrive: function (ip) {
            var apiurl = `/api/storage/getcloudrive`
            this.$http.get(apiurl).then(response => {
				if (response.data.err === null ) {
            		this.data = response.data.res
				} else {
					alert ("获取数据失败（"+response.data.err.Message+")")
					}
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

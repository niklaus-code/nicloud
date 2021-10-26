<template>
	<div>
	<nicloudhead></nicloudhead>
	<vmleft></vmleft>
  	<div class="content whisper-content leacots-content details-content col-md-11 col-md-offset-2" style="background-color:white; float:left">
		<div class="col-sm-2 col-sm-offset-4" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>vlan</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="vlan" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>网桥</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="bridge" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>地址段</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="network" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>子网掩码</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="prefix" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>网关</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="gateway" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
		<div class="form-group" style="margin-top:20px" >
			<div class="col-sm-2 col-sm-offset-4">
  				<button type="submit" @click="createvlan" class="btn btn-default btn-sm">提交</button>
			</div>
		</div>
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
			vlan: "",
			bridge: "",
			network: "",
			prefix: "",
			gateway: "",
        }
    },

    components: {
        foot, nicloudhead, vmleft
    },


    methods: {
		check: function (vlan, bridge, network, prefix, gateway) {
			if (typeof vlan === 'undefined' || vlan === null || vlan === ''|| typeof bridge === 'undefined' || bridge === null || bridge === '' || typeof network === 'undefined' || network === null || network === '' ||typeof prefix === 'undefined' || prefix === null || prefix === ''|| typeof gateway === 'undefined' || gateway === null || gateway === '' ) {
				alert("缺少信息")
                return true
            } else {
				return false
				}
			},

		createvlan: function () {
			if (this.check(this.vlan, this.bridge, this.network, this.prefix,this.gateway)) {
				return 
				}
		

            var apiurl = `/api/networks/createvlan`

            this.$http.get(apiurl, { params: { vlan: this.vlan, bridge:this.bridge, network: this.network, prefix: this.prefix, gateway: this.gateway} }).then(response => {
				if (response.data.res) {
					alert("创建成功! 是否查看宿主机列表")
					this.$router.push('/network')
				} else {
					alert("创建失败(" + response.data.err.Message+ ")" )
					}
			})
			},

        }
  }
</script>
<style scoped>

.col-sm-4 label{
	float: right;
}
select{
    font-family: "微软雅黑";
    border: 1px #1a1a1a solid;
    border-radius: 5px;
}

.content {
    box-shadow: 0 0 10px rgba(0,0,0,8);
    border-radius: 10px/10px;
    z-index: -1;
    padding: 100px 0px 100px 0px;
    margin-left: 0px;
    margin-TOP: 50px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
	font-weight : 400;
	margin-top: 5px;
}
</style>

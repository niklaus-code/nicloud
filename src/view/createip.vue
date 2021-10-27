<template>
	<div>
	<nicloudhead></nicloudhead>
	<vmleft></vmleft>
  	<div class="content whisper-content leacots-content details-content col-md-11 col-md-offset-2" style="background-color:white; float:left">
		<div class="col-sm-8 col-sm-offset-2" >
			<div class="col-sm-12">
			<span class="vlaninfo">vlan信息</span>
			</div>
			<div class="col-sm-3">
				vlan名称：{{vlan.Vlan}}
			</div>
			<div class="col-sm-3">
				网桥名称：{{vlan.Bridge}}
			</div>
			<div class="col-sm-3">
				地址段：{{vlan.Network}}/{{vlan.Prefix}}
			</div>
			<div class="col-sm-3">
				网关：{{vlan.Gateway}}
			</div>
		</div>
		<div class="col-sm-8 col-sm-offset-2" >
			<div class="col-sm-12">
			<span class="createip">生成IP</span>
			</div>
			<div class="col-sm-5 startip">
				<form class="form-inline" style="width:100%" >
  					<div class="form-group" style="width:100%" >
    					<label class="sr-only" for="exampleInputAmount">Amount (in dollars)</label>
    					<div class="input-group" style="width:100%" >
      						<div class="input-group-addon">起始IP</div>
      						<input style="width:100%" type="text" class="form-control"  placeholder="10.0.0.1">
    					</div>
  					</div>
				</form>
			</div>
			<div  class="col-sm-5 endip">
				<form class="form-inline" style="width:90%" >
  					<div class="form-group" style="width:100%" >
    					<label class="sr-only" for="exampleInputAmount">Amount (in dollars)</label>
    					<div class="input-group"  style="width:100%" >
      						<div class="input-group-addon">结束IP</div>
      						<input type="text" class="form-control"  placeholder="10.0.0.254">
    					</div>
  					</div>
				</form>
			</div>
			<div  class="col-sm-2">
  					<button type="submit" class="btn btn-sm btn-primary">生成IP</button>
  					<button type="submit" class="btn btn-sm btn-primary">导入数据库</button>
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
			vlan: {},
        }
    },

    components: {
        foot, nicloudhead, vmleft
    },
	
    created: function () {
        this.vlaninfo()
    },


    methods: {
		vlaninfo: function () {
			this.vlan = JSON.parse(this.$route.query.vlan)
			},
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

.createip {
	font-weight:500
}

.vlaninfo {
	font-weight:501
}
.col-sm-2 {
	padding-left:0;
	margin-top:10px;
}

.col-sm-3 {
	margin-top:10px;
}
.col-sm-8 {
	padding: 10px;
	border-style: solid;
	border-color: #ddd;
	border-width: 1px;
	border-radius: 4px 4px 0 0;
}

.startip {
	margin-top: 10px;
	padding-right: 0px;
}

.endip {
	margin-top: 10px;
	padding-right: 0px;
	padding-left: 0px;
}

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
    padding: 80px 0px 80px 0px;
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

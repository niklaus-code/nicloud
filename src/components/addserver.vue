<template>
	<div>
		<div class="col-sm-4 col-sm-offset-2" style="margin-top:50px">
			<form class="form-horizontal">
  				<div class="form-group">
    				<label class="col-sm-4 control-label">资产名称*</label>
    				<div class="col-sm-8">
      					<input type="email" class="form-control" v-model="zichanmingcheng" placeholder="服务器 || 交换机">
    				</div>
  				</div>
  				<div class="form-group">
    				<label  class="col-sm-4 control-label">品牌*</label>
    				<div class="col-sm-8">
      				<input class="form-control" v-model="pinpai"  placeholder="DELL">
    				</div>
  				</div>
  				<div class="form-group">
    				<label  class="col-sm-4 control-label">型号*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="xinghao"  placeholder="R740xd">
    				</div>
  				</div>
  				<div class="form-group">
    				<label  class="col-sm-4 control-label">原厂序列号*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="xuliehao"  placeholder="12345678">
    				</div>
  				</div>
  				<div class="form-group">
    				<label  class="col-sm-4 control-label">资产标签*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="zichanbiaoqian"  placeholder="abcdefg">
    				</div>
  				</div>
  				<div class="form-group">
    				<label class="col-sm-4 control-label">单位*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="danwei"  placeholder="cnic">
    				</div>
  				</div>
  				<div class="form-group">
    				<label class="col-sm-4 control-label">所属部门*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="suoshubumen"  placeholder="大数据部">
    				</div>
  				</div>
  				<div class="form-group">
    				<label  class="col-sm-4 control-label">责任部门*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="zichanzerenbumen"  placeholder="大数据部">
    				</div>
  				</div>
  				<div class="form-group">
    				<label class="col-sm-4 control-label">责任人*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="zerenren"  placeholder="满宇帅">
    				</div>
  				</div>
  				<div class="form-group">
    				<label class="col-sm-4 control-label">机房*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="suoshujifang"  placeholder="201">
    				</div>
  				</div>
			</form>
		</div>
		<div class="col-sm-4" style="margin-top:50px">
			<form class="form-horizontal">
  				<div class="form-group">
    				<label  class="col-sm-4 control-label">机柜*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="jigui"  placeholder="h03">
    				</div>
  				</div>
  				<div class="form-group">
    				<label class="col-sm-4 control-label">机柜资产标签*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="jiguizichanbiaoqian"  placeholder="123456">
    				</div>
  				</div>
  				<div class="form-group">
    				<label  class="col-sm-4 control-label">机柜位置*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="weizhi"  placeholder="17-18u">
    				</div>
  				</div>
  				<div class="form-group">
    				<label  class="col-sm-4 control-label">高度*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="gaodu"  placeholder="2u">
    				</div>
  				</div>
  				<div class="form-group">
    				<label class="col-sm-4 control-label">设备状态*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="shebeizhuangtai"  placeholder="开机">
    				</div>
  				</div>
  				<div class="form-group">
    				<label class="col-sm-4 control-label">额定功率*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="edinggonglv"  placeholder="100kw">
    				</div>
  				</div>
  				<div class="form-group">
    				<label class="col-sm-4 control-label">用电等级*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="yongdiandengji"  placeholder="4级">
    				</div>
  				</div>
  				<div class="form-group">
    				<label class="col-sm-4 control-label">管理IP*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="guanliip"  placeholder="10.0.80.83">
    				</div>
  				</div>
  				<div class="form-group">
    				<label class="col-sm-4 control-label">业务IP*</label>
    				<div class="col-sm-8">
      					<input class="form-control"  v-model="yewuip" placeholder="10.0.82.83">
    				</div>
  				</div>
  				<div class="form-group">
    				<label  class="col-sm-4 control-label">备注*</label>
    				<div class="col-sm-8">
      					<input class="form-control" v-model="beizhu"  placeholder="evcloud服务器">
    				</div>
  				</div>
			<form>
		</div>
		<div class="col-sm-4 col-sm-offset-2" style="margin-top:50px;">
		    <div class="col-sm-2 col-sm-offset-4" style="padding-left:10" >
  			    <div class="form-group">
			        <button class="btn btn-success" @click="addmachine">提交</button>
		        </div>
		    </div>
		</div>
	</div>
</template>
<script>
import nicloudhead from '@/components/nicloudhead'

export default {
    data () {
        return {
			zichanmingcheng: "",
			pinpai: "",
			xinghao: "",
			xuliehao: "",
			zichanbiaoqian: "",
			danwei: "",
			suoshubumen: "",
			zichanzerenbumen: "",
			zerenren: "",
			suoshujifang: "",
			jigui: "",
			jiguizichanbiaoqian: "",
			weizhi: "",
			gaodu: "",
			shebeizhuangtai: "", 
			edinggonglv: "",
			yongdiandengji: "",
			guanliip: "",
			yewuip: "",
			beizhu: "",	
			data: [],
			}
		},
    components: {
        nicloudhead
    },
	mounted: function () {
		this.getmachinelist()
		},

	methods: {
		getmachinelist: function () {
            sessionStorage.setItem('router', "addserver")
			var apiurl = `/api/machine/getmachinelist`
			this.$http.get(apiurl).then(response => {
                this.data = response.data.res
            })
			},

		addmachine: function () {
			var apiurl = `/api/machine/addmachine`
			this.$http.get(apiurl, { params: { zichanmingcheng: this.zichanmingcheng, pinpai: this.pinpai, xinghao: this.xinghao, xuliehao: this.xuliehao, zichanbiaoqian: this.zichanbiaoqian, danwei: this.danwei, suoshubumen: this.suoshubumen, zichanzerenbumen: this.zichanzerenbumen, zerenren: this.zerenren, suoshujifang: this.suoshujifang, jigui: this.jigui, jiguizichanbiaoqian: this.jiguizichanbiaoqian, weizhi: this.weizhi, gaodu: this.gaodu, shebeizhuangtai: this.shebeizhuangtai, edinggonglv: this.edinggonglv, yongdiandengji: this.yongdiandengji, guanliip:this.guanliip, yewuip: this.yewuip, beizhu: this.beizhu}}).then(response => {
                if (response.data.res) {
					alert("保存成功")
					}
            })
			},
		}
	}
</script>

<style>
input {
	margin-right: 5px;
	height: 29px;
	}

.machine {
	margin-top: 100px;
}
</style>

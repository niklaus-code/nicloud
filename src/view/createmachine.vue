<template>
	<div>
		<nicloudhead> </nicloudhead>
		<div class="col-sm-4 col-sm-offset-4">
			<form class="form-horizontal">
  				<div class="form-group">
    				<label for="inputEmail3" class="col-sm-2 control-label">资产名称*</label>
    				<div class="col-sm-10">
      					<input type="email" class="form-control" id="inputEmail3" placeholder="服务器 || 交换机">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">品牌*</label>
    				<div class="col-sm-10">
      				<input type="password" class="form-control" id="inputPassword3" placeholder="DELL">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">型号*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="R740xd">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">原厂序列号*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="12345678">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">资产标签*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="abcdefg">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">单位*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="cnic">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">所属部门*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="大数据部">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">责任部门*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="大数据部">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">责任人*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="满宇帅">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">机房*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="201">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">机柜*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="h03">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">机柜资产标签*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="123456">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">机柜位置*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="17-18u">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">高度*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="2u">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">设备状态*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="开机">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">额定功率*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="100kw">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">用电等级*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="4级">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">管理IP*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="10.0.80.83">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">业务IP*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="10.0.82.83">
    				</div>
  				</div>
  				<div class="form-group">
    				<label for="inputPassword3" class="col-sm-2 control-label">备注*</label>
    				<div class="col-sm-10">
      					<input type="password" class="form-control" id="inputPassword3" placeholder="evcloud服务器">
    				</div>
  				</div>
			</form>
		</div>
	</div>
</template>
<script>
import nicloudhead from '@/components/nicloudhead'

export default {
    data () {
        return {
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
			var apiurl = `/api/machine/getmachinelist`
			this.$http.get(apiurl).then(response => {
                this.data = response.data.res
            })
			}
		}
	}
</script>

<style>
input {
	margin-right: 5px;
	height: 29px;
	}

.machine {
	margin-top: 80px;
}
</style>

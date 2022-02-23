<template>
<div>
    <div class="col-sm-12 form-group" style="margin-top:30px;border-bottom: 1px green solid">
        <h4>创建存储</h4>
    </div>

	<div class="col-sm-8 col-sm-offset-1" style="margin-top:30px; margin-bottom:30px">
		<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>数据中心</label>
				</div>
				<div class="col-sm-8">
				    <select class="col-sm-12" v-model="centervalue">
                        <option  v-for="c in datacenter" :value="c.Datacenter">
                            {{ c.Datacenter }}
                        </option>
                    </select>
				</div>
    		</div>
    	</div>
		<div class="col-sm-12" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>存储名称</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="storagename" placeholder="">
  						</div>
					</form>
				</div>
			</div>
    	</div>
		<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>ceph-uuid</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="uuid" placeholder="">
  						</div>
					</form>
				</div>
			</div>
    	</div>
		<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>数据池</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="pool" placeholder="">
  						</div>
					</form>
				</div>
			</div>
    	</div>
		<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>密钥</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="ceph_secret" placeholder="">
  						</div>
					</form>
				</div>
			</div>
    	</div>
		<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>hosts</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="ips" placeholder="">
  						</div>
					</form>
				</div>
			</div>
    	</div>
		<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>端口</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="port" placeholder="">
  						</div>
					</form>
				</div>
			</div>
    	</div>
		<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>备注</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="comment" placeholder="">
  						</div>
					</form>
				</div>
			</div>
    	</div>
		<div class="form-group" style="margin-top:20px" >
			<div class="col-sm-2 col-sm-offset-7">
  				<button type="submit" @click="createceph" class="btn btn-success">提交</button>
			</div>
		</div>
	</div>
</div>
</template>
<script>
export default {
    data () {
        return {
			centervalue: "",
            datacenter: [],

			uuid: "",
			pool: "",
			ceph_secret: "",
			ips: "",
			port: "",
			comment: "",
            storagename: "",
        }
    },

	mounted: function() {
        this.getdatacenter()
        },

    methods: {
        getdatacenter: function () {
            var apiurl = `/api/datacenter/getdatacenter`
            
            this.$http.get(apiurl).then(response => {
                if (response.data.err === null) {
                    this.datacenter = response.data.res
                    this.centervalue = response.data.res[0].Datacenter
                } else {
                    alert("获取数据失败(" + response.data.err.Message+ ")" )
                    }
            })
            },

		check: function (osname, cephblockdevice, snapimage, xml) {
			if (typeof osname === 'undefined' || osname === null || osname === ''|| typeof cephblockdevice === 'undefined' || cephblockdevice === null || cephblockdevice === '' || typeof snapimage === 'undefined' || snapimage === null || snapimage === '' ||typeof xml === 'undefined' || xml === null || xml === '') {
				alert("缺少信息")
                return true
            } else {
				return false
				}
			},

		createceph: function () {
            var apiurl = `/api/storage/add`

            this.$http.post(apiurl,  this.$qs.stringify({ uuid: this.uuid, storagename: this.storagename,  pool: this.pool, datacenter: this.centervalue, ceph_secret: this.ceph_secret, port: this.port, ips: this.ips, comment: this.comment})).then(response => {
				if (response.data.err === null) {
					alert("创建成功!")
					this.$emit("toParent", "storage");
				} else {
					alert(response.data.err.Message)
					}
			})
			},

        }
  }
</script>
<style scoped>
.form-control {
    height:30px;
}

.col-sm-4 label{
	float: right;
}

select{
	height:30px;
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
}

label {
	font-weight : 400;
	margin-top: 5px;
}
</style>

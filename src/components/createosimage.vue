<template>
<div>
      <div class="col-sm-12 form-group" style="border-bottom: 1px green solid; margin-top:30px">
                <h4>创建镜像</h4>
            </div>

		<div class="col-sm-8 col-sm-offset-1">
			<div class="col-sm-12">
	 			<div class="form-group">
					<div class="col-sm-2 col-sm-offset-2">
        				<label>数据中心</label>
					</div>
					<div class="col-sm-8">
				    	<select class="col-sm-12" v-model="centervalue" @change="getstorage(centervalue)">
							<option value="">--请选择--</option>
                        	<option  v-for="c in datacenter" :value="c.Datacenter">
                            	{{ c.Datacenter }}
                        	</option>
                    	</select>
					</div>
				</div>
    		</div>
			<div class="col-sm-12" style="margin-top:20px">
	 			<div class="form-group">
					<div class="col-sm-2 col-sm-offset-2">
        				<label>存储集群</label>
					</div>
					<div class="col-sm-8">
				    	<select class="col-sm-12" v-model="storagevalue">
							<option value="">--请选择--</option>
                        	<option  v-for="c in storage" :value="c.Uuid">
                            	{{ c.Name }}
                        	</option>
                    	</select>
					</div>
				</div>
    		</div>
			<div class="col-sm-12" style="margin-top:20px">
	 			<div class="form-group">
					<div class="col-sm-2 col-sm-offset-2">
        				<label>镜像类别</label>
					</div>
					<div class="col-sm-8">
				    	<select class="col-sm-12" v-model="sortvalue">
							<option value="">--请选择--</option>
                        	<option v-for="s in sort" :value="s">
                            	{{ s.sort }}
                        	</option>
                    	</select>
					</div>
				</div>
    		</div>
			<div class="col-sm-12" style="margin-top:20px">
	 		    <div class="form-group">
				    <div class="col-sm-2 col-sm-offset-2">
        			    <label>镜像名称</label>
				    </div>
				    <div class="col-sm-8">
					    <form role="form">
  						    <div class="form-group">
    						    <input type="text" class="form-control" v-model="osimage" placeholder="">
  						    </div>
					    </form>
				    </div>
				</div>
    		</div>
			<div class="col-sm-12" style="margin-top: 6px">
	 		    <div class="form-group">
				    <div class="col-sm-2 col-sm-offset-2">
        			    <label>ceph块名称</label>
				    </div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="cephblockdevice" placeholder="">
  						</div>
					</form>
				</div>
				</div>
    		</div>
			<div class="col-sm-12" style="margin-top: 6px">
	 		    <div class="form-group">
				    <div class="col-sm-2 col-sm-offset-2">
        			    <label>是否创建快照</label>
				    </div>
				    <div class="col-sm-8">
                        <div class="checkbox" style="margin-top:0">
                            <label style="float: left">
                                <input type="checkbox" @click="checkbox" style="margin-top: -6px"><span style="color: #999">勾选创建快照并以此为基础克隆镜像</span>
                            </label>
                        </div>
				    </div>
				</div>
    		</div>
			<div class="col-sm-12" style="margin-top: 10px">
	 			<div class="form-group">
					<div class="col-sm-2 col-sm-offset-2">
        				<label>镜像xml</label>
					</div>
					<div class="col-sm-8">
						<form role="form">
  							<div class="form-group">
    							<textarea class="form-control" v-model="xml" rows="16"></textarea>
  							</div>
						</form>
					</div>
				</div>
    		</div>
			<div class="col-sm-12" style=" margin-bottom:20px">
				<div class="form-group" style="margin-top:20px;" >
					<div class="col-sm-2 col-sm-offset-2">
					</div>
					<div class="col-sm-8" style="margin:0 auto; text-align: center;">
  						<button type="submit" style="margin:0 auto" @click="createosimage" class="btn btn-success btn-sm">提交</button>
					</div>
				</div>
			</div>
	</div>
</div>
</template>
<script>
export default {
    data () {
        return {
            sortvalue: "",
            sort: [
                {id: 1, sort: "基础镜像"},
                {id: 2, sort: "用户镜像"},
                ],

           	centervalue: "",
            datacenter: [],

			storage : [],
            storagevalue: "",

			osimage: "",
			cephblockdevice: "",
			xml: "",

            checkboxobj: false,
        }
    },

    created: function () {
        this.vlaninfo()
		this.getdatacenter()
    },

    methods: {
        checkbox: function () {
            if (this.checkboxobj == false ) {
                alert("已勾选")
                this.checkboxobj = true
                } else {
                    alert("已取消")
                    this.checkboxobj = false
                }
            },

        getdatacenter: function () {
            var apiurl = `/api/datacenter/getdatacenter`

            this.$http.get(apiurl).then(response => {
                if (response.data.err === null) {
                    this.datacenter = response.data.res
                } else {
                    alert("获取数据失败(" + response.data.err.Message+ ")" )
                    }
            })
            },

       getstorage: function (centervalue) {
            var apiurl = `/api/storage/get`
            this.$http.get(apiurl, { params: { datacenter: centervalue}}).then(response => {
                if (response.data.err === null) {
                    this.storage = response.data.res
                    this.storagevalue = response.data.res[0].Uuid
                } else {
                    alert("获取数据失败(" + response.data.err.Message+ ")" )
                    }
            })
        },


		vlaninfo: function () {
            this.osimage = this.$route.query.osimage
            this.cephblockdevice = this.$route.query.cephblockdevice
            this.xml = this.$route.query.xml
            },

		check: function (osname, cephblockdevice, xml) {
			if (typeof osname === 'undefined' || osname === null || osname === ''|| typeof cephblockdevice === 'undefined' || cephblockdevice === null || cephblockdevice === '' ||typeof xml === 'undefined' || xml === null || xml === '') {
				alert("缺少信息")
                return true
            } else {
				return false
				}
			},

		createosimage: function () {
			if (this.check(this.osimage, this.cephblockdevice, this.xml)) {
				return 
				}
		

            var apiurl = `/api/osimage/createimage`

            this.$http.post(apiurl, this.$qs.stringify({osname: this.osimage, datacenter: this.centervalue, storage: this.storagevalue, cephblockdevice: this.cephblockdevice, createsnap: this.checkboxobj, xml: this.xml, sort: this.sortvalue.id})).then(response => {
				if (response.data.err === null) {
					alert("创建成功!")
                	this.$emit("toParent", "osimage");
				} else {
					alert(response.data.err.Message)
					}
			})
			},

        }
  }
</script>
<style scoped>
.btn-sm {
    line-height: 0;
    padding: 3px 3px;
}

select{
	height:30px;
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
}

label {
	float: right;
	font-weight : 400;
	margin-top: 5px;
}
</style>

<template>
<div>
      <div class="col-sm-12 form-group" style="border-bottom: 1px green solid; margin-top:30px">
                <h4>创建镜像</h4>
            </div>

		<div class="col-sm-12" style="background: rgb(249, 231, 186); padding: 10px">
			<div class="col-sm-12">
	 			<div class="form-group">
					<div class="col-sm-2 col-sm-offset-2">
        				<label>数据中心</label>
					</div>
					<div class="col-sm-6">
                        <label :class = "datacenterclick == index ? 'addclass' : 'labelbackcolordefault' " v-for="(x, index) in datacenter" @click="clickdatacenter(index)">{{x.Datacenter}}</label>
					</div>
				</div>
    		</div>
			<div class="col-sm-12" style="margin-top:10px">
	 			<div class="form-group">
					<div class="col-sm-2 col-sm-offset-2">
        				<label>存储集群</label>
					</div>
					<div class="col-sm-6">
                        <label :class = "storageclick == index ? 'addclass' : 'labelbackcolordefault' " v-for="(x, index) in storage" @click="clickstorage(index)">{{x.Name}}</label>
					</div>
				</div>
    		</div>
			<div class="col-sm-12" style="margin-top:10px">
	 			<div class="form-group">
					<div class="col-sm-2 col-sm-offset-2">
        				<label>镜像类别</label>
					</div>
					<div class="col-sm-6">
                        <label :class = "sortclick == index ? 'addclass' : 'labelbackcolordefault' " v-for="(x, index) in sort" @click="clicksort(index)">{{x.Sort}}</label>
					</div>
				</div>
			</div>
			<div class="col-sm-12" style="margin-top:10px">
	 			<div class="form-group">
					<div class="col-sm-2 col-sm-offset-2">
        				<label>系统标签</label>
					</div>
					<div class="col-sm-6">
                        <label :class = "tagclick == index ? 'addclass' : 'labelbackcolordefault' " v-for="(x, index) in ostags" @click="clicktag(index)">{{x.Tag}}</label>
					</div>
				</div>
			</div>
			<div class="col-sm-12" style="margin-top:10px">
	 	        <div class="form-group">
			        <div class="col-sm-2 col-sm-offset-2">
        		        <label>镜像名称</label>
			        </div>
			        <div class="col-sm-6">
				        <form role="form">
  					        <div class="form-group">
    					        <input type="text" class="form-control" v-model="osimage" placeholder="">
  					        </div>
				        </form>
			        </div>
    		    </div>
    		</div>
			<div class="col-sm-12" style="margin-top:10px">
	 	        <div class="form-group">
			        <div class="col-sm-2 col-sm-offset-2">
        		        <label>ceph块设备名称</label>
			        </div>
			        <div class="col-sm-6">
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
				    <div class="col-sm-6">
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
        				<label>镜像XML</label>
					</div>
					<div class="col-sm-6">
                        <label :class = "xmlclick == index ? 'addclass' : 'labelbackcolordefault' " v-for="(x, index) in xmllist" @click="clickxml(index)">{{x.Comment}}</label>
					</div>
				</div>
			</div>
			<div class="col-sm-12" style="margin-top: 20px">
				<div class="col-sm-2 col-sm-offset-4">
  					<button type="submit" style="margin-left: 5px" @click="createosimage" class="btn btn-success btn-sm">提交</button>
				</div>
			</div>
	</div>
</div>
</template>
<script>
export default {
    data () {
        return {
            datacenterclick: -1,
            storageclick: -1,
            xmlclick: -1,
            tagclick: -1,
            sortclick: -1,
            xmllist: [],
            size: "",
            sortvalue: "",
            sort: [],
            ostags: [],
            tagvalue: "",

           	centervalue: "",
            datacenter: [],

			storage : [],
            storagevalue: "",

			osimage: "",
			cephblockdevice: "",
			xmlvalue: "",

            checkboxobj: false,
        }
    },

    created: function () {
        this.osimageinfo()
		this.getdatacenter()
        this.getimagesort()
        this.getostag()
        this.getosxml()
    },

    methods: {
        clickstorage: function (index) {
            this.storageclick = index
            this.storagevalue = this.storage[index].Uuid
            },

        clickdatacenter: function (index) {
            this.datacenterclick = index
            this.centervalue = this.datacenter[index].Datacenter
            },

        clicksort: function (index) {
            this.sortclick = index
            this.sortvalue = this.sort[index]
            },

        clicktag: function (index) {
            this.tagclick = index
            this.tagvalue = this.ostags[index]
            },

        clickxml: function (index) {
            this.xmlclick = index
            this.xmlvalue = this.xmllist[index].Id
            },

        getosxml: function () {
            var apiurl = `/api/osimage/getosimagexml`

            this.$http.get(apiurl).then(response => {
                if (response.data.err === null) {
                    this.xmllist = response.data.res
                } else {
                    alert(response.data.err.Message)
                    }
                })

            },

        checkbox: function () {
            if (this.checkboxobj == false) {
                alert("已勾选")
                this.checkboxobj = true
                } else {
                    alert("已取消")
                    this.checkboxobj = false
                }
            },

        getostag: function () {
            var apiurl = `/api/osimage/getiostags`

            this.$http.get(apiurl).then(response => {
                if (response.data.err === null) {
                    this.ostags = response.data.res
                } else {
                    alert(response.data.err.Message)
                    }
                })
            },

        getimagesort: function () {
            var apiurl = `/api/osimage/getimagesort`

            this.$http.get(apiurl).then(response => {
                if (response.data.err === null) {
                    this.sort = response.data.res
                } else {
                    alert(response.data.err.Message)
                    }
                })
            },

        getdatacenter: function () {
            var apiurl = `/api/datacenter/getdatacenter`

            this.$http.get(apiurl).then(response => {
                if (response.data.err === null) {
                    this.datacenter = response.data.res
                    this.getstorage(this.centervalue)
                    this.centervalue = response.data.res[0].Datacenter
                } else {
                    alert(response.data.err.Message)
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


		osimageinfo: function () {
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
			if (this.check(this.osimage, this.cephblockdevice, this.xmlvalue)) {
				return 
				}

            var apiurl = `/api/osimage/createimage`

            this.$http.post(apiurl, this.$qs.stringify({osname: this.osimage, datacenter: this.centervalue, storage: this.storagevalue, cephblockdevice: this.cephblockdevice, createsnap: this.checkboxobj, xml: this.xmlvalue, ossort: this.sortvalue.Id, tag: this.tagvalue.Id})).then(response => {
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
input {
    max-width: 300px;
    }

.col-sm-2 label {
    float: right;
}

label {
    border-radius: 4px;
    font-weight: 400;
    margin-top: 2px;
    margin-left: 5px;
}

.labelbackcolordefault {
    padding-top: 6px;
    padding-bottom: 5px;
    padding-left: 10px;
    padding-right: 10px;
    background: #eaeaea;
}

select {
    margin-left: 5px;
	height:30px;
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
}

input {
    margin-left: 5px;
    }

.addclass{
    padding-top: 5px;
    padding-bottom: 5px;
    padding-left: 10px;
    padding-right: 10px;
    background-color: green;
}
</style>

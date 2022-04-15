<template>
    <div>
        <div class="col-sm-12 form-group" style="border-bottom: 1px green solid">
            <h4>编辑镜像</h4>
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
                                {{ c.Uuid }}
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
				    <div class="col-sm-3">
					    <form role="form">
  						    <div class="form-group">
    						    <input type="text" class="form-control" v-model="osimage" placeholder="">
  						    </div>
					    </form>
				    </div>
				    <div class="col-sm-2">
        			    <label>ceph块</label>
				    </div>
				    <div class="col-sm-3">
					    <form role="form">
  						    <div class="form-group">
    						    <input type="text" class="form-control" v-model="cephblockdevice" placeholder="">
  						    </div>
					    </form>
				    </div>
				</div>
    		</div>
			<div class="col-sm-12">
	 		    <div class="form-group">
				    <div class="col-sm-2 col-sm-offset-2">
        			    <label>镜像类别</label>
				    </div>
				    <div class="col-sm-3">
					    <form role="form">
  						    <div class="form-group">
    						    <input type="text" class="form-control" v-model="sort.Sort" placeholder="">
  						    </div>
					    </form>
				    </div>
				    <div class="col-sm-2">
        			    <label>镜像标签</label>
				    </div>
				    <div class="col-sm-3">
					    <form role="form">
  						    <div class="form-group">
    						    <input type="text" class="form-control" v-model="tag.Tag" placeholder="">
  						    </div>
					    </form>
				    </div>
				</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-2 col-sm-offset-2">
        			<label>快照名称</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="snapimage" placeholder="">
  						</div>
					</form>
				</div>
				</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-2 col-sm-offset-2">
        			<label>xml</label>
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
		<div class="form-group" style="margin-top:20px" >
			<div class="col-sm-1 col-sm-offset-7">
  				<button type="submit" @click="commit" class="btn btn-success">提交</button>
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

			storage : [],
            storagevalue: "",

			osimage: "",
			cephblockdevice: "",
			snapimage: "",
			xml: "",
			id: "",
            tag: "",
            sort: "",
        }
    },

    created: function () {
        this.osinfo()
		this.getdatacenter()
    },

    methods: {
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

		osinfo: function () {
            var v = this.$store.state.osimage.osname
            if (v === null || typeof v === 'undefined' || v === '' || v === "undefined") {
                this.osimage = sessionStorage.getItem('osimage')
                this.cephblockdevice = sessionStorage.getItem('cephblockdevice')
                this.snapimage = sessionStorage.getItem('snapimage')
                this.xml = sessionStorage.getItem('xml')
                this.id = sessionStorage.getItem('id')
                this.tag = JSON.parse(sessionStorage.getItem('tag'))
                this.sort = JSON.parse(sessionStorage.getItem('sort'))
            } else {
                this.osimage = this.$store.state.osimage.osname
                this.cephblockdevice = this.$store.state.osimage.cephblockdevice
                this.snapimage = this.$store.state.osimage.snap
                this.xml = this.$store.state.osimage.xml
                this.id = this.$store.state.osimage.id
                this.tag = this.$store.state.osimage.tag
                this.sort = this.$store.state.osimage.sort
                sessionStorage.setItem('osimage', this.$store.state.osimage.osname)
                sessionStorage.setItem('cephblockdevice', this.$store.state.osimage.cephblockdevice)
                sessionStorage.setItem('snapimage', this.$store.state.osimage.snap)
                sessionStorage.setItem('xml', this.$store.state.osimage.xml)
                sessionStorage.setItem('id', this.$store.state.osimage.id)
                sessionStorage.setItem('tag', JSON.stringify(this.$store.state.osimage.tag))
                sessionStorage.setItem('sort', JSON.stringify(this.$store.state.osimage.sort))
                }
        },

		check: function (osname, cephblockdevice, snapimage, xml) {
			if (typeof osname === 'undefined' || osname === null || osname === ''|| typeof cephblockdevice === 'undefined' || cephblockdevice === null || cephblockdevice === '' || typeof snapimage === 'undefined' || snapimage === null || snapimage === '' ||typeof xml === 'undefined' || xml === null || xml === '') {
				alert("缺少信息")
                return true
            } else {
				return false
				}
			},

		commit: function () {
            var apiurl = `/api/osimage/updateimage`
            this.$http.post(apiurl,this.$qs.stringify({sort:this.sort.Id, id: this.id , datacenter: this.centervalue, storage: this.storagevalue, osname: this.osimage, cephblockdevice: this.cephblockdevice, snapimage: this.snapimage, xml: this.xml, tag: this.tag.Id, sort: this.sort.Id})).then(response => {
				if (response.data.err === null) {
					alert("更新成功!")
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

label {
	float: right;
}

select{
	height:30px;
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
	font-weight : 400;
	margin-top: 5px;
}
</style>

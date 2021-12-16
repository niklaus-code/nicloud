<template>
    <div>
	 		<div class="col-sm-12 form-group" style="border-bottom: 1px green solid">
                <h4>创建云主机</h4>
	 		</div>
		<div  class="col-sm-9 col-sm-offset-1" style="margin-top:20px">
		<div  class="col-sm-10 col-sm-offset-1" style="margin-top:20px">
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-4">
        			<label>数据中心</label>
				</div>
				<div class="col-sm-8">
        			<select class="col-sm-10" v-model="centervalue" @change="getvlan(centervalue)" @change="getstorage(centervalue)">
					  <option value="">--请选择--</option>
  						<option  v-for="c in datacenter" :value="c.Datacenter">
							{{ c.Datacenter }}
						</option>
        			</select>
				</div>
    		</div>
           <div class="col-sm-12 form-group">
                <div class="col-sm-4">
                    <label>存储集群</label>
                </div>
                <div class="col-sm-8">
                    <select class="col-sm-10" v-model="storagevalue" @change="getpool()" @change="getimageby()">
					  <option value="">--请选择--</option>
                        <option  v-for="c in storage" :value="c.Uuid">
                            {{ c.Uuid }}
                        </option>
                    </select>
                </div>
            </div>

	 		<div class="col-sm-12 form-group">
				<div class="col-sm-4">
        			<label>VLAN</label>
				</div>
				<div class="col-sm-8">
        			<select class="col-sm-10" v-model="vlanvalue" @change="getip" @change="gethosts(centervalue)">
					  <option value="">--请选择--</option>
  						<option  v-for="v in vlanlist" :value="v.Vlan">
							{{ v.Vlan }}
						</option>
        			</select>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-4">
        			<label>IP地址</label>
				</div>
				<div class="col-sm-8">
	                <select class="col-sm-10" v-model="ipvalue">
					  	<option value="">--请选择--</option>
                    	<option  v-for="ip in iplist" :value="ip.Ipv4">
                        	{{ ip.Ipv4 }}
						</option>
        			</select>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-4">
        			<label>cpu/内存</label>
				</div>
				<div class="col-sm-8">
        			<select class="col-sm-10" v-model="flavorvalue">
  						<option  v-for="f in flavorlist" :value="f">
							{{ f.Cpu}}核 / {{f.Mem}}G
						</option>
        			</select>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-4">
        			<label>宿主机</label>
				</div>
				<div class="col-sm-8 title">
        			<select class="col-sm-10" v-model="hostvalue">
					  	<option value="">--请选择--</option>
  						<option  v-for="host in hostlist" :value="host.Ipv4">
							 {{ host.Ipv4 }} (&nbsp{{host.Usedcpu}}/{{host.Cpu}}&nbsp核，{{host.Usedmem}}/{{host.Mem}}&nbspG ，{{host.count}}/{{host.Max_vms}}&nbsp数量&nbsp)
						</option>
        			</select>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-4">
        			<label>镜像</label>
				</div>
				<div class="col-sm-8">
        			<select class="col-sm-10" v-model="imagevalue">
					  <option value="">--请选择--</option>
  						<option  v-for="image in imagelist" :value="image.Osname">
							{{ image.Osname }}
						</option>
        			</select>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group create">
				<div class="col-sm-2 col-sm-offset-4">
					<button class="btn btn-success btn-sm"  @click="createvm()">创建</button>
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
			storage : [],
			storagevalue: "",

            centervalue: "",
            datacenter: [],

			vlanvalue: "",
			vlanlist: [],

			imagevalue: "",
            imagelist: [],

			hostvalue: "",
            hostlist: [],

			ipvalue: "",
            iplist: [],

			poolvalue: "",
			pool: [],

			flavorvalue: {},
			flavorlist: [],
			status: {
				0: "关机",
				1: "运行"
				}
        }
    },

	mounted: function () {
		this.getdatacenter()
        this.getflavor () 
		},

    methods: {
       	getpool: function () {
            var apiurl = `/api/storage/getpool`
            this.$http.get(apiurl, { params: { datacenter: this.centervalue, storage: this.storagevalue}}).then(response => {
                if (response.data.err === null) {
                	this.pool = response.data.res
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
                } else {
                    alert("获取数据失败(" + response.data.err.Message+ ")" )
                    }
            })
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

		createvm: function () {
            var apiurl = `/api/vm/create`

			if (typeof this.ipvalue === 'undefined' || this.ipvalue == null || this.ipvalue === '') {
				alert("缺少信息!")
				return
			}
            this.$http.post(apiurl, this.$qs.stringify({datacenter: this.centervalue, storage: this.storagevalue, vlan: this.vlanvalue,  cpu: this.flavorvalue.Cpu, mem:this.flavorvalue.Mem, ip: this.ipvalue, host: this.hostvalue, os: this.imagevalue})).then(response => {
				if (response.data.err === null) {
					alert("创建成功! 是否查看虚拟机列表")
					this.$emit("toParent", "vm");
				} else {
					alert("创建失败('"+response.data.err.Message+"')")
					}
			})
			},

        getimageby: function () {
           	var apiurl = `/api/osimage/getimageby`
            this.$http.get(apiurl, { params: {datacenter:this.centervalue, storage: this.storagevalue}}).then(response => {
				if (response.data.err === null) {
					this.imagelist = response.data.res
				} else {
					alert("创建失败('"+response.data.err.Message+"')")
					}
            	})
        },

        gethosts: function (datacenter) {
            var apiurl = `/api/hosts/gethostsbydatacenter`
           	this.$http.get(apiurl, { params: {datacenter: datacenter, vlan: this.vlanvalue}} ).then(response => {
				if (response.data.err === null) {
            		this.hostlist = response.data.res
					} else {
						alert("获取数据失败('"+response.data.err.Message+"')")
					}
            	})
        },

        getvlan: function (datacenter) {
            var apiurl = `/api/networks/getvlanbydatacenter`
            this.$http.get(apiurl, { params: {datacenter: datacenter}} ).then(response => {
			if (response.data.err === null) {
            	this.vlanlist = response.data.res
				} else {
					alert("创建失败('"+response.data.err.Message+"')")
				}
            })
        },

        getip: function () {
            	var apiurl = `/api/networks/getip`
            	this.$http.get(apiurl, { params: { vlan: this.vlanvalue}}).then(response => {
            	this.iplist = response.data.res
            	})
        },

        getflavor: function () {
            var apiurl = `/api/vm/getflavor`
            this.$http.get(apiurl).then(response => {
            this.flavorlist = response.data.res
			this.flavorvalue = response.data.res[0]
            })
        }
    }
  }
</script>
<style scoped>
.title {
	text-align: center;
}

.col-sm-9 {
	padding-left:0px;
}

.col-sm-3 {
	padding-right:30px;
}
label {
	float: right;
    font-weight : 400;
    margin-top: 2px;
}

select{
	height:30px;
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
}

.create {
    margin-top:20px
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}
</style>

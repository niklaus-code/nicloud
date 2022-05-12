<template>
    <div style="background-color: black">
	 	<div class="col-sm-12 form-group" style="margin-top:30px; border-bottom: 1px green solid">
            <h4>创建云主机
            </h4>
	    </div>
		<div  class="col-sm-12" style="margin-top:20px;background-color: #f9e7ba; padding-top: 20px; box-shadow: 0 0.75rem 1.5rem rgba(18, 38, 63, 0.03)">
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>数据中心</label>
				</div>
				<div class="col-sm-9">
                    <label :class = "datacenterclick == index ? 'addclass' : 'labelbackcolordefault' " v-for="(c, index) in datacenter" @click="clickdatacenter(index)">{{c.Datacenter}}</label>
				</div>
    		</div>
           <div class="col-sm-12 form-group">
                <div class="col-sm-3">
                    <label>存储集群</label>
                </div>
                <div class="col-sm-9">
                    <label :class = "storageclick == index ? 'addclass' : 'labelbackcolordefault' " v-for="(c, index) in storage" @click="clickstorage(index)">{{c.Name}}</label>
                </div>
            </div>

	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>VLAN子网 & IP</label>
				</div>
				<div class="col-sm-3">
        			<select style="margin-left: 5px" class="col-sm-12" v-model="vlanvalue" @change="getip" @change="gethosts(centervalue)">
					  <option value="">--请选择--</option>
  						<option  v-for="v in vlanlist" :value="v.Vlan">
							{{ v.Vlan }}
						</option>
        			</select>
				</div>
				<div class="col-sm-4">
	                <select class="col-sm-12" v-model="ipvalue" style="margin-left: 10px">
					  	<option value="">--请选择--</option>
                    	<option  v-for="ip in iplist" :value="ip.Ipv4">
                        	{{ ip.Ipv4 }}
						</option>
        			</select>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>处理器核心</label>
				</div>
				<div class="col-sm-9">
                    <label :class = "cpuclick == index ? 'addclass' : 'labelbackcolordefault' " v-for="(c, index) in cpulist" @click="clickcpu(index)">{{c}}核心</label>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>内存</label>
				</div>
				<div class="col-sm-9">
                    <label :class = "memclick == index ? 'addclass' : 'labelbackcolordefault' " v-for="(c, index) in memlist" @click="clickmem(index)">{{c}}G</label>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>镜像</label>
				</div>
				<div class="col-sm-3">
        			<select style="margin-left: 5px" class="col-sm-12" v-model="tagvalue" @change="getimagebytag">
					  <option value="">--请选择--</option>
  						<option  v-for="o in ostaglist" :value="o.Id">
							{{ o.Tag }}
						</option>
        			</select>
				</div>
				<div class="col-sm-4">
        			<select style="margin-left: 5px" class="col-sm-12" v-model="imagevalue">
					  <option value="">--请选择--</option>
  						<option  v-for="image in imagelist" :value="image.Id">
							{{ image.Osname }}
						</option>
        			</select>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>宿主机</label>
				</div>
				<div class="col-sm-3 title">
        			<select class="col-sm-12" v-model="hostvalue" style="margin-left: 5px">
					  	<option value="">--请选择--</option>
  						<option  v-for="host in hostlist" :value="host.Ipv4">
							 {{ host.Ipv4 }} ({{host.Usedcpu}}/{{host.Cpu}}&nbsp核，{{host.Usedmem}}/{{host.Mem}}&nbspG ，{{host.count}}/{{host.Max_vms}}&nbsp个)
						</option>
        			</select>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group" style="margin-bottom:-10px">
				<div class="col-sm-3">
        			<label>备注</label>
				</div>
				<div class="col-sm-4">
                    <form role="form" style="margin-left: 5px">
                        <div class="form-group">
                            <input type="text" class="form-control" v-model="comment" placeholder="">
                        </div>
                    </form>
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-8 col-md-offset-4" style="color: #C0C0C0">
    		    </div>
    		</div>
	 		<div class="col-sm-12 form-group create">
				<div class="col-sm-2 col-sm-offset-3">
					<button class="btn btn-success btn-sm" style="margin-left: 5px"  @click="createvm()">创建</button>
				</div>
    		</div>
		</div>
	</div>
</template>
<script>
export default {
    data () {
        return {
            tagvalue: "",

            datacenterclick: -1,
            storageclick: -1,
            cpuclick: -1,
            memclick: -1,

            memvalue: "",
            cpuvalue: "",

            isactive: -1,
            comment: "",

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

            cpulist : [],
            memlist : [],

            ostaglist : [],

			status: {
				0: "关机",
				1: "运行"
				}
        }
    },

	mounted: function () {
        this.datacenterclick = 0
        this.storageclick = 0
        this.vlanclick = 0
		this.getdatacenter()
        this.getflavor() 
        this.getostag()
        this.getip()
		},

    methods: {
         getostag: function () {
            var apiurl = `/api/osimage/getiostags`

            this.$http.get(apiurl).then(response => {
                if (response.data.err === null) {
                    this.ostaglist = response.data.res
                } else {
                    alert(response.data.err.Message)
                    }
                })
            },

        clickmem: function (index) {
            this.memclick = index
            this.memvalue = this.memlist[index]
            },

        clickcpu: function (index) {
            this.cpuclick = index
            this.cpuvalue = this.cpulist[index]
            },

        clickstorage: function (index) {
            this.storageclick = index
            this.storagevalue = this.storage[index]
            },

        clickdatacenter: function (index) {
            this.datacenterclick = index
            this.datacentervalue = this.datacenter[index]
            },

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
                	this.storagevalue = response.data.res[0].Uuid
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
                    this.centervalue = response.data.res[0].Datacenter
                    this.getstorage(this.centervalue)
                    this.getvlan(this.centervalue)
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
            this.$http.post(apiurl, this.$qs.stringify({datacenter: this.centervalue, storage: this.storagevalue, vlan: this.vlanvalue,  cpu: this.cpuvalue, mem:this.memvalue, ip: this.ipvalue, host: this.hostvalue, os: this.imagevalue, comment: this.comment})).then(response => {
				if (response.data.err === null) {
					alert("创建成功! 是否查看虚拟机列表")
					this.$emit("toParent", "vm");
				} else {
					alert(response.data.err.Message)
					}
			    })
			},

        getimagebytag: function () {
           	var apiurl = `/api/osimage/getimagebytag`
            this.$http.get(apiurl, { params: {datacenter:this.centervalue, storage: this.storagevalue, tag: this.tagvalue}}).then(response => {
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
                var l = new Array()
                var m = new Array()
			    if (response.data.err === null) {
                    for (var v in response.data.res) {
                        l.push(response.data.res[v]["Cpu"])
                        m.push(response.data.res[v]["Mem"])
                    }
                 this.cpulist = l
                 this.memlist = m
                 this.clickcpu(0)
                 this.clickmem(0)
                 }
            })
        }
    }
}
</script>
<style scoped>
.col-sm-3 label {
    float: right;
}

.labelbackcolordefault {
    padding-top: 6px;
    padding-bottom: 5px;
    padding-left: 10px;
    padding-right: 10px;
    background: #eaeaea;
}

label {
    border-radius: 4px;
    margin-left: 5px;
    font-weight: 400;
    margin-top: 2px;
}

select{
	height:30px;
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
}

.create {
    margin-top:8px
}

.addclass{
    padding-top: 5px;
    padding-bottom: 5px;
    padding-left: 10px;
    padding-right: 10px;
    background-color: green;
}
</style>

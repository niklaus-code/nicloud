<template>
	<div>
        <div class="col-sm-12 form-group" style="margin-top:10px; border-bottom: 1px green solid">
            <h4>热迁移云主机 <span style="color: #C0C0C0; font-size:13px">&nbsp *迁移目标宿主机需开启selinux</span></h4>
            </div>
		<div class="col-sm-8 col-sm-offset-2" style="margin-top:20px">
	 		<div class="col-sm-12">
				<div class="col-sm-3">
				    <div class="col-sm-10 col-sm-offset-1">
        			<label>ip</label>
				    </div>
				</div>
				<div class="col-sm-9">
					{{ip}}
				</div>
    		</div>
	 		<div class="col-sm-12">
				<div class="col-sm-3">
				    <div class="col-sm-10 col-sm-offset-1">
        			    <label>uuid</label>
				    </div>
				</div>
				<div class="col-sm-9">
					{{uuid}}
				</div>
    		</div>
	 		<div class="col-sm-12">
				<div class="col-sm-3">
				    <div class="col-sm-10 col-sm-offset-1">
        			<label>宿主机</label>
				    </div>
				</div>
				<div class="col-sm-9">
					{{vmhost}}
				</div>
    		</div>
	 		<div class="col-sm-12">
				<div class="col-sm-3">
				    <div class="col-sm-10 col-sm-offset-1">
        			<label>CPU / 内存</label>
				    </div>
				</div>
				<div class="col-sm-9">
					{{cpu}}&nbsp核 / {{mem}}&nbspG
				</div>
    		</div>
	 		<div class="col-sm-12">
				<div class="col-sm-3">
				    <div class="col-sm-10 col-sm-offset-1">
        			<label>系统</label>
				    </div>
				</div>
				<div class="col-sm-9">
					{{os}}
				</div>
    		</div>
	 		<div class="col-sm-12">
				<div class="col-sm-3">
				    <div class="col-sm-10 col-sm-offset-1">
        			<label>创建者</label>
				    </div>
				</div>
				<div class="col-sm-9">
					{{ owner }}
				</div>
    		</div>
	 		<div class="col-sm-12">
				<div class="col-sm-3">
				    <div class="col-sm-10 col-sm-offset-1">
        			<label>备注</label>
				    </div>
				</div>
				<div class="col-sm-9">
					{{ comment }}
				</div>
    		</div>
	 		<div class="col-sm-12 form-group" style="margin-top:20px">
				<div class="col-sm-3">
				    <div class="col-sm-10 col-sm-offset-1">
        			    <button class="btn btn-success btn-sm" style="float:right" @click="migratevmlive()">迁移</button>
				    </div>
				</div>
				<div class="col-sm-9">
				    <select class="form-select col-sm-10" v-model="hostvalue" style="max-width: 400px">
                        <option  v-for="h in host"  :value="h.Ipv4" >
                            {{h.Ipv4}}&nbsp(cpu:{{h.Usedcpu}}/{{h.Cpu}}， &nbsp mem:{{h.Usedmem}}/{{h.Mem}}， &nbsp 数量:{{h.count}}/{{h.Max_vms}})
                        </option>
                    </select>
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
			hostvalue: "",
			uuid: "",
			host: [],
			vmhost: "",
			comment: "",
        }
    },

	created (){
     this.ipvalue = this.iplist[1]
     this.flavorvalue =this.flavorlist[1]
 
	},

    components: {
        foot, nicloudhead, vmleft
    },

    created: function () {
		this.getvminfo()
		this.gethostinfo()
    },

    methods: {
        migratevmlive: function () {
            var apiurl = `/api/vm/migratelive`
            this.$http.get(apiurl, { params: { uuid: this.uuid , migratehost: this.hostvalue} } ).then(response => {
            	if (response.data.res) {
						alert("迁移失败('"+response.data.res.Message+"')")
					} else {
						alert("迁移成功")
					}
            })
        },

        getvminfo: function () {
			var u = this.$store.state.vm.uuid
            if (u === null || typeof u === 'undefined' || u === '' ) {
			    this.uuid = sessionStorage.getItem('uuid')
			    this.ip = sessionStorage.getItem('ip')
			    this.vmhost = sessionStorage.getItem('vmhost')
			    this.cpu = sessionStorage.getItem('cpu')
			    this.mem = sessionStorage.getItem('mem')
			    this.owner = sessionStorage.getItem('owner')
			    this.os = sessionStorage.getItem('os')
			    this.comment = sessionStorage.getItem('comment')
				} else {
				this.uuid = this.$store.state.vm.uuid
				this.ip = this.$store.state.vm.ip
				this.vmhost = this.$store.state.vm.host
				this.cpu = this.$store.state.vm.cpu
				this.mem = this.$store.state.vm.mem
				this.owner = this.$store.state.vm.owner
				this.comment = this.$store.state.vm.comment
				this.os = this.$store.state.vm.os
			    sessionStorage.setItem('uuid', this.$store.state.vm.uuid)
			    sessionStorage.setItem('ip', this.$store.state.vm.ip)
			    sessionStorage.setItem('vmhost', this.$store.state.vm.host)
			    sessionStorage.setItem('cpu', this.$store.state.vm.cpu)
			    sessionStorage.setItem('mem', this.$store.state.vm.mem)
			    sessionStorage.setItem('owner', this.$store.state.vm.owner)
			    sessionStorage.setItem('os', this.$store.state.vm.os)
			    sessionStorage.setItem('comment', this.$store.state.vm.comment)
				}
        },

        gethostinfo: function () {
			var host = this.$route.query.host
            var apiurl = `/api/hosts/gethostsby`
            this.$http.get(apiurl, { params: { ip: host} } ).then(response => {
            	this.host = response.data.res
            })
        }
    }
  }
</script>

<style scoped>

.form-group {
	margin-bottom:6px;
}

.col-sm-3 {
    padding-right: 0px;
}

.col-sm-10 {
    padding-right: 0px;
}

label {
    font-weight : 500;
    margin-top: 0px;
    float: right;
    font-weight: 600;
}

select{
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
    height: 30px;
}

.content {
    box-shadow: 0 0 10px rgba(0,0,0,8);
    border-radius: 10px/10px;
    z-index: -1;
    padding: 70px 0px 70px 0px;
    margin-left: 0px;
    margin-TOP: 50px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}
</style>

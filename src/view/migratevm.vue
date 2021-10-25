<template>
	<div>
	<nicloudhead></nicloudhead>
	<vmleft></vmleft>
  	<div class="content whisper-content leacots-content details-content col-md-11 col-md-offset-2" style="background-color:white; float:left">
		<div  class="col-sm-5 col-sm-offset-3" style="margin-top:20px">
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>uuid</label>
				</div>
				<div class="col-sm-9">
					{{data.Uuid}}
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>ip</label>
				</div>
				<div class="col-sm-9">
					{{data.Ip}}
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>宿主机</label>
				</div>
				<div class="col-sm-9">
					{{data.Host}}
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>cpu</label>
				</div>
				<div class="col-sm-9">
					{{data.Cpu}}
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>内存</label>
				</div>
				<div class="col-sm-9">
					{{ data.Mem }}&nbspG
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>系统</label>
				</div>
				<div class="col-sm-9">
					{{ data.Os }}
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>创建者</label>
				</div>
				<div class="col-sm-9">
					{{ data.Owner }}
				</div>
    		</div>
	 		<div class="col-sm-12 form-group">
				<div class="col-sm-3">
        			<label>备注</label>
				</div>
				<div class="col-sm-9">
			a		{{ data.Comment }}
				</div>
    		</div>
	 		<div class="col-sm-12 form-group" style="margin-top:20px">
				<div class="col-sm-3">
        			<button class="btn btn-primary btn-sm" @click="migratevm(data.Uuid)">迁移</button>
				</div>
				<div class="col-sm-9">
				    <select class="form-select col-sm-10" v-model="hostvalue">
                        <option  v-for="h in host"  :value="h.Ipv4" >
                            {{h.Ipv4}}&nbsp(cpu:{{h.Usedcpu}}/{{h.Cpu}}， &nbsp mem:{{h.Usedmem}}/{{h.Mem}}， &nbsp 数量:{{h.count}}/{{h.Max_vms}})
                        </option>
                    </select>
				</div>
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
			hostvalue: "",
			data: {},
			host: [],
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
        migratevm: function (uuid) {
            var apiurl = `/api/vm/migratevm`
            this.$http.get(apiurl, { params: { uuid: uuid , migratehost: this.hostvalue} } ).then(response => {
            	if (response.data.res) {
						alert("迁移失败("+response.data.res.Message+")")
					} else {
						alert("迁移成功")
					}
            })
        },
        getvminfo: function () {
			var uuid = this.$route.query.uuid
            var apiurl = `/api/vm/getvminfo`
            this.$http.get(apiurl, { params: { uuid: uuid} } ).then(response => {
            	this.data = response.data.res
            })
        },
        gethostinfo: function () {
			var host = this.$route.query.host
            var apiurl = `/api/vm/gethostinfo`
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

label {
    font-weight : 500;
    margin-top: 0px;
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
    padding: 70px 0px 70px 0px;
    margin-left: 0px;
    margin-TOP: 50px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}
</style>

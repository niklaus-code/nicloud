<template>
<div>
      <div class="col-sm-12 form-group" style="border-bottom: 1px green solid">
                <h4>重置镜像</h4>
            </div>

        <div class="col-sm-8 col-sm-offset-2 choose">
            <div class="col-sm-12" >
                <div class="col-sm-2" >
                    <label>IP&nbsp:</label>
                </div>
                <div class="col-sm-10" >
                    {{ip}}
                </div>
            </div>
            <div class="col-sm-12" >
                <div class="col-sm-2" >
                    <label>UUID&nbsp:</label>
                </div>
                <div class="col-sm-10" >
                    {{uuid}}
                </div>
            </div>
            <div class="col-sm-12" >
                <div class="col-sm-2" >
                    <label>备注&nbsp:</label>
                </div>
                <div class="col-sm-10" >
                    {{comment}}
                </div>
            </div>
            <div class="col-sm-12" >
                <div class="col-sm-2" >
                    <label>存储集群&nbsp:</label>
                </div>
                <div class="col-sm-10" >
                    {{storage}}
                </div>
            </div>
            <div class="col-sm-12" >
                <div class="col-sm-2" >
                    <label>数据中心&nbsp:</label>
                </div>
                <div class="col-sm-10" >
                    {{datacenter}}
                </div>
            </div>
        </div>
		<div class="col-sm-8 col-sm-offset-2 choose"  style="margin-top:30px; margin-bottom:30px" >
            <div class="col-sm-12">
                <div class="col-sm-2">
                    <h5>选择镜像</h5>
                </div>
                <div class="col-sm-3" >
                    <div class="form-group">
                        <select class="col-sm-12" v-model="osimagevalue">
                            <option value="">--请选择--</option>
                            <option  v-for="c in osimage" :value="c.Id">
                                {{ c.Osname}}
                            </option>
                        </select>
                    </div>
                </div>
		    </div>
            <div class="col-sm-2 col-sm-offset-2" style="margin-top:5px">
                <button @click="restore" type="button" class="btn btn-success btn-xs">提交</button>
		        </div>
		    </div>
		</div>
</div>
</template>
<script>
export default {
    data () {
        return {
            osimage: [],
            osimagevalue: "",
            uuid: "",
            ip: "",
            host: "",
            datacenter: "",
            storage: "",
            comment: "",
        }
    },
	
    created: function () {
        this.vminfo()
        this.getimg()
    },


    methods: {
        restore: function () {
            var apiurl = `/api/vm/rebuild`
            this.$http.get(apiurl, { params: { uuid: this.uuid, datacenter: this.datacenter, storage: this.storage, osname: this.osimagevalue, host: this.host}}).then(response => {
                if (response.data.err == null) {
                    alert("重置成功")
                } else {    
                    alert(response.data.err.Message)
                }
            })
        },

        getimg: function () {
            var apiurl = `/api/osimage/getimageby`
            this.$http.get(apiurl, { params: {datacenter: this.datacenter, storage: this.storage} }).then(response => {
				if (response.data.err === null) {
                    this.osimage = response.data.res
					} else {
					alert("获取镜像失败'(" + response.data.err.Message+"')")
					}
                
				})
            },

		vminfo: function () {
			var v = this.$store.state.changeparam.uuid
			if (v === null || typeof v === 'undefined' || v === '' || v === "undefined") {
			 	this.uuid = sessionStorage.getItem('uuid')
			 	this.ip = sessionStorage.getItem('ip')
			 	this.os = sessionStorage.getItem('os')
			 	this.host = sessionStorage.getItem('host')
			 	this.datacenter = sessionStorage.getItem('datacenter')
			 	this.storage = sessionStorage.getItem('storage')
			 	this.owner = sessionStorage.getItem('owner')
			 	this.comment = sessionStorage.getItem('comment')
				} else {
				    this.uuid = this.$store.state.changeparam.uuid
				    this.ip = this.$store.state.changeparam.ip
				    this.os = this.$store.state.changeparam.os
				    this.host = this.$store.state.changeparam.host
			 	    this.datacenter = this.$store.state.changeparam.datacenter
			 	    this.storage = this.$store.state.changeparam.storage
				    this.owner = this.$store.state.changeparam.owner
				    this.comment = this.$store.state.changeparam.comment
			 	    sessionStorage.setItem('uuid', this.$store.state.changeparam.uuid)
			 	    sessionStorage.setItem('ip', this.$store.state.changeparam.ip)
			 	    sessionStorage.setItem('os', this.$store.state.changeparam.os)
			 	    sessionStorage.setItem('host', this.$store.state.changeparam.host)
			 	    sessionStorage.setItem('datacenter', this.$store.state.changeparam.datacenter)
			 	    sessionStorage.setItem('storage', this.$store.state.changeparam.storage)
			 	    sessionStorage.setItem('owner', this.$store.state.changeparam.owner)
			 	    sessionStorage.setItem('comment', this.$store.state.changeparam.comment)
				}
			},
        }
  }
</script>
<style scoped>

.createip {
	font-weight:500
}

.vlaninfo {
	font-weight:501
}

.choose {
	padding: 10px;
	border-style: solid;
	border-color: #ddd;
	border-width: 1px;
	border-radius: 4px 4px 0 0;
}

h5 {
    float: right
}

.col-sm-2 label {
    float: right
}

.col-sm-3 {
    padding-left: 0px
}

.col-sm-10 {
    padding-left: 0px
}

select{
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    height: 30px;
    border-radius: 5px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

th {
text-align: center
}

.info {
    border-bottom: 1px solid #ccc
}

.form-control {
    height:30px;
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
}
</style>

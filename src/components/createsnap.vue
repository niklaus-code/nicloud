<template>
<div>
      <div class="col-sm-12 form-group" style="border-bottom: 1px green solid">
                <h4>云主机 & 创建快照</h4>
            </div>

        <div class="col-sm-10 col-sm-offset-1 choose">
            <div class="col-sm-12" >
                <div class="col-sm-2" >
                    <label>IP:</label>
                </div>
                <div class="col-sm-10" >
                    {{ip}}
                </div>
            </div>
            <div class="col-sm-12" >
                <div class="col-sm-2" >
                    <label>UUID:</label>
                </div>
                <div class="col-sm-10" >
                    {{uuid}}
                </div>
            </div>
            <div class="col-sm-12" >
                <div class="col-sm-2" >
                    <label>备注:</label>
                </div>
                <div class="col-sm-10" >
                    {{comment}}
                </div>
            </div>
            <div class="col-sm-12" >
                <div class="col-sm-2" >
                    <label>存储集群:</label>
                </div>
                <div class="col-sm-10" >
                    {{storage}}
                </div>
            </div>
            <div class="col-sm-12" >
                <div class="col-sm-2" >
                    <label>数据中心:</label>
                </div>
                <div class="col-sm-10" >
                    {{datacenter}}
                </div>
            </div>
            <div class="col-sm-12" style="margin-top:30px">
                <table class="table table-bordered" style="text-align: center" v-if="lensnap">
                    <thead>
                        <tr>
                            <th>快照列表</th>
                            <th>创建时间</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="c in snap">
                            <td>{{c.Snap}}</td>
                            <td>{{c.Create_time}}</td>
                            <td>
                                <button type="button" class="btn btn-success btn-xs" @click="rollback(c.Snap)">从此快照恢复</button>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
		<div class="col-sm-10 col-sm-offset-1 choose"  style="margin-top:30px; margin-bottom:30px" >
            <div class="col-sm-12" >
                    <div class="col-sm-2">
                        <h5>输入名称</h5>
                    </div>
                    <div class="col-sm-4">
                        <form role="form">
                            <div class="form-group">
                                <input type="text" class="form-control" v-model="snapvalue" placeholder="">
                            </div>
                        </form>
                    </div>
              </div>
            <div class="col-sm-12" >
                <div class="col-sm-2 col-sm-offset-2">
                    <button @click="createsnap" type="button" class="btn btn-success btn-xs">创建快照</button>
		        </div>
		    </div>
		</div>
</div>
</template>
<script>
export default {
    data () {
        return {
            snap: [],
            lensnap: 0,
            uuid: "",
            ip: "",
            host: "",
            datacenter: "",
            storage: "",
            comment: "",
            snapvalue: ""
        }
    },
	
    created: function () {
        this.vminfo()
        this.getsnap()
    },


    methods: {
        rollback: function (snapname) {
            alert("恢复快照，需要联系管理员手动操作 =^_^=")
            return
            var apiurl = `/api/vm/rollback`
            this.$http.get(apiurl, { params: {uuid: this.uuid, datacenter: this.datacenter, storage: this.storage, snapname: snapname} }).then(response => {
                if (response.data.err === null) {
                    alert("回滚成功")
                    } else {
                    alert("回滚失败'(" + response.data.err.Message+"')")
                    }
                
                })
            },

        createsnap: function () {
            var apiurl = `/api/vm/createsnap`
            this.$http.post(apiurl, this.$qs.stringify({ uuid: this.uuid, datacenter: this.datacenter, storage: this.storage, snapname: this.snapvalue})).then(response => {
                if (response.data.err == null) {
                    alert("创建成功")
                } else {    
                    alert(response.data.err.Message)
                }
            })
        },

        getsnap: function () {
            var apiurl = `/api/vm/getsnap`
            this.$http.get(apiurl, { params: {uuid: this.uuid, datacenter: this.datacenter, storage: this.storage} }).then(response => {
				if (response.data.err === null) {
                    this.snap = response.data.res
                    this.lensnap = response.data.res.length
					} else {
					alert("获取镜像快照失败'(" + response.data.err.Message+"')")
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
.col-sm-2 {
	padding-left:0;
}

.choose {
	padding: 10px;
	border-style: solid;
	border-color: #ddd;
	border-width: 1px;
	border-radius: 4px 4px 0 0;
}

.col-sm-2 label {
    float: right
}

.col-sm-2 h5 {
    float: right
}

.col-sm-10  {
    padding-left: 0px;
}

.col-sm-4  {
    padding-left: 0px;
}

select{
    font-family: "微软雅黑";
    border: 1px #ccc solid;
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
       height: 30px;
    }
</style>

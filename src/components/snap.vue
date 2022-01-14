<template>
<div>
      <div class="col-sm-12 form-group" style="border-bottom: 1px green solid">
                <h4>云主机 & 快照</h4>
            </div>

        <div class="col-sm-8 col-sm-offset-2 choose">
                <ul>
                    <li><strong>uuid</strong>:&nbsp&nbsp{{uuid}}</li>
                    <li><strong>数据中心</strong>:&nbsp&nbsp{{datacenter}}</li>
                    <li><strong>存储</strong>:&nbsp&nbsp{{storage}}</li>
                    <li><strong>备注</strong>:&nbsp&nbsp{{comment}}</li>
                </ul>
            <table class="table table-bordered" style="text-align: center; margin-top:30px">
                <thead>
                    <tr>
                        <th>快照名称</th>
                        <th>创建时间</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="c in snap">
                        <td>{{c.Snap}}</td>
                        <td>{{c.Create_time}}</td>
                    </tr>
                </tbody>
            </table>
        </div>
		<div class="col-sm-8 col-sm-offset-2 choose"  style="margin-top:30px; margin-bottom:30px" >
            <div class="col-sm-8">
                <div class="col-sm-2">从快照恢复</div>
                <div class="col-sm-9">
                    <select class="col-sm-12" v-model="snapvalue">
                        <option value="">--请选择--</option>
                        <option  v-for="s in snap" :value="s.Snap">
                            {{s.Snap}}
                        </option>
                    </select>
		        </div>
                <div class="col-sm-9 col-sm-offset-2">
                    <h5 style="color: #C0C0C0; margin-top:10px">*回滚快照时间较长，建议联系管理员操作</h5>
		        </div>
                <div class="col-sm-2 col-sm-offset-2">
                    <button @click="rollback" type="button" class="btn btn-success btn-xs">提交</button>
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
        getsnap: function () {
            var apiurl = `/api/vm/getsnap`
            this.$http.get(apiurl, { params: {uuid: this.uuid, datacenter: this.datacenter, storage: this.storage} }).then(response => {
				if (response.data.err === null) {
                    this.snap = response.data.res
					} else {
					alert("获取镜像快照失败'(" + response.data.err.Message+"')")
					}
                
				})
            },

        rollback: function () {
            alert("回滚快照请联系管理员操作 (*＾-＾*)")
            return
            var apiurl = `/api/vm/rollback`
            this.$http.get(apiurl, { params: {uuid: this.uuid, datacenter: this.datacenter, storage: this.storage, snapname: this.snapvalue} }).then(response => {
				if (response.data.err === null) {
					alert("回滚成功")
					} else {
					alert("回滚失败'(" + response.data.err.Message+"')")
					}
                
				})
            },
		create: function () {
			if (typeof this.startip === 'undefined' || this.startip === null || this.startip === ''|| typeof this.endip === 'undefined' || this.endip === null || this.endip === '') {
				alert("输入为空")
				return
				}

            var apiurl = `/api/networks/createip`
            this.$http.get(apiurl, { params: {startip: this.startip, endip: this.endip, vlan: this.vlan} }).then(response => {
				if (response.data.res === null) {
					alert("创建成功")
					} else {
					alert(response.data.res.Message)
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
			 	    this.datacenter = sessionStorage.getItem('datacenter')
			 	    this.storage = sessionStorage.getItem('storage')
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

.col-sm-6 {
}

.startip {
	margin-top: 10px;
}

.endip {
	margin-top: 10px;
}

.col-sm-4 label{
	float: right;
}
select{
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
    height: 30px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
	font-weight : 400;
	margin-top: 5px;
}

th {
text-align: center
}
.info {
    border-bottom: 1px solid #ccc
}

.col-sm-9 {
    padding-left:0px;
}
</style>

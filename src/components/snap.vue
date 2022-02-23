<template>
<div>
    <div class="col-sm-12 form-group" style="border-bottom: 1px green solid; margin-top: 20px">
        <div style="width: 200px; float: left;">
            <h3>云主机 & 快照管理</h3>
        </div>
        <div style="margin-left: 30px; width: 160px; float: left; color: red">
            <h5>*删除快照为永久性删除</h5>
        </div>
        <div style="margin-left: 5px; width: 160px; float: left; color: red">
            <h5>*有子镜像快照无法删除</h5>
        </div>
        <div style="margin-left: 5px; width: 180px; float: left; color: red">
            <h5>*恢复快照需联系管理员操作</h5>
        </div>
    </div>

    <div class="col-sm-10 col-sm-offset-1 choose" >
        <div class="col-sm-12" >
            <div class="col-sm-2" >
                <label>IP :</label>
            </div>
            <div class="col-sm-10" >
               {{ip}}
            </div>
        </div>
        <div class="col-sm-12" >
            <div class="col-sm-2" >
                <label>UUID :</label>
            </div>
            <div class="col-sm-10" >
                {{uuid}}
            </div>
        </div>
        <div class="col-sm-12" >
            <div class="col-sm-2" >
                <label>备注 :</label>
            </div>
            <div class="col-sm-10" >
                {{comment}}
            </div>
        </div>
        <div class="col-sm-12" >
            <div class="col-sm-2" >
                <label>存储集群 :</label>
            </div>
            <div class="col-sm-10" >
                {{storage}}
            </div>
        </div>
        <div class="col-sm-12" >
            <div class="col-sm-2" >
                <label>数据中心 :</label>
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
                    <tr v-for="(c, index) in snap" v-if="c.Status">
                        <td>{{c.Snap}}</td>
                        <td>{{c.Create_time}}</td>
                        <td>
                            <button style="border: none; margin-right: 7px" v-if="createwait">
                                <div style="width: 20px; height:20px;background-color: #FFF; margin-bottom: -6px">
                                    <spinner></spinner>
                                </div>
                            </button>
                            <button v-else type="button" class="btn btn-success btn-xs" @click="snaptoimage(c.Snap, true)">以此创建镜像</button>
                            <button type="button" class="btn btn-primary btn-xs" @click="rollback(c.Snap)">以此快照恢复</button>
                            <button type="button" class="btn btn-danger btn-xs" @click="rmsnap(c.Snap, index)">删除快照</button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
	<div class="col-sm-10 col-sm-offset-1 choose"  style="margin-top:30px; margin-bottom:30px" >
        <div class="col-sm-12" >
           <div class="col-sm-2">
               <h5 style="margin-top: 8px">输入快照名称&nbsp</h5>
           </div>
           <div class="col-sm-4">
               <form role="form">
                   <div class="form-group">
                       <input type="text" class="form-control" v-model="snapvalue" placeholder="">
                   </div>
               </form>
           </div>
          </div>
        <div class="col-sm-12" style="padding-bottom: 10px; padding-top: 10px">
            <div class="col-sm-2 col-sm-offset-2">
                <button @click="createsnap(false)" type="button" class="btn btn-success btn-xs">创建</button>
		    </div>
		</div>
	</div>
	<div class="col-sm-10 col-sm-offset-1 "  style="margin-top:30px; margin-bottom:30px" >
        <div class="col-sm-12" style="padding-bottom: 10px; color: red">
	    </div>
	</div>
</div>
</template>
<script>
import spinner from '@/components/spinner'


export default {
    data () {
        return {
            createwait: false,
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


    components: {
        spinner
        },

    methods: {
        rollback: function (snapname) {
            alert("恢复快照，需要联系管理员操作 =^_^=")
            return
            },

        rmsnap: function (snapname, index) {
            var param = {uuid: this.uuid, datacenter: this.datacenter, storage: this.storage, snapname: snapname}
            var apiurl = `/api/vm/delsnap`
            this.$http.delete(apiurl, {params: param} ).then(response => {
                if (response.data.err == null) {
                    alert("删除成功")
                    this.snap[index].Status = false
                } else {    
                    alert(response.data.err.Message)
                }
            })
        },

        snaptoimage: function (snapname, protect) {
            this.createwait = true
            var apiurl = `/api/vm/createsnap`
            this.$http.post(apiurl, this.$qs.stringify({ uuid: this.uuid, datacenter: this.datacenter, storage: this.storage, snapname: snapname, protect: protect})).then(response => {
                if (response.data.err === null) {
                    alert("创建成功")
                    this.getsnap()
                } else {    
                    alert(response.data.err.Message)
                    }
                this.createwait = false
                })
            },

        createsnap: function (protect) {
            var apiurl = `/api/vm/createsnap`
            this.$http.post(apiurl, this.$qs.stringify({ uuid: this.uuid, datacenter: this.datacenter, storage: this.storage, snapname: this.snapvalue, protect: protect})).then(response => {
                if (response.data.err == null) {
                    alert("创建成功")
                    this.getsnap()
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
h3 {
    margin-bottom: 0;
}

h5 {
    margin-top: 29px;
    margin-bottom: 0;
}

.form-group {
    margin-bottom: 0;
}

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
    margin-top: 30px;
    padding-top: 20px;
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

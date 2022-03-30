<template>
<div>
    <div class="col-sm-12 form-group" style="border-bottom: 1px green solid; margin-top: 20px">
        <div style="width: 200px; float: left;">
            <h3>云主机配置管理</h3>
        </div>
    </div>

    <div class="col-sm-10 col-sm-offset-1 choose" >
        <div class="col-sm-12" style="margin-top:30px">
            <table class="table table-bordered" style="text-align: center">
                <thead>
                    <tr>
                        <th>编号</th>
                        <th>CPU（ 核 ）</th>
                        <th>内存（ G ）</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(c, index) in flavorlist">
                        <td>{{index+1}}</td>
                        <td>{{c.Cpu}}</td>
                        <td>{{c.Mem}}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
	<div class="col-sm-10 col-sm-offset-1 choose"  style="margin-top:30px; margin-bottom:30px" >
        <div class="col-sm-12" >
           <div class="col-sm-2">
               <h5 style="margin-top: 8px">CPU&nbsp</h5>
           </div>
           <div class="col-sm-3">
               <form role="form">
                   <div class="form-group">
                       <input type="text" class="form-control" v-model="cpu" placeholder="">
                   </div>
               </form>
           </div>
           <div class="col-sm-2">
               <h5 style="margin-top: 8px">内存&nbsp</h5>
           </div>
           <div class="col-sm-3">
               <form role="form">
                   <div class="form-group">
                       <input type="text" class="form-control" v-model="mem" placeholder="">
                   </div>
               </form>
           </div>
        </div>
        <div class="col-sm-12" style="padding-bottom: 10px; padding-top: 10px">
            <div class="col-sm-2 col-sm-offset-2">
                <button @click="createflavor()" type="button" class="btn btn-success btn-xs">新增配置</button>
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
            flavorlist: [],
            cpu: "",
            mem: ""
        }
    },
	
    created: function () {
        this.getflavor()
    },


    components: {
        },

    methods: {
       getflavor: function () {
            var apiurl = `/api/vm/getflavor`
            this.$http.get(apiurl).then(response => {
            this.flavorlist = response.data.res
            })
        },

       createflavor: function () {
            var apiurl = `/api/vm/createflavor`
            this.$http.get(apiurl, { params: { cpu: this.cpu, mem: this.mem}}).then(response => {
             if (response.data.err === null) {
                this.getflavor()
                } else {
                    alert(response.data.err.Message)
                    }
            })
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

.col-sm-3  {
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

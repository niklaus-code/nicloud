<template>
	<div>
	<nicloudhead></nicloudhead>
	<vmleft></vmleft>
  	<div class="content whisper-content leacots-content details-content col-md-11 col-md-offset-2" style="background-color:white; float:left">
		<div class="col-sm-2 col-sm-offset-4" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-3">
        			<label>cpu</label>
				</div>
				<div class="col-sm-9">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="cpu" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
	 		<div class="form-group">
				<div class="col-sm-3">
        			<label>内存</label>
				</div>
				<div class="col-sm-9">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="mem" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
	 		<div class="form-group">
				<div class="col-sm-3">
        			<label>IP</label>
				</div>
				<div class="col-sm-9">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="ip" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
	 		<div class="form-group">
				<div class="col-sm-3">
        			<label>数量</label>
				</div>
				<div class="col-sm-9">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="num" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
		<div class="form-group" style="margin-top:20px" >
			<div class="col-sm-3 col-sm-offset-3">
  				<button type="submit" @click="createhost" class="btn btn-default btn-sm">提交</button>
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
			cpu: "",
			mem: "",
			ip: "",
			num: "",
        }
    },

    components: {
        foot, nicloudhead, vmleft
    },


    methods: {
		createhost: function () {
            var apiurl = `/api/vm/createhost`

            this.$http.get(apiurl, { params: { cpu: this.cpu, mem:this.mem, ip: this.ip, num: this.num} }).then(response => {
				if (response.data.res) {
					alert("创建成功! 是否查看宿主机列表")
					this.$router.push('/host')
				} else {
					alert("创建失败")
					}
			})
			},

        }
  }
</script>
<style>

select{
    font-family: "微软雅黑";
    border: 1px #1a1a1a solid;
    border-radius: 5px;
}

.content {
    box-shadow: 0 0 10px rgba(0,0,0,8);
    border-radius: 10px/10px;
    z-index: -1;
    padding: 100px 0px 100px 0px;
    margin-left: 0px;
    margin-TOP: 50px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
	font-weight : 400;
	margin-top: 5px;
}
</style>

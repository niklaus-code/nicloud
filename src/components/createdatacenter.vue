<template>
<div>
    <div class="col-sm-12 form-group" style="border-bottom: 1px green solid">
        <h4>创建数据中心</h4>
    </div>
	<div class="col-sm-12">
        <div  class="col-sm-6 col-sm-offset-3">
		    <div class="col-sm-3">
	            <div class="form-group">
        	        <label>数据中心</label>
			    </div>
		    </div>
		    <div class="col-sm-9">
			    <form role="form">
  				    <div class="form-group">
    				    <input type="text" class="form-control" v-model="datacenter" placeholder="">
  				    </div>
			    </form>
		    </div>
        </div>
    </div>
	<div class="col-sm-12">
        <div  class="col-sm-6 col-sm-offset-3">
		    <div class="col-sm-3">
	            <div class="form-group">
        	        <label>备注</label>
			    </div>
		    </div>
		    <div class="col-sm-9">
			    <form role="form">
  				    <div class="form-group">
    				    <input type="text" class="form-control" v-model="comment" placeholder="">
  				    </div>
			    </form>
		    </div>
        </div>
    </div>
	<div class="col-sm-12">
		<div class="form-group">
            <div  class="col-sm-6 col-sm-offset-3">
			    <div class="col-sm-1 col-sm-offset-3" style="margin-top:20px" >
  				    <button type="submit" @click="commit" class="btn btn-success btn-sm">提交</button>
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
           	datacenter: "",
            comment: "",
        }
    },

    methods: {
		commit: function () {
            var apiurl = `/api/datacenter/adddatacenter`

            this.$http.post(apiurl, this.$qs.stringify({datacenter: this.datacenter, comment: this.comment})).then(response => {
				if (response.data.err === null) {
					alert("创建成功! 是否查看数据中心列表")
					this.$emit("toParent", "datacenter");
				} else {
					alert(response.data.err)
					}
			})
			},

        }
  }
</script>

<style scoped>
.form-control {
	height:30px;
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
}

label {
	font-weight : 400;
	margin-top: 5px;
    float: right;
}
</style>

<template>
<div>
    <div class="col-sm-12 form-group" style="margin-top:30px; border-bottom: 1px green solid">
        <h4>修改密码</h4>
    </div>
    <div class="col-sm-6 col-sm-offset-2" style="margin-top:20px">
		<div class="col-sm-10 col-sm-offset-2" style="margin-top:20px">
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>用户名 || EMAIL</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="username" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>原密码</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="oldpasswd" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>新密码</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="newpasswd1" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>重复新密码</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="newpasswd2" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
    		</div>
				<div class="col-sm-12">
		<div class="form-group" style="margin-top:20px" >
			<div class="col-sm-2 col-sm-offset-4">
  				<button type="submit" @click="commit" class="btn btn-success btn-sm">提交</button>
			</div>
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
            username: "",
            oldpasswd: "",
            newpasswd1 : "",
            nwepasswd2 : "",
        }
    },

	mounted: function() {
		},

    methods: {
		commit: function () {            
            if (this.newpasswd1 != this.newpasswd2) {
                alert("2次输入的新密码不一致")
                return
                }

            var apiurl = `/api/user/changepasswd`
            this.$http.post(apiurl,  this.$qs.stringify({username: this.username, oldpasswd: this.oldpasswd, newpasswd1: this.newpasswd1, newpasswd2: this.newpasswd2})).then(response => {
				if (response.data.err === null) {
					alert("修改成功")
					this.$emit("toParent", "login");
				} else {
					alert(response.data.err.Message)
					}
			})
			},

        }
  }
</script>
<style scoped>
.form-control {
	height:30px;
}

.col-sm-4 label{
	float: right;
}
select{
	height:30px;
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
	font-weight : 400;
	margin-top: 5px;
}
</style>

<template>
<div>
      <div class="col-sm-12 form-group" style="border-bottom: 1px green solid">
                <h4>创建用户</h4>
            </div>

		<div class="col-sm-4 col-sm-offset-3" style="margin-top:30px; margin-bottom:30px">
				<div class="col-sm-12" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>用户名*</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="user" placeholder="">
  						</div>
					</form>
				</div>
				</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>密码*</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="passwd" placeholder="">
  						</div>
					</form>
				</div>
				</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>邮箱*</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="email" placeholder="">
  						</div>
					</form>
				</div>
				</div>
    		</div>
				<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>电话*</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="mobile" placeholder="">
  						</div>
					</form>
				</div>
				</div>
    		</div>
			<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>角色*</label>
				</div>
				<div class="col-sm-8">
                        <select class="col-sm-12" v-model="rolevalue">
                            <option value="">--请选择--</option>
                            <option v-for="r in roles" :value="r.Rolename">
                                {{r.Rolename }}
                            </option>
                        </select>
				</div>
			</div>
    	</div>
		<div class="col-sm-12" style="margin-top:20px" >
			<div class="col-sm-2 col-sm-offset-4">
  				<button type="submit" @click="createuser" class="btn btn-success btn-sm">提交</button>
			</div>
		</div>
	</div>
</div>
</template>
<script>
export default {
    data () {
        return {
            rolevalue: "",
            roles: [],
            user: "",
            passwd: "",
            email: "",
            mobile: "",
            role: "",
        }
    },

	mounted: function() {
        this.getroles()
        },

    methods: {
        getroles: function () {
            var apiurl = `/api/user/getroles`
            this.$http.get(apiurl).then(response => {
				if (response.data.err === null) {
                    this.roles = response.data.res
				} else {
					alert(response.data.err.Message)
					}
			})
            },
        
		createuser: function () {
            var apiurl = `/api/user/createuser`

            this.$http.post(apiurl,  this.$qs.stringify({ username: this.user, passwd: this.passwd, email: this.email, mobile: this.mobile, role: this.rolevalue})).then(response => {
				if (response.data.err === null) {
					alert("创建成功!")
				} else {
					alert(1111)
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

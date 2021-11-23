<template>
<div class="login col-sm-12">
	<div class="col-sm-2 col-sm-offset-5" style="margin-top:150px">
		<form class="form-horizontal">
  		<div class="form-group">
    		<div class="col-sm-10">
      			<input v-model="username" class="form-control" placeholder="账号">
    		</div>
  		</div>
  		<div class="form-group">
    		<div class="col-sm-10">
      			<input v-model="passwd" type="password" class="form-control" placeholder="密码">
    		</div>
  		</div>
  		<div class="form-group">
    		<div class="col-sm-10">
      			<input type="button" @click="login" value="登陆">
    		</div>
  		</div>
		</form>
  </div>
</div>
</template>
<script>
export default {
    data () {
        return {
            username: "",
            passwd: "",
        }
    },
    methods: {
		login: function (index, uuid) {
            var apiurl = `/api/user/login`
            this.$http.post(apiurl, this.$qs.stringify({ username: this.username, passwd: this.passwd} )).then(response => {
					
				if (response.data.err === null ) {
					alert("登陆成功")
					this.$store.state.token = response.data.res
					this.$store.commit('set_token', response.data.res);
					this.$router.push({name:"nicloud"})
					} else {
					alert("登陆失败")
						}
            	})
            },
	}
}

</script>
<style scoped>
.login {
	position:fixed;
	height:100%;
	background-color:#ffffff;
	width:100%
	}

input {
	width: 100%;
	font-family: "微软雅黑";
}

</style>

<template>
<div class="login col-sm-12" :style="login">
	<div class="col-sm-4 col-sm-offset-4" style="margin-top:250px; display: flex; 　align-items:center;">
		<form class="form-horizontal col-sm-12">
			<div class="col-sm-10 col-sm-offset-1 a">
			<div class="col-sm-8 col-sm-offset-2 ">
  				<div class="form-group">
    				<div class="col-sm-12">
      					<input v-model="username" class="form-control passwd" :style="userimage" placeholder="账号">
    				</div>
  				</div>
  				<div class="form-group">
    				<div class="col-sm-12">
      					<input v-model="passwd" type="password" :style="passimage" class="form-control passwd" placeholder="密码">
    				</div>
  				</div>
  				<div class="form-group">
    				<div class="col-sm-12">
      					<input type="button" @click="login" value="登陆">
    				</div>
  				</div>
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
			login: {
			backgroundImage:"url(" + require("./fish.jpg") + ")",
        	},
			passimage: {
				backgroundImage:"url(" + require("./pass.svg") + ")",
        	},
			userimage: {
				backgroundImage:"url(" + require("./user.svg") + ")",
        	}
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
.passwd {
	padding: 0 10px 0 35px;
	background-position: 10px 8px;
	background-repeat: no-repeat;
}

.login {
	position:fixed;
	height:100%;
	#background-color:#ffffff;
	width:100%
	}

.a {
	background-color:rgba(255,255,255,0.30);
}

.glyphicon {
	display: inline-block;
	 background-repeat: no-repeat;
	 top: 24%;
	z-index: 2;
}

input {
	width: 100%;
	font-family: "微软雅黑";
}

.col-sm-8 {
	margin-top:30px;
	margin-bottom:10px;
}
</style>

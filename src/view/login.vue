<template>
<div class="logincss" >
	<div style="padding-top:150px; max-width: 400px; margin: 0 auto;">
		<form class="form-horizontal col-sm-12">

			<div class="col-sm-6 col-sm-offset-3"> 
                <img src="./logo.png" width="100%">
            </div>

			<div class="col-sm-12 login">
  				<div class="form-group">
    				<div class="col-sm-12">
      					<input v-model="username" class="form-control passwd" :style="userimage" placeholder="账号">
    				</div>
  				</div>
  				<div class="form-group" style="margin-bottom:5px;">
    				<div class="col-sm-12">
      					<input v-model="passwd" type="password" :style="passimage" class="form-control passwd" placeholder="密码">
    				</div>
  				</div>
                <p style="margin-left:3px; color: #59BF74; font-weight:300; font-size:13px">忘记密码？<p>
  				<div class="form-group">
    				<div class="col-sm-12">
      					<input type="button" style="background-color: #59BF74" @click="loginfun" value="登陆">
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
			loginbkimage: {
				backgroundImage:"url(" + require("./fish.jpg") + ")",
        	},
			passimage: {
				backgroundImage:"url(" + require("./key-fill.svg") + ")",
        	},
			userimage: {
				backgroundImage:"url(" + require("./user.svg") + ")",
        	},
		}	
    },

    mounted: function () {
        let self = this
        document.onkeydown = function(e) {
            let ev = document.all ? window.event : e
            if (ev.keyCode === 13) {
                self.loginfun()
                }
            }
        },

    methods: {
		loginfun: function () {
            var apiurl = `/api/user/login`
            this.$http.post(apiurl, this.$qs.stringify({ username: this.username, passwd: this.passwd} )).then(response => {
					
				if (response.data.err === null ) {
					this.$store.state.username = response.data.username
					this.$store.commit('set_token', response.data.token);
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

.logincss {
	position:fixed;
	height:100%;
	#background-color:#ffffff;
	width:100%
	}

.glyphicon {
	display: inline-block;
	 background-repeat: no-repeat;
	 top: 24%;
	z-index: 2;
}

input {
    font-weight: 500px;
	width: 100%;
	font-family: "微软雅黑";
}

.login {
	border-radius: 15px;
	background-color:rgba(255,255,255,0.30);
    font-weight:900;
	margin-top:30px;
	margin-bottom:10px;
}
</style>

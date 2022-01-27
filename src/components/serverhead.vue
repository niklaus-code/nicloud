<template>
<div class="abc col-sm-12">
    	<div class="col-sm-3">
            <strong>
				<span @click="nicloud">NICLOUD&nbsp/</span>
				<span @click="serverlist">SERVEROOM</span>
            </strong>
		</div>
	</ul>
</div>

</template>
<script>
export default {
    data () {
        return {
            username: "",
            charge: "",
            active: "",
        }
    },
       created: function () {
        this.getuser()
    },
    methods: {
        nicloud: function () {
            sessionStorage.setItem('router', "vm")
            this.$router.push('nicloud')
            },

        serverlist: function () {
            this.$emit("toParent", "server");
            },

        index: function () {
            this.$emit("toParent", "vm");
            },

        getuser: function () {
            var u = this.$store.state.username
            if (u === null || typeof u === 'undefined' || u === '' || u === "undefined") {
                this.username = sessionStorage.getItem('username')
            } else {
                sessionStorage.setItem('username', this.$store.state.username)
                this.username =  this.$store.state.username
                }
            },
        out: function () {
            sessionStorage.removeItem("token");
            this.$router.push({name:"login"});
            },

        leave: function () {
            this.active = "color: #FFF";
            },
        
        mouseOver: function () {
            this.active = "color: #3090C7";
            },
        },
    }

</script>

<style scoped>
.outstyle {
    background-color: red;
}

a {
	color: #FFF;
}

.abc {
	color: #FFF;
    padding-top: 6px;
    border-top-left-radius: 4px;
    border-top-right-radius: 4px;
    border-bottom-right-radius: 4px;
    border-bottom-left-radius: 4px;
    font-size: 15px;
	background-color: #5B5B5B;
}
</style>

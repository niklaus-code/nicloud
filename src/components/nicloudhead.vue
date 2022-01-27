<template>
<div style="padding-bottom: 5px">
   	<div class="col-sm-3">
        <strong>
		<span @click="index">NICLOUD&nbsp/</span>
		<span @click="serveroom">SERVEROOM</span>
        </strong>
	</div>
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
        serveroom: function () {
            sessionStorage.setItem('router', "server")

            this.$router.push('serveroom')
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
a {
	color: #FFF;
}

</style>

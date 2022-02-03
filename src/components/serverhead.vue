<template>
<div style="padding-left:10px; background-color: #5b5b5b; color:#FFA500;padding-top:6px; padding-bottom: 6px; ">
        <strong>
		<span @click="index">NICLOUD&nbsp/</span>
		<span @click="serveroom">SERVEROOM</span>
        </strong>
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
             this.$emit("toParent", "server");
            },
        index: function () {
            sessionStorage.setItem('router', "vm")
            this.$router.push('nicloud')
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
	color: #FFA500;
}

</style>

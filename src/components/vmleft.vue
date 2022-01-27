<template>
    <div>
        <div style="margin-top:50px ;text-shadow: black 1px;">
            <strong class="glyphicon glyphicon-user">{{username}} </strong> | <span @click="logout">Logout</span>
            
        </div>
        <div style="padding-left: 20px;">
            <ul class="list-group">
                <li class="list-group-item" v-for="(item, index) in routelist" @click="choose(index)" @click="toParent(item.router)">
                    <p :class=item.class> </p><p style="margin-left: 5px">{{item.name}}</p>
                </li>
            </ul>
        </div>
    </div>
</template>

<script>

export default {
    data () {
        return {
            username: "",
			selected: 0,
			routelist: [
				{
				name: "云主机",
				router: "vm",
                class: "glyphicon glyphicon-th-large"
					},
				{
				name: "云盘",
				router: "disk",
                class: "glyphicon glyphicon-hdd"
					},
				{
				name: "系统镜像",
				router: "osimage",
                class: "glyphicon glyphicon-paperclip"
					},
				{
				name: "网络",
				router: "network",
                class: "glyphicon glyphicon-plane"
					},
				{
				name: "宿主机",
				router: "hosts",
                class: "glyphicon glyphicon-home"
					},
				{
				name: "存储集群",
				router: "storage",
                class: "glyphicon glyphicon-list-alt"
					},
				{
				name: "数据中心",
				router: "datacenter",
                class: "glyphicon glyphicon-globe"
					},
				{
				name: "统计",
				router: "count",
                class: "glyphicon glyphicon-eye-open"
					},
				],
            }
        },

    created: function () {
        this.getuser()
        },

	methods: {
        logout: function () {
            sessionStorage.removeItem("token");
            this.$router.push({name:"login"});
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

       	choose(index){
			this.selected = index;
            },

       	toParent: function (item) {
            if (item === "count") {
                alert("暂未开放")
                return
                }
		    this.$emit("toParent", item);
            },
		}
}
</script>

<style scoped>

li {
	text-align: left 
}

.list-group-item {
    background-color: #778899;
	border: 0px;
    color: white;
}

.list-group-item {
    padding-right: 0px;
    padding-left: 0px;
}

.col-md-1 {
    padding-left: 5px;
    padding-right: 5px;
}

.list-group {
    margin-top:100px;
	}

span {
	font-weight: 600;
}

a{
	color: white;
}

p {
    display:inline-block;
}
</style>

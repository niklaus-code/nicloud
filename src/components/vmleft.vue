<template>
    <div  class="col-md-1" style="float:left">
        <ul class="list-group col-md-11 col-md-offset-1">
            <li class="list-group-item" v-for="(item, index) in routelist" v-bind:class="{aaa: is === index}">
                    <span  @click="toParent(item.router)" >{{item.name}}</span>
            </li>
        </ul>
    </div>
</template>

<script>

export default {
    data () {
        return {
			routelist: [
				{
				name: "云主机",
				router: "vm"
					},
				{
				name: "云盘",
				router: "disk"
					},
				{
				name: "镜像",
				router: "osimage"
					},
				{
				name: "网络",
				router: "network"
					},
				{
				name: "宿主机",
				router: "hosts"
					},
				{
				name: "存储集群",
				router: "storage"
					},
				{
				name: "统计",
				router: "count"
					},
				],
        }
    },

	methods: {
        toParent: function (item) {
			this.$emit("toParent", item);
            },
	}
}
</script>

<style scoped>
li {
	text-align: center 
}

.aaa {
	background-color: red;
}

.list-group {
    margin-top:100px;
	}

span {
	font-weight: 600;
}

a{
	color: black;
}

.list-group-item {
    background-color: transparent;
    border: none;
}

.breadcrumb {
    background-color: #5B5B5B
}

</style>

<template>
    <div  class="col-md-1" style="float:left">
        <ul class="list-group col-md-11 col-md-offset-1">
            <li class="list-group-item" v-for="(item, index) in routelist" :class='{bg:index==isactive}'>
                <router-link :to="{name: item.router}">
                    <span @click="mouseOver(index)" >{{item.name}}</span>
                </router-link>
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
				router: "nicloud"
					},
				{
				name: "云盘",
				router: "nicloud"
					},
				{
				name: "镜像",
				router: "os"
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
				name: "ceph",
				router: "ceph"
					},
				{
				name: "统计",
				router: "count"
					},
				],
			isactive:8,
        }
    },

	methods: {
        mouseOver: function (index) {
            this.isactive = index;
            },

        mouseLeave: function (index) {
            this.isactive = index;
            }
	}
}
</script>

<style scoped>
li {
	text-align: center 
}

.bg {
	background-color: red;
}

.list-group {
    margin-top:100px;
	}

.list-group-item {
    background-color: transparent;
    border: none;
    color: blue;
}

.breadcrumb {
    background-color: #5B5B5B
}
span {
	font-weight: 600;
}

</style>

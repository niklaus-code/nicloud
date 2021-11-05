<template>
    <div  class="col-md-1" style="float:left">
        <ul class="list-group col-md-11 col-md-offset-1">
            <li class="list-group-item" v-for="(item, index) in routelist"  @click="up(index)"  @click="route(item.router)" v-bind:class="{aaa: is === index}">
                    <span>{{item.name}}</span>
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
				router: "cloudrive"
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
				router: "ceph"
					},
				{
				name: "统计",
				router: "count"
					},
				],
			is: 8,
        }
    },

	methods: {
        route: function (label) {
            this.$router.push({
            	path: label,
            }) 
            },

        up: function (index) {
            this.is = index;
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

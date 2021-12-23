<template>
    <div  class="col-md-1">
        <ul class="list-group col-md-11 col-md-offset-1">
            <li class="list-group-item" v-for="(item, index) in routelist" :class="{ sss: selected === index }" @click="choose(index)" @click="toParent(item.router)">
                    <span >{{item.name}}</span>
            </li>
        </ul>
    </div>
</template>

<script>

export default {
    data () {
        return {
			selected: 0,
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
				name: "系统镜像",
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
				name: "数据中心",
				router: "datacenter"
					},
				{
				name: "统计",
					},
				],
        }
    },

	methods:{
        	choose(index){
				this.selected = index;
              },

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
.list-group-item {
	background-color: #e3e3e3;
	border: 0px;
}

.sss {
	#background-color: #fff;
	color: red;
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
	color: black;
}

</style>

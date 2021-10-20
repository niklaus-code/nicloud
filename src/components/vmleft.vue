<template>
    <div  class="col-md-1" style="float:left">
        <ul class="list-group col-md-11 col-md-offset-1">
            <li class="list-group-item" @mouseover="mouseOver(index)" @mouseleave="mouseLeave()">
                <router-link :to="{name:'nicloud'}">
                    <span>云主机</span>
                </router-link>
            </li>
            <li class="list-group-item" @mouseover="mouseOver(index)" @mouseleave="mouseLeave(index)">
                <router-link :to="{name:'nicloud'}">
                <span>云硬盘</span>
                </router-link>
            </li>
            <li class="list-group-item" @mouseover="mouseOver(index)" @mouseleave="mouseLeave(index)">
                <router-link :to="{name:'nicloud'}">
                <span>镜像</span>
                </router-link>
            </li>
            <li class="list-group-item" @mouseover="mouseOver(index)" @mouseleave="mouseLeave(index)">
                <router-link :to="{name:'nicloud'}">
                <span>网络</span>
                </router-link>
            </li>
            <li class="list-group-item" @mouseover="mouseOver(index)" @mouseleave="mouseLeave(index)">
                <router-link :to="{name:'nicloud'}">
                <span>统计</span>
                </router-link>
            </li>
            <li class="list-group-item" @mouseover="mouseOver(index)" @mouseleave="mouseLeave(index)">
                <router-link :to="{name:'createhost'}">
                <span>宿主机</span>
                </router-link>
            </li>
            <li class="list-group-item" @mouseover="mouseOver(index)" @mouseleave="mouseLeave(index)">
                <router-link :to="{name:'nicloud'}">
                <span>ceph</span>
                </router-link>
            </li>
        </ul>
    </div>
</template>

<script>

export default {
    data () {
        return {
            active: "",
            checkvalue: false,
            content: "",
            stat: {
                "运行": "btn btn-success btn-xs", 
                "关机": "btn btn-warning btn-xs", 
                },
            statclass: "btn btn-danger",
            data: "",
            status: {
                0: "关机",
                1: "运行",
                2: "已删除",
            }
        }
    },

	methods: {
        mouseOver: function (index) {
            this.daohang[index].active = "background-color: red";
            },

        mouseLeave: function (index) {
            this.daohang[index].active = "background-color: #e3e3e3";
            }
	}
}
</script>

<style scoped>
	.list-group {
    margin-top:100px;
	}

.list-group-item {
    background-color: #e3e3e3;
    border: none;
    color: blue;
}

.breadcrumb {
    background-color: #5B5B5B
}

</style>

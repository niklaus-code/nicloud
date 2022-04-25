<template>
<div>
    <div style="margin-top: 35px">
         <img src="./logo.jpg" width="100%">
    </div>
    <div style="width:140; text-align: center;margin-top: 50px; z-index: 2500">
        <div align="center" style="color: white; text-align: left; margin-left: 30px">
            <ul>
                <li v-for="(item, index) in routelist" @mouseover="mouseOver(index)" @mouseLeave="mouseLeave(index)" @click="toParent(item.router)">
                    <div @click="fun_downmenu(index)">
                    <p :class=item.class> </p>
                    <p class="leftitem"><a :class = "isactive == index ? 'addclass' : '' " href="javascript:void(0)">{{item.name}}</a></p>
                    <p style="padding-top: 12px; float: right; margin-right: 30px"><span v-if="item.check_downmenu_icon" style="top: 1px; height: 6px;width: 6px ;" class="glyphicon glyphicon-triangle-bottom"></span></p>
                    </div>
                    <ul v-if="item.check_downmenu">
                        <li v-for="(item, index) in item.downmenu">
                            <p @click="toParent(item.router)" style="margin-left: 16px;"><a href="javascript:void(0)" style="color: rgb(74, 242, 147);">{{item.name}}</a><p>
                        </li>
                    </ul>
                </li>
            </ul>
        </div>
    </div>
    <div class="bottom" style="">
        <p>Copyright <span class="glyphicon glyphicon-copyright-mark" ></span> 2021</p>
        <p>关于•联系•反馈</p>
    </div>
</div>
</template>

<script>

export default {
    data () {
        return {
            isactive: -1,
            username: "",
			selected: 0,
			routelist: [
				{
                id: 0,
				name: "云主机",
                router: "none",
                class: "glyphicon glyphicon-th-large",
                downmenu: [
                    {
                    "name": "云主机列表",
                    "router": "vm",
                    },
                    {
                    "name": "规格列表",
                    "router": "vm_flavor",
                    },
                    {
                    "name": "归档列表",
                    "router": "vm_archive",
                    },
                    ],
                check_downmenu: false,
                check_downmenu_icon: true
					},
				{
                id: 1,
				name: "云盘",
                router: "none",
                class: "glyphicon glyphicon-hdd",
                downmenu:[
                    {
                    "name": "云盘列表",
                    "router": "disk",
                    },
                    {
                    "name": "归档列表",
                    "router": "vdisk_archive",
                    },
                    ],
                check_downmenu: false,
                check_downmenu_icon: true
					},
				{
                id: 2,
				name: "系统镜像",
				router: "osimage",
                class: "glyphicon glyphicon-modal-window",
                router: "osimage",
					},
				{
                id: 3,
				name: "网络",
				router: "network",
                class: "glyphicon glyphicon-plane",
                downmenu: false,
					},
				{
                id: 4,
				name: "宿主机",
				router: "hosts",
                class: "glyphicon glyphicon-home",
                downmenu: false,
					},
				{
                id: 5,
				name: "存储集群",
				router: "storage",
                class: "glyphicon glyphicon-list-alt",
                downmenu: false,
					},
				{
                id: 6,
				name: "数据中心",
				router: "datacenter",
                class: "glyphicon glyphicon-globe",
                downmenu: false,
					},
				],
            }
        },

    created: function () {
        this.getuser()
        },

	methods: {
        mouseLeave(index) {
            this.isactive = index
        },

        mouseOver(index) {
            this.isactive = index
        },

        fun_downmenu: function (index) {
            if (this.routelist[index].check_downmenu) {
                this.routelist[index].check_downmenu = false
                } else {
                    this.routelist[index].check_downmenu = true
                    }
                },

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

       	toParent: function (item) {
            if (item === "none") {
                return
                } else {
		            this.$emit("toParent", item);
                }
            },
		}
}
</script>

<style scoped>
li {
    font-size: 0.8125rem;
    text-shadow: 1px 1px 0 rgb(0 0 0 / 10%);
    border-bottom: 1px solid rgb(85 110 230 / 5%);
    color: #c9d4f6;
}

td {
    float: left;
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

li p {
    color: #aebbcc;
    font-size: 13px;
    margin-bottom: 6px;
    display:inline-block;
}

.glyphicon {
    top: 2px;
}

a {
    color: #aebbcc;
}

.addclass{
    color : red;
}

.leftitem {
    margin-left: 6px;margin-top:12px;padding-right: 10px
}

.bottom {
    position: absolute;
    bottom: 5px;
    width:100%;
    color: white;
    text-align:center;
}

.bottom p {
    color: #aebbcc;
    margin: 0;
}
</style>

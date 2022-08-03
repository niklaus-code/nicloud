<template>
<div>
    <div style="margin-top: 35px">
        <img src="data:image/png;base64,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"></img>"
    </div>
    <div style="width:140; text-align: center;margin-top: 50px; z-index: 2500">
        <div align="center" style="color: white; text-align: left;">
            <ul>
                <li v-for="(item, index) in routelist" @mouseover="mouseOver(index)" @mouseLeave="mouseLeave(index)" @click="toParent(item.router)">
                    <div @click="fun_downmenu(index)">
                    <p :class=item.class> </p>
                    <p class="leftitem"><a :class = "isactive == index ? 'addclass' : '' " href="javascript:void(0)">{{item.name}}</a></p>
                    <p style="padding-top: 12px; float: right; margin-right: 30px"><span v-if="item.check_downmenu_icon" style="top: 1px; height: 6px;width: 6px ;" :class="item.class_down"></span></p>
                    </div>
                    <ul v-if="item.check_downmenu" style="background-color: #000c20;">
                        <li v-for="(item, index) in item.downmenu">
                            <p v-if="index === 0" @click="toParent(item.router)" style="margin-left: 46px; padding-top: 10px"><a href="javascript:void(0)" style="color: rgb(74, 242, 147);">{{item.name}}</a></p>
                            <p v-else @click="toParent(item.router)" style="margin-left: 46px;"><a href="javascript:void(0)" style="color: rgb(74, 242, 147);">{{item.name}}</a></p>
                        </li>
                    </ul>
                </li>
            </ul>
        </div>
    </div>
    <div class="bottom" style="">
        <p>Copyright <span class="glyphicon glyphicon-copyright-mark" ></span> 2021</p>
        <p>联系•反馈</p>
    </div>
</div>
</template>

<script>

export default {
    data () {
        return {
            class_down: "glyphicon glyphicon-chevron-down",
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
                class_down: "glyphicon glyphicon-chevron-down",
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
                class_down: "glyphicon glyphicon-chevron-down",
                check_downmenu_icon: true
					},
				{
                id: 2,
				name: "系统镜像",
                class: "glyphicon glyphicon-modal-window",
                router: "none",
                check_downmenu: false,
                class_down: "glyphicon glyphicon-chevron-down",
                check_downmenu_icon: true,
                downmenu:[
                    {
                    "name": "镜像列表",
                    "router": "osimage",
                    },
                    {
                    "name": "镜像配置文件xml",
                    "router": "xml",
                    },
                    ]
				},
				{
                id: 3,
				name: "VLAN子网",
				router: "network",
                class: "glyphicon glyphicon-plane",
                downmenu: false,
                check_downmenu: false,
					},
				{
                id: 4,
				name: "宿主机",
				router: "hosts",
                class: "glyphicon glyphicon-home",
                downmenu: false,
                check_downmenu: false,
					},
				{
                id: 5,
				name: "存储集群",
				router: "storage",
                class: "glyphicon glyphicon-list-alt",
                downmenu: false,
                check_downmenu: false,
					},
				{
                id: 6,
				name: "数据中心",
				router: "datacenter",
                class: "glyphicon glyphicon-globe",
                downmenu: false,
                check_downmenu: false,
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
            if (this.routelist[index].class_down === "glyphicon glyphicon-chevron-down") {
                this.routelist[index].class_down = "glyphicon glyphicon-chevron-up"
                } else {
                    this.routelist[index].class_down = "glyphicon glyphicon-chevron-down"
                    }

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
    margin-left: 30px;
    color: #ffffff;
    font-size: 13px;
    margin-bottom: 6px;
    display:inline-block;
}

.glyphicon {
    top: 2px;
}

a {
    color: #ffffff;
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

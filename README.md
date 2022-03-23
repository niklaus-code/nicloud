# NICLOUD简介
**NICLOUD是一个轻量级虚拟机管理平台，后端采用golang gin 框架，前端用了VUE，集成了ceph，libvirtd， dhcp等服务**
**功能包括常用的：**
- 云主机创建删除开关机快照等
- 云主迁移热迁移，热迁移可以在不停止服务的情况下，更换宿主机
- 云盘创建，自定义容量，删除挂载等
- 把系统快照保存为基础镜像
- 归档：删除的云主机，云盘并非真正删除，可以根据需求保留时间，恢复
### 云主机列表页面
![avatar](./static/vm.png)
### 云硬盘列表页面
![avatar](./static/vdisk.png)
### 镜像列表页面
![avatar](./static/os.png)
### 快照详情页
![avatar](./static/os.png)
### 网络列表页面
![avatar](./static/os.png)

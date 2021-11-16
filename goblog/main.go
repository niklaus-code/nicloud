//+build linux
package main

import (
  "github.com/gin-gonic/gin"
  "goblog/apis"
  "goblog/apis/cephapis"
  "goblog/apis/hostapis"
  "goblog/apis/machineapis"
  "goblog/apis/networkapis"
  "goblog/apis/osimage"
  "goblog/apis/vmapis"
  "goblog/apis/datacenterapis"
  "goblog/apis/vdisk"
)

func main() {
  r := gin.Default()

	v1 := r.Group("/api/blog/get_blog")
	{
		v1.GET("/get_blog_read", apis.Get_read)
		v1.GET("/get_blog_thoughts", apis.Get_thoughts)
		v1.GET("/get_blog_by_id/:id", apis.Get_blog_by_id)
		v1.GET("/get_blog/:pagenumber", apis.Get_blog)
	}

	v2 := r.Group("/api/vm")
	{
		v2.GET("getvm", vmapis.Getvmlist)
		v2.GET("create", vmapis.Createvm)
		v2.GET("operation/:id", vmapis.Operation)
		v2.GET("delete", vmapis.DeleteVM)
		v2.GET("getflavor", vmapis.GetFlavor)
    v2.GET("vnc", vmapis.Vnc)
    v2.GET("search", vmapis.Search)
    v2.GET("getstatus", vmapis.GetVmStatus)
    v2.GET("addcomment", vmapis.Addcomment)
    v2.GET("getvminfo", vmapis.GetVminfo)
    v2.GET("migratevm", vmapis.MigrateVm)
	}

  v3 := r.Group("/api/machine")
  {
    v3.GET("getmachinelist", machineapis.Getmachinelist)
    v3.GET("addmachine", machineapis.Addmachine)
    v3.GET("delmachine", machineapis.Delmachine)
    v3.GET("getpage", machineapis.Getpage)
    v3.GET("search", machineapis.Search)
    v3.GET("update", machineapis.Update)
    v3.GET("ping", machineapis.Ping)
  }

  v4 := r.Group("/api/networks")
  {
    v4.GET("createvlan", networkapis.Add)
    v4.GET("getvlan", networkapis.Get)
    v4.GET("getvlanbydatacenter", networkapis.Getvlanbydatacenter)
    v4.GET("getip", networkapis.GetIplist)
    v4.GET("getallip", networkapis.GetallIp)
    v4.GET("createip", networkapis.CreateIp)
    v4.GET("downip", networkapis.DownIp)
    v4.GET("upip", networkapis.UpIp)
    v4.GET("delete", networkapis.Delete)
    v4.GET("deleteip", networkapis.Deleteip)
  }

  v5 := r.Group("/api/hosts")
  {
    v5.GET("delete", hostapis.Delhost)
    v5.GET("gethostinfo", hostapis.Gethostinfo)
    v5.GET("createhost", hostapis.Createhost)
    v5.GET("gethosts", hostapis.GetHosts)
    v5.GET("gethostsbydatacenter", hostapis.GetHostsbydatacenter)
  }

  v6 := r.Group("/api/osimage")
  {
    v6.GET("getimageby", osimage.GetImageby)
    v6.GET("getimage", osimage.GetImage)
    v6.GET("updateimage", osimage.UpdateImage)
    v6.GET("createimage", osimage.AddImage)
    v6.GET("delimage", osimage.DelImage)
  }

  v7 := r.Group("/api/storage")
  {
    v7.GET("get", cephapis.GetStorage)
    v7.GET("add", cephapis.Addceph)
    v7.GET("delete", cephapis.Delete)
    v7.GET("getpool", cephapis.Getpool)

  }

  v8 := r.Group("/api/datacenter")
  {
    v8.GET("getdatacenter", datacenterapis.GetDatacenter)
  }

  v9 := r.Group("/api/vdisk")
  {
    v9.GET("umountdisk", vdisk.Umountdisk)
    v9.GET("mountdisk", vdisk.Mountdisk)
    v9.GET("deletevdisk", vdisk.Deletevdisk)
    v9.GET("createvdisk", vdisk.Createvdisk)
    v9.GET("getvdisk", vdisk.GetVdisk)
  }


  r.Run("127.0.0.1:1992")
}

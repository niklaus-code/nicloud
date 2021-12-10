//+build linux
package main

import (
  "github.com/gin-gonic/gin"
  "nicloud/apis/cephapis"
  "nicloud/apis/datacenterapis"
  "nicloud/apis/hostapis"
  "nicloud/apis/machineapis"
  "nicloud/apis/networkapis"
  "nicloud/apis/osimage"
  "nicloud/apis/userapis"
  "nicloud/apis/vdisk"
  "nicloud/apis/vmapis"
  "nicloud/utils"
)

func main() {
  r := gin.Default()
  v1 := r.Group("/api/user")
  {
    v1.POST("login", userapis.Login)
  }
	v2 := r.Group("/api/vm")
	{
	  v2.Use(utils.Tokenauth())

		v2.GET("getvm", vmapis.Getvmlist)
		v2.POST("create", vmapis.Createvm)
		v2.GET("operation/:id", vmapis.Operation)
		v2.GET("delete", vmapis.DeleteVM)
		v2.GET("getflavor", vmapis.GetFlavor)
    v2.GET("vnc", vmapis.Vnc)
    v2.GET("search", vmapis.Search)
    v2.GET("getstatus", vmapis.GetVmStatus)
    v2.GET("addcomment", vmapis.Addcomment)
    v2.GET("getvminfo", vmapis.GetVminfo)
    v2.GET("migratevm", vmapis.MigrateVm)
    v2.GET("changeconfig", vmapis.Changeconfig)
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
    v2.Use(utils.Tokenauth())
    v4.POST("createvlan", networkapis.Add)
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
    v2.Use(utils.Tokenauth())
    v5.GET("delete", hostapis.Delhost)
    v5.GET("gethostsby", hostapis.Gethostinfo)
    v5.POST("createhost", hostapis.Createhost)
    v5.GET("gethosts", hostapis.GetHosts)
    v5.GET("gethostsbydatacenter", hostapis.GetHostsbydatacenter)
  }

  v6 := r.Group("/api/osimage")
  {
    v2.Use(utils.Tokenauth())
    v6.GET("getimageby", osimage.GetImageby)
    v6.GET("getimage", osimage.GetImage)
    v6.POST("updateimage", osimage.UpdateImage)
    v6.POST("createimage", osimage.AddImage)
    v6.GET("delimage", osimage.DelImage)
  }

  v7 := r.Group("/api/storage")
  {
    v2.Use(utils.Tokenauth())
    v7.GET("get", cephapis.GetStorage)
    v7.POST("add", cephapis.Addceph)
    v7.GET("delete", cephapis.Delete)
    v7.GET("getpool", cephapis.Getpool)

  }

  v8 := r.Group("/api/datacenter")
  {
    v2.Use(utils.Tokenauth())
    v8.GET("getdatacenter", datacenterapis.GetDatacenter)
    v8.POST("adddatacenter", datacenterapis.AddDatacenter)
    v8.GET("deldatacenter", datacenterapis.DelDatacenter)
  }

  v9 := r.Group("/api/vdisk")
  {
    v2.Use(utils.Tokenauth())
    v9.GET("umountdisk", vdisk.Umountdisk)
    v9.GET("mountdisk", vdisk.Mountdisk)
    v9.GET("deletevdisk", vdisk.Deletevdisk)
    v9.POST("createvdisk", vdisk.Createvdisk)
    v9.GET("getvdisk", vdisk.GetVdisk)
  }

  r.Run("127.0.0.1:1992")
}

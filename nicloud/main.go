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

  "github.com/swaggo/gin-swagger"
  "github.com/swaggo/gin-swagger/swaggerFiles"
  _ "nicloud/docs"
)

// @title NILCOUD
// @version 1.0
// @description PRIVATE CLOUD PLATFORM
// @termsOfService https://github.com/niklaus-code/nicloud

// @contact.name NIKLAUS
// @contact.url https://github.com/niklaus-code/nicloud
// @contact.email 1309584951@qq.com
func main() {
  r := gin.Default()

  //utils.Bind(r, userapis.User{})
  r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  v1 := r.Group("/api/user")
  {
    v1.POST("login", userapis.Login)
    v1.Use(utils.Tokenauth())
    v1.Use(utils.RoleAuth())
    v1.GET("getuser", userapis.GetUser)
    v1.POST("createuser", userapis.Createuser)
    v1.GET("getroles", userapis.GetAllRoles)
    v1.GET("deluser", userapis.DelUser)
    v1.POST("changepasswd", userapis.Changepasswd)
  }

	v2 := r.Group("/api/vm")
	{  vmdetails := r.Group("/api/vm/details")
    {
      ////vmdetails.Use(utils.Tokenauth())
      vmdetails.GET("cpuinfo", vmapis.Cpudetails)
      vmdetails.GET("meminfo", vmapis.Memdetails)
      vmdetails.GET("netinfo", vmapis.Netdetails)
      vmdetails.GET("diskinfo", vmapis.Diskdetails)
    }
	  v2.Use(utils.Tokenauth())
		v2.GET("getvm", vmapis.Getvmlist)
		v2.POST("create", vmapis.Createvm)
		v2.GET("operation/:id", vmapis.Operation)
		v2.GET("delete", vmapis.DeleteVM)
		v2.GET("getflavor", vmapis.GetFlavor)
    v2.GET("vnc", vmapis.Vnc)
    v2.GET("search", vmapis.Search)
    v2.GET("migratelive", vmapis.MigrateVmlive)
    v2.GET("getstatus", vmapis.GetVmStatus)
    v2.GET("addcomment", vmapis.Addcomment)
    v2.GET("getvminfo", vmapis.GetVminfo)
    v2.GET("migratevm", vmapis.MigrateVm)
    v2.GET("changeconfig", vmapis.Changeconfig)
    v2.GET("rebuild", vmapis.Rebuild)
    v2.POST("createsnap", vmapis.Createsnap)
    v2.GET("getsnap", vmapis.Getsnap)
    v2.DELETE("delsnap", vmapis.DelSnap)
    v2.GET("rollback", vmapis.Rollback)
    v2.GET("getvmarchive", vmapis.GetVmArchive)
    v2.GET("delvmpermanent", vmapis.Delvmpermanent)
    v2.GET("createflavor", vmapis.CreateFlavor)
    v2.GET("searchvmachives", vmapis.SearchVMArchive)
	}

  v3 := r.Group("/api/machine")
  {
    v3.Use(utils.Tokenauth())
    v3.Use(utils.RoleAuth())
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
    v4.GET("download_excel", networkapis.DownloadExcel)

    v4.Use(utils.Tokenauth())
    v4.GET("getvlan", networkapis.Get)
    v4.GET("getvlanbydatacenter", networkapis.Getvlanbydatacenter)
    v4.GET("getip", networkapis.GetIplist)
    v4.GET("getallip", networkapis.GetallIp)

    v4.Use(utils.RoleAuth())
    v4.POST("createvlan", networkapis.Add)
    v4.GET("createip", networkapis.CreateIp)
    v4.GET("downip", networkapis.DownIp)
    v4.GET("upip", networkapis.UpIp)
    v4.GET("delete", networkapis.Delete)
    v4.GET("deleteip", networkapis.Deleteip)

  }

  v5 := r.Group("/api/hosts")
  {
    v5.Use(utils.Tokenauth())
    v5.GET("gethostsbydatacenter", hostapis.GetHostsbyvlan)
    v5.GET("gethosts", hostapis.GetHosts)
    v5.GET("countdomains", hostapis.ListDomains)
    v5.GET("gethostsby", hostapis.Gethostinfo)
    v5.GET("gethostsbyip", hostapis.Gethostbyip)

    v5.Use(utils.RoleAuth())
    v5.GET("counthosts", hostapis.Counthost)
    v5.GET("delete", hostapis.Delhost)
    v5.POST("createhost", hostapis.Createhost)
    v5.GET("addcomment", hostapis.Addcomment)
    v5.POST("updatehost", hostapis.Updatehostinfo)
  }

  v6 := r.Group("/api/osimage")
  {
    v6.Use(utils.Tokenauth())
    v6.GET("getimagebytag", osimage.GetImagebytag)
    v6.GET("getimageby", osimage.GetImageby)
    v6.GET("getimage", osimage.GetImage)
    v6.GET("getimagesort", osimage.GetImageSort)
    v6.GET("getiostags", osimage.GetOsTag)
    v6.GET("getosimagexml", osimage.Getosimagexml)

    v6.Use(utils.RoleAuth())
    v6.POST("updateimage", osimage.UpdateImage)
    v6.POST("createimage", osimage.AddImage)
    v6.GET("delimage", osimage.DelImage)
    v6.POST("createxml", osimage.Addosimagexml)
    v6.GET("delxml", osimage.Delosimagexml)
  }

  v7 := r.Group("/api/storage")
  {
    v7.Use(utils.Tokenauth())
    v7.GET("get", cephapis.GetStorage)
    v7.GET("getpool", cephapis.Getpool)

    v7.Use(utils.RoleAuth())
    v7.POST("add", cephapis.Addceph)
    v7.GET("delete", cephapis.Delete)
  }

  v8 := r.Group("/api/datacenter")
  {
    v8.Use(utils.Tokenauth())
    v8.GET("getdatacenter", datacenterapis.GetDatacenter)

    v8.Use(utils.RoleAuth())
    v8.POST("adddatacenter", datacenterapis.AddDatacenter)
    v8.GET("deldatacenter", datacenterapis.DelDatacenter)
  }

  v9 := r.Group("/api/vdisk")
  {
    v9.Use(utils.Tokenauth())
    v9.GET("umountdisk", vdisk.Umountdisk)
    v9.GET("mountdisk", vdisk.Mountdisk)
    v9.GET("deletevdisk", vdisk.Deletevdisk)
    v9.POST("createvdisk", vdisk.Createvdisk)
    v9.GET("getvdisk", vdisk.GetVdisk)
    v9.POST("addcomment", vdisk.AddComment)
    v9.GET("getvdiskarchives", vdisk.GetVdiskArchive)
  }

  r.Run("127.0.0.1:1992")
}

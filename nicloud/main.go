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

  u := r.Group("/api/user")
  {
    u.POST("login", userapis.Login)
    u.Use(utils.Tokenauth())
    u.Use(utils.RoleAuth())
    u.GET("getuser", userapis.GetUser)
    u.POST("createuser", userapis.Createuser)
    u.GET("getroles", userapis.GetAllRoles)
    u.GET("deluser", userapis.DelUser)
    u.POST("changepasswd", userapis.Changepasswd)
  }

	v := r.Group("/api/vm")
	{  vmdetails := r.Group("/api/vm/details")
    {
      ////vmdetails.Use(utils.Tokenauth())
      vmdetails.GET("cpuinfo", vmapis.Cpudetails)
      vmdetails.GET("meminfo", vmapis.Memdetails)
      vmdetails.GET("netinfo", vmapis.Netdetails)
      vmdetails.GET("diskinfo", vmapis.Diskdetails)
    }

	  v.Use(utils.Tokenauth())
    v.GET("getvm", vmapis.Getvmlist)
		v.POST("create", vmapis.Createvm)
		v.GET("operation/:id", vmapis.Operation)
		v.GET("delete", vmapis.DeleteVM)
		v.GET("getflavor", vmapis.GetFlavor)
    v.GET("vnc", vmapis.Vnc)
    v.GET("search", vmapis.Search)
    v.GET("migratelive", vmapis.MigrateVmlive)
    v.GET("getstatus", vmapis.GetVmStatus)
    v.GET("addcomment", vmapis.Addcomment)
    v.GET("getvminfo", vmapis.GetVminfo)
    v.GET("migratevm", vmapis.MigrateVm)
    v.GET("changeconfig", vmapis.Changeconfig)
    v.GET("rebuild", vmapis.Rebuild)
    v.POST("createsnap", vmapis.Createsnap)
    v.GET("getsnap", vmapis.Getsnap)
    v.DELETE("delsnap", vmapis.DelSnap)
    v.GET("rollback", vmapis.Rollback)
    v.GET("getvmarchive", vmapis.GetVmArchive)
    v.GET("delvmpermanent", vmapis.Delvmpermanent)
    v.GET("createflavor", vmapis.CreateFlavor)
    v.GET("searchvmachives", vmapis.SearchVMArchive)
	}

  m := r.Group("/api/machine")
  {
    m.Use(utils.Tokenauth())
    m.Use(utils.RoleAuth())
    m.GET("getmachinelist", machineapis.Getmachinelist)
    m.GET("addmachine", machineapis.Addmachine)
    m.GET("delmachine", machineapis.Delmachine)
    m.GET("getpage", machineapis.Getpage)
    m.GET("search", machineapis.Search)
    m.GET("update", machineapis.Update)
    m.GET("ping", machineapis.Ping)
  }

  n := r.Group("/api/networks")
  {
    n.GET("download_excel", networkapis.DownloadExcel)

    n.Use(utils.Tokenauth())
    n.GET("getvlan", networkapis.Get)
    n.GET("getvlanbydatacenter", networkapis.Getvlanbydatacenter)
    n.GET("getip", networkapis.GetIplist)
    n.GET("getallip", networkapis.GetallIp)

    n.Use(utils.RoleAuth())
    n.POST("createvlan", networkapis.Add)
    n.GET("createip", networkapis.CreateIp)
    n.GET("downip", networkapis.DownIp)
    n.GET("upip", networkapis.UpIp)
    n.GET("delete", networkapis.Delete)
    n.GET("deleteip", networkapis.Deleteip)

  }

  h := r.Group("/api/hosts")
  {
    h.Use(utils.Tokenauth())
    h.GET("gethostsbydatacenter", hostapis.GetHostsbyvlan)
    h.GET("gethosts", hostapis.GetHosts)
    h.GET("countdomains", hostapis.ListDomains)
    h.GET("gethostsby", hostapis.Gethostinfo)
    h.GET("gethostsbyip", hostapis.Gethostbyip)

    h.Use(utils.RoleAuth())
    h.GET("counthosts", hostapis.Counthost)
    h.GET("delete", hostapis.Delhost)
    h.POST("createhost", hostapis.Createhost)
    h.GET("addcomment", hostapis.Addcomment)
    h.POST("updatehost", hostapis.Updatehostinfo)
  }

  o := r.Group("/api/osimage")
  {
    o.Use(utils.Tokenauth())
    o.GET("getimagebytag", osimage.GetImagebytag)
    o.GET("getimageby", osimage.GetImageby)
    o.GET("getimage", osimage.GetImage)
    o.GET("getimagesort", osimage.GetImageSort)
    o.GET("getiostags", osimage.GetOsTag)
    o.GET("getosimagexml", osimage.Getosimagexml)

    o.Use(utils.RoleAuth())
    o.POST("updateimage", osimage.UpdateImage)
    o.POST("createimage", osimage.AddImage)
    o.GET("delimage", osimage.DelImage)
    o.POST("createxml", osimage.Addosimagexml)
    o.GET("delxml", osimage.Delosimagexml)
  }

  s := r.Group("/api/storage")
  {
    s.Use(utils.Tokenauth())
    s.GET("get", cephapis.GetStorage)
    s.GET("getpool", cephapis.Getpool)

    s.Use(utils.RoleAuth())
    s.POST("add", cephapis.Addceph)
    s.GET("delete", cephapis.Delete)
  }

  d := r.Group("/api/datacenter")
  {
    d.Use(utils.Tokenauth())
    d.GET("getdatacenter", datacenterapis.GetDatacenter)

    d.Use(utils.RoleAuth())
    d.POST("adddatacenter", datacenterapis.AddDatacenter)
    d.GET("deldatacenter", datacenterapis.DelDatacenter)
  }

  vd := r.Group("/api/vdisk")
  {
    vd.Use(utils.Tokenauth())
    vd.GET("umountdisk", vdisk.Umountdisk)
    vd.GET("mountdisk", vdisk.Mountdisk)
    vd.GET("deletevdisk", vdisk.Deletevdisk)
    vd.POST("createvdisk", vdisk.Createvdisk)
    vd.GET("getvdisk", vdisk.GetVdisk)
    vd.POST("addcomment", vdisk.AddComment)
    vd.GET("getvdiskarchives", vdisk.GetVdiskArchive)
  }

  r.Run("127.0.0.1:1992")
}

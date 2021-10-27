//+build linux
package main

import (
  "github.com/gin-gonic/gin"
  "goblog/apis"
  "goblog/apis/machineapis"
  "goblog/apis/networkapis"
  "goblog/apis/vmapis"
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
		v2.GET("getimage", vmapis.GetImage)
    v2.GET("gethosts", vmapis.GetHosts)
		v2.GET("getflavor", vmapis.GetFlavor)
    v2.GET("vnc", vmapis.Vnc)
    v2.GET("search", vmapis.Search)
    v2.GET("getstatus", vmapis.GetVmStatus)
    v2.GET("addcomment", vmapis.Addcomment)
    v2.GET("createhost", vmapis.Createhost)
    v2.GET("delhost", vmapis.Delhost)
    v2.GET("getvminfo", vmapis.GetVminfo)
    v2.GET("gethostinfo", vmapis.Gethostinfo)
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
    v4.GET("getip", networkapis.GetIplist)
    v4.GET("createip", networkapis.CreateIp)
  }
	r.Run("127.0.0.1:1992")
}

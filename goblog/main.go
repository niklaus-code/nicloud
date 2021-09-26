package main

import (
	"goblog/apis"
	"goblog/apis/vmapis"

	"github.com/gin-gonic/gin"
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
		v2.GET("getip", vmapis.GetIplist)
		v2.GET("gethost", vmapis.GetHosts)
		v2.GET("getflavor", vmapis.GetFlavor)
    v2.GET("vnc", vmapis.Vnc)
	}

	r.Run("127.0.0.1:1992")
}

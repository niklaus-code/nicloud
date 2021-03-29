package main

import "github.com/gin-gonic/gin"
import "goblog/apis"

func main() {

	r:=gin.Default()

	v1 :=r.Group("/api/blog/get_blog")
	{
	  v1.GET("/get_blog_read", apis.Get_read)
		v1.GET("/get_blog_thoughts", apis.Get_thoughts)
		v1.GET("/get_blog_by_id/:id", apis.Get_blog_by_id)
		v1.GET("/get_blog/:pagenumber", apis.Get_blog)
	}

	//v2 := r.Group("/api/vm")
	//{
	//  v2.Get("getvmlist", apis.)
  //}

	r.Run("0.0.0.0:1992")
}

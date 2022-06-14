package datacenterapis

import (
  "github.com/gin-gonic/gin"
  "nicloud/datacenter"
  "nicloud/vmerror"
)

func GetDatacenter(c *gin.Context) {
  r, err := datacenter.Get()
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, r)
}

func AddDatacenter(c *gin.Context) {
  d := c.PostForm("datacenter")
  comment := c.PostForm("comment")

  err := datacenter.Add(d, comment)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, nil)
}

func DelDatacenter(c *gin.Context) {
  d := c.Query("datacenter")
  err := datacenter.Del(d)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

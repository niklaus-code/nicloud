package datacenterapis

import (
  "github.com/gin-gonic/gin"
  "nicloud/datacenter"
)

func GetDatacenter(c *gin.Context) {

  res := make(map[string]interface{})
  r, err := datacenter.Get()

  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}

func AddDatacenter(c *gin.Context) {
  d := c.PostForm("datacenter")
  comment := c.PostForm("comment")

  res := make(map[string]interface{})
  err := datacenter.Add(d, comment)

  res["err"] = err
  c.JSON(200, res)
}

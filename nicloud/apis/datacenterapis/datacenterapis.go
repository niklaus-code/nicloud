package datacenterapis

import (
  "github.com/gin-gonic/gin"
  "nicloud/datacenter"
)

func GetDatacenter(c *gin.Context) {
  hostlist, err := datacenter.Get()
  res := make(map[string]interface{})
  res["res"] = hostlist
  res["err"] = err

  c.JSON(200, res)
}

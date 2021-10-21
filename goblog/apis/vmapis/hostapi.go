package vmapis

import (
  "github.com/gin-gonic/gin"
  "goblog/vmcommon"
  "strconv"
)

func GetHosts(c *gin.Context) {
  hostlist := vmcommon.Hosts()
  res := make(map[string]interface{})
  res["res"] = hostlist

  c.JSON(200, res)
}

func Createhost(c *gin.Context) {
  cpu, _ := strconv.Atoi(c.Query("cpu"))
  mem, _ := strconv.Atoi(c.Query("mem"))
  ip := c.Query("ip")
  num,_ := strconv.Atoi(c.Query("mem"))

  res := make(map[string]interface{})
  err := vmcommon.Createhost(cpu, mem, ip, num)
  res["res"] = err
  c.JSON(200, res)
}

func Delhost(c *gin.Context) {
  ip := c.Query("ip")
  r := vmcommon.Delhost(ip)
  res := make(map[string]interface{})
  res["res"] = r

  c.JSON(200, res)
}



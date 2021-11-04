package hostapis

import (
  "github.com/gin-gonic/gin"
  "goblog/vmcommon"
  "goblog/vmerror"
  "strconv"
)

func GetHosts(c *gin.Context) {
  hostlist := vmcommon.Hosts()
  res := make(map[string]interface{})
  res["res"] = hostlist

  c.JSON(200, res)
}

func GetHostsbydatacenter(c *gin.Context) {
  datacenter := c.Query("datacenter")
  vlan := c.Query("vlan")
  hostlist, err := vmcommon.GetHostsbydatacenter(datacenter, vlan)
  res := make(map[string]interface{})
  res["res"] = hostlist
  res["err"] = err

  c.JSON(200, res)
}

func Createhost(c *gin.Context) {
  cpu, _ := strconv.Atoi(c.Query("cpu"))
  mem, _ := strconv.Atoi(c.Query("mem"))
  ip := c.Query("ip")
  num,_ := strconv.Atoi(c.Query("mem"))
  vlan := c.Query("vlan")
  datacenter := c.Query("datacenter")

  res := make(map[string]interface{})
  err := vmcommon.Createhost(datacenter, cpu, mem, ip, num, vlan)
  res["res"] = err
  c.JSON(200, res)
}

func Delhost(c *gin.Context) {
  res := make(map[string]interface{})
  ip := c.Query("ip")
  status, err := strconv.Atoi(c.Query("status"))
  if err != nil {
    r := vmerror.Error{Message: "参数错误"}
    res["res"] = r
    c.JSON(400, res)
  }
  r := vmcommon.Restore(ip, status)
  res["res"] = r
  c.JSON(200, res)
}

func Gethostinfo(c *gin.Context) {
  ip := c.Query("ip")
  r := vmcommon.Gethostinfo(ip)
  res := make(map[string]interface{})
  res["res"] = r

  c.JSON(200, res)
}

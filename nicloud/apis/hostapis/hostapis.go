package hostapis

import (
  "github.com/gin-gonic/gin"
  "goblog/vm"
  "strconv"
)

func GetHosts(c *gin.Context) {
  hostlist := vm.Hosts()
  res := make(map[string]interface{})
  res["res"] = hostlist

  c.JSON(200, res)
}

func GetHostsbydatacenter(c *gin.Context) {
  datacenter := c.Query("datacenter")
  vlan := c.Query("vlan")
  hostlist, err := vm.GetHostsbydatacenter(datacenter, vlan)
  res := make(map[string]interface{})
  res["res"] = hostlist
  res["err"] = err

  c.JSON(200, res)
}

func Createhost(c *gin.Context) {
  cpu, _ := strconv.Atoi(c.Query("cpu"))
  mem, _ := strconv.Atoi(c.Query("mem"))
  ip := c.Query("ip")
  num,_ := strconv.Atoi(c.Query("num"))
  vlan := c.Query("vlan")
  datacenter := c.Query("datacenter")

  res := make(map[string]interface{})
  err := vm.Createhost(datacenter, cpu, mem, ip, num, vlan)
  res["res"] = err
  c.JSON(200, res)
}

func Delhost(c *gin.Context) {
  res := make(map[string]interface{})
  ip := c.Query("ip")

  r := vm.Deletehost(ip)
  res["res"] = r
  c.JSON(200, res)
}

func Gethostinfo(c *gin.Context) {
  ip := c.Query("ip")
  r := vm.Gethostinfo(ip)
  res := make(map[string]interface{})
  res["res"] = r

  c.JSON(200, res)
}

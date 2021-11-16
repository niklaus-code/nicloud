package networkapis

import (
  "github.com/gin-gonic/gin"
  "goblog/networks"
  "goblog/vm"
  "goblog/vmerror"
  "strconv"
)

func Add(c *gin.Context) {
  vlan := c.Query("vlan")
  bridge := c.Query("bridge")
  network := c.Query("network")
  prefix, err := strconv.Atoi(c.Query("prefix"))
  gateway := c.Query("gateway")
  datacenter := c.Query("datacenter")

  err1 := networks.AddVlan(datacenter, vlan, bridge, network, prefix, gateway)
  res := make(map[string]interface{})
  if err != nil {
    res["err"] = err
    res["res"] = false
    c.JSON(200, res)
  }

  if err1 != nil {
    res["res"] = false
    res["err"] = err1
  } else {
    res["res"] = true
  }

  c.JSON(200, res)
}

func Get(c *gin.Context) {
  res := make(map[string]interface{})

  vlans, err := networks.Getvlan()
  res["res"] = vlans
  res["err"] = err
  c.JSON(200, res)
}

func Getvlanbydatacenter(c *gin.Context) {
  datacenter := c.Query("datacenter")
  res := make(map[string]interface{})
  vlans, err := networks.Getvlanbydatacenter(datacenter)
  res["res"] = vlans
  res["err"] = err
  c.JSON(200, res)
}

func Delete(c *gin.Context) {
  vlan := c.Query("vlan")

  res := make(map[string]interface{})

  err := networks.Delete(vlan)
  res["err"] = err
  c.JSON(200, res)
}


func Deleteip(c *gin.Context) {
 vlan := c.Query("vlan")
 ip := c.Query("ip")
 res := make(map[string]interface{})
 ipcheck := vm.GetVmByIp(ip)

 if ipcheck.Ip != "" {
   res["err"] = vmerror.Error{Message: "ip已被占用"}
 } else {
   res["err"] = networks.Deleteip(ip, vlan)
 }
 c.JSON(200, res)
}

func CreateIp(c *gin.Context) {
  startip := c.Query("startip")
  endip := c.Query("endip")
  vlan := c.Query("vlan")
  res := make(map[string]interface{})

  err := networks.Createip(startip, endip, vlan)
  res["res"] = err
  c.JSON(200, res)
}

func UpIp(c *gin.Context) {
  ipv4 := c.Query("ipv4")
  vlan := c.Query("vlan")
  res := make(map[string]interface{})

  err := networks.OpIP(ipv4, vlan, 0)
  res["res"] = err
  c.JSON(200, res)
}


func DownIp(c *gin.Context) {
  ipv4 := c.Query("ipv4")
  vlan := c.Query("vlan")
  res := make(map[string]interface{})

  err := networks.OpIP(ipv4, vlan, 1)
  res["res"] = err
  c.JSON(200, res)
}

func GetIplist(c *gin.Context) {
  vlan := c.Query("vlan")
  iplist := networks.IPlist(vlan)
  res := make(map[string]interface{})
  res["res"] = iplist

  c.JSON(200, res)
}

func GetallIp(c *gin.Context) {
  vlan := c.Query("vlan")
  iplist := networks.AllIP(vlan)
  res := make(map[string]interface{})
  res["res"] = iplist

  c.JSON(200, res)
}

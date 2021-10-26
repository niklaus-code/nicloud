package networkapis

import (
  "github.com/gin-gonic/gin"
  "goblog/networks"
  "strconv"
)

func Add(c *gin.Context) {
  vlan := c.Query("vlan")
  bridge := c.Query("bridge")
  network := c.Query("network")
  prefix, err := strconv.Atoi(c.Query("prefix"))
  gateway := c.Query("gateway")

  err1 := networks.AddVlan(vlan, bridge, network, prefix, gateway)
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

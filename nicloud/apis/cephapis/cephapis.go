package cephapis

import (
  "github.com/gin-gonic/gin"
  "nicloud/cephcommon"
)

func  Getpool(c *gin.Context) {
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")
  cephinfo, err := cephcommon.Getpool(datacenter, storage)
  res := make(map[string]interface{})
  res["res"] = cephinfo
  res["err"] = err

  c.JSON(200, res)
}

func  GetStorage(c *gin.Context) {
  cephinfo, err := cephcommon.Get()
  res := make(map[string]interface{})
  res["res"] = cephinfo
  res["err"] = err

  c.JSON(200, res)
}

func Delete(c *gin.Context) {
  res := make(map[string]interface{})
  name := c.Query("name")
  err := cephcommon.Delete(name)
  res["err"] = err

  c.JSON(200, res)
}

func Addceph(c *gin.Context) {
  name := c.Query("name")
  pool := c.Query("pool")
  datacenter := c.Query("datacenter")
  ceph_secret := c.Query("ceph_secret")
  port := c.Query("port")
  ips := c.Query("ips")
  comment := c.Query("comment")
  err := cephcommon.Add(name, pool, datacenter, ceph_secret, ips, port, comment)
  res := make(map[string]interface{})
  res["res"] = err

  c.JSON(200, res)
}

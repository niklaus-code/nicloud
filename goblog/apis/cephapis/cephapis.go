package cephapis

import (
  "github.com/gin-gonic/gin"
  "goblog/ceph"
  "goblog/vmerror"
  "strconv"
)

func  Getcloudrive(c *gin.Context) {
  r, err := ceph.Get_cloudrive()
  res := make(map[string]interface{})
  res["res"] = r
  res["err"] = err

  c.JSON(200, res)
}


func  Addcloudrive(c *gin.Context) {
  res := make(map[string]interface{})
  contain, err := strconv.Atoi(c.Query("contain"))
  if err != nil {
    res["err"] = err
    c.JSON(400, vmerror.Error{Message: "param error"})
  }
  pool := c.Query("pool")
  storage := c.Query("storage")
  datacenter := c.Query("datacenter")
  user := "nicloud"
  r, err := ceph.Add_cloudrive(contain, pool, storage, datacenter, user)

  res["res"] = r
  res["err"] = err

  c.JSON(200, res)
}

func  Getpool(c *gin.Context) {
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")
  cephinfo, err := ceph.Getpool(datacenter, storage)
  res := make(map[string]interface{})
  res["res"] = cephinfo
  res["err"] = err

  c.JSON(200, res)
}

func  GetStorage(c *gin.Context) {
  cephinfo, err := ceph.Get()
  res := make(map[string]interface{})
  res["res"] = cephinfo
  res["err"] = err

  c.JSON(200, res)
}

func Delete(c *gin.Context) {
  res := make(map[string]interface{})
  name := c.Query("name")
  err := ceph.Delete(name)
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
  err := ceph.Add(name, pool, datacenter, ceph_secret, ips, port, comment)
  res := make(map[string]interface{})
  res["res"] = err

  c.JSON(200, res)
}

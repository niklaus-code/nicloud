package cephapis

import (
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "nicloud/cephcommon"
  "nicloud/vmerror"
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
  name := c.PostForm("name")
  pool := c.PostForm("pool")
  datacenter := c.PostForm("datacenter")
  ceph_secret := c.PostForm("ceph_secret")
  port := c.PostForm("port")
  ips := c.PostForm("ips")
  comment := c.PostForm("comment")

  ce := cephcommon.Vms_Ceph{
    Pool: pool,
    Datacenter: datacenter,
    Ceph_secret: ceph_secret,
    Ips: ips,
    Port: port,
  }

  res := make(map[string]interface{})

  validate := validator.New()
  err := validate.Struct(&ce)
  if err != nil {
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(400, res)
    return
  }

  err = cephcommon.Add(name, pool, datacenter, ceph_secret, ips, port, comment)
  res["res"] = err

  c.JSON(200, res)
}

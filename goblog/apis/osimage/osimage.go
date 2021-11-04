package osimage

import (
  "github.com/gin-gonic/gin"
  "goblog/osimage"
  "goblog/vmerror"
  "strconv"
)

func DelImage(c *gin.Context) {
  osname := c.Query("osname")
  res := make(map[string]interface{})
  r := osimage.Del(osname)

  res["res"] = r
  c.JSON(200, res)
}

func UpdateImage(c *gin.Context) {
  res := make(map[string]interface{})
  id, errparam := strconv.Atoi(c.Query("id"))
  if errparam != nil {
    res["res"] = vmerror.Error{Message: "param err by ysman"}
    c.JSON(400, res)
  }
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")
  osname := c.Query("osname")
  snapname := c.Query("snapimage")
  cephblockdevice := c.Query("cephblockdevice")
  xml := c.Query("xml")
  err := osimage.Update(id, datacenter, storage, osname, snapname, cephblockdevice, xml)

  res["res"] = err
  c.JSON(200, res)
}

func GetImage(c *gin.Context) {

  res := make(map[string]interface{})
  r, err := osimage.Get()

  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}

func AddImage(c *gin.Context) {
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")
  osname := c.Query("osname")
  snapname := c.Query("snapimage")
  cephblockdevice := c.Query("cephblockdevice")
  xml := c.Query("xml")
  res := make(map[string]interface{})
  err := osimage.Add(datacenter, storage, osname, cephblockdevice, snapname, xml)

  res["res"] = err
  c.JSON(200, res)
}

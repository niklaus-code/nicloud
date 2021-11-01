package osimage
import (
  "github.com/gin-gonic/gin"
  "goblog/osimage"
)

func DelImage(c *gin.Context) {
  osname := c.Query("osname")
  res := make(map[string]interface{})
  r := osimage.Del(osname)

  res["res"] = r
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
  osname := c.Query("osname")
  snapname := c.Query("snapimage")
  cephblockdevice := c.Query("cephblockdevice")
  xml := c.Query("xml")
  res := make(map[string]interface{})
  err := osimage.Add(osname, cephblockdevice, snapname, xml)

  res["res"] = err
  c.JSON(200, res)
}

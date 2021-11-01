package osimage
import (
  "github.com/gin-gonic/gin"
  "goblog/osimage"
)

func GetImage(c *gin.Context) {

  res := make(map[string]interface{})
  r, err := osimage.Get()

  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}

func AddImage(c *gin.Context) {
  osname := c.Query("osname")
  snapname := c.Query("snapname")
  xml := c.Query("xml")
  res := make(map[string]interface{})
  err := osimage.Add(osname, snapname, xml)

  res["res"] = err
  c.JSON(200, res)
}

package machineapis

import (
  "github.com/gin-gonic/gin"
  "goblog/machinecommon"
)

func Getmachinelist(c *gin.Context) {

res := make(map[string]interface{})
r, err := machinecommon.Machinelist()

res["res"] = r
res["err"] = err
c.JSON(200, res)
}

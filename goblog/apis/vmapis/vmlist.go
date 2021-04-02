package vmapis

import (
  "github.com/gin-gonic/gin"
  "goblog/vmcommon"
)

func Getvmlist(c *gin.Context) {

  vmlist := vmcommon.GetVmList()
  res := make(map[string]interface{})
  res["res"] = vmlist

  c.JSON(200, res)
}

func Createvm(c *gin.Context) {

  create, err := vmcommon.Create("3ee18210-3761-4fdc-9141-f13879878725")
  res := make(map[string]interface{})
  res["res"] = create
  res["err"] = err.Error()

  c.JSON(200, res)
}

package vmapis

import (
  "fmt"
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
  fmt.Println(err)
  res := make(map[string]interface{})
  res["res"] = create
  res["err"] = string(err.Error())

  c.JSON(200, res)
}

func GetStatus(c *gin.Context) {
  s, _ := vmcommon.GetVmStatus("31a803b2-5f11-4f14-875f-d14347db13fb")
  res := make(map[string]interface{})
  res["res"] = s
  c.JSON(200, res)
}

func  VmStart(c *gin.Context)  {
  s, err := vmcommon.Start("31a803b2-5f11-4f14-875f-d14347db13fb")
  fmt.Println(err)
  res := make(map[string]interface{})
  res["res"] = s
  res["err"] = err
  c.JSON(200, res)
}

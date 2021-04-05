package vmapis

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "goblog/vmcommon"
  "strconv"
)

func Getvmlist(c *gin.Context) {

  vmlist := vmcommon.VmList()
  res := make(map[string]interface{})
  res["res"] = vmlist

  c.JSON(200, res)
}

func Createvm(c *gin.Context) {

  create, err := vmcommon.Create("3ee18210-3761-4fdc-9141-f13879878725")
  res := make(map[string]interface{})
  res["res"] = create
  res["err"] = string(err.Error())

  c.JSON(200, res)
}

func GetStatus(c *gin.Context) {
  s, _ := vmcommon.VmStatus("31a803b2-5f11-4f14-875f-d14347db13fb")
  res := make(map[string]interface{})
  res["res"] = s
  c.JSON(200, res)
}

func  Operation(c *gin.Context)  {
  uuid := c.Query("uuid")
  fmt.Println(uuid)
  res := make(map[string]interface{})

  o, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(400, res)
  }

  var s *vmcommon.Vms
  switch o {
  case 1: s, _ = vmcommon.Start(uuid)
  case 0: s, _ = vmcommon.Shutdown(uuid)
  }


  res["res"] = s
  res["err"] = err
  c.JSON(200, res)
}

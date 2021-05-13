package vmapis

import (
  "github.com/gin-gonic/gin"
  "goblog/vmcommon"
  "strconv"
)


func GetHosts(c *gin.Context)  {
  hostlist := vmcommon.Hosts()
  res := make(map[string]interface{})
  res["res"] = hostlist

  c.JSON(200, res)
}

func GetIplist(c *gin.Context)  {
  iplist := vmcommon.IPlist()
  res := make(map[string]interface{})
  res["res"] = iplist

  c.JSON(200, res)
}

func Getvmlist(c *gin.Context) {

  vmlist := vmcommon.VmList()
  res := make(map[string]interface{})
  res["res"] = vmlist

  c.JSON(200, res)
}

func Createvm(c *gin.Context) {
  cpu, _  := strconv.Atoi(c.Query("cpu"))
  mem, _ := strconv.Atoi(c.Query("mem"))
  ip := c.Query("ip")

  create, err := vmcommon.Create(cpu, mem, ip)
  if err != nil {
    c.JSON(500, err)
  }
  res := make(map[string]interface{})
  res["res"] = create

  c.JSON(200, res)
}

func GetStatus(c *gin.Context) {
  s, _ := vmcommon.VmStatus("31a803b2-5f11-4f14-875f-d14347db13fb")
  res := make(map[string]interface{})
  res["res"] = s
  c.JSON(200, res)
}

func DeleteVM(c *gin.Context)  {
  uuid := c.Query("uuid")
  ip := c.Query("ip")

  res := make(map[string]interface{})
  r, _ := vmcommon.Delete(uuid, ip)

  res["res"] = r
  res["err"] = nil
  c.JSON(200, res)
}

func  Operation(c *gin.Context)  {
  uuid := c.Query("uuid")
  res := make(map[string]interface{})

  var err error

  o, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(400, res)
  }

  var s *vmcommon.Vms
  switch o {
  case 0: s, err = vmcommon.Shutdown(uuid)
  case 1: s, err = vmcommon.Start(uuid)
  }

  res["res"] = s
  res["err"] = err
  c.JSON(200, res)
}


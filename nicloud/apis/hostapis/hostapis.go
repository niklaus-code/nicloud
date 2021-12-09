package hostapis

import (
  "github.com/gin-gonic/gin"
  validator "github.com/go-playground/validator/v10"
  "nicloud/vm"
  "nicloud/vmerror"
  "strconv"
)

func GetHosts(c *gin.Context) {
  hostlist := vm.Hosts()
  res := make(map[string]interface{})
  res["res"] = hostlist

  c.JSON(200, res)
}

func GetHostsbydatacenter(c *gin.Context) {
  datacenter := c.Query("datacenter")
  vlan := c.Query("vlan")
  hostlist, err := vm.GetHostsbydatacenter(datacenter, vlan)
  res := make(map[string]interface{})
  res["res"] = hostlist
  res["err"] = err

  c.JSON(200, res)
}


func Createhost(c *gin.Context) {
  res := make(map[string]interface{})

  cpu, _ := strconv.Atoi(c.PostForm("cpu"))
  mem, _ := strconv.Atoi(c.PostForm("mem"))
  ipv4 := c.PostForm("ipv4")
  max_vms,_ := strconv.Atoi(c.PostForm("max_vms"))
  vlan := c.PostForm("vlan")
  datacenter := c.PostForm("datacenter")
  h := vm.Vm_hosts{
    Cpu: cpu,
    Mem: mem,
    Ipv4: ipv4,
    Max_vms: max_vms,
    Vlan: vlan,
    Datacenter: datacenter,
  }
  validate := validator.New()
  err := validate.Struct(h)
  if err != nil {
    res["res"] = vmerror.Error{Message: err.Error()}
    c.JSON(400, res)
    return
  }

  err = vm.Createhost(datacenter, cpu, mem, ipv4, max_vms, vlan)
  res["res"] = err
  c.JSON(200, res)
}

func Delhost(c *gin.Context) {
  res := make(map[string]interface{})
  ip := c.Query("ip")

  r := vm.Deletehost(ip)
  res["res"] = r
  c.JSON(200, res)
}

func Gethostinfo(c *gin.Context) {
  ip := c.Query("ip")
  r := vm.Gethostinfo(ip)
  res := make(map[string]interface{})
  res["res"] = r

  c.JSON(200, res)
}

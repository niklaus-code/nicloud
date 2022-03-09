package hostapis

import (
  "github.com/gin-gonic/gin"
  validator "github.com/go-playground/validator/v10"
  "nicloud/vm"
  "nicloud/vmerror"
  "strconv"
)

func ListDomains(c *gin.Context) {
  host := c.Query("host")
  countvm, err := vm.ListDomains(host)
  res := make(map[string]interface{})
  res["res"] = countvm
  res["err"] = err

  c.JSON(200, res)
}

func GetHosts(c *gin.Context) {
  hostlist, err := vm.Hosts()
  res := make(map[string]interface{})
  res["res"] = hostlist
  res["err"] = err

  c.JSON(200, res)
}

func Addcomment(c *gin.Context) {
  ip := c.Query("ip")
  comment := c.Query("comment")

  err := vm.Addcomment(ip, comment)
  res := make(map[string]interface{})
  res["err"] = err

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
    res["err"] = vmerror.Error{Message: err.Error()}
    c.JSON(400, res)
    return
  }

  err = vm.Createhost(datacenter, cpu, mem, ipv4, max_vms, vlan)
  res["err"] = err
  c.JSON(200, res)
}

func Delhost(c *gin.Context) {
  res := make(map[string]interface{})
  ip := c.Query("ip")

  err := vm.Deletehost(ip)
  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
  }

  c.JSON(200, res)
}

func Gethostinfo(c *gin.Context) {
  ip := c.Query("ip")
  r := vm.Gethostinfo(ip)
  res := make(map[string]interface{})
  res["res"] = r

  c.JSON(200, res)
}

func Counthost(c *gin.Context) {
  res := make(map[string]interface{})
  r, err := vm.CountHost()
  res["res"] = r
  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
  }

  c.JSON(200, res)
}

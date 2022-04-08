package hostapis

import (
  "encoding/json"
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

func GetHostsbyvlan(c *gin.Context) {
  datacenter := c.Query("datacenter")
  vlan := c.Query("vlan")

  h := vm.Vm_hosts{}
  hostlist, err := h.GetHostsbyVlan(datacenter, vlan)
  res := make(map[string]interface{})

  res["err"] = nil
  res["res"] = hostlist
  if err != nil {
    res["err"] = err.Error()
  }

  c.JSON(200, res)
}


func Createhost(c *gin.Context) {
  res := make(map[string]interface{})

  cpu, _ := strconv.Atoi(c.PostForm("cpu"))
  mem, _ := strconv.Atoi(c.PostForm("mem"))
  ipv4 := c.PostForm("ipv4")
  max_vms,_ := strconv.Atoi(c.PostForm("max_vms"))
  vlan, _ := c.GetPostForm("vlan")
  datacenter := c.PostForm("datacenter")
  h := vm.Vm_hosts{
    Cpu: uint(cpu),
    Mem: uint(mem),
    Ipv4: ipv4,
    Max_vms: uint(max_vms),
    Datacenter: datacenter,
  }
  validate := validator.New()
  err := validate.Struct(h)
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
    c.JSON(400, res)
    return
  }

  var ss []string
  err = json.Unmarshal([]byte(vlan), &ss)
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
    c.JSON(400, res)
    return
  }

  err = h.Createhost(datacenter, uint(cpu), uint(mem), ipv4, uint(max_vms), ss)
  res["err"] = err
  c.JSON(200, res)
}

func Delhost(c *gin.Context) {
  res := make(map[string]interface{})
  ip := c.Query("ip")
  h := vm.Vm_hosts{}

  err := h.Deletehost(ip)
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

func Gethostbyip(c *gin.Context) {
  ip := c.Query("ip")
  r, err := vm.Maphost(ip)
  res := make(map[string]interface{})
  res["err"] = nil
  if err != nil {
    res["err"] = err.Error()
  }
  res["res"] = r
  c.JSON(200, res)
}

func Updatehostinfo(c *gin.Context) {
  h := vm.Vm_hosts{}
  ip := c.PostForm("ip")
  cpu, _ := strconv.Atoi(c.PostForm("cpu"))
  mem, _ := strconv.Atoi(c.PostForm("mem"))
  num, _ := strconv.Atoi(c.PostForm("num"))
  vlanlist := c.PostForm("vlan")
  res := make(map[string]interface{})

  var ss []string
  err := json.Unmarshal([]byte(vlanlist), &ss)
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
    c.JSON(400, res)
    return
  }

  err = h.Updatehostinfo(ip, cpu, mem, num, ss)

  res["err"] = nil
  if err != nil {
    res["err"] = err.Error()
  }

  c.JSON(200, res)
}

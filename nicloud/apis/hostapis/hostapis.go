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
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, countvm)
}

func GetHosts(c *gin.Context) {
  hostlist, err := vm.Hosts()
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, hostlist)
}

func Addcomment(c *gin.Context) {
  ip := c.Query("ip")
  comment := c.Query("comment")

  err := vm.Addcomment(ip, comment)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func GetHostsbyvlan(c *gin.Context) {
  datacenter := c.Query("datacenter")
  vlan := c.Query("vlan")

  h := vm.Vm_hosts{}
  hostlist, err := h.GetHostsbyVlan(datacenter, vlan)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, hostlist)
}


func Createhost(c *gin.Context) {
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
    vmerror.REQUESTERROR(c, err)
    return
  }

  var ss []string
  err = json.Unmarshal([]byte(vlan), &ss)
  if err != nil {
    vmerror.SERVERERROR(c ,err)
    return
  }

  err = h.Createhost(datacenter, uint(cpu), uint(mem), ipv4, uint(max_vms), ss)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func Delhost(c *gin.Context) {
  ip := c.Query("ip")
  h := vm.Vm_hosts{}

  err := h.Deletehost(ip)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func Gethostinfo(c *gin.Context) {
  ip := c.Query("ip")
  r, err := vm.Gethostinfo(ip)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, r)
}

func Counthost(c *gin.Context) {
  r, err := vm.CountHost()
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, r)
}

func Gethostbyip(c *gin.Context) {
  ip := c.Query("ip")
  r, err := vm.Maphost(ip)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, r)
}

func Updatehostinfo(c *gin.Context) {
  h := vm.Vm_hosts{}
  ip := c.PostForm("ip")
  cpu, _ := strconv.Atoi(c.PostForm("cpu"))
  mem, _ := strconv.Atoi(c.PostForm("mem"))
  num, _ := strconv.Atoi(c.PostForm("num"))
  vlanlist := c.PostForm("vlan")

  var ss []string
  err := json.Unmarshal([]byte(vlanlist), &ss)
  if err != nil {
    vmerror.REQUESTERROR(c, err)
    return
  }

  err = h.Updatehostinfo(ip, cpu, mem, num, ss)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

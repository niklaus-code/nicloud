package vdisk

import (
  "github.com/gin-gonic/gin"
  vdisk "goblog/vdsik"
  "goblog/vm"
  "goblog/vmerror"
  "strconv"
)

func Mountdisk(c *gin.Context) {
  vmid := c.Query("vmid")
  ip := c.Query("ip")
  storage := c.Query("storage")
  datacenter := c.Query("datacenter")
  pool := c.Query("pool")
  host := c.Query("host")
  cloudriveid := c.Query("cloudriveid")
  vms := vm.Vms{}
  res := make(map[string]interface{})

  s, err := vm.VmStatus(vmid, host)
  if err != nil {
    res["err"] = err
  }

  xml, err := vm.Getvmxmlby(ip, storage, datacenter)
  if err != nil {
    res["err"] = err
  }

  if s != "关机" {
    res["err"] = vmerror.Error{Message: "cont mount disk, vm is " + s}
  }
  r := vdisk.Mountdisk(ip,  host, storage, pool, datacenter, cloudriveid, vms, xml)

  res["err"] = r
  c.JSON(200, res)
}

func Deletevdisk(c *gin.Context)  {
  uuid := c.Query("uuid")
  err := vdisk.Deletevdisk(uuid)
  res := make(map[string]interface{})
  res["err"] = err

  c.JSON(200, res)
}


func Createvdisk(c *gin.Context) {
  res := make(map[string]interface{})
  contain, err := strconv.Atoi(c.Query("contain"))
  if err != nil {
    res["err"] = err
    c.JSON(400, vmerror.Error{Message: "param error"})
  }
  pool := c.Query("pool")
  storage := c.Query("storage")
  datacenter := c.Query("datacenter")
  user := "nicloud"
  r, err := vdisk.Add_cloudrive(contain, pool, storage, datacenter, user)

  res["res"] = r
  res["err"] = err

  c.JSON(200, res)
}

func Umountdisk(c *gin.Context) {
  vmip := c.Query("vmip")
  storage := c.Query("storage")
  datacenter := c.Query("datacenter")
  vdiskid := c.Query("vdiskid")
  res := make(map[string]interface{})
  vminfo := vm.GetVmByIp(vmip)

  s, err := vm.VmStatus(vminfo.Uuid, vminfo.Host)
  if err != nil {
    res["err"] = err
  }

  if s != "关机" {
    res["err"] = vmerror.Error{Message: "cont mount disk, vm is " + s}
  }

  xml, err := vm.Getvmxmlby(vmip, storage, datacenter)
  if err != nil {
    res["err"] = err
  } else {
    v := vm.Vms{}
    r := vdisk.Umountdisk(vmip, storage, datacenter, vdiskid, xml, vminfo.Host, v)
    res["err"] = r
  }
  c.JSON(200, res)
}

func  GetVdisk(c *gin.Context) {
  r, err := vdisk.Getvdisk()
  res := make(map[string]interface{})
  res["res"] = r
  res["err"] = err

  c.JSON(200, res)
}
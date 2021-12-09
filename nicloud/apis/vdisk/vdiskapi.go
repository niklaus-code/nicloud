package vdisk

import (
  "github.com/gin-gonic/gin"
  vdisk "nicloud/vdisk"
  "nicloud/vm"
  "nicloud/vmerror"
  "strconv"
)

func Mountdisk(c *gin.Context) {
  vmid := c.Query("vmid")
  ip := c.Query("ip")
  storage := c.Query("storage")
  datacenter := c.Query("datacenter")
  pool := c.Query("pool")
  host := c.Query("host")
  vdiskid := c.Query("vdiskid")

  vms := vm.Vms{}
  res := make(map[string]interface{})

  s, err := vm.VmStatus(vmid, host)
  if err != nil {
    res["err"] = err
    c.Abort()
    c.JSON(200, res)
  }

  if s != "关机" {
    res["err"] = vmerror.Error{Message: "cont mount disk, vm is " + s}
    c.Abort()
    c.JSON(200, res)
  }

  xml, err := vm.Getvmxmlby(ip, storage, datacenter)
  if err != nil {
    res["err"] = err
    c.Abort()
    c.JSON(200, res)
  }

  err = vdisk.Mountdisk(ip,  host, storage, pool, datacenter, vdiskid, vms, xml)
  if err != nil {
    res["err"] = err
    c.Abort()
    c.JSON(200, res)
  } else {
    res["err"] = nil
    c.JSON(200, res)
  }
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
    err = vmerror.Error{Message: "param error"}
    res["err"] = err
    c.Abort()
    c.JSON(200, res)
  }
  pool := c.Query("pool")
  storage := c.Query("storage")
  datacenter := c.Query("datacenter")
  user := "nicloud"
  err = vdisk.Add_vdisk(contain, pool, storage, datacenter, user)

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
    c.Abort()
    c.JSON(200, res)
  }

  if s != "关机" {
    res["err"] = vmerror.Error{Message: "cont mount disk, vm is " + s}
    c.Abort()
    c.JSON(200, res)
  }

  xml, err := vm.Getvmxmlby(vmip, storage, datacenter)
  if err != nil {
    res["err"] = err
    c.Abort()
    c.JSON(200, res)
  }

  v := vm.Vms{}
  err = vdisk.Umountdisk(vmip, storage, datacenter, vdiskid, xml, vminfo.Host, v)
  if err != nil {
    res["err"] = err
    c.Abort()
    c.JSON(200, res)
  } else {
    res["err"] = nil
    c.JSON(200, res)
  }
}

func  GetVdisk(c *gin.Context) {
  vmip := c.Query("vmip")
  r, err := vdisk.Getvdisk(vmip)
  res := make(map[string]interface{})
  res["res"] = r
  res["err"] = err

  c.JSON(200, res)
}

package vdisk

import (
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "nicloud/utils"
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
    c.JSON(200, res)
    return
  }

  if s != "关机" {
    res["err"] = vmerror.Error{Message: "cont mount disk, vm is " + s}
    c.JSON(200, res)
    return
  }

  xml, err := vm.Getvmxmlby(ip, storage, datacenter)
  if err != nil {
    res["err"] = vmerror.Error{Message: "获取云主机信息失败"}
    c.JSON(200, res)
    return
  }

  err = vdisk.Mountdisk(ip,  host, storage, pool, datacenter, vdiskid, vms, xml)
  if err != nil {
    res["err"] = err
    c.Abort()
    c.JSON(200, res)
    return
  }

  res["err"] = nil
  c.JSON(200, res)

}

func Deletevdisk(c *gin.Context)  {
  uuid := c.Query("uuid")
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")
  err := vdisk.Deletevdisk(uuid, datacenter, storage)

  res := make(map[string]interface{})
  res["err"] = err

  c.JSON(200, res)
}


func Createvdisk(c *gin.Context) {
  res := make(map[string]interface{})
  contain, _ := strconv.Atoi(c.PostForm("contain"))
  pool := c.PostForm("pool")
  storage := c.PostForm("storage")
  datacenter := c.PostForm("datacenter")

  token := c.Request.Header.Get("token")
  user, err := utils.ParseToken(token)
  if err != nil {
    res["err"] = vmerror.Error{Message: "认证失败"}
    c.JSON(200, res)
    return
  }

  d := vdisk.Vms_vdisks{
    Contain: contain,
    Pool: pool,
    Storage: storage,
    Datacenter: datacenter,
    User: user,
  }
  validate := validator.New()
  err = validate.Struct(d)
  if err != nil {
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(400, res)
    return
  }

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
    c.JSON(200, res)
    return
  }

  if s != "关机" {
    res["err"] = vmerror.Error{Message: "cont mount disk, vm is " + s}
    c.Abort()
    c.JSON(200, res)
    return
  }

  xml, err := vm.Getvmxmlby(vmip, storage, datacenter)
  if err != nil {
    res["err"] = err
    c.JSON(200, res)
    return
  }

  v := vm.Vms{}
  err = vdisk.Umountdisk(vmip, storage, datacenter, vdiskid, xml, vminfo.Host, v)
  if err != nil {
    res["err"] = err
    c.JSON(200, res)
    return
  }
  res["err"] = nil
  c.JSON(200, res)
}

func  GetVdisk(c *gin.Context) {
  res := make(map[string]interface{})

  token := c.Request.Header.Get("token")
  user, err := utils.ParseToken(token)
  if err != nil {
    res["err"] = vmerror.Error{Message: "认证失败"}
    c.JSON(200, res)
    return
  }

  r, err := vdisk.Getvdisk(user)

  res["res"] = r
  res["err"] = err

  c.JSON(200, res)
}

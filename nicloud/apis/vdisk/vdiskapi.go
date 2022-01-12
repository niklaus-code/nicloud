package vdisk

import (
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "nicloud/cephcommon"
  "nicloud/utils"
  vdisk "nicloud/vdisk"
  "nicloud/vm"
  "nicloud/vmerror"
  "strconv"
  "sync"
)

func Mountdisk(c *gin.Context) {
  vmid := c.Query("vmid")
  vdiskid := c.Query("vdiskid")

  vms := vm.Vms{}
  res := make(map[string]interface{})

  vminfo, err := vm.GetVmByUuid(vmid)
  if err != nil {
    res["err"] = vmerror.Error{Message: "获取云主机信息失败"}
    c.JSON(200, res)
    return
  }

  s, err := vm.VmStatus(vmid, vminfo.Host)
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

  storageinfo, err := cephcommon.Cephinfobyname(vminfo.Datacenter, vminfo.Storage)
  if err != nil {
    res["err"] = vmerror.Error{Message: "获取云主机信息失败"}
    c.JSON(200, res)
    return
  }

  var rwLock sync.RWMutex
  rwLock.Lock()
  err = vdisk.Mountdisk(vminfo.Ip, vminfo.Host, vminfo.Storage, storageinfo.Pool, vminfo.Datacenter, vdiskid, vms, vminfo.Vmxml)
  rwLock.Unlock()
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

  var rwLock sync.RWMutex
  rwLock.Lock()
  err := vdisk.Deletevdisk(uuid)
  rwLock.Unlock()
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
  userid, err := utils.ParseToken(token)
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
    User: userid,
  }
  validate := validator.New()
  err = validate.Struct(d)
  if err != nil {
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(400, res)
    return
  }

  var rwLock sync.RWMutex
  rwLock.Lock()
  err = vdisk.Add_vdisk(contain, pool, storage, datacenter, userid)
  rwLock.Unlock()
  res["err"] = err

  c.JSON(200, res)
}

func Umountdisk(c *gin.Context) {
  vmip := c.Query("vmip")
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

  xml, err := vm.Getvmxmlby(vmip, vminfo.Storage, vminfo.Datacenter)
  if err != nil {
    res["err"] = err
    c.JSON(200, res)
    return
  }

  v := vm.Vms{}
  var rwLock sync.RWMutex
  rwLock.Lock()
  err = vdisk.Umountdisk(vmip, vminfo.Storage, vminfo.Datacenter, vdiskid, xml, vminfo.Host, v)
  rwLock.Unlock()
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
  userid, err := utils.ParseToken(token)
  if err != nil {
    res["err"] = vmerror.Error{Message: "认证失败"}
    c.JSON(200, res)
    return
  }

  r, err := vdisk.Getvdisk(userid)

  res["res"] = r
  res["err"] = err

  c.JSON(200, res)
}

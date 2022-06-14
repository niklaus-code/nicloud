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
var ceph cephcommon.Vms_Ceph

func Mountdisk(c *gin.Context) {
  vmid := c.Query("vmid")
  vdiskid := c.Query("vdiskid")

  vms := vm.Vms{}

  vminfo, err := vm.GetVmByUuid(vmid)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  s, err := vm.VmStatus(vmid, vminfo.Host)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  if s != "关机" {
    vmerror.SUCCESS(c, vmerror.Error{Message: "cont mount disk, vm is " + s})
    return
  }

  storageinfo, err := ceph.Cephinfobyuuid(vminfo.Storage)
  if err != nil {
    vmerror.SERVERERROR(c, vmerror.Error{Message: "获取云主机信息失败"})
    return
  }

  var rwLock sync.RWMutex
  rwLock.Lock()
  err = vdisk.Mountdisk(vminfo.Ip, vminfo.Host, vminfo.Storage, storageinfo.Pool, vminfo.Datacenter, vdiskid, vms, vminfo.Vmxml, vminfo.Os)
  rwLock.Unlock()
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, nil)
}

func Deletevdisk(c *gin.Context)  {
  uuid := c.Query("uuid")
  comment := c.Query("comment")

  var rwLock sync.RWMutex
  rwLock.Lock()
  err := vdisk.Deletevdisk(uuid, comment)
  rwLock.Unlock()
  res := make(map[string]interface{})
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, res)
}

func Createvdisk(c *gin.Context) {
  contain, _ := strconv.Atoi(c.PostForm("contain"))
  pool := c.PostForm("pool")
  cephid := c.PostForm("storage")
  datacenter := c.PostForm("datacenter")
  comment := c.PostForm("comment")

  token := c.Request.Header.Get("token")
  userid, err := utils.ParseToken(token)
  if err != nil {
    vmerror.SUCCESS(c, vmerror.Error{Message: "认证失败"})
    return
  }

  d := vdisk.Vms_vdisks{
    Contain: contain,
    Pool: pool,
    Storage: cephid,
    Datacenter: datacenter,
    User: userid,
  }
  validate := validator.New()
  err = validate.Struct(d)
  if err != nil {
    vmerror.REQUESTERROR(c, err)
    return
  }

  var rwLock sync.RWMutex
  rwLock.Lock()
  err = d.Create(contain, pool, cephid, datacenter, userid, comment)
  rwLock.Unlock()
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, nil)
}

func Umountdisk(c *gin.Context) {
  vmip := c.Query("vmip")
  vdiskid := c.Query("vdiskid")
  vminfo := vm.GetVmByIp(vmip)

  s, err := vm.VmStatus(vminfo.Uuid, vminfo.Host)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  if s != "关机" {
    vmerror.SUCCESS(c, vmerror.Error{Message: "卸载云盘，需要云主机处于关机状态"})
    return
  }

  xml, err := vm.Getvmxmlby(vmip, vminfo.Storage, vminfo.Datacenter)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  v := vm.Vms{}
  d := vdisk.Vms_vdisks{}
  var rwLock sync.RWMutex
  rwLock.Lock()
  err = d.Umountdisk(vmip, vminfo.Storage, vminfo.Datacenter, vdiskid, xml, vminfo.Host, v)
  rwLock.Unlock()

  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, nil)
}

func GetVdisk(c *gin.Context) {
  token := c.Request.Header.Get("token")
  userid, err := utils.ParseToken(token)
  if err != nil {
    vmerror.SUCCESS(c, vmerror.Error{Message: "认证失败"})
    return
  }

  r, err := vdisk.Getvdisk(userid)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, r)
}

func AddComment(c *gin.Context)  {
  res := make(map[string]interface{})
  uuid := c.PostForm("uuid")
  comment := c.PostForm("comment")
  v := vdisk.Vms_vdisks{}
  err := v.Addcomment(uuid, comment)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, res)
}

func GetVdiskArchive(c *gin.Context)  {
  vd := vdisk.Vms_vdisks_archives{}

  r, err := vd.GetVmArchive()
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, r)
}

package vmapis

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "net/http"
  "nicloud/utils"
  "nicloud/vm"
  "nicloud/vdisk"
  "nicloud/vmerror"
  "strconv"
  "sync"
)

func Vnc(c *gin.Context)  {
  uuid := c.Query("uuid")
  host := c.Query("host")
  res := fmt.Sprintf("http://%s/novnc/vnc.html?path=websockify/?token=%s", host, uuid)
  c.JSON(200, res)
}

  //xlsx := excelize.NewFile()
  //xlsx.SetCellValue("Sheet1", "A2", "asdas")
  //
  //c.Header("Content-Type", "application/octet-stream")
  //c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
  //c.Header("Content-Transfer-Encoding", "binary")
  //_ = xlsx.Write(c.Writer)

func Search(c *gin.Context)  {
  ct := c.Query("content")
  vms, err := vm.SearchVm(ct)

  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, vms)
}

func GetVminfo(c *gin.Context) {
  uuid := c.Query("uuid")
  iplist, err := vm.GetVmByUuid(uuid)

  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, iplist)
}

func GetVmStatus(c *gin.Context) {
  //之前把res（map） 放到了公共变量， 导致获取vmlist , 还有获取vm状态 都要使用这个res（map）， 并发访问就会 gouteline泄漏

  host := c.Query("host")
  uuid := c.Query("uuid")

  vmstate, err := vm.VmStatus(uuid, host)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, vmstate)
}

func Getvmlist(c *gin.Context) {
  start, err := strconv.Atoi(c.Query("start"))
  item := c.Query("item")
  if err != nil {
    vmerror.REQUESTERROR(c, err)
    return
  }
  token := c.Request.Header.Get("token")
  userid, err := utils.ParseToken(token)
  if err != nil {
   vmerror.SERVERERROR(c, vmerror.Error{Message: "认证失败"})
   return
  }

	pagenumber, vmcount,  err := vm.Getpagenumber(userid)
	if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmlist, err := vm.VmList(userid, start, item)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  var res = make(map[string]interface{})
	res["res"] = vmlist
  res["pagenumber"] = pagenumber
  res["vmcount"] = vmcount
  res["err"] = nil

	c.JSON(http.StatusOK, res)
}

func MigrateVmlive(c *gin.Context) {
  uuid := c.Query("uuid")
  migratehost := c.Query("migratehost")

  err := vm.MigrateVmlive(uuid, migratehost)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func MigrateVm(c *gin.Context) {
  uuid := c.Query("uuid")
  migratehost := c.Query("migratehost")

  err := vm.MigrateVm(uuid, migratehost)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, nil)
}

func Createvm(c *gin.Context) {
  ip := c.PostForm("ip")
  cpu, _ := strconv.Atoi(c.PostForm("cpu"))
  mem, _ := strconv.Atoi(c.PostForm("mem"))
  host := c.PostForm("host")
  osid, _ := strconv.Atoi(c.PostForm("os"))
  datacenter := c.PostForm("datacenter")
  storage := c.PostForm("storage")
  vlan := c.PostForm("vlan")
  comment := c.PostForm("comment")

  token := c.Request.Header.Get("token")
  user, err := utils.ParseToken(token)

  if err != nil {
    vmerror.SERVERERROR(c, vmerror.Error{Message: "认证失败"})
    return
  }

  v := vm.Vms{
    Ip: ip,
    Cpu: uint(cpu),
    Mem: uint(mem),
    Host: host,
    Os: osid,
    Datacenter: datacenter,
    Storage: storage,
    Owner: user,
  }

  validate := validator.New()
  err = validate.Struct(v)
  if err != nil {
    vmerror.REQUESTERROR(c, err)
    return
  }
  var rwLock sync.RWMutex
  rwLock.Lock()
  err = v.Create(datacenter, storage, vlan, uint(cpu), uint(mem), ip, host, osid, user, comment)
  rwLock.Unlock()

  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func Addcomment(c *gin.Context) {
  uuid := c.Query("uuid")
  comment := c.Query("comment")
  r, err := vm.Updatecomments(uuid, comment)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, r)
}

func GetFlavor(c *gin.Context) {
	s, err := vm.Flavor()
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
	vmerror.SUCCESS(c, s)
}

func Changeconfig(c *gin.Context) {
  id := c.Query("uuid")
  host := c.Query("host")
  vmhost := c.Query("vmhost")
  cpu, err := strconv.Atoi(c.Query("cpu"))
  if err != nil {
    vmerror.REQUESTERROR(c, err)
    return
  }

  oldcpu, err := strconv.Atoi(c.Query("oldcpu"))
  if err != nil {
    vmerror.REQUESTERROR(c, err)
    return
  }

  oldmem, err := strconv.Atoi(c.Query("oldmem"))
  if err != nil {
    vmerror.REQUESTERROR(c, err)
    return
  }

  mem, err := strconv.Atoi(c.Query("mem"))
  if err != nil {
    vmerror.REQUESTERROR(c, err)
    return
  }
  err = vm.Changeconfig(id, host, uint(cpu), uint(oldcpu), uint(mem), uint(oldmem), vmhost)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func DeleteVM(c *gin.Context) {
	uuid := c.Query("uuid")
  storage := c.Query("storage")

  var rwLock sync.RWMutex
	rwLock.Lock()
	err := vm.Delete(uuid, storage)
	rwLock.Unlock()

	if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
	vmerror.SUCCESS(c, nil)
}

func Operation(c *gin.Context) {
	uuid := c.Query("uuid")
	host := c.Query("host")

	var err error

	o, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		vmerror.REQUESTERROR(c, err)
    return
	}

	var s *vm.Vms
	switch o {
	case 0:
		err = vm.Shutdown(uuid, host)
  case 1:
    err = vm.Destroy(uuid, host)
	case 2:
		err = vm.Start(uuid, host)
  case 3:
    err = vm.PauseVm(uuid, host)
  case 4:
    err = vm.Reboot(uuid, host)
	}

	if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
	vmerror.SUCCESS(c, s)
}

func Rebuild(c *gin.Context)  {
  uuid := c.Query("uuid")
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")
  osid, _ := strconv.Atoi(c.Query("osname"))
  host := c.Query("host")

  v := vm.Vms{}
  err := v.Rebuildimg(osid, storage, datacenter, uuid, host)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func Createsnap(c *gin.Context)  {
  token := c.Request.Header.Get("token")
  userid, err := utils.ParseToken(token)
  if err != nil {
    vmerror.SERVERERROR(c, vmerror.Error{Message: "认证失败"})
    return
  }

  snapname := c.PostForm("snapname")

  if len(snapname) == 0 {
    vmerror.REQUESTERROR(c, nil)
    return
  }

  uuid := c.PostForm("uuid")
  datacenter := c.PostForm("datacenter")
  storage := c.PostForm("storage")
  protect, err := strconv.ParseBool(c.PostForm("protect"))
  if err != nil {
    vmerror.REQUESTERROR(c, err)
    return
  }

  if protect == false {
    err = vm.CreatSnap(uuid, datacenter, storage, snapname)
  } else {
    err = vm.SaveSnapToImg(uuid, datacenter, storage, snapname, userid)
  }

  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, nil)
}

func Getsnap(c *gin.Context)  {
  uuid := c.Query("uuid")
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")

  s, err := vm.Getsnap(datacenter, storage, uuid)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, s)
}

func Rollback(c *gin.Context)  {
  uuid := c.Query("uuid")
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")
  snapname := c.Query("snapname")

  err := vm.RollbackSnap(uuid, snapname,  datacenter, storage)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func DelSnap(c *gin.Context)  {

  uuid := c.Query("uuid")
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")
  snapname := c.Query("snapname")

  err := vm.DelSnap(uuid, snapname,  datacenter, storage)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, nil)
}

func GetVmArchive(c *gin.Context)  {
  startpage, err := strconv.Atoi(c.Query("startpage"))
  if err != nil {
    vmerror.REQUESTERROR(c, err)
    return
  }
  ar := vm.Vms_archives{}

  token := c.Request.Header.Get("token")
  userid, err := utils.ParseToken(token)
  if err != nil {
    vmerror.SERVERERROR(c, vmerror.Error{Message: "认证失败"})
    return
  }
  pagenumber, vmcount,  err := vm.Getvmarchivepagenumber(userid)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  r, err := ar.GetVmArchive(startpage)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  var res = make(map[string]interface{})
  res["res"] = r
  res["pagenumber"] = pagenumber
  res["vmcount"] = vmcount
  res["err"] = nil


  c.JSON(http.StatusOK, res)
}

func Delvmpermanent(c *gin.Context) {
  var res = make(map[string]interface{})
  //暂时不开放次接口
  res["err"] = vmerror.Error{Message: "暂时不开放此接口"}
  c.JSON(http.StatusOK, res)
  return

  uuid := c.Query("uuid")
  storage := c.Query("storage")


  v := vm.Vms_archives{}

  del := v.Delvmpermanent(storage, uuid)
  res["err"] = nil
  if del != nil {
    res["err"] = vmerror.Error{Message: del.Error()}
    c.JSON(http.StatusInternalServerError, res)
    return
  }
  c.JSON(http.StatusOK, res)
}

func SearchVMArchive(c *gin.Context) {
  content:= c.Query("content")
  v := vm.Vms_archives{}

  s, err := v.SearchVMArchives(content)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, s)
}

func CreateFlavor(c *gin.Context) {
  cpu, _ := strconv.Atoi(c.Query("cpu"))
  mem, _ := strconv.Atoi(c.Query("mem"))

  f := vm.Vm_flavors{
    Cpu: cpu,
    Mem: mem,
  }

  validate := validator.New()
  err := validate.Struct(f)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  err = f.Createflavor(&f)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func Vmchangeowner(c *gin.Context) {
  uuid := c.PostForm("uuid")
  vmip:= c.PostForm("ip")
  userid, err := strconv.Atoi(c.PostForm("userid"))
  if err != nil {
    vmerror.REQUESTERROR(c, err)
  }

  v := vm.Vms{}
  err = v.ChanegOwner(userid, uuid)
  if err != nil {
    vmerror.SERVERERROR(c, err)
  }

  vd := vdisk.Vms_vdisks{}
  err = vd.ChanegOwner(userid, vmip)
  if err != nil {
    vmerror.SERVERERROR(c, err)
  }

  vmerror.SUCCESS(c, nil)
}


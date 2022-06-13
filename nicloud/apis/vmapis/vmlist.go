package vmapis

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "net/http"
  "nicloud/utils"
  "nicloud/vm"
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
  var res = make(map[string]interface{})
  ct := c.Query("content")
  vms, err := vm.SearchVm(ct)

  if err != nil {
    c.JSON(200, res)
  }

  res["res"] = vms
  c.JSON(200, res)
}

func GetVminfo(c *gin.Context) {
  var res = make(map[string]interface{})
  uuid := c.Query("uuid")
  iplist, err := vm.GetVmByUuid(uuid)

  res["res"] = iplist
  res["err"] = nil
  if err != nil {
    res["err"] = err
    c.JSON(200, res)
    return
  }

  c.JSON(200, res)
}

func GetVmStatus(c *gin.Context) {
  var res = make(map[string]interface{})
  //之前把res（map） 放到了公共变量， 导致获取vmlist , 还有获取vm状态 都要使用这个res（map）， 并发访问就会 gouteline泄漏

  var mux sync.Mutex
  host := c.Query("host")
  uuid := c.Query("uuid")

  mux.Lock()
  vmstate, err := vm.VmStatus(uuid, host)

  res["res"] = vmstate
  res["err"] = err
  mux.Unlock()

  c.JSON(200, res)
}

func Getvmlist(c *gin.Context) {
  var res = make(map[string]interface{})
  start, err := strconv.Atoi(c.Query("start"))
  item := c.Query("item")
  if err != nil {
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(200, res)
    return
  }

  token := c.Request.Header.Get("token")
  userid, err := utils.ParseToken(token)
  if err != nil {
    res["err"] = vmerror.Error{Message: "认证失败"}
    c.JSON(200, res)
    return
  }

	pagenumber, vmcount,  err := vm.Getpagenumber(userid)
	if err != nil {
    res["res"] = nil
    res["err"] = err
    c.JSON(200, res)
    return
  }

  vmlist, err := vm.VmList(userid, start, item)

	res["res"] = vmlist
  res["pagenumber"] = pagenumber
  res["vmcount"] = vmcount
  res["err"] = err

	c.JSON(200, res)
}

func MigrateVmlive(c *gin.Context) {
  var res = make(map[string]interface{})
  uuid := c.Query("uuid")
  migratehost := c.Query("migratehost")

  vmlist := vm.MigrateVmlive(uuid, migratehost)
  res["res"] = vmlist

  c.JSON(200, res)
}

func MigrateVm(c *gin.Context) {
  var res = make(map[string]interface{})
  uuid := c.Query("uuid")
  migratehost := c.Query("migratehost")

  vmlist := vm.MigrateVm(uuid, migratehost)
  res["res"] = vmlist

  c.JSON(200, res)
}

func Createvm(c *gin.Context) {
  var res = make(map[string]interface{})
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
    res["err"] = vmerror.Error{Message: "认证失败"}
    c.JSON(200, res)
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
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(400, res)
    return
  }
  var rwLock sync.RWMutex
  rwLock.Lock()
  err = v.Create(datacenter, storage, vlan, uint(cpu), uint(mem), ip, host, osid, user, comment)
  rwLock.Unlock()

  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: "创建失败: " + err.Error()}
  }
  c.JSON(200, res)
}

func Addcomment(c *gin.Context) {
  var res = make(map[string]interface{})
  uuid := c.Query("uuid")
  comment := c.Query("comment")
  r, err := vm.Updatecomments(uuid, comment)

  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}

func GetFlavor(c *gin.Context) {
  var res = make(map[string]interface{})
	s, err := vm.Flavor()
	res["res"] = s
  res["err"] = nil

	if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
		c.JSON(200, res)
	}
	c.JSON(200, res)
}

func Changeconfig(c *gin.Context) {
  var res = make(map[string]interface{})
  id := c.Query("uuid")
  host := c.Query("host")
  vmhost := c.Query("vmhost")
  cpu, err := strconv.Atoi(c.Query("cpu"))
  if err != nil {
    c.JSON(400, vmerror.Error{Message: "参数错误"})
    return
  }

  oldcpu, err := strconv.Atoi(c.Query("oldcpu"))
  if err != nil {
    c.JSON(400, vmerror.Error{Message: "参数错误"})
    return
  }

  oldmem, err := strconv.Atoi(c.Query("oldmem"))
  if err != nil {
    c.JSON(400, vmerror.Error{Message: "参数错误"})
    return
  }

  mem, err := strconv.Atoi(c.Query("mem"))
  if err != nil {
    c.JSON(400, vmerror.Error{Message: "参数错误"})
    return
  }
  err = vm.Changeconfig(id, host, uint(cpu), uint(oldcpu), uint(mem), uint(oldmem), vmhost)

  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
  }
  c.JSON(200, res)
}

func DeleteVM(c *gin.Context) {
  var res = make(map[string]interface{})
	uuid := c.Query("uuid")
  //datacenter := c.Query("datacenter")
  storage := c.Query("storage")

  var rwLock sync.RWMutex
	rwLock.Lock()
	err := vm.Delete(uuid, storage)
	rwLock.Unlock()

	res["err"] = nil
	if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
  }
	c.JSON(200, res)
}

func Operation(c *gin.Context) {
  var res = make(map[string]interface{})
	uuid := c.Query("uuid")
	host := c.Query("host")

	var err error

	o, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, res)
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

	res["res"] = s
	res["err"] = nil
	if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
  }
	c.JSON(http.StatusOK, res)
}

func Rebuild(c *gin.Context)  {
  var res = make(map[string]interface{})
  uuid := c.Query("uuid")
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")
  osid, _ := strconv.Atoi(c.Query("osname"))
  host := c.Query("host")

  v := vm.Vms{}
  err := v.Rebuildimg(osid, storage, datacenter, uuid, host)

  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: "重置失败: " + err.Error()}
    c.JSON(http.StatusInternalServerError, res)
  }
  c.JSON(http.StatusOK, res)
}

func Createsnap(c *gin.Context)  {
  var res = make(map[string]interface{})
  var err error

  token := c.Request.Header.Get("token")
  userid, err := utils.ParseToken(token)
  if err != nil {
    res["err"] = vmerror.Error{Message: "认证失败"}
    c.JSON(http.StatusOK, res)
    return
  }

  snapname := c.PostForm("snapname")

  if len(snapname) == 0 {
    c.JSON(http.StatusBadRequest, res)
    return
  }

  uuid := c.PostForm("uuid")
  datacenter := c.PostForm("datacenter")
  storage := c.PostForm("storage")
  protect, err := strconv.ParseBool(c.PostForm("protect"))
  if err != nil {
    c.JSON(http.StatusBadRequest, res)
    return
  }

  if protect == false {
    err = vm.CreatSnap(uuid, datacenter, storage, snapname)
  } else {
    err = vm.SaveSnapToImg(uuid, datacenter, storage, snapname, userid)
  }

  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: "创建失败: " + err.Error()}
  }

  c.JSON(http.StatusOK, res)
}

func Getsnap(c *gin.Context)  {
  var res = make(map[string]interface{})
  uuid := c.Query("uuid")
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")

  s, err := vm.Getsnap(datacenter, storage, uuid)
  res["res"] = s

  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
  }
  c.JSON(http.StatusOK, res)
}

func Rollback(c *gin.Context)  {
  var res = make(map[string]interface{})
  uuid := c.Query("uuid")
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")
  snapname := c.Query("snapname")

  err := vm.RollbackSnap(uuid, snapname,  datacenter, storage)

  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
  }
  c.JSON(http.StatusOK, res)
}

func DelSnap(c *gin.Context)  {
  var res = make(map[string]interface{})
  uuid := c.Query("uuid")
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")
  snapname := c.Query("snapname")

  err := vm.DelSnap(uuid, snapname,  datacenter, storage)

  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: "删除快照失败: " + err.Error()}
  }

  c.JSON(http.StatusOK, res)
}

func GetVmArchive(c *gin.Context)  {
  res := make(map[string]interface{})
  startpage, err := strconv.Atoi(c.Query("startpage"))
  if err != nil {
    res["err"] = vmerror.Error{Message: "认证失败"}
    c.JSON(http.StatusBadRequest, res)
    return
  }
  ar := vm.Vms_archives{}

  token := c.Request.Header.Get("token")
  userid, err := utils.ParseToken(token)
  if err != nil {
    res["err"] = vmerror.Error{Message: "认证失败"}
    c.JSON(http.StatusOK, res)
    return
  }
  pagenumber, vmcount,  err := vm.Getvmarchivepagenumber(userid)
  if err != nil {
    res["res"] = nil
    res["err"] = err
    c.JSON(http.StatusInternalServerError, res)
    return
  }

  r, err := ar.GetVmArchive(startpage)

  res["res"] = r
  res["pagenumber"] = pagenumber
  res["vmcount"] = vmcount
  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
    c.JSON(http.StatusInternalServerError, res)
  }

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
  var res = make(map[string]interface{})
  content:= c.Query("content")

  v := vm.Vms_archives{}

  s, err := v.SearchVMArchives(content)
  res["err"] = nil
  res["res"] = s
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
    c.JSON(http.StatusInternalServerError, res)
    return
  }
  c.JSON(http.StatusOK, res)
}

func CreateFlavor(c *gin.Context) {
  var res = make(map[string]interface{})
  cpu, _ := strconv.Atoi(c.Query("cpu"))
  mem, _ := strconv.Atoi(c.Query("mem"))

  f := vm.Vm_flavors{
    Cpu: cpu,
    Mem: mem,
  }

  validate := validator.New()
  err := validate.Struct(f)
  if err != nil {
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(http.StatusBadRequest, res)
    return
  }

  err = f.Createflavor(&f)
  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
    c.JSON(http.StatusInternalServerError, res)
    return
  }
  c.JSON(http.StatusOK, res)
}

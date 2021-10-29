package vmapis

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "goblog/vmcommon"
  "strconv"
)

func Vnc(c *gin.Context)  {
  uuid := c.Query("uuid")
  host := c.Query("host")
  res := fmt.Sprintf("http://%s/novnc/vnc.html?path=websockify/?token=%s", host, uuid)
  c.JSON(200, res)
}

func Search(c *gin.Context)  {
  ct := c.Query("content")
  vms, err := vmcommon.SearchVm(ct)
  res := make(map[string]interface{})
  if err != nil {
    c.JSON(200, res)
  }

  res["res"] = vms
  c.JSON(200, res)
}

func GetVminfo(c *gin.Context) {
  uuid := c.Query("uuid")
  iplist := vmcommon.GetVmByUuid(uuid)
  res := make(map[string]interface{})
  res["res"] = iplist

  c.JSON(200, res)
}


func GetVmStatus(c *gin.Context) {
  host := c.Query("host")
  uuid := c.Query("uuid")

  res := make(map[string]interface{})
  vmstate, err := vmcommon.VmStatus(uuid, host)

  if err != nil {
    res["res"] = vmstate
  }

  res["res"] = vmstate
  c.JSON(200, res)
}

func Getvmlist(c *gin.Context) {
	host := c.Query("host")
	vmlist := vmcommon.VmList(host)
	res := make(map[string]interface{})
	res["res"] = vmlist

	c.JSON(200, res)
}

func MigrateVm(c *gin.Context) {
  uuid := c.Query("uuid")
  migratehost := c.Query("migratehost")

  vmlist := vmcommon.MigrateVm(uuid, migratehost)
  res := make(map[string]interface{})
  res["res"] = vmlist

  c.JSON(200, res)
}

func Createvm(c *gin.Context) {
	cpu, _ := strconv.Atoi(c.Query("cpu"))
	mem, _ := strconv.Atoi(c.Query("mem"))
	ip := c.Query("ip")
	mac := c.Query("mac")
	host := c.Query("host")
  image := c.Query("image")
	create, err := vmcommon.Create(cpu, mem, ip, mac, host, image)
  res := make(map[string]interface{})

	res["res"] = create
  res["err"] = err

	c.JSON(200, res)
}

func GetImage(c *gin.Context) {

  res := make(map[string]interface{})
  r, err := vmcommon.GetImages()

  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}

func Addcomment(c *gin.Context) {
  uuid := c.Query("uuid")
  comment := c.Query("comment")
  res := make(map[string]interface{})
  r, err := vmcommon.Updatecomments(uuid, comment)

  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}

func GetFlavor(c *gin.Context) {
	res := make(map[string]interface{})
	s, err := vmcommon.Flavor()
	res["res"] = s
	res["err"] = err
	if err != nil {
		c.JSON(200, res)
	}

	c.JSON(200, res)
}

func DeleteVM(c *gin.Context) {
	uuid := c.Query("uuid")

	res := make(map[string]interface{})
	r, err := vmcommon.Delete(uuid)

	res["res"] = r
	res["err"] = err
	c.JSON(200, res)
}

func Operation(c *gin.Context) {
	uuid := c.Query("uuid")
	host := c.Query("host")
	res := make(map[string]interface{})

	var err error

	o, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, res)
	}

	var s *vmcommon.Vms
	switch o {
	case 0:
		s, err = vmcommon.Shutdown(uuid, host)
	case 1:
		s, err = vmcommon.Start(uuid, host)
  case 3:
    s, err = vmcommon.PauseVm(uuid, host)
	}

	res["res"] = s
	res["err"] = err
	c.JSON(200, res)
}

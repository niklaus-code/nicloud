package vmapis

import (
	"goblog/vmcommon"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetHosts(c *gin.Context) {
	hostlist := vmcommon.Hosts()
	res := make(map[string]interface{})
	res["res"] = hostlist

	c.JSON(200, res)
}

func GetIplist(c *gin.Context) {
	iplist := vmcommon.IPlist()
	res := make(map[string]interface{})
	res["res"] = iplist

	c.JSON(200, res)
}

func Getvmlist(c *gin.Context) {
	host := c.Query("host")
	vmlist := vmcommon.VmList(host)
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

	create, err := vmcommon.Create(cpu, mem, ip, mac, host)
	res := make(map[string]interface{})
	res["res"] = create
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
	ip := c.Query("ip")
	host := c.Query("host")

	res := make(map[string]interface{})
	r, err := vmcommon.Delete(uuid, ip, host)

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
	}

	res["res"] = s
	res["err"] = err
	c.JSON(200, res)
}

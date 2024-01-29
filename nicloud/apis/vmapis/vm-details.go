package vmapis

import (
	"github.com/gin-gonic/gin"
	"nicloud/vm"
	"nicloud/vmerror"
)

func Cpudetails(c *gin.Context) {
	uuid := c.Query("uuid")
	host := c.Query("host")

	cpu, err := vm.Cpuinfo(host, uuid)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	vmerror.SUCCESS(c, cpu)
}

func Memdetails(c *gin.Context) {
	uuid := c.Query("uuid")
	host := c.Query("host")
	m, err := vm.Meminfo(host, uuid)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	vmerror.SUCCESS(c, m)
}

func Netdetails(c *gin.Context) {
	uuid := c.Query("uuid")
	host := c.Query("host")

	n, err := vm.Netinfo(host, uuid)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	vmerror.SUCCESS(c, n)
}

func Diskdetails(c *gin.Context) {
	uuid := c.Query("uuid")
	host := c.Query("host")

	dio, err := vm.Diskinfo(host, uuid, "linux")
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	vmerror.SUCCESS(c, dio)
}

package cephapis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"nicloud/cephcommon"
	"nicloud/vmerror"
)

func Getpool(c *gin.Context) {
	datacenter := c.Query("datacenter")
	storage := c.Query("storage")
	cephinfo, err := cephcommon.Getpool(datacenter, storage)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}
	vmerror.SUCCESS(c, cephinfo)
}

func GetStorage(c *gin.Context) {
	var ceph cephcommon.Vms_Ceph
	cephinfo, err := ceph.Get()
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}
	vmerror.SUCCESS(c, cephinfo)
}

func Delete(c *gin.Context) {
	name := c.Query("name")
	err := cephcommon.Delete(name)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}
	vmerror.SUCCESS(c, nil)
}

func Addceph(c *gin.Context) {
	uuid := c.PostForm("uuid")
	name := c.PostForm("storagename")
	pool := c.PostForm("pool")
	datacenter := c.PostForm("datacenter")
	ceph_secret := c.PostForm("ceph_secret")
	port := c.PostForm("port")
	ips := c.PostForm("ips")
	comment := c.PostForm("comment")

	ceph := cephcommon.Vms_Ceph{
		Uuid: uuid,
		Name: name,
		Pool: pool,
		//Contain: contain,
		Datacenter:  datacenter,
		Ceph_secret: ceph_secret,
		Ips:         ips,
		Port:        port,
	}

	validate := validator.New()
	err := validate.Struct(&ceph)
	if err != nil {
		vmerror.REQUESTERROR(c, err)
		return
	}

	err = ceph.Add(uuid, name, pool, datacenter, ceph_secret, ips, port, comment)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}
	vmerror.SUCCESS(c, nil)
}

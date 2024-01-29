package osimage

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"nicloud/cephcommon"
	"nicloud/osimage"
	"nicloud/utils"
	"nicloud/vm"
	"nicloud/vmerror"
	"strconv"
)

var ceph cephcommon.Vms_Ceph

func DelImage(c *gin.Context) {
	osid, err := strconv.Atoi(c.Query("osid"))
	if err != nil {
		vmerror.REQUESTERROR(c, err)
		return
	}

	checkvmsandos, err := vm.GetVmbyOsId(osid)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	if checkvmsandos {
		err := osimage.Del(osid)
		if err != nil {
			vmerror.SERVERERROR(c, err)
			return
		}
		vmerror.SUCCESS(c, nil)
		return
	}

	vmerror.SERVERERROR(c, vmerror.Error{Message: "有关联云主机，无法删除"})
}

func UpdateImage(c *gin.Context) {
	id, errparam := strconv.Atoi(c.PostForm("id"))
	if errparam != nil {
		vmerror.REQUESTERROR(c, errparam)
		return
	}
	datacenter := c.PostForm("datacenter")
	storage := c.PostForm("storage")
	osname := c.PostForm("osname")
	snapname := c.PostForm("snapimage")
	cephblockdevice := c.PostForm("cephblockdevice")
	xml, _ := strconv.Atoi(c.PostForm("xml"))
	tag, _ := strconv.Atoi(c.PostForm("tag"))
	sort, _ := strconv.Atoi(c.PostForm("sort"))

	o := osimage.Vms_os{
		Id:              id,
		Datacenter:      datacenter,
		Storage:         storage,
		Osname:          osname,
		Snapimage:       snapname,
		Cephblockdevice: cephblockdevice,
		Xml:             xml,
		Tag:             tag,
		Sort:            sort,
	}

	validate := validator.New()
	err := validate.Struct(o)

	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	err = osimage.Update(id, datacenter, storage, osname, snapname, cephblockdevice, xml)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	vmerror.SUCCESS(c, nil)
}

func GetImage(c *gin.Context) {
	token := c.Request.Header.Get("token")
	user, err := utils.ParseToken(token)

	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	sort, err := strconv.Atoi(c.Query("sort"))
	if err != nil {
		vmerror.REQUESTERROR(c, err)
		return
	}
	r, err := osimage.Maposimage(user, sort)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}
	vmerror.SUCCESS(c, r)
}

func GetImageSort(c *gin.Context) {
	r, err := osimage.Get_osimage_sort()
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}
	vmerror.SUCCESS(c, r)
}

func GetOsTag(c *gin.Context) {
	t := osimage.Vms_os_tags{}
	r, err := t.Getostags()
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}
	vmerror.SUCCESS(c, r)
}

func GetImageby(c *gin.Context) {
	datacenter := c.Query("datacenter")
	storage := c.Query("storage")

	r, err := osimage.Getimageby(datacenter, storage)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}
	vmerror.SUCCESS(c, r)
}

func GetImagebytag(c *gin.Context) {
	datacenter := c.Query("datacenter")
	storage := c.Query("storage")
	tag := c.Query("tag")

	r, err := osimage.Getimagebytag(datacenter, storage, tag)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}
	vmerror.SUCCESS(c, r)
}

func AddImage(c *gin.Context) {
	datacenter := c.PostForm("datacenter")
	storage := c.PostForm("storage")
	osname := c.PostForm("osname")
	tag, _ := strconv.Atoi(c.PostForm("tag"))
	createsnap, _ := strconv.ParseBool(c.PostForm("createsnap"))

	cephblockdevice := c.PostForm("cephblockdevice")
	xml, _ := strconv.Atoi(c.PostForm("xml"))
	sort, _ := strconv.Atoi(c.PostForm("ossort"))

	token := c.Request.Header.Get("token")
	user, err := utils.ParseToken(token)

	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	o := osimage.Vms_os{
		Sort:            sort,
		Owner:           user,
		Datacenter:      datacenter,
		Storage:         storage,
		Osname:          osname,
		Cephblockdevice: cephblockdevice,
		Xml:             xml,
		Tag:             tag,
	}

	validate := validator.New()
	err = validate.Struct(o)
	if err != nil {
		vmerror.REQUESTERROR(c, err)
		return
	}

	snap := ""
	if createsnap {
		storageinfo, err := ceph.Cephinfobyuuid(storage)
		if err != nil {
			vmerror.SERVERERROR(c, err)
			return
		}

		ceph := cephcommon.Vms_Ceph{}
		snap, err = ceph.CreateSnapAndProtect(storageinfo.Pool, cephblockdevice)
		if err != nil {
			vmerror.SERVERERROR(c, err)
			return
		}
	}

	err = o.Add(datacenter, storage, osname, cephblockdevice, xml, sort, user, snap, tag)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	vmerror.SUCCESS(c, nil)
}

func Addosimagexml(c *gin.Context) {
	sort, _ := strconv.Atoi(c.PostForm("tag"))
	xml := c.PostForm("xml")
	comment := c.PostForm("comment")

	x := osimage.Vms_osimage_xmls{
		Sort:    sort,
		Xml:     xml,
		Comment: comment,
	}
	err := x.Addxml(&x)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	vmerror.SUCCESS(c, nil)
}

func Getosimagexml(c *gin.Context) {
	data, err := osimage.Maposimagexml()
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	vmerror.SUCCESS(c, data)
}

func Delosimagexml(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	x := osimage.Vms_osimage_xmls{}
	err := x.Delxml(id)
	if err != nil {
		vmerror.SERVERERROR(c, err)
	}

	vmerror.SUCCESS(c, nil)
}

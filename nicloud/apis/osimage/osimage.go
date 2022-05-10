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
  res := make(map[string]interface{})
  osid, err := strconv.Atoi(c.Query("osid"))
  if err != nil {
    res["err"] = nil
    c.JSON(400, res)
    return
  }

  checkvmsandos, err := vm.GetVmbyOsId(osid)
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
    c.JSON(200, res)
    return
  }

  if checkvmsandos {
    err := osimage.Del(osid)
    if err != nil {
      res["err"] = vmerror.Error{Message: err.Error()}
    }
    res["err"] = nil
  } else {
    res["err"] = vmerror.Error{Message: "有关联云主机，无法删除"}
  }

  c.JSON(200, res)
}

func UpdateImage(c *gin.Context) {
  res := make(map[string]interface{})
  id, errparam := strconv.Atoi(c.PostForm("id"))
  if errparam != nil {
    res["res"] = vmerror.Error{Message: "param err"}
    c.JSON(400, res)
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
    Id: id,
    Datacenter: datacenter,
    Storage: storage,
    Osname: osname,
    Snapimage: snapname,
    Cephblockdevice: cephblockdevice,
    Xml: xml,
    Tag: tag,
    Sort: sort,
  }

  validate := validator.New()
  err := validate.Struct(o)

  if err != nil {
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(400, res)
    return
  }

  err = osimage.Update(id, datacenter, storage, osname, snapname, cephblockdevice, xml)

  res["err"] = err
  c.JSON(200, res)
}

func GetImage(c *gin.Context) {
  res := make(map[string]interface{})
  token := c.Request.Header.Get("token")
  user, err := utils.ParseToken(token)

  if err != nil {
    res["err"] = vmerror.Error{Message: "认证失败"}
    c.JSON(200, res)
    return
  }

  sort,_ := strconv.Atoi(c.Query("sort"))
  r, err := osimage.Maposimage(user, sort)

  res["res"] = r
  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
  }

  c.JSON(200, res)
}


func GetImageSort(c *gin.Context) {

  res := make(map[string]interface{})
  r, err := osimage.Get_osimage_sort()

  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}

func GetOsTag(c *gin.Context) {
  t := osimage.Vms_os_tags{}

  res := make(map[string]interface{})
  r, err := t.Getostags()

  res["res"] = r
  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
  }
  c.JSON(200, res)
}

func GetImageby(c *gin.Context) {
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")

  res := make(map[string]interface{})
  r, err := osimage.Getimageby(datacenter, storage)

  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}

func GetImagebytag(c *gin.Context) {
  datacenter := c.Query("datacenter")
  storage := c.Query("storage")
  tag := c.Query("tag")

  res := make(map[string]interface{})
  r, err := osimage.Getimagebytag(datacenter, storage, tag)

  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}

func AddImage(c *gin.Context) {
  res := make(map[string]interface{})

  datacenter := c.PostForm("datacenter")
  storage := c.PostForm("storage")
  osname := c.PostForm("osname")
  tag, _ := strconv.Atoi(c.PostForm("tag"))
  createsnap, err := strconv.ParseBool(c.PostForm("createsnap"))
  if err != nil {
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(200, res)
    return
  }

  cephblockdevice := c.PostForm("cephblockdevice")
  xml, _ := strconv.Atoi(c.PostForm("xml"))
  sort,_ := strconv.Atoi(c.PostForm("ossort"))

  token := c.Request.Header.Get("token")
  user, err := utils.ParseToken(token)

  if err != nil {
    res["err"] = vmerror.Error{Message: "认证失败"}
    c.JSON(200, res)
    return
  }

  o := osimage.Vms_os{
    Sort: sort,
    Owner: user,
    Datacenter: datacenter,
    Storage: storage,
    Osname: osname,
    Cephblockdevice: cephblockdevice,
    Xml: xml,
    Tag: tag,
  }

  validate := validator.New()
  err = validate.Struct(o)
  if err != nil {
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(400, res)
    return
  }

  snap := ""
  if createsnap {
    storageinfo, err := ceph.Cephinfobyuuid(storage)
    if err != nil {
      res["err"] = vmerror.Error{Message: err.Error()}
      c.JSON(200, res)
      return
    }

    ceph := cephcommon.Vms_Ceph{}
    snap, err = ceph.CreateSnapAndProtect(storageinfo.Pool, cephblockdevice)
    if err != nil {
      res["err"] = vmerror.Error{Message: err.Error()}
      c.JSON(200, res)
      return
    }
  }

  err = o.Add(datacenter, storage, osname, cephblockdevice,  xml, sort , user, snap, tag)

  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: "创建失败: "+ err.Error()}
  }

  c.JSON(200, res)
}

func Addosimagexml(c *gin.Context) {
  res := make(map[string]interface{})

  sort, _ := strconv.Atoi(c.PostForm("tag"))
  xml := c.PostForm("xml")
  comment := c.PostForm("comment")

  x := osimage.Vms_osimage_xmls{
    Sort: sort,
    Xml: xml,
    Comment: comment,
  }
  err := x.Addxml(&x)
  if err != nil {
    res["err"] = vmerror.Error{Message:  err.Error()}
    c.JSON(200, res)
    return
  }

  res["err"] = nil
  c.JSON(200, res)
}

func Getosimagexml(c *gin.Context) {
  res := make(map[string]interface{})

  data, err := osimage.Maposimagexml()
  if err != nil {
    res["err"] = vmerror.Error{Message:  err.Error()}
    c.JSON(200, res)
    return
  }

  res["err"] = nil
  res["res"] = data
  c.JSON(200, res)
}

func Delosimagexml(c *gin.Context) {
  res := make(map[string]interface{})
  id, _ := strconv.Atoi(c.Query("id"))
  x := osimage.Vms_osimage_xmls{}
  err := x.Delxml(id)
  if err != nil {
    res["err"] = vmerror.Error{Message:  err.Error()}
    c.JSON(200, res)
    return
  }

  res["err"] = nil
  c.JSON(200, res)
}

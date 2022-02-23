package osimage

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "nicloud/cephcommon"
  "nicloud/osimage"
  "nicloud/utils"
  "nicloud/vm"
  "nicloud/vmerror"
  "strconv"
)

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
  xml := c.PostForm("xml")

  o := osimage.Vms_os{
    Id: id,
    Datacenter: datacenter,
    Storage: storage,
    Osname: osname,
    Snapimage: snapname,
    Cephblockdevice: cephblockdevice,
    Xml: xml,
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

  var r []map[string]interface{}

  r, err = osimage.Maposimage(user, sort)
  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}


func GetImageSort(c *gin.Context) {

  res := make(map[string]interface{})
  r, err := osimage.Get_osimage_sort()

  res["res"] = r
  res["err"] = err
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

func AddImage(c *gin.Context) {
  res := make(map[string]interface{})

  datacenter := c.PostForm("datacenter")
  storage := c.PostForm("storage")
  osname := c.PostForm("osname")
  createsnap, err := strconv.ParseBool(c.PostForm("createsnap"))
  if err != nil {
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(200, res)
    return
  }

  cephblockdevice := c.PostForm("cephblockdevice")
  xml :=c.PostForm("xml")
  sort,_ := strconv.Atoi(c.PostForm("sort"))

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
    storageinfo, err := cephcommon.Cephinfobyuuid(datacenter, storage)
    if err != nil {
      res["err"] = vmerror.Error{Message: err.Error()}
      c.JSON(200, res)
      return
    }
    snap, err = cephcommon.CreateSnapAndProtect(storageinfo.Pool, cephblockdevice)
    if err != nil {
      res["err"] = vmerror.Error{Message: err.Error()}
      fmt.Println(err)
      c.JSON(200, res)
      return
    }
  }

  os := osimage.Vms_os{}
  err = os.Add(datacenter, storage, osname, cephblockdevice,  xml, sort , user, snap)

  res["err"] = nil
  if err != nil {
    res["err"] = vmerror.Error{Message: "创建失败: "+ err.Error()}
  }

  c.JSON(200, res)
}

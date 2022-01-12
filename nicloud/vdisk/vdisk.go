package vdisk

import (
  "fmt"
  "github.com/beevik/etree"
  "nicloud/cephcommon"
  db "nicloud/dbs"
  "nicloud/libvirtd"
  "nicloud/users"
  "nicloud/utils"
  "nicloud/vmerror"
  "reflect"
  "strings"
  "time"
)

type Vms_vdisks struct {
  Vdiskid string
  Contain int `json:"contain" validate:"min=0,max=1024"`
  Diskname string
  Pool string `json:"pool" validate:"required"`
  Storage string  `json:"storage" validate:"required"`
  Datacenter string `json:"datacenter" validate:"required"`
  Vm_ip string
  User int `json:"user" validate:"required"`
  Exist int8
  Status int
  Comment string `json:"Comment" validate:"required"`
  Createtime string
}

func Getdiskbyvm(vmip string) ([]*Vms_vdisks, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  c := []*Vms_vdisks{}
  dbs.Select("contain, diskname").Where("vm_ip=?", vmip).Find(&c)
  return c, nil
}

func UpdateMountvmstatus(datacenter string, storage string, vdiskid string, vmip string, diskname string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  errdb := dbs.Model(Vms_vdisks{}).Where("datacenter=? and storage=? and vdiskid=?", datacenter, storage, vdiskid).Update(map[string]interface{}{"vm_ip": vmip, "status": 0, "diskname": diskname})
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}

func Updatevdiskbydelvm(datacenter string, storage string, vmip string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  errdb := dbs.Model(Vms_vdisks{}).Where("datacenter=? and storage=? and vm_ip=?", datacenter, storage, vmip).Update(map[string]interface{}{"vm_ip": "", "status": 1, "diskname": ""})
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}


func Create_vdisk(contain int, pool string, storage string, datacenter string, userid int, comment string) error {
  vdiskid := utils.Createuuid()
  c := &Vms_vdisks{
    Vdiskid: vdiskid,
    Contain: contain,
    Pool: pool,
    Storage: storage,
    Datacenter: datacenter,
    User: userid,
    Exist: 1,
    Status: 1,
    Comment: comment,
    Createtime: time.Now().Format("2006-01-02 15:04:05"),
  }

  err := cephcommon.Createcephblock(vdiskid, contain, pool)
  if err != nil {
    return err
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  errdb := dbs.Create(&c)
  if errdb.Error != nil {
    return vmerror.Error{Message: errdb.Error.Error()}
  }
  return err
}

type Vms_vdisks_archives struct {
  Vdiskid string
  Owner int
  Pool string
  Storage string
  Datacenter string
  Comment string
  Create_time time.Time
}

func addiskachives(uuid string, pool string, storage string, datacenter string, ownerid int, comment string) (string, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return "", err
  }

  diskachi := &Vms_vdisks_archives{
    Vdiskid: uuid,
    Owner: ownerid,
    Pool: pool,
    Storage: storage,
    Datacenter: datacenter,
    Comment: comment,
    Create_time: time.Now(),
  }

  errdb := dbs.Create(diskachi)
  if errdb.Error != nil {
    return "", errdb.Error
  }
  return diskachi.Vdiskid, nil
}

func deletedisk(uuid string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb := dbs.Where("uuid=?").Delete(&Vms_vdisks_archives{})
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}

func getdiskinfobyid(uuid string) (*Vms_vdisks, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  vdiskinfo := &Vms_vdisks{}
  errdb := dbs.Where("Vdiskid=?", uuid).First(vdiskinfo)
  if errdb.Error != nil {
    return nil, errdb.Error
  }
  return vdiskinfo, err
}

func Deletevdisk(uuid string, comment string) error {
  checkmount, err := Getdiskstatus(uuid)
  if err != nil {
    return err
  }

  if checkmount == 0 {
    return vmerror.Error{Message: "硬盘已挂载，请卸载后删除"}
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  vdiskinfo, err := getdiskinfobyid(uuid)
  if err != nil {
    return err
  }

  addachives, err := addiskachives(uuid, vdiskinfo.Pool, vdiskinfo.Storage, vdiskinfo.Datacenter, vdiskinfo.User, comment)
  if err != nil {
    return err
  }

  errdb := dbs.Where("vdiskid=?", uuid).Delete(Vms_vdisks{})
  if errdb.Error != nil {
    err = deletedisk(addachives)
    if err != nil {
      return err
    }
    return vmerror.Error{Message: "删除硬盘失败"}
  }

  err = cephcommon.Rm_image(uuid, vdiskinfo.Pool)
  if err != nil {
    return vmerror.Error{Message: "删除块设备失败"}
  }
  return nil
}

func Umountdisk(vmip string, storage string, datacenter string, vdiskid string, xml string, host string, vms interface{}) error {
  checkmount, err := checkmount(vdiskid)
  if err != nil {
    return err
  }
  if checkmount == 1 {
    return vmerror.Error{Message: "vdisk has been mouunted"}
  }


  doc := etree.NewDocument()
  err = doc.ReadFromString(xml)
  if err != nil {
    return err
  }
  device := doc.FindElements("./domain/devices/disk")
  d:= doc.FindElement("./domain/devices/")
  for _, v := range device {
    source := v.FindElement("./source")
    vmdisk := source.SelectAttr("name").Value
    uuid := strings.Split(vmdisk, "/")

    if len(uuid)> 1 && uuid[1] == vdiskid {
      d.RemoveChild(v)
      var docstring string
      docstring, err = doc.WriteToString()
      libvirtd.DefineVm(docstring, host)

      err := Umountvmstatus(datacenter, storage, vdiskid)
      if err != nil {
        return err
      }

      dbs, err := db.NicloudDb()
      if err != nil {
        return err
      }

      errdb := dbs.Model(vms).Where("ip=?", vmip).Update("vmxml", docstring)
      if errdb.Error != nil {
        return vmerror.Error{Message: errdb.Error.Error()}
      }

      if err != nil {
        return err
      }
      return nil
    }
  }
  return vmerror.Error{
    Message: "disk not found",
  }
}

func Getdiskstatus(uuid string) (int, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return 0, err
  }

  vdisk := &Vms_vdisks{}
  errdb := dbs.Where("vdiskid=?", uuid).First(vdisk)
  if errdb.Error != nil {
    return 0, errdb.Error
  }
  return vdisk.Status, nil
}


func mapvdisk(obj  []Vms_vdisks) ([]map[string]interface{}, error)  {
  var mapc []map[string]interface{}

  for _, v := range obj {
    c := make(map[string]interface{})
    m := reflect.TypeOf(v)
    n := reflect.ValueOf(v)
    for i := 0; i < m.NumField(); i++ {
      c[m.Field(i).Name] = n.Field(i).Interface()
    }

    username, err := users.GetUsernameById(v.User)
    if err != nil {
      return nil, err
    }
    c["username"] = username
    mapc = append(mapc, c)
  }
  return mapc, nil
}

func Getvdisk(userid int) ([]map[string]interface{}, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var c  []Vms_vdisks

  user, err := users.GetUsernameById(userid)
  if err != nil {
    return nil, err
  }
  if user == "admin" {
    dbs.Order("createtime desc").Find(&c)
  } else {
    dbs.Where("user=?", userid).Order("createtime desc").Find(&c)
  }
  mapvdisk, err := mapvdisk(c)
  if err != nil {
    return nil, err
  }
  return mapvdisk, err
}

func Umountvmstatus(datacenter string, storage string, vdiskid string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  errdb := dbs.Model(Vms_vdisks{}).Where("datacenter=? and storage=? and vdiskid=?", datacenter, storage, vdiskid).Update("vm_ip", "").Update("status", 1)
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}

var Disknametype = []string{"vdb", "vdc", "vdd", "vde", "vdf"}
var slot = map[string] int {"vdb": 11, "vdc": 12, "vdd": 13, "vde": 14, "vdf": 15}

func next(items []*Vms_vdisks, item string) bool {
  for _, i := range items {
    c := i
    if c.Diskname == item {
      return true
    }
  }

  return false
}


func namedisk(vmip string) (string, error) {
  disklist, err := Getdiskbyvm(vmip)
  if err != nil {
    return "", err
  }

  for _, a := range Disknametype {
    b := next(disklist, a)
    if (b == false ) {
      return a, err
    }
  }
  return "vdb", err
}

func checkmount(vdiskid string) (int, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return 0, err
  }

  v := &Vms_vdisks{}
  dbs.Where("vdiskid=?", vdiskid).First(v)
  return v.Status, nil
}

func Mountdisk(ip string, vmhost string, storage string, pool string, datacenter string, vdiskid string, vms interface{}, xml string) error {
  checkmount, err := checkmount(vdiskid)
  if err != nil {
    return err
  }
  if checkmount == 0 {
    return vmerror.Error{Message: "vdisk has been mouunted"}
  }

  disknum, err := Getdiskbyvm(ip)
  if err != nil {
    return err
  }

  if len(disknum) >= 5 {
    return vmerror.Error{Message: "Maximum number of mounted to 5"}
  }
  storageinfo, err := cephcommon.Cephinfobyname(datacenter, storage)
  if err != nil {
    return err
  }

  doc := etree.NewDocument()
  err = doc.ReadFromString(xml)
  if err != nil {
    return err
  }
  device := doc.FindElement("./domain/devices")
  disk := device.CreateElement("disk")
  disk.CreateAttr("type", "network")
  disk.CreateAttr("device", "disk")

  driver := disk.CreateElement("driver")
  driver.CreateAttr("name", "qemu")
  driver.CreateAttr("type", "raw")

  auth := disk.CreateElement("auth")
  auth.CreateAttr("username", "admin")
  secret := auth.CreateElement("secret")
  secret.CreateAttr("type", "ceph")
  secret.CreateAttr("uuid", storageinfo.Ceph_secret)

  source := disk.CreateElement("source")
  source.CreateAttr("protocol", "rbd")

  source.CreateAttr("name", fmt.Sprintf("%s/%s",pool, vdiskid))
  host := source.CreateElement("host")

  var iplist []string
  iplist = strings.Split(storageinfo.Ips, ",")
  for _, v := range iplist {
    host.CreateAttr("name", v)
  }
  host.CreateAttr("port", storageinfo.Port)

  diskname, err := namedisk(ip)
  if err != nil {
    return err
  }

  target := disk.CreateElement("target")
  target.CreateAttr("dev", diskname)
  target.CreateAttr("bus", "virtio")

  address := disk.CreateElement("address")
  address.CreateAttr("type", "pci")
  address.CreateAttr("domain", "0x0000")
  address.CreateAttr("bus", "0x00")
  slot := fmt.Sprintf("0x%d", slot[diskname])
  address.CreateAttr("slot", slot)
  address.CreateAttr("function", "0x0")
  doc.Indent(2)
  var docstring string
  docstring, err = doc.WriteToString()

  updatexml := updatexmlbyip(docstring, ip, vms)
  if updatexml != nil {
    return updatexml
  }

  updatevm := UpdateMountvmstatus(datacenter, storage, vdiskid, ip, diskname)
  if updatevm != nil {
    updatexml := updatexmlbyip(xml, ip, vms)
    if updatexml != nil {
      return updatexml
    }
    return updatevm
  }

  err = libvirtd.DefineVm(docstring, vmhost)
  if err != nil {
    updatexml := updatexmlbyip(xml, ip, vms)
    if updatexml != nil {
      return updatexml
    }
    updatevm := UpdateMountvmstatus(datacenter, storage, vdiskid, "", "")
    if updatevm != nil {
      return updatevm
    }
    return err
  }

  return nil
}

func updatexmlbyip(xml string, ip string, vms interface{}) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb:= dbs.Model(vms).Where("ip=?", ip).Update("vmxml", xml)
  if errdb.Error != nil {
    return vmerror.Error{Message: errdb.Error.Error()}
  }

  return nil
}


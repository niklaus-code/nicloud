package vdisk

import (
  "fmt"
  "github.com/beevik/etree"
  "goblog/ceph"
  db "goblog/dbs"
  "goblog/libvirtd"
  "goblog/utils"
  "goblog/vmerror"
  "strings"
)

type Vms_vdisks struct {
  Vdiskid string
  Contain int
  Diskname string
  Pool string
  Storage string
  Datacenter string
  Vm_ip string
  User string
  Exist int8
  Status int
}

func Getdiskbyvm(vmip string) ([]*Vms_vdisks) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return []*Vms_vdisks{}
  }
  c := []*Vms_vdisks{}
  dbs.Select("contain, diskname").Where("vm_ip=?", vmip).Find(&c)
  return c
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
  errdb := dbs.Model(Vms_vdisks{}).Where("datacenter=? and storage=?", datacenter, storage).Update("vm_ip", "").Update("status", 1).Where("vm_ip=?", vmip)
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}


func Add_cloudrive(contain int, pool string, storage string, datacenter string, user string) ([]*Vms_vdisks, error) {
  vdiskid := utils.Createuuid()
  c := &Vms_vdisks{
    Vdiskid: vdiskid,
    Contain: contain,
    Pool: pool,
    Storage: storage,
    Datacenter: datacenter,
    User: "niklaus",
    Exist: 1,
    Status: 1,
  }

  err := ceph.Createcephblock(vdiskid, contain)
  if err != nil {
    return nil, err
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  errdb := dbs.Create(&c)
  if errdb.Error != nil {
    return nil, vmerror.Error{Message: errdb.Error.Error()}
  }
  return nil, err
}

type Vms_vdisks_archives struct {
  Vdiskid string
  Pool string
  Storage string
  Datacenter string
}

func adddiskachives(uuid string, pool string, storage string, datacenter string) (string, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return "", err
  }

  diskachi := &Vms_vdisks_archives{
    Vdiskid: uuid,
    Pool: pool,
    Storage: storage,
    Datacenter: datacenter,
  }

  errdb := dbs.Create(diskachi)
  if errdb.Error != nil {
    return "", errdb.Error
  }
  return diskachi.Vdiskid, nil
}

func deletediskachives(uuid string) error {
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

func Deletevdisk(uuid string) error {
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

  addachives, err := adddiskachives(uuid, vdiskinfo.Pool, vdiskinfo.Storage, vdiskinfo.Datacenter)
  if err != nil {
    return err
  }

  errdb := dbs.Where("vdiskid=?", uuid).Delete(Vms_vdisks{})
  if errdb.Error != nil {
    err = deletediskachives(addachives)
    if err != nil {
      return err
    }
    return vmerror.Error{Message: "delete vdisk fail"}
  }

  err = ceph.Rm_image(uuid)
  if err != nil {
    return err
  }
  return nil
}

func Umountdisk(vmip string, storage string, datacenter string, vdiskid string, xml string, host string, vms interface{}) error {
  doc := etree.NewDocument()
  err := doc.ReadFromString(xml)
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

  fmt.Println(uuid)

  vdisk := &Vms_vdisks{}
  errdb := dbs.Where("vdiskid=?", uuid).First(vdisk)
  if errdb.Error != nil {
    return 0, errdb.Error
  }
  return vdisk.Status, nil
}

func Getvdisk() ([]*Vms_vdisks, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  c := []*Vms_vdisks{}
  dbs.Find(&c)
  return c, err
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

func Disknametype(num int) string {
  switch (num) {
  case 0: return "vdb"
  case 1: return "vdc"
  case 2: return "vdd"
  case 3: return "vde"
  case 4: return "vdf"
  default:         return "UNKNOWN"
  }
}

func Mountdisk(ip string, vmhost string, storage string, pool string, datacenter string, vdiskid string, vms interface{}, xml string) error {
  storageinfo, err := ceph.Cephinfobyname(datacenter, storage)
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

  source.CreateAttr("name", pool+"/"+vdiskid)
  host := source.CreateElement("host")

  disknum := len(Getdiskbyvm(ip))
  var iplist []string
  iplist = strings.Split(storageinfo.Ips, ",")
  for _, v := range iplist {
    host.CreateAttr("name", v)
  }
  host.CreateAttr("port", storageinfo.Port)

  diskname := Disknametype(disknum)

  target := disk.CreateElement("target")
  target.CreateAttr("dev", diskname)
  target.CreateAttr("bus", "virtio")

  address := disk.CreateElement("address")
  address.CreateAttr("type", "pci")
  address.CreateAttr("domain", "0x0000")
  address.CreateAttr("bus", "0x00")
  slot := fmt.Sprintf("0x%d", 10+disknum)
  address.CreateAttr("slot", slot)
  address.CreateAttr("function", "0x0")
  doc.Indent(2)
  var docstring string
  docstring, err = doc.WriteToString()

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  err = libvirtd.DefineVm(docstring, vmhost)
  if err != nil {
    return err
  }

  errdb:= dbs.Model(vms).Where("ip=?", ip).Update("vmxml", docstring)
  if errdb.Error != nil {
    return vmerror.Error{Message: errdb.Error.Error()}
  }

  updatevm := UpdateMountvmstatus(datacenter, storage, vdiskid, ip, diskname)
  if updatevm != nil {
    return updatevm
  }
  return nil
}


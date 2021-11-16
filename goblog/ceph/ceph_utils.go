package ceph

import (
  "github.com/beevik/etree"
  "github.com/ceph/go-ceph/rados"
  rbd "github.com/ceph/go-ceph/rbd"
  "goblog/dbs"
  "goblog/libvirtd"
  "goblog/utils"
  "goblog/vmerror"
  "strings"
  "time"
)

type Vms_Ceph struct {
  Uuid string
  Pool string
  Datacenter string
  Ceph_secret string
  Ips string
  Port string
  Comment string
  Status int8
}

func Delete(uuid string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  dberr := dbs.Where("uuid=?", uuid).Delete(Vms_Ceph{})
  if dberr.Error != nil {
    return dberr.Error
  }

  return nil
}

func Add(name string, pool string, datacenter string, ceph_secret string, ips string, port string, comment  string) error {
  c := &Vms_Ceph{
    Uuid: name,
    Pool: pool,
    Datacenter: datacenter,
    Ceph_secret: ceph_secret,
    Ips: ips,
    Port: port,
    Comment: comment,
    Status: 1,
  }
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  errdb := dbs.Create(c)
  if errdb.Error != nil {
    return errdb.Error
  }
  dbs.NewRecord(c)
  return nil
}

func Get()([]*Vms_Ceph, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  c := []*Vms_Ceph{}
  dbs.Find(&c)
  return c, nil
}

func Getpool(datacenter string, storage string)([]*Vms_Ceph, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  c := []*Vms_Ceph{}
  dbs.Where("datacenter=? and uuid=?", datacenter, storage).Find(&c)
  return c, nil
}

func Cephinfobyname(datacenter string, storage string)([]*Vms_Ceph, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  c := []*Vms_Ceph{}
  dbs.Where("datacenter=? and uuid=?", datacenter, storage).Find(&c)
  return c, nil
}

func CephConn() (*rados.Conn, error) {
  conn, err := rados.NewConn()
  if err != nil {
    return nil, err
  }
  err = conn.ReadDefaultConfigFile()
  if err != nil {
    return nil, err
  }
  err = conn.Connect()
  if err != nil {
    return nil, err
  }

  return conn, nil
}

func Rm_image(uuid string) (error) {
  conn, err := CephConn()
  if err != nil {
    return err
  }

  ioctx, err := conn.OpenIOContext("vm")
  if err != nil {
    return err
  }

  img := rbd.GetImage(ioctx, uuid)
  Archive_img := "x_"+(time.Now().Format("200601021504"))+uuid
  img.Rename(Archive_img)

  return nil
}

func RbdClone(id string) (string, error) {
  conn, err := CephConn()
  if err != nil {
    return "", err
  }

  ioctx, _ := conn.OpenIOContext("vm")
  img := rbd.GetImage(ioctx, "0000_demo_centos7")
  _, e := img.Clone("20210806_095737", ioctx, id, rbd.RbdFeatureLayering, 0)

  if e != nil {
    return "", e
  }
  return id, nil
}

func Createcloudrive( uuid string, contain int) error {
  conn, err := CephConn()
  if err != nil {
    return err
  }

  ioctx, _ := conn.OpenIOContext("vm")
  _, err = rbd.Create(ioctx, uuid, uint64(1024*1024*1024*contain), 0)
  if err != nil {
    return err
  }
  return nil
}

type Vms_cloudrive struct {
  Cloudriveid string
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


func Getvdisk() ([]*Vms_cloudrive, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  c := []*Vms_cloudrive{}
  dbs.Find(&c)
  return c, err
}

func Getdiskbyvm(vmip string) ([]*Vms_cloudrive) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return []*Vms_cloudrive{}
  }
  c := []*Vms_cloudrive{}
  dbs.Select("contain, diskname").Where("vm_ip=?", vmip).Find(&c)
  return c
}

func UpdateMountvmstatus(datacenter string, storage string, cloudriveid string, vmip string, diskname string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  errdb := dbs.Model(Vms_cloudrive{}).Where("datacenter=? and storage=? and cloudriveid=?", datacenter, storage, cloudriveid).Update(map[string]interface{}{"vm_ip": vmip, "status": 0, "diskname": diskname})
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
  errdb := dbs.Model(Vms_cloudrive{}).Where("datacenter=? and storage=?", datacenter, storage).Update("vm_ip", "").Update("status", 1).Where("vm_ip=?", vmip)
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}

func Umountvmstatus(datacenter string, storage string, cloudriveid string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  errdb := dbs.Model(Vms_cloudrive{}).Where("datacenter=? and storage=? and cloudriveid=?", datacenter, storage, cloudriveid).Update("vm_ip", "").Update("status", 1)
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}

func Add_cloudrive(contain int, pool string, storage string, datacenter string, user string) ([]*Vms_cloudrive, error) {
  cloudriveid := utils.Createuuid()
  c := &Vms_cloudrive{
    Cloudriveid: cloudriveid,
    Contain: contain,
    Pool: pool,
    Storage: storage,
    Datacenter: datacenter,
    User: "niklaus",
    Exist: 1,
    Status: 1,
  }

  err := Createcloudrive(cloudriveid, contain)
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

func Deletevdisk(uuid string) error {
  checkmount, err := Getdiskstatus(uuid)
  if err != nil {
    return err
  }

  if checkmount == 0 {
    return vmerror.Error{Message: "硬盘已挂载，请卸载后删除"}
  }

  err = Rm_image(uuid)
  if err != nil {
    return err
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb := dbs.Where("cloudriveid=?", uuid).Delete(Vms_cloudrive{})
  if errdb.Error != nil {
    return vmerror.Error{Message: "delete vdisk fail"}
  }
  return nil
}

func Getdiskstatus(uuid string) (int, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return 0, err
  }

  vdisk := &Vms_cloudrive{}
  errdb := dbs.Where("cloudriveid=?", uuid).First(vdisk)
  if errdb.Error != nil {
    return 0, errdb.Error
  }
  return vdisk.Status, nil
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

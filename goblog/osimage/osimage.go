package osimage

import (
  "fmt"
  "goblog/ceph"
  db "goblog/dbs"
  "goblog/networks"

  "github.com/beevik/etree"

  "goblog/vmerror"
)


type Vms_os struct {
  Osname string
  Cephblockdevice string
  Snapimage string
  Xml string
  Status int8
}

func Del(osname string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  err1 := dbs.Model(Vms_os{}).Where("osname=?", osname).Update("status", 0)
  if err1.Error != nil {
    return err1.Error
  }
  return nil
}

func Add(osname string, cephblockdevice string, snapimage string, xml string) error {
  os := &Vms_os{
    Osname: osname,
    Cephblockdevice: cephblockdevice,
    Snapimage: snapimage,
    Xml: xml,
    Status: 1,
  }
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb := dbs.Create(*os)
  if errdb.Error != nil {
    return errdb.Error
  }
  booldb := dbs.NewRecord(*os)
  if booldb == false {
    return vmerror.Error{
      Message: "数据库错误",
    }
  }
  return nil
}

func Get() ([]*Vms_os, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var v []*Vms_os
  dbs.Find(&v)
  return v, nil
}

func getxml(osname string) (string, error) {
  db, err := db.NicloudDb()
  if err != nil {
    return "", err
  }

  var v []*Vms_os
  db.Where("osname=?", osname).Find(&v)
  return v[0].Xml, nil
}


func Xml(cephnaem string, vcpu int, vmem int, uuid string, mac string, image_name string, osname string) (string, error) {
  cephinfo, err := ceph.Get()
  if err != nil {
    return "", err
  }
  var ceph_secret = cephinfo[0].Ceph_secret
  ips := cephinfo[0].Ips
  port := cephinfo[0].Port

  br, err := networks.Getvlan()
  if err != nil {
    return "", err
  }

  f, err := getxml(osname)
  if err != nil {
    return "", err
  }

  doc := etree.NewDocument()
  err = doc.ReadFromString(f)
  if err != nil {
    return "", err
  }

  cephsecret := doc.FindElement("./domain/devices/disk/auth/secret")
  cephsecret.CreateAttr("uuid", ceph_secret)

  cpu := doc.FindElement("./domain/vcpu")
  cpu.CreateText(fmt.Sprintf("%d", vcpu))

  id := doc.FindElement("./domain/uuid")
  id.CreateText(uuid)

  name := doc.FindElement("./domain/name")
  name.CreateText(uuid)

  mem := doc.FindElement("./domain/memory")
  mem.CreateText(fmt.Sprintf("%d", vmem))

  currentMemory := doc.FindElement("./domain/currentMemory")
  currentMemory.CreateText(fmt.Sprintf("%d", vmem))

  bridge := doc.FindElement("./domain/devices/interface/source")
  bridge.CreateAttr("bridge", fmt.Sprintf("%s", br))

  macaddr := doc.FindElement("./domain/devices/interface/mac")
  macaddr.CreateAttr("address", fmt.Sprintf("%s", mac))

  for _, e := range doc.FindElements("./domain/devices[1]/*") {
    if e.Tag == "disk" {
      for _, v := range e.ChildElements() {
        if v.Tag == "source" {
          v.CreateAttr("name", fmt.Sprintf("vm/%s", image_name))

          for ip_k, ip := range ips {
            v.CreateElement("host")
            v.ChildElements()[ip_k].CreateAttr("name", string(ip))
            v.ChildElements()[ip_k].CreateAttr("port", port)
          }
        }
      }
    }
  }
  doc.Indent(2)
  var docstring string
  docstring, err = doc.WriteToString()
  if err != nil {
    return "", err
  }

  return docstring, nil
}


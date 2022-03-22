package osimage

import (
  "fmt"
  "github.com/beevik/etree"
  "nicloud/cephcommon"
  db "nicloud/dbs"
  "nicloud/networks"
  "nicloud/users"
  "reflect"
  "strings"
)


type Vms_os struct {
  Id int
  Sort int
  Owner int
  Size int `json:"Size" validate:"required"`
  Osname string `json:"Osname" validate:"required"`
  Datacenter string `json:"Datacenter" validate:"required"`
  Storage string  `json:"Storage" validate:"required"`
  Cephblockdevice string  `json:"Cephblockdevice" validate:"required"`
  Snapimage string
  Xml string  `json:"Xml" validate:"required"`
  Status int8
}

type Vms_osimage_sort struct {
  Id int
  Sort string
}

func get_osimage_sortbyid(id int) (*Vms_osimage_sort, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var o Vms_osimage_sort
  data := dbs.Where("id=?", id).First(&o)
  if data.Error != nil {
    return nil, data.Error
  }
  return &o, nil
}


func Get_osimage_sort() ([]*Vms_osimage_sort, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var o []*Vms_osimage_sort
  data := dbs.Find(&o)
  if data.Error != nil {
    return nil, data.Error
  }
  return o, nil
}

func Del(osid int) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  err1 := dbs.Where("id=?", osid).Delete(Vms_os{})
  if err1.Error != nil {
    return err1.Error
  }
  return nil
}

func Update(id int, datacenter string, storage string, osname string,  snapimage string, cephblockdevice string, xml string) error {
  os := &Vms_os{
    Id: id,
    Datacenter: datacenter,
    Storage: storage,
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

  errdb := dbs.Model(&Vms_os{}).Where("id=?", id).Update(os)
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}

func (vmsos *Vms_os) Add(datacenter string, storage string, osname string, cephblockdevice string, xml string, sort int, owner int, snap string, size int) error {
  os := Vms_os{
    Datacenter: datacenter,
    Storage: storage,
    Size: size,
    Osname: osname,
    Cephblockdevice: cephblockdevice,
    Snapimage: snap,
    Xml: xml,
    Status: 1,
    Sort: sort,
    Owner: owner,
  }
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb := dbs.Create(&os)
  if errdb.Error != nil {
    return errdb.Error
  }

  return nil
}

func Maposimage(user int, sort int) ([]map[string]interface{}, error)  {
  var mapc []map[string]interface{}
  var obj []Vms_os
  var err error
  if sort == 0 {
    obj, err = Get(user, sort)
    if err != nil  {
      return nil, err
    }
  } else {
    obj, err = Getimagebysort(user, sort)
    if err != nil  {
      return nil, err
    }
  }

  for _, v := range obj {
    c := make(map[string]interface{})
    m := reflect.TypeOf(v)
    n := reflect.ValueOf(v)
    for i := 0; i < m.NumField(); i++ {
      c[m.Field(i).Name] = n.Field(i).Interface()
    }

    sort, err := get_osimage_sortbyid(v.Sort)
    if err != nil {
      c["sort"] = nil
    } else {
      c["sort"] = sort.Sort
    }

    ceph := cephcommon.Vms_Ceph{}
    storageinfo, err := ceph.Cephinfobyuuid(v.Storage)
    if err != nil {
      c["storagename"] = nil
    } else {
      c["storagename"] = storageinfo.Name
    }

    owner, err := users.GetUserByUserID(v.Owner)
    if err != nil {
      c["owner"] = nil
    } else {
      c["owner"] = owner.Username
    }
    mapc = append(mapc, c)
  }
  return mapc, nil
}

func Get(user int, sort int) ([]Vms_os, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var v []Vms_os
  data := dbs.Find(&v)
  if data.Error != nil {
    return nil, data.Error
  }
  return v, nil
}

func Getimagebysort(userid int, sortid int) ([]Vms_os, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var v []Vms_os
  dbs.Where("owner=? and sort=?", userid, sortid).Find(&v)
  return v, nil
}

func Getimageby(datacenter string, storage string) ([]*Vms_os, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var v []*Vms_os
  dbs.Where("datacenter=? and storage=?", datacenter, storage).Find(&v)
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

func Xml(datacenter string, storage string, vlan string,  vcpu int, vmem int, uuid string, mac string, image_name string, osid int, pool string) (string, error) {
  ceph := cephcommon.Vms_Ceph{}
  storagename, err := ceph.Cephinfobyuuid(storage)
  if err != nil {
    return "", err
  }

  var ceph_secret = storagename.Ceph_secret
  ips := strings.Split(storagename.Ips, ",")
  port := storagename.Port

  br, err := networks.Getbridge(datacenter, vlan)
  if err != nil {
    return "", err
  }

  osinfo, err := GetOsInfoById(storage, osid)
  if err != nil {
    return "", err
  }

  doc := etree.NewDocument()
  err = doc.ReadFromString(osinfo.Xml)
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
          v.CreateAttr("name", fmt.Sprintf("%s/%s", pool, image_name))

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

func GetOsInfoById(storage string, id int) (*Vms_os, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  o := &Vms_os{}
  dbs.Where("id=? and storage=?", id, storage).First(o)
  return o, nil
}


func (o Vms_os)CheckOsbyUuid(uuid string) (bool, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return false, err
  }

  os := []*Vms_os{}
  dberr := dbs.Where("Cephblockdevice=?", uuid).First(&os)
  if dberr.Error != nil {
    return false, dberr.Error
  }
  if len(os) > 0 {
    return true, err
  }
  return false, nil
}

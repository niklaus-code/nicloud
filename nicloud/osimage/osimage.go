package osimage

import (
  "nicloud/cephcommon"
  db "nicloud/dbs"
  "nicloud/users"
  "reflect"
)


type Vms_os struct {
  Id int `gorm:"primary_key;AUTO_INCREMENT"`
  Sort int `json:"Sort" validate:"required"`
  Owner int
  Osname string `gorm:"unique" json:"Osname" validate:"required"`
  Datacenter string `json:"Datacenter" validate:"required"`
  Storage string  `json:"Storage" validate:"required"`
  Cephblockdevice string  `json:"Cephblockdevice" validate:"required"`
  Snapimage string
  Xml string  `gorm:"size:65535" json:"Xml" validate:"required"`
  Tag int `json:"Tag" validate:"required"`
  Status int8
}

type Vms_osimage_sort struct {
  Id int `gorm:"primary_key;AUTO_INCREMENT"`
  Sort string
}

type Vms_os_tags struct {
  Id int `gorm:"primary_key;AUTO_INCREMENT"`
  Tag string
}

func (t Vms_os_tags)Getostags() ([]*Vms_os_tags, error){
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  var tags []*Vms_os_tags
  err = dbs.Find(&tags).Error
  if err != nil {
    return nil, err
  }

  return tags, nil
}

func (t Vms_os_tags)GetostagByid(id int) (*Vms_os_tags, error){
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  err = dbs.Where("id=?", id).First(&t).Error
  if err != nil {
    return nil, err
  }

  return &t, nil
}

func get_osimage_sortbyid(id int) (*Vms_osimage_sort, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var o Vms_osimage_sort
  err = dbs.Where("id=?", id).First(&o).Error
  if err != nil {
    return nil, err
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

func (vmsos *Vms_os)Add (datacenter string, storage string, osname string, cephblockdevice string, xml string, sort int, owner int, snap string, tag int) error {
  os := Vms_os{
    Datacenter: datacenter,
    Storage: storage,
    Osname: osname,
    Cephblockdevice: cephblockdevice,
    Snapimage: snap,
    Xml: xml,
    Status: 1,
    Sort: sort,
    Owner: owner,
    Tag: tag,
  }
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  err = dbs.Create(&os).Error
  if err != nil {
    return err
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
      c["Sort"] = nil
    } else {
      c["Sort"] = sort
    }

    os_tag := Vms_os_tags{}
    tag, err := os_tag.GetostagByid(v.Tag)

    if err != nil {
      c["Tag"] = ""
    } else {
      c["Tag"] = tag
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
  err = dbs.Where("datacenter=? and storage=?", datacenter, storage).Find(&v).Error
  if err != nil {
    return nil, err
  }
  return v, nil
}

func (o Vms_os)GetOsInfoById(storage string, id int) (*Vms_os, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  err = dbs.Where("id=? and storage=?", id, storage).First(o).Error
  if err != nil {
    return nil, err
  }
  return &o, nil
}


func (o Vms_os)CheckOsbyUuid(uuid string) (bool, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return false, err
  }

  os := []*Vms_os{}
  err = dbs.Where("Cephblockdevice=?", uuid).First(&os).Error
  if err != nil {
    return false, err
  }
  if len(os) > 0 {
    return true, err
  }
  return false, nil
}

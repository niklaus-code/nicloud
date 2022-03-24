package datacenter

import (
  "nicloud/dbs"
)

type Vms_datacenter struct {
  Datacenter string `gorm:"unique"`
  Comment string
  Status int
}

func Get() ([]*Vms_datacenter, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var center []*Vms_datacenter
  dberr := dbs.Find(&center)
  if dberr.Error != nil {
    return nil, dberr.Error
  }
  return center, nil
}

func Add (datacenter string, comment string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  c := Vms_datacenter{
    Datacenter: datacenter,
    Comment: comment,
    Status: 1,
  }

  errdb := dbs.Create(&c)
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}

func Del(datacenter string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb := dbs.Where("datacenter=?", datacenter).Delete(&Vms_datacenter{})
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}

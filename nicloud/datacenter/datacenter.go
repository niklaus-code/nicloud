package datacenter

import "nicloud/dbs"

type Vms_datacenter struct {
  Datacenter string
  Comment string
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

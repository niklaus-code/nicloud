package networks

import (
	"goblog/dbs"
)

type Vms_vlans struct {
  Vlan string
  Bridge string
  Network string
  Prefix int
  Gateway string
  Status bool
}

func AddVlan(vlan string, bridge string, network string, prefix int, gateway string) error {
  v := &Vms_vlans{
    Vlan: vlan,
    Bridge: bridge,
    Network: network,
    Prefix: prefix,
    Gateway: gateway,
    Status: true,
  }
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  err1 := dbs.Create(*v)
  if err1.Error != nil {
    return err1.Error
  }
  return err1.Error
}

func Getvlan() ([]*Vms_vlans, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  if err != nil {
    return nil, err
  }
  var f []*Vms_vlans
  dbs.Find(&f)
  return f, nil
}

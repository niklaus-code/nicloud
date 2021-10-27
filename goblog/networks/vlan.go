package networks

import (
  "fmt"
  "goblog/dbs"
  "strconv"
  "strings"
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

func split(item string) (bool, []string) {
  start := item
  l := strings.Split(start, ".")

  if len(l) != 4 {
    return  false, nil
  }
  return true, l
}

func Createip(startip string, endip string) bool {
  b, l := split(startip)

  if b == false {
    return false
  }

  for _,v := range l {
    _, err := strconv.Atoi(v)
    if err != nil {
      return false
    }
  }


  c, d := split(endip)
  if c == false {
    return false
  }

  for _,v := range d {
    _, err := strconv.Atoi(v)
    if err != nil {
      return false
    }
  }

  startnum, _ := strconv.Atoi(l[3])
  endnum, _ := strconv.Atoi(d[3])
  for i:= startnum; i <= endnum ; i++ {
    fmt.Println(l[0]+"."+l[1]+"."+l[2]+"."+strconv.Itoa(i))
  }


  return true
}

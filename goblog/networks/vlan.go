package networks

import (
  "fmt"
  "goblog/dbs"
  "math/rand"
  "strconv"
  "strings"
  "time"
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


type Vms_ips struct {
  Ipv4 string
  Macaddr string
  Status int8
  Vlan string
}


func IPlist(vlan string) []*Vms_ips {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil
  }
  var ip []*Vms_ips
  dbs.Where("status=0 and vlan=?", vlan).Find(&ip)

  return ip
}

func Ipresource(ip string, mac string) bool {
  dbs, err := db.NicloudDb()
  if err != nil {
    return false
  }
  var ipnet []*Vms_ips
  dbs.Where("ipv4=?", ip).Where("macaddr=?", mac).Find(&ipnet)
  for _, v := range ipnet {
    if v.Status == 0 {
      return false
    }
  }
  return true
}

func Updateipstatus(ipv4 string) (error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  dbs.Model(&Vms_ips{}).Where("ipv4=?", ipv4).Update("status", 1)
  return nil
}


func Createip(startip string, endip string, vlan string) bool {
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

  if startnum > endnum {
    return false
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return false
  }
  for i:= startnum; i <= endnum ; i++ {
    i := &Vms_ips{
      Ipv4: l[0]+"."+l[1]+"."+l[2]+"."+strconv.Itoa(i),
      Macaddr: NewRandomMac().String(),
      Vlan: vlan,
      Status: 0,
    }
    err := dbs.Create(*i)
    if err.Error != nil {
      return false
    }
    dbs.NewRecord(*i)
  }
  return true
}


type Mac [6]byte

func (m Mac) String() string {
  return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",m[0],m[1],m[2],m[3],m[4],m[5])
}

func NewRandomMac() Mac{
  var m [6]byte

  rand.Seed(time.Now().UnixNano())
  for i:=0;i<6;i++ {
    mac_byte := rand.Intn(256)
    m[i] = byte(mac_byte)

    rand.Seed(int64(mac_byte))
  }

  return Mac(m)
}

package networks

import (
  "fmt"
  "goblog/dbs"
  vmerror "goblog/vmerror"
  "math/rand"
  "strconv"
  "strings"
  "time"
)

type Vms_vlans struct {
  Datacenter string
  Vlan string
  Bridge string
  Network string
  Prefix int
  Gateway string
  Status bool
}

func Delete(vlan string) error {
  existips := IPlist(vlan)
  if len(existips) > 0 {
    return vmerror.Error{Message: "存在vlan相关IP 无法删除"}
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  dberr := dbs.Where("vlan=?", vlan).Delete(&Vms_vlans{})
  if dberr.Error != nil {
    return dberr.Error
  }

  return nil
}

func AddVlan(datacenter string, vlan string, bridge string, network string, prefix int, gateway string) error {
  v := &Vms_vlans{
    Datacenter: datacenter,
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

func Getvlanbydatacenter(datacenter string) ([]*Vms_vlans, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  var f []*Vms_vlans
  dbs.Where("datacenter=?", datacenter).Find(&f)
  return f, nil
}

func Getbridge(datacenter string, vlan string) (string, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return "", err
  }
  var f []*Vms_vlans
  dbs.Where("datacenter=? and vlan=?", datacenter, vlan).Find(&f)
  return f[0].Bridge, nil
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
  Exist int8
}

func AllIP(vlan string) []*Vms_ips {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil
  }
  var ip []*Vms_ips
  dbs.Where("vlan=?", vlan).Find(&ip)

  return ip
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

func Ipresource(ip string) (string, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return "", err
  }
  var ipnet []*Vms_ips
  dbs.Where("ipv4=?", ip).Find(&ipnet)
  for _, v := range ipnet {
    if v.Status == 1 {
      return "", vmerror.Error{
        Message: "IP已经被占用",
      }
    }
  }
  return ipnet[0].Macaddr, nil
}

func Deleteip(ipv4 string, vlan string) error {
 dbs, err := db.NicloudDb()
 if err != nil {
   return err
 }

 dberr := dbs.Where("vlan=? and ipv4=?", vlan, ipv4 ).Delete(&Vms_ips{})
 if dberr.Error != nil {
   return dberr.Error
 }

 return nil
}

func Updateipstatus(ipv4 string, status int) (error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  dbs.Model(&Vms_ips{}).Where("ipv4=?", ipv4).Update("status", status)
  return nil
}

func OpIP(ipv4 string, vlan string, op int) error {
  dbs, err1 := db.NicloudDb()
  if err1 != nil {
    return vmerror.Error{
      Message: err1.Error(),
    }
  }

  dbs, err := db.NicloudDb()
  if err!= nil {
    return vmerror.Error{
      Message: err.Error(),
    }
  }
  err2 := dbs.Model(&Vms_ips{}).Where("ipv4=? and vlan=?", ipv4, vlan).Update("status", op)
  if err2.Error != nil {
    return vmerror.Error{
      Message: err2.Error.Error(),
    }
  }
  return nil
}

func Createip(startip string, endip string, vlan string) error {
  b, l := split(startip)

  if b == false {
    return vmerror.Error{
      Message: "数据格式错误",
    }
  }

  for _,v := range l {
    _, err := strconv.Atoi(v)
    if err != nil {
      return vmerror.Error{
        Message: "数据格式错误",
      }
    }
  }

  c, d := split(endip)
  if c == false {
    return vmerror.Error{
      Message: "数据格式错误",
    }
  }

  for _,v := range d {
    _, err := strconv.Atoi(v)
    if err != nil {
      return vmerror.Error{
        Message: "数据格式错误",
      }
    }
  }

  startnum, _ := strconv.Atoi(l[3])
  endnum, _ := strconv.Atoi(d[3])

  if startnum > endnum {
    return vmerror.Error{
      Message: "数据格式错误",
    }
  }

  dbs, err1 := db.NicloudDb()
  if err1 != nil {
    return vmerror.Error{
      Message: err1.Error(),
    }
  }

  for i:= startnum; i <= endnum ; i++ {
    ips := &Vms_ips{
      Ipv4: l[0]+"."+l[1]+"."+l[2]+"."+strconv.Itoa(i),
      Macaddr: NewRandomMac().String(vlan, i),
      Vlan: vlan,
      Status: 0,
      Exist: 1,
    }

    err := dbs.Create(*ips)
    if err.Error != nil {
      return err.Error
    }
    dbs.NewRecord(*ips)
  }
  return nil
}


type Mac [3]byte

func (m Mac) String(vlan string, end int) string {
  if len(vlan) == 1 {
    vlan = "0"+vlan
  }

  if len(vlan) > 2 {
    vlan = vlan[len(vlan)-2 : len(vlan)]
  }

  return fmt.Sprintf("c8:00:%02x:%02x:%s:%02x",m[0],m[1], vlan, end)
}

func NewRandomMac() Mac{
  var m [3]byte

  rand.Seed(time.Now().UnixNano())
  for i:=0;i<3;i++ {
    mac_byte := rand.Intn(256)
    m[i] = byte(mac_byte)

    rand.Seed(int64(mac_byte))
  }

  return Mac(m)
}

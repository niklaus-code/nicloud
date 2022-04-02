package networks

import (
  "fmt"
  "github.com/jinzhu/gorm"
  "nicloud/dbs"
  vmerror "nicloud/vmerror"
  "math/rand"
  "strconv"
  "strings"
  "time"
)

type Vms_vlans struct {
  Datacenter string  `gorm:"primary_key" json:"Datacenter" validate:"required"`
  Vlan string  `json:"Vlan" validate:"required"`
  Bridge string  `json:"Bridge" validate:"required"`
  Network string   `json:"Network" validate:"required"`
  Prefix int   `json:"Prefix" validate:"required"`
  Gateway string  `json:"Gateway" validate:"min=8,max=15"`
  Status bool
}

func DeleteVlan(vlan string) error {
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

  var f []*Vms_vlans
  dbs.Find(&f)
  return f, nil
}

func (h Vms_vlans)Gethostunselectedvlan (vlanlist []string) ([]*Vms_vlans, error) {
  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var v []*Vms_vlans
  db.Not("vlan", vlanlist).Find(&v)
  if len(vlanlist) == 0 {
    db.Find(&v)
  }
  return v, nil
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
  Ipv4 string `gorm:"primary_key;"`
  Macaddr string `gorm:"unique"`
  Status int8
  Vlan string
  Exist int8
  Create_time time.Time
}

func AllIP(vlan string) ([]*Vms_ips, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var ip []*Vms_ips
  errdb := dbs.Where("vlan=?", vlan).Order("length(ipv4)").Order("ipv4").Find(&ip)
  if errdb.Error != nil {
    return nil, vmerror.Error{Message: errdb.Error.Error()}
  }
  return ip, nil
}

func Downloadips(vlan string) (string, error) {
  ips, err := AllIP(vlan)
  if err != nil {
    return "", err
  }
  var ipliststr string
  str1 := "host v_"
  for _, v := range ips {
      ipliststr += str1
      ipliststr += v.Ipv4
      ipliststr += "{hardware ethernet "
      ipliststr += v.Macaddr
      ipliststr += ";fixed-address "
      ipliststr += v.Ipv4
      ipliststr += ";}"
      ipliststr += "\n"
  }
  return ipliststr, nil
}

func IPlist(vlan string) []*Vms_ips {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil
  }
  var ips []*Vms_ips
  dbs.Where("vlan=? and status=0", vlan).Order("length(ipv4)").Order("ipv4").Find(&ips)

  return ips
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

func Updateipstatus(ipv4 string, status int) (*gorm.DB, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  tx := dbs.Begin()
  err = tx.Model(&Vms_ips{}).Where("ipv4=?", ipv4).Update("status", status).Error
  if err != nil {
    tx.Rollback()
    return nil, err
  }
  return tx, nil
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

func Createip(startip string, endip string, vlan string, prefix int, gateway string) error {
  b, l := split(startip)
  e, f := split(gateway)

  if b == false || e == false {
    return vmerror.Error{
      Message: "数据格式错误",
    }
  }

  if prefix >= 8 && prefix < 16 {
    if l[0] != f[0] {
      return vmerror.Error{
        Message: "数据格式错误",
      }
    }
  }

  if prefix >= 16 && prefix < 24 {
    if l[0] != f[0] || l[1] != f[1] {
      return vmerror.Error{
        Message: "数据格式错误",
      }
    }
  }

  if prefix >= 24 && prefix <= 32 {
    if l[0] != f[0] || l[1] != f[1] || l[2] != f[2] {
      return vmerror.Error{
        Message: "数据格式错误",
      }
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
      Macaddr: NewRandomMac().String(l[2], i),
      Vlan: vlan,
      Status: 0,
      Exist: 1,
      Create_time: time.Now(),
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
  vlani, _ := strconv.Atoi(vlan)
  buf := make([]byte, 255)
  //fmt.Println(fmt.Sprintf("%02X, %02x", vlan, end))
  return fmt.Sprintf("c8:00:%02x:%02x:%02x:%02x",m[0],m[1], buf[vlani], end)
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

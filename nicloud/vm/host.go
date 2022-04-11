package vm

import (
  "fmt"
  "github.com/jinzhu/gorm"
  db "nicloud/dbs"
  "nicloud/libvirtd"
  "nicloud/networks"
  "nicloud/vmerror"
  "reflect"
)

type Vm_hosts struct {
  Ipv4        string `gorm:"primary_key" json:"ipv4" validate:"min=8,max=15"`
  Mem         uint `json:"mem" validate:"gt=0"`
  Cpu         uint `json:"cpu" validate:"gt=0"`
  Max_vms     uint `json:"max_vms" validate:"gt=0"`
  Created_vms uint
  Usedmem     uint
  Usedcpu     uint
  Datacenter  string `json:"datacenter" validate:"required"`
  Status      int8
  Comment     string
}

type Vms_vlan_map_hosts struct {
  id int `gorm:"primary_key;AUTO_INCREMENT"`
  Vlan string
  Hosts string
}

func (h Vm_hosts)gethostunselectedvlan (ip string) ([]*networks.Vms_vlans, error) {
  vlan := networks.Vms_vlans{}
  hostmapvlan := Vms_vlan_map_hosts{}

  hostvlans, err := hostmapvlan.Getvlanbyhost(ip)
  if err != nil {
    return nil, err
  }

  var vlanlist []string
  for _, v := range hostvlans {
    vlanlist = append(vlanlist, v.Vlan)
  }

  unseletdvlan, err := vlan.Gethostunselectedvlan(vlanlist)
  if err != nil {
    return nil, err
  }

  return unseletdvlan, nil
}

func Allhosts(hosts []Vm_hosts) []map[string]interface{}  {
  mapvlanhost := Vms_vlan_map_hosts{}
  var mapc []map[string]interface{}

  for _, v := range hosts {
    c := make(map[string]interface{})
    c["count"] = CountHosts(v.Ipv4)

    m := reflect.TypeOf(v)
    n := reflect.ValueOf(v)
    for i := 0; i < m.NumField(); i++ {
      c[m.Field(i).Name] = n.Field(i).Interface()
    }
    c["vmnum"] = ""
    c["vlan"], _ = mapvlanhost.Getvlanbyhost(v.Ipv4)
    mapc = append(mapc, c)
  }
  return mapc
}

func Maphost(ip string) (map[string]interface{}, error) {
  h := Vm_hosts{}
  host, err := h.Gethostsbyip(ip)
  if err != nil {
    return nil, err
  }
  mapvlanhost := Vms_vlan_map_hosts{}
  c := make(map[string]interface{})
  c["count"] = CountHosts(host.Ipv4)

  m := reflect.TypeOf(*host)
  n := reflect.ValueOf(*host)
  c["vmnum"] = ""
  for i := 0; i < m.NumField(); i++ {
    c[m.Field(i).Name] = n.Field(i).Interface()
  }
  c["vlan"], _ = mapvlanhost.Getvlanbyhost(host.Ipv4)
  c["unselectvlan"], _ = h.gethostunselectedvlan(ip)
  return c, nil
}

func CountHosts(ip string) int {
  db, err := db.NicloudDb()
  if err != nil {
    return 0
  }
  var c int
  db.Model(&Vms{}).Where("host=?", ip).Count(&c)
  return c
}

func (h Vm_hosts)Createhost(datacenter string, cpu uint, mem uint, ip string, num uint, vlan []string) error {
  db, err := db.NicloudDb()
  if err != nil {
   return err
  }
  host := &Vm_hosts{
   Datacenter: datacenter,
   Cpu: cpu,
   Mem: mem,
   Ipv4: ip,
   Max_vms: num,
   Status: 1,
   Usedcpu: 0,
   Usedmem: 0,
  }

  res := db.Create(&host)
  if res.Error != nil {
   return vmerror.Error{Message: "数据插入失败"}
  }

  vlanmaphost := Vms_vlan_map_hosts{}
  for _, v := range vlan {
    err = vlanmaphost.Add(v, ip)
    if err != nil {
      return err
    }
  }
  return nil
}

func (vh Vms_vlan_map_hosts)Add (vlan string, host string) error {
  db, err := db.NicloudDb()
  if err != nil {
    return err
  }
  s := &Vms_vlan_map_hosts{
    Vlan: vlan,
    Hosts: host,
  }

  res := db.Create(&s)
  if res.Error != nil {
    return vmerror.Error{Message: "数据插入失败"}
  }
  return nil
}

func (vh Vms_vlan_map_hosts)Getvlanbyhost (host string) ([]*Vms_vlan_map_hosts, error) {
  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  v := []*Vms_vlan_map_hosts{}
  dberr := db.Where("hosts = ?", host).Find(&v)
  if dberr.Error != nil {
    return nil, dberr.Error
  }
  return v, err
}

func (vh Vms_vlan_map_hosts)Gethostbyvlan (vlan string) ([]*Vms_vlan_map_hosts, error) {
  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  v := []*Vms_vlan_map_hosts{}
  dberr := db.Where("vlan = ?", vlan).Find(&v)
  if dberr.Error != nil {
    return nil, dberr.Error
  }
  return v, err
}

func (vh Vms_vlan_map_hosts)Delhostmapvlan (hostip string) (error) {
  db, err := db.NicloudDb()
  if err != nil {
    return err
  }
  v := []*Vms_vlan_map_hosts{}
  dberr := db.Where("hosts = ?", hostip).Delete(&v)
  if dberr.Error != nil {
    return dberr.Error
  }
  return err
}

func (h Vm_hosts)checkcpumem(ip string, cpu uint, mem uint) error {
  db, err := db.NicloudDb()
  if err != nil {
    return err
  }
  v := &Vm_hosts{}
  db.Where("ipv4 = ?", ip).Find(&v)

  if cpu + v.Usedcpu > v.Cpu {
    return vmerror.Error{
      Message: "cpu资源不够",
    }
  }
  if mem + v.Usedmem > v.Mem {
    return vmerror.Error{
      "内存资源不够",
    }
  }

  return nil
}

func GetHostVmNumber(ip string) (uint, uint, error) {
  db, err := db.NicloudDb()
  if err != nil {
    return 0,0, err
  }
  h := &Vm_hosts{}
  errdb := db.Where("ipv4=?", ip).First(h)
  if errdb.Error != nil {
    return 0, 0, errdb.Error
  }
  return h.Created_vms, h.Max_vms, err
}

func (h Vm_hosts)UpdateCpuMem (ip string, cpu int, mem int) (*gorm.DB, error) {
  c, m, err := h.UpdateCpuMem_(ip, cpu, mem)
  if err != nil {
    return nil, err
  }

  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  tx := db.Begin()
  err = tx.Model(&Vm_hosts{}).Where("ipv4=?", ip).Update("usedcpu", c).Update("usedmem", m).Error
  if err!= nil {
    tx.Rollback()
    return nil, err
  }
  return tx, nil
}

func (h Vm_hosts)Createvmonhost(ip string, cpu uint, mem uint) (*gorm.DB, error) {
  c, m, err := h.UpdateCpuMem_(ip, int(cpu), int(mem))
  if err != nil {
    return nil, err
  }

  vms_num, max_num, err := GetHostVmNumber(ip)
  if err != nil {
    return nil, err
  }

  if vms_num + 1 > max_num {
    return nil, vmerror.Error{Message: "超过宿主机可以创建的最大数量"}
  }

  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  tx := db.Begin()
  err = tx.Model(&Vm_hosts{}).Where("ipv4=?", ip).Update("usedcpu", c).Update("usedmem", m).Update(" Created_vms", vms_num + 1).Error
  if err != nil {
    tx.Rollback()
    return nil, err
  }

  return tx, nil
}

func (h Vm_hosts)UpdateCpuMem_ (ip string, cpu int, mem int) (uint, uint, error) {
  host, err := h.Gethostsbyip(ip)
  if err != nil {
    return 0, 0, err
  }
  c := int(host.Usedcpu) + cpu
  m := int(host.Usedmem) + mem

  if c < 0 {
    c =0
  }
  if m < 0 {
    m= 0
  }

  return uint(c), uint(m), nil
}

func (h Vm_hosts)Deletehost (ip string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  dberr := dbs.Where("ipv4=?", ip).Delete(&Vm_hosts{})
  if dberr.Error != nil {
    return dberr.Error
  }

  vlanmaphost := Vms_vlan_map_hosts{}
  delvlanmaphost := vlanmaphost.Delhostmapvlan(ip)
  if delvlanmaphost != nil {
    return delvlanmaphost
  }

  return nil
}

func Gethostinfo(ip string) []map[string]interface{} {
  db, err := db.NicloudDb()
  if err != nil {
    return nil
  }
  var v []Vm_hosts
  db.Where("status=1 and ipv4 != ?", ip).Find(&v)
  res := Allhosts(v)
  return res
}

func Hosts() ([]map[string]interface{}, error) {
    db, err := db.NicloudDb()
    if err != nil {
        return nil, err
      }
    var hosts []Vm_hosts
    db.Where("status=1").Find(&hosts)

    res := Allhosts(hosts)
    return res, nil
  }

func (h Vm_hosts)GetHostsbyVlan(datacenter string, vlan string) ([]map[string]interface{},  error) {
  vlanmaphost := Vms_vlan_map_hosts{}
  vlanhosts, err  := vlanmaphost.Gethostbyvlan(vlan)
  if err != nil {
    return nil, err
  }
  var hostlist []string
  for _, v := range vlanhosts {
    hostlist = append(hostlist, v.Hosts)
  }

  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var hosts []Vm_hosts
  db.Where("ipv4 IN (?)", hostlist).Find(&hosts)

  res := Allhosts(hosts)
  return res, nil
}

func (h Vm_hosts)Gethostsbyip (ip string) (*Vm_hosts,  error) {
  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var host Vm_hosts
  db.Where("status=1 and ipv4=?", ip).First(&host)

  return &host, nil
}

func Addcomment(ip string, c string) error {
  db, err := db.NicloudDb()
  if err != nil {
    return err
  }

  dberr := db.Model(&Vm_hosts{}).Where("ipv4 = ?", ip).Update("comment", c)
  if dberr.Error != nil {
    return dberr.Error
  }
  return nil
}

func ListDomains(host string) (int, error) {
  l, err := libvirtd.Listdomains(host)
  return len(l), err
}

type counthosts struct {
  Mem int
  Cpu int
  Usedmem int
  Usedcpu int
  Counthosts int
  Cpu_percent string
  Mem_percent string
  Datacenter string
}

func CountHost() (*counthosts, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  var c counthosts
  sql := "select sum(mem) as mem, sum(cpu) as cpu, sum(usedcpu) as usedcpu, sum(usedmem) usedmem, count(ipv4) as counthosts, datacenter from vm_hosts;"
  dberr := dbs.Raw(sql).Scan(&c)
  if dberr.Error != nil {
    return nil, dberr.Error
  }
  c.Cpu_percent = fmt.Sprintf("%.2f", float64(c.Usedcpu)/float64(c.Cpu)*100)
  c.Mem_percent = fmt.Sprintf("%.2f", float64(c.Usedmem)/float64(c.Mem)*100)

  return &c, nil
}

func (h Vm_hosts)Updatehostinfo(ip string, cpu int, mem int, maxnum int, vlanlist []string) error {
  db, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb := db.Model(&Vm_hosts{}).Where("ipv4=?", ip).Update("cpu", cpu).Update("mem", mem).Update("Max_vms", maxnum)
  if errdb.Error != nil {
    return errdb.Error
  }

  hostmapvlan := Vms_vlan_map_hosts{}
  delhostmapvlan := hostmapvlan.Delhostmapvlan(ip)
  if delhostmapvlan != nil {
    return delhostmapvlan
  }

  for _, v := range vlanlist {
    err := hostmapvlan.Add(v, ip)
    if err != nil {
      return err
    }
  }
  return nil
}

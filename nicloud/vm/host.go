package vm

import (
  "fmt"
  db "nicloud/dbs"
  "nicloud/libvirtd"
  "nicloud/vmerror"
  "reflect"
)

type Vm_hosts struct {
  Datacenter  string  `json:"datacenter" validate:"required"`
  Ipv4        string  `json:"ipv4" validate:"min=8,max=15"`
  Mem         int `json:"mem" validate:"gt=0"`
  Cpu         int `json:"cpu" validate:"gt=0"`
  Max_vms     int `json:"max_vms" validate:"gt=0"`
  Created_vms int
  Usedmem     int
  Usedcpu     int
  Status      int8
  Comment     string
}

type Vms_vlan_map_hosts struct {
  id int
  Vlan string
  Hosts string
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

func CountHosts(ip string) int {
  db, err := db.NicloudDb()
  if err != nil {
    return 0
  }
  var c int
  db.Model(&Vms{}).Where("host=?", ip).Count(&c)
  return c
}

func (h Vm_hosts)Createhost(datacenter string, cpu int, mem int, ip string, num int, vlan []string) error {
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

func (h Vm_hosts)checkcpumem(ip string, cpu int, mem int) error {
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

func (h Vm_hosts)downcpumem (ip string, cpu int, mem int) (int, int, error) {
  host, err := h.gethostsbyip(ip)
  if err != nil {
    return 0, 0, err
  }

  c := host.Usedcpu - cpu
  m := host.Usedmem - mem

  if c < 0 {
    c = 0
  }

  if m < 0 {
    m = 0
  }
  return c, m, nil
}

func (h Vm_hosts)freecpumem (ip string, cpu int, mem int) error {
  c, m, err := h.downcpumem(ip, cpu, mem)
  db, err := db.NicloudDb()
  if err != nil {
    return err
  }

  var v Vm_hosts
  db.Where("ipv4 = ?", ip).Find(&v)

  db.Model(&Vm_hosts{}).Where("ipv4=?", v.Ipv4).Update("usedcpu", c)
  db.Model(&Vm_hosts{}).Where("ipv4=?", v.Ipv4).Update("usedmem", m)
  return nil
}

func GetHostVmNumber(ip string) (int, int, error) {
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

func (h Vm_hosts)Updatehostcpumem (ip string, cpu int, mem int) error {
  c, m, err := h.Addcpumem(ip, cpu, mem)
  if err != nil {
    return err
  }

  db, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb := db.Model(&Vm_hosts{}).Where("ipv4=?", ip).Update("usedcpu", c).Update("usedmem", m)
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}

func (h Vm_hosts)Updatehost(ip string, cpu int, mem int) error {
  c, m, err := h.Addcpumem(ip, cpu, mem)
  if err != nil {
    return err
  }

  vms_num, max_num, err := GetHostVmNumber(ip)
  if err != nil {
    return err
  }

  if vms_num + 1 > max_num {
    return vmerror.Error{Message: "超过宿主机可以创建的最大数量"}
  }

  db, err := db.NicloudDb()
  if err != nil {
    return err
  }

  err1 := db.Model(&Vm_hosts{}).Where("ipv4=?", ip).Update("usedcpu", c).Update("usedmem", m).Update(" Created_vms", vms_num + 1)
  if err1.Error != nil {
    return err1.Error
  }

  return nil
}

func (h Vm_hosts)Addcpumem (ip string, cpu int, mem int) (int, int, error) {
  host, err := h.gethostsbyip(ip)
  if err != nil {
    return 0, 0, err
  }
  c := host.Usedcpu + cpu
  m := host.Usedmem + mem

  return c, m, nil
}

func (h Vm_hosts)Deletehost(ip string) error {
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

func (h Vm_hosts)gethostsbyip (ip string) (*Vm_hosts,  error) {
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

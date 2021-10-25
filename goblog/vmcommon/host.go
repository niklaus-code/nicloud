package vmcommon

import (
  db "goblog/dbs"
  "reflect"
)

type Vm_hosts struct {
  Ipv4        string
  Mem         int
  Cpu         int
  Max_vms     int
  Created_vms int
  Usedmem     int
  Usedcpu     int
  Status      int8
}

func Allhosts(obj []Vm_hosts) []map[string]interface{}  {
  var mapc []map[string]interface{}

  for _, v := range obj {
    c := make(map[string]interface{})
    c["count"] = CountHosts(v.Ipv4)

    m := reflect.TypeOf(v)
    n := reflect.ValueOf(v)
    for i := 0; i < m.NumField(); i++ {
      c[m.Field(i).Name] = n.Field(i).Interface()
    }
    mapc = append(mapc, c)
  }
  return mapc
}

func Hosts() []map[string]interface{} {
  db, err := db.NicloudDb()
  if err != nil {
    return nil
  }
  var hosts []Vm_hosts
  db.Where("status=1").Find(&hosts)

  res := Allhosts(hosts)
  return res
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



func Createhost(cpu int, mem int, ip string, num int) bool {
  db, err := db.NicloudDb()
  if err != nil {
    return false
  }
  h := &Vm_hosts{
    Cpu: cpu,
    Mem: mem,
    Ipv4: ip,
    Max_vms: num,
    Status: 1,
    Usedcpu: 0,
    Usedmem: 0,
  }

  err1 := db.Create(*h)
  if err1.Error != nil {
    return false
  }

  //return bool
  res := db.NewRecord(&h)
  return res
}

func getcpumem(ip string, cpu int, mem int) map[string]int {
  db, err := db.NicloudDb()
  if err != nil {
    return nil
  }
  v := &Vm_hosts{}
  db.Where("ipv4 = ?", ip).Find(&v)
  c := make(map[string]int)

  if cpu + v.Usedcpu > v.Cpu {
    return nil
  }
  if mem + v.Usedmem > v.Mem {
    return nil
  }
  c["cpu"] = cpu + v.Usedcpu
  c["mem"] = mem +v.Usedmem
  return c
}

func Updatehost(ip string, cpu int, mem int) bool {
  db, err := db.NicloudDb()
  if err != nil {
    return false
  }

  t := getcpumem(ip, cpu, mem)
  if t == nil {
    return false
  }
  c := t["cpu"]
  m := t["mem"]

  err1 := db.Model(&Vm_hosts{}).Where("ipv4=?", ip).Update("usedcpu", c).Update("usedmem", m)
  if err1.Error != nil {
    return false
  }
  return true
}

func Delhost(ip string) bool {
  db, err := db.NicloudDb()
  if err != nil {
    return false
  }
  err1 := db.Model(&Vm_hosts{}).Where("ipv4=?", ip).Update("status", 0)
  if err1.Error != nil {
    return false
  }
  return true
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

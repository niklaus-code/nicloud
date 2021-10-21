package vmcommon

import db "goblog/dbs"

type Vm_hosts struct {
  Ipv4        string
  Mem         int
  Cpu         int
  Max_vms     int
  Created_vms int
  Status      int8
}

func Hosts() []*Vm_hosts {
  db, err := db.NicloudDb()
  if err != nil {
    return nil
  }
  var hosts []*Vm_hosts
  db.Where("status=1").Find(&hosts)
  return hosts
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
  }

  err1 := db.Create(*h)
  if err1.Error != nil {
    return false
  }

  //return bool
  res := db.NewRecord(&h)
  return res
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

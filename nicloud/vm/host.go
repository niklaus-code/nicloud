package vm

import (
  db "nicloud/dbs"
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
  Vlan        string  `json:"vlan" validate:"required"`
  Comment     string
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


func CountHosts(ip string) int {
  db, err := db.NicloudDb()
  if err != nil {
    return 0
  }
  var c int
  db.Model(&Vms{}).Where("host=?", ip).Count(&c)
  return c
}

func Createhost(datacenter string, cpu int, mem int, ip string, num int, vlan string) error {
  db, err := db.NicloudDb()
  if err != nil {
    return err
  }
  h := &Vm_hosts{
    Datacenter: datacenter,
    Cpu: cpu,
    Mem: mem,
    Ipv4: ip,
    Max_vms: num,
    Status: 1,
    Usedcpu: 0,
    Usedmem: 0,
    Vlan: vlan,
  }

  err1 := db.Create(*h)
  if err1.Error != nil {
    return err1.Error
  }

  //return bool
  res := db.NewRecord(&h)
  if res == false {
    return vmerror.Error{Message: "数据插入失败"}
  }
  return nil
}

func getcpumem(ip string, cpu int, mem int) (map[string]int, error) {
  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  v := &Vm_hosts{}
  db.Where("ipv4 = ?", ip).Find(&v)
  c := make(map[string]int)

  if cpu + v.Usedcpu > v.Cpu {
    return nil, vmerror.Error{
      Message: "cpu资源不够",
    }
  }
  if mem + v.Usedmem > v.Mem {
    return nil, vmerror.Error{
      "内存资源不够",
    }
  }
  c["cpu"] = cpu + v.Usedcpu
  c["mem"] = mem +v.Usedmem

  return c, nil
}

func Freehost(ip string, cpu int, mem int) error {
  db, err := db.NicloudDb()
  if err != nil {
    return err
  }
  var v Vm_hosts
  db.Where("ipv4 = ?", ip).Find(&v)

  db.Model(&Vm_hosts{}).Where("ipv4=?", v.Ipv4).Update("usedcpu", v.Usedcpu-cpu)
  db.Model(&Vm_hosts{}).Where("ipv4=?", v.Ipv4).Update("usedmem", v.Usedmem-mem)
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

func Updatehostcpumem(ip string, cpu int, mem int) error {
  c, m, err := checkcpumem(ip, cpu, mem)
  if err != nil {
    return err
  }
  db, err := db.NicloudDb()
  if err != nil {
    return err
  }

  err1 := db.Model(&Vm_hosts{}).Where("ipv4=?", ip).Update("usedcpu", c).Update("usedmem", m)
  if err1.Error != nil {
    return err1.Error
  }
  return nil
}

func Updatehost(ip string, cpu int, mem int) error {
  c, m, err := checkcpumem(ip, cpu, mem)
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

func checkcpumem(ip string, cpu int, mem int) (int, int, error) {
  t, err := getcpumem(ip, cpu, mem)
  if err != nil {
    return 0, 0, err
  }

  c := t["cpu"]
  m := t["mem"]

  return c, m, nil
}

func Deletehost(ip string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  dberr := dbs.Where("ipv4=?", ip).Delete(&Vm_hosts{})
  if dberr.Error != nil {
    return dberr.Error
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

func GetHostsbydatacenter(datacenter string, vlan string) ([]map[string]interface{},  error) {
  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var hosts []Vm_hosts
  db.Where("status=1 and datacenter=? and vlan=?", datacenter, vlan).Find(&hosts)

  res := Allhosts(hosts)
  return res, nil
}

func GetHostsbyvmip(vmip string) (*Vms,  error) {
  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  vm := &Vms{}
  errdb := db.Where("ip=?", vmip).Find(vm)
  if errdb.Error != nil {
    return nil, errdb.Error
  }

  return vm, nil
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

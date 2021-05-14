package vmcommon

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
  uuid "github.com/satori/go.uuid"
  libvirt "libvirt.org/libvirt-go"
  "time"
  "errors"
)

type Vms struct {
  Uuid       string
  Name       string
  Cpu        int
  Mem        int
  Create_time time.Time
  Owner      string
  Comment    string
  Vmxml      string
  Status     interface{}
  Exist       int
  Ip          string
  Host        string
}


func (v Vms) Error(info string) error {
  errorinfo := fmt.Sprintf("%s", info)
  return errors.New(errorinfo)
}

func vmdb() *gorm.DB {
  db, errDb := gorm.Open("mysql", "modis:modis@(127.0.0.1:3306)/gocloud?parseTime=true")
  if errDb != nil {
    fmt.Println(errDb)
  }
  return db
}


type Vm_xmls struct {
  Ostype string
  Osxml string
}

func libvirtconn(host string) (*libvirt.Connect, error) {
  conn, err := libvirt.NewConnect(fmt.Sprintf("qemu+ssh://%s/system", host))
  if err != nil {
    fmt.Println(err)
  }
  return conn, err
}

func VmStatus(uuid string, host string) (string, error) {
  conn, err := libvirtconn(host)
  if err != nil {
    return "", err
  }
  vm, err := conn.LookupDomainByUUIDString(uuid)
  if err != nil {
    return "vm not found", err
  }

  state, _ , err1  := vm.GetState()

  if err1 != nil {
    return "vm not found", err1
  }

  return Vmstate[state], err1
}

var Vmstate = map[libvirt.DomainState]string{
  1: "运行",
  5: "关机",
  2: "deleted",
}

type Error struct {
  Code    int16
  Message string
}

func (err Error) Error() string {
  return fmt.Sprintf("vm is running, con't delete")
}

func Delete(uuid string, ip string, host string) ([]*Vms, error) {
  vmstat, err := VmStatus(uuid, host)
  if err != nil {
    return nil, err
  }

  if vmstat == "开机" {
    //e.Message := fmt.Sprintf("domain is runnin, con't delete")
    return nil, Error{
      Code: 501,
      Message: "vm is running, con't delete",
    }
  }

  db := vmdb()

  //undefine vm
  conn, err := libvirtconn(host)
  if err !=nil {
    return nil, err
  }

  vm, err1 := conn.LookupDomainByUUIDString(uuid)
  if err1 != nil {
    return nil, err1
  }
  vm.Undefine()

  db.Model(&Vms{}).Where("uuid=?", uuid).Update("exist", 0)
  db.Model(&vm_networks{}).Where("ipv4=?", ip ).Update("status", 0)
  vmlist := VmList(host)
  return vmlist,err
}

func Shutdown(uuid string, host string) (*Vms, error) {
  /*start vm*/
  conn, err := libvirtconn(host)
  if err != nil {
    return nil, err
  }
  vm, err4 := conn.LookupDomainByUUIDString(uuid)
  if err4 != nil {
    return nil, err4
  }
  err1 := vm.Destroy()
  if err1 != nil {
    return nil, err1
  }

  db := vmdb()
  var v = &Vms{}
  db.Where("uuid = ?", uuid).First(&v)

  s, err2 := VmStatus(uuid, host)
  v.Status = s
  if err2 != nil {
    fmt.Println(err2)
  }
  return v, err2
}

func Start(uuid string, host string) (*Vms, error) {
  /*start vm*/
  conn, connerr := libvirtconn(host)
  if connerr != nil {
    return nil, connerr
  }
  vm, err := conn.LookupDomainByUUIDString(uuid)

  if err != nil {
    return nil, err
  }

  err1 := vm.Create()
  if err1 != nil {
    return nil, err1
  }

  db := vmdb()
  var v = &Vms{}
  db.Where("uuid = ?", uuid).First(&v)

  s, err2 := VmStatus(uuid, host)
  v.Status = s
  if err2 != nil {
    return nil, err2
  }
  return v, err2
}

func Createuuid() string {
  u := uuid.NewV4().String()
  return u
}

func savevm(uuid string, cpu int, mem int, vmxml string, ip string, host string) bool {
  db := vmdb()
  vm := &Vms{
    Uuid: uuid,
    Name: uuid,
    Cpu: cpu,
    Mem: mem,
    Vmxml: vmxml,
    Create_time: time.Now(),
    Status: 1,
    Exist: 1,
    Ip: ip,
    Host: host,
  }
  db.Create(vm)

  //return bool
  res := db.NewRecord(&vm)
  return res
}

func Create(cpu int, mem int, ip string, host string) (bool, error) {
  /*create a vm*/

  vcpu := cpu
  vmem := mem*1024*1024

  u := Createuuid()

  db := vmdb()
  var x Vm_xmls
  db.First(&x, "ostype = ?", "linux")


  vmxml := fmt.Sprintf(x.Osxml, u, u, vmem, vmem, vcpu)
  err := savevm(u, vcpu, vmem, vmxml, ip, host)

  dba := vmdb()
  dba.Model(&vm_networks{}).Where("ipv4=?", ip).Update("status", 1)

  if err == false {
    return false, nil
  }

  conn, connerr := libvirtconn(host)
  if connerr != nil {
    return false, connerr
  }
  _, err1 := conn.DomainDefineXML(vmxml)

  if err1 != nil {
    return false, err1
  }
  return true, err1
}

func VmList(host string) []*Vms {
  db := vmdb()
  var v []*Vms
  db.Where("exist=1").Find(&v)
  for _, e := range(v) {
    s, _ := VmStatus(e.Uuid, host)
    e.Status = s
  }
  return v
}

type vm_networks struct {
  Ipv4 string
  Macaddr string
  Status int8
}

func IPlist() []*vm_networks {
  db := vmdb()
  var ip []*vm_networks
  db.Where("status=0").Find(&ip)

  return ip
}

type Vm_hosts struct {
  Ipv4 string
  Mem int8
  Cpu int8
  Max_vms int8
  Created_vms int8
  Status int8
}

func Hosts() []*Vm_hosts {
  db := vmdb()
  var hosts []*Vm_hosts
  db.Where("status=0").Find(&hosts)
  return hosts
}

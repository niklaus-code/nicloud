package vmcommon

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
  uuid "github.com/satori/go.uuid"
  libvirt "libvirt.org/libvirt-go"
  "time"
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

func libvirtconn() (*libvirt.Connect, error) {
  conn, err := libvirt.NewConnect("qemu:///system")
  if err != nil {
    fmt.Println(err)
  }
  return conn, err
}

func VmStatus(uuid string) (string, error) {
  //var stats map[libvirt.DomainState]string
  //stats = make(map[libvirt.DomainState]string)
  //stats[0] = "关机"
  //stats[1] = "运行"
  //stats[2] = "已删除"

  conn, err := libvirtconn()
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
  1: "开机",
  5: "关机",
  2: "deleted",
}


func Delete(uuid string) *Vms {
  var v = &Vms{}

  db := vmdb()
  db.Model(&v).Update("exist", 0).Where("uuid=?", uuid)

  db.Where("uuid = ?", uuid).First(&v)
  return v
}

func Shutdown(uuid string) (*Vms, error) {
  /*start vm*/
  conn, err := libvirtconn()
  if err != nil {
    return nil, err
  }
  vm, err4 := conn.LookupDomainByUUIDString(uuid)
  if err4 != nil {
    fmt.Println(err4)
    return nil, err4
  }
  err1 := vm.Destroy()
  if err1 != nil {
    return nil, err1
  }

  db := vmdb()
  var v = &Vms{}
  db.Where("uuid = ?", uuid).First(&v)

  s, err2 := VmStatus(uuid)
  v.Status = s
  if err2 != nil {
    fmt.Println(err2)
  }
  return v, err2
}

func Start(uuid string) (*Vms, error) {
  /*start vm*/
  conn, connerr := libvirtconn()
  if connerr != nil {
    return nil, connerr
  }
  vm, err := conn.LookupDomainByUUIDString(uuid)

  if err != nil {
    fmt.Println(err)
  }

  err1 := vm.Create()
  if err1 != nil {
    fmt.Println(err1)
  }

  db := vmdb()
  var v = &Vms{}
  db.Where("uuid = ?", uuid).First(&v)

  s, err2 := VmStatus(uuid)
  v.Status = s
  if err2 != nil {
    fmt.Println(err2)
  }
  return v, err2
}

func Createuuid() string {
  u := uuid.NewV4().String()
  return u
}

func savevm(uuid string, cpu int, mem int, vmxml string) bool {
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
  }
  db.Create(vm)

  //return bool
  res := db.NewRecord(&vm)
  return res
}

func Create(cpu int, mem int) (bool, error) {
  /*create a vm*/

  vcpu := cpu
  vmem := mem*1024*1024

  u := Createuuid()

  db := vmdb()
  var x Vm_xmls
  db.First(&x, "ostype = ?", "linux")

  vmxml := fmt.Sprintf(x.Osxml, u, u, vmem, vmem, vcpu)
  err := savevm(u, vcpu, vmem, vmxml)

  if err == false {
    fmt.Println("insert sql fail")
    return false, nil
  }

  conn, connerr := libvirtconn()
  if connerr != nil {
    return false, connerr
  }
  _, err1 := conn.DomainDefineXML(vmxml)

  if err1 != nil {
    return false, err1
  }
  return true, err1
}

func VmList() []*Vms {
  db := vmdb()
  var v []*Vms
  db.Where("exist=1").Find(&v)
  for _, e := range(v) {
    s, _ := VmStatus(e.Uuid)
    e.Status = s
  }
  return v
}

package vmcommon

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
  libvirt "libvirt.org/libvirt-go"
)

type Vms struct {
  Uuid string
  Name string
  Cpu int8
  Mem int8
  Createtime string
  Owner string
  Comment string
  Status int
}

func vmdb() *gorm.DB {
  db, errDb := gorm.Open("mysql", "modis:modis@(10.0.90.151:3306)/gocloud")
  if errDb != nil {
    fmt.Println(errDb)
  }
  return db
}


type Vm_xmls struct {
  Ostype string
  Osxml string
}

func libvirtconn() *libvirt.Connect {
  conn, err := libvirt.NewConnect("qemu:///system")
  if err != nil {
    fmt.Println(err)
  }
  return conn
}

type Mes struct {
  Res bool
}

func (m Mes)Error() string {
  return "vm already exists"
}

func vmexists(uuid string) (bool, error) {
  conn := libvirtconn()
  _, err := conn.LookupDomainByUUIDString(uuid)
  if err != nil {
    return false, err
  }
  return true, Mes{}
}

func Create(uuid string) (bool, error) {
  exists, err := vmexists(uuid)
  if exists == true {
    return false, err
  }

  db := vmdb()
  var x Vm_xmls
  db.First(&x, "ostype = ?", "linux")

  conn := libvirtconn()
  _, err1 := conn.DomainDefineXML(x.Osxml)

  if err1 != nil {
    return false, err1
  }
  return true, err1
}

func GetVmList() []*Vms {
  db := vmdb()
  var v []*Vms
  db.Find(&v)
  return v
}

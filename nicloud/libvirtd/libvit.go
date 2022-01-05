package libvirtd

import (
  "fmt"
  libvirt "github.com/libvirt/libvirt-go"
)

var Vmstate = map[libvirt.DomainState]string{
  1: "运行",
  5: "关机",
  3: "暂停",
  2: "deleted",
}


func Libvirtconn(host string) (*libvirt.Connect, error) {
  conn, err := libvirt.NewConnect(fmt.Sprintf("qemu+ssh://%s/system", host))
  if err != nil {
    return nil, err
  }
  return conn, err
}

func Migratevmlive() error {
  c, err := Libvirtconn("10.0.85.99")
  d, err := Libvirtconn("10.0.85.92")
  if err != nil {
    return err
  }
  a, err := c.LookupDomainByUUIDString("777e8cf6-1271-4e4f-81d4-ea4298cfb241")
  if err != nil {
    return err
  }
  _, err = a.Migrate(d, 1, "ysman", "tcp://10.0.85.92", 1024)
  if err != nil {
    return err
  }
  return nil
}

func DefineVm(xml string, host string) error {
  conn, err := Libvirtconn(host)
  if err != nil {
    return err
  }

  _, err = conn.DomainDefineXML(xml)
  if err != nil {
    return err
  }
  return err
}

func Undefine(host string, uuid string) error {
  //undefine vm
  conn, err :=  Libvirtconn(host)
  if err != nil {
    return err
  }

  vm, err := conn.LookupDomainByUUIDString(uuid)
  if err != nil {
    return err
  }
  err = vm.Undefine()
  if err != nil {
    return err
  }
  return nil
}

func GetDomain(host string, uuid string) (*libvirt.Domain, error) {
  conn, err :=  Libvirtconn(host)
  if err != nil {
    return nil, err
  }

  vm, err := conn.LookupDomainByUUIDString(uuid)

  if err != nil {
    return nil, err
  }
  return vm, nil
}

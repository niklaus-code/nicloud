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
  vm.Undefine()
  return nil
}

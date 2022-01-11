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

func Migratevmlive(uuid string, shost string, dhost string) error {
  c, err := Libvirtconn(shost)
  d, err := Libvirtconn(dhost)
  if err != nil {
    return err
  }
  a, err := c.LookupDomainByUUIDString(uuid)
  if err != nil {
    return err
  }
  _, err = a.Migrate(d, 1, uuid, fmt.Sprintf("tcp://%s", dhost), 1024)
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

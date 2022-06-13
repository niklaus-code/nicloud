package libvirtd

import "C"
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
  if err != nil {
    return err
  }
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

  defer conn.Close()
  dom, err := conn.LookupDomainByUUIDString(uuid)

  if err != nil {

    return nil, err
  }
  return dom, nil
}


//返回所有虚拟机接口参数
//const (
//  CONNECT_LIST_DOMAINS_ACTIVE         = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_ACTIVE)
//  CONNECT_LIST_DOMAINS_INACTIVE       = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_INACTIVE)
//  CONNECT_LIST_DOMAINS_PERSISTENT     = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_PERSISTENT)
//  CONNECT_LIST_DOMAINS_TRANSIENT      = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_TRANSIENT)
//  CONNECT_LIST_DOMAINS_RUNNING        = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_RUNNING)
//  CONNECT_LIST_DOMAINS_PAUSED         = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_PAUSED)
//  CONNECT_LIST_DOMAINS_SHUTOFF        = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_SHUTOFF)
//  CONNECT_LIST_DOMAINS_OTHER          = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_OTHER)
//  CONNECT_LIST_DOMAINS_MANAGEDSAVE    = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_MANAGEDSAVE)
//  CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE)
//  CONNECT_LIST_DOMAINS_AUTOSTART      = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_AUTOSTART)
//  CONNECT_LIST_DOMAINS_NO_AUTOSTART   = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_AUTOSTART)
//  CONNECT_LIST_DOMAINS_HAS_SNAPSHOT   = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_HAS_SNAPSHOT)
//  CONNECT_LIST_DOMAINS_NO_SNAPSHOT    = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_SNAPSHOT)
//  CONNECT_LIST_DOMAINS_HAS_CHECKPOINT = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_HAS_CHECKPOINT)
//  CONNECT_LIST_DOMAINS_NO_CHECKPOINT  = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_CHECKPOINT)
//)


func Listdomains(host string) ([]libvirt.Domain, error) {
  conn, err :=  Libvirtconn(host)
  if err != nil {
    return nil, err
  }

  defer conn.Close()
  vm, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_PERSISTENT)
  if err != nil {
    return nil, err
  }

  return vm, nil
}

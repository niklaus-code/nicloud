package libvirtd

import (
  "fmt"
  //"golang.org/x/text/language/display"

  //"github.com/beevik/etree"
  goxml "github.com/libvirt/libvirt-go-xml"
  "nicloud/cephcommon"
  "nicloud/networks"
  "nicloud/osimage"
  "strings"
)
//
//enum virNodeDeviceEventLifecycleType {
//
//VIR_NODE_DEVICE_EVENT_CREATED 	= 	0 (0x0)
//VIR_NODE_DEVICE_EVENT_DELETED 	= 	1 (0x1)
//VIR_NODE_DEVICE_EVENT_DEFINED 	= 	2 (0x2)
//VIR_NODE_DEVICE_EVENT_UNDEFINED 	= 	3 (0x3)
//VIR_NODE_DEVICE_EVENT_LAST 	= 	4 (0x4
//
//}

const(
  Name string = "qemu"
  Type string = "raw"
  Username string = "admin"
  Protocol string = "rbd"
  Device string = "disk"
  Secret_type = "ceph"
)
var Slotlist = map[string] uint {"vda": 10, "vdb": 11, "vdc": 12, "vdd": 13, "vde": 14, "vdf": 15, "bridge": 16}
var Win_Slotlist = map[string] uint {"sda": 10, "sdb": 11, "sdc": 12, "sdd": 13, "sde": 14, "sdf": 15, "bridge": 16}

func xml_Pci(slot uint) *goxml.DomainAddressPCI {
  var Domain uint = 00
  var Bus uint = 00
  var Function uint = 0
  PCI := &goxml.DomainAddressPCI{
    Domain:   &Domain,
    Bus:      &Bus,
    Slot:     &slot,
    Function: &Function,
  }
  return PCI
}

func xml_DomainDiskDriver() *goxml.DomainDiskDriver {
  driver := goxml.DomainDiskDriver{
    Name: Name,
    Type: Type,
  }
  return &driver
}

func xml_DomainDiskSecret(s string) goxml.DomainDiskSecret {
  secret := goxml.DomainDiskSecret{
    Type: Secret_type,
    UUID: s,
  }
  return secret
}

func xml_DomainDiskAuth(s goxml.DomainDiskSecret) *goxml.DomainDiskAuth {
  auth := goxml.DomainDiskAuth{
    Username: Username,
    Secret: &s,
  }
  return &auth
}

func xml_DomainDiskSourceHost(h []string, port string) []goxml.DomainDiskSourceHost {
  host := []goxml.DomainDiskSourceHost{}
  for _, v := range h {
    hostobj := goxml.DomainDiskSourceHost{
      Name: v,
      Port: port,
    }
    host = append(host, hostobj)
  }

  return host
}

func xml_DomainDiskSourceNetwork(h []goxml.DomainDiskSourceHost, pool string, uuid string) *goxml.DomainDiskSourceNetwork {
  Network := &goxml.DomainDiskSourceNetwork{
    Protocol: Protocol,
    Name: fmt.Sprintf("%s/%s", pool, uuid),
    Hosts: h,
  }
  return Network
}

func xml_DomainDiskSource(h *goxml.DomainDiskSourceNetwork) *goxml.DomainDiskSource {
  source := goxml.DomainDiskSource{
    Network: h,
  }
  return &source
}

func xml_DomainDeviceBoot(o uint) *goxml.DomainDeviceBoot {
  order := &goxml.DomainDeviceBoot{
    Order: o,
  }
  return order
}

func diskxml(iplist[]string, port string, poolname string, uuid string, secret string, diskname string, order_check bool, os string) (goxml.DomainDisk, error) {
  disk := goxml.DomainDisk{}
  disk.Device = Device
  disk.Driver = xml_DomainDiskDriver()
  disk.Auth = xml_DomainDiskAuth(xml_DomainDiskSecret(secret))
  disk.Source = xml_DomainDiskSource(xml_DomainDiskSourceNetwork(xml_DomainDiskSourceHost(iplist, port), poolname, uuid))
  disk.Address = &goxml.DomainAddress{
      PCI: xml_Pci(Slotlist[diskname]),
    }
  if os == "LINUX" {
    if order_check {
      disk.Boot = xml_DomainDeviceBoot(1)
    }
    disk.Target = &goxml.DomainDiskTarget{
      Dev: diskname,
      Bus: "virtio",
    }
  } else {
    disk.Target = &goxml.DomainDiskTarget{
      Dev: diskname,
      Bus: "sata",
    }
  }

  return disk, nil
}

func xml_DomainVCPU(cpu uint) *goxml.DomainVCPU {
  c := goxml.DomainVCPU{
    Placement: "static",
    Value: cpu,
  }
  return &c
}

func xml_DomainMemory(mem uint) *goxml.DomainMemory {
  m := goxml.DomainMemory{
    Unit: "KiB",
    Value: mem,
  }
  return &m
}

func xml_DomainCurrentMemory(cmem uint) *goxml.DomainCurrentMemory {
  m := goxml.DomainCurrentMemory{
    Unit: "KiB",
    Value: cmem,
  }
  return &m
}

func xml_DomainInterfaceMAC(mac string) *goxml.DomainInterfaceMAC {
  m := &goxml.DomainInterfaceMAC{
    Address: mac,
  }
  return m
}

func xml_DomainInterfaceModel() *goxml.DomainInterfaceModel {
  i := &goxml.DomainInterfaceModel{
    Type: "virtio",
  }
  return i
}

func xml_DomainInterfaceSource(bridge string) *goxml.DomainInterfaceSource {
  b := &goxml.DomainInterfaceSource{
    Bridge: &goxml.DomainInterfaceSourceBridge{
      Bridge: bridge,
    },
  }
  return b
}

func xml_DomainAddress(adderss uint) *goxml.DomainAddress {
  a := &goxml.DomainAddress{
    PCI: xml_Pci(adderss),
  }
  return a
}

func xml_bridge(bridge string, mac string, os string) goxml.DomainInterface {
  m := goxml.DomainInterface{}
  m.Model = xml_DomainInterfaceModel()
  if os == "LINUX" {
    m.Boot = xml_DomainDeviceBoot(2)
  }
  m.MAC = xml_DomainInterfaceMAC(mac)
  m.Source = xml_DomainInterfaceSource(bridge)
  m.Address = xml_DomainAddress(Slotlist["bridge"])
  return m
}

func CreateVmXml(datacenter string, storage string, vlan string,  vcpu uint, vmem uint, uuid string, mac string, image_name string, osid int, pool string, os string) (string, error) {
  ceph := cephcommon.Vms_Ceph{}
  storagename, err := ceph.Cephinfobyuuid(storage)
  if err != nil {
    return "", err
  }

  ips := strings.Split(storagename.Ips, ",")
  port := storagename.Port

  bridge, err := networks.Getbridge(datacenter, vlan)
  if err != nil {
    return "", err
  }

  osinfo, err := osimage.GetOsInfoById(storage, osid)
  if err != nil {
    return "", err
  }

  domcfg := &goxml.Domain{}
  err = domcfg.Unmarshal(osinfo.Xml)
  if err != nil {
    return "", err
  }

  disk, err := diskxml(ips, port, pool, image_name, storagename.Ceph_secret, "vda", true, os)
  if err != nil {
    return "", err
  }
  domcfg.Devices.Disks = append(domcfg.Devices.Disks, disk)
  domcfg.VCPU = xml_DomainVCPU(vcpu)
  domcfg.UUID = uuid
  domcfg.Name = uuid
  domcfg.Memory = xml_DomainMemory(vmem)
  domcfg.CurrentMemory = xml_DomainCurrentMemory(vmem)
  domcfg.Devices.Interfaces = append(domcfg.Devices.Interfaces, xml_bridge(bridge, mac, os))

  xmlstr, err := domcfg.Marshal()
  if err != nil {
    return "", err
  }

  return xmlstr, nil
}

func CreateDiskXml(xml string, ceph_block string, ips []string, port string, pool string, disknum int, diskname string, secret string, os string) (string, error) {
  domcfg := &goxml.Domain{}
  err := domcfg.Unmarshal(xml)
  if err != nil {
    return "", err
  }

  disk, err := diskxml(ips, port, pool, ceph_block, secret, diskname, false, os)
  if err != nil {
    return "", err
  }

  domcfg.Devices.Disks = append(domcfg.Devices.Disks, disk)

  xmlstr, err := domcfg.Marshal()
  if err != nil {
    return "", err
  }

  return xmlstr, nil
}

func RemoveDiskXml(xml string, ceph_block string, pool string) (string, error) {
  domcfg := &goxml.Domain{}
  err := domcfg.Unmarshal(xml)
  if err != nil {
    return "", err
  }

  disklist := domcfg.Devices.Disks
  for i :=0; i <= len(disklist); i++ {
    if disklist[i].Source.Network.Name == fmt.Sprintf("%s/%s", pool, ceph_block) {
      disklist = append(disklist[:i ], disklist[i+1:]...)
    }
  }

  domcfg.Devices.Disks = disklist

  xmlstr, err := domcfg.Marshal()
  if err != nil {
    return "", err
  }

  return xmlstr, nil
}

func UpdateCpuMem(xml string, vcpu uint, vmem uint) (string, error) {
  domcfg := &goxml.Domain{}
  err := domcfg.Unmarshal(xml)
  if err != nil {
    return "", err
  }

  domcfg.VCPU = xml_DomainVCPU(vcpu)
  domcfg.Memory = xml_DomainMemory(vmem)
  domcfg.CurrentMemory = xml_DomainCurrentMemory(vmem)

  xmlstr, err := domcfg.Marshal()
  if err != nil {
    return "", err
  }
  return xmlstr, nil
}

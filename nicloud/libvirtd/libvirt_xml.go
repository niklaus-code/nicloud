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

func XML_Pci(slot uint) *goxml.DomainAddressPCI {
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

func XML_Driver() *goxml.DomainDiskDriver {
  driver := goxml.DomainDiskDriver{
    Name: Name,
    Type: Type,
  }
  return &driver
}

func XML_Secret(s string) goxml.DomainDiskSecret {
  secret := goxml.DomainDiskSecret{
    Type: Secret_type,
    UUID: s,
  }
  return secret
}

func XML_Auth(s goxml.DomainDiskSecret) *goxml.DomainDiskAuth {
  auth := goxml.DomainDiskAuth{
    Username: Username,
    Secret: &s,
  }
  return &auth
}

func XML_Host(h []string, port string) []goxml.DomainDiskSourceHost {
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

func XML_Network(h []goxml.DomainDiskSourceHost, pool string, uuid string) *goxml.DomainDiskSourceNetwork {
  Network := &goxml.DomainDiskSourceNetwork{
    Protocol: Protocol,
    Name: fmt.Sprintf("%s/%s", pool, uuid),
    Hosts: h,
  }
  return Network
}

func XML_Source(h *goxml.DomainDiskSourceNetwork) *goxml.DomainDiskSource {
  source := goxml.DomainDiskSource{
    Network: h,
  }
  return &source
}

func xml_order() *goxml.DomainDeviceBoot {
  order := &goxml.DomainDeviceBoot{
    Order: 1,
  }
  return order
}

func diskxml(iplist[]string, port string, poolname string, uuid string, secret string, diskname string, order_check bool) (goxml.DomainDisk, error) {
  disk := goxml.DomainDisk{}
  disk.Device = Device
  disk.Driver = XML_Driver()
  disk.Auth = XML_Auth(XML_Secret(secret))
  disk.Source = XML_Source(XML_Network(XML_Host(iplist, port), poolname, uuid))
  disk.Target = &goxml.DomainDiskTarget{
      Dev: diskname,
      Bus: "virtio",
    }
  disk.Address = &goxml.DomainAddress{
      PCI: XML_Pci(Slotlist[diskname]),
    }
  if order_check {
    disk.Boot = xml_order()
  }

  return disk, nil
}

func xml_cpu(cpu uint) *goxml.DomainVCPU {
  c := goxml.DomainVCPU{
    Placement: "static",
    Value: cpu,
  }
  return &c
}

func xml_mem(mem uint) *goxml.DomainMemory {
  m := goxml.DomainMemory{
    Unit: "KiB",
    Value: mem,
  }
  return &m
}

func xml_currentmem(cmem uint) *goxml.DomainCurrentMemory {
  m := goxml.DomainCurrentMemory{
    Unit: "KiB",
    Value: cmem,
  }
  return &m
}

func xml_bridge(bridge string, mac string) goxml.DomainInterface {
  m := goxml.DomainInterface{
    Model: &goxml.DomainInterfaceModel{
      Type: "virtio",
    },
    Boot: &goxml.DomainDeviceBoot{
      Order: 2,
    },
    MAC: &goxml.DomainInterfaceMAC{
      Address: mac,
    },
    Source: &goxml.DomainInterfaceSource{
      Bridge: &goxml.DomainInterfaceSourceBridge{
        Bridge: bridge,
      },
    },
    Address: &goxml.DomainAddress{
      PCI: XML_Pci(Slotlist["bridge"]),
    },
  }
  return m
}

func CreateVmXml(datacenter string, storage string, vlan string,  vcpu uint, vmem uint, uuid string, mac string, image_name string, osid int, pool string) (string, error) {
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

  disk, err := diskxml(ips, port, pool, image_name, storagename.Ceph_secret, "vda", true)
  if err != nil {
    return "", err
  }
  domcfg.Devices.Disks[0] = disk
  domcfg.VCPU = xml_cpu(vcpu)
  domcfg.UUID = uuid
  domcfg.Name = uuid
  domcfg.Memory = xml_mem(vmem)
  domcfg.CurrentMemory = xml_currentmem(vmem)
  domcfg.Devices.Interfaces[0] = xml_bridge(bridge, mac)

  xmlstr, err := domcfg.Marshal()
  if err != nil {
    return "", err
  }

  return xmlstr, nil
}

func CreateDiskXml(xml string, ceph_block string, ips []string, port string, pool string, disknum int, diskname string, secret string) (string, error) {
  domcfg := &goxml.Domain{}
  err := domcfg.Unmarshal(xml)
  if err != nil {
    return "", err
  }

  disk, err := diskxml(ips, port, pool, ceph_block, secret, diskname, false)
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
  for index, disk := range disklist {
    if disk.Source.Network.Name == fmt.Sprintf("%s/%s", pool, ceph_block) {
      disklist = append(disklist[:index], disklist[index+1:]...)
    }
  }

  domcfg.Devices.Disks = disklist

  xmlstr, err := domcfg.Marshal()
  if err != nil {
    return "", err
  }

  return xmlstr, nil
}


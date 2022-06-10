package vm

import (
  "fmt"
  "nicloud/libvirtd"
  goxml "github.com/libvirt/libvirt-go-xml"
  "nicloud/osimage"
  "strconv"
  "time"
)

var LINUX_disk = "vda"
var WIN_disk = "sda"

type DiskIO struct {
  Read float64
  Write float64
  Ctime string
}

func Diskinfo(host string, uuid string, ostype string) (*DiskIO, error)  {
  vm, err := GetVmByUuid(uuid)
  if err != nil {
    return nil, err
  }

  o := osimage.Vms_os{}
  os, err := o.GetOsInfoById(vm.Storage, vm.Os)
  if err != nil {
    return nil, err
  }

  t := osimage.Vms_os_tags{}
  tag, err := t.GetostagByid(os.Tag)
  if err != nil {
    return nil, err
  }

  domain, err := libvirtd.GetDomain(host, uuid)
  if err != nil {
    return nil, err
  }

  var diskname string

  //需要配置化
  if tag.Tag == "LINUX" {
    diskname = LINUX_disk
  } else {
    diskname = WIN_disk
  }

  vmstat, err := domain.BlockStats(diskname)
  if err != nil {
    return nil, err
  }
  readbytes_before := vmstat.RdBytes
  writebytes_before := vmstat.WrBytes

  Now := time.Now()
  ticker := time.NewTicker(1 * time.Second)
  <-ticker.C

  vmstat_after, err := domain.BlockStats(diskname)
  if err != nil {
    return nil, err
  }

  readbytes_after := vmstat_after.RdBytes
  writebytes_after := vmstat_after.WrBytes

  read, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(readbytes_after-readbytes_before)/float64(1024)), 64)
  write, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(writebytes_after-writebytes_before)/float64(1024)), 64)

  dio := &DiskIO{
    Read: read,
    Write: write,
    Ctime: Now.Format("15:04:05"),
  }

  return dio, nil
}

type MemCount struct {
  Mem_used float64
  Mem_available float64
  Ctime string
}
func Meminfo(host string, uuid string) (*MemCount, error){
  Now := time.Now()
  domain, err := libvirtd.GetDomain(host, uuid)
  if err != nil {
    return nil, err
  }

  info, err := domain.MemoryStats(12, 0)
  //虚拟机分配的总内存
  //mem_actual := info[0].Val

  //虚拟机未使用的内存
  mem_unused := info[7].Val

  //虚拟机识别到的总内存
  mem_available := info[6].Val

  //占用宿主机内存
  //mem_rss := info[9].Val

  mem_total, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(mem_available)/float64(1024*1024)), 64)
  if err != nil {
    return nil, err
  }
  mem_used, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(mem_available-mem_unused)/float64(1024*1024)), 64)
  if err != nil {
    return nil, err
  }

  m := &MemCount{
    Mem_used: mem_used,
    Mem_available: mem_total,
    Ctime: Now.Format("15:04:05"),
  }
  return m, nil
}

type CpuLoad struct {
  Load float64
  Ctime string
}

func Cpuinfo(host string, uuid string) (*CpuLoad, error){
  domain, err := libvirtd.GetDomain(host, uuid)
  if err != nil {
    return nil, err
  }

  v, err := domain.GetInfo()
  if err != nil {
    return nil, err
  }

  Now := time.Now()
  cpu_before, err := domain.GetCPUStats(-1, 1, 0)
  if err != nil {
    return nil, err
  }

  ticker := time.NewTicker(1 * time.Second)
  <-ticker.C

  cpu_after, err := domain.GetCPUStats(-1, 1, 0)
  if err != nil {
    return nil, err
  }
  cpu_load, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64((cpu_after[0].CpuTime-cpu_before[0].CpuTime)*uint64(v.NrVirtCpu))/float64(1e9)*100), 64)
  if err != nil {
    return nil, err
  }
  c := &CpuLoad{
    Load: cpu_load,
    Ctime: Now.Format("15:04:05"),
  }

  return c, nil
}

type NetSpeed struct {
  Tx float64
  Rx float64
  Ctime string
}

func Netinfo(host string, uuid string) (*NetSpeed, error){
  domain, err := libvirtd.GetDomain(host, uuid)
  if err != nil {
    return nil, err
  }

  xml, err := domain.GetXMLDesc(1)
  if err != nil {
    return nil, err
  }

  domcfg := &goxml.Domain{}
  err = domcfg.Unmarshal(xml)
  if err != nil {
    return nil, err
  }

  dev := domcfg.Devices.Interfaces[0].Target.Dev

  net_before, err := domain.InterfaceStats(dev)
  if err != nil {
   return nil, err
  }

  rxbytes_before := net_before.RxBytes
  txbytes_before := net_before.TxBytes

  Now := time.Now()
  ticker := time.NewTicker(1 * time.Second)
  <-ticker.C

  net_after, err := domain.InterfaceStats(dev)
  rxbytes_after := net_after.RxBytes
  txbytes_after := net_after.TxBytes

  rx, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(rxbytes_after-rxbytes_before)/float64(1024)), 64)
  tx, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(txbytes_after-txbytes_before)/float64(1024)), 64)

  n := &NetSpeed{
    Rx: rx,
    Tx: tx,
    Ctime: Now.Format("15:04:05"),
  }

  return n, nil
}

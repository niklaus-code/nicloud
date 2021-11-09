package vm

import (
  "errors"
  "fmt"
  "github.com/beevik/etree"
  _ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
  "goblog/ceph"
  "goblog/dbs"
  "goblog/libvirtd"
  "goblog/networks"
  "goblog/osimage"
  "goblog/utils"
  vmerror "goblog/vmerror"
  "strings"
  "time"
)

type Vms struct {
	Uuid        string
	Name        string
	Cpu         int
	Mem         int
	Create_time time.Time
	Owner       string
	Comment     string
	Vmxml       string
	Status      string
	Exist       int
	Ip          string
	Host        string
	Os          string
	Datacenter  string
	Storage     string
}

func GetVmByUuid(uuid string) *Vms {
  dbs, err := db.NicloudDb()
  v := &Vms{}
  if err != nil {
    return nil
  }
  dbs.Where("uuid = ?", uuid).First(v)
  return v
}

func (v Vms) Error(info string) error {
	errorinfo := fmt.Sprintf("%s", info)
	return errors.New(errorinfo)
}

type Vm_xmls struct {
	Ostype string
	Osxml  string
}

func VmStatus(uuid string, host string) (string, error) {
	conn, err := libvirtd.Libvirtconn(host)

	if err != nil {
		return "", err
	}
	vm, err := conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return "vm not found", err
	}

	state, _, err1 := vm.GetState()

	if err1 != nil {
		return "vm not found", err1
	}

	return libvirtd.Vmstate[state], err1
}

type Vms_archive struct {
  Uuid string
  Create_time time.Time
  Owner string
  Comment string
  Vmxml string
  Ip string
  Datacenter string
  Storage string
}

func Delete(uuid string) ([]*Vms, error) {
  vminfo := GetVmByUuid(uuid)
  host := vminfo.Host

	vmstat, err := VmStatus(uuid, host)
	if err != nil {
		return nil, err
	}

	if vmstat == "运行" {
		return nil, vmerror.Error{
			Message: "vm is running, con't delete",
		}
	}

  if vmstat == "暂停" {
    return nil, vmerror.Error{
      Message: "vm is paused, con't delete",
    }
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return  nil, err
  }

  err = libvirtd.Undefine(host, uuid)
  if err != nil {
    return nil, err
  }

	err = ceph.Rm_image(uuid)
  if err != nil {
    return nil, err
  }

	dbs.Model(&Vms{}).Where("uuid=?", uuid).Delete(&Vms{})
	err = networks.Updateipstatus(vminfo.Ip, 0)
	if err != nil {
	  return nil, err
  }
	v := &Vms_archive{
	  Uuid: vminfo.Uuid,
	  Owner: vminfo.Owner,
	  Comment: vminfo.Comment,
	  Ip: vminfo.Ip,
	  Vmxml: vminfo.Vmxml,
	  Datacenter: vminfo.Datacenter,
	  Storage: vminfo.Storage,
  }
  err2 := dbs.Create(*v)
  if err2.Error != nil {
    return nil, err2.Error
  }

  err = Freehost(vminfo.Host, vminfo.Cpu, vminfo.Mem)
  if err != nil {
    return nil, err
  }
	vmlist := VmList(host)
	return vmlist, err
}

func PauseVm(uuid string, host string) (*Vms, error) {
  conn, err := libvirtd.Libvirtconn(host)
  if err != nil {
    return nil, err
  }
  vm, err1 := conn.LookupDomainByUUIDString(uuid)
  if err1 != nil {
    return nil, err1
  }

  err = vm.Suspend()
  if err != nil {
    return nil, err
  }

  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var v = &Vms{}
  db.Where("uuid = ?", uuid).First(&v)

  s, err := VmStatus(uuid, host)
  v.Status = s
  if err != nil {
    return nil, err
  }
  return v, err
}

func Shutdown(uuid string, host string) (*Vms, error) {
	/*start vm*/
	conn, err := libvirtd.Libvirtconn(host)
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

  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
	var v = &Vms{}
	dbs.Where("uuid = ?", uuid).First(&v)

	s, err2 := VmStatus(uuid, host)
	v.Status = s
	if err2 != nil {
		return nil, err2
	}
	return v, err2
}

func Start(uuid string, host string) (*Vms, error) {
	/*start vm*/

	conn, connerr := libvirtd.Libvirtconn(host)
	if connerr != nil {
		return nil, connerr
	}
	vm, err := conn.LookupDomainByUUIDString(uuid)

	if err != nil {
		return nil, err
	}

  vm1, err1 := VmStatus(uuid, host)
	if err1 != nil {
	  return nil, err1
  }

  if vm1 == "暂停" {
    eer := vm.Resume()
    if eer != nil {
      return nil, eer
    }
  } else {
    err2 := vm.Create()
    if err2 != nil {
      return nil, err2
    }
  }

  dbs, err3 := db.NicloudDb()
  if err3 != nil {
    return nil, err3
  }
	var v = &Vms{}
	dbs.Where("uuid = ?", uuid).First(&v)

	s, err4 := VmStatus(uuid, host)
	v.Status = s
	if err4 != nil {
		return nil, err4
	}

	return v, err4
}

func savevm(datacenter string, cephname string, uuid string, cpu int, mem int, vmxml string, ip string, host string, image string) (bool, error) {
  /*save config to db*/
  dbs, err := db.NicloudDb()
  if err != nil {
    return false, err
  }
	vm := &Vms{
		Uuid:        uuid,
		Name:        uuid,
		Cpu:         cpu,
		Mem:         mem,
		Vmxml:       vmxml,
		Create_time: time.Now(),
		Exist:       1,
		Status:      "",
		Ip:          ip,
		Host:        host,
		Owner:       "Niklaus",
		Os:          image,
		Datacenter: datacenter,
		Storage: cephname,
	}

	err1 := dbs.Create(*vm)
	if err1.Error != nil {
	    return false, err1.Error
  }

	//return bool
	res := dbs.NewRecord(&vm)
	return res, err1.Error
}


func MigrateVm(uuid string, migrate_host string) error {
  vm := GetVmByUuid(uuid)
  err := libvirtd.DefineVm(vm.Vmxml, migrate_host)
  if err != nil {
    return err
  }
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  err1 := dbs.Model(&Vms{}).Where("uuid=?", uuid).Update("host", migrate_host)
  if err1.Error != nil {
      return err1.Error
  }

  err = libvirtd.Undefine(vm.Host, vm.Uuid)
  if err != nil {
    return err
  }

  return err
}

func Create(datacenter string,  storage string, vlan string, cpu int, mem int, ip string, host string, image string) (bool, error) {
  mac, err := networks.Ipresource(ip)
  if err != nil {
    return false, err
  }

	/*create a vm*/
	vcpu := cpu
	vmem := mem * 1024 * 1024

	//create a uuid
	u := utils.Createuuid()

	//create baseimage
	imge_name, err := ceph.RbdClone(u)
	if err != nil {
	 return false, err
  }

	f, err := osimage.Xml(datacenter, storage, vlan,  vcpu, vmem, u, mac, imge_name, image)
	if err != nil {
	  return false, err
  }

	err = libvirtd.DefineVm(f, host)
	if err != nil {
	  return false, err
  }

  err = Updatehost(host, cpu, mem)
  if  err != nil {
    return false, err
  }
	svm, err := savevm(datacenter, storage, u, cpu, mem, f, ip, host, image)
	if err != nil {
	  return svm, err
  }

  err = networks.Updateipstatus(ip, 1)
  if err != nil {
    return false, err
  }

	return true, err
}

func Getvmxmlby (ip string, storage string, datacenter string) (string, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return "", err
  }
  v := &Vms{}
  errdb := dbs.Where("ip=? and storage=? and datacenter=?",ip, storage, datacenter).Find(v)
  if errdb.Error != nil {
    return "", vmerror.Error{Message: errdb.Error.Error()}
  }
  return v.Vmxml, nil
}

func Umountdisk(vmip string,  storage string, datacenter string, diskid string) error {
  f, err := Getvmxmlby(vmip, storage, datacenter)
  if err != nil {
    return err
  }

  host, err := GetHostsbyvmip(vmip)

  if err != nil {
    return err
  }

  doc := etree.NewDocument()
  err = doc.ReadFromString(f)
  device := doc.FindElements("./domain/devices/disk")
  d:= doc.FindElement("./domain/devices/")
  for _, v := range device {
    source := v.FindElement("./source")
    vmdisk := source.SelectAttr("name").Value
    uuid := strings.Split(vmdisk, "/")
    if len(uuid)> 1 && uuid[1] == diskid {
      d.RemoveChild(v)
      var docstring string
      docstring, err = doc.WriteToString()
      libvirtd.DefineVm(docstring, host.Host)

      err := ceph.Umountvmstatus(datacenter, storage, diskid)
      if err != nil {
        return err
      }
      return nil
    }
  }

  return vmerror.Error{
    Message: "disk not found",
  }
}

func Updatexml(vmid string, ip string, vmhost string, storage string, pool string, datacenter string, cloudriveid string) error {
  s, err := VmStatus(vmid, vmhost)
  if err != nil {
    return err
  }
  if s != "关机" {
    return vmerror.Error{Message: "cont mount disk, vm is " + s}
  }

  storageinfo, err := ceph.Cephinfobyname(datacenter, storage)
  if err != nil {
    return err
  }

  f, err := Getvmxmlby(ip, storage, datacenter)
  if err != nil {
    return err
  }

  doc := etree.NewDocument()
  err = doc.ReadFromString(f)
  if err != nil {
    return err
  }
  device := doc.FindElement("./domain/devices")
  disk := device.CreateElement("disk")
  disk.CreateAttr("type", "network")
  disk.CreateAttr("device", "disk")

  driver := disk.CreateElement("driver")
  driver.CreateAttr("name", "qemu")
  driver.CreateAttr("type", "raw")

  auth := disk.CreateElement("auth")
  auth.CreateAttr("username", "admin")
  secret := auth.CreateElement("secret")
  secret.CreateAttr("type", "ceph")
  secret.CreateAttr("uuid", storageinfo[0].Ceph_secret)

  source := disk.CreateElement("source")
  source.CreateAttr("protocol", "rbd")

  source.CreateAttr("name", pool+"/"+cloudriveid)
  host := source.CreateElement("host")


  var iplist []string
  iplist = strings.Split(storageinfo[0].Ips, ",")
  for _, v := range iplist {
    host.CreateAttr("name", v)
  }
  host.CreateAttr("port", storageinfo[0].Port)

  target := disk.CreateElement("target")
  target.CreateAttr("dev", "vdb")
  target.CreateAttr("bus", "virtio")

  address := disk.CreateElement("address")
  address.CreateAttr("type", "pci")
  address.CreateAttr("domain", "0x0000")
  address.CreateAttr("bus", "0x00")
  address.CreateAttr("slot", "0x19")
  address.CreateAttr("function", "0x0")
  doc.Indent(2)
  var docstring string
  docstring, err = doc.WriteToString()

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb := dbs.Model(Vms{}).Where("ip=?", ip).Update("vmxml", docstring)
  if errdb.Error != nil {
    return vmerror.Error{Message: errdb.Error.Error()}
  }

  err = libvirtd.DefineVm(docstring, vmhost)
  if err != nil {
    return err
  }

  updatevm := ceph.Mountvmstatus(datacenter, storage, cloudriveid, ip)
  if updatevm != nil {
    return updatevm
  }
  return nil
}

func VmList(host string) []*Vms {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil
  }
	var v []*Vms
	dbs.Where("exist=1").Find(&v)

	return v
}

type Vm_flavors struct {
	Cpu int
	Mem int
}

func Flavor() ([]*Vm_flavors, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
	var f []*Vm_flavors
	dbs.Find(&f)
	return f, nil
}

func SearchVm(c string) ([]*Vms, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var v []*Vms
  i := fmt.Sprintf("ip like %s", "'"+c+"%'")
  dbs.Where(i).Find(&v)

  return v, nil
}

func Updatecomments(uuid string, comment string) (bool, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return true, err
  }
  dbs.Model(&Vms{}).Where("uuid=?", uuid).Update("comment", comment)
  return true, nil
}


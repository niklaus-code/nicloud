package vm

import (
  "errors"
  "fmt"
  "github.com/beevik/etree"
  _ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
  "nicloud/cephcommon"
  "nicloud/dbs"
  "nicloud/libvirtd"
  "nicloud/networks"
  "nicloud/osimage"
  "nicloud/utils"
  vdisk "nicloud/vdisk"
  vmerror "nicloud/vmerror"
  "reflect"
  "time"
  "encoding/base64"
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

func updatexmlbyuuid(xml string, uuid string, vcpu int, vmem int) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb:= dbs.Model(&Vms{}).Where("uuid=?", uuid).Update("vmxml", xml).Update("cpu", vcpu).Update("mem", vmem)
  if errdb.Error != nil {
    return vmerror.Error{Message: errdb.Error.Error()}
  }

  return nil
}

func Changeconfig(uuid string, host string, vcpu int, vmem int, vmhost string) error {
  m := vmem * 1024 * 1024
  s, err := VmStatus(uuid, host)
  if s != "关机" {
    return vmerror.Error{Message: "云主机需要关机状态"}
  }

  updatehost := Updatehostcpumem(host, vcpu, vmem)
  if updatehost != nil {
    return updatehost
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  v := &Vms{}
  errdb := dbs.Where("uuid = ?", uuid).First(v)
  if errdb.Error != nil {
    return vmerror.Error{Message: "未发现云主机"}
  }

  xml := v.Vmxml
  doc := etree.NewDocument()
  err = doc.ReadFromString(xml)
  if err != nil {
    return err
  }

  cpu := doc.FindElement("./domain/vcpu")
  cpu.SetText(fmt.Sprintf("%d", vcpu))

  mem := doc.FindElement("./domain/memory")
  mem.SetText(fmt.Sprintf("%d", m))

  currentMemory := doc.FindElement("./domain/currentMemory")
  currentMemory.SetText(fmt.Sprintf("%d", m))

  doc.Indent(2)
  var docstring string
  docstring, err = doc.WriteToString()

  updatexml := updatexmlbyuuid(docstring, uuid, vcpu, vmem)
  if updatexml != nil {
    return updatexml
  }

  err = libvirtd.DefineVm(docstring, vmhost)
  if err != nil {
    return err
  }
  return nil
}

func GetVmByUuid(uuid string) *Vms {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil
  }

  v := &Vms{}
  dbs.Where("uuid = ?", uuid).First(v)
  return v
}

func GetVmByIp(ip string) *Vms {
  dbs, err := db.NicloudDb()
  v := &Vms{}
  if err != nil {
    return nil
  }
  dbs.Where("ip = ?", ip).First(v)
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

func Delete(uuid string) (error) {
  vminfo := GetVmByUuid(uuid)
  host := vminfo.Host

	vmstat, err := VmStatus(uuid, host)
	if err != nil {
		return err
	}

	if vmstat == "运行" {
		return vmerror.Error{
			Message: "vm is running, con't delete",
		}
	}

  if vmstat == "暂停" {
    return vmerror.Error{
      Message: "vm is paused, con't delete",
    }
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  err = libvirtd.Undefine(host, uuid)
  if err != nil {
    return err
  }

	err = cephcommon.Rm_image(uuid)
  if err != nil {
    return err
  }

	dbs.Model(&Vms{}).Where("uuid=?", uuid).Delete(&Vms{})
	err = networks.Updateipstatus(vminfo.Ip, 0)
	if err != nil {
	  return err
  }

	v := &Vms_archive{
	  Uuid: vminfo.Uuid,
	  Owner: vminfo.Owner,
	  Comment: vminfo.Comment,
	  Ip: vminfo.Ip,
	  Vmxml: vminfo.Vmxml,
	  Datacenter: vminfo.Datacenter,
	  Storage: vminfo.Storage,
	  Create_time: time.Now(),
  }
  err2 := dbs.Create(*v)
  if err2.Error != nil {
    return err2.Error
  }

  err = Freehost(vminfo.Host, vminfo.Cpu, vminfo.Mem)
  if err != nil {
    return err
  }

  updatevdisk := vdisk.Updatevdiskbydelvm(vminfo.Datacenter, vminfo.Storage, vminfo.Ip)
  if updatevdisk != nil {
    return updatevdisk
  }
	return nil
}

func PauseVm(uuid string, host string) error {
  conn, err := libvirtd.Libvirtconn(host)
  if err != nil {
    return  err
  }
  vm, err1 := conn.LookupDomainByUUIDString(uuid)
  if err1 != nil {
    return err1
  }

  err = vm.Suspend()
  if err != nil {
    return err
  }
  return nil
}

func Shutdown(uuid string, host string) error {
  /*start vm*/
  conn, err := libvirtd.Libvirtconn(host)
  if err != nil {
    return err
  }
  vm, err4 := conn.LookupDomainByUUIDString(uuid)
  if err4 != nil {
    return err4
  }
  err1 := vm.Shutdown()
  if err1 != nil {
    return err1
  }

  return nil
}

func Destroy(uuid string, host string) error {
	/*start vm*/
	conn, err := libvirtd.Libvirtconn(host)
	if err != nil {
		return err
	}
	vm, err4 := conn.LookupDomainByUUIDString(uuid)
	if err4 != nil {
		return err4
	}
	err1 := vm.Destroy()
	if err1 != nil {
		return err1
	}
	return nil
}

func Start(uuid string, host string) error {
	/*start vm*/

	conn, connerr := libvirtd.Libvirtconn(host)
	if connerr != nil {
		return connerr
	}
	vm, err := conn.LookupDomainByUUIDString(uuid)

	if err != nil {
		return err
	}

  vm1, err1 := VmStatus(uuid, host)
	if err1 != nil {
	  return err1
  }

  if vm1 == "暂停" {
    eer := vm.Resume()
    if eer != nil {
      return  eer
    }
  } else {
    err2 := vm.Create()
    if err2 != nil {
      return err2
    }
  }
	return nil
}

func savevm(datacenter string, cephname string, uuid string, cpu int, mem int, vmxml string, ip string, host string, image string) (string, error) {
  /*save config to db*/
  dbs, err := db.NicloudDb()
  if err != nil {
    return "", err
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

	err1 := dbs.Create(vm)
	if err1.Error != nil {
	    return "", err1.Error
  }
	return vm.Uuid, err1.Error
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

func deletevmbyid(uuid string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  errdb := dbs.Model(Vms{}).Where("uuid=?", uuid).Delete(Vms{})
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}

func Create(datacenter string,  storage string, vlan string, cpu int, mem int, ip string, host string, image string, pool string) (error) {
  mac, err := networks.Ipresource(ip)
  if err != nil {
    return err
  }

	/*create a vm*/
	vcpu := cpu
	vmem := mem * 1024 * 1024

	//create a uuid
	u := utils.Createuuid()

	osinfo, err := osimage.Getosinfobyosname(image, storage)
	if err != nil {
	  return err
  }

	//create baseimage
	imge_name, err := cephcommon.RbdClone(u, osinfo.Cephblockdevice, osinfo.Snapimage, pool)
	if err != nil {
	 return err
  }

	f, err := osimage.Xml(datacenter, storage, vlan,  vcpu, vmem, u, mac, imge_name, image, pool)
	if err != nil {
	  cephcommon.Rm_image(u)
	  return err
  }

	err = libvirtd.DefineVm(f, host)
	if err != nil {
	  cephcommon.Rm_image(u)
	  return err
  }

  err = Updatehost(host, cpu, mem)
  if  err != nil {
    cephcommon.Rm_image(u)
    libvirtd.Undefine(host, u)
    return err
  }
	newvm, err := savevm(datacenter, storage, u, cpu, mem, f, ip, host, image)
	if err != nil {
    cephcommon.Rm_image(u)
    libvirtd.Undefine(host, u)
    Freehost(host, cpu, mem)
	  return err
  }

  err = networks.Updateipstatus(ip, 1)
  if err != nil {
    cephcommon.Rm_image(u)
    libvirtd.Undefine(host, u)
    Freehost(host, cpu, mem)
    deletevmbyid(newvm)
    return  err
  }

	return nil
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

func base(vmid string, vmip string) string {
  b := []byte(vmid + "," + vmip)
  encodeString := base64.URLEncoding.EncodeToString(b)
  return encodeString
}

func allvm(obj []Vms) []map[string]interface{}  {
  var mapc []map[string]interface{}

  for _, v := range obj {
    c := make(map[string]interface{})
    m := reflect.TypeOf(v)
    n := reflect.ValueOf(v)
    for i := 0; i < m.NumField(); i++ {
      c[m.Field(i).Name] = n.Field(i).Interface()
    }

    vdisk, err := vdisk.Getdiskbyvm(v.Ip)
    if err != nil {
      return nil
    }
    c["disk"] = vdisk

    vncid := base(v.Uuid, v.Host)
    c["vncid"] = vncid
    mapc = append(mapc, c)
  }
  return mapc
}

func VmList() ([]map[string]interface{}, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
	var v []Vms
	dbs.Table("vms").Order("create_time desc").Select([]string{"uuid", "name", "cpu", "mem", "owner", "comment", "status", "storage", "datacenter", "exist", "ip" , "host", "os"}).Scan(&v)

	return allvm(v), nil
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
	dbs.Order("cpu").Find(&f)
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


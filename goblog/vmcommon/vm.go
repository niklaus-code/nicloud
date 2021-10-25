package vmcommon

import (
  "errors"
  "fmt"
  _ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
  uuid "github.com/satori/go.uuid"
  "goblog/ceph"
  "goblog/dbs"
  "goblog/libvirtd"
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
}

func GetVmByUuid(uuid string) *Vms {
  db, err := db.NicloudDb()
  v := &Vms{}
  if err != nil {
    return nil
  }
  db.Where("uuid = ?", uuid).First(v)
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


type Error struct {
	Code    int16
	Message string
}

func (err Error) Error() string {
	return fmt.Sprintf("vm is running, con't delete")
}

func Delete(uuid string, ip string, host string) ([]*Vms, error) {
	vmstat, err := VmStatus(uuid, host)
	if err != nil {
		return nil, err
	}

	if vmstat == "运行" {
		return nil, Error{
			Code:    501,
			Message: "vm is running, con't delete",
		}
	}

  db, err := db.NicloudDb()
  if err != nil {
    return  nil, err
  }

	//undefine vm
	conn, err := libvirtd.Libvirtconn(host)
	if err != nil {
		return nil, Error{
			Code:    502,
			Message: err.Error(),
		}
	}

	vm, err1 := conn.LookupDomainByUUIDString(uuid)
	if err1 != nil {
		return nil, err1
	}
	vm.Undefine()
	err = ceph.Rm_image(uuid)
  if err != nil {
    return nil, err
  }

	db.Model(&Vms{}).Where("uuid=?", uuid).Delete(&Vms{})
	db.Model(&vm_networks{}).Where("ipv4=?", ip).Update("status", 0)
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

  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
	var v = &Vms{}
	db.Where("uuid = ?", uuid).First(&v)

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

  db, err3 := db.NicloudDb()
  if err3 != nil {
    return nil, err3
  }
	var v = &Vms{}
	db.Where("uuid = ?", uuid).First(&v)

	s, err4 := VmStatus(uuid, host)
	v.Status = s
	if err4 != nil {
		return nil, err4
	}

	return v, err4
}

func Createuuid() string {
  /*create uuid*/
	u := uuid.NewV4().String()
	return u
}

func savevm(uuid string, cpu int, mem int, vmxml string, ip string, host string, image string) (bool, error) {
  /*save config to db*/
  db, err := db.NicloudDb()
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
	}
	err1 := db.Create(*vm)
	if err1.Error != nil {
	    return false, err1.Error
  }

	//return bool
	res := db.NewRecord(&vm)
	return res, err1.Error
}

func Ipresource(ip string, mac string) bool {
  db, err := db.NicloudDb()
  if err != nil {
    return false
  }
  var ipnet []*vm_networks
  db.Where("ipv4=?", ip).Where("macaddr=?", mac).Find(&ipnet)
  for _, v := range ipnet {
    if v.Status == 0 {
      return false
    }
  }
  return true
}

func MigrateVm(uuid string, migrate_host string) error {
  vm := GetVmByUuid(uuid)
  err := libvirtd.DefineVm(vm.Vmxml, migrate_host)
  if err != nil {
    return err
  }
  db, err := db.NicloudDb()
  if err != nil {
    return err
  }
  err1 := db.Model(&Vms{}).Where("uuid=?", uuid).Update("host", migrate_host)
  if err1.Error != nil {
      return err1.Error
  }

  _, err = Delete(uuid, vm.Ip, vm.Host)
  if err != nil {
    return err
  }

  return err
}

func Create(cpu int, mem int, ip string, mac string, host string, image string) (bool, error) {
  if Ipresource(ip, mac) {
    return false, nil
  }

	/*create a vm*/
	vcpu := cpu
	vmem := mem * 1024 * 1024

	//create a uuid
	u := Createuuid()

	//create baseimage
	imge_name, err := ceph.RbdClone(u)
	if err != nil {
	 return false, err
  }

	f, err := ceph.Xml(vcpu, vmem, u, mac, imge_name, image)

	err = libvirtd.DefineVm(f, host)
	if err != nil {
	  return false, err
  }

  hosterr := Updatehost(host, cpu, mem)
  if hosterr == false {
    return hosterr, nil
  }

	svm, err := savevm(u, cpu, mem, f, ip, host, image)
	if err != nil {
	  return svm, err
  }

  _, err = updateipstatus(ip)
  if err != nil {
    return false, err
  }

	return true, err
}

func updateipstatus(ipv4 string) (bool, error) {
  db, err := db.NicloudDb()
  if err != nil {
    return false, err
  }
	db.Model(&vm_networks{}).Where("ipv4=?", ipv4).Update("status", 1)
	return true, nil
}

func VmList(host string) []*Vms {
  db, err := db.NicloudDb()
  if err != nil {
    return nil
  }
	var v []*Vms
	db.Where("exist=1").Find(&v)

	return v
}

type vm_networks struct {
	Ipv4    string
	Macaddr string
	Status  int8
}

func IPlist() []*vm_networks {
  db, err := db.NicloudDb()
  if err != nil {
    return nil
  }
	var ip []*vm_networks
	db.Where("status=0").Find(&ip)

	return ip
}

type Vm_flavors struct {
	Cpu int
	Mem int
}

func Flavor() ([]*Vm_flavors, error) {
  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
	var f []*Vm_flavors
	db.Find(&f)
	return f, nil
}

func SearchVm(c string) ([]*Vms, error) {
  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var v []*Vms
  i := fmt.Sprintf("ip like %s", "'"+c+"%'")
  db.Where(i).Find(&v)

  return v, nil
}

type Vms_os struct {
  Osname string
  Snapimage string
}

func GetImages() ([]*Vms_os, error) {
  db, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var v []*Vms_os
  db.Find(&v)
  return v, nil
}

func Updatecomments(uuid string, comment string) (bool, error) {
  db, err := db.NicloudDb()
  if err != nil {
    return true, err
  }
  db.Model(&Vms{}).Where("uuid=?", uuid).Update("comment", comment)
  return true, nil
}


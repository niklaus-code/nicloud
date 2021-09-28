package vmcommon

import (
  "errors"
  "fmt"
  "goblog/ceph"
  "time"

  "github.com/ceph/go-ceph/rbd"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
  uuid "github.com/satori/go.uuid"
  libvirt "libvirt.org/libvirt-go"
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
	Status      interface{}
	Exist       int
	Ip          string
	Host        string
	Checkout    bool
}

func (v Vms) Error(info string) error {
	errorinfo := fmt.Sprintf("%s", info)
	return errors.New(errorinfo)
}

func vmdb() *gorm.DB {
	db, errDb := gorm.Open("mysql", "modis:modis@(127.0.0.1:3306)/nicloud?parseTime=true")
	if errDb != nil {
		return  nil
	}
	return db
}

type Vm_xmls struct {
	Ostype string
	Osxml  string
}

func libvirtconn(host string) (*libvirt.Connect, error) {
	conn, err := libvirt.NewConnect(fmt.Sprintf("qemu+ssh://%s/system", host))
	if err != nil {
		return nil, err
	}
	return conn, err
}

func VmStatus(uuid string, host string) (string, error) {
	conn, err := libvirtconn(host)

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

	return Vmstate[state], err1
}

var Vmstate = map[libvirt.DomainState]string{
	1: "运行",
	5: "关机",
	2: "deleted",
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

	db := vmdb()

	//undefine vm
	conn, err := libvirtconn(host)
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

func Shutdown(uuid string, host string) (*Vms, error) {
	/*start vm*/
	conn, err := libvirtconn(host)
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

	db := vmdb()
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
	conn, connerr := libvirtconn(host)
	if connerr != nil {
		return nil, connerr
	}
	vm, err := conn.LookupDomainByUUIDString(uuid)

	if err != nil {
		return nil, err
	}

	err1 := vm.Create()
	if err1 != nil {
		return nil, err1
	}

	db := vmdb()
	var v = &Vms{}
	db.Where("uuid = ?", uuid).First(&v)

	s, err2 := VmStatus(uuid, host)
	v.Status = s
	if err2 != nil {
		return nil, err2
	}
	return v, err2
}

func Createuuid() string {
  /*create uuid*/
	u := uuid.NewV4().String()
	return u
}

func savevm(uuid string, cpu int, mem int, vmxml string, ip string, host string) bool {
  /*save config to db*/
	db := vmdb()
	vm := &Vms{
		Uuid:        uuid,
		Name:        uuid,
		Cpu:         cpu,
		Mem:         mem,
		Vmxml:       vmxml,
		Create_time: time.Now(),
		Status:      1,
		Exist:       1,
		Ip:          ip,
		Host:        host,
		Owner:       "Niklaus",
	}
	db.Create(vm)

	//return bool
	res := db.NewRecord(&vm)
	return res
}

func Ipresource(ip string, mac string) bool {
  db := vmdb()
  var ipnet []*vm_networks
  db.Where("ipv4=?", ip).Where("macaddr=?", mac).Find(&ipnet)
  for _, v := range ipnet {
    if v.Status == 0 {
      return false
    }
  }
  return true
}


func Create(cpu int, mem int, ip string, mac string, host string) (bool, error) {
  if Ipresource(ip, mac) {
    return false, nil
  }

	/*create a vm*/
	vcpu := cpu
	vmem := mem * 1024 * 1024

	//create a uuid
	u := Createuuid()

	//create baseimage
	imge_name, err := RbdClone(u)
	if err != nil {
	 return false, err
  }

	f, err := ceph.Xml(vcpu, vmem, u, mac, imge_name)

	conn, connerr := libvirtconn(host)
	if connerr != nil {
		return false, connerr
	}

	_, err = conn.DomainDefineXML(f)
	if err != nil {
		return false, err
	}
	_, err = updateipstatus(ip)
	if err != nil {
		return false, err
	}

	savevm(u, cpu, mem, f, ip, host)
	if err != nil {
		return false, err
	}

	return true, err
}

func updateipstatus(ipv4 string) (bool, error) {
	db := vmdb()
	db.Model(&vm_networks{}).Where("ipv4=?", ipv4).Update("status", 1)
	return true, nil
}

func VmList(host string) []*Vms {
	db := vmdb()
	var v []*Vms
	db.Where("exist=1").Find(&v)
	//for _, e := range v {
	//	s, err := VmStatus(e.Uuid, e.Host)
	//	e.Checkout = false
	//	if err != nil {
	//		e.Status = err.Error()
	//	} else {
	//		e.Status = s
	//	}
	//}
	return v
}

type vm_networks struct {
	Ipv4    string
	Macaddr string
	Status  int8
}

func IPlist() []*vm_networks {
	db := vmdb()
	var ip []*vm_networks
	db.Where("status=0").Find(&ip)

	return ip
}

type Vm_hosts struct {
	Ipv4        string
	Mem         int8
	Cpu         int8
	Max_vms     int8
	Created_vms int8
	Status      int8
}

func Hosts() []*Vm_hosts {
	db := vmdb()
	var hosts []*Vm_hosts
	db.Where("status=0").Find(&hosts)
	return hosts
}

type Vm_flavors struct {
	Cpu int
	Mem int
}

func Flavor() ([]*Vm_flavors, error) {
	db := vmdb()
	var f []*Vm_flavors
	db.Find(&f)
	return f, nil
}

func RbdClone(id string) (string, error) {

	conn, err := ceph.CephConn()
	if err != nil {
		return "", err
	}

	ioctx, _ := conn.OpenIOContext("vm")
	img := rbd.GetImage(ioctx, "0000_demo_centos7")
	_, e := img.Clone("20210806_095737", ioctx, id, rbd.RbdFeatureLayering, 0)

	if e != nil {
		return "", e
	}
	return id, nil
}

func SearchVm(c string) []*Vms {
  db := vmdb()
  var v []*Vms
  i := fmt.Sprintf("ip like %s", "'"+c+"%'")
  db.Where(i).Find(&v)
  //for _, e := range v {
  //  s, err := VmStatus(e.Uuid, e.Host)
  //  if err != nil {
  //    e.Status = err.Error()
  //  } else {
  //    e.Status = s
  //  }
  //}
  return v
}

package vm

import (
  "encoding/base64"
  "errors"
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
  "nicloud/cephcommon"
  c "nicloud/config"
  "nicloud/dbs"
  "nicloud/libvirtd"
  "nicloud/networks"
  "nicloud/osimage"
  "nicloud/users"
  "nicloud/utils"
  vdisk "nicloud/vdisk"
  vmerror "nicloud/vmerror"
  "reflect"
  "time"
)
var ceph cephcommon.Vms_Ceph
type Vms struct {
	Uuid        string `gorm:"primary_key;"`
	Name        string
	Cpu         uint `json:"cpu" validate:"gt=0"`
	Mem         uint `json:"mem" validate:"gt=0"`
	Create_time time.Time `json:"Create_time"`
	Owner       int  `json:"Owner" validate:"required"`
	Comment     string
	Vmxml       string `gorm:"size:65535"`
  Status      string
	Exist       int
	Ip          string  `gorm:"unique" json:"ip"`
	Host        string  `json:"host" validate:"min=8,max=15"`
	Os          int  `json:"os" validate:"required"`
	Datacenter  string  `json:"datacenter" validate:"required"`
	Storage     string  `json:"storage" validate:"required"`
}

func updatexmlbyuuid(xml string, uuid string, vcpu uint, vmem uint) (*gorm.DB, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  tx := dbs.Begin()
  err = tx.Model(&Vms{}).Where("uuid=?", uuid).Update("vmxml", xml).Update("cpu", vcpu).Update("mem", vmem).Error
  if err != nil {
    tx.Rollback()
    return nil, err
  }

  return tx, nil
}

func Changeconfig(uuid string, host string, vcpu uint, oldcpu uint,  vmem uint, oldmem uint,  vmhost string) error {
  mem_kb := vmem * 1024 * 1024
  s, err := VmStatus(uuid, host)
  if s != "关机" {
    return vmerror.Error{Message: "云主机需要关机状态"}
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

  h := Vm_hosts{}

  tx_updatehost, err := h.UpdateCpuMem(host, int(vcpu-oldcpu), int(vmem-oldmem))
  if err != nil {
    return err
  }

  xmlstr, err := libvirtd.UpdateCpuMem(v.Vmxml, vcpu, mem_kb)
  if err != nil {
    tx_updatehost.Rollback()
    return err
  }

  tx_updatexml, err := updatexmlbyuuid(xmlstr, uuid, vcpu, vmem)
  if err != nil {
    return err
  }

  err = libvirtd.DefineVm(xmlstr, vmhost)
  if err != nil {
    tx_updatehost.Rollback()
    tx_updatexml.Rollback()
    return err
  }
  tx_updatexml.Commit()
  tx_updatehost.Commit()
  return nil
}

func GetVmByUuid(uuid string) (*Vms, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, nil
  }

  v := &Vms{}
  dbs.Where("uuid = ?", uuid).First(v)
  return v, nil
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
	vm, err := libvirtd.GetDomain(host, uuid)
	if err != nil {
		return "未发现云主机", err
	}

	state, _, err1 := vm.GetState()

	if err1 != nil {
		return "未发现云主机", err1
	}
    vm.Free()

	return libvirtd.Vmstate[state], nil
}

type Vms_archives struct {
  Uuid string
  Create_time time.Time
  Archive_time time.Time
  Owner int
  Comment string
  Vmxml string `gorm:"size:65535"`
  Ip string
  Host string
  Os int
  Datacenter string
  Storage string
}

func (a Vms_archives)savevmarchives(uuid string, create_time time.Time, owner int, comment string, vmxml string, ip string, host string, os int, datacenter string, storage string) (*gorm.DB, error) {
  archives := Vms_archives{
    Uuid: uuid,
    Create_time: create_time,
    Archive_time: time.Now(),
    Owner: owner,
    Comment: comment,
    Vmxml: vmxml,
    Ip: ip,
    Host: host,
    Os: os,
    Datacenter: datacenter,
    Storage: storage,
  }
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  tx := dbs.Begin()
  err = tx.Create(&archives).Error
  if err != nil {
    tx.Rollback()
    return nil, err
  }
  return tx, nil
}

func Getvmarchivepagenumber(userid int) (int, int, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return 0, 0, err
  }

  var v []Vms_archives
  user, err := users.GetUserByUserID(userid)
  if err != nil {
    return 0, 0, err
  }

  role, err := users.GetRoleByRoleId(user.Role)
  if err != nil {
    return 0, 0, err
  }

  if role.Rolename == "admin" {
    dbs.Table("vms_archives").Order("archive_time desc").Find(&v)
  } else {
    dbs.Table("vms_archives").Order("archive_time desc").Where("owner=?", userid).Order("create_time desc").Find(&v)
  }
  remainder := len(v)%offset
  var pagenumber int
  if remainder > 0 {
    pagenumber = len(v)/offset + 1
  } else {
    pagenumber = len(v)/offset
  }
  return pagenumber, len(v), nil
}

func Mapvmarchive(obj []Vms_archives) []map[string]interface{}  {
  var mapc []map[string]interface{}

  for _, v := range obj {
    c := make(map[string]interface{})
    m := reflect.TypeOf(v)
    n := reflect.ValueOf(v)
    for i := 0; i < m.NumField(); i++ {
      c[m.Field(i).Name] = n.Field(i).Interface()
    }

    o := osimage.Vms_os{}
    osinfo, err := o.GetOsInfoById(v.Storage, v.Os)

    if err != nil {
      c["osname"] = nil
    } else {
      c["osname"] = osinfo.Osname
    }

    owner, err := users.GetUserByUserID(v.Owner)
    c["owner"] = owner.Username
    if err != nil {
      c["owner"] = nil
    }


    mapc = append(mapc, c)
  }
  return mapc
}

func (arch Vms_archives)GetVmArchive(startpage int) ([]map[string]interface{}, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  a := []Vms_archives{}
  dberr := dbs.Order("archive_time desc").Order("create_time desc").Limit(offset).Offset((startpage-1)*offset).Find(&a)
  if dberr.Error != nil {
    return nil, err
  }

  return Mapvmarchive(a), nil
}

func Delete(uuid string, storage string) (error) {
  o := osimage.Vms_os{}
  check_cephblock_relate_os, err := o.CheckOsbyUuid(uuid)
  if err != nil {
    return  err
  }

  if check_cephblock_relate_os {
    return vmerror.Error{Message: "系统快照关联基础镜像, 无法删除"}
  }

  vminfo, err := GetVmByUuid(uuid)
  if err != nil {
    return err
  }
  host := vminfo.Host

  storageinfo, err := ceph.Cephinfobyuuid(storage)
  if err != nil {
    return err
  }

	vmstat, err := VmStatus(uuid, host)
	if err != nil {
		return err
	}

	if vmstat == "运行" {
		return vmerror.Error{
			Message: "虚拟机正在运行，无法删除",
		}
	}

  if vmstat == "暂停" {
    return vmerror.Error{
      Message: "虚拟机处于暂停状态，无法删除",
    }
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  dblist := []*gorm.DB{}

  c := cephcommon.Vms_Ceph{}
	delimgid, err := c.Rm_image(uuid, storageinfo.Pool)
  if err != nil {
    return vmerror.Error{"删除块设备错误"}
  }

  tx := dbs.Begin()
	tx_delvm := dbs.Model(&Vms{}).Where("uuid=?", uuid).Delete(&Vms{}).Error
	if tx_delvm != nil {
    db.Tx_rollback(append(dblist, tx))
	  c.RenameBlock(delimgid, uuid)
  }

  tx_updateipstatus, err := networks.Updateipstatus(vminfo.Ip, 0)
	if err != nil {
    db.Tx_rollback(append(dblist, tx))
    c.RenameBlock(delimgid, uuid)
	  return err
  }

  archives := Vms_archives{}
  tx_savearchives, err := archives.savevmarchives(delimgid, vminfo.Create_time, vminfo.Owner, vminfo.Comment, vminfo.Vmxml, vminfo.Ip, vminfo.Host, vminfo.Os, vminfo.Datacenter, vminfo.Storage)
  if err != nil {
    db.Tx_rollback(append(dblist, tx, tx_updateipstatus))
    c.RenameBlock(delimgid, uuid)
    return err
  }

  h := Vm_hosts{}
  tx_UpdateCpuMem, err := h.UpdateCpuMem(vminfo.Host, int(0-vminfo.Cpu), int(0-vminfo.Mem))
  if err != nil {
    db.Tx_rollback(append(dblist, tx, tx_updateipstatus, tx_savearchives))
    ceph.RenameBlock(delimgid, uuid)
    return err
  }

  tx_updatevdisk, err := vdisk.Updatevdiskbydelvm(vminfo.Datacenter, vminfo.Storage, vminfo.Ip)
  if err != nil {
    db.Tx_rollback(append(dblist, tx, tx_updateipstatus, tx_savearchives, tx_UpdateCpuMem))
    ceph.RenameBlock(delimgid, uuid)
    return err
  }

  err = libvirtd.Undefine(host, uuid)
  if err != nil {
    db.Tx_rollback(append(dblist, tx, tx_updateipstatus, tx_savearchives, tx_UpdateCpuMem, tx_updatevdisk))
    ceph.RenameBlock(delimgid, uuid)
    return err
  }

  db.Tx_commot(append(dblist, tx,tx_updateipstatus, tx_savearchives, tx_UpdateCpuMem,tx_updatevdisk))
	return nil
}

func PauseVm(uuid string, host string) error {
  vm, err := libvirtd.GetDomain(host, uuid)
  if err != nil {
    return err
  }

  err = vm.Suspend()
  if err != nil {
    return err
  }
  return nil
}

/*
enum virDomainRebootFlagValues {
  VIR_DOMAIN_REBOOT_DEFAULT 	= 	0 (0x0) hypervisor choice
  VIR_DOMAIN_REBOOT_ACPI_POWER_BTN 	= 	1 (0x1; 1 << 0) Send ACPI event
  VIR_DOMAIN_REBOOT_GUEST_AGENT 	= 	2 (0x2; 1 << 1) Use guest agent
  VIR_DOMAIN_REBOOT_INITCTL 	= 	4 (0x4; 1 << 2) Use initctl
  VIR_DOMAIN_REBOOT_SIGNAL 	= 	8 (0x8; 1 << 3) Send a signal
  VIR_DOMAIN_REBOOT_PARAVIRT 	= 	16 (0x10; 1 << 4) Use paravirt guest control
}
*/

func Reboot(uuid string, host string) error {
  /*start vm*/
  vm, err4 := libvirtd.GetDomain(host, uuid)
  if err4 != nil {
    return err4
  }

  err1 := vm.Reboot(0)
  if err1 != nil {
    return err1
  }

  return nil
}

func Shutdown(uuid string, host string) error {
  /*start vm*/
  vm, err4 := libvirtd.GetDomain(host, uuid)
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
	vm, err4 := libvirtd.GetDomain(host, uuid)
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
	vm, err := libvirtd.GetDomain(host, uuid)

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

func savevm(datacenter string, cephname string, uuid string, cpu uint, mem uint, vmxml string, ip string, host string, image int, owner int, comment string) (*gorm.DB, error) {
  /*save config to db*/
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
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
		Owner:       owner,
		Os:          image,
		Datacenter: datacenter,
		Storage: cephname,
		Comment: comment,
	}

	tx := dbs.Begin()
	err = tx.Create(vm).Error
	if err != nil {
	  tx.Rollback()
	  return tx, err
  }
	return tx, err
}

func MigrateVm(uuid string, migrate_host string) error {
  vm, err := GetVmByUuid(uuid)
  if err != nil {
    return err
  }

  h := Vm_hosts{}
  check := h.checkcpumem(migrate_host, vm.Cpu, vm.Mem)
  if check != nil {
    return check
  }

  s, err := VmStatus(uuid, vm.Host)
  if s == "开机" {
    return vmerror.Error{Message: "云主机需要关机状态"}
  }

  err = libvirtd.DefineVm(vm.Vmxml, migrate_host)
  if err != nil {
    return err
  }
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  dblist := []*gorm.DB{}
  tx := dbs.Begin()
  err = tx.Model(&Vms{}).Where("uuid=?", uuid).Update("host", migrate_host).Error
  if err != nil {
    libvirtd.Undefine(migrate_host, uuid)
    tx.Rollback()
    return err
  }

  tx_freecpumem, err := h.UpdateCpuMem(vm.Host, int(0-vm.Cpu), int(0-vm.Mem))
  if err != nil {
    libvirtd.Undefine(migrate_host, uuid)
    tx.Rollback()
    return err
  }

  tx_updatehost, err := h.Createvmonhost(migrate_host, vm.Cpu, vm.Mem)
  if  err != nil {
    libvirtd.Undefine(migrate_host, uuid)
    db.Tx_rollback(append(dblist, tx, tx_freecpumem))
    return err
  }

  err = libvirtd.Undefine(vm.Host, vm.Uuid)
  if err != nil {
    db.Tx_rollback(append(dblist, tx, tx_updatehost, tx_freecpumem))
    libvirtd.Undefine(migrate_host, uuid)
    return err
  }

  db.Tx_commot(append(dblist, tx, tx_freecpumem, tx_updatehost))
  return err
}

func MigrateVmlive(uuid string,  desthost string) error {
  vm, err := GetVmByUuid(uuid)
  if err != nil {
    return err
  }

  h := Vm_hosts{}
  checkhost := h.checkcpumem(desthost, vm.Cpu, vm.Mem)
  if checkhost != nil {
    return checkhost
  }

  s, err := VmStatus(uuid, vm.Host)
  if s != "运行" && s != "暂停" {
    return vmerror.Error{Message: "云主机需要开机或者暂停状态"}
  }

  //动态迁移生存在内存里面,重启之后消失
  migratelive := libvirtd.Migratevmlive(uuid, vm.Host, desthost)
  if migratelive != nil {
    return migratelive
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  dblist := []*gorm.DB{}
  tx := dbs.Begin()
  err = dbs.Model(&Vms{}).Where("uuid=?", uuid).Update("host", desthost).Error
  if err != nil {
    db.Tx_rollback(append(dblist, tx))
    return err
  }

  //update desination  host
  tx_updatehost, err := h.Createvmonhost(desthost, vm.Cpu, vm.Mem)
  if err != nil {
    db.Tx_rollback(append(dblist, tx_updatehost))
    return err
  }

  //update src host
  tx_freecpumem, err := h.UpdateCpuMem(vm.Host, int(0-vm.Cpu), int(0-vm.Mem))
  if err != nil {
    db.Tx_rollback(append(dblist, tx, tx_updatehost))
    return err
  }

  err = libvirtd.Undefine(vm.Host, uuid)
  if err != nil {
    db.Tx_rollback(append(dblist, tx, tx_updatehost, tx_freecpumem))
    return err
  }

  //防止重启之后消失，新建
  err = libvirtd.DefineVm(vm.Vmxml, desthost)
  if err != nil {
    db.Tx_rollback(append(dblist, tx, tx_updatehost, tx_freecpumem))
    libvirtd.DefineVm(vm.Vmxml, vm.Host)
    return err
  }

  db.Tx_commot(append(dblist, tx, tx_updatehost, tx_freecpumem))
  return err
}

func (v Vms)Create (datacenter string,  storage string, vlan string, cpu uint, mem uint, ip string, host string, osid int, owner int, comment string) (error) {
  h := Vm_hosts{}
  checkresoures := h.checkcpumem(host, cpu, mem)
  if checkresoures != nil {
    return checkresoures
  }

  mac, err := networks.Ipresource(ip)
  if err != nil {
    return err
  }

	/*create a vm*/
	vcpu := cpu
	vmem := mem * 1024 * 1024

	//create a uuid
	u := utils.Createuuid()

	o := osimage.Vms_os{}
	osinfo, err := o.GetOsInfoById(storage, osid)
	if err != nil {
	  return err
  }
  t := osimage.Vms_os_tags{}
  ostag, err := t.GetostagByid(osinfo.Tag)
  if err != nil {
    return err
  }

  storageinfo, err := ceph.Cephinfobyuuid(storage)
  if err != nil {
    return err
  }

	//create baseimage
  c := cephcommon.Vms_Ceph{}
	imge_name, err := c.RbdClone(u, osinfo.Cephblockdevice, osinfo.Snapimage, storageinfo.Pool)
	if err != nil {
	 return err
  }

	f, err := libvirtd.CreateVmXml(datacenter, storage, vlan, vcpu, vmem, u, mac, imge_name, osid, storageinfo.Pool, ostag.Tag)
	if err != nil {
	  c.Rm_image(u, storageinfo.Pool)
	  return err
  }

	err = libvirtd.DefineVm(f, host)
	if err != nil {
	  c.Rm_image(u, storageinfo.Pool)
	  return err
  }

  dblist := []*gorm.DB{}

  tx_updatehost, err := h.Createvmonhost(host, cpu, mem)
  if  err != nil {
    c.Rm_image(u, storageinfo.Pool)
    libvirtd.Undefine(host, u)
    return err
  }

  tx_savevm, err := savevm(datacenter, storage, u, cpu, mem, f, ip, host, osid, owner, comment)
	if err != nil {
    c.Rm_image(u, storageinfo.Pool)
    libvirtd.Undefine(host, u)
    db.Tx_rollback(append(dblist, tx_savevm))
	  return err
  }

  tx_updateipstatus, err := networks.Updateipstatus(ip, 1)
  if err != nil {
    c.Rm_image(u, storageinfo.Pool)
    libvirtd.Undefine(host, u)
    db.Tx_rollback(append(dblist, tx_savevm, tx_updatehost))
    return  err
  }

  db.Tx_commot(append(dblist, tx_updatehost, tx_savevm, tx_updateipstatus))

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
    return "", vmerror.Error{Message: "未找到数据"}
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

  for k, v := range obj {
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

    o := osimage.Vms_os{}
    osinfo, err := o.GetOsInfoById(v.Storage, v.Os)
    if err != nil {
      return nil
    }
    c["osname"] = osinfo.Osname
    c["Create_time"] = obj[k].Create_time.Format("2006-01-02 15:04:05")

    vncid := base(v.Uuid, v.Host)
    c["vncid"] = vncid
    owner, err := users.GetUserByUserID(v.Owner)
    if err != nil {
      c["Owner"] = nil
    } else {
    	c["Owner"] = owner.Username
	  }
    mapc = append(mapc, c)
  }
  return mapc
}

var (
	config, _ = c.Exportconfig()
	offset = config.Page.Offset
)

func Getpagenumber(userid int) (int, int, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return 0, 0, err
  }

  var v []Vms
  user, err := users.GetUserByUserID(userid)
  if err != nil {
    return 0, 0, err
  }

  role, err := users.GetRoleByRoleId(user.Role)
  if err != nil {
    return 0, 0, err
  }

  if role.Rolename == "admin" {
    dbs.Table("vms").Order("create_time desc").Select([]string{"uuid", "name", "cpu", "mem", "owner", "comment", "status", "storage", "datacenter", "exist", "ip", "host", "os"}).Scan(&v)
  } else {
    dbs.Table("vms").Order("create_time desc").Where("owner=?", userid).Order("create_time desc").Select([]string{"uuid", "name", "cpu", "mem", "owner", "comment", "status", "storage", "datacenter", "exist", "ip", "host", "os"}).Scan(&v)
  }
  remainder := len(v)%offset
  var pagenumber int
  if remainder > 0 {
    pagenumber = len(v)/offset + 1
  } else {
    pagenumber = len(v)/offset
  }
  return pagenumber, len(v), nil
}

//判断奇数偶数
var odd = 0
var startpage = 1
var order = "desc"

func VmList(userid int, start int, item string) ([]map[string]interface{}, error) {
  odd = odd + 1
  var order string
  if odd%2 == 0 || item == "create_time" {
    order = "desc"
  } else {
    order = "asc"
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
	var v []Vms

  user, err := users.GetUserByUserID(userid)
  if err != nil {
    return nil, err
  }

  role, err := users.GetRoleByRoleId(user.Role)
  if err != nil {
    return nil, err
  }

  if role.Rolename == "admin" {
    dbs.Raw(fmt.Sprintf("select * from (select * from vms  order by create_time desc limit %d offset %d) v order by %s %s", offset, (start-1)*offset,item, order)).Scan(&v)
  } else {
    dbs.Table("vms").Order(fmt.Sprintf("%s %s", item, order)).Where("owner=?", userid).Order("create_time desc").Select([]string{"uuid", "name", "cpu", "mem", "owner", "comment", "status", "storage", "datacenter", "exist", "ip", "host", "os"}).Limit(offset).Offset((start-1)*offset).Scan(&v)
  }

  return allvm(v), nil
}

type Vm_flavors struct {
  Cpu int `gorm:"index;unique_index:name_d;unique" json:"Cpu" validate:"gt=0"`
  Mem int `gorm:"unique_index:name_d;unique" json:"Mem" validate:"gt=0"`
}

func (f Vm_flavors)Createflavor(flavors *Vm_flavors) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  err = dbs.Create(&flavors).Error
  if err != nil {
    return err
  }
  return nil
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

func SearchVm(c string) ([]map[string]interface{}, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var v []Vms
  i := fmt.Sprintf("uuid like %s or ip like %s or comment like %s or host like %s", "'%"+c+"%'", "'%"+c+"%'", "'%"+c+"%'", "'%"+c+"%'")
  dbs.Where(i).Order("create_time desc").Find(&v)
  return allvm(v), nil
}

func Updatecomments(uuid string, comment string) (bool, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return true, err
  }
  dbs.Model(&Vms{}).Where("uuid=?", uuid).Update("comment", comment)
  return true, nil
}

func (v Vms)Rebuildimg(osid int, storage string, datacenter string, old_uuid string, host string) error {
  vmstat, err := VmStatus(old_uuid, host)
  if err != nil {
    return err
  }

  if vmstat != "关机" {
    return vmerror.Error{
      Message: "虚拟机正在运行，无法重置",
    }
  }

  o := osimage.Vms_os{}
  osinfo, err := o.GetOsInfoById(storage, osid)
  if err != nil {
    return err
  }

  storageinfo, err := ceph.Cephinfobyuuid(storage)
  if err != nil {
    return err
  }

  uuid := utils.Createuuid()
  c := cephcommon.Vms_Ceph{}
  err = c.Changename(uuid, osinfo.Cephblockdevice, osinfo.Snapimage, storageinfo.Pool, old_uuid)
  if err != nil {
    return err
  }

  updateos := v.Updataosbyuuid(old_uuid, osid)
  if updateos != nil {
    return updateos
  }
  return nil
}

func CreatSnap(vmid string, datacenter string, storage string, snapname string) error {

  storageinfo, err := ceph.Cephinfobyuuid(storage)
  if err != nil {
    return err
  }
  c := cephcommon.Vms_Ceph{}
  create := c.Createimgsnap(vmid, snapname, storageinfo.Pool)
  if create != nil {
    return create
  }

  s := cephcommon.Vms_snaps{
    Vm_uuid: vmid,
    Datacenter: datacenter,
    Storage: storage,
    Snap: snapname,
    Create_time: time.Now(),
    Status: true,
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb := dbs.Create(&s)
  if errdb.Error !=nil {
    return errdb.Error
  }
  return nil
}

func SaveSnapToImg(vmid string, datacenter string, storage string, snapname string, userid int) error {
  vminfo, err  := GetVmByUuid(vmid)
  if err != nil {
    return err
  }

  o := osimage.Vms_os{}
  osinfo, err := o.GetOsInfoById(storage, vminfo.Os)

  storageinfo, err := ceph.Cephinfobyuuid(storage)
  if err != nil {
    return err
  }

  c := cephcommon.Vms_Ceph{}
  snapprotect := c.SnapProtect(vmid, storageinfo.Pool, snapname)
  if snapprotect != nil {
    return snapprotect
  }

  os := osimage.Vms_os{}
  err = os.Add(datacenter, storage, snapname+"_"+osinfo.Osname, vmid, osinfo.Xml, 2, userid, snapname, osinfo.Tag)
  if err != nil {
    return err
  }

  return nil
}


func mapsnap(snap []cephcommon.Vms_snaps) []map[string]interface{}  {
  var mapc []map[string]interface{}

  for _, v := range snap {
    c := make(map[string]interface{})
    m := reflect.TypeOf(v)
    n := reflect.ValueOf(v)
    for i := 0; i < m.NumField(); i++ {
      c[m.Field(i).Name] = n.Field(i).Interface()
    }

    chilimages, err := ceph.ListChildernImages(v.Storage, v.Vm_uuid)
    if err != nil {
      c["chilimages"] = nil
    } else {
      c["chilimages"] = chilimages
    }
    mapc = append(mapc, c)
  }
  return mapc
}

func Getsnap(datacenter string, storage string, vmid string) ([]map[string]interface{} , error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  snap := []cephcommon.Vms_snaps{}
  errdb := dbs.Where("datacenter=? and storage= ? and vm_uuid =?", datacenter, storage, vmid).Find(&snap)
  if errdb.Error != nil {
    return nil, errdb.Error
  }
  return mapsnap(snap), nil
}

func RollbackSnap(vmid string, snapname string, datacenter string, storage string) error {
  c := cephcommon.Vms_Ceph{}
  storageinfo, err := ceph.Cephinfobyuuid(storage)
  if err != nil {
    return err
  }

  r := c.Rollback(vmid, snapname, storageinfo.Pool)
  if r != nil {
    return r
  }
  return nil
}

func DelSnap(vmid string, snapname string, datacenter string, storage string) error {
  c := cephcommon.Vms_Ceph{}
  storageinfo, err := ceph.Cephinfobyuuid(storage)
  if err != nil {
    return err
  }

  r := c.Delsnap(vmid, snapname, storageinfo.Pool)
  if r != nil {
    return r
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb := dbs.Model(Vms{}).Where("vm_uuid=? and snap=?", vmid, snapname).Delete(cephcommon.Vms_snaps{})
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}

func Checkuser(userid int) error{
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  var v []*Vms
  dbs.Where("owner=?", userid).Find(&v)

  if len(v) > 0 {
    return vmerror.Error{Message: "请先删除用户关联云主机"}
  }
  return nil
}

func GetVmbyOsId(osid int) (bool, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return false, err
  }

  var count int
  errdb := dbs.Model(&Vms{}).Where("os=?", osid).Count(&count)
  if errdb.Error != nil {
    return false, errdb.Error
  }

  if count > 0 {
    return false, err
  }
  return true, nil
}

func (v Vms)Updataosbyuuid(uuid string, osid int) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  errdb := dbs.Model(&Vms{}).Where("uuid=?", uuid).Update("os", osid)
  if errdb.Error != nil {
    return err
  }
  return nil
}

func (v Vms_archives)Delvmpermanent(storage string, uuid string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  c := cephcommon.Vms_Ceph{}
  delcephimg:= c.Delimgpermanent(storage, uuid)
  if delcephimg != nil {
    return delcephimg
  }

  errdb := dbs.Model(&Vms_archives{}).Where("uuid=?", uuid).Delete(&Vms_archives{})
  if errdb.Error != nil {
    return err
  }
  return nil
}

func (v Vms_archives)SearchVMArchives(content string) ([]*Vms_archives, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var vmarchives []*Vms_archives
  i := fmt.Sprintf("ip like %s or comment like %s", "'%"+content+"%'", "'%"+content+"%'")
  errdb := dbs.Where(i).Order("archive_time desc").Find(&vmarchives)

  if errdb.Error != nil {
    return nil, err
  }
  return vmarchives, nil
}

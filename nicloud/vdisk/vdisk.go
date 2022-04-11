package vdisk

import (
  "github.com/jinzhu/gorm"
  "nicloud/cephcommon"
  c "nicloud/config"
  db "nicloud/dbs"
  "nicloud/libvirtd"
  "nicloud/users"
  "nicloud/utils"
  "nicloud/vmerror"
  "reflect"
  "strings"
  "time"
)

type Vms_vdisks struct {
  Vdiskid string `gorm:"primary_key;"`
  Contain int `json:"contain" validate:"min=0,max=1024"`
  Diskname string
  Pool string `json:"pool" validate:"required"`
  Storage string  `json:"storage" validate:"required"`
  Datacenter string `json:"datacenter" validate:"required"`
  Vm_ip string
  User int `json:"user" validate:"required"`
  Exist int8
  Status int
  Comment string
  Createtime time.Time
}

func (* Vms_vdisks) Addcomment(vdiskid string, comment string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb := dbs.Model(Vms_vdisks{}).Where("vdiskid=?", vdiskid).Update("comment", comment)
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}

func Getdiskbyvm(vmip string) ([]*Vms_vdisks, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  c := []*Vms_vdisks{}
  dbs.Select("contain, diskname").Where("vm_ip=?", vmip).Find(&c)
  return c, nil
}

func UpdateMountvmstatus(datacenter string, storage string, vdiskid string, vmip string, diskname string) (*gorm.DB, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  tx := dbs.Begin()
  err = dbs.Model(Vms_vdisks{}).Where("datacenter=? and storage=? and vdiskid=?", datacenter, storage, vdiskid).Update(map[string]interface{}{"vm_ip": vmip, "status": 0, "diskname": diskname}).Error
  if err != nil {
    tx.Rollback()
    return nil, err
  }
  return tx, nil
}

func Updatevdiskbydelvm(datacenter string, storage string, vmip string) (*gorm.DB, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  tx := dbs.Begin()
  err = tx.Model(Vms_vdisks{}).Where("datacenter=? and storage=? and vm_ip=?", datacenter, storage, vmip).Update(map[string]interface{}{"vm_ip": "", "status": 1, "diskname": ""}).Error
  if err != nil {
    tx.Rollback()
    return nil, err
  }
  return tx, nil
}


func (d Vms_vdisks)Create(contain int, pool string, cephid string, datacenter string, userid int, comment string) error {
  ceph := cephcommon.Vms_Ceph{}
  vdiskid := utils.Createuuid()

  c := &Vms_vdisks{
    Vdiskid: vdiskid,
    Contain: contain,
    Pool: pool,
    Storage: cephid,
    Datacenter: datacenter,
    User: userid,
    Exist: 1,
    Status: 1,
    Comment: comment,
    Createtime: time.Now(),
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  tx := dbs.Begin()
  err = tx.Create(&c).Error
  if err != nil {
    tx.Rollback()
    return err
  }

  err = ceph.Createcephblock(vdiskid, contain, pool)
  if err != nil {
    tx.Rollback()
    return err
  }
  tx.Commit()
  return err
}

type Vms_vdisks_archives struct {
  Vdiskid string `gorm:"size:100"`
  Owner int
  Pool string
  Storage string
  Datacenter string
  Comment string
  Create_time time.Time
  Archive_time time.Time
}

func addiskachives(uuid string, pool string, storage string, datacenter string, ownerid int, comment string, create_time time.Time) (string, *gorm.DB, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return "", nil ,err
  }

  diskachi := &Vms_vdisks_archives{
    Vdiskid: uuid,
    Owner: ownerid,
    Pool: pool,
    Storage: storage,
    Datacenter: datacenter,
    Comment: comment,
    Create_time: create_time,
    Archive_time: time.Now(),
  }

  tx := dbs.Begin()
  err = dbs.Create(diskachi).Error
  if err != nil {
    tx.Rollback()
    return "", nil, err
  }
  return diskachi.Vdiskid, tx, nil
}

func deletedisk(uuid string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  errdb := dbs.Where("uuid=?").Delete(&Vms_vdisks_archives{})
  if errdb.Error != nil {
    return errdb.Error
  }
  return nil
}

func getdiskinfobyid(uuid string) (*Vms_vdisks, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  vdiskinfo := &Vms_vdisks{}
  errdb := dbs.Where("Vdiskid=?", uuid).First(vdiskinfo)
  if errdb.Error != nil {
    return nil, errdb.Error
  }
  return vdiskinfo, err
}

func Deletevdisk(uuid string, comment string) error {
  checkmount, err := Getdiskstatus(uuid)
  if err != nil {
    return err
  }

  if checkmount == 0 {
    return vmerror.Error{Message: "硬盘已挂载，请卸载后删除"}
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  vdiskinfo, err := getdiskinfobyid(uuid)
  if err != nil {
    return err
  }

  tx := dbs.Begin()
  err = tx.Where("vdiskid=?", uuid).Delete(Vms_vdisks{}).Error
  if err != nil {
    tx.Rollback()
    return err
  }

  c := cephcommon.Vms_Ceph{}
  delid, err := c.Rm_image(uuid, vdiskinfo.Pool)
  if err != nil {
    tx.Rollback()
    return err
  }

  _, _, err = addiskachives(delid, vdiskinfo.Pool, vdiskinfo.Storage, vdiskinfo.Datacenter, vdiskinfo.User, comment, vdiskinfo.Createtime)
  if err != nil {
    tx.Rollback()
    c.RenameBlock(delid, uuid)
    return err
  }
  return nil
}

func (d Vms_vdisks)Umountdisk(vmip string, storage string, datacenter string, vdiskid string, xml string, host string, vms interface{}) error {
  checkmount, err := checkmount(vdiskid)
  if err != nil {
    return err
  }
  if checkmount == 1 {
    return vmerror.Error{Message: "vdisk has been mouunted"}
  }

  c := cephcommon.Vms_Ceph{}
  storageinfo, err := c.Cephinfobyuuid(storage)
  if err != nil {
    return err
  }

  updatexml, err := libvirtd.RemoveDiskXml(xml, vdiskid, storageinfo.Pool)
  if err != nil {
    return err
  }

  tx_Umountvmstatus, err := Umountvmstatus(datacenter, storage, vdiskid)
  if err != nil {
    return err
  }

  dblist := []*gorm.DB{}
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  tx := dbs.Begin()
  err = tx.Model(vms).Where("ip=?", vmip).Update("vmxml", updatexml).Error
  if err != nil {
    db.Tx_rollback(append(dblist, tx_Umountvmstatus, tx))
    return err
  }

  err = libvirtd.DefineVm(updatexml, host)
  if err != nil {
    db.Tx_rollback(append(dblist, tx_Umountvmstatus, tx))
  }

  db.Tx_commot(append(dblist, tx, tx_Umountvmstatus))
  return nil
}

func Getdiskstatus(uuid string) (int, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return 0, err
  }

  vdisk := &Vms_vdisks{}
  errdb := dbs.Where("vdiskid=?", uuid).First(vdisk)
  if errdb.Error != nil {
    return 0, errdb.Error
  }
  return vdisk.Status, nil
}


func mapvdisk(obj  []Vms_vdisks) ([]map[string]interface{}, error)  {
  var mapc []map[string]interface{}

  for _, v := range obj {
    c := make(map[string]interface{})
    m := reflect.TypeOf(v)
    n := reflect.ValueOf(v)
    for i := 0; i < m.NumField(); i++ {
      c[m.Field(i).Name] = n.Field(i).Interface()
    }

    u, err := users.GetUserByUserID(v.User)
    if err != nil {
      return nil, err
    }

    c["username"] = u.Username
    mapc = append(mapc, c)
  }
  return mapc, nil
}

func Getvdisk(userid int) ([]map[string]interface{}, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var c  []Vms_vdisks

  user, err := users.GetUserByUserID(userid)
  if err != nil {
    return nil, err
  }

  role, err := users.GetRoleByRoleId(user.Role)
  if err != nil {
    return nil, err
  }

  if  role.Rolename == "admin" {
    dbs.Order("createtime desc").Find(&c)
  } else {
    dbs.Where("user=?", userid).Order("createtime desc").Find(&c)
  }
  mapvdisk, err := mapvdisk(c)
  if err != nil {
    return nil, err
  }
  return mapvdisk, err
}

func Umountvmstatus(datacenter string, storage string, vdiskid string) (*gorm.DB, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  tx := dbs.Begin()
  err = tx.Model(Vms_vdisks{}).Where("datacenter=? and storage=? and vdiskid=?", datacenter, storage, vdiskid).Update("vm_ip", "").Update("status", 1).Error
  if err != nil {
    tx.Rollback()
    return nil, err
  }
  return tx, nil
}

var Disknametype = []string{"vdb", "vdc", "vdd", "vde", "vdf"}
var slot = map[string] int {"vdb": 11, "vdc": 12, "vdd": 13, "vde": 14, "vdf": 15}

func next(items []*Vms_vdisks, item string) bool {
  for _, i := range items {
    c := i
    if c.Diskname == item {
      return true
    }
  }

  return false
}

func namedisk(vmip string) (string, error) {
  disklist, err := Getdiskbyvm(vmip)
  if err != nil {
    return "", err
  }

  for _, a := range Disknametype {
    b := next(disklist, a)
    if (b == false ) {
      return a, err
    }
  }
  return "vdb", err
}

func checkmount(vdiskid string) (int, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return 0, err
  }

  v := &Vms_vdisks{}
  dbs.Where("vdiskid=?", vdiskid).First(v)
  return v.Status, nil
}

func Mountdisk(ip string, vmhost string, storage string, pool string, datacenter string, vdiskid string, vms interface{}, xml string) error {
  checkmount, err := checkmount(vdiskid)
  if err != nil {
    return err
  }
  if checkmount == 0 {
    return vmerror.Error{Message: "vdisk has been mouunted"}
  }

  disknum, err := Getdiskbyvm(ip)
  if err != nil {
    return err
  }

  ceph := cephcommon.Vms_Ceph{}
  if len(disknum) >= 5 {
    return vmerror.Error{Message: "Maximum number of mounted to 5"}
  }
  storageinfo, err := ceph.Cephinfobyuuid(storage)
  if err != nil {
    return err
  }

  ips := strings.Split(storageinfo.Ips, ",")

  diskname, err := namedisk(ip)
  if err != nil {
    return err
  }

  docstring, err := libvirtd.CreateDiskXml(xml, vdiskid, ips, storageinfo.Port, pool, len(disknum), diskname, storageinfo.Ceph_secret)
  if err != nil {
    return err
  }

  tx_updatexml, err := updatexmlbyip(docstring, ip, vms)
  if err != nil {
    return err
  }

  dblist := []*gorm.DB{}
  tx_updatevm, err := UpdateMountvmstatus(datacenter, storage, vdiskid, ip, diskname)
  if err != nil {
    db.Tx_rollback(append(dblist, tx_updatexml))
    return err
  }

  err = libvirtd.DefineVm(docstring, vmhost)
  if err != nil {
    db.Tx_rollback(append(dblist, tx_updatexml, tx_updatevm))
    tx_updatexml.Rollback()
    tx_updatevm.Rollback()
    return err
  }

  db.Tx_commot(append(dblist, tx_updatevm, tx_updatexml))
  return nil
}

func updatexmlbyip(xml string, ip string, vms interface{}) (*gorm.DB, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  tx := dbs.Begin()
  err = tx.Model(vms).Where("ip=?", ip).Update("vmxml", xml).Error
  if err != nil {
    tx.Rollback()
    return nil, err
  }

  return tx, nil
}

var (
  config, _ = c.Exportconfig()
  offset = config.Page.Offset
)

func Mapvdiskarchive(obj []Vms_vdisks_archives) []map[string]interface{}  {
  var mapc []map[string]interface{}

  for _, v := range obj {
    c := make(map[string]interface{})
    m := reflect.TypeOf(v)
    n := reflect.ValueOf(v)
    for i := 0; i < m.NumField(); i++ {
      c[m.Field(i).Name] = n.Field(i).Interface()
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

func (vd Vms_vdisks_archives)GetVmArchive() ([]map[string]interface{}, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  d := []Vms_vdisks_archives{}
  dberr := dbs.Order("archive_time desc").Order("create_time desc").Find(&d)
  if dberr.Error != nil {
    return nil, err
  }

  return Mapvdiskarchive(d), nil
}


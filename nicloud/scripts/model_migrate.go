package main

import (
  "flag"
  "fmt"
  "nicloud/utils"
  "nicloud/vdisk"
  "nicloud/vm"
  "nicloud/dbs"
  "nicloud/cephcommon"
  "nicloud/datacenter"
  "nicloud/networks"
  "nicloud/osimage"
  "nicloud/users"
)

func main()  {
  var username string
  var passwd string
  flag.StringVar(&username, "username","", "")
  flag.StringVar(&passwd, "passwd","", "")
  flag.Parse()

  if username == "" || passwd == "" {
    fmt.Println("Usage: go run scripts/model_migrate.go --username admin  --passwd 123456")
    return
  }

  passwd = utils.Encryption(passwd)

  db, err := db.NicloudDb()
  if err != nil {
    return
  }

  vmobj := vm.Vms{}
  errdb := db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vmobj)
  if errdb.Error != nil {
    return
  }

  hostobj := vm.Vm_hosts{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&hostobj).Error
  if err != nil {
    return
  }

  vlanmaphost := vm.Vms_vlan_map_hosts{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vlanmaphost).Error
  if err != nil {
    return
  }

  vms_archives := vm.Vms_archives{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_archives).Error
  if err != nil {
    return
  }

  vms_cephs := cephcommon.Vms_Ceph{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_cephs).Error
  if err != nil {
    return
  }

  vms_snaps := cephcommon.Vms_snaps{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_snaps).Error
  if err != nil {
    return
  }

  vms_datacenter := datacenter.Vms_datacenter{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_datacenter).Error
  if err != nil {
    return
  }

  vms_netwoks:= networks.Vms_vlans{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_netwoks).Error
  if err != nil {
    return
  }

  vms_ips := networks.Vms_ips{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_ips).Error
  if err != nil {
    return
  }

  vms_os := osimage.Vms_os{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_os).Error
  if err != nil {
    return
  }

  vms_os_tags := osimage.Vms_os_tags{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_os_tags).Error
  if err != nil {
    return
  }

  tag_linux := osimage.Vms_os_tags{
    Id: 1,
    Tag: "LINUX",
  }

  db.Create(&tag_linux)

  tag_windows := osimage.Vms_os_tags{
    Id: 2,
    Tag: "WINDOWS",
  }

  db.Create(&tag_windows)


  vms_osimage_sort := osimage.Vms_osimage_sort{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_osimage_sort).Error
  if err != nil {
    return
  }

  os_base := osimage.Vms_osimage_sort{
    Id: 1,
    Sort: "基础镜像",
  }
  db.Create(&os_base)

  os_user := osimage.Vms_osimage_sort{
    Id: 2,
    Sort: "用户镜像",
  }
  db.Create(&os_user)

  vms_roles := users.Vms_roles{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_roles).Error
  if err != nil {
    return
  }

  admin := users.Vms_roles{
    Id: 1,
    Rolename: "admin",
  }
  db.Create(&admin)

  user := users.Vms_roles{
    Id: 2,
    Rolename: "user",
  }
  db.Create(&user)

  vms_users := users.Vms_users{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_users).Error
  if err != nil {
    return
  }

  u := users.Vms_users{
    Username: username,
    Passwd: passwd,
    Email: "",
    Role: 1,
    Mobile: "",
  }
  db.Create(&u)

  vms_vdisk := vdisk.Vms_vdisks{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_vdisk).Error
  if err != nil {
    return
  }

  vms_vdisks_archives := vdisk.Vms_vdisks_archives{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_vdisks_archives).Error
  if err != nil {
    return
  }

  vms_flavors := vm.Vm_flavors{}
  err = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_flavors).Error
  if err != nil {
    return
  }

  return
}


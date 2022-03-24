package main

import (
  "nicloud/vdisk"
  "nicloud/vm"
)
import "nicloud/dbs"
import "nicloud/cephcommon"
import "nicloud/datacenter"
import "nicloud/networks"
import "nicloud/osimage"
import "nicloud/users"


func main()  {
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
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&hostobj)
  if errdb.Error != nil {
    return
  }

  vlanmaphost := vm.Vms_vlan_map_hosts{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vlanmaphost)
  if errdb.Error != nil {
    return
  }

  vms_archives := vm.Vms_archive{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_archives)
  if errdb.Error != nil {
    return
  }

  vms_cephs := cephcommon.Vms_Ceph{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_cephs)
  if errdb.Error != nil {
    return
  }

  vms_snaps := cephcommon.Vms_snaps{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_snaps)
  if errdb.Error != nil {
    return
  }

  vms_datacenter := datacenter.Vms_datacenter{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_datacenter)
  if errdb.Error != nil {
    return
  }

  vms_netwoks:= networks.Vms_vlans{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_netwoks)
  if errdb.Error != nil {
    return
  }

  vms_ips := networks.Vms_ips{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_ips)
  if errdb.Error != nil {
    return
  }

  vms_os := osimage.Vms_os{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_os)
  if errdb.Error != nil {
    return
  }

  vms_osimage_sort := osimage.Vms_osimage_sort{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_osimage_sort)
  if errdb.Error != nil {
    return
  }

  vms_users := users.Vms_users{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_users)
  if errdb.Error != nil {
    return
  }

  vms_roles := users.Vms_roles{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_roles)
  if errdb.Error != nil {
    return
  }

  vms_vdisk := vdisk.Vms_vdisks{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_vdisk)
  if errdb.Error != nil {
    return
  }

  vms_vdisks_archives := vdisk.Vms_vdisks_archives{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_vdisks_archives)
  if errdb.Error != nil {
    return
  }

  vms_flavors := vm.Vm_flavors{}
  errdb = db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&vms_flavors)
  if errdb.Error != nil {
    return
  }

  return
}


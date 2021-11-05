package ceph

import (
  "github.com/ceph/go-ceph/rados"
  rbd "github.com/ceph/go-ceph/rbd"
  "goblog/dbs"
  "goblog/vmerror"
  "time"
)

type Vms_Ceph struct {
  Name string
  Pool string
  Datacenter string
  Ceph_secret string
  Ips string
  Port string
  Comment string
  Status int8
}


func Restore(vlan string, status int) error {
  var s int

  if status == 0 {
    s = 1
  } else {
    s = 0
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  dberr := dbs.Model(Vms_Ceph{}).Where("name=?", vlan).Update("status", s)
  if dberr.Error != nil {
    return dberr.Error
  }

  return nil
}

func Add(name string, pool string, datacenter string, ceph_secret string, ips string, port string, comment  string) error {
  c := &Vms_Ceph{
    Name: name,
    Pool: pool,
    Datacenter: datacenter,
    Ceph_secret: ceph_secret,
    Ips: ips,
    Port: port,
    Comment: comment,
  }
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  errdb := dbs.Create(c)
  if errdb.Error != nil {
    return errdb.Error
  }
  dbs.NewRecord(c)
  return nil
}

func Get()([]*Vms_Ceph, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  c := []*Vms_Ceph{}
  dbs.Find(&c)
  return c, nil
}

func Getpool(datacenter string, storage string)([]*Vms_Ceph, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  c := []*Vms_Ceph{}
  dbs.Where("datacenter=? and name=?", datacenter, storage).Find(&c)
  return c, nil
}

func Cephinfobyname(datacenter string, storage string)([]*Vms_Ceph, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  c := []*Vms_Ceph{}
  dbs.Where("datacenter=? and name=?", datacenter, storage).Find(&c)
  return c, nil
}


func CephConn() (*rados.Conn, error) {
  conn, err := rados.NewConn()
  if err != nil {
    return nil, err
  }
  err = conn.ReadDefaultConfigFile()
  if err != nil {
    return nil, err
  }
  err = conn.Connect()
  if err != nil {
    return nil, err
  }

  return conn, nil
}

func Rm_image(uuid string) (error) {
  conn, err := CephConn()
  if err != nil {
    return err
  }

  ioctx, err := conn.OpenIOContext("vm")
  if err != nil {
    return err
  }

  img := rbd.GetImage(ioctx, uuid)
  Archive_img := "x_"+(time.Now().Format("200601021504"))+uuid
  img.Rename(Archive_img)

  return nil
}

func RbdClone(id string) (string, error) {

  conn, err := CephConn()
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

type Vms_cloudrive struct {
  Cloudriveid string
  Contain string
  Pool string
  Storage string
  Datacenter string
  Vm_ip string
  User string
  Status int
}

func Add_cloudrive(contain string, pool string, storage string, datacenter string, user string) ([]*Vms_cloudrive, error) {
  cloudriveid := "123cnasdasdaweqwe"
  c := &Vms_cloudrive{
    Cloudriveid: cloudriveid,
    Contain: contain,
    Pool: pool,
    Storage: storage,
    Datacenter: datacenter,
    User: user,
    Status: 1,
  }

  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  errdb := dbs.Create(&c)
  if errdb.Error != nil {
    return nil, vmerror.Error{Message: errdb.Error.Error()}
  }
  return nil, err
}

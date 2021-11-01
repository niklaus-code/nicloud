package ceph

import (
  "github.com/ceph/go-ceph/rados"
  rbd "github.com/ceph/go-ceph/rbd"
  "goblog/dbs"
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
